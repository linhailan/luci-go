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

	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/proto/google"
	"go.chromium.org/luci/common/retry"
	"go.chromium.org/luci/grpc/prpc"

	api "go.chromium.org/luci/cipd/api/cipd/v1"
	"go.chromium.org/luci/cipd/common"
)

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
						Retries: 5,
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
// Type converters, until callers switch to using API types directly.

func apiInstanceToInfo(inst *api.Instance) InstanceInfo {
	return InstanceInfo{
		Pin: common.Pin{
			PackageName: inst.Package,
			InstanceID:  common.ObjectRefToInstanceID(inst.Instance),
		},
		RegisteredBy: inst.RegisteredBy,
		RegisteredTs: UnixTime(google.TimeFromProto(inst.RegisteredTs)),
	}
}

func apiRefToInfo(r *api.Ref) RefInfo {
	return RefInfo{
		Ref:        r.Name,
		InstanceID: common.ObjectRefToInstanceID(r.Instance),
		ModifiedBy: r.ModifiedBy,
		ModifiedTs: UnixTime(google.TimeFromProto(r.ModifiedTs)),
	}
}

func apiTagToInfo(t *api.Tag) TagInfo {
	return TagInfo{
		Tag:          common.JoinInstanceTag(t),
		RegisteredBy: t.AttachedBy,
		RegisteredTs: UnixTime(google.TimeFromProto(t.AttachedTs)),
	}
}

////////////////////////////////////////////////////////////////////////////////
// ACLs.

var legacyRoles = map[string]api.Role{
	"READER": api.Role_READER,
	"WRITER": api.Role_WRITER,
	"OWNER":  api.Role_OWNER,
}

func grantRole(m *api.PrefixMetadata, role api.Role, principal string) bool {
	var roleAcl *api.PrefixMetadata_ACL
	for _, acl := range m.Acls {
		if acl.Role != role {
			continue
		}
		for _, p := range acl.Principals {
			if p == principal {
				return false // already have it
			}
		}
		roleAcl = acl
	}

	if roleAcl != nil {
		// Append to the existing ACL.
		roleAcl.Principals = append(roleAcl.Principals, principal)
	} else {
		// Add new ACL for this role, this is the first one.
		m.Acls = append(m.Acls, &api.PrefixMetadata_ACL{
			Role:       role,
			Principals: []string{principal},
		})
	}

	return true
}

func revokeRole(m *api.PrefixMetadata, role api.Role, principal string) bool {
	dirty := false
	for _, acl := range m.Acls {
		if acl.Role != role {
			continue
		}
		filtered := acl.Principals[:0]
		for _, p := range acl.Principals {
			if p != principal {
				filtered = append(filtered, p)
			}
		}
		if len(filtered) != len(acl.Principals) {
			acl.Principals = filtered
			dirty = true
		}
	}

	if !dirty {
		return false
	}

	// Kick out empty ACL entries.
	acls := m.Acls[:0]
	for _, acl := range m.Acls {
		if len(acl.Principals) != 0 {
			acls = append(acls, acl)
		}
	}
	if len(acls) == 0 {
		m.Acls = nil
	} else {
		m.Acls = acls
	}
	return true
}

func (r *prpcRemoteImpl) fetchACL(ctx context.Context, packagePath string) ([]PackageACL, error) {
	resp, err := r.repo.GetInheritedPrefixMetadata(ctx, &api.PrefixRequest{
		Prefix: packagePath,
	})
	if err != nil {
		return nil, err
	}
	var out []PackageACL
	for _, p := range resp.PerPrefixMetadata {
		var acls []PackageACL
		for _, acl := range p.Acls {
			role := acl.Role.String()
			found := false
			for i, existing := range acls {
				if existing.Role == role {
					acls[i].Principals = append(acls[i].Principals, acl.Principals...)
					found = true
					break
				}
			}
			if !found {
				acls = append(acls, PackageACL{
					PackagePath: p.Prefix,
					Role:        role,
					Principals:  acl.Principals,
					ModifiedBy:  p.UpdateUser,
					ModifiedTs:  UnixTime(google.TimeFromProto(p.UpdateTime)),
				})
			}
		}
		out = append(out, acls...)
	}
	return out, nil
}

