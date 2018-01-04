package gitiles

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/proto/git"
	"go.chromium.org/luci/common/proto/gitiles"
)

// DefaultLimit is the default maximum number of commits to load.
// It is used in PagingLog.
const DefaultLimit = 1000

// Helper functions for Gitiles.Log RPC.

// PagingLog is a wrapper around Gitiles.Log RPC that pages though commits.
// If req.PageToken is not empty, paging will continue from there.
//
// req.PageSize specifies maximum number of commits to load in each page.
//
// Limit specifies the maximum number of commits to load.
// 0 means use DefaultLimit.
func PagingLog(ctx context.Context, client gitiles.GitilesClient, req gitiles.LogRequest, limit int, opts ...grpc.CallOption) ([]*git.Commit, error) {
	// Note: we intentionally receive req as struct (not pointer)
	// because we need to mutate it.

	switch {
	case limit < 0:
		return nil, errors.New("limit must not be negative")
	case limit == 0:
		limit = DefaultLimit
	}

	var combinedLog []*git.Commit
	for {
		remaining := limit - len(combinedLog)
		if remaining <= 0 {
			break
		}
		if req.PageSize == 0 || remaining < int(req.PageSize) {
			// Do not fetch more than we need.
			req.PageSize = int32(remaining)
		}

		res, err := client.Log(ctx, &req, opts...)
		if err != nil {
			return combinedLog, err
		}

		// req was capped, so this should not exceed limit.
		combinedLog = append(combinedLog, res.Log...)
		req.PageToken = res.NextPageToken
	}
	return combinedLog, nil
}
