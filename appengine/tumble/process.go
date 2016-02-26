// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package tumble

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/luci/gae/filter/txnBuf"
	"github.com/luci/gae/service/datastore"
	"github.com/luci/gae/service/datastore/serialize"
	"github.com/luci/gae/service/memcache"
	"github.com/luci/luci-go/appengine/memlock"
	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/common/errors"
	"github.com/luci/luci-go/common/logging"
	"github.com/luci/luci-go/common/parallel"
	"github.com/luci/luci-go/common/stringset"
	"golang.org/x/net/context"
)

// expandedShardBounds returns the boundary of the expandedShard order that
// currently corresponds to this shard number. If Shard is < 0 or > NumShards
// (the currently configured number of shards), this will return a low > high.
// Otherwise low < high.
func expandedShardBounds(c context.Context, shard uint64) (low, high int64) {
	cfg := GetConfig(c)

	if shard < 0 || uint64(shard) >= cfg.NumShards {
		logging.Warningf(c, "Invalid shard: %d", shard)
		// return inverted bounds
		return 0, -1
	}

	expandedShardsPerShard := int64(math.MaxUint64 / cfg.NumShards)
	low = math.MinInt64 + (int64(shard) * expandedShardsPerShard)
	if uint64(shard) == cfg.NumShards-1 {
		high = math.MaxInt64
	} else {
		high = low + expandedShardsPerShard
	}
	return
}

var dustSettleTimeout = 2 * time.Second

// ProcessShardHandler is a http handler suitable for installation into
// a httprouter. It expects `logging` and `luci/gae` services to be installed
// into the context.
//
// ProcessShardHandler verifies that its being run as a taskqueue task and that
// the following parameters exist and are well-formed:
//   * timestamp: decimal-encoded UNIX/UTC timestamp in seconds.
//   * shard_id: decimal-encoded shard identifier.
//
// ProcessShardHandler then invokes ProcessShard with the parsed parameters.
func ProcessShardHandler(c context.Context, rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tstampStr := p.ByName("timestamp")
	sidStr := p.ByName("shard_id")

	tstamp, err := strconv.ParseInt(tstampStr, 10, 64)
	if err != nil {
		logging.Errorf(c, "bad timestamp %q", tstampStr)
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "bad timestamp")
		return
	}

	sid, err := strconv.ParseUint(sidStr, 10, 64)
	if err != nil {
		logging.Errorf(c, "bad shardID %q", tstampStr)
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "bad shardID")
		return
	}

	err = ProcessShard(c, time.Unix(tstamp, 0).UTC(), sid)
	if err != nil {
		logging.Errorf(c, "failure! %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "error: %s", err)
	} else {
		rw.Write([]byte("ok"))
	}
}

