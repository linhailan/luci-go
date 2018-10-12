// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/common/proto/git/commit.proto

package git

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Commit_TreeDiff_ChangeType int32

const (
	Commit_TreeDiff_ADD    Commit_TreeDiff_ChangeType = 0
	Commit_TreeDiff_COPY   Commit_TreeDiff_ChangeType = 1
	Commit_TreeDiff_DELETE Commit_TreeDiff_ChangeType = 2
	Commit_TreeDiff_MODIFY Commit_TreeDiff_ChangeType = 3
	Commit_TreeDiff_RENAME Commit_TreeDiff_ChangeType = 4
)

var Commit_TreeDiff_ChangeType_name = map[int32]string{
	0: "ADD",
	1: "COPY",
	2: "DELETE",
	3: "MODIFY",
	4: "RENAME",
}

var Commit_TreeDiff_ChangeType_value = map[string]int32{
	"ADD":    0,
	"COPY":   1,
	"DELETE": 2,
	"MODIFY": 3,
	"RENAME": 4,
}

func (x Commit_TreeDiff_ChangeType) String() string {
	return proto.EnumName(Commit_TreeDiff_ChangeType_name, int32(x))
}

func (Commit_TreeDiff_ChangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d74b2aaad703343, []int{0, 1, 0}
}

// Commit is a single parsed commit as represented in a git log or git show
// expression.
type Commit struct {
	// The hex sha1 of the commit.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The hex sha1 of the tree for this commit.
	Tree string `protobuf:"bytes,2,opt,name=tree,proto3" json:"tree,omitempty"`
	// The hex sha1's of each of this commits' parents.
	Parents   []string     `protobuf:"bytes,3,rep,name=parents,proto3" json:"parents,omitempty"`
	Author    *Commit_User `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Committer *Commit_User `protobuf:"bytes,5,opt,name=committer,proto3" json:"committer,omitempty"`
	// This is the entire unaltered message body.
	Message              string             `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	TreeDiff             []*Commit_TreeDiff `protobuf:"bytes,7,rep,name=tree_diff,json=treeDiff,proto3" json:"tree_diff,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d74b2aaad703343, []int{0}
}

func (m *Commit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit.Unmarshal(m, b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
}
func (m *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(m, src)
}
func (m *Commit) XXX_Size() int {
	return xxx_messageInfo_Commit.Size(m)
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Commit) GetTree() string {
	if m != nil {
		return m.Tree
	}
	return ""
}

func (m *Commit) GetParents() []string {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *Commit) GetAuthor() *Commit_User {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Commit) GetCommitter() *Commit_User {
	if m != nil {
		return m.Committer
	}
	return nil
}

func (m *Commit) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Commit) GetTreeDiff() []*Commit_TreeDiff {
	if m != nil {
		return m.TreeDiff
	}
	return nil
}

// User represents the (name, email, timestamp) Commit header for author and/or
// commtter.
type Commit_User struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Commit_User) Reset()         { *m = Commit_User{} }
func (m *Commit_User) String() string { return proto.CompactTextString(m) }
func (*Commit_User) ProtoMessage()    {}
func (*Commit_User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d74b2aaad703343, []int{0, 0}
}

func (m *Commit_User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit_User.Unmarshal(m, b)
}
func (m *Commit_User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit_User.Marshal(b, m, deterministic)
}
func (m *Commit_User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit_User.Merge(m, src)
}
func (m *Commit_User) XXX_Size() int {
	return xxx_messageInfo_Commit_User.Size(m)
}
func (m *Commit_User) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit_User.DiscardUnknown(m)
}

var xxx_messageInfo_Commit_User proto.InternalMessageInfo

func (m *Commit_User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Commit_User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Commit_User) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

// Each TreeDiff represents a single file that's changed between this commit
// and the "previous" commit, where "previous" depends on the context of how
// this Commit object was produced (i.e. the specific `git log` invocation, or
// similar command).
//
// Note that these are an artifact of the `git log` expression, not of the
// commit itself (since git log has different ways that it could sort the
// commits in the log, and thus different ways it could calculate these
// diffs). In particular, you should avoid caching the TreeDiff data using
// only the Commit.id as the key.
//
// The old_* fields correspond to the matching file in the previous commit (in
// the case of COPY/DELETE/MODIFY/RENAME), telling its blob hash, file mode
// and path name.
//
// The new_* fields correspond to the matching file in this commit (in the
// case of ADD/COPY/MODIFY/RENAME), telling its blob hash, file mode and path
// name.
type Commit_TreeDiff struct {
	// How this file changed.
	Type                 Commit_TreeDiff_ChangeType `protobuf:"varint,1,opt,name=type,proto3,enum=git.Commit_TreeDiff_ChangeType" json:"type,omitempty"`
	OldId                string                     `protobuf:"bytes,2,opt,name=old_id,json=oldId,proto3" json:"old_id,omitempty"`
	OldMode              uint32                     `protobuf:"varint,3,opt,name=old_mode,json=oldMode,proto3" json:"old_mode,omitempty"`
	OldPath              string                     `protobuf:"bytes,4,opt,name=old_path,json=oldPath,proto3" json:"old_path,omitempty"`
	NewId                string                     `protobuf:"bytes,5,opt,name=new_id,json=newId,proto3" json:"new_id,omitempty"`
	NewMode              uint32                     `protobuf:"varint,6,opt,name=new_mode,json=newMode,proto3" json:"new_mode,omitempty"`
	NewPath              string                     `protobuf:"bytes,7,opt,name=new_path,json=newPath,proto3" json:"new_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Commit_TreeDiff) Reset()         { *m = Commit_TreeDiff{} }
func (m *Commit_TreeDiff) String() string { return proto.CompactTextString(m) }
func (*Commit_TreeDiff) ProtoMessage()    {}
func (*Commit_TreeDiff) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d74b2aaad703343, []int{0, 1}
}

