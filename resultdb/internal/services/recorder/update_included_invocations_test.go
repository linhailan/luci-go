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

package recorder

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"go.chromium.org/luci/resultdb/internal/span"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
	. "go.chromium.org/luci/resultdb/internal/testutil"
)

func TestValidateUpdateIncludedInvocationsRequest(t *testing.T) {
	t.Parallel()
	Convey(`TestValidateUpdateIncludedInvocationsRequest`, t, func() {
		Convey(`Valid`, func() {
			err := validateUpdateIncludedInvocationsRequest(&pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "invocations/a",
				AddInvocations:      []string{"invocations/b"},
				RemoveInvocations:   []string{"invocations/c"},
			})
			So(err, ShouldBeNil)
		})

		Convey(`Invalid including_invocation`, func() {
			err := validateUpdateIncludedInvocationsRequest(&pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "x",
				AddInvocations:      []string{"invocations/b"},
				RemoveInvocations:   []string{"invocations/c"},
			})
			So(err, ShouldErrLike, `including_invocation: does not match`)
		})
		Convey(`Invalid add_invocations`, func() {
			err := validateUpdateIncludedInvocationsRequest(&pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "invocations/a",
				AddInvocations:      []string{"x"},
				RemoveInvocations:   []string{"invocations/c"},
			})
			So(err, ShouldErrLike, `add_invocations: "x": does not match`)
		})
		Convey(`Invalid remove_invocations`, func() {
			err := validateUpdateIncludedInvocationsRequest(&pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "invocations/a",
				AddInvocations:      []string{"invocations/b"},
				RemoveInvocations:   []string{"x"},
			})
			So(err, ShouldErrLike, `remove_invocations: "x": does not match`)
		})
		Convey(`Include itself`, func() {
			err := validateUpdateIncludedInvocationsRequest(&pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "invocations/a",
				AddInvocations:      []string{"invocations/a"},
				RemoveInvocations:   []string{"invocations/c"},
			})
			So(err, ShouldErrLike, `cannot include itself`)
		})
		Convey(`Add and remove same invocation`, func() {
			err := validateUpdateIncludedInvocationsRequest(&pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "invocations/a",
				AddInvocations:      []string{"invocations/b"},
				RemoveInvocations:   []string{"invocations/b"},
			})
			So(err, ShouldErrLike, `cannot add and remove the same invocation(s) at the same time: ["invocations/b"]`)
		})
	})
}

func TestUpdateIncludedInvocations(t *testing.T) {
	Convey(`TestIncludedInvocations`, t, func() {
		ctx := SpannerTestContext(t)
		recorder := newTestRecorderServer()

		const token = "update token"
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(UpdateTokenMetadataKey, token))

		insInv := InsertInvocation

		assertIncluded := func(includedInvID span.InvocationID) {
			var throwAway span.InvocationID
			MustReadRow(ctx, "IncludedInvocations", span.InclusionKey("including", includedInvID), map[string]interface{}{
				"IncludedInvocationID": &throwAway,
			})
		}
		assertNotIncluded := func(includedInvID span.InvocationID) {
			var throwAway span.InvocationID
			MustNotFindRow(ctx, "IncludedInvocations", span.InclusionKey("including", includedInvID), map[string]interface{}{
				"IncludedInvocationID": &throwAway,
			})
		}

		Convey(`Invalid request`, func() {
			_, err := recorder.UpdateIncludedInvocations(ctx, &pb.UpdateIncludedInvocationsRequest{})
			So(err, ShouldHaveAppStatus, codes.InvalidArgument, `bad request: including_invocation: unspecified`)
		})

		Convey(`With valid request`, func() {
			req := &pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "invocations/including",
				AddInvocations: []string{
					"invocations/included",
					"invocations/included2",
				},
				RemoveInvocations: []string{
					"invocations/toberemoved",
					"invocations/neverexisted",
				},
			}

			Convey(`No including invocation`, func() {
				_, err := recorder.UpdateIncludedInvocations(ctx, req)
				So(err, ShouldHaveAppStatus, codes.NotFound, `invocations/including not found`)
			})

			Convey(`With existing inclusion`, func() {
				MustApply(ctx,
					insInv("including", pb.Invocation_ACTIVE, map[string]interface{}{"UpdateToken": token}),
					insInv("toberemoved", pb.Invocation_FINALIZED, nil),
				)
				_, err := recorder.UpdateIncludedInvocations(ctx, &pb.UpdateIncludedInvocationsRequest{
					IncludingInvocation: "invocations/including",
					AddInvocations:      []string{"invocations/toberemoved"},
				})
				So(err, ShouldBeNil)
				assertIncluded("toberemoved")

				Convey(`No included invocation`, func() {
					_, err := recorder.UpdateIncludedInvocations(ctx, req)
					So(err, ShouldHaveAppStatus, codes.NotFound, `one of the included invocations does not exist`)
				})

				Convey(`Success - idempotent`, func() {
					MustApply(ctx,
						insInv("included", pb.Invocation_FINALIZED, nil),
						insInv("included2", pb.Invocation_FINALIZED, nil),
					)

					_, err := recorder.UpdateIncludedInvocations(ctx, req)
					So(err, ShouldBeNil)
					assertIncluded("included")
					assertIncluded("included2")
					assertNotIncluded("toberemoved")

					_, err = recorder.UpdateIncludedInvocations(ctx, req)
					So(err, ShouldBeNil)
					assertIncluded("included")
					assertIncluded("included2")
					assertNotIncluded("toberemoved")
				})
			})
		})
	})
}
