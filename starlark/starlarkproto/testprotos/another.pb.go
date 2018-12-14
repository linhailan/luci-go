// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/starlark/starlarkproto/testprotos/another.proto

package testprotos

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AnotherMessage struct {
	I                    int64    `protobuf:"varint,1,opt,name=i,proto3" json:"i,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnotherMessage) Reset()         { *m = AnotherMessage{} }
func (m *AnotherMessage) String() string { return proto.CompactTextString(m) }
func (*AnotherMessage) ProtoMessage()    {}
func (*AnotherMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9042e40bd6bdc252, []int{0}
}

func (m *AnotherMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnotherMessage.Unmarshal(m, b)
}
func (m *AnotherMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnotherMessage.Marshal(b, m, deterministic)
}
func (m *AnotherMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnotherMessage.Merge(m, src)
}
func (m *AnotherMessage) XXX_Size() int {
	return xxx_messageInfo_AnotherMessage.Size(m)
}
func (m *AnotherMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AnotherMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AnotherMessage proto.InternalMessageInfo

func (m *AnotherMessage) GetI() int64 {
	if m != nil {
		return m.I
	}
	return 0
}

func init() {
	proto.RegisterType((*AnotherMessage)(nil), "testprotos.AnotherMessage")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/starlark/starlarkproto/testprotos/another.proto", fileDescriptor_9042e40bd6bdc252)
}

var fileDescriptor_9042e40bd6bdc252 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x49, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x2f, 0x2e, 0x49, 0x2c, 0xca, 0x49, 0x2c, 0xca, 0x86, 0x33, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0xf5, 0x4b, 0x52, 0x8b, 0x4b, 0xc0, 0xac, 0x62, 0xfd, 0xc4, 0xbc, 0xfc, 0x92, 0x8c, 0xd4, 0x22,
	0x3d, 0x30, 0x57, 0x88, 0x0b, 0x21, 0xa3, 0x24, 0xc7, 0xc5, 0xe7, 0x08, 0x91, 0xf4, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0xe2, 0xe1, 0x62, 0xcc, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e,
	0x62, 0xcc, 0x4c, 0x62, 0x03, 0xab, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x90, 0x77, 0x38,
	0xba, 0x7a, 0x00, 0x00, 0x00,
}