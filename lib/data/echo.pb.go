// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

package data

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EchoMessage struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoMessage) Reset()         { *m = EchoMessage{} }
func (m *EchoMessage) String() string { return proto.CompactTextString(m) }
func (*EchoMessage) ProtoMessage()    {}
func (*EchoMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{0}
}

func (m *EchoMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoMessage.Unmarshal(m, b)
}
func (m *EchoMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoMessage.Marshal(b, m, deterministic)
}
func (m *EchoMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoMessage.Merge(m, src)
}
func (m *EchoMessage) XXX_Size() int {
	return xxx_messageInfo_EchoMessage.Size(m)
}
func (m *EchoMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EchoMessage proto.InternalMessageInfo

func (m *EchoMessage) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoMessage)(nil), "data.EchoMessage")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor_08134aea513e0001) }

var fileDescriptor_08134aea513e0001 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x54, 0xd2, 0xe4, 0xe2,
	0x76, 0x4d, 0xce, 0xc8, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x92, 0xe2, 0xe2, 0x48,
	0x2f, 0x4a, 0x4d, 0x2d, 0xc9, 0xcc, 0x4b, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3,
	0x8d, 0xac, 0xb8, 0x58, 0x40, 0x4a, 0x85, 0x8c, 0xb8, 0x38, 0x82, 0x13, 0x2b, 0x3d, 0x52, 0x73,
	0x72, 0xf2, 0x85, 0x04, 0xf5, 0x40, 0xa6, 0xe8, 0x21, 0x19, 0x21, 0x85, 0x29, 0xa4, 0xc4, 0x90,
	0xc4, 0x06, 0xb6, 0xd3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x0a, 0x91, 0x4a, 0x81, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	SayHello(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) SayHello(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error) {
	out := new(EchoMessage)
	err := c.cc.Invoke(ctx, "/data.Echo/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	SayHello(context.Context, *EchoMessage) (*EchoMessage, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.Echo/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).SayHello(ctx, req.(*EchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Echo_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}