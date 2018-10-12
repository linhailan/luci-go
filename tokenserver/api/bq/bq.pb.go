// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/tokenserver/api/bq/bq.proto

package bq

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	messages "go.chromium.org/luci/server/auth/delegation/messages"
	api "go.chromium.org/luci/tokenserver/api"
	v1 "go.chromium.org/luci/tokenserver/api/minter/v1"
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

// Issued delegation tokens.
type DelegationToken struct {
	// First 16 bytes of SHA256 of the token body, hex-encoded.
	Fingerprint string `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Kind of the token.
	TokenKind messages.Subtoken_Kind `protobuf:"varint,2,opt,name=token_kind,json=tokenKind,proto3,enum=messages.Subtoken_Kind" json:"token_kind,omitempty"`
	// Identifier of this token as generated by the token server.
	TokenId string `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// Identity whose authority is delegated.
	DelegatedIdentity string `protobuf:"bytes,4,opt,name=delegated_identity,json=delegatedIdentity,proto3" json:"delegated_identity,omitempty"`
	// Who requested this token.
	RequestorIdentity string `protobuf:"bytes,5,opt,name=requestor_identity,json=requestorIdentity,proto3" json:"requestor_identity,omitempty"`
	// When the token was generated.
	IssuedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	// When the token expires.
	Expiration *timestamp.Timestamp `protobuf:"bytes,7,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// Who can present this token.
	TargetAudience []string `protobuf:"bytes,8,rep,name=target_audience,json=targetAudience,proto3" json:"target_audience,omitempty"`
	// What services should accept this token
	TargetServices []string `protobuf:"bytes,9,rep,name=target_services,json=targetServices,proto3" json:"target_services,omitempty"`
	// Token validity duration (in seconds), as requested by the caller.
	RequestedValidity int64 `protobuf:"varint,10,opt,name=requested_validity,json=requestedValidity,proto3" json:"requested_validity,omitempty"`
	// An intent string provided by the caller.
	RequestedIntent string `protobuf:"bytes,11,opt,name=requested_intent,json=requestedIntent,proto3" json:"requested_intent,omitempty"`
	// Arbitrary key:value pairs embedded into the token.
	Tags []string `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	// Revision of the luci-config repo with rules.
	ConfigRev string `protobuf:"bytes,13,opt,name=config_rev,json=configRev,proto3" json:"config_rev,omitempty"`
	// Name of the rule used to authorize this call.
	ConfigRule string `protobuf:"bytes,14,opt,name=config_rule,json=configRule,proto3" json:"config_rule,omitempty"`
	// IP address of the caller.
	PeerIp string `protobuf:"bytes,15,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	// Identifier of the token server GAE app and version.
	ServiceVersion string `protobuf:"bytes,16,opt,name=service_version,json=serviceVersion,proto3" json:"service_version,omitempty"`
	// ID of the GAE request that handled the call.
	GaeRequestId string `protobuf:"bytes,17,opt,name=gae_request_id,json=gaeRequestId,proto3" json:"gae_request_id,omitempty"`
	// Revision of groups database used to authorize this call.
	AuthDbRev            int64    `protobuf:"varint,18,opt,name=auth_db_rev,json=authDbRev,proto3" json:"auth_db_rev,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelegationToken) Reset()         { *m = DelegationToken{} }
func (m *DelegationToken) String() string { return proto.CompactTextString(m) }
func (*DelegationToken) ProtoMessage()    {}
func (*DelegationToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0cc6154e6d463f7, []int{0}
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

func (m *DelegationToken) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *DelegationToken) GetTokenKind() messages.Subtoken_Kind {
	if m != nil {
		return m.TokenKind
	}
	return messages.Subtoken_UNKNOWN_KIND
}

func (m *DelegationToken) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *DelegationToken) GetDelegatedIdentity() string {
	if m != nil {
		return m.DelegatedIdentity
	}
	return ""
}

func (m *DelegationToken) GetRequestorIdentity() string {
	if m != nil {
		return m.RequestorIdentity
	}
	return ""
}

func (m *DelegationToken) GetIssuedAt() *timestamp.Timestamp {
	if m != nil {
		return m.IssuedAt
	}
	return nil
}

func (m *DelegationToken) GetExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *DelegationToken) GetTargetAudience() []string {
	if m != nil {
		return m.TargetAudience
	}
	return nil
}

func (m *DelegationToken) GetTargetServices() []string {
	if m != nil {
		return m.TargetServices
	}
	return nil
}

func (m *DelegationToken) GetRequestedValidity() int64 {
	if m != nil {
		return m.RequestedValidity
	}
	return 0
}

func (m *DelegationToken) GetRequestedIntent() string {
	if m != nil {
		return m.RequestedIntent
	}
	return ""
}

func (m *DelegationToken) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DelegationToken) GetConfigRev() string {
	if m != nil {
		return m.ConfigRev
	}
	return ""
}

func (m *DelegationToken) GetConfigRule() string {
	if m != nil {
		return m.ConfigRule
	}
	return ""
}

func (m *DelegationToken) GetPeerIp() string {
	if m != nil {
		return m.PeerIp
	}
	return ""
}

func (m *DelegationToken) GetServiceVersion() string {
	if m != nil {
		return m.ServiceVersion
	}
	return ""
}

func (m *DelegationToken) GetGaeRequestId() string {
	if m != nil {
		return m.GaeRequestId
	}
	return ""
}

func (m *DelegationToken) GetAuthDbRev() int64 {
	if m != nil {
		return m.AuthDbRev
	}
	return 0
}

// Issued machine tokens.
type MachineToken struct {
	// First 16 bytes of SHA256 of the token body, hex-encoded.
	Fingerprint string `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Machine domain name encoded in the token.
	MachineFqdn string `protobuf:"bytes,2,opt,name=machine_fqdn,json=machineFqdn,proto3" json:"machine_fqdn,omitempty"`
	// Type of the machine token.
	TokenType api.MachineTokenType `protobuf:"varint,3,opt,name=token_type,json=tokenType,proto3,enum=tokenserver.MachineTokenType" json:"token_type,omitempty"`
	// When the token was generated.
	IssuedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	// When the token expires.
	Expiration *timestamp.Timestamp `protobuf:"bytes,5,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// Serial number of the peer certificate.
	CertSerialNumber string `protobuf:"bytes,6,opt,name=cert_serial_number,json=certSerialNumber,proto3" json:"cert_serial_number,omitempty"`
	// Type of the signature used to proof possession of the private key.
	SignatureAlgorithm v1.SignatureAlgorithm `protobuf:"varint,7,opt,name=signature_algorithm,json=signatureAlgorithm,proto3,enum=tokenserver.minter.SignatureAlgorithm" json:"signature_algorithm,omitempty"`
	// Common Name of a CA that signed the peer certificate.
	CaCommonName string `protobuf:"bytes,8,opt,name=ca_common_name,json=caCommonName,proto3" json:"ca_common_name,omitempty"`
	// Revision of the luci-config repo that contains the CA.
	CaConfigRev string `protobuf:"bytes,9,opt,name=ca_config_rev,json=caConfigRev,proto3" json:"ca_config_rev,omitempty"`
	// IP address of the caller.
	PeerIp string `protobuf:"bytes,10,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	// Identifier of the token server GAE app and version.
	ServiceVersion string `protobuf:"bytes,11,opt,name=service_version,json=serviceVersion,proto3" json:"service_version,omitempty"`
	// ID of the GAE request that handled the call.
	GaeRequestId         string   `protobuf:"bytes,12,opt,name=gae_request_id,json=gaeRequestId,proto3" json:"gae_request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MachineToken) Reset()         { *m = MachineToken{} }
func (m *MachineToken) String() string { return proto.CompactTextString(m) }
func (*MachineToken) ProtoMessage()    {}
func (*MachineToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0cc6154e6d463f7, []int{1}
}

func (m *MachineToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MachineToken.Unmarshal(m, b)
}
func (m *MachineToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MachineToken.Marshal(b, m, deterministic)
}
func (m *MachineToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineToken.Merge(m, src)
}
func (m *MachineToken) XXX_Size() int {
	return xxx_messageInfo_MachineToken.Size(m)
}
func (m *MachineToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineToken.DiscardUnknown(m)
}

var xxx_messageInfo_MachineToken proto.InternalMessageInfo

func (m *MachineToken) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *MachineToken) GetMachineFqdn() string {
	if m != nil {
		return m.MachineFqdn
	}
	return ""
}

func (m *MachineToken) GetTokenType() api.MachineTokenType {
	if m != nil {
		return m.TokenType
	}
	return api.MachineTokenType_UNKNOWN_TYPE
}

func (m *MachineToken) GetIssuedAt() *timestamp.Timestamp {
	if m != nil {
		return m.IssuedAt
	}
	return nil
}

func (m *MachineToken) GetExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *MachineToken) GetCertSerialNumber() string {
	if m != nil {
		return m.CertSerialNumber
	}
	return ""
}

func (m *MachineToken) GetSignatureAlgorithm() v1.SignatureAlgorithm {
	if m != nil {
		return m.SignatureAlgorithm
	}
	return v1.SignatureAlgorithm_UNKNOWN_ALGO
}

func (m *MachineToken) GetCaCommonName() string {
	if m != nil {
		return m.CaCommonName
	}
	return ""
}

func (m *MachineToken) GetCaConfigRev() string {
	if m != nil {
		return m.CaConfigRev
	}
	return ""
}

func (m *MachineToken) GetPeerIp() string {
	if m != nil {
		return m.PeerIp
	}
	return ""
}

func (m *MachineToken) GetServiceVersion() string {
	if m != nil {
		return m.ServiceVersion
	}
	return ""
}

func (m *MachineToken) GetGaeRequestId() string {
	if m != nil {
		return m.GaeRequestId
	}
	return ""
}

// Issued OAuth token grants.
type OAuthTokenGrant struct {
	// First 16 bytes of SHA256 of the token body, hex-encoded.
	Fingerprint string `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Identifier of this token as generated by the token server.
	TokenId string `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// Service account email the end user wants to act as.
	ServiceAccount string `protobuf:"bytes,3,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// Who requested and who can use this token.
	ProxyIdentity string `protobuf:"bytes,4,opt,name=proxy_identity,json=proxyIdentity,proto3" json:"proxy_identity,omitempty"`
	// On whose behalf the token is generated.
	EndUserIdentity string `protobuf:"bytes,5,opt,name=end_user_identity,json=endUserIdentity,proto3" json:"end_user_identity,omitempty"`
	// When the token was generated.
	IssuedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	// When the token expires.
	Expiration *timestamp.Timestamp `protobuf:"bytes,7,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// Arbitrary key:value pairs provided by the caller.
	AuditTags []string `protobuf:"bytes,8,rep,name=audit_tags,json=auditTags,proto3" json:"audit_tags,omitempty"`
	// Revision of the luci-config repo with rules.
	ConfigRev string `protobuf:"bytes,9,opt,name=config_rev,json=configRev,proto3" json:"config_rev,omitempty"`
	// Name of the rule used to authorize this call.
	ConfigRule string `protobuf:"bytes,10,opt,name=config_rule,json=configRule,proto3" json:"config_rule,omitempty"`
	// IP address of the caller.
	PeerIp string `protobuf:"bytes,11,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	// Identifier of the token server GAE app and version.
	ServiceVersion string `protobuf:"bytes,12,opt,name=service_version,json=serviceVersion,proto3" json:"service_version,omitempty"`
	// ID of the GAE request that handled the call.
	GaeRequestId string `protobuf:"bytes,13,opt,name=gae_request_id,json=gaeRequestId,proto3" json:"gae_request_id,omitempty"`
	// Revision of groups database used to authorize this call.
	AuthDbRev            int64    `protobuf:"varint,14,opt,name=auth_db_rev,json=authDbRev,proto3" json:"auth_db_rev,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthTokenGrant) Reset()         { *m = OAuthTokenGrant{} }
func (m *OAuthTokenGrant) String() string { return proto.CompactTextString(m) }
func (*OAuthTokenGrant) ProtoMessage()    {}
func (*OAuthTokenGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0cc6154e6d463f7, []int{2}
}

func (m *OAuthTokenGrant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthTokenGrant.Unmarshal(m, b)
}
func (m *OAuthTokenGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthTokenGrant.Marshal(b, m, deterministic)
}
func (m *OAuthTokenGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthTokenGrant.Merge(m, src)
}
func (m *OAuthTokenGrant) XXX_Size() int {
	return xxx_messageInfo_OAuthTokenGrant.Size(m)
}
func (m *OAuthTokenGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthTokenGrant.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthTokenGrant proto.InternalMessageInfo

func (m *OAuthTokenGrant) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *OAuthTokenGrant) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *OAuthTokenGrant) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func (m *OAuthTokenGrant) GetProxyIdentity() string {
	if m != nil {
		return m.ProxyIdentity
	}
	return ""
}