// ProcessShard is the tumble backend endpoint. This accepts a shard number which
// is expected to be < GlobalConfig.NumShards.
func ProcessShard(c context.Context, timestamp time.Time, shard uint64) error {
	low, high := expandedShardBounds(c, shard)
	if low > high {
		return nil
	}

	l := logging.Get(logging.SetField(c, "shard", shard))

	cfg := GetConfig(c)

	lockKey := fmt.Sprintf("%s.%d.lock", cfg.Name, shard)
	clientID := fmt.Sprintf("%d_%d", timestamp.Unix(), shard)

	// this last key allows buffered tasks to early exit if some other shard
	// processor has already processed past this task's target timestamp.
	lastKey := fmt.Sprintf("%s.%d.last", cfg.Name, shard)
	mc := memcache.Get(c)
	lastItm, err := mc.Get(lastKey)
	if err != nil {
		if err != memcache.ErrCacheMiss {
			l.Warningf("couldn't obtain last timestamp: %s", err)
		}
	} else {
		val := lastItm.Value()
		last, err := serialize.ReadTime(bytes.NewBuffer(val))
		if err != nil {
			l.Warningf("could not decode timestamp %v: %s", val, err)
		} else {
			last = last.Add(time.Duration(cfg.TemporalRoundFactor))
			if last.After(timestamp) {
				l.Infof("early exit, %s > %s", last, timestamp)
				return nil
			}
		}
	}
	err = nil

	q := datastore.NewQuery("tumble.Mutation").
		Gte("ExpandedShard", low).Lte("ExpandedShard", high).
		Project("TargetRoot").Distinct(true).
		Limit(cfg.ProcessMaxBatchSize)

	banSets := map[string]stringset.Set{}

	for try := 0; try < 2; try++ {
		err = memlock.TryWithLock(c, lockKey, clientID, func(c context.Context) error {
			l.Infof("Got lock (try %d)", try)

			for {
				processCounters := []*int64{}
				err := parallel.WorkPool(int(cfg.NumGoroutines), func(ch chan<- func() error) {
					err := datastore.Get(c).Run(q, func(pm datastore.PropertyMap) error {
						root := pm["TargetRoot"][0].Value().(*datastore.Key)
						encRoot := root.Encode()

						// TODO(riannucci): make banSets remove keys from the banSet which
						// weren't hit. Once they stop showing up, they'll never show up
						// again.

						bs := banSets[encRoot]
						if bs == nil {
							bs = stringset.New(0)
							banSets[encRoot] = bs
						}
						counter := new(int64)
						processCounters = append(processCounters, counter)

						ch <- func() error {
							return processRoot(c, root, bs, counter)
						}

						if c.Err() != nil {
							l.Warningf("Lost lock! %s", c.Err())
							return datastore.Stop
						}
						return nil
					})
					if err != nil {
						l.Errorf("Failure to query: %s", err)
						ch <- func() error {
							return err
						}
					}
				})
				if err != nil {
					return err
				}
				numProcessed := int64(0)
				for _, n := range processCounters {
					numProcessed += *n
				}
				l.Infof("cumulatively processed %d items", numProcessed)
				if numProcessed == 0 {
					break
				}

				err = mc.Set(mc.NewItem(lastKey).SetValue(serialize.ToBytes(clock.Now(c).UTC())))
				if err != nil {
					l.Warningf("could not update last process memcache key %s: %s", lastKey, err)
				}

				if tr := clock.Sleep(c, dustSettleTimeout); tr.Incomplete() {
					l.Warningf("sleep interrupted, context is done: %v", tr.Err)
					return tr.Err
				}
			}
			return nil
		})
		if err != memlock.ErrFailedToLock {
			break
		}
		l.Infof("Couldn't obtain lock (try %d) (sleeping 2s)", try+1)
		if tr := clock.Sleep(c, time.Second*2); tr.Incomplete() {
			l.Warningf("sleep interrupted, context is done: %v", tr.Err)
			return tr.Err
		}
	}
	if err == memlock.ErrFailedToLock {
		l.Infof("Couldn't obtain lock (giving up): %s", err)
		err = nil
	}
	return err
}

func getBatchByRoot(c context.Context, root *datastore.Key, banSet stringset.Set) ([]*realMutation, error) {
	cfg := GetConfig(c)
	ds := datastore.Get(c)
	q := datastore.NewQuery("tumble.Mutation").Eq("TargetRoot", root)
	if cfg.DelayedMutations {
		q = q.Lte("ProcessAfter", clock.Now(c).UTC())
	}

	toFetch := make([]*realMutation, 0, cfg.ProcessMaxBatchSize)
	err := ds.Run(q, func(k *datastore.Key) error {
		if !banSet.Has(k.Encode()) {
			toFetch = append(toFetch, &realMutation{
				ID:     k.StringID(),
				Parent: k.Parent(),
			})
		}
		if len(toFetch) < cap(toFetch) {
			return nil
		}
		return datastore.Stop
	})
	return toFetch, err
}