func (r *prpcRemoteImpl) modifyACL(ctx context.Context, packagePath string, changes []PackageACLChange) error {
	// Fetch existing metadata, if any.
	meta, err := r.repo.GetPrefixMetadata(ctx, &api.PrefixRequest{
		Prefix: packagePath,
	}, prpc.ExpectedCode(codes.NotFound))
	if code := grpc.Code(err); code != codes.OK && code != codes.NotFound {
		return err
	}

	// Construct new empty metadata for codes.NotFound.
	if meta == nil {
		meta = &api.PrefixMetadata{Prefix: packagePath}
	}

	// Apply mutations.
	dirty := false
	for _, ch := range changes {
		role, ok := legacyRoles[ch.Role]
		if !ok {
			// Just log and ignore. Aborting with error here breaks 'acl-edit -revoke'
			// functionality, since it always tries to revoke all possible roles,
			// including unsupported COUNTER_WRITER.
			logging.Warningf(ctx, "Ignoring role %q not supported in v2", ch.Role)
			continue
		}
		changed := false
		switch ch.Action {
		case GrantRole:
			changed = grantRole(meta, role, ch.Principal)
		case RevokeRole:
			changed = revokeRole(meta, role, ch.Principal)
		default:
			return fmt.Errorf("unrecognized PackageACLChangeAction %q", ch.Action)
		}
		dirty = dirty || changed
	}

	if !dirty {
		return nil
	}

	// Store the new metadata. This call will check meta.Fingerprint.
	_, err = r.repo.UpdatePrefixMetadata(ctx, meta)
	return err
}

func (r *prpcRemoteImpl) fetchRoles(ctx context.Context, packagePath string) ([]string, error) {
	resp, err := r.repo.GetRolesInPrefix(ctx, &api.PrefixRequest{
		Prefix: packagePath,
	})
	if err != nil {
		return nil, err
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
	})
	if err != nil {
		return nil, err
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

func (r *prpcRemoteImpl) resolveVersion(ctx context.Context, packageName, version string) (pin common.Pin, err error) {
	resp, err := r.repo.ResolveVersion(ctx, &api.ResolveVersionRequest{
		Package: packageName,
		Version: version,
	}, prpc.ExpectedCode(codes.NotFound, codes.FailedPrecondition))
	switch grpc.Code(err) {
	case codes.OK:
		pin = common.Pin{
			PackageName: packageName,
			InstanceID:  common.ObjectRefToInstanceID(resp.Instance),
		}
	case codes.NotFound, codes.FailedPrecondition:
		// Return a friendlier looking error message without gRPC framing.
		err = errors.New(grpc.ErrorDesc(err))
	}
	return
}

func (r *prpcRemoteImpl) fetchPackageRefs(ctx context.Context, packageName string) ([]RefInfo, error) {
	resp, err := r.repo.ListRefs(ctx, &api.ListRefsRequest{
		Package: packageName,
	})
	if err != nil {
		return nil, err
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
	}, prpc.ExpectedCode(codes.NotFound))
	switch grpc.Code(err) {
	case codes.OK:
		return resp.SignedUrl, nil
	case codes.NotFound:
		return "", errors.New(grpc.ErrorDesc(err))
	default:
		return "", err
	}
}

func (r *prpcRemoteImpl) fetchClientBinaryInfo(ctx context.Context, pin common.Pin) (*clientBinary, error) {
	resp, err := r.repo.DescribeClient(ctx, &api.DescribeClientRequest{
		Package:  pin.PackageName,
		Instance: common.InstanceIDToObjectRef(pin.InstanceID),
	})
	if err != nil {
		return nil, err
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
	})
	if err != nil {
		return nil, err
	}
	desc := &InstanceDescription{
		InstanceInfo: apiInstanceToInfo(resp.Instance),
		Refs:         make([]RefInfo, len(resp.Refs)),
		Tags:         make([]TagInfo, len(resp.Tags)),
	}
	for i, r := range resp.Refs {
		desc.Refs[i] = apiRefToInfo(r)
	}
	for i, t := range resp.Tags {
		desc.Tags[i] = apiTagToInfo(t)
	}
	return desc, nil
}

////////////////////////////////////////////////////////////////////////////////
// Refs and tags.

func (r *prpcRemoteImpl) setRef(ctx context.Context, ref string, pin common.Pin) error {
	_, err := r.repo.CreateRef(ctx, &api.Ref{
		Name:     ref,
		Package:  pin.PackageName,
		Instance: common.InstanceIDToObjectRef(pin.InstanceID),
	}, prpc.ExpectedCode(codes.FailedPrecondition))
	if grpc.Code(err) == codes.FailedPrecondition {
		return &pendingProcessingError{grpc.ErrorDesc(err)}
	}
	return err
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
	})
	if grpc.Code(err) == codes.FailedPrecondition {
		return &pendingProcessingError{grpc.ErrorDesc(err)}
	}
	return err
}

////////////////////////////////////////////////////////////////////////////////
// Misc.

func (r *prpcRemoteImpl) listPackages(ctx context.Context, path string, recursive, includeHidden bool) ([]string, []string, error) {
	resp, err := r.repo.ListPrefix(ctx, &api.ListPrefixRequest{
		Prefix:        path,
		Recursive:     recursive,
		IncludeHidden: includeHidden,
	})
	if err != nil {
		return nil, nil, err
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
	})
	if err != nil {
		return nil, err
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
	})
	if err != nil {
		return nil, err
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
