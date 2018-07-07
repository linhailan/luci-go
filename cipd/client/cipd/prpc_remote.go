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

package cipd

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/proto/google"
	"go.chromium.org/luci/common/retry"
	"go.chromium.org/luci/grpc/prpc"

	api "go.chromium.org/luci/cipd/api/cipd/v1"
	"go.chromium.org/luci/cipd/common"
)

// gRPC errors that may be returned by api.RepositoryClient that we recognize
// and handle ourselves. They will not be logged by the pRPC library.
var expectedCodes = prpc.ExpectedCode(
	codes.Aborted,
	codes.AlreadyExists,
	codes.FailedPrecondition,
	codes.NotFound,
	codes.PermissionDenied,
)

// humanErr takes gRPC errors and returns a human readable error that can be
// presented in the CLI.
//
// It basically strips scary looking gRPC framing around the error message.
func humanErr(err error) error {
	if err != nil {
		if status, ok := status.FromError(err); ok {
			return errors.New(status.Message())
		}
	}
	return err
}

// prpcRemoteImpl implements v1 'remote' interface using v2 protocol.
//
// It exists temporarily during the transition from v1 to v2 protocol. Once
// v2 is fully implemented and becomes the default, v1 implementation of the
// 'remote' interface will be deleted, and the interface itself will be adjusted
// to  match v2 protocol better (or deleted completely, since pRPC-level
// interface is good enough on its own).
type prpcRemoteImpl struct {
	serviceURL string
	userAgent  string
	client     *http.Client

	cas  api.StorageClient
	repo api.RepositoryClient
}