func (m *OAuthTokenGrant) GetEndUserIdentity() string {
	if m != nil {
		return m.EndUserIdentity
	}
	return ""
}

func (m *OAuthTokenGrant) GetIssuedAt() *timestamp.Timestamp {
	if m != nil {
		return m.IssuedAt
	}
	return nil
}

func (m *OAuthTokenGrant) GetExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *OAuthTokenGrant) GetAuditTags() []string {
	if m != nil {
		return m.AuditTags
	}
	return nil
}

func (m *OAuthTokenGrant) GetConfigRev() string {
	if m != nil {
		return m.ConfigRev
	}
	return ""
}

func (m *OAuthTokenGrant) GetConfigRule() string {
	if m != nil {
		return m.ConfigRule
	}
	return ""
}

func (m *OAuthTokenGrant) GetPeerIp() string {
	if m != nil {
		return m.PeerIp
	}
	return ""
}

func (m *OAuthTokenGrant) GetServiceVersion() string {
	if m != nil {
		return m.ServiceVersion
	}
	return ""
}

func (m *OAuthTokenGrant) GetGaeRequestId() string {
	if m != nil {
		return m.GaeRequestId
	}
	return ""
}

func (m *OAuthTokenGrant) GetAuthDbRev() int64 {
	if m != nil {
		return m.AuthDbRev
	}
	return 0
}

