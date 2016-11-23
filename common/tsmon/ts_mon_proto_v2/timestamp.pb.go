// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/common/tsmon/ts_mon_proto_v2/timestamp.proto
// DO NOT EDIT!

package ts_mon_proto_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Timestamp struct {
	Seconds          *int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	Nanos            *int32 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Timestamp) GetSeconds() int64 {
	if m != nil && m.Seconds != nil {
		return *m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int32 {
	if m != nil && m.Nanos != nil {
		return *m.Nanos
	}
	return 0
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "ts_mon.proto.v2.Timestamp")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/common/tsmon/ts_mon_proto_v2/timestamp.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x72, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x29, 0x4d, 0xce, 0x04, 0x13, 0xba, 0xe9, 0xf9,
	0xfa, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0xfa, 0x25, 0xc5, 0x10, 0x32, 0x3e, 0x37, 0x3f, 0x2f,
	0xbe, 0xa0, 0x28, 0xbf, 0x24, 0x3f, 0xbe, 0xcc, 0x48, 0xbf, 0x24, 0x33, 0x37, 0xb5, 0xb8, 0x24,
	0x31, 0xb7, 0x40, 0x0f, 0x2c, 0x24, 0xc4, 0x0f, 0x51, 0x00, 0xe1, 0xe9, 0x95, 0x19, 0x29, 0x59,
	0x73, 0x71, 0x86, 0xc0, 0xd4, 0x08, 0x49, 0x70, 0xb1, 0x17, 0xa7, 0x26, 0xe7, 0xe7, 0xa5, 0x14,
	0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x79, 0x89, 0x79,
	0xf9, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x0e, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x04, 0x0a, 0x9b, 0x32, 0x94, 0x00, 0x00, 0x00,
}
