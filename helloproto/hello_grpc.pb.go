// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: hello.proto

package helloproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	HelloService_SayHello_FullMethodName    = "/helloproto.HelloService/SayHello"
	HelloService_Read_FullMethodName        = "/helloproto.HelloService/Read"
	HelloService_Write_FullMethodName       = "/helloproto.HelloService/Write"
	HelloService_ReadStream_FullMethodName  = "/helloproto.HelloService/ReadStream"
	HelloService_WriteStream_FullMethodName = "/helloproto.HelloService/WriteStream"
)

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadReply, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteReply, error)
	ReadStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_ReadStreamClient, error)
	WriteStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_WriteStreamClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, HelloService_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadReply, error) {
	out := new(ReadReply)
	err := c.cc.Invoke(ctx, HelloService_Read_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteReply, error) {
	out := new(WriteReply)
	err := c.cc.Invoke(ctx, HelloService_Write_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) ReadStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_ReadStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], HelloService_ReadStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceReadStreamClient{stream}
	return x, nil
}

type HelloService_ReadStreamClient interface {
	Send(*ReadRequest) error
	Recv() (*ReadReply, error)
	grpc.ClientStream
}

type helloServiceReadStreamClient struct {
	grpc.ClientStream
}

func (x *helloServiceReadStreamClient) Send(m *ReadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceReadStreamClient) Recv() (*ReadReply, error) {
	m := new(ReadReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) WriteStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_WriteStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[1], HelloService_WriteStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceWriteStreamClient{stream}
	return x, nil
}

type HelloService_WriteStreamClient interface {
	Send(*WriteRequest) error
	Recv() (*WriteReply, error)
	grpc.ClientStream
}

type helloServiceWriteStreamClient struct {
	grpc.ClientStream
}

func (x *helloServiceWriteStreamClient) Send(m *WriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceWriteStreamClient) Recv() (*WriteReply, error) {
	m := new(WriteReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations should embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Read(context.Context, *ReadRequest) (*ReadReply, error)
	Write(context.Context, *WriteRequest) (*WriteReply, error)
	ReadStream(HelloService_ReadStreamServer) error
	WriteStream(HelloService_WriteStreamServer) error
}

// UnimplementedHelloServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServiceServer) Read(context.Context, *ReadRequest) (*ReadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedHelloServiceServer) Write(context.Context, *WriteRequest) (*WriteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedHelloServiceServer) ReadStream(HelloService_ReadStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadStream not implemented")
}
func (UnimplementedHelloServiceServer) WriteStream(HelloService_WriteStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteStream not implemented")
}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloService_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloService_Write_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_ReadStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).ReadStream(&helloServiceReadStreamServer{stream})
}

type HelloService_ReadStreamServer interface {
	Send(*ReadReply) error
	Recv() (*ReadRequest, error)
	grpc.ServerStream
}

type helloServiceReadStreamServer struct {
	grpc.ServerStream
}

func (x *helloServiceReadStreamServer) Send(m *ReadReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceReadStreamServer) Recv() (*ReadRequest, error) {
	m := new(ReadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloService_WriteStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).WriteStream(&helloServiceWriteStreamServer{stream})
}

type HelloService_WriteStreamServer interface {
	Send(*WriteReply) error
	Recv() (*WriteRequest, error)
	grpc.ServerStream
}

type helloServiceWriteStreamServer struct {
	grpc.ServerStream
}

func (x *helloServiceWriteStreamServer) Send(m *WriteReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceWriteStreamServer) Recv() (*WriteRequest, error) {
	m := new(WriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloproto.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _HelloService_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _HelloService_Write_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadStream",
			Handler:       _HelloService_ReadStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "WriteStream",
			Handler:       _HelloService_WriteStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}