// Issued OAuth tokens.
type OAuthToken struct {
	// First 16 bytes of SHA256 of the token body, hex-encoded.
	Fingerprint string `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// First 16 bytes of SHA256 of the oauth grant, hex-encoded.
	GrantFingerprint string `protobuf:"bytes,2,opt,name=grant_fingerprint,json=grantFingerprint,proto3" json:"grant_fingerprint,omitempty"`
	// Service account email the end user wants to act as.
	ServiceAccount string `protobuf:"bytes,3,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// Requested OAuth scopes.
	OauthScopes []string `protobuf:"bytes,4,rep,name=oauth_scopes,json=oauthScopes,proto3" json:"oauth_scopes,omitempty"`
	// Who requested and who can use this token.
	ProxyIdentity string `protobuf:"bytes,5,opt,name=proxy_identity,json=proxyIdentity,proto3" json:"proxy_identity,omitempty"`
	// On whose behalf the token is generated.
	EndUserIdentity string `protobuf:"bytes,6,opt,name=end_user_identity,json=endUserIdentity,proto3" json:"end_user_identity,omitempty"`
	// When this request happened.
	RequestedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=requested_at,json=requestedAt,proto3" json:"requested_at,omitempty"`
	// When the token expires.
	Expiration *timestamp.Timestamp `protobuf:"bytes,8,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// Arbitrary key:value pairs provided by the caller.
	AuditTags []string `protobuf:"bytes,9,rep,name=audit_tags,json=auditTags,proto3" json:"audit_tags,omitempty"`
	// Revision of the luci-config repo with rules.
	ConfigRev string `protobuf:"bytes,10,opt,name=config_rev,json=configRev,proto3" json:"config_rev,omitempty"`
	// Name of the rule used to authorize this call.
	ConfigRule string `protobuf:"bytes,11,opt,name=config_rule,json=configRule,proto3" json:"config_rule,omitempty"`
	// IP address of the caller.
	PeerIp string `protobuf:"bytes,12,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	// Identifier of the token server GAE app and version.
	ServiceVersion string `protobuf:"bytes,13,opt,name=service_version,json=serviceVersion,proto3" json:"service_version,omitempty"`
	// ID of the GAE request that handled the call.
	GaeRequestId string `protobuf:"bytes,14,opt,name=gae_request_id,json=gaeRequestId,proto3" json:"gae_request_id,omitempty"`
	// Revision of groups database used to authorize this call.
	AuthDbRev            int64    `protobuf:"varint,15,opt,name=auth_db_rev,json=authDbRev,proto3" json:"auth_db_rev,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthToken) Reset()         { *m = OAuthToken{} }
func (m *OAuthToken) String() string { return proto.CompactTextString(m) }
func (*OAuthToken) ProtoMessage()    {}
func (*OAuthToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0cc6154e6d463f7, []int{3}
}

func (m *OAuthToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthToken.Unmarshal(m, b)
}
func (m *OAuthToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthToken.Marshal(b, m, deterministic)
}
func (m *OAuthToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthToken.Merge(m, src)
}
func (m *OAuthToken) XXX_Size() int {
	return xxx_messageInfo_OAuthToken.Size(m)
}
func (m *OAuthToken) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthToken.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthToken proto.InternalMessageInfo

func (m *OAuthToken) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *OAuthToken) GetGrantFingerprint() string {
	if m != nil {
		return m.GrantFingerprint
	}
	return ""
}

func (m *OAuthToken) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func (m *OAuthToken) GetOauthScopes() []string {
	if m != nil {
		return m.OauthScopes
	}
	return nil
}

func (m *OAuthToken) GetProxyIdentity() string {
	if m != nil {
		return m.ProxyIdentity
	}
	return ""
}

func (m *OAuthToken) GetEndUserIdentity() string {
	if m != nil {
		return m.EndUserIdentity
	}
	return ""
}

func (m *OAuthToken) GetRequestedAt() *timestamp.Timestamp {
	if m != nil {
		return m.RequestedAt
	}
	return nil
}

func (m *OAuthToken) GetExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *OAuthToken) GetAuditTags() []string {
	if m != nil {
		return m.AuditTags
	}
	return nil
}

func (m *OAuthToken) GetConfigRev() string {
	if m != nil {
		return m.ConfigRev
	}
	return ""
}

func (m *OAuthToken) GetConfigRule() string {
	if m != nil {
		return m.ConfigRule
	}
	return ""
}

func (m *OAuthToken) GetPeerIp() string {
	if m != nil {
		return m.PeerIp
	}
	return ""
}

func (m *OAuthToken) GetServiceVersion() string {
	if m != nil {
		return m.ServiceVersion
	}
	return ""
}

func (m *OAuthToken) GetGaeRequestId() string {
	if m != nil {
		return m.GaeRequestId
	}
	return ""
}

func (m *OAuthToken) GetAuthDbRev() int64 {
	if m != nil {
		return m.AuthDbRev
	}
	return 0
}

func init() {
	proto.RegisterType((*DelegationToken)(nil), "tokenserver.bq.DelegationToken")
	proto.RegisterType((*MachineToken)(nil), "tokenserver.bq.MachineToken")
	proto.RegisterType((*OAuthTokenGrant)(nil), "tokenserver.bq.OAuthTokenGrant")
	proto.RegisterType((*OAuthToken)(nil), "tokenserver.bq.OAuthToken")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/tokenserver/api/bq/bq.proto", fileDescriptor_e0cc6154e6d463f7)
}

var fileDescriptor_e0cc6154e6d463f7 = []byte{
	// 934 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x95, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0x95, 0x6d, 0xfa, 0xe1, 0xe3, 0xc4, 0x69, 0x87, 0x8b, 0x35, 0x95, 0xca, 0x66, 0x2b,
	0x3e, 0xc2, 0x97, 0x0d, 0x45, 0x02, 0x84, 0xe0, 0x22, 0xb0, 0x2c, 0x8a, 0x10, 0x8b, 0xe4, 0x94,
	0x45, 0xe2, 0xc6, 0x1a, 0xdb, 0xa7, 0xce, 0x68, 0xe3, 0xb1, 0x33, 0x1e, 0x47, 0xdb, 0x17, 0xe2,
	0x35, 0x78, 0x11, 0x6e, 0x78, 0x13, 0xe4, 0x63, 0x3b, 0x75, 0x9b, 0xaa, 0x31, 0x70, 0xc1, 0xdd,
	0xe4, 0x7f, 0xfe, 0x33, 0x9e, 0x99, 0xf3, 0x9b, 0x7f, 0xe0, 0x93, 0x38, 0x75, 0xc2, 0x85, 0x4a,
	0x13, 0x51, 0x24, 0x4e, 0xaa, 0x62, 0x77, 0x59, 0x84, 0xc2, 0xd5, 0xe9, 0x2b, 0x94, 0x39, 0xaa,
	0x35, 0x2a, 0x97, 0x67, 0xc2, 0x0d, 0x56, 0x6e, 0xb0, 0x72, 0x32, 0x95, 0xea, 0x94, 0x59, 0xad,
	0xa2, 0x13, 0xac, 0x4e, 0x9f, 0xc4, 0x69, 0x1a, 0x2f, 0xd1, 0xa5, 0x6a, 0x50, 0x5c, 0xb9, 0x5a,
	0x24, 0x98, 0x6b, 0x9e, 0x64, 0xd5, 0x84, 0xd3, 0xef, 0xef, 0xfd, 0x44, 0xb3, 0x7a, 0xa1, 0x17,
	0x6e, 0x84, 0x4b, 0x8c, 0xb9, 0x16, 0xa9, 0x74, 0x13, 0xcc, 0x73, 0x1e, 0x63, 0xde, 0xd2, 0xea,
	0x65, 0xbe, 0xec, 0xb4, 0xd3, 0x84, 0x87, 0x0b, 0x21, 0xd1, 0x27, 0xbd, 0x9e, 0x39, 0xed, 0x36,
	0x53, 0x48, 0x8d, 0xca, 0x5d, 0x7f, 0x5a, 0x55, 0xfc, 0xea, 0x77, 0xb5, 0xc4, 0xf9, 0x5f, 0xfb,
	0x30, 0x7a, 0xb6, 0xd9, 0xd1, 0x65, 0x69, 0x60, 0x63, 0x30, 0xaf, 0x84, 0x8c, 0x51, 0x65, 0x4a,
	0x48, 0x6d, 0xf7, 0xc6, 0xbd, 0x89, 0xe1, 0xb5, 0x25, 0xf6, 0x39, 0x40, 0xb5, 0xd6, 0x2b, 0x21,
	0x23, 0xfb, 0xd1, 0xb8, 0x37, 0xb1, 0x2e, 0x1e, 0x3b, 0xcd, 0x11, 0x9d, 0x79, 0x11, 0x54, 0xdb,
	0xfc, 0x51, 0xc8, 0xc8, 0x33, 0x68, 0x5c, 0x0e, 0xd9, 0x9b, 0x70, 0x54, 0xcd, 0x13, 0x91, 0xbd,
	0x47, 0xcb, 0x1e, 0xd2, 0xef, 0x59, 0xc4, 0x3e, 0x06, 0x56, 0xdf, 0x0c, 0x46, 0xbe, 0x88, 0x50,
	0x6a, 0xa1, 0xaf, 0xed, 0x3e, 0x99, 0x4e, 0x36, 0x95, 0x59, 0x5d, 0x28, 0xed, 0x0a, 0x57, 0x05,
	0xe6, 0x3a, 0x55, 0x37, 0xf6, 0xfd, 0xca, 0xbe, 0xa9, 0x6c, 0xec, 0x5f, 0x80, 0x21, 0xf2, 0xbc,
	0xc0, 0xc8, 0xe7, 0xda, 0x3e, 0x18, 0xf7, 0x26, 0xe6, 0xc5, 0xa9, 0x53, 0xf5, 0xd7, 0x69, 0xfa,
	0xeb, 0x5c, 0x36, 0xfd, 0xf5, 0x8e, 0x2a, 0xf3, 0x54, 0xb3, 0xaf, 0x00, 0xf0, 0x75, 0x26, 0x14,
	0x5d, 0x8f, 0x7d, 0xb8, 0x73, 0x66, 0xcb, 0xcd, 0xde, 0x83, 0x91, 0xe6, 0x2a, 0x46, 0xed, 0xf3,
	0x22, 0x12, 0x28, 0x43, 0xb4, 0x8f, 0xc6, 0x7b, 0x13, 0xc3, 0xb3, 0x2a, 0x79, 0x5a, 0xab, 0x2d,
	0x63, 0xd9, 0x36, 0x11, 0x62, 0x6e, 0x1b, 0x6d, 0xe3, 0xbc, 0x56, 0x5b, 0xa7, 0xc6, 0xc8, 0x5f,
	0xf3, 0xa5, 0x88, 0xca, 0x53, 0xc3, 0xb8, 0x37, 0xd9, 0xdb, 0x9c, 0x1a, 0xa3, 0x97, 0x75, 0x81,
	0xbd, 0x0f, 0xc7, 0x37, 0xf6, 0xb2, 0xeb, 0x52, 0xdb, 0x26, 0x5d, 0xd1, 0x68, 0xa3, 0xcf, 0x48,
	0x66, 0x0c, 0xfa, 0x9a, 0xc7, 0xb9, 0x3d, 0xa0, 0xef, 0xd2, 0x98, 0x9d, 0x01, 0x84, 0xa9, 0xbc,
	0x12, 0xb1, 0xaf, 0x70, 0x6d, 0x0f, 0x69, 0xa2, 0x51, 0x29, 0x1e, 0xae, 0xd9, 0x13, 0x30, 0x9b,
	0x72, 0xb1, 0x44, 0xdb, 0xa2, 0x7a, 0x3d, 0xc3, 0x2b, 0x96, 0xc8, 0x1e, 0xc3, 0x61, 0x86, 0xa8,
	0x7c, 0x91, 0xd9, 0x23, 0x2a, 0x1e, 0x94, 0x3f, 0x67, 0x59, 0x79, 0xde, 0xfa, 0xa0, 0xfe, 0x1a,
	0x55, 0x5e, 0xde, 0xec, 0x31, 0x19, 0xac, 0x5a, 0x7e, 0x59, 0xa9, 0xec, 0x6d, 0xb0, 0x62, 0x8e,
	0x7e, 0xbd, 0xd9, 0x92, 0x9a, 0x13, 0xf2, 0x0d, 0x62, 0x8e, 0x5e, 0x25, 0xce, 0x22, 0xf6, 0x16,
	0x98, 0xe5, 0x6b, 0xf3, 0xa3, 0x80, 0x36, 0xca, 0xe8, 0x3a, 0x8c, 0x52, 0x7a, 0x16, 0x78, 0xb8,
	0x3e, 0xff, 0xa3, 0x0f, 0x83, 0x9f, 0xaa, 0xe7, 0xd3, 0x15, 0xf0, 0xa7, 0x30, 0x68, 0x1e, 0xdc,
	0xd5, 0x2a, 0x92, 0x84, 0xb8, 0xe1, 0x99, 0xb5, 0xf6, 0x7c, 0x15, 0x49, 0xf6, 0x75, 0xf3, 0x06,
	0xf4, 0x75, 0x86, 0x44, 0xb3, 0x75, 0x71, 0xe6, 0xb4, 0x33, 0xa4, 0xfd, 0xcd, 0xcb, 0xeb, 0x0c,
	0xeb, 0x97, 0x50, 0x0e, 0x6f, 0x03, 0xd9, 0xff, 0xd7, 0x40, 0xee, 0xff, 0x23, 0x20, 0x3f, 0x02,
	0x16, 0xa2, 0x22, 0xca, 0x04, 0x5f, 0xfa, 0xb2, 0x48, 0x02, 0x54, 0xf4, 0x1c, 0x0c, 0xef, 0xb8,
	0xac, 0xcc, 0xa9, 0xf0, 0x82, 0x74, 0xf6, 0x2b, 0xbc, 0x91, 0x8b, 0x58, 0x72, 0x5d, 0x28, 0xf4,
	0xf9, 0x32, 0x4e, 0x95, 0xd0, 0x8b, 0x84, 0xde, 0x80, 0x75, 0xf1, 0xee, 0xad, 0x93, 0xd6, 0x91,
	0x32, 0x6f, 0xec, 0xd3, 0xc6, 0xed, 0xb1, 0x7c, 0x4b, 0x2b, 0xbb, 0x1a, 0x72, 0x3f, 0x4c, 0x93,
	0x24, 0x95, 0xbe, 0xe4, 0x49, 0xf9, 0x2c, 0xa8, 0xab, 0x21, 0xff, 0x8e, 0xc4, 0x17, 0x3c, 0x41,
	0x76, 0x0e, 0x43, 0x72, 0x6d, 0x00, 0x34, 0xaa, 0x1e, 0x94, 0xa6, 0x06, 0xc1, 0x16, 0x61, 0xb0,
	0x8b, 0x30, 0xb3, 0x23, 0x61, 0x83, 0x6d, 0xc2, 0xce, 0x7f, 0xef, 0xc3, 0xe8, 0xe7, 0x69, 0xa1,
	0x17, 0xd4, 0xcb, 0x1f, 0x14, 0x97, 0xba, 0x03, 0x44, 0xed, 0xb4, 0x7b, 0x74, 0x3b, 0xed, 0x5a,
	0xfb, 0xe3, 0x61, 0x98, 0x16, 0x52, 0xd7, 0x79, 0xd8, 0xec, 0x6f, 0x5a, 0xa9, 0xec, 0x1d, 0xb0,
	0x32, 0x95, 0xbe, 0xbe, 0xbe, 0x1b, 0x89, 0x43, 0x52, 0x37, 0xf9, 0xf6, 0x01, 0x9c, 0xa0, 0x8c,
	0xfc, 0x22, 0xc7, 0xad, 0x34, 0x1c, 0xa1, 0x8c, 0x7e, 0xc9, 0xf1, 0x7f, 0xce, 0xc2, 0x33, 0x80,
	0x32, 0x04, 0xb5, 0x4f, 0x29, 0x53, 0xc5, 0xa0, 0x41, 0xca, 0xe5, 0x76, 0xd4, 0x18, 0x3b, 0xa2,
	0x06, 0x1e, 0x8a, 0x1a, 0x73, 0x17, 0x08, 0x83, 0x8e, 0x20, 0x0c, 0x77, 0x47, 0x8d, 0x75, 0x37,
	0x6a, 0xfe, 0xec, 0x03, 0xdc, 0x80, 0xd2, 0x81, 0x91, 0x0f, 0xe1, 0x24, 0x2e, 0x71, 0xf2, 0xdb,
	0xbe, 0x0a, 0x96, 0x63, 0x2a, 0x3c, 0x6f, 0x99, 0x3b, 0x53, 0xf3, 0x14, 0x06, 0x29, 0xed, 0x33,
	0x0f, 0xd3, 0x0c, 0x73, 0xbb, 0x4f, 0xf7, 0x6d, 0x92, 0x36, 0x27, 0xe9, 0x1e, 0xb0, 0xf6, 0x3b,
	0x83, 0x75, 0x70, 0x3f, 0x58, 0xdf, 0xc0, 0xe0, 0xe6, 0xef, 0x86, 0xeb, 0x0e, 0x84, 0x98, 0x1b,
	0xff, 0x16, 0x5e, 0x47, 0xff, 0x01, 0x2f, 0xe3, 0x61, 0xbc, 0x60, 0x07, 0x5e, 0xe6, 0x43, 0x78,
	0x0d, 0x76, 0xe1, 0x35, 0xec, 0x88, 0x97, 0xb5, 0x1b, 0xaf, 0xd1, 0x1d, 0xbc, 0xbe, 0xed, 0xff,
	0xf6, 0x28, 0x58, 0x05, 0x07, 0x74, 0x19, 0x9f, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x66, 0x44,
	0xce, 0x5b, 0xe3, 0x0a, 0x00, 0x00,
}
