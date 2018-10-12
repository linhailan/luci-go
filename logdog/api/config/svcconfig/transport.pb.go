// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/logdog/api/config/svcconfig/transport.proto

package svcconfig

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

// Transport is the transport configuration.
type Transport struct {
	// Type is the transport configuration that is being used.
	//
	// Types that are valid to be assigned to Type:
	//	*Transport_Pubsub
	Type                 isTransport_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Transport) Reset()         { *m = Transport{} }
func (m *Transport) String() string { return proto.CompactTextString(m) }
func (*Transport) ProtoMessage()    {}
func (*Transport) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf4c1f6cd2b2e8c, []int{0}
}

func (m *Transport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transport.Unmarshal(m, b)
}
func (m *Transport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transport.Marshal(b, m, deterministic)
}
func (m *Transport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transport.Merge(m, src)
}
func (m *Transport) XXX_Size() int {
	return xxx_messageInfo_Transport.Size(m)
}
func (m *Transport) XXX_DiscardUnknown() {
	xxx_messageInfo_Transport.DiscardUnknown(m)
}

var xxx_messageInfo_Transport proto.InternalMessageInfo

type isTransport_Type interface {
	isTransport_Type()
}

type Transport_Pubsub struct {
	Pubsub *Transport_PubSub `protobuf:"bytes,1,opt,name=pubsub,proto3,oneof"`
}

func (*Transport_Pubsub) isTransport_Type() {}

func (m *Transport) GetType() isTransport_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Transport) GetPubsub() *Transport_PubSub {
	if x, ok := m.GetType().(*Transport_Pubsub); ok {
		return x.Pubsub
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Transport) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Transport_OneofMarshaler, _Transport_OneofUnmarshaler, _Transport_OneofSizer, []interface{}{
		(*Transport_Pubsub)(nil),
	}
}

func _Transport_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Transport)
	// Type
	switch x := m.Type.(type) {
	case *Transport_Pubsub:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pubsub); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Transport.Type has unexpected type %T", x)
	}
	return nil
}

func _Transport_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Transport)
	switch tag {
	case 1: // Type.pubsub
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transport_PubSub)
		err := b.DecodeMessage(msg)
		m.Type = &Transport_Pubsub{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Transport_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Transport)
	// Type
	switch x := m.Type.(type) {
	case *Transport_Pubsub:
		s := proto.Size(x.Pubsub)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// PubSub is a transport configuration for Google Cloud Pub/Sub.
type Transport_PubSub struct {
	// The name of the authentication group for administrators.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The name of the authentication group for administrators.
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	// The name of the authentication group for administrators.
	Subscription         string   `protobuf:"bytes,3,opt,name=subscription,proto3" json:"subscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transport_PubSub) Reset()         { *m = Transport_PubSub{} }
func (m *Transport_PubSub) String() string { return proto.CompactTextString(m) }
func (*Transport_PubSub) ProtoMessage()    {}
func (*Transport_PubSub) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf4c1f6cd2b2e8c, []int{0, 0}
}

func (m *Transport_PubSub) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transport_PubSub.Unmarshal(m, b)
}
func (m *Transport_PubSub) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transport_PubSub.Marshal(b, m, deterministic)
}
func (m *Transport_PubSub) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transport_PubSub.Merge(m, src)
}
func (m *Transport_PubSub) XXX_Size() int {
	return xxx_messageInfo_Transport_PubSub.Size(m)
}
func (m *Transport_PubSub) XXX_DiscardUnknown() {
	xxx_messageInfo_Transport_PubSub.DiscardUnknown(m)
}

var xxx_messageInfo_Transport_PubSub proto.InternalMessageInfo

func (m *Transport_PubSub) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Transport_PubSub) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Transport_PubSub) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

func init() {
	proto.RegisterType((*Transport)(nil), "svcconfig.Transport")
	proto.RegisterType((*Transport_PubSub)(nil), "svcconfig.Transport.PubSub")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/logdog/api/config/svcconfig/transport.proto", fileDescriptor_abf4c1f6cd2b2e8c)
}

var fileDescriptor_abf4c1f6cd2b2e8c = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0x3b, 0x4f, 0xc3, 0x30,
	0x14, 0x46, 0x09, 0x0f, 0xa3, 0x5c, 0x98, 0x2c, 0x06, 0x0b, 0x16, 0x94, 0x89, 0xc9, 0x96, 0x40,
	0xec, 0x88, 0x89, 0x11, 0x85, 0x8c, 0x2c, 0xb1, 0x9b, 0xba, 0xae, 0x92, 0xdc, 0x2b, 0x3f, 0x2a,
	0xf5, 0x1f, 0xf5, 0x67, 0x56, 0x72, 0x1e, 0x52, 0x37, 0x7f, 0xe7, 0x1c, 0x59, 0x17, 0xbe, 0x2c,
	0x4a, 0xb3, 0xf3, 0x38, 0xb8, 0x34, 0x48, 0xf4, 0x56, 0xf5, 0xc9, 0x38, 0xd5, 0xa3, 0xdd, 0xa0,
	0x55, 0x2d, 0x39, 0x65, 0x70, 0xdc, 0x3a, 0xab, 0xc2, 0xc1, 0xcc, 0xaf, 0xe8, 0xdb, 0x31, 0x10,
	0xfa, 0x28, 0xc9, 0x63, 0x44, 0x5e, 0xae, 0xaa, 0x3a, 0x15, 0x50, 0x36, 0x8b, 0xe6, 0x9f, 0xc0,
	0x28, 0xe9, 0x90, 0xb4, 0x28, 0x5e, 0x8b, 0xb7, 0x87, 0xf7, 0x17, 0xb9, 0x96, 0x72, 0xad, 0xe4,
	0x6f, 0xd2, 0x7f, 0x49, 0xff, 0x5c, 0xd5, 0x73, 0xfc, 0xfc, 0x0f, 0x6c, 0x62, 0x5c, 0xc0, 0x3d,
	0x79, 0xdc, 0x77, 0x26, 0xe6, 0x1f, 0xca, 0x7a, 0x99, 0xfc, 0x09, 0xee, 0x22, 0x92, 0x33, 0xe2,
	0x3a, 0xf3, 0x69, 0xf0, 0x0a, 0x1e, 0x43, 0xd2, 0xc1, 0x78, 0x47, 0xd1, 0xe1, 0x28, 0x6e, 0xb2,
	0xbc, 0x60, 0xdf, 0x0c, 0x6e, 0x9b, 0x23, 0x75, 0x9a, 0xe5, 0xe3, 0x3f, 0xce, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x87, 0xa9, 0x2f, 0x29, 0x00, 0x01, 0x00, 0x00,
}
