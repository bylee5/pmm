// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agentlocal/agent_local.proto

package agentlocal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_local_f7cd48981737bb93, []int{0}
}
func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (dst *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(dst, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusResponse struct {
	NodeId               uint32   `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_local_f7cd48981737bb93, []int{1}
}
func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (dst *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(dst, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "agentlocal.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "agentlocal.StatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentLocalClient is the client API for AgentLocal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentLocalClient interface {
	// Status returns current pmm-agent status.
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type agentLocalClient struct {
	cc *grpc.ClientConn
}

func NewAgentLocalClient(cc *grpc.ClientConn) AgentLocalClient {
	return &agentLocalClient{cc}
}

func (c *agentLocalClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/agentlocal.AgentLocal/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentLocalServer is the server API for AgentLocal service.
type AgentLocalServer interface {
	// Status returns current pmm-agent status.
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
}

func RegisterAgentLocalServer(s *grpc.Server, srv AgentLocalServer) {
	s.RegisterService(&_AgentLocal_serviceDesc, srv)
}

func _AgentLocal_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentLocalServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentlocal.AgentLocal/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentLocalServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentLocal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agentlocal.AgentLocal",
	HandlerType: (*AgentLocalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _AgentLocal_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agentlocal/agent_local.proto",
}

func init() {
	proto.RegisterFile("agentlocal/agent_local.proto", fileDescriptor_agent_local_f7cd48981737bb93)
}

var fileDescriptor_agent_local_f7cd48981737bb93 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x4c, 0x4f, 0xcd,
	0x2b, 0xc9, 0xc9, 0x4f, 0x4e, 0xcc, 0xd1, 0x07, 0x33, 0xe3, 0xc1, 0x6c, 0xbd, 0x82, 0xa2, 0xfc,
	0x92, 0x7c, 0x21, 0x2e, 0x84, 0xac, 0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62,
	0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0xa5,
	0x12, 0x3f, 0x17, 0x6f, 0x70, 0x49, 0x62, 0x49, 0x69, 0x71, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71,
	0x89, 0x92, 0x26, 0x17, 0x1f, 0x4c, 0xa0, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x9c, 0x8b,
	0x3d, 0x2f, 0x3f, 0x25, 0x35, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x37, 0x88, 0x0d,
	0xc4, 0xf5, 0x4c, 0x31, 0xca, 0xe2, 0xe2, 0x72, 0x04, 0xd9, 0xe3, 0x03, 0xb2, 0x47, 0x28, 0x86,
	0x8b, 0x0d, 0xa2, 0x51, 0x48, 0x52, 0x0f, 0x61, 0xbd, 0x1e, 0x8a, 0xe9, 0x52, 0x52, 0xd8, 0xa4,
	0x20, 0xf6, 0x28, 0x49, 0x37, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x54, 0x49, 0x40, 0xbf, 0xcc, 0x40,
	0x1f, 0xe2, 0x33, 0x88, 0x0a, 0x2b, 0x46, 0xad, 0x24, 0x36, 0xb0, 0x73, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x10, 0x7b, 0x0f, 0x61, 0xf8, 0x00, 0x00, 0x00,
}
