// Code generated by protoc-gen-go.
// source: atomic-counter.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	atomic-counter.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package helloworld

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

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	FromClient(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// Called when executing FromClient
	// Test for nested rpc call
	IncreaseCounter(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHelloAgain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) FromClient(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/FromClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) IncreaseCounter(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/IncreaseCounter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error)
	FromClient(context.Context, *HelloRequest) (*HelloReply, error)
	// Called when executing FromClient
	// Test for nested rpc call
	IncreaseCounter(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHelloAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHelloAgain(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_FromClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).FromClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/FromClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).FromClient(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_IncreaseCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).IncreaseCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/IncreaseCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).IncreaseCounter(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _Greeter_SayHelloAgain_Handler,
		},
		{
			MethodName: "FromClient",
			Handler:    _Greeter_FromClient_Handler,
		},
		{
			MethodName: "IncreaseCounter",
			Handler:    _Greeter_IncreaseCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("atomic-counter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x3f, 0x4b, 0x04, 0x31,
	0x10, 0xc5, 0x5d, 0x11, 0x4f, 0x07, 0xe5, 0x20, 0x88, 0x2c, 0xda, 0x48, 0x0a, 0xb1, 0x31, 0x88,
	0xf6, 0xa2, 0xb7, 0xf8, 0xaf, 0x3b, 0xce, 0xc2, 0x3a, 0xc6, 0x61, 0x0d, 0x24, 0x99, 0x38, 0xc9,
	0xa1, 0xfb, 0x49, 0xfc, 0xba, 0xb2, 0xd1, 0xc5, 0x2b, 0xae, 0xda, 0x6e, 0x1e, 0xf9, 0xbd, 0xf7,
	0x20, 0x0f, 0x0e, 0x74, 0x26, 0x6f, 0xcd, 0xb9, 0xa1, 0x65, 0xc8, 0xc8, 0x2a, 0x32, 0x65, 0x12,
	0xf0, 0x8e, 0xce, 0xd1, 0x27, 0xb1, 0x7b, 0x93, 0x12, 0xf6, 0x1e, 0x7b, 0xb5, 0xc0, 0x8f, 0x25,
	0xa6, 0x2c, 0x04, 0x6c, 0x05, 0xed, 0xb1, 0xae, 0x4e, 0xaa, 0xb3, 0xdd, 0x45, 0xb9, 0xe5, 0x29,
	0xc0, 0x1f, 0x13, 0x5d, 0x27, 0x6a, 0x98, 0x78, 0x4c, 0x49, 0xb7, 0x03, 0x34, 0xc8, 0xcb, 0xef,
	0x4d, 0x98, 0x3c, 0x30, 0x62, 0x46, 0x16, 0xd7, 0xb0, 0xf3, 0xac, 0xbb, 0x62, 0x13, 0xb5, 0xfa,
	0x2f, 0x54, 0xab, 0x6d, 0x47, 0x87, 0x6b, 0x5e, 0xa2, 0xeb, 0xe4, 0x86, 0x68, 0x60, 0x7f, 0xf0,
	0xdf, 0xb6, 0xda, 0x86, 0x51, 0x21, 0x37, 0x00, 0xf7, 0x4c, 0xbe, 0x71, 0x16, 0x43, 0x1e, 0x95,
	0x70, 0x07, 0xd3, 0xa7, 0x60, 0x18, 0x75, 0xc2, 0xe6, 0xf7, 0x0f, 0xc7, 0xc4, 0xcc, 0x2e, 0xe0,
	0xd8, 0x92, 0x6a, 0x39, 0x1a, 0x85, 0x5f, 0xda, 0x47, 0x87, 0x69, 0x85, 0x9d, 0x4d, 0x0b, 0xfc,
	0xd2, 0xdf, 0xf3, 0x7e, 0xa1, 0x79, 0xf5, 0xba, 0x5d, 0xa6, 0xba, 0xfa, 0x09, 0x00, 0x00, 0xff,
	0xff, 0xde, 0x07, 0x0f, 0x71, 0xc2, 0x01, 0x00, 0x00,
}