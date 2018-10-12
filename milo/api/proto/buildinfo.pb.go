// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/milo/api/proto/buildinfo.proto

package milo

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	milo "go.chromium.org/luci/common/proto/milo"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BuildInfoRequest struct {
	// Types that are valid to be assigned to Build:
	//	*BuildInfoRequest_Buildbot
	//	*BuildInfoRequest_Swarming_
	//	*BuildInfoRequest_Buildbucket_
	Build isBuildInfoRequest_Build `protobuf_oneof:"build"`
	// Project hint is a LUCI project suggestion for this build. Some builds,
	// notably older ones, may not contain enough metadata to resolve their
	// project. Resolution may succeed if this hint is provided and correct.
	//
	// This field is optional, and its use is discouraged unless necessary.
	ProjectHint          string   `protobuf:"bytes,11,opt,name=project_hint,json=projectHint,proto3" json:"project_hint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildInfoRequest) Reset()         { *m = BuildInfoRequest{} }
func (m *BuildInfoRequest) String() string { return proto.CompactTextString(m) }
func (*BuildInfoRequest) ProtoMessage()    {}
func (*BuildInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98f9739304de6fa6, []int{0}
}

func (m *BuildInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildInfoRequest.Unmarshal(m, b)
}
func (m *BuildInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildInfoRequest.Marshal(b, m, deterministic)
}
func (m *BuildInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildInfoRequest.Merge(m, src)
}
func (m *BuildInfoRequest) XXX_Size() int {
	return xxx_messageInfo_BuildInfoRequest.Size(m)
}
func (m *BuildInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildInfoRequest proto.InternalMessageInfo

type isBuildInfoRequest_Build interface {
	isBuildInfoRequest_Build()
}

type BuildInfoRequest_Buildbot struct {
	Buildbot *BuildInfoRequest_BuildBot `protobuf:"bytes,1,opt,name=buildbot,proto3,oneof"`
}

type BuildInfoRequest_Swarming_ struct {
	Swarming *BuildInfoRequest_Swarming `protobuf:"bytes,2,opt,name=swarming,proto3,oneof"`
}

type BuildInfoRequest_Buildbucket_ struct {
	Buildbucket *BuildInfoRequest_Buildbucket `protobuf:"bytes,3,opt,name=buildbucket,proto3,oneof"`
}

func (*BuildInfoRequest_Buildbot) isBuildInfoRequest_Build() {}

func (*BuildInfoRequest_Swarming_) isBuildInfoRequest_Build() {}

func (*BuildInfoRequest_Buildbucket_) isBuildInfoRequest_Build() {}

func (m *BuildInfoRequest) GetBuild() isBuildInfoRequest_Build {
	if m != nil {
		return m.Build
	}
	return nil
}

func (m *BuildInfoRequest) GetBuildbot() *BuildInfoRequest_BuildBot {
	if x, ok := m.GetBuild().(*BuildInfoRequest_Buildbot); ok {
		return x.Buildbot
	}
	return nil
}

func (m *BuildInfoRequest) GetSwarming() *BuildInfoRequest_Swarming {
	if x, ok := m.GetBuild().(*BuildInfoRequest_Swarming_); ok {
		return x.Swarming
	}
	return nil
}

func (m *BuildInfoRequest) GetBuildbucket() *BuildInfoRequest_Buildbucket {
	if x, ok := m.GetBuild().(*BuildInfoRequest_Buildbucket_); ok {
		return x.Buildbucket
	}
	return nil
}

func (m *BuildInfoRequest) GetProjectHint() string {
	if m != nil {
		return m.ProjectHint
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BuildInfoRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BuildInfoRequest_OneofMarshaler, _BuildInfoRequest_OneofUnmarshaler, _BuildInfoRequest_OneofSizer, []interface{}{
		(*BuildInfoRequest_Buildbot)(nil),
		(*BuildInfoRequest_Swarming_)(nil),
		(*BuildInfoRequest_Buildbucket_)(nil),
	}
}

func _BuildInfoRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BuildInfoRequest)
	// build
	switch x := m.Build.(type) {
	case *BuildInfoRequest_Buildbot:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Buildbot); err != nil {
			return err
		}
	case *BuildInfoRequest_Swarming_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Swarming); err != nil {
			return err
		}
	case *BuildInfoRequest_Buildbucket_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Buildbucket); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BuildInfoRequest.Build has unexpected type %T", x)
	}
	return nil
}

func _BuildInfoRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BuildInfoRequest)
	switch tag {
	case 1: // build.buildbot
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuildInfoRequest_BuildBot)
		err := b.DecodeMessage(msg)
		m.Build = &BuildInfoRequest_Buildbot{msg}
		return true, err
	case 2: // build.swarming
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuildInfoRequest_Swarming)
		err := b.DecodeMessage(msg)
		m.Build = &BuildInfoRequest_Swarming_{msg}
		return true, err
	case 3: // build.buildbucket
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuildInfoRequest_Buildbucket)
		err := b.DecodeMessage(msg)
		m.Build = &BuildInfoRequest_Buildbucket_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BuildInfoRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BuildInfoRequest)
	// build
	switch x := m.Build.(type) {
	case *BuildInfoRequest_Buildbot:
		s := proto.Size(x.Buildbot)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BuildInfoRequest_Swarming_:
		s := proto.Size(x.Swarming)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BuildInfoRequest_Buildbucket_:
		s := proto.Size(x.Buildbucket)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The request for the name of a BuildBot built.
type BuildInfoRequest_BuildBot struct {
	// The master name.
	MasterName string `protobuf:"bytes,1,opt,name=master_name,json=masterName,proto3" json:"master_name,omitempty"`
	// The builder name server.
	BuilderName string `protobuf:"bytes,2,opt,name=builder_name,json=builderName,proto3" json:"builder_name,omitempty"`
	// The build number.
	BuildNumber          int64    `protobuf:"varint,3,opt,name=build_number,json=buildNumber,proto3" json:"build_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildInfoRequest_BuildBot) Reset()         { *m = BuildInfoRequest_BuildBot{} }
func (m *BuildInfoRequest_BuildBot) String() string { return proto.CompactTextString(m) }
func (*BuildInfoRequest_BuildBot) ProtoMessage()    {}
func (*BuildInfoRequest_BuildBot) Descriptor() ([]byte, []int) {
	return fileDescriptor_98f9739304de6fa6, []int{0, 0}
}

func (m *BuildInfoRequest_BuildBot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildInfoRequest_BuildBot.Unmarshal(m, b)
}
func (m *BuildInfoRequest_BuildBot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildInfoRequest_BuildBot.Marshal(b, m, deterministic)
}
func (m *BuildInfoRequest_BuildBot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildInfoRequest_BuildBot.Merge(m, src)
}
func (m *BuildInfoRequest_BuildBot) XXX_Size() int {
	return xxx_messageInfo_BuildInfoRequest_BuildBot.Size(m)
}
func (m *BuildInfoRequest_BuildBot) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildInfoRequest_BuildBot.DiscardUnknown(m)
}

