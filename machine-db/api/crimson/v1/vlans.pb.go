// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/vlans.proto

package crimson

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "go.chromium.org/luci/machine-db/api/common/v1"
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

// A VLAN in the database.
type VLAN struct {
	// The ID of this VLAN. Uniquely identifies this VLAN.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// An alias for this VLAN.
	Alias string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	// A description of this VLAN.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The state of this VLAN.
	State v1.State `protobuf:"varint,4,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	// The block of IPv4 addresses belonging to this VLAN.
	CidrBlock            string   `protobuf:"bytes,5,opt,name=cidr_block,json=cidrBlock,proto3" json:"cidr_block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VLAN) Reset()         { *m = VLAN{} }
func (m *VLAN) String() string { return proto.CompactTextString(m) }
func (*VLAN) ProtoMessage()    {}
func (*VLAN) Descriptor() ([]byte, []int) {
	return fileDescriptor_674503a3aa3cd9ab, []int{0}
}

func (m *VLAN) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VLAN.Unmarshal(m, b)
}
func (m *VLAN) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VLAN.Marshal(b, m, deterministic)
}
func (m *VLAN) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VLAN.Merge(m, src)
}
func (m *VLAN) XXX_Size() int {
	return xxx_messageInfo_VLAN.Size(m)
}
func (m *VLAN) XXX_DiscardUnknown() {
	xxx_messageInfo_VLAN.DiscardUnknown(m)
}

var xxx_messageInfo_VLAN proto.InternalMessageInfo

func (m *VLAN) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VLAN) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *VLAN) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *VLAN) GetState() v1.State {
	if m != nil {
		return m.State
	}
	return v1.State_STATE_UNSPECIFIED
}

func (m *VLAN) GetCidrBlock() string {
	if m != nil {
		return m.CidrBlock
	}
	return ""
}

