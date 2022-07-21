// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: clientserver.proto

package proto

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

// ClientServerServiceClient is the client API for ClientServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServerServiceClient interface {
	// Unary
	CreateBlockchain(ctx context.Context, in *CreationRequest, opts ...grpc.CallOption) (*CreationResponse, error)
	// Client Streaming
	AddBlock(ctx context.Context, opts ...grpc.CallOption) (ClientServerService_AddBlockClient, error)
	// Server Streaming - Return response to all client if the blockchain is corrupted
	IsValid(ctx context.Context, in *IsValidRequest, opts ...grpc.CallOption) (ClientServerService_IsValidClient, error)
}

type clientServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServerServiceClient(cc grpc.ClientConnInterface) ClientServerServiceClient {
	return &clientServerServiceClient{cc}
}

func (c *clientServerServiceClient) CreateBlockchain(ctx context.Context, in *CreationRequest, opts ...grpc.CallOption) (*CreationResponse, error) {
	out := new(CreationResponse)
	err := c.cc.Invoke(ctx, "/clientserver.ClientServerService/CreateBlockchain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServerServiceClient) AddBlock(ctx context.Context, opts ...grpc.CallOption) (ClientServerService_AddBlockClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientServerService_ServiceDesc.Streams[0], "/clientserver.ClientServerService/AddBlock", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServerServiceAddBlockClient{stream}
	return x, nil
}

type ClientServerService_AddBlockClient interface {
	Send(*AddBlockRequest) error
	CloseAndRecv() (*AddBlockResponse, error)
	grpc.ClientStream
}

type clientServerServiceAddBlockClient struct {
	grpc.ClientStream
}

func (x *clientServerServiceAddBlockClient) Send(m *AddBlockRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientServerServiceAddBlockClient) CloseAndRecv() (*AddBlockResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddBlockResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServerServiceClient) IsValid(ctx context.Context, in *IsValidRequest, opts ...grpc.CallOption) (ClientServerService_IsValidClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientServerService_ServiceDesc.Streams[1], "/clientserver.ClientServerService/IsValid", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServerServiceIsValidClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientServerService_IsValidClient interface {
	Recv() (*IsValidResponse, error)
	grpc.ClientStream
}

type clientServerServiceIsValidClient struct {
	grpc.ClientStream
}

func (x *clientServerServiceIsValidClient) Recv() (*IsValidResponse, error) {
	m := new(IsValidResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientServerServiceServer is the server API for ClientServerService service.
// All implementations must embed UnimplementedClientServerServiceServer
// for forward compatibility
type ClientServerServiceServer interface {
	// Unary
	CreateBlockchain(context.Context, *CreationRequest) (*CreationResponse, error)
	// Client Streaming
	AddBlock(ClientServerService_AddBlockServer) error
	// Server Streaming - Return response to all client if the blockchain is corrupted
	IsValid(*IsValidRequest, ClientServerService_IsValidServer) error
	mustEmbedUnimplementedClientServerServiceServer()
}

// UnimplementedClientServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServerServiceServer struct {
}

func (UnimplementedClientServerServiceServer) CreateBlockchain(context.Context, *CreationRequest) (*CreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlockchain not implemented")
}
func (UnimplementedClientServerServiceServer) AddBlock(ClientServerService_AddBlockServer) error {
	return status.Errorf(codes.Unimplemented, "method AddBlock not implemented")
}
func (UnimplementedClientServerServiceServer) IsValid(*IsValidRequest, ClientServerService_IsValidServer) error {
	return status.Errorf(codes.Unimplemented, "method IsValid not implemented")
}
func (UnimplementedClientServerServiceServer) mustEmbedUnimplementedClientServerServiceServer() {}

// UnsafeClientServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServerServiceServer will
// result in compilation errors.
type UnsafeClientServerServiceServer interface {
	mustEmbedUnimplementedClientServerServiceServer()
}

func RegisterClientServerServiceServer(s grpc.ServiceRegistrar, srv ClientServerServiceServer) {
	s.RegisterService(&ClientServerService_ServiceDesc, srv)
}

func _ClientServerService_CreateBlockchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServerServiceServer).CreateBlockchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientserver.ClientServerService/CreateBlockchain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServerServiceServer).CreateBlockchain(ctx, req.(*CreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientServerService_AddBlock_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientServerServiceServer).AddBlock(&clientServerServiceAddBlockServer{stream})
}

type ClientServerService_AddBlockServer interface {
	SendAndClose(*AddBlockResponse) error
	Recv() (*AddBlockRequest, error)
	grpc.ServerStream
}

type clientServerServiceAddBlockServer struct {
	grpc.ServerStream
}

func (x *clientServerServiceAddBlockServer) SendAndClose(m *AddBlockResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientServerServiceAddBlockServer) Recv() (*AddBlockRequest, error) {
	m := new(AddBlockRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClientServerService_IsValid_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IsValidRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServerServiceServer).IsValid(m, &clientServerServiceIsValidServer{stream})
}

type ClientServerService_IsValidServer interface {
	Send(*IsValidResponse) error
	grpc.ServerStream
}

type clientServerServiceIsValidServer struct {
	grpc.ServerStream
}

func (x *clientServerServiceIsValidServer) Send(m *IsValidResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ClientServerService_ServiceDesc is the grpc.ServiceDesc for ClientServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clientserver.ClientServerService",
	HandlerType: (*ClientServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlockchain",
			Handler:    _ClientServerService_CreateBlockchain_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddBlock",
			Handler:       _ClientServerService_AddBlock_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "IsValid",
			Handler:       _ClientServerService_IsValid_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "clientserver.proto",
}
