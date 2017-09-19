// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/appengine/internal/types.proto

package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Trigger is task.Trigger equivalent.
type Trigger struct {
	Id           string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	JobId        string `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	InvocationId int64  `protobuf:"varint,3,opt,name=invocation_id,json=invocationId" json:"invocation_id,omitempty"`
	Created      int64  `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
	Title        string `protobuf:"bytes,5,opt,name=title" json:"title,omitempty"`
	Url          string `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	Payload      []byte `protobuf:"bytes,7,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Trigger) Reset()                    { *m = Trigger{} }
func (m *Trigger) String() string            { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()               {}
func (*Trigger) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Trigger) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Trigger) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Trigger) GetInvocationId() int64 {
	if m != nil {
		return m.InvocationId
	}
	return 0
}

func (m *Trigger) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Trigger) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Trigger) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Trigger) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Trigger)(nil), "internal.types.Trigger")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/appengine/internal/types.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xcf, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x06, 0x60, 0xd2, 0xda, 0x16, 0xc3, 0xba, 0x48, 0x50, 0xc8, 0xb1, 0xe8, 0xa5, 0xa7, 0xe6,
	0xe0, 0xd9, 0x07, 0xd8, 0x6b, 0xf1, 0x2e, 0x69, 0x32, 0x64, 0x67, 0xc9, 0x66, 0x42, 0x36, 0x15,
	0xfa, 0x58, 0xbe, 0xa1, 0x34, 0x52, 0xf6, 0x36, 0xff, 0xf7, 0xc3, 0x0f, 0xc3, 0x3f, 0x1d, 0x8d,
	0xe6, 0x9c, 0xe8, 0x8a, 0xcb, 0x75, 0xa4, 0xe4, 0x94, 0x5f, 0x0c, 0xaa, 0x9b, 0x39, 0x83, 0x5d,
	0x3c, 0x24, 0xa5, 0x63, 0x84, 0xe0, 0x30, 0x80, 0xc2, 0x90, 0x21, 0x05, 0xed, 0x55, 0x5e, 0x23,
	0xdc, 0xc6, 0x98, 0x28, 0x93, 0x38, 0xee, 0x3a, 0x16, 0x7d, 0xfb, 0x65, 0xbc, 0xfb, 0x4a, 0xe8,
	0x1c, 0x24, 0x71, 0xe4, 0x15, 0x5a, 0xc9, 0x7a, 0x36, 0x3c, 0x4e, 0x15, 0x5a, 0xf1, 0xca, 0xdb,
	0x0b, 0xcd, 0xdf, 0x68, 0x65, 0x55, 0xac, 0xb9, 0xd0, 0x7c, 0xb2, 0xe2, 0x9d, 0x3f, 0x61, 0xf8,
	0x21, 0xa3, 0x33, 0x52, 0xd8, 0xda, 0xba, 0x67, 0x43, 0x3d, 0x1d, 0xee, 0x78, 0xb2, 0x42, 0xf2,
	0xce, 0x24, 0xd0, 0x19, 0xac, 0x7c, 0x28, 0xf5, 0x1e, 0xc5, 0x0b, 0x6f, 0x32, 0x66, 0x0f, 0xb2,
	0xf9, 0x1f, 0x2d, 0x41, 0x3c, 0xf3, 0x7a, 0x49, 0x5e, 0xb6, 0xc5, 0xb6, 0x73, 0x5b, 0x88, 0x7a,
	0xf5, 0xa4, 0xad, 0xec, 0x7a, 0x36, 0x1c, 0xa6, 0x3d, 0xce, 0x6d, 0x79, 0xe5, 0xe3, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0xa5, 0x67, 0x12, 0xd4, 0x0b, 0x01, 0x00, 0x00,
}
