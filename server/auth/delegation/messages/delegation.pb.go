// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/server/auth/delegation/messages/delegation.proto

package messages

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

type Subtoken_Kind int32

const (
	// This is to catch old tokens that don't have 'kind' field yet.
	//
	// Tokens of this kind are interpreted as 'BEARER_DELEGATION_TOKEN' for now,
	// for compatibility. But eventually (when all backends are updated), they
	// will become invalid (and there will be no way to generate them). This is
	// needed to avoid old servers accidentally interpret tokens of kind != 0 as
	// BEARER_DELEGATION_TOKEN tokens.
	Subtoken_UNKNOWN_KIND Subtoken_Kind = 0
	// The token of this kind can be sent in X-Delegation-Token-V1 HTTP header.
	// The services will check all restrictions of the token, and will
	// authenticate requests as coming from 'delegated_identity'.
	Subtoken_BEARER_DELEGATION_TOKEN Subtoken_Kind = 1
)

var Subtoken_Kind_name = map[int32]string{
	0: "UNKNOWN_KIND",
	1: "BEARER_DELEGATION_TOKEN",
}

var Subtoken_Kind_value = map[string]int32{
	"UNKNOWN_KIND":            0,
	"BEARER_DELEGATION_TOKEN": 1,
}

func (x Subtoken_Kind) String() string {
	return proto.EnumName(Subtoken_Kind_name, int32(x))
}

func (Subtoken_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67dd1eacea24f38c, []int{1, 0}
}

