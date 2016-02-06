// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package services is a generated protocol buffer package.

It is generated from these files:
	service.proto
	state.proto

It has these top-level messages:
	GetConfigResponse
	RegisterStreamRequest
	TerminateStreamRequest
	LogStreamState
*/
package services

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import logpb "github.com/luci/luci-go/common/proto/logdog/logpb"
import google_protobuf2 "github.com/luci/luci-go/common/proto/google"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// GetConfigResponse is the response structure for the user
// "GetConfig" endpoint.
type GetConfigResponse struct {
	// The API URL of the base "luci-config" service. If empty, the default
	// service URL will be used.
	ConfigServiceUrl string `protobuf:"bytes,1,opt,name=config_service_url" json:"config_service_url,omitempty"`
	// The name of the configuration set to load from.
	ConfigSet string `protobuf:"bytes,2,opt,name=config_set" json:"config_set,omitempty"`
	// The path of the text-serialized configuration protobuf.
	ConfigPath string `protobuf:"bytes,3,opt,name=config_path" json:"config_path,omitempty"`
}

func (m *GetConfigResponse) Reset()                    { *m = GetConfigResponse{} }
func (m *GetConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*GetConfigResponse) ProtoMessage()               {}
func (*GetConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// RegisterStreamRequest is the set of caller-supplied data for the
// RegisterStream Coordinator service endpoint.
type RegisterStreamRequest struct {
	// The log stream's path.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// The log stream's secret.
	Secret []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	// The protobuf version string for this stream.
	ProtoVersion string `protobuf:"bytes,3,opt,name=proto_version" json:"proto_version,omitempty"`
	// The serialized LogStreamDescriptor protobuf for this stream.
	Desc *logpb.LogStreamDescriptor `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
}

func (m *RegisterStreamRequest) Reset()                    { *m = RegisterStreamRequest{} }
func (m *RegisterStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterStreamRequest) ProtoMessage()               {}
func (*RegisterStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterStreamRequest) GetDesc() *logpb.LogStreamDescriptor {
	if m != nil {
		return m.Desc
	}
	return nil
}

// TerminateStreamRequest is the set of caller-supplied data for the
// TerminateStream Coordinator service endpoint.
type TerminateStreamRequest struct {
	// The log stream's path.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// The log stream's secret.
	Secret []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	// The terminal index of the stream.
	TerminalIndex int64 `protobuf:"varint,3,opt,name=terminal_index" json:"terminal_index,omitempty"`
}

func (m *TerminateStreamRequest) Reset()                    { *m = TerminateStreamRequest{} }
func (m *TerminateStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*TerminateStreamRequest) ProtoMessage()               {}
func (*TerminateStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*GetConfigResponse)(nil), "services.GetConfigResponse")
	proto.RegisterType((*RegisterStreamRequest)(nil), "services.RegisterStreamRequest")
	proto.RegisterType((*TerminateStreamRequest)(nil), "services.TerminateStreamRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Services service

type ServicesClient interface { // GetConfig allows a service to retrieve the current service configuration
	// parameters.
	GetConfig(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*GetConfigResponse, error)
	// RegisterStream is an idempotent stream state register operation.
	RegisterStream(ctx context.Context, in *RegisterStreamRequest, opts ...grpc.CallOption) (*LogStreamState, error)
	// TerminateStream is an idempotent operation to update the stream's terminal
	// index.
	TerminateStream(ctx context.Context, in *TerminateStreamRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}
type servicesPRPCClient struct {
	client *prpccommon.Client
}

func NewServicesPRPCClient(client *prpccommon.Client) ServicesClient {
	return &servicesPRPCClient{client}
}

func (c *servicesPRPCClient) GetConfig(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.client.Call(ctx, "services.Services", "GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesPRPCClient) RegisterStream(ctx context.Context, in *RegisterStreamRequest, opts ...grpc.CallOption) (*LogStreamState, error) {
	out := new(LogStreamState)
	err := c.client.Call(ctx, "services.Services", "RegisterStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesPRPCClient) TerminateStream(ctx context.Context, in *TerminateStreamRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := c.client.Call(ctx, "services.Services", "TerminateStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type servicesClient struct {
	cc *grpc.ClientConn
}

func NewServicesClient(cc *grpc.ClientConn) ServicesClient {
	return &servicesClient{cc}
}

func (c *servicesClient) GetConfig(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := grpc.Invoke(ctx, "/services.Services/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) RegisterStream(ctx context.Context, in *RegisterStreamRequest, opts ...grpc.CallOption) (*LogStreamState, error) {
	out := new(LogStreamState)
	err := grpc.Invoke(ctx, "/services.Services/RegisterStream", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) TerminateStream(ctx context.Context, in *TerminateStreamRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/services.Services/TerminateStream", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Services service

type ServicesServer interface {
	// GetConfig allows a service to retrieve the current service configuration
	// parameters.
	GetConfig(context.Context, *google_protobuf2.Empty) (*GetConfigResponse, error)
	// RegisterStream is an idempotent stream state register operation.
	RegisterStream(context.Context, *RegisterStreamRequest) (*LogStreamState, error)
	// TerminateStream is an idempotent operation to update the stream's terminal
	// index.
	TerminateStream(context.Context, *TerminateStreamRequest) (*google_protobuf2.Empty, error)
}

func RegisterServicesServer(s prpc.Registrar, srv ServicesServer) {
	s.RegisterService(&_Services_serviceDesc, srv)
}

func _Services_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf2.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ServicesServer).GetConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Services_RegisterStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RegisterStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ServicesServer).RegisterStream(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Services_TerminateStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(TerminateStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ServicesServer).TerminateStream(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Services_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.Services",
	HandlerType: (*ServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _Services_GetConfig_Handler,
		},
		{
			MethodName: "RegisterStream",
			Handler:    _Services_RegisterStream_Handler,
		},
		{
			MethodName: "TerminateStream",
			Handler:    _Services_TerminateStream_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x83, 0x10, 0x02, 0xc3, 0x1f, 0xe3, 0x1a, 0x48, 0x53, 0x0e, 0x12, 0x4e, 0x5c, 0xdc,
	0x26, 0x78, 0xf4, 0x62, 0xa2, 0xc6, 0x83, 0xc6, 0x03, 0x78, 0xf0, 0xd6, 0xd0, 0x32, 0x2c, 0x9b,
	0xb4, 0x9d, 0xba, 0xbb, 0x25, 0xfa, 0x85, 0xfd, 0x1c, 0xb6, 0xdb, 0x02, 0x51, 0xf1, 0xe0, 0x65,
	0x9b, 0x7d, 0xf3, 0xf6, 0xed, 0xcc, 0x6f, 0x0b, 0x3d, 0x8d, 0x6a, 0x2b, 0x43, 0xe4, 0xa9, 0x22,
	0x43, 0xac, 0x55, 0x6d, 0xb5, 0xdb, 0xd1, 0x66, 0x69, 0x2a, 0xd9, 0xbd, 0x16, 0xd2, 0x6c, 0xb2,
	0x80, 0x87, 0x14, 0x7b, 0x51, 0x16, 0x4a, 0xbb, 0x5c, 0x0a, 0xf2, 0x72, 0x21, 0xa6, 0xc4, 0xb3,
	0x2e, 0x2f, 0x22, 0xb1, 0x22, 0x51, 0x7c, 0xd2, 0xa0, 0x58, 0xab, 0xc3, 0x23, 0x41, 0x24, 0x22,
	0x2c, 0x4d, 0x41, 0xb6, 0xf6, 0x30, 0x4e, 0xcd, 0x47, 0x59, 0x9c, 0xbc, 0xc2, 0xd9, 0x03, 0x9a,
	0x5b, 0x4a, 0xd6, 0x52, 0xcc, 0x51, 0xa7, 0x94, 0x68, 0x64, 0x2e, 0xb0, 0xd0, 0x2a, 0x7e, 0xd5,
	0x8e, 0x9f, 0xa9, 0xc8, 0xa9, 0x8d, 0x6b, 0xd3, 0x36, 0x63, 0x00, 0xfb, 0x9a, 0x71, 0x4e, 0xac,
	0x76, 0x0e, 0x9d, 0x4a, 0x4b, 0x97, 0x66, 0xe3, 0xd4, 0x0b, 0x71, 0xb2, 0x85, 0xc1, 0x1c, 0x85,
	0xd4, 0x06, 0xd5, 0xc2, 0x28, 0x5c, 0xc6, 0x73, 0x7c, 0xcb, 0x50, 0x1b, 0xd6, 0x85, 0x86, 0xb5,
	0x95, 0x79, 0x7d, 0x68, 0x6a, 0x0c, 0x55, 0x95, 0xd5, 0x65, 0x03, 0xe8, 0xd9, 0xce, 0xfc, 0x2d,
	0x2a, 0x2d, 0x29, 0x29, 0xd3, 0xd8, 0x14, 0x1a, 0x2b, 0xd4, 0xa1, 0xd3, 0xc8, 0x77, 0x9d, 0x99,
	0xcb, 0xed, 0x90, 0xfc, 0x89, 0x44, 0x99, 0x7d, 0x97, 0xd7, 0x94, 0x4c, 0x0d, 0xa9, 0xc9, 0x33,
	0x0c, 0x5f, 0x50, 0xc5, 0x32, 0xc9, 0xf1, 0xfd, 0xe7, 0xe2, 0x21, 0xf4, 0x4d, 0x79, 0x2e, 0xf2,
	0x65, 0xb2, 0xc2, 0x77, 0x7b, 0x73, 0x7d, 0xf6, 0x59, 0x83, 0xd6, 0xa2, 0x7a, 0x15, 0x76, 0x03,
	0xed, 0x3d, 0x2e, 0x36, 0xe4, 0x25, 0x59, 0xbe, 0x23, 0xcb, 0xef, 0x0b, 0xb2, 0xee, 0x88, 0xef,
	0x5e, 0x91, 0xff, 0x66, 0xfb, 0x08, 0xfd, 0xef, 0x58, 0xd8, 0xc5, 0xc1, 0x7e, 0x14, 0x98, 0xeb,
	0x1c, 0x0c, 0xfb, 0x81, 0x17, 0xc5, 0xdf, 0x91, 0x87, 0x9d, 0xfe, 0x98, 0x95, 0x8d, 0x0f, 0xe6,
	0xe3, 0x18, 0xdc, 0x3f, 0xda, 0x0e, 0x9a, 0x76, 0x7f, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0xee,
	0xdf, 0xb7, 0xdc, 0x93, 0x02, 0x00, 0x00,
}
