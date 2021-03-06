// Copyright 2020 The LUCI Authors.
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

// Package purger deletes expired test results from Spanner.
package purger

import (
	"context"
	"time"

	"cloud.google.com/go/spanner"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/server"

	"go.chromium.org/luci/resultdb/internal/cron"
	"go.chromium.org/luci/resultdb/internal/span"
)

// maxTestVariantsToFilter is the maximum number of test variants to include
// in the exclusion clause of the paritioned delete statement used to purge
// expired test results. Invocations that have more than this number of test
// variant combinations with unexpected results will not be purged, until
// the whole invocation expires.
const maxTestVariantsToFilter = 1000

// Options is purger server configuration.
type Options struct {
	// ForceCronInterval forces minimum interval in cron jobs.
	// Useful in integration tests to reduce the test time.
	ForceCronInterval time.Duration
}

// InitServer initializes a purger server.
func InitServer(srv *server.Server, opts Options) {
	srv.RunInBackground("resultdb.purge", func(ctx context.Context) {
		minInterval := time.Minute
		if opts.ForceCronInterval > 0 {
			minInterval = opts.ForceCronInterval
		}
		run(ctx, minInterval)
	})
}

// run continuously purges expired test results.
// It blocks until context is canceled.
func run(ctx context.Context, minInterval time.Duration) {
	maxShard, err := span.CurrentMaxShard(ctx)
	switch {
	case err == span.ErrNoResults:
		maxShard = span.InvocationShards - 1
	case err != nil:
		panic(errors.Annotate(err, "failed to determine number of shards").Err())
	}

	// Start one cron job for each shard of the database.
	cron.Group(ctx, maxShard+1, minInterval, purgeOneShard)
}

func purgeOneShard(ctx context.Context, shard int) error {
	st := spanner.NewStatement(`
		SELECT InvocationId
		FROM Invocations@{FORCE_INDEX=InvocationsByExpectedTestResultsExpiration}
		WHERE ShardId = @shardId
		AND ExpectedTestResultsExpirationTime IS NOT NULL
		AND ExpectedTestResultsExpirationTime <= CURRENT_TIMESTAMP()
	`)
	st.Params["shardId"] = shard
	return span.Query(ctx, span.Client(ctx).Single(), st, func(row *spanner.Row) error {
		var id span.InvocationID
		if err := span.FromSpanner(row, &id); err != nil {
			return err
		}

		if err := purgeOneInvocation(ctx, id); err != nil {
			logging.Errorf(ctx, "failed to process %s: %s", id, err)
		}
		return nil
	})
}

func purgeOneInvocation(ctx context.Context, invID span.InvocationID) error {
	txn := span.Client(ctx).ReadOnlyTransaction()
	defer txn.Close()

	// Check that invocation hasn't been purged already.
	var expirationTime spanner.NullTime
	ptrs := map[string]interface{}{"ExpectedTestResultsExpirationTime": &expirationTime}
	if err := span.ReadInvocation(ctx, txn, invID, ptrs); err != nil {
		return err
	}
	if expirationTime.IsNull() {
		// Invocation was purged by other worker.
		return nil
	}

	// Stream test results that need to be purged, and set Purge=true on them,
	// in batches.
	// Note that we cannot use Partitioned UPDATE here because its time complexity
	// is currently O(table size).
	var ms []*spanner.Mutation
	count := 0
	err := testResultsToPurge(ctx, txn, invID, func(testID, resultID string) error {
		count++
		ms = append(ms, spanner.Delete("TestResults", invID.Key(testID, resultID)))
		// Flush if the batch is too large.
		// Cloud Spanner limitation is 20k mutations per txn.
		// One deletion is one mutation.
		// Flush at 19k boundary.
		if len(ms) > 19000 {
			if _, err := span.Client(ctx).Apply(ctx, ms); err != nil {
				return err
			}
			span.IncRowCount(ctx, len(ms), span.TestResults, span.Deleted)
			ms = ms[:0]
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Flush the last batch.
	if len(ms) > 0 {
		if _, err := span.Client(ctx).Apply(ctx, ms); err != nil {
			return err
		}
		span.IncRowCount(ctx, len(ms), span.TestResults, span.Deleted)
	}

	// Set the invocation's result expiration to null.
	if err := unsetInvocationResultsExpiration(ctx, invID); err != nil {
		return err
	}

	logging.Debugf(ctx, "Deleted %d test results in %s", count, invID.Name())
	return nil
}

// testResultsToPurge calls f for test results that should be purged.
func testResultsToPurge(ctx context.Context, txn *spanner.ReadOnlyTransaction, inv span.InvocationID, f func(testID, resultID string) error) error {
	st := spanner.NewStatement(`
		WITH DoNotPurge AS (
			SELECT DISTINCT TestId, VariantHash
			FROM TestResults@{FORCE_INDEX=UnexpectedTestResults}
			WHERE InvocationId = @invocationId
			  AND IsUnexpected = TRUE
		)
		SELECT tr.TestId, tr.ResultId
		FROM TestResults tr
		LEFT JOIN DoNotPurge dnp ON tr.TestId = dnp.TestId AND tr.VariantHash = dnp.VariantHash
		WHERE InvocationId = @invocationId
			AND dnp.VariantHash IS NULL
	`)

	st.Params["invocationId"] = inv
	return span.Query(ctx, txn, st, func(row *spanner.Row) error {
		var testID, resultID string
		if err := row.Columns(&testID, &resultID); err != nil {
			return err
		}
		return f(testID, resultID)
	})
}

func unsetInvocationResultsExpiration(ctx context.Context, id span.InvocationID) error {
	_, err := span.Client(ctx).Apply(ctx, []*spanner.Mutation{
		span.UpdateMap("Invocations", map[string]interface{}{
			"InvocationID":                      id,
			"ExpectedTestResultsExpirationTime": nil,
		}),
	})
	if err != nil {
		return err
	}
	return nil
}