func (m *Commit_TreeDiff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit_TreeDiff.Unmarshal(m, b)
}
func (m *Commit_TreeDiff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit_TreeDiff.Marshal(b, m, deterministic)
}
func (m *Commit_TreeDiff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit_TreeDiff.Merge(m, src)
}
func (m *Commit_TreeDiff) XXX_Size() int {
	return xxx_messageInfo_Commit_TreeDiff.Size(m)
}
func (m *Commit_TreeDiff) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit_TreeDiff.DiscardUnknown(m)
}

var xxx_messageInfo_Commit_TreeDiff proto.InternalMessageInfo

func (m *Commit_TreeDiff) GetType() Commit_TreeDiff_ChangeType {
	if m != nil {
		return m.Type
	}
	return Commit_TreeDiff_ADD
}

func (m *Commit_TreeDiff) GetOldId() string {
	if m != nil {
		return m.OldId
	}
	return ""
}

func (m *Commit_TreeDiff) GetOldMode() uint32 {
	if m != nil {
		return m.OldMode
	}
	return 0
}

func (m *Commit_TreeDiff) GetOldPath() string {
	if m != nil {
		return m.OldPath
	}
	return ""
}

func (m *Commit_TreeDiff) GetNewId() string {
	if m != nil {
		return m.NewId
	}
	return ""
}

func (m *Commit_TreeDiff) GetNewMode() uint32 {
	if m != nil {
		return m.NewMode
	}
	return 0
}

func (m *Commit_TreeDiff) GetNewPath() string {
	if m != nil {
		return m.NewPath
	}
	return ""
}

func init() {
	proto.RegisterEnum("git.Commit_TreeDiff_ChangeType", Commit_TreeDiff_ChangeType_name, Commit_TreeDiff_ChangeType_value)
	proto.RegisterType((*Commit)(nil), "git.Commit")
	proto.RegisterType((*Commit_User)(nil), "git.Commit.User")
	proto.RegisterType((*Commit_TreeDiff)(nil), "git.Commit.TreeDiff")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/common/proto/git/commit.proto", fileDescriptor_9d74b2aaad703343)
}