func loadFilteredMutations(c context.Context, rms []*realMutation) ([]*datastore.Key, []Mutation, error) {
	ds := datastore.Get(c)

	mutKeys := make([]*datastore.Key, 0, len(rms))
	muts := make([]Mutation, 0, len(rms))
	err := ds.GetMulti(rms)
	me, ok := err.(errors.MultiError)
	if !ok && err != nil {
		return nil, nil, err
	}

	for i, rm := range rms {
		err = nil
		if me != nil {
			err = me[i]
		}
		if err == nil {
			if rm.Version != getAppVersion(c) {
				logging.Fields{
					"mut_version": rm.Version,
					"cur_version": getAppVersion(c),
				}.Warningf(c, "loading mutation with different code version")
			}
			m, err := rm.GetMutation()
			if err != nil {
				logging.Errorf(c, "couldn't load mutation: %s", err)
				continue
			}
			muts = append(muts, m)
			mutKeys = append(mutKeys, ds.KeyForObj(rm))
		} else if err != datastore.ErrNoSuchEntity {
			return nil, nil, me
		}
	}

	return mutKeys, muts, nil
}

type overrideRoot struct {
	Mutation

	root *datastore.Key
}

func (o overrideRoot) Root(context.Context) *datastore.Key {
	return o.root
}

func processRoot(c context.Context, root *datastore.Key, banSet stringset.Set, counter *int64) error {
	cfg := GetConfig(c)
	l := logging.Get(c)

	toFetch, err := getBatchByRoot(c, root, banSet)
	if err != nil || len(toFetch) == 0 {
		return err
	}

	mutKeys, muts, err := loadFilteredMutations(c, toFetch)
	if err != nil {
		return err
	}

	if c.Err() != nil {
		l.Warningf("Lost lock during processRoot")
		return nil
	}

	allShards := map[taskShard]struct{}{}

	toDel := make([]*datastore.Key, 0, len(muts))
	numMuts := uint64(0)
	deletedMuts := uint64(0)
	processedMuts := uint64(0)
	err = datastore.Get(txnBuf.FilterRDS(c)).RunInTransaction(func(c context.Context) error {
		toDel = toDel[:0]
		numMuts = 0
		deletedMuts = 0
		processedMuts = 0

		iterMuts := muts
		iterMutKeys := mutKeys

		for i := 0; i < len(iterMuts); i++ {
			m := iterMuts[i]

			shards, newMuts, newMutKeys, err := enterTransactionInternal(c, overrideRoot{m, root}, uint64(i))
			if err != nil {
				l.Errorf("Executing decoded gob(%T) failed: %q: %+v", m, err, m)
				continue
			}
			processedMuts++
			for j, nm := range newMuts {
				if nm.Root(c).HasAncestor(root) {
					runNow := !cfg.DelayedMutations
					if !runNow {
						dm, isDelayedMutation := nm.(DelayedMutation)
						runNow = !isDelayedMutation || clock.Now(c).UTC().After(dm.ProcessAfter())
					}
					if runNow {
						iterMuts = append(iterMuts, nm)
						iterMutKeys = append(iterMutKeys, newMutKeys[j])
					}
				}
			}

			key := iterMutKeys[i]
			if key.HasAncestor(root) {
				// try to delete it as part of the same transaction.
				if err := datastore.Get(c).Delete(key); err == nil {
					deletedMuts++
				} else {
					toDel = append(toDel, key)
				}
			} else {
				toDel = append(toDel, key)
			}

			numMuts += uint64(len(newMuts))
			for shard := range shards {
				allShards[shard] = struct{}{}
			}
		}

		return nil
	}, nil)
	if err != nil {
		l.Errorf("failed running transaction: %s", err)
		return err
	}
	numMuts -= deletedMuts

	fireTasks(c, allShards)
	l.Infof("successfully processed %d mutations (%d tail-call), adding %d more", processedMuts, deletedMuts, numMuts)

	if len(toDel) > 0 {
		atomic.StoreInt64(counter, int64(len(toDel)))

		for _, k := range toDel {
			banSet.Add(k.Encode())
		}
		if err := datastore.Get(c).DeleteMulti(toDel); err != nil {
			l.Warningf("error deleting finished mutations: %s", err)
		}
	}

	return nil
}