// A request to list VLANs in the database.
type ListVLANsRequest struct {
	// The IDs of VLANs to retrieve.
	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	// The aliases of VLANs to retrieve.
	Aliases              []string `protobuf:"bytes,2,rep,name=aliases,proto3" json:"aliases,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListVLANsRequest) Reset()         { *m = ListVLANsRequest{} }
func (m *ListVLANsRequest) String() string { return proto.CompactTextString(m) }
func (*ListVLANsRequest) ProtoMessage()    {}
func (*ListVLANsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_674503a3aa3cd9ab, []int{1}
}

func (m *ListVLANsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVLANsRequest.Unmarshal(m, b)
}
func (m *ListVLANsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVLANsRequest.Marshal(b, m, deterministic)
}
func (m *ListVLANsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVLANsRequest.Merge(m, src)
}
func (m *ListVLANsRequest) XXX_Size() int {
	return xxx_messageInfo_ListVLANsRequest.Size(m)
}
func (m *ListVLANsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVLANsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListVLANsRequest proto.InternalMessageInfo

func (m *ListVLANsRequest) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ListVLANsRequest) GetAliases() []string {
	if m != nil {
		return m.Aliases
	}
	return nil
}

// A response containing a list of VLANs in the database.
type ListVLANsResponse struct {
	// The VLANs matching the request.
	Vlans                []*VLAN  `protobuf:"bytes,1,rep,name=vlans,proto3" json:"vlans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListVLANsResponse) Reset()         { *m = ListVLANsResponse{} }
func (m *ListVLANsResponse) String() string { return proto.CompactTextString(m) }
func (*ListVLANsResponse) ProtoMessage()    {}
func (*ListVLANsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_674503a3aa3cd9ab, []int{2}
}

func (m *ListVLANsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVLANsResponse.Unmarshal(m, b)
}
func (m *ListVLANsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVLANsResponse.Marshal(b, m, deterministic)
}
func (m *ListVLANsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVLANsResponse.Merge(m, src)
}
func (m *ListVLANsResponse) XXX_Size() int {
	return xxx_messageInfo_ListVLANsResponse.Size(m)
}
func (m *ListVLANsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVLANsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListVLANsResponse proto.InternalMessageInfo

func (m *ListVLANsResponse) GetVlans() []*VLAN {
	if m != nil {
		return m.Vlans
	}
	return nil
}

func init() {
	proto.RegisterType((*VLAN)(nil), "crimson.VLAN")
	proto.RegisterType((*ListVLANsRequest)(nil), "crimson.ListVLANsRequest")
	proto.RegisterType((*ListVLANsResponse)(nil), "crimson.ListVLANsResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/vlans.proto", fileDescriptor_674503a3aa3cd9ab)
}

var fileDescriptor_674503a3aa3cd9ab = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xd9, 0x6e, 0x6b, 0xe9, 0x14, 0x4b, 0x0d, 0x1e, 0x82, 0x20, 0x2c, 0xf5, 0xb2, 0x17,
	0x13, 0xd4, 0x8b, 0x78, 0x10, 0xf4, 0x5c, 0x3c, 0x44, 0xf0, 0x2a, 0x69, 0x12, 0xda, 0xc1, 0xcd,
	0x66, 0x4d, 0xd2, 0xfe, 0x10, 0x7f, 0xb1, 0x24, 0xa9, 0xd0, 0xa3, 0xb7, 0xcc, 0x7b, 0x99, 0x37,
	0x1f, 0x0f, 0x9e, 0xb6, 0x8e, 0xa9, 0x9d, 0x77, 0x16, 0xf7, 0x96, 0x39, 0xbf, 0xe5, 0xdd, 0x5e,
	0x21, 0xb7, 0x52, 0xed, 0xb0, 0x37, 0xb7, 0x7a, 0xc3, 0xe5, 0x80, 0x5c, 0x79, 0xb4, 0xc1, 0xf5,
	0xfc, 0x70, 0xc7, 0x0f, 0x9d, 0xec, 0x03, 0x1b, 0xbc, 0x8b, 0x8e, 0x4c, 0x8f, 0xfa, 0xd5, 0xff,
	0x42, 0x9c, 0xb5, 0x25, 0x23, 0x44, 0x19, 0xcd, 0x31, 0x64, 0xf5, 0x53, 0xc1, 0xf8, 0x63, 0xfd,
	0xf2, 0x46, 0x16, 0x30, 0x42, 0x4d, 0xab, 0xa6, 0x6a, 0x6b, 0x31, 0x42, 0x4d, 0x2e, 0x61, 0x22,
	0x3b, 0x94, 0x81, 0x8e, 0x9a, 0xaa, 0x9d, 0x89, 0x32, 0x90, 0x06, 0xe6, 0xda, 0x04, 0xe5, 0x71,
	0x88, 0xe8, 0x7a, 0x5a, 0x67, 0xef, 0x54, 0x22, 0x37, 0x30, 0xc9, 0x07, 0xe8, 0xb8, 0xa9, 0xda,
	0xc5, 0xfd, 0x39, 0x2b, 0x87, 0xd9, 0x7b, 0x12, 0x45, 0xf1, 0xc8, 0x35, 0x80, 0x42, 0xed, 0x3f,
	0x37, 0x9d, 0x53, 0x5f, 0x74, 0x92, 0x53, 0x66, 0x49, 0x79, 0x4d, 0xc2, 0xea, 0x19, 0x96, 0x6b,
	0x0c, 0x31, 0x71, 0x05, 0x61, 0xbe, 0xf7, 0x26, 0x44, 0xb2, 0x84, 0x1a, 0x75, 0xa0, 0x55, 0x53,
	0xb7, 0xb5, 0x48, 0x4f, 0x42, 0x61, 0x9a, 0xa1, 0x4c, 0x62, 0xac, 0xdb, 0x99, 0xf8, 0x1b, 0x57,
	0x8f, 0x70, 0x71, 0xb2, 0x1f, 0x06, 0xd7, 0x07, 0x93, 0xc0, 0x72, 0x7b, 0x39, 0x62, 0x9e, 0xc0,
	0x4a, 0x7d, 0x2c, 0x7d, 0x13, 0xc5, 0xdb, 0x9c, 0xe5, 0x56, 0x1e, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0x6e, 0xca, 0xc5, 0x98, 0x01, 0x00, 0x00,
}
