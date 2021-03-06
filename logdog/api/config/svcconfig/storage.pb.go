// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/logdog/api/config/svcconfig/storage.proto

package svcconfig

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
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

// Storage is the in-transit storage configuration.
type Storage struct {
	// Type is the transport configuration that is being used.
	//
	// Types that are valid to be assigned to Type:
	//	*Storage_Bigtable
	Type                 isStorage_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_955b461662b6fa9d, []int{0}
}

func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

type isStorage_Type interface {
	isStorage_Type()
}

type Storage_Bigtable struct {
	Bigtable *Storage_BigTable `protobuf:"bytes,1,opt,name=bigtable,proto3,oneof"`
}

func (*Storage_Bigtable) isStorage_Type() {}

func (m *Storage) GetType() isStorage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Storage) GetBigtable() *Storage_BigTable {
	if x, ok := m.GetType().(*Storage_Bigtable); ok {
		return x.Bigtable
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Storage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Storage_Bigtable)(nil),
	}
}

// BigTable is the set of BigTable configuration parameters.
type Storage_BigTable struct {
	// The name of the BigTable instance project.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The name of the BigTable instance.
	Instance string `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
	// The name of the BigTable instance's log table.
	LogTableName         string   `protobuf:"bytes,3,opt,name=log_table_name,json=logTableName,proto3" json:"log_table_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Storage_BigTable) Reset()         { *m = Storage_BigTable{} }
func (m *Storage_BigTable) String() string { return proto.CompactTextString(m) }
func (*Storage_BigTable) ProtoMessage()    {}
func (*Storage_BigTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_955b461662b6fa9d, []int{0, 0}
}

func (m *Storage_BigTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage_BigTable.Unmarshal(m, b)
}
func (m *Storage_BigTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage_BigTable.Marshal(b, m, deterministic)
}
func (m *Storage_BigTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage_BigTable.Merge(m, src)
}
func (m *Storage_BigTable) XXX_Size() int {
	return xxx_messageInfo_Storage_BigTable.Size(m)
}
func (m *Storage_BigTable) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage_BigTable.DiscardUnknown(m)
}

var xxx_messageInfo_Storage_BigTable proto.InternalMessageInfo

func (m *Storage_BigTable) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Storage_BigTable) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Storage_BigTable) GetLogTableName() string {
	if m != nil {
		return m.LogTableName
	}
	return ""
}

func init() {
	proto.RegisterType((*Storage)(nil), "svcconfig.Storage")
	proto.RegisterType((*Storage_BigTable)(nil), "svcconfig.Storage.BigTable")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/logdog/api/config/svcconfig/storage.proto", fileDescriptor_955b461662b6fa9d)
}

var fileDescriptor_955b461662b6fa9d = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x49, 0x5b, 0xb5, 0xa9, 0x8b, 0x10, 0xf2, 0x14, 0x05, 0x09, 0x21, 0xc4, 0xc0, 0x64,
	0x4b, 0x30, 0xb1, 0x30, 0x74, 0x42, 0x0c, 0x0c, 0xa1, 0x7b, 0xe4, 0xb8, 0x97, 0xc3, 0xc8, 0xf6,
	0x45, 0xae, 0x83, 0xe0, 0x17, 0xf2, 0xb7, 0x50, 0x1c, 0x92, 0xcd, 0xcf, 0xef, 0x7b, 0x9f, 0x74,
	0xec, 0x19, 0x49, 0xe8, 0x8f, 0x40, 0xce, 0xf4, 0x4e, 0x50, 0x40, 0x69, 0x7b, 0x6d, 0xa4, 0x25,
	0x3c, 0x12, 0x4a, 0xd5, 0x19, 0xa9, 0xc9, 0xb7, 0x06, 0xe5, 0xe9, 0x4b, 0x4f, 0xaf, 0x48, 0x41,
	0x21, 0x88, 0x2e, 0x50, 0x24, 0xbe, 0x9d, 0x8b, 0xf2, 0x1a, 0x89, 0xd0, 0x82, 0x4c, 0x45, 0xd3,
	0xb7, 0xf2, 0xd8, 0x07, 0x15, 0x0d, 0xf9, 0x11, 0xbd, 0xfd, 0xcd, 0xd8, 0xe6, 0x7d, 0x1c, 0xf3,
	0x27, 0x96, 0x37, 0x06, 0xa3, 0x6a, 0x2c, 0x14, 0xd9, 0x4d, 0x76, 0xbf, 0x7b, 0xb8, 0x12, 0xb3,
	0x49, 0xfc, 0x53, 0x62, 0x6f, 0xf0, 0x30, 0x20, 0x2f, 0x67, 0xd5, 0x8c, 0x97, 0x2d, 0xcb, 0xa7,
	0x7f, 0x5e, 0xb0, 0x4d, 0x17, 0xe8, 0x13, 0x74, 0x4c, 0x96, 0x6d, 0x35, 0x45, 0x5e, 0xb2, 0xdc,
	0xf8, 0x53, 0x54, 0x5e, 0x43, 0xb1, 0x48, 0xd5, 0x9c, 0xf9, 0x1d, 0xbb, 0xb0, 0x84, 0x75, 0xd2,
	0xd5, 0x5e, 0x39, 0x28, 0x96, 0x89, 0x38, 0xb7, 0x34, 0x7a, 0xdf, 0x94, 0x83, 0xfd, 0x9a, 0xad,
	0x0e, 0x3f, 0x1d, 0xbc, 0xae, 0xf2, 0xc5, 0xe5, 0xb2, 0xda, 0x39, 0xf5, 0x5d, 0x0f, 0x2b, 0x85,
	0xd0, 0xac, 0xd3, 0x41, 0x8f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x29, 0xf0, 0xa4, 0x64, 0x3d,
	0x01, 0x00, 0x00,
}
