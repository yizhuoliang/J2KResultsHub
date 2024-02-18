// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: J2KSerializer.proto

package __

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

// BrokerServiceClient is the client API for BrokerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerServiceClient interface {
	// NOTE: claim cell finished implies having no more FS access
	ClaimCellFinished(ctx context.Context, in *VarResults, opts ...grpc.CallOption) (*Empty, error)
	FetchVarResult(ctx context.Context, in *FetchVarResultRequest, opts ...grpc.CallOption) (*VarResult, error)
	// The Ping-Pong service for cluster testing, remove in future
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type brokerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerServiceClient(cc grpc.ClientConnInterface) BrokerServiceClient {
	return &brokerServiceClient{cc}
}

func (c *brokerServiceClient) ClaimCellFinished(ctx context.Context, in *VarResults, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/J2KSerializer.BrokerService/ClaimCellFinished", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerServiceClient) FetchVarResult(ctx context.Context, in *FetchVarResultRequest, opts ...grpc.CallOption) (*VarResult, error) {
	out := new(VarResult)
	err := c.cc.Invoke(ctx, "/J2KSerializer.BrokerService/FetchVarResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/J2KSerializer.BrokerService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServiceServer is the server API for BrokerService service.
// All implementations must embed UnimplementedBrokerServiceServer
// for forward compatibility
type BrokerServiceServer interface {
	// NOTE: claim cell finished implies having no more FS access
	ClaimCellFinished(context.Context, *VarResults) (*Empty, error)
	FetchVarResult(context.Context, *FetchVarResultRequest) (*VarResult, error)
	// The Ping-Pong service for cluster testing, remove in future
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedBrokerServiceServer()
}

// UnimplementedBrokerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBrokerServiceServer struct {
}

func (UnimplementedBrokerServiceServer) ClaimCellFinished(context.Context, *VarResults) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimCellFinished not implemented")
}
func (UnimplementedBrokerServiceServer) FetchVarResult(context.Context, *FetchVarResultRequest) (*VarResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchVarResult not implemented")
}
func (UnimplementedBrokerServiceServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedBrokerServiceServer) mustEmbedUnimplementedBrokerServiceServer() {}

// UnsafeBrokerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerServiceServer will
// result in compilation errors.
type UnsafeBrokerServiceServer interface {
	mustEmbedUnimplementedBrokerServiceServer()
}

func RegisterBrokerServiceServer(s grpc.ServiceRegistrar, srv BrokerServiceServer) {
	s.RegisterService(&BrokerService_ServiceDesc, srv)
}

func _BrokerService_ClaimCellFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VarResults)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).ClaimCellFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/J2KSerializer.BrokerService/ClaimCellFinished",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).ClaimCellFinished(ctx, req.(*VarResults))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerService_FetchVarResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchVarResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).FetchVarResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/J2KSerializer.BrokerService/FetchVarResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).FetchVarResult(ctx, req.(*FetchVarResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/J2KSerializer.BrokerService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BrokerService_ServiceDesc is the grpc.ServiceDesc for BrokerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrokerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "J2KSerializer.BrokerService",
	HandlerType: (*BrokerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClaimCellFinished",
			Handler:    _BrokerService_ClaimCellFinished_Handler,
		},
		{
			MethodName: "FetchVarResult",
			Handler:    _BrokerService_FetchVarResult_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _BrokerService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "J2KSerializer.proto",
}
