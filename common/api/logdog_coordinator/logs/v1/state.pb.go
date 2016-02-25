// Code generated by protoc-gen-go.
// source: state.proto
// DO NOT EDIT!

package logs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// LogStreamState is a bidirectional state value used in UpdateStream calls.
//
// LogStreamState is embeddable in Endpoints request/response structs.
type LogStreamState struct {
	// ProtoVersion is the protobuf version for this stream.
	ProtoVersion string `protobuf:"bytes,1,opt,name=proto_version" json:"proto_version,omitempty"`
	// The time when the log stream was registered with the Coordinator.
	Created *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
	// The time when the log stream's state was last updated.
	Updated *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=updated" json:"updated,omitempty"`
	// The stream index of the log stream's terminal message. If the value is -1,
	// the log is still streaming.
	TerminalIndex int64 `protobuf:"varint,4,opt,name=terminal_index" json:"terminal_index,omitempty"`
	// If non-nil, the log stream is archived, and this field contains archival
	// details.
	Archive *LogStreamState_ArchiveInfo `protobuf:"bytes,5,opt,name=archive" json:"archive,omitempty"`
	// Indicates the purged state of a log. A log that has been purged is only
	// acknowledged to administrative clients.
	Purged bool `protobuf:"varint,6,opt,name=purged" json:"purged,omitempty"`
}

func (m *LogStreamState) Reset()                    { *m = LogStreamState{} }
func (m *LogStreamState) String() string            { return proto.CompactTextString(m) }
func (*LogStreamState) ProtoMessage()               {}
func (*LogStreamState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *LogStreamState) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *LogStreamState) GetUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *LogStreamState) GetArchive() *LogStreamState_ArchiveInfo {
	if m != nil {
		return m.Archive
	}
	return nil
}

// ArchiveInfo contains archive details for the log stream.
type LogStreamState_ArchiveInfo struct {
	// The Google Storage URL where the log stream's index is archived.
	IndexUrl string `protobuf:"bytes,1,opt,name=index_url" json:"index_url,omitempty"`
	// The Google Storage URL where the log stream's raw stream data is archived.
	StreamUrl string `protobuf:"bytes,2,opt,name=stream_url" json:"stream_url,omitempty"`
	// The Google Storage URL where the log stream's assembled data is archived.
	DataUrl string `protobuf:"bytes,3,opt,name=data_url" json:"data_url,omitempty"`
	// If true, all log entries between 0 and terminal_index were archived. If
	// false, this indicates that the log stream was not completely loaded into
	// intermediate storage when the archival interval expired.
	Whole bool `protobuf:"varint,4,opt,name=whole" json:"whole,omitempty"`
}

func (m *LogStreamState_ArchiveInfo) Reset()                    { *m = LogStreamState_ArchiveInfo{} }
func (m *LogStreamState_ArchiveInfo) String() string            { return proto.CompactTextString(m) }
func (*LogStreamState_ArchiveInfo) ProtoMessage()               {}
func (*LogStreamState_ArchiveInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func init() {
	proto.RegisterType((*LogStreamState)(nil), "logs.LogStreamState")
	proto.RegisterType((*LogStreamState_ArchiveInfo)(nil), "logs.LogStreamState.ArchiveInfo")
}

var fileDescriptor1 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x69, 0xd3, 0xa6, 0xe9, 0x84, 0x06, 0x5d, 0x50, 0x42, 0x2e, 0x16, 0x4f, 0x82, 0xb0,
	0x45, 0x7d, 0x02, 0x8f, 0x82, 0xb7, 0xea, 0x39, 0x6c, 0x9b, 0xe9, 0x76, 0x61, 0x93, 0x09, 0x9b,
	0x4d, 0xf5, 0xbd, 0x7c, 0x41, 0x37, 0x13, 0x0b, 0x7a, 0xf2, 0xfa, 0xcd, 0x37, 0xff, 0xcc, 0x0f,
	0x69, 0xe7, 0x95, 0x47, 0xd9, 0x3a, 0xf2, 0x24, 0x66, 0x96, 0x74, 0x57, 0xdc, 0x68, 0x22, 0x6d,
	0x71, 0xc3, 0x6c, 0xd7, 0x1f, 0x36, 0xde, 0xd4, 0x18, 0xb4, 0xba, 0x1d, 0xb5, 0xdb, 0xaf, 0x29,
	0x64, 0xaf, 0xa4, 0xb7, 0xde, 0xa1, 0xaa, 0xb7, 0xc3, 0xbe, 0xb8, 0x82, 0x15, 0xcf, 0xca, 0x13,
	0xba, 0xce, 0x50, 0x93, 0x4f, 0xd6, 0x93, 0xbb, 0xa5, 0xb8, 0x87, 0xc5, 0x3e, 0x48, 0x1e, 0xab,
	0x7c, 0x1a, 0x40, 0xfa, 0x58, 0xc8, 0x31, 0x5c, 0x9e, 0xc3, 0xe5, 0xdb, 0x39, 0x7c, 0x90, 0xfb,
	0xb6, 0x62, 0x39, 0xfa, 0x57, 0xbe, 0x86, 0xcc, 0xa3, 0xab, 0x4d, 0xa3, 0x6c, 0x69, 0x9a, 0x0a,
	0x3f, 0xf3, 0x59, 0xd8, 0x89, 0xc4, 0x03, 0x2c, 0x94, 0xdb, 0x1f, 0xcd, 0x09, 0xf3, 0x39, 0x87,
	0xac, 0xe5, 0x50, 0x4a, 0xfe, 0xfd, 0x57, 0x3e, 0x8f, 0xce, 0x4b, 0x73, 0x20, 0x91, 0x41, 0xdc,
	0xf6, 0x4e, 0x87, 0xb3, 0x71, 0xd8, 0x48, 0x8a, 0x77, 0x48, 0x7f, 0x8f, 0x2f, 0x61, 0xc9, 0x07,
	0xca, 0xde, 0xd9, 0x9f, 0x5a, 0x02, 0xa0, 0xe3, 0x30, 0x66, 0x53, 0x66, 0x17, 0x90, 0x84, 0xdf,
	0x15, 0x93, 0x88, 0xc9, 0x0a, 0xe6, 0x1f, 0x47, 0xb2, 0xc8, 0x9f, 0x25, 0xbb, 0x98, 0x5b, 0x3c,
	0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x96, 0x53, 0x74, 0x72, 0x01, 0x00, 0x00,
}
