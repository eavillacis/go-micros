// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: solana.proto

package solana

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

// SolanaServiceClient is the client API for SolanaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SolanaServiceClient interface {
	GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
}

type solanaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSolanaServiceClient(cc grpc.ClientConnInterface) SolanaServiceClient {
	return &solanaServiceClient{cc}
}

func (c *solanaServiceClient) GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/solana.SolanaService/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolanaServiceServer is the server API for SolanaService service.
// All implementations must embed UnimplementedSolanaServiceServer
// for forward compatibility
type SolanaServiceServer interface {
	GetBalance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	mustEmbedUnimplementedSolanaServiceServer()
}

// UnimplementedSolanaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSolanaServiceServer struct {
}

func (UnimplementedSolanaServiceServer) GetBalance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedSolanaServiceServer) mustEmbedUnimplementedSolanaServiceServer() {}

// UnsafeSolanaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SolanaServiceServer will
// result in compilation errors.
type UnsafeSolanaServiceServer interface {
	mustEmbedUnimplementedSolanaServiceServer()
}

func RegisterSolanaServiceServer(s grpc.ServiceRegistrar, srv SolanaServiceServer) {
	s.RegisterService(&SolanaService_ServiceDesc, srv)
}

func _SolanaService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolanaServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solana.SolanaService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolanaServiceServer).GetBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SolanaService_ServiceDesc is the grpc.ServiceDesc for SolanaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SolanaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "solana.SolanaService",
	HandlerType: (*SolanaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _SolanaService_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "solana.proto",
}