// Signed serialized Subtoken.
//
// This message is just an envelope that carries the serialized Subtoken message
// and its signature.
//
// Next ID: 6.
type DelegationToken struct {
	// Identity of a service that signed this token.
	//
	// It can be a 'service:<app-id>' string or 'user:<service-account-email>'
	// string.
	//
	// In both cases the appropriate certificate store will be queried (via SSL)
	// for the public key to use for signature verification.
	SignerId string `protobuf:"bytes,2,opt,name=signer_id,json=signerId,proto3" json:"signer_id,omitempty"`
	// ID of a key used for making the signature.
	//
	// There can be multiple active keys at any moment in time: one used for new
	// signatures, and one being rotated out (but still valid for verification).
	//
	// The lifetime of the token indirectly depends on the lifetime of the signing
	// key, which is 24h. So delegation tokens can't live longer than 24h.
	SigningKeyId string `protobuf:"bytes,3,opt,name=signing_key_id,json=signingKeyId,proto3" json:"signing_key_id,omitempty"`
	// The signature: PKCS1_v1_5+SHA256(serialized_subtoken, signing_key_id).
	Pkcs1Sha256Sig []byte `protobuf:"bytes,4,opt,name=pkcs1_sha256_sig,json=pkcs1Sha256Sig,proto3" json:"pkcs1_sha256_sig,omitempty"`
	// Serialized Subtoken message. It's signature is stored in pkcs1_sha256_sig.
	SerializedSubtoken   []byte   `protobuf:"bytes,5,opt,name=serialized_subtoken,json=serializedSubtoken,proto3" json:"serialized_subtoken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelegationToken) Reset()         { *m = DelegationToken{} }
func (m *DelegationToken) String() string { return proto.CompactTextString(m) }
func (*DelegationToken) ProtoMessage()    {}
func (*DelegationToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_67dd1eacea24f38c, []int{0}
}

func (m *DelegationToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelegationToken.Unmarshal(m, b)
}
func (m *DelegationToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelegationToken.Marshal(b, m, deterministic)
}
func (m *DelegationToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegationToken.Merge(m, src)
}
func (m *DelegationToken) XXX_Size() int {
	return xxx_messageInfo_DelegationToken.Size(m)
}
func (m *DelegationToken) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegationToken.DiscardUnknown(m)
}

var xxx_messageInfo_DelegationToken proto.InternalMessageInfo

func (m *DelegationToken) GetSignerId() string {
	if m != nil {
		return m.SignerId
	}
	return ""
}

func (m *DelegationToken) GetSigningKeyId() string {
	if m != nil {
		return m.SigningKeyId
	}
	return ""
}

func (m *DelegationToken) GetPkcs1Sha256Sig() []byte {
	if m != nil {
		return m.Pkcs1Sha256Sig
	}
	return nil
}

func (m *DelegationToken) GetSerializedSubtoken() []byte {
	if m != nil {
		return m.SerializedSubtoken
	}
	return nil
}

// Identifies who delegates what authority to whom where.
//
// Next ID: 10.
type Subtoken struct {
	// What kind of token is this.
	//
	// Defines how it can be used. See comments for Kind enum.
	Kind Subtoken_Kind `protobuf:"varint,8,opt,name=kind,proto3,enum=messages.Subtoken_Kind" json:"kind,omitempty"`
	// Identifier of this subtoken as generated by the token server.
	//
	// Used for logging and tracking purposes.
	SubtokenId int64 `protobuf:"varint,4,opt,name=subtoken_id,json=subtokenId,proto3" json:"subtoken_id,omitempty"`
	// Identity whose authority is delegated.
	//
	// A string of the form "user:<email>".
	DelegatedIdentity string `protobuf:"bytes,1,opt,name=delegated_identity,json=delegatedIdentity,proto3" json:"delegated_identity,omitempty"`
	// Who requested this token.
	//
	// This can match delegated_identity if the user is delegating their own
	// identity or it can be a different id if the token is actually
	// an impersonation token.
	RequestorIdentity string `protobuf:"bytes,7,opt,name=requestor_identity,json=requestorIdentity,proto3" json:"requestor_identity,omitempty"`
	// When the token was generated (and when it becomes valid).
	//
	// Number of seconds since epoch (Unix timestamp).
	CreationTime int64 `protobuf:"varint,2,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	// How long the token is considered valid (in seconds).
	ValidityDuration int32 `protobuf:"varint,3,opt,name=validity_duration,json=validityDuration,proto3" json:"validity_duration,omitempty"`
	// Who can present this token.
	//
	// Each item can be an identity string (e.g. "user:<email>"), a "group:<name>"
	// string, or special "*" string which means "Any bearer can use the token".
	Audience []string `protobuf:"bytes,5,rep,name=audience,proto3" json:"audience,omitempty"`
	// What services should accept this token.
	//
	// List of services (specified as service identities, e.g. "service:app-id")
	// that should accept this token. May also contain special "*" string, which
	// means "All services".
	Services []string `protobuf:"bytes,6,rep,name=services,proto3" json:"services,omitempty"`
	// Arbitrary key:value pairs embedded into the token by whoever requested it.
	// Convey circumstance of why the token is created.
	//
	// Services that accept the token may use them for additional authorization
	// decisions. Please use extremely carefully, only when you control both sides
	// of the delegation link and can guarantee that services involved understand
	// the tags.
	Tags                 []string `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subtoken) Reset()         { *m = Subtoken{} }
func (m *Subtoken) String() string { return proto.CompactTextString(m) }
func (*Subtoken) ProtoMessage()    {}
func (*Subtoken) Descriptor() ([]byte, []int) {
	return fileDescriptor_67dd1eacea24f38c, []int{1}
}

func (m *Subtoken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subtoken.Unmarshal(m, b)
}
func (m *Subtoken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subtoken.Marshal(b, m, deterministic)
}
func (m *Subtoken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subtoken.Merge(m, src)
}
func (m *Subtoken) XXX_Size() int {
	return xxx_messageInfo_Subtoken.Size(m)
}
func (m *Subtoken) XXX_DiscardUnknown() {
	xxx_messageInfo_Subtoken.DiscardUnknown(m)
}

var xxx_messageInfo_Subtoken proto.InternalMessageInfo

func (m *Subtoken) GetKind() Subtoken_Kind {
	if m != nil {
		return m.Kind
	}
	return Subtoken_UNKNOWN_KIND
}

func (m *Subtoken) GetSubtokenId() int64 {
	if m != nil {
		return m.SubtokenId
	}
	return 0
}

func (m *Subtoken) GetDelegatedIdentity() string {
	if m != nil {
		return m.DelegatedIdentity
	}
	return ""
}

func (m *Subtoken) GetRequestorIdentity() string {
	if m != nil {
		return m.RequestorIdentity
	}
	return ""
}

func (m *Subtoken) GetCreationTime() int64 {
	if m != nil {
		return m.CreationTime
	}
	return 0
}

func (m *Subtoken) GetValidityDuration() int32 {
	if m != nil {
		return m.ValidityDuration
	}
	return 0
}

func (m *Subtoken) GetAudience() []string {
	if m != nil {
		return m.Audience
	}
	return nil
}

func (m *Subtoken) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Subtoken) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterEnum("messages.Subtoken_Kind", Subtoken_Kind_name, Subtoken_Kind_value)
	proto.RegisterType((*DelegationToken)(nil), "messages.DelegationToken")
	proto.RegisterType((*Subtoken)(nil), "messages.Subtoken")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/server/auth/delegation/messages/delegation.proto", fileDescriptor_67dd1eacea24f38c)
}

var fileDescriptor_67dd1eacea24f38c = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0xc9, 0x92, 0x8e, 0xd4, 0x94, 0x92, 0x99, 0x8b, 0x45, 0xec, 0x82, 0xaa, 0x70, 0x51,
	0x69, 0x22, 0x11, 0x43, 0xe3, 0x7e, 0xa8, 0x11, 0x0a, 0x41, 0xa9, 0x94, 0x16, 0x71, 0x69, 0x65,
	0xf1, 0x51, 0x7a, 0xd4, 0x26, 0x19, 0xb6, 0x33, 0xa9, 0xbc, 0x02, 0xcf, 0xc3, 0xfb, 0x21, 0xbb,
	0x71, 0xc7, 0x9d, 0xfd, 0x7d, 0xbf, 0x94, 0xe3, 0x3f, 0x87, 0x24, 0x75, 0x17, 0x55, 0x5b, 0xd1,
	0x35, 0xd8, 0x37, 0x51, 0x27, 0xea, 0x78, 0xdf, 0x57, 0x18, 0x4b, 0x10, 0x8f, 0x20, 0xe2, 0xb2,
	0x57, 0xdb, 0x98, 0xc3, 0x1e, 0xea, 0x52, 0x61, 0xd7, 0xc6, 0x0d, 0x48, 0x59, 0xd6, 0x20, 0xff,
	0x63, 0xd1, 0x83, 0xe8, 0x54, 0x47, 0x7d, 0xab, 0xe6, 0x7f, 0x1d, 0xf2, 0x6a, 0x79, 0xd2, 0x9b,
	0x6e, 0x07, 0x2d, 0xbd, 0x22, 0x63, 0x89, 0x75, 0x0b, 0x82, 0x21, 0x0f, 0xcf, 0x66, 0xce, 0x62,
	0x5c, 0xf8, 0x47, 0x90, 0x72, 0xfa, 0x9e, 0x4c, 0xf5, 0x19, 0xdb, 0x9a, 0xed, 0xe0, 0xa0, 0x13,
	0xae, 0x49, 0x4c, 0x06, 0x9a, 0xc1, 0x21, 0xe5, 0x74, 0x41, 0x82, 0x87, 0x5d, 0x25, 0x3f, 0x32,
	0xb9, 0x2d, 0x6f, 0x6e, 0x3f, 0x33, 0x89, 0x75, 0xe8, 0xcd, 0x9c, 0xc5, 0xa4, 0x98, 0x1a, 0xbe,
	0x36, 0x78, 0x8d, 0x35, 0x8d, 0xc9, 0x6b, 0x09, 0x02, 0xcb, 0x3d, 0xfe, 0x06, 0xce, 0x64, 0x7f,
	0xaf, 0xf4, 0x0c, 0xe1, 0xc8, 0x84, 0xe9, 0x93, 0x5a, 0x0f, 0xe6, 0x9b, 0xe7, 0x3b, 0xc1, 0xd9,
	0xfc, 0x8f, 0x4b, 0x7c, 0x8b, 0xe8, 0x35, 0xf1, 0x76, 0xd8, 0xf2, 0xd0, 0x9f, 0x39, 0x8b, 0xe9,
	0xcd, 0x65, 0x64, 0x5f, 0x17, 0xd9, 0x44, 0x94, 0x61, 0xcb, 0x0b, 0x13, 0xa2, 0x6f, 0xc9, 0x0b,
	0xfb, 0x15, 0x3d, 0xbd, 0x9e, 0xca, 0x2d, 0x88, 0x45, 0x29, 0xa7, 0x1f, 0x08, 0x1d, 0x0a, 0x03,
	0xce, 0x90, 0x43, 0xab, 0x50, 0x1d, 0x42, 0xc7, 0xbc, 0xf2, 0xe2, 0x64, 0xd2, 0x41, 0xe8, 0xb8,
	0x80, 0x5f, 0x3d, 0x48, 0xd5, 0x89, 0xa7, 0xf8, 0xf3, 0x63, 0xfc, 0x64, 0x4e, 0xf1, 0x77, 0xe4,
	0x65, 0x25, 0xc0, 0xb4, 0xcd, 0x14, 0x36, 0x60, 0x0a, 0x76, 0x8b, 0x89, 0x85, 0x1b, 0x6c, 0x80,
	0x5e, 0x93, 0x8b, 0xc7, 0x72, 0x8f, 0x1c, 0xd5, 0x81, 0xf1, 0x5e, 0x18, 0x61, 0x7a, 0x1e, 0x15,
	0x81, 0x15, 0xcb, 0x81, 0xd3, 0x37, 0xc4, 0x2f, 0x7b, 0x8e, 0xd0, 0x56, 0x10, 0x8e, 0x66, 0xae,
	0xfe, 0x5b, 0xf6, 0xae, 0x9d, 0x5e, 0x0e, 0xac, 0x40, 0x86, 0xe7, 0x47, 0x67, 0xef, 0x94, 0x12,
	0x4f, 0x95, 0xb5, 0x0c, 0xc7, 0x86, 0x9b, 0xf3, 0xfc, 0x96, 0x78, 0xba, 0x2a, 0x1a, 0x90, 0xc9,
	0x8f, 0x3c, 0xcb, 0x57, 0x3f, 0x73, 0x96, 0xa5, 0xf9, 0x32, 0x78, 0x46, 0xaf, 0xc8, 0xe5, 0x97,
	0xe4, 0xae, 0x48, 0x0a, 0xb6, 0x4c, 0xbe, 0x27, 0x5f, 0xef, 0x36, 0xe9, 0x2a, 0x67, 0x9b, 0x55,
	0x96, 0xe4, 0x81, 0x73, 0x7f, 0x6e, 0xd6, 0xea, 0xd3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09,
	0x07, 0xd9, 0xb4, 0x9f, 0x02, 0x00, 0x00,
}
