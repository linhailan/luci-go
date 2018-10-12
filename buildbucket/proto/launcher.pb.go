// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/launcher.proto

package buildbucketpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// A collection of build-related secrets we might pass from Buildbucket to Kitchen.
type BuildSecrets struct {
	// Token to identify RPCs associated with the same build.
	BuildToken           []byte   `protobuf:"bytes,1,opt,name=build_token,json=buildToken,proto3" json:"build_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildSecrets) Reset()         { *m = BuildSecrets{} }
func (m *BuildSecrets) String() string { return proto.CompactTextString(m) }
func (*BuildSecrets) ProtoMessage()    {}
func (*BuildSecrets) Descriptor() ([]byte, []int) {
	return fileDescriptor_45f9e6fb262a81d2, []int{0}
}

func (m *BuildSecrets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildSecrets.Unmarshal(m, b)
}
func (m *BuildSecrets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildSecrets.Marshal(b, m, deterministic)
}
func (m *BuildSecrets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildSecrets.Merge(m, src)
}
func (m *BuildSecrets) XXX_Size() int {
	return xxx_messageInfo_BuildSecrets.Size(m)
}
func (m *BuildSecrets) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildSecrets.DiscardUnknown(m)
}

var xxx_messageInfo_BuildSecrets proto.InternalMessageInfo

func (m *BuildSecrets) GetBuildToken() []byte {
	if m != nil {
		return m.BuildToken
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildSecrets)(nil), "buildbucket.v2.BuildSecrets")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/launcher.proto", fileDescriptor_45f9e6fb262a81d2)
}

var fileDescriptor_45f9e6fb262a81d2 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0x2a, 0xcd, 0xcc, 0x49, 0x49, 0x2a, 0x4d, 0xce, 0x4e, 0x2d, 0xd1, 0x2f, 0x28, 0xca,
	0x2f, 0xc9, 0xd7, 0xcf, 0x49, 0x2c, 0xcd, 0x4b, 0xce, 0x48, 0x2d, 0xd2, 0x03, 0x73, 0x85, 0xf8,
	0x90, 0x54, 0xe8, 0x95, 0x19, 0x29, 0xe9, 0x73, 0xf1, 0x38, 0x81, 0x44, 0x82, 0x53, 0x93, 0x8b,
	0x52, 0x4b, 0x8a, 0x85, 0xe4, 0xb9, 0xb8, 0xc1, 0x2a, 0xe2, 0x4b, 0xf2, 0xb3, 0x53, 0xf3, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xb8, 0xc0, 0x42, 0x21, 0x20, 0x11, 0x27, 0xb3, 0x28, 0x13,
	0xe2, 0x6c, 0xb6, 0x46, 0x12, 0x29, 0x48, 0x4a, 0x62, 0x03, 0x0b, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x41, 0x6d, 0x70, 0xe4, 0xb8, 0x00, 0x00, 0x00,
}