func (r *prpcRemoteImpl) init() error {
	// Note: serviceURL is ALWAYS "http(s)://<host>" here, as setup by NewClient.
	parsed, err := url.Parse(r.serviceURL)
	if err != nil {
		panic(err)
	}
	prpcC := &prpc.Client{
		C:    r.client,
		Host: parsed.Host,
		Options: &prpc.Options{
			UserAgent: r.userAgent,
			Insecure:  parsed.Scheme == "http", // for testing with local dev server
			Retry: func() retry.Iterator {
				return &retry.ExponentialBackoff{
					Limited: retry.Limited{
						Delay:   time.Second,
						Retries: 10,
					},
				}
			},
		},
	}

	r.cas = api.NewStoragePRPCClient(prpcC)
	r.repo = api.NewRepositoryPRPCClient(prpcC)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// ACLs.

func (r *prpcRemoteImpl) fetchACL(ctx context.Context, packagePath string) ([]PackageACL, error) {
	resp, err := r.repo.GetInheritedPrefixMetadata(ctx, &api.PrefixRequest{
		Prefix: packagePath,
	}, expectedCodes)
	if err != nil {
		return nil, humanErr(err)
	}
	return prefixMetadataToACLs(resp), nil
}

func (r *prpcRemoteImpl) modifyACL(ctx context.Context, packagePath string, changes []PackageACLChange) error {
	// Fetch existing metadata, if any.
	meta, err := r.repo.GetPrefixMetadata(ctx, &api.PrefixRequest{
		Prefix: packagePath,
	}, expectedCodes)
	if code := grpc.Code(err); code != codes.OK && code != codes.NotFound {
		return humanErr(err)
	}

	// Construct new empty metadata for codes.NotFound.
	if meta == nil {
		meta = &api.PrefixMetadata{Prefix: packagePath}
	}

	// Apply mutations.
	if dirty, err := mutateACLs(meta, changes); !dirty || err != nil {
		return err
	}

	// Store the new metadata. This call will check meta.Fingerprint.
	_, err = r.repo.UpdatePrefixMetadata(ctx, meta, expectedCodes)
	return humanErr(err)
}

func (r *prpcRemoteImpl) fetchRoles(ctx context.Context, packagePath string) ([]string, error) {
	resp, err := r.repo.GetRolesInPrefix(ctx, &api.PrefixRequest{
		Prefix: packagePath,
	}, expectedCodes)
	if err != nil {
		return nil, humanErr(err)
	}
	out := make([]string, len(resp.Roles))
	for i, r := range resp.Roles {
		out[i] = r.Role.String()
	}
	return out, nil
}

////////////////////////////////////////////////////////////////////////////////
// Upload.

func (r *prpcRemoteImpl) initiateUpload(ctx context.Context, sha1 string) (*UploadSession, error) {
	op, err := r.cas.BeginUpload(ctx, &api.BeginUploadRequest{
		Object: &api.ObjectRef{
			HashAlgo:  api.HashAlgo_SHA1,
			HexDigest: sha1,
		},
	}, prpc.ExpectedCode(codes.AlreadyExists))
	switch grpc.Code(err) {
	case codes.OK:
		return &UploadSession{op.OperationId, op.UploadUrl}, nil
	case codes.AlreadyExists:
		return nil, nil
	default:
		return nil, err
	}
}

func (r *prpcRemoteImpl) finalizeUpload(ctx context.Context, sessionID string) (bool, error) {
	op, err := r.cas.FinishUpload(ctx, &api.FinishUploadRequest{
		UploadOperationId: sessionID,
	})
	if err != nil {
		return false, err
	}
	switch op.Status {
	case api.UploadStatus_UPLOADING, api.UploadStatus_VERIFYING:
		return false, nil // still verifying
	case api.UploadStatus_PUBLISHED:
		return true, nil // verified!
	case api.UploadStatus_ERRORED:
		return false, errors.New(op.ErrorMessage)
	default:
		return false, fmt.Errorf("unrecognized upload operation status %s", op.Status)
	}
}

func (r *prpcRemoteImpl) registerInstance(ctx context.Context, pin common.Pin) (*registerInstanceResponse, error) {
	resp, err := r.repo.RegisterInstance(ctx, &api.Instance{
		Package: pin.PackageName,
		Instance: &api.ObjectRef{
			HashAlgo:  api.HashAlgo_SHA1,
			HexDigest: pin.InstanceID,
		},
	}, expectedCodes)
	if err != nil {
		return nil, humanErr(err)
	}
	switch resp.Status {
	case api.RegistrationStatus_REGISTERED, api.RegistrationStatus_ALREADY_REGISTERED:
		return &registerInstanceResponse{
			alreadyRegistered: resp.Status == api.RegistrationStatus_ALREADY_REGISTERED,
			registeredBy:      resp.Instance.RegisteredBy,
			registeredTs:      google.TimeFromProto(resp.Instance.RegisteredTs),
		}, nil
	case api.RegistrationStatus_NOT_UPLOADED:
		return &registerInstanceResponse{
			uploadSession: &UploadSession{resp.UploadOp.OperationId, resp.UploadOp.UploadUrl},
		}, nil
	default:
		return nil, fmt.Errorf("unrecognized package registration status %s", resp.Status)
	}
}

////////////////////////////////////////////////////////////////////////////////
// Fetching.

func (r *prpcRemoteImpl) resolveVersion(ctx context.Context, packageName, version string) (common.Pin, error) {
	resp, err := r.repo.ResolveVersion(ctx, &api.ResolveVersionRequest{
		Package: packageName,
		Version: version,
	}, expectedCodes)
	if err != nil {
		return common.Pin{}, humanErr(err)
	}
	return common.Pin{
		PackageName: packageName,
		InstanceID:  common.ObjectRefToInstanceID(resp.Instance),
	}, nil
}

func (r *prpcRemoteImpl) fetchPackageRefs(ctx context.Context, packageName string) ([]RefInfo, error) {
	resp, err := r.repo.ListRefs(ctx, &api.ListRefsRequest{
		Package: packageName,
	}, expectedCodes)
	if err != nil {
		return nil, humanErr(err)
	}
	refs := make([]RefInfo, len(resp.Refs))
	for i, r := range resp.Refs {
		refs[i] = apiRefToInfo(r)
	}
	return refs, nil
}

func (r *prpcRemoteImpl) fetchInstanceURL(ctx context.Context, pin common.Pin) (string, error) {
	resp, err := r.repo.GetInstanceURL(ctx, &api.GetInstanceURLRequest{
		Package:  pin.PackageName,
		Instance: common.InstanceIDToObjectRef(pin.InstanceID),
	}, expectedCodes)
	if err != nil {
		return "", humanErr(err)
	}
	return resp.SignedUrl, nil
}

func (r *prpcRemoteImpl) fetchClientBinaryInfo(ctx context.Context, pin common.Pin) (*clientBinary, error) {
	resp, err := r.repo.DescribeClient(ctx, &api.DescribeClientRequest{
		Package:  pin.PackageName,
		Instance: common.InstanceIDToObjectRef(pin.InstanceID),
	}, expectedCodes)
	if err != nil {
		return nil, humanErr(err)
	}
	return &clientBinary{
		SHA1:     resp.LegacySha1,
		FetchURL: resp.ClientBinary.SignedUrl,
	}, nil
}

func (r *prpcRemoteImpl) describeInstance(ctx context.Context, pin common.Pin, opts *DescribeInstanceOpts) (*InstanceDescription, error) {
	if opts == nil {
		opts = &DescribeInstanceOpts{}
	}
	resp, err := r.repo.DescribeInstance(ctx, &api.DescribeInstanceRequest{
		Package:      pin.PackageName,
		Instance:     common.InstanceIDToObjectRef(pin.InstanceID),
		DescribeRefs: opts.DescribeRefs,
		DescribeTags: opts.DescribeTags,
	}, expectedCodes)
	if err != nil {
		return nil, humanErr(err)
	}
	return apiDescToInfo(resp), nil
}

////////////////////////////////////////////////////////////////////////////////
// Refs and tags.

func (r *prpcRemoteImpl) setRef(ctx context.Context, ref string, pin common.Pin) error {
	_, err := r.repo.CreateRef(ctx, &api.Ref{
		Name:     ref,
		Package:  pin.PackageName,
		Instance: common.InstanceIDToObjectRef(pin.InstanceID),
	}, expectedCodes)
	if grpc.Code(err) == codes.FailedPrecondition {
		return &pendingProcessingError{grpc.ErrorDesc(err)}
	}
	return humanErr(err)
}

func (r *prpcRemoteImpl) attachTags(ctx context.Context, pin common.Pin, tags []string) error {
	apiTags := make([]*api.Tag, len(tags))
	for i, t := range tags {
		var err error
		if apiTags[i], err = common.ParseInstanceTag(t); err != nil {
			return err
		}
	}
	_, err := r.repo.AttachTags(ctx, &api.AttachTagsRequest{
		Package:  pin.PackageName,
		Instance: common.InstanceIDToObjectRef(pin.InstanceID),
		Tags:     apiTags,
	}, expectedCodes)
	if grpc.Code(err) == codes.FailedPrecondition {
		return &pendingProcessingError{grpc.ErrorDesc(err)}
	}
	return humanErr(err)
}

////////////////////////////////////////////////////////////////////////////////
// Misc.

func (r *prpcRemoteImpl) listPackages(ctx context.Context, path string, recursive, includeHidden bool) ([]string, []string, error) {
	resp, err := r.repo.ListPrefix(ctx, &api.ListPrefixRequest{
		Prefix:        path,
		Recursive:     recursive,
		IncludeHidden: includeHidden,
	}, expectedCodes)
	if err != nil {
		return nil, nil, humanErr(err)
	}
	return resp.Packages, resp.Prefixes, nil
}

func (r *prpcRemoteImpl) searchInstances(ctx context.Context, packageName string, tags []string) (common.PinSlice, error) {
	apiTags := make([]*api.Tag, len(tags))
	for i, t := range tags {
		var err error
		if apiTags[i], err = common.ParseInstanceTag(t); err != nil {
			return nil, err
		}
	}

	resp, err := r.repo.SearchInstances(ctx, &api.SearchInstancesRequest{
		Package:  packageName,
		Tags:     apiTags,
		PageSize: 1000, // TODO(vadimsh): Support pagination on the client.
	}, expectedCodes)
	if err != nil {
		return nil, humanErr(err)
	}

	out := make(common.PinSlice, len(resp.Instances))
	for i, inst := range resp.Instances {
		out[i] = apiInstanceToInfo(inst).Pin
	}

	if resp.NextPageToken != "" {
		logging.Warningf(ctx, "Truncating the result only to first %d instances", len(resp.Instances))
	}
	return out, nil
}

func (r *prpcRemoteImpl) listInstances(ctx context.Context, packageName string, limit int, cursor string) (*listInstancesResponse, error) {
	resp, err := r.repo.ListInstances(ctx, &api.ListInstancesRequest{
		Package:   packageName,
		PageSize:  int32(limit),
		PageToken: cursor,
	}, expectedCodes)
	if err != nil {
		return nil, humanErr(err)
	}
	out := &listInstancesResponse{
		instances: make([]InstanceInfo, len(resp.Instances)),
		cursor:    resp.NextPageToken,
	}
	for i, inst := range resp.Instances {
		out.instances[i] = apiInstanceToInfo(inst)
	}
	return out, nil
}
