// Code generated by protoc-gen-go.
// source: grpclb.proto
// DO NOT EDIT!

/*
Package grpc_lb_v1 is a generated protocol buffer package.

It is generated from these files:
	grpclb.proto

It has these top-level messages:
	Duration
	LoadBalanceRequest
	InitialLoadBalanceRequest
	ClientStats
	LoadBalanceResponse
	InitialLoadBalanceResponse
	ServerList
	Server
*/
package grpc_lb_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Duration struct {
	// Signed seconds of the span of time. Must be from -315,576,000,000
	// to +315,576,000,000 inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Signed fractions of a second at nanosecond resolution of the span
	// of time. Durations less than one second are represented with a 0
	// `seconds` field and a positive or negative `nanos` field. For durations
	// of one second or more, a non-zero value for the `nanos` field must be
	// of the same sign as the `seconds` field. Must be from -999,999,999
	// to +999,999,999 inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *Duration) Reset()                    { *m = Duration{} }
func (m *Duration) String() string            { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()               {}
func (*Duration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LoadBalanceRequest struct {
	// Types that are valid to be assigned to LoadBalanceRequestType:
	//	*LoadBalanceRequest_InitialRequest
	//	*LoadBalanceRequest_ClientStats
	LoadBalanceRequestType isLoadBalanceRequest_LoadBalanceRequestType `protobuf_oneof:"load_balance_request_type"`
}

func (m *LoadBalanceRequest) Reset()                    { *m = LoadBalanceRequest{} }
func (m *LoadBalanceRequest) String() string            { return proto.CompactTextString(m) }
func (*LoadBalanceRequest) ProtoMessage()               {}
func (*LoadBalanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isLoadBalanceRequest_LoadBalanceRequestType interface {
	isLoadBalanceRequest_LoadBalanceRequestType()
}

type LoadBalanceRequest_InitialRequest struct {
	InitialRequest *InitialLoadBalanceRequest `protobuf:"bytes,1,opt,name=initial_request,oneof"`
}
type LoadBalanceRequest_ClientStats struct {
	ClientStats *ClientStats `protobuf:"bytes,2,opt,name=client_stats,oneof"`
}

func (*LoadBalanceRequest_InitialRequest) isLoadBalanceRequest_LoadBalanceRequestType() {}
func (*LoadBalanceRequest_ClientStats) isLoadBalanceRequest_LoadBalanceRequestType()    {}

func (m *LoadBalanceRequest) GetLoadBalanceRequestType() isLoadBalanceRequest_LoadBalanceRequestType {
	if m != nil {
		return m.LoadBalanceRequestType
	}
	return nil
}

func (m *LoadBalanceRequest) GetInitialRequest() *InitialLoadBalanceRequest {
	if x, ok := m.GetLoadBalanceRequestType().(*LoadBalanceRequest_InitialRequest); ok {
		return x.InitialRequest
	}
	return nil
}

func (m *LoadBalanceRequest) GetClientStats() *ClientStats {
	if x, ok := m.GetLoadBalanceRequestType().(*LoadBalanceRequest_ClientStats); ok {
		return x.ClientStats
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LoadBalanceRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LoadBalanceRequest_OneofMarshaler, _LoadBalanceRequest_OneofUnmarshaler, _LoadBalanceRequest_OneofSizer, []interface{}{
		(*LoadBalanceRequest_InitialRequest)(nil),
		(*LoadBalanceRequest_ClientStats)(nil),
	}
}

func _LoadBalanceRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LoadBalanceRequest)
	// load_balance_request_type
	switch x := m.LoadBalanceRequestType.(type) {
	case *LoadBalanceRequest_InitialRequest:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InitialRequest); err != nil {
			return err
		}
	case *LoadBalanceRequest_ClientStats:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClientStats); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LoadBalanceRequest.LoadBalanceRequestType has unexpected type %T", x)
	}
	return nil
}

func _LoadBalanceRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LoadBalanceRequest)
	switch tag {
	case 1: // load_balance_request_type.initial_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InitialLoadBalanceRequest)
		err := b.DecodeMessage(msg)
		m.LoadBalanceRequestType = &LoadBalanceRequest_InitialRequest{msg}
		return true, err
	case 2: // load_balance_request_type.client_stats
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientStats)
		err := b.DecodeMessage(msg)
		m.LoadBalanceRequestType = &LoadBalanceRequest_ClientStats{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LoadBalanceRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LoadBalanceRequest)
	// load_balance_request_type
	switch x := m.LoadBalanceRequestType.(type) {
	case *LoadBalanceRequest_InitialRequest:
		s := proto.Size(x.InitialRequest)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LoadBalanceRequest_ClientStats:
		s := proto.Size(x.ClientStats)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type InitialLoadBalanceRequest struct {
	// Name of load balanced service (IE, service.grpc.gslb.google.com)
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *InitialLoadBalanceRequest) Reset()                    { *m = InitialLoadBalanceRequest{} }
func (m *InitialLoadBalanceRequest) String() string            { return proto.CompactTextString(m) }
func (*InitialLoadBalanceRequest) ProtoMessage()               {}
func (*InitialLoadBalanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Contains client level statistics that are useful to load balancing. Each
// count should be reset to zero after reporting the stats.
type ClientStats struct {
	// The total number of requests sent by the client since the last report.
	TotalRequests int64 `protobuf:"varint,1,opt,name=total_requests" json:"total_requests,omitempty"`
	// The number of client rpc errors since the last report.
	ClientRpcErrors int64 `protobuf:"varint,2,opt,name=client_rpc_errors" json:"client_rpc_errors,omitempty"`
	// The number of dropped requests since the last report.
	DroppedRequests int64 `protobuf:"varint,3,opt,name=dropped_requests" json:"dropped_requests,omitempty"`
}

func (m *ClientStats) Reset()                    { *m = ClientStats{} }
func (m *ClientStats) String() string            { return proto.CompactTextString(m) }
func (*ClientStats) ProtoMessage()               {}
func (*ClientStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type LoadBalanceResponse struct {
	// Types that are valid to be assigned to LoadBalanceResponseType:
	//	*LoadBalanceResponse_InitialResponse
	//	*LoadBalanceResponse_ServerList
	LoadBalanceResponseType isLoadBalanceResponse_LoadBalanceResponseType `protobuf_oneof:"load_balance_response_type"`
}

func (m *LoadBalanceResponse) Reset()                    { *m = LoadBalanceResponse{} }
func (m *LoadBalanceResponse) String() string            { return proto.CompactTextString(m) }
func (*LoadBalanceResponse) ProtoMessage()               {}
func (*LoadBalanceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isLoadBalanceResponse_LoadBalanceResponseType interface {
	isLoadBalanceResponse_LoadBalanceResponseType()
}

type LoadBalanceResponse_InitialResponse struct {
	InitialResponse *InitialLoadBalanceResponse `protobuf:"bytes,1,opt,name=initial_response,oneof"`
}
type LoadBalanceResponse_ServerList struct {
	ServerList *ServerList `protobuf:"bytes,2,opt,name=server_list,oneof"`
}

func (*LoadBalanceResponse_InitialResponse) isLoadBalanceResponse_LoadBalanceResponseType() {}
func (*LoadBalanceResponse_ServerList) isLoadBalanceResponse_LoadBalanceResponseType()      {}

func (m *LoadBalanceResponse) GetLoadBalanceResponseType() isLoadBalanceResponse_LoadBalanceResponseType {
	if m != nil {
		return m.LoadBalanceResponseType
	}
	return nil
}

func (m *LoadBalanceResponse) GetInitialResponse() *InitialLoadBalanceResponse {
	if x, ok := m.GetLoadBalanceResponseType().(*LoadBalanceResponse_InitialResponse); ok {
		return x.InitialResponse
	}
	return nil
}

func (m *LoadBalanceResponse) GetServerList() *ServerList {
	if x, ok := m.GetLoadBalanceResponseType().(*LoadBalanceResponse_ServerList); ok {
		return x.ServerList
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LoadBalanceResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LoadBalanceResponse_OneofMarshaler, _LoadBalanceResponse_OneofUnmarshaler, _LoadBalanceResponse_OneofSizer, []interface{}{
		(*LoadBalanceResponse_InitialResponse)(nil),
		(*LoadBalanceResponse_ServerList)(nil),
	}
}

func _LoadBalanceResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LoadBalanceResponse)
	// load_balance_response_type
	switch x := m.LoadBalanceResponseType.(type) {
	case *LoadBalanceResponse_InitialResponse:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InitialResponse); err != nil {
			return err
		}
	case *LoadBalanceResponse_ServerList:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ServerList); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LoadBalanceResponse.LoadBalanceResponseType has unexpected type %T", x)
	}
	return nil
}

func _LoadBalanceResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LoadBalanceResponse)
	switch tag {
	case 1: // load_balance_response_type.initial_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InitialLoadBalanceResponse)
		err := b.DecodeMessage(msg)
		m.LoadBalanceResponseType = &LoadBalanceResponse_InitialResponse{msg}
		return true, err
	case 2: // load_balance_response_type.server_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ServerList)
		err := b.DecodeMessage(msg)
		m.LoadBalanceResponseType = &LoadBalanceResponse_ServerList{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LoadBalanceResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LoadBalanceResponse)
	// load_balance_response_type
	switch x := m.LoadBalanceResponseType.(type) {
	case *LoadBalanceResponse_InitialResponse:
		s := proto.Size(x.InitialResponse)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LoadBalanceResponse_ServerList:
		s := proto.Size(x.ServerList)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type InitialLoadBalanceResponse struct {
	// This is an application layer redirect that indicates the client should use
	// the specified server for load balancing. When this field is non-empty in
	// the response, the client should open a separate connection to the
	// load_balancer_delegate and call the BalanceLoad method.
	LoadBalancerDelegate string `protobuf:"bytes,1,opt,name=load_balancer_delegate" json:"load_balancer_delegate,omitempty"`
	// This interval defines how often the client should send the client stats
	// to the load balancer. Stats should only be reported when the duration is
	// positive.
	ClientStatsReportInterval *Duration `protobuf:"bytes,3,opt,name=client_stats_report_interval" json:"client_stats_report_interval,omitempty"`
}

func (m *InitialLoadBalanceResponse) Reset()                    { *m = InitialLoadBalanceResponse{} }
func (m *InitialLoadBalanceResponse) String() string            { return proto.CompactTextString(m) }
func (*InitialLoadBalanceResponse) ProtoMessage()               {}
func (*InitialLoadBalanceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *InitialLoadBalanceResponse) GetClientStatsReportInterval() *Duration {
	if m != nil {
		return m.ClientStatsReportInterval
	}
	return nil
}

type ServerList struct {
	// Contains a list of servers selected by the load balancer. The list will
	// be updated when server resolutions change or as needed to balance load
	// across more servers. The client should consume the server list in order
	// unless instructed otherwise via the client_config.
	Servers []*Server `protobuf:"bytes,1,rep,name=servers" json:"servers,omitempty"`
	// Indicates the amount of time that the client should consider this server
	// list as valid. It may be considered stale after waiting this interval of
	// time after receiving the list. If the interval is not positive, the
	// client can assume the list is valid until the next list is received.
	ExpirationInterval *Duration `protobuf:"bytes,3,opt,name=expiration_interval" json:"expiration_interval,omitempty"`
}

func (m *ServerList) Reset()                    { *m = ServerList{} }
func (m *ServerList) String() string            { return proto.CompactTextString(m) }
func (*ServerList) ProtoMessage()               {}
func (*ServerList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ServerList) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

func (m *ServerList) GetExpirationInterval() *Duration {
	if m != nil {
		return m.ExpirationInterval
	}
	return nil
}

type Server struct {
	// A resolved address for the server, serialized in network-byte-order. It may
	// either be an IPv4 or IPv6 address.
	IpAddress []byte `protobuf:"bytes,1,opt,name=ip_address,proto3" json:"ip_address,omitempty"`
	// A resolved port number for the server.
	Port int32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	// An opaque but printable token given to the frontend for each pick. All
	// frontend requests for that pick must include the token in its initial
	// metadata. The token is used by the backend to verify the request and to
	// allow the backend to report load to the gRPC LB system.
	LoadBalanceToken string `protobuf:"bytes,3,opt,name=load_balance_token" json:"load_balance_token,omitempty"`
	// Indicates whether this particular request should be dropped by the client
	// when this server is chosen from the list.
	DropRequest bool `protobuf:"varint,4,opt,name=drop_request" json:"drop_request,omitempty"`
}

func (m *Server) Reset()                    { *m = Server{} }
func (m *Server) String() string            { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()               {}
func (*Server) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Duration)(nil), "grpc.lb.v1.Duration")
	proto.RegisterType((*LoadBalanceRequest)(nil), "grpc.lb.v1.LoadBalanceRequest")
	proto.RegisterType((*InitialLoadBalanceRequest)(nil), "grpc.lb.v1.InitialLoadBalanceRequest")
	proto.RegisterType((*ClientStats)(nil), "grpc.lb.v1.ClientStats")
	proto.RegisterType((*LoadBalanceResponse)(nil), "grpc.lb.v1.LoadBalanceResponse")
	proto.RegisterType((*InitialLoadBalanceResponse)(nil), "grpc.lb.v1.InitialLoadBalanceResponse")
	proto.RegisterType((*ServerList)(nil), "grpc.lb.v1.ServerList")
	proto.RegisterType((*Server)(nil), "grpc.lb.v1.Server")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LoadBalancer service

type LoadBalancerClient interface {
	// Bidirectional rpc to get a list of servers.
	BalanceLoad(ctx context.Context, opts ...grpc.CallOption) (LoadBalancer_BalanceLoadClient, error)
}

type loadBalancerClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancerClient(cc *grpc.ClientConn) LoadBalancerClient {
	return &loadBalancerClient{cc}
}

func (c *loadBalancerClient) BalanceLoad(ctx context.Context, opts ...grpc.CallOption) (LoadBalancer_BalanceLoadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_LoadBalancer_serviceDesc.Streams[0], c.cc, "/grpc.lb.v1.LoadBalancer/BalanceLoad", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadBalancerBalanceLoadClient{stream}
	return x, nil
}

type LoadBalancer_BalanceLoadClient interface {
	Send(*LoadBalanceRequest) error
	Recv() (*LoadBalanceResponse, error)
	grpc.ClientStream
}

type loadBalancerBalanceLoadClient struct {
	grpc.ClientStream
}

func (x *loadBalancerBalanceLoadClient) Send(m *LoadBalanceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadBalancerBalanceLoadClient) Recv() (*LoadBalanceResponse, error) {
	m := new(LoadBalanceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for LoadBalancer service

type LoadBalancerServer interface {
	// Bidirectional rpc to get a list of servers.
	BalanceLoad(LoadBalancer_BalanceLoadServer) error
}

func RegisterLoadBalancerServer(s *grpc.Server, srv LoadBalancerServer) {
	s.RegisterService(&_LoadBalancer_serviceDesc, srv)
}

func _LoadBalancer_BalanceLoad_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadBalancerServer).BalanceLoad(&loadBalancerBalanceLoadServer{stream})
}

type LoadBalancer_BalanceLoadServer interface {
	Send(*LoadBalanceResponse) error
	Recv() (*LoadBalanceRequest, error)
	grpc.ServerStream
}

type loadBalancerBalanceLoadServer struct {
	grpc.ServerStream
}

func (x *loadBalancerBalanceLoadServer) Send(m *LoadBalanceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadBalancerBalanceLoadServer) Recv() (*LoadBalanceRequest, error) {
	m := new(LoadBalanceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadBalancer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.lb.v1.LoadBalancer",
	HandlerType: (*LoadBalancerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BalanceLoad",
			Handler:       _LoadBalancer_BalanceLoad_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpclb.proto",
}

func init() { proto.RegisterFile("grpclb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0x51, 0x6f, 0xd3, 0x3e,
	0x14, 0xc5, 0x9b, 0x7f, 0xb7, 0xfd, 0xb7, 0x9b, 0xc0, 0xc6, 0xdd, 0x54, 0xda, 0x32, 0x8d, 0x2a,
	0x08, 0x54, 0x90, 0x08, 0x2c, 0xbc, 0xf1, 0x84, 0x0a, 0x0f, 0x45, 0xda, 0xd3, 0xf6, 0x86, 0x90,
	0x2c, 0x27, 0xb9, 0x9a, 0x2c, 0x82, 0x6d, 0x6c, 0xaf, 0x1a, 0xdf, 0x07, 0xf1, 0x39, 0x91, 0xe3,
	0x94, 0x64, 0x54, 0x15, 0xbc, 0xc5, 0xbe, 0x3e, 0xf7, 0x1e, 0xff, 0x7c, 0x02, 0xc9, 0xb5, 0xd1,
	0x65, 0x5d, 0x64, 0xda, 0x28, 0xa7, 0x10, 0xfc, 0x2a, 0xab, 0x8b, 0x6c, 0x75, 0x9e, 0xbe, 0x80,
	0xfd, 0x0f, 0x37, 0x86, 0x3b, 0xa1, 0x24, 0x1e, 0xc2, 0xff, 0x96, 0x4a, 0x25, 0x2b, 0x3b, 0x8e,
	0x66, 0xd1, 0x7c, 0x88, 0xf7, 0x60, 0x57, 0x72, 0xa9, 0xec, 0xf8, 0xbf, 0x59, 0x34, 0xdf, 0x4d,
	0x7f, 0x44, 0x80, 0x17, 0x8a, 0x57, 0x0b, 0x5e, 0x73, 0x59, 0xd2, 0x25, 0x7d, 0xbb, 0x21, 0xeb,
	0xf0, 0x1d, 0x1c, 0x0a, 0x29, 0x9c, 0xe0, 0x35, 0x33, 0x61, 0xab, 0x91, 0xc7, 0xf9, 0xd3, 0xac,
	0x1b, 0x94, 0x7d, 0x0c, 0x47, 0x36, 0xf5, 0xcb, 0x01, 0xbe, 0x82, 0xa4, 0xac, 0x05, 0x49, 0xc7,
	0xac, 0xe3, 0x2e, 0x8c, 0x8b, 0xf3, 0x87, 0x7d, 0xf9, 0xfb, 0xa6, 0x7e, 0xe5, 0xcb, 0xcb, 0xc1,
	0xe2, 0x11, 0x4c, 0x6a, 0xc5, 0x2b, 0x56, 0x84, 0x4e, 0xeb, 0xb9, 0xcc, 0x7d, 0xd7, 0x94, 0x3e,
	0x87, 0xc9, 0xd6, 0x61, 0x98, 0xc0, 0x8e, 0xe4, 0x5f, 0xa9, 0x71, 0x78, 0x90, 0x7e, 0x82, 0xb8,
	0xd7, 0x18, 0x47, 0x70, 0xdf, 0x29, 0xd7, 0xdd, 0x63, 0xcd, 0x61, 0x02, 0x0f, 0x5a, 0x7f, 0x46,
	0x97, 0x8c, 0x8c, 0x51, 0x26, 0x98, 0x1c, 0xe2, 0x18, 0x8e, 0x2a, 0xa3, 0xb4, 0xa6, 0xaa, 0x13,
	0x0d, 0x7d, 0x25, 0xfd, 0x19, 0xc1, 0xf1, 0x1d, 0x03, 0x56, 0x2b, 0x69, 0x09, 0x17, 0x70, 0xd4,
	0xe1, 0x0a, 0x7b, 0x2d, 0xaf, 0x67, 0x7f, 0xe3, 0x15, 0x4e, 0x2f, 0x07, 0xf8, 0x12, 0x62, 0x4b,
	0x66, 0x45, 0x86, 0xd5, 0xc2, 0xba, 0x96, 0xd7, 0xa8, 0x2f, 0xbf, 0x6a, 0xca, 0x17, 0xc2, 0xf3,
	0x5d, 0x9c, 0xc2, 0xf4, 0x0f, 0x5c, 0xa1, 0x53, 0xe0, 0x75, 0x0b, 0xd3, 0xed, 0xc3, 0xf0, 0x0c,
	0x46, 0x7d, 0xad, 0x61, 0x15, 0xd5, 0x74, 0xcd, 0x5d, 0x8b, 0x10, 0xdf, 0xc2, 0x69, 0xff, 0xed,
	0x98, 0x21, 0xad, 0x8c, 0x63, 0x42, 0x3a, 0x32, 0x2b, 0x5e, 0x37, 0x30, 0xe2, 0xfc, 0xa4, 0xef,
	0x6d, 0x1d, 0xb8, 0xb4, 0x02, 0xe8, 0x7c, 0xe2, 0x13, 0x1f, 0x3f, 0xbf, 0xf2, 0xd8, 0x87, 0xf3,
	0x38, 0xc7, 0xcd, 0x0b, 0xe1, 0x39, 0x1c, 0xd3, 0xad, 0x16, 0xa1, 0xc1, 0xbf, 0x4d, 0xf9, 0x0c,
	0x7b, 0xad, 0x18, 0x01, 0x84, 0x66, 0xbc, 0xaa, 0x0c, 0xd9, 0xf0, 0xb6, 0x89, 0x0f, 0x84, 0x37,
	0x1c, 0x22, 0x8e, 0x53, 0xc0, 0x3b, 0xa4, 0x9c, 0xfa, 0x42, 0xb2, 0xe9, 0x7e, 0x80, 0x27, 0x90,
	0xf8, 0xa7, 0xfe, 0x1d, 0xf2, 0x9d, 0x59, 0x34, 0xdf, 0xcf, 0x0b, 0x48, 0x7a, 0xd8, 0x0c, 0x5e,
	0x42, 0xdc, 0x7e, 0xfb, 0x6d, 0x3c, 0xeb, 0x5b, 0xda, 0xcc, 0xe3, 0xf4, 0xf1, 0xd6, 0x7a, 0xe0,
	0x3f, 0x8f, 0x5e, 0x47, 0xc5, 0x5e, 0xf3, 0xdf, 0xbe, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0x01,
	0x8b, 0xc9, 0x26, 0xc7, 0x03, 0x00, 0x00,
}