var xxx_messageInfo_BuildInfoRequest_BuildBot proto.InternalMessageInfo

func (m *BuildInfoRequest_BuildBot) GetMasterName() string {
	if m != nil {
		return m.MasterName
	}
	return ""
}

func (m *BuildInfoRequest_BuildBot) GetBuilderName() string {
	if m != nil {
		return m.BuilderName
	}
	return ""
}

func (m *BuildInfoRequest_BuildBot) GetBuildNumber() int64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

// The request containing a Swarming task.
type BuildInfoRequest_Swarming struct {
	// Host is the hostname of the Swarming server to connect to
	// (e.g., "swarming.example.com").
	//
	// This is optional. If omitted or empty, Milo's default Swarming server
	// will be used.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The Swarming task name.
	Task                 string   `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildInfoRequest_Swarming) Reset()         { *m = BuildInfoRequest_Swarming{} }
func (m *BuildInfoRequest_Swarming) String() string { return proto.CompactTextString(m) }
func (*BuildInfoRequest_Swarming) ProtoMessage()    {}
func (*BuildInfoRequest_Swarming) Descriptor() ([]byte, []int) {
	return fileDescriptor_98f9739304de6fa6, []int{0, 1}
}

func (m *BuildInfoRequest_Swarming) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildInfoRequest_Swarming.Unmarshal(m, b)
}
func (m *BuildInfoRequest_Swarming) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildInfoRequest_Swarming.Marshal(b, m, deterministic)
}
func (m *BuildInfoRequest_Swarming) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildInfoRequest_Swarming.Merge(m, src)
}
func (m *BuildInfoRequest_Swarming) XXX_Size() int {
	return xxx_messageInfo_BuildInfoRequest_Swarming.Size(m)
}
func (m *BuildInfoRequest_Swarming) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildInfoRequest_Swarming.DiscardUnknown(m)
}

var xxx_messageInfo_BuildInfoRequest_Swarming proto.InternalMessageInfo

func (m *BuildInfoRequest_Swarming) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *BuildInfoRequest_Swarming) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

// The request containing a BuildBucket build.
type BuildInfoRequest_Buildbucket struct {
	// The build ID of the buildbucket build.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildInfoRequest_Buildbucket) Reset()         { *m = BuildInfoRequest_Buildbucket{} }
func (m *BuildInfoRequest_Buildbucket) String() string { return proto.CompactTextString(m) }
func (*BuildInfoRequest_Buildbucket) ProtoMessage()    {}
func (*BuildInfoRequest_Buildbucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_98f9739304de6fa6, []int{0, 2}
}

func (m *BuildInfoRequest_Buildbucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildInfoRequest_Buildbucket.Unmarshal(m, b)
}
func (m *BuildInfoRequest_Buildbucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildInfoRequest_Buildbucket.Marshal(b, m, deterministic)
}
func (m *BuildInfoRequest_Buildbucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildInfoRequest_Buildbucket.Merge(m, src)
}
func (m *BuildInfoRequest_Buildbucket) XXX_Size() int {
	return xxx_messageInfo_BuildInfoRequest_Buildbucket.Size(m)
}
func (m *BuildInfoRequest_Buildbucket) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildInfoRequest_Buildbucket.DiscardUnknown(m)
}

var xxx_messageInfo_BuildInfoRequest_Buildbucket proto.InternalMessageInfo

func (m *BuildInfoRequest_Buildbucket) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// The request containing the name of the master.
type BuildInfoResponse struct {
	// The LUCI project that this build belongs to.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The main build step.
	Step *milo.Step `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
	// The LogDog annotation stream for this build. The Prefix will be populated
	// and can be used as the prefix for any un-prefixed LogdogStream in "step".
	AnnotationStream     *milo.LogdogStream `protobuf:"bytes,3,opt,name=annotation_stream,json=annotationStream,proto3" json:"annotation_stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BuildInfoResponse) Reset()         { *m = BuildInfoResponse{} }
func (m *BuildInfoResponse) String() string { return proto.CompactTextString(m) }
func (*BuildInfoResponse) ProtoMessage()    {}
func (*BuildInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98f9739304de6fa6, []int{1}
}

func (m *BuildInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildInfoResponse.Unmarshal(m, b)
}
func (m *BuildInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildInfoResponse.Marshal(b, m, deterministic)
}
func (m *BuildInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildInfoResponse.Merge(m, src)
}
func (m *BuildInfoResponse) XXX_Size() int {
	return xxx_messageInfo_BuildInfoResponse.Size(m)
}
func (m *BuildInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuildInfoResponse proto.InternalMessageInfo

func (m *BuildInfoResponse) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *BuildInfoResponse) GetStep() *milo.Step {
	if m != nil {
		return m.Step
	}
	return nil
}

func (m *BuildInfoResponse) GetAnnotationStream() *milo.LogdogStream {
	if m != nil {
		return m.AnnotationStream
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildInfoRequest)(nil), "milo.BuildInfoRequest")
	proto.RegisterType((*BuildInfoRequest_BuildBot)(nil), "milo.BuildInfoRequest.BuildBot")
	proto.RegisterType((*BuildInfoRequest_Swarming)(nil), "milo.BuildInfoRequest.Swarming")
	proto.RegisterType((*BuildInfoRequest_Buildbucket)(nil), "milo.BuildInfoRequest.Buildbucket")
	proto.RegisterType((*BuildInfoResponse)(nil), "milo.BuildInfoResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/milo/api/proto/buildinfo.proto", fileDescriptor_98f9739304de6fa6)
}

var fileDescriptor_98f9739304de6fa6 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0xd7, 0xa6, 0xb0, 0xf6, 0x04, 0xa1, 0xcd, 0x17, 0x10, 0x45, 0x82, 0x8d, 0x5e, 0xed,
	0x2a, 0x91, 0x3a, 0x09, 0x71, 0x83, 0x90, 0x8a, 0x04, 0x45, 0x42, 0xbb, 0x70, 0x1f, 0xa0, 0x72,
	0x13, 0x2f, 0x35, 0xab, 0x7d, 0x32, 0xfb, 0x44, 0xbc, 0x05, 0x0f, 0xc6, 0x53, 0xa1, 0xd8, 0xce,
	0x56, 0x60, 0xea, 0x9d, 0xfd, 0xfb, 0xff, 0xfe, 0x73, 0x8e, 0x6d, 0xb8, 0x6e, 0xb0, 0xa8, 0x76,
	0x16, 0xb5, 0xea, 0x74, 0x81, 0xb6, 0x29, 0xf7, 0x5d, 0xa5, 0x4a, 0xad, 0xf6, 0x58, 0x8a, 0x56,
	0x95, 0xad, 0x45, 0xc2, 0x72, 0xdb, 0xa9, 0x7d, 0xad, 0xcc, 0x2d, 0x16, 0x7e, 0xcf, 0x26, 0xfd,
	0x79, 0xfe, 0xe1, 0x49, 0xb4, 0x42, 0xad, 0xd1, 0x44, 0x30, 0xe4, 0x18, 0x83, 0x24, 0x48, 0xa1,
	0x71, 0x81, 0x9f, 0xff, 0x4e, 0xe0, 0x6c, 0xd9, 0x67, 0x7e, 0x33, 0xb7, 0xc8, 0xe5, 0x7d, 0x27,
	0x1d, 0xb1, 0x8f, 0x30, 0xf5, 0x75, 0xb6, 0x48, 0xd9, 0xe8, 0x72, 0x74, 0x95, 0x2e, 0x2e, 0x8a,
	0x9e, 0x2f, 0xfe, 0x75, 0x06, 0x61, 0x89, 0xb4, 0x3a, 0xe1, 0x0f, 0x48, 0x8f, 0xbb, 0x9f, 0xc2,
	0x6a, 0x65, 0x9a, 0x6c, 0x7c, 0x14, 0x5f, 0x47, 0x5b, 0x8f, 0x0f, 0x08, 0xfb, 0x02, 0x69, 0x88,
	0xea, 0xaa, 0x3b, 0x49, 0x59, 0xe2, 0x13, 0xe6, 0xc7, 0x1a, 0x08, 0xce, 0xd5, 0x09, 0x3f, 0x04,
	0xd9, 0x3b, 0x78, 0xd1, 0x5a, 0xfc, 0x21, 0x2b, 0xda, 0xec, 0x94, 0xa1, 0x2c, 0xbd, 0x1c, 0x5d,
	0xcd, 0x78, 0x1a, 0xb5, 0x95, 0x32, 0x94, 0xdf, 0xc3, 0x74, 0x98, 0x80, 0x5d, 0x40, 0xaa, 0x85,
	0x23, 0x69, 0x37, 0x46, 0x68, 0xe9, 0xe7, 0x9e, 0x71, 0x08, 0xd2, 0x8d, 0xd0, 0xb2, 0xcf, 0xf3,
	0xf1, 0x83, 0x63, 0x1c, 0xf2, 0xa2, 0xf6, 0x97, 0x65, 0x63, 0x3a, 0xbd, 0x95, 0xd6, 0xf7, 0x9e,
	0x44, 0xcb, 0x8d, 0x97, 0xf2, 0x05, 0x4c, 0x87, 0xa9, 0x19, 0x83, 0xc9, 0x0e, 0x1d, 0xc5, 0x5a,
	0x7e, 0xdd, 0x6b, 0x24, 0xdc, 0x5d, 0x4c, 0xf7, 0xeb, 0xfc, 0x0d, 0xa4, 0x07, 0x73, 0xb2, 0x97,
	0x30, 0x56, 0xb5, 0x87, 0x12, 0x3e, 0x56, 0xf5, 0xf2, 0x14, 0x9e, 0xf9, 0x0a, 0xf3, 0x5f, 0x23,
	0x38, 0x3f, 0xb8, 0x21, 0xd7, 0xa2, 0x71, 0x92, 0x65, 0x70, 0x1a, 0x67, 0x8e, 0x85, 0x86, 0x2d,
	0x7b, 0x0b, 0x13, 0x47, 0xb2, 0x8d, 0x8f, 0x04, 0xe1, 0x8a, 0xd7, 0x24, 0x5b, 0xee, 0x75, 0xf6,
	0x09, 0xce, 0x1f, 0x7f, 0xcc, 0xc6, 0x91, 0x95, 0x42, 0xc7, 0xf7, 0x60, 0xc1, 0xfc, 0x1d, 0x9b,
	0x1a, 0x9b, 0xb5, 0x3f, 0xe1, 0x67, 0x8f, 0xe6, 0xa0, 0x2c, 0x3e, 0xc3, 0xec, 0xa1, 0x1f, 0xf6,
	0x1e, 0x92, 0xaf, 0x92, 0xd8, 0xab, 0xa7, 0x5f, 0x32, 0x7f, 0xfd, 0x9f, 0x1e, 0xfa, 0xdf, 0x3e,
	0xf7, 0x3f, 0xf5, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x6f, 0x78, 0x53, 0x20, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BuildInfoClient is the client API for BuildInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildInfoClient interface {
	Get(ctx context.Context, in *BuildInfoRequest, opts ...grpc.CallOption) (*BuildInfoResponse, error)
}
type buildInfoPRPCClient struct {
	client *prpc.Client
}

func NewBuildInfoPRPCClient(client *prpc.Client) BuildInfoClient {
	return &buildInfoPRPCClient{client}
}

func (c *buildInfoPRPCClient) Get(ctx context.Context, in *BuildInfoRequest, opts ...grpc.CallOption) (*BuildInfoResponse, error) {
	out := new(BuildInfoResponse)
	err := c.client.Call(ctx, "milo.BuildInfo", "Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type buildInfoClient struct {
	cc *grpc.ClientConn
}

func NewBuildInfoClient(cc *grpc.ClientConn) BuildInfoClient {
	return &buildInfoClient{cc}
}

func (c *buildInfoClient) Get(ctx context.Context, in *BuildInfoRequest, opts ...grpc.CallOption) (*BuildInfoResponse, error) {
	out := new(BuildInfoResponse)
	err := c.cc.Invoke(ctx, "/milo.BuildInfo/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildInfoServer is the server API for BuildInfo service.
type BuildInfoServer interface {
	Get(context.Context, *BuildInfoRequest) (*BuildInfoResponse, error)
}

func RegisterBuildInfoServer(s prpc.Registrar, srv BuildInfoServer) {
	s.RegisterService(&_BuildInfo_serviceDesc, srv)
}

func _BuildInfo_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildInfoServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milo.BuildInfo/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildInfoServer).Get(ctx, req.(*BuildInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BuildInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milo.BuildInfo",
	HandlerType: (*BuildInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BuildInfo_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/milo/api/proto/buildinfo.proto",
}