var fileDescriptor_9d74b2aaad703343 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x6e, 0x9c, 0x30,
	0x10, 0x86, 0xbb, 0x40, 0x60, 0x99, 0xa8, 0x11, 0xb2, 0x52, 0x89, 0xee, 0x25, 0xab, 0x9c, 0x38,
	0x19, 0x75, 0xf3, 0x04, 0xd1, 0x42, 0xa5, 0x48, 0xdd, 0x26, 0x42, 0xdb, 0x43, 0x4e, 0x29, 0x59,
	0x0f, 0x60, 0x09, 0x63, 0x64, 0xbc, 0x5a, 0xe5, 0x65, 0xfa, 0x8c, 0x7d, 0x84, 0xca, 0xf6, 0xa2,
	0xf4, 0xd0, 0xde, 0xc6, 0x33, 0xbf, 0xbf, 0xf9, 0xe7, 0x87, 0x4d, 0x2b, 0xe9, 0xa1, 0x53, 0x52,
	0xf0, 0xa3, 0xa0, 0x52, 0xb5, 0x79, 0x7f, 0x3c, 0xf0, 0xfc, 0x20, 0x85, 0x90, 0x43, 0x3e, 0x2a,
	0xa9, 0x65, 0xde, 0x72, 0x6d, 0x1b, 0x5c, 0x53, 0xdb, 0x20, 0x7e, 0xcb, 0xf5, 0xea, 0xa6, 0x95,
	0xb2, 0xed, 0xd1, 0x69, 0x5e, 0x8f, 0x4d, 0xae, 0xb9, 0xc0, 0x49, 0xd7, 0x62, 0x74, 0xaa, 0xdb,
	0xdf, 0x01, 0x84, 0x5b, 0xfb, 0x8d, 0x5c, 0x81, 0xc7, 0x59, 0xba, 0x58, 0x2f, 0xb2, 0xb8, 0xf2,
	0x38, 0x23, 0x04, 0x02, 0xad, 0x10, 0x53, 0xcf, 0x76, 0x6c, 0x4d, 0x52, 0x88, 0xc6, 0x5a, 0xe1,
	0xa0, 0xa7, 0xd4, 0x5f, 0xfb, 0x59, 0x5c, 0xcd, 0x4f, 0x92, 0x41, 0x58, 0x1f, 0x75, 0x27, 0x55,
	0x1a, 0xac, 0x17, 0xd9, 0xe5, 0x26, 0xa1, 0x2d, 0xd7, 0xd4, 0xa1, 0xe9, 0x8f, 0x09, 0x55, 0x75,
	0x9e, 0x13, 0x0a, 0xb1, 0x33, 0xaa, 0x51, 0xa5, 0x17, 0xff, 0x11, 0xbf, 0x4b, 0xcc, 0x4e, 0x81,
	0xd3, 0x54, 0xb7, 0x98, 0x86, 0xd6, 0xca, 0xfc, 0x24, 0x5f, 0x20, 0x36, 0xae, 0x5e, 0x18, 0x6f,
	0x9a, 0x34, 0x5a, 0xfb, 0xd9, 0xe5, 0xe6, 0xfa, 0x6f, 0xd2, 0x5e, 0x21, 0x16, 0xbc, 0x69, 0xaa,
	0xa5, 0x3e, 0x57, 0xab, 0x9f, 0x10, 0x18, 0xbe, 0x39, 0x6e, 0xa8, 0x05, 0x9e, 0xcf, 0xb5, 0x35,
	0xb9, 0x86, 0x0b, 0x14, 0x35, 0xef, 0xcf, 0x17, 0xbb, 0x07, 0xa1, 0x10, 0x98, 0xd0, 0x52, 0xdf,
	0x3a, 0x5d, 0x51, 0x97, 0x28, 0x9d, 0x13, 0xa5, 0xfb, 0x39, 0xd1, 0xca, 0xea, 0x56, 0xbf, 0x3c,
	0x58, 0xce, 0x8b, 0xc9, 0x1d, 0x04, 0xfa, 0x6d, 0x74, 0x6b, 0xae, 0x36, 0x37, 0xff, 0x32, 0x47,
	0xb7, 0x5d, 0x3d, 0xb4, 0xb8, 0x7f, 0x1b, 0xb1, 0xb2, 0x62, 0xf2, 0x09, 0x42, 0xd9, 0xb3, 0x17,
	0xce, 0x66, 0x23, 0xb2, 0x67, 0x0f, 0x8c, 0x7c, 0x86, 0xa5, 0x69, 0x0b, 0xc9, 0x9c, 0x99, 0x8f,
	0x55, 0x24, 0x7b, 0xb6, 0x93, 0x0c, 0xe7, 0xd1, 0x58, 0xeb, 0xce, 0xc6, 0x1f, 0xdb, 0xd1, 0x53,
	0xad, 0x3b, 0x03, 0x1b, 0xf0, 0x64, 0x60, 0x17, 0x0e, 0x36, 0xe0, 0xc9, 0xc1, 0x4c, 0xdb, 0xc2,
	0x42, 0x07, 0x1b, 0xf0, 0x34, 0xc3, 0xcc, 0xc8, 0xc2, 0x22, 0x07, 0x1b, 0xf0, 0x64, 0x60, 0xb7,
	0x5b, 0x80, 0x77, 0xb7, 0x24, 0x02, 0xff, 0xbe, 0x28, 0x92, 0x0f, 0x64, 0x09, 0xc1, 0xf6, 0xf1,
	0xe9, 0x39, 0x59, 0x10, 0x80, 0xb0, 0x28, 0xbf, 0x95, 0xfb, 0x32, 0xf1, 0x4c, 0xbd, 0x7b, 0x2c,
	0x1e, 0xbe, 0x3e, 0x27, 0xbe, 0xa9, 0xab, 0xf2, 0xfb, 0xfd, 0xae, 0x4c, 0x82, 0xd7, 0xd0, 0x46,
	0x77, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xc7, 0x7d, 0x8d, 0xd5, 0x02, 0x00, 0x00,
}
