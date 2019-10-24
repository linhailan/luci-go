// Copyright 2019 The LUCI Authors.
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

package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/golang/protobuf/ptypes"
	tspb "github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/grpcutil"
	"go.chromium.org/luci/grpc/prpc"
	"go.chromium.org/luci/server/auth"

	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"
)

const (
	// createInvocationGroup is a CIA group that can create invocations.
	// TODO(crbug.com/1013316): remove in favor of realms.
	createInvocationGroup = "luci-resultdb-access"
)

// validateInvocationDeadline returns a non-nil error if deadline is invalid.
func validateInvocationDeadline(deadline *tspb.Timestamp, now time.Time) error {
	switch deadline, err := ptypes.Timestamp(deadline); {
	case err != nil:
		return err

	case deadline.Sub(now) < 10*time.Second:
		return errors.Reason("must be at least 10 seconds in the future").Err()

	case deadline.Sub(now) > 2*24*time.Hour:
		return errors.Reason("must be before 48h in the future").Err()

	default:
		return nil
	}
}

// validateCreateInvocationRequest returns an error if req is determined to be
// invalid.
func validateCreateInvocationRequest(req *pb.CreateInvocationRequest, now time.Time) error {
	if err := pbutil.ValidateInvocationID(req.InvocationId); err != nil {
		return errors.Annotate(err, "invocation_id").Err()
	}

	if err := pbutil.ValidateRequestID(req.RequestId); err != nil {
		return errors.Annotate(err, "request_id").Err()
	}

	inv := req.Invocation
	if inv == nil {
		return nil
	}

	if err := pbutil.ValidateStringPairs(inv.GetTags()); err != nil {
		return errors.Annotate(err, "invocation.tags").Err()
	}

	if inv.GetDeadline() != nil {
		if err := validateInvocationDeadline(inv.Deadline, now); err != nil {
			return errors.Annotate(err, "invocation: deadline").Err()
		}
	}

	if inv.GetBaseTestVariantDef() != nil {
		if err := pbutil.ValidateVariantDef(inv.BaseTestVariantDef); err != nil {
			return errors.Annotate(err, "invocation.base_test_variant_def").Err()
		}
	}

	return nil
}

// CreateInvocation implements pb.RecorderServer.
func (s *recorderServer) CreateInvocation(ctx context.Context, in *pb.CreateInvocationRequest) (*pb.Invocation, error) {
	now := clock.Now(ctx)

	if err := mayCreateInvocation(ctx); err != nil {
		return nil, errors.Annotate(err, "").Err()
	}
	if err := validateCreateInvocationRequest(in, now); err != nil {
		return nil, errors.Annotate(err, "bad request").Tag(grpcutil.InvalidArgumentTag).Err()
	}

	// Return update token to the client.
	updateToken, err := generateUpdateToken()
	if err != nil {
		return nil, err
	}
	prpc.SetHeader(ctx, metadata.Pairs(updateTokenMetadataKey, updateToken))

	// Prepare the invocation we will return.
	inv := &pb.Invocation{
		Name:               pbutil.InvocationName(in.InvocationId),
		State:              pb.Invocation_ACTIVE,
		CreateTime:         pbutil.MustTimestampProto(now),
		Deadline:           in.Invocation.GetDeadline(),
		BaseTestVariantDef: in.Invocation.GetBaseTestVariantDef(),
		Tags:               in.Invocation.GetTags(),
	}

	// Determine the deadline and expiration times.
	if inv.Deadline == nil {
		inv.Deadline = pbutil.MustTimestampProto(now.Add(defaultInvocationDeadlineDuration))
	}

	pbutil.NormalizeInvocation(inv)

	_, err = span.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		// Dedup the request if possible.
		if in.RequestId != "" {
			var curRequestID spanner.NullString
			err := span.ReadInvocation(ctx, txn, in.InvocationId, map[string]interface{}{
				"CreateRequestId": &curRequestID,
			})
			switch {
			case grpcutil.Code(err) == codes.NotFound:
				// Continue to creation.

			case err != nil:
				return err

			case curRequestID.Valid && curRequestID.StringVal == in.RequestId:
				// Dedup the request.
				inv, err = span.ReadInvocationFull(ctx, txn, in.InvocationId)
				return err

			default:
				return invocationAlreadyExists()
			}
		}

		muts := insertInvocationsByTag(in.InvocationId, inv)
		muts = append(muts, insertInvocation(ctx, inv, updateToken, in.RequestId))

		return txn.BufferWrite(muts)
	})

	switch {
	case spanner.ErrCode(err) == codes.AlreadyExists:
		return nil, invocationAlreadyExists()
	case err != nil:
		return nil, err
	default:
		return inv, nil
	}
}

func invocationAlreadyExists() error {
	return errors.Reason("invocation already exists").Tag(grpcutil.AlreadyExistsTag).Err()
}

func mayCreateInvocation(ctx context.Context) error {
	// TODO(crbug.com/1013316): use realms.
	switch allowed, err := auth.IsMember(ctx, createInvocationGroup); {
	case err != nil:
		return err
	case !allowed:
		return errors.
			Reason("%s is not allowed to create invocations", auth.CurrentIdentity(ctx)).
			Tag(grpcutil.PermissionDeniedTag).
			Err()
	default:
		return nil
	}
}

func generateUpdateToken() (string, error) {
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	return hex.EncodeToString(buf), err
}