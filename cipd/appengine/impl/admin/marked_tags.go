// Copyright 2018 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package admin

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/luci/appengine/mapper"
	"go.chromium.org/luci/cipd/appengine/impl/model"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/retry/transient"
)

// markedTag is a root entity created for each tag marked by some particular
// mapper operation.
//
// We can create many of these at arbitrary rate. Once the mapper finishes
// running and (eventually consistent) datastore indexes settle, we can query
// for all such marked tags using Eq() filter and deal with them in arbitrary
// order or at arbitrary rate.
//
// The core assumption here is that number of "interesting" tags found by
// a mapper is much less than total number of tags, but they may be clustered
// close together (so naively updating them inside the mapper causes transaction
// collisions).
type markedTag struct {
	_kind  string                `gae:"$kind,mapper.MarkedTag"`
	_extra datastore.PropertyMap `gae:"-,extra"`

	ID  string       `gae:"$id"` // see .genID()
	Job mapper.JobID // ID of a mapping job that produced it, for queries

	Key *datastore.Key `gae:",noindex"` // key of the corresponding Tag entity
	Tag string         `gae:",noindex"` // the original tag string (k:v)
	Why string         `gae:",noindex"` // why the tag was marked, for humans
}

// genID derives an ID for this markedTag entity.
//
// Called internally by visitAndMarkTags().
//
// Each unique tag marked by some particular job gets its own key, i.e. each
// job has its own namespace for marked tags.
func (t *markedTag) genID() {
	h := sha256.New()
	fmt.Fprintf(h, "%d\n%s", t.Job, t.Key.Encode())
	t.ID = base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

// queryMarkedTags returns a query for markedTags entities produced by a job.
func queryMarkedTags(job mapper.JobID) *datastore.Query {
	return datastore.NewQuery("mapper.MarkedTag").Eq("Job", job)
}

// visitAndMarkTags fetches Tag entities given their keys and passes them to
// the callback, which may choose to mark a tag by returning non-empty string
// with the human-readable reason why the tag was marked.
//
// Such marked tags are then stored in the datastore and later can be queried.
func visitAndMarkTags(c context.Context, job mapper.JobID, keys []*datastore.Key, cb func(*model.Tag) string) error {
	tags := make([]model.Tag, len(keys))
	for i, k := range keys {
		tags[i] = model.Tag{
			ID:       k.StringID(),
			Instance: k.Parent(),
		}
	}

	errAt := func(idx int) error { return nil }
	if err := datastore.Get(c, tags); err != nil {
		merr, ok := err.(errors.MultiError)
		if !ok {
			return errors.Annotate(err, "GetMulti RPC error when fetching %d tags", len(tags)).Tag(transient.Tag).Err()
		}
		errAt = func(idx int) error { return merr[idx] }
	}

	var marked []*markedTag
	for i, t := range tags {
		switch err := errAt(i); {
		case err == datastore.ErrNoSuchEntity:
			continue // just skip, no big deal
		case err != nil:
			return errors.Annotate(err, "failed to fetch tag entity with key %s", keys[i]).Err()
		}
		if why := cb(&t); why != "" {
			mt := &markedTag{
				Job: job,
				Key: keys[i],
				Tag: t.Tag,
				Why: why,
			}
			mt.genID()
			marked = append(marked, mt)
		}
	}

	if err := datastore.Put(c, marked); err != nil {
		return errors.Annotate(err, "failed to store %d markedTag(s)", len(marked)).Tag(transient.Tag).Err()
	}
	return nil
}