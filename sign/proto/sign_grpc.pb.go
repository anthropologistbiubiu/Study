// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: proto/sign.proto

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

const (
	SignServiceRequest_GetSign_FullMethodName = "/SignServiceRequest/GetSign"
)

// SignServiceRequestClient is the client API for SignServiceRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignServiceRequestClient interface {
	GetSign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignReponse, error)
}

type signServiceRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewSignServiceRequestClient(cc grpc.ClientConnInterface) SignServiceRequestClient {
	return &signServiceRequestClient{cc}
}

func (c *signServiceRequestClient) GetSign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignReponse, error) {
	out := new(SignReponse)
	err := c.cc.Invoke(ctx, SignServiceRequest_GetSign_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignServiceRequestServer is the server API for SignServiceRequest service.
// All implementations must embed UnimplementedSignServiceRequestServer
// for forward compatibility
type SignServiceRequestServer interface {
	GetSign(context.Context, *SignRequest) (*SignReponse, error)
	//mustEmbedUnimplementedSignServiceRequestServer()
}

// UnimplementedSignServiceRequestServer must be embedded to have forward compatible implementations.
type UnimplementedSignServiceRequestServer struct {
}

func (UnimplementedSignServiceRequestServer) GetSign(context.Context, *SignRequest) (*SignReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSign not implemented")
}
func (UnimplementedSignServiceRequestServer) mustEmbedUnimplementedSignServiceRequestServer() {}

// UnsafeSignServiceRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignServiceRequestServer will
// result in compilation errors.
type UnsafeSignServiceRequestServer interface {
	mustEmbedUnimplementedSignServiceRequestServer()
}

func RegisterSignServiceRequestServer(s grpc.ServiceRegistrar, srv SignServiceRequestServer) {
	s.RegisterService(&SignServiceRequest_ServiceDesc, srv)
}

func _SignServiceRequest_GetSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServiceRequestServer).GetSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SignServiceRequest_GetSign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServiceRequestServer).GetSign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignServiceRequest_ServiceDesc is the grpc.ServiceDesc for SignServiceRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignServiceRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SignServiceRequest",
	HandlerType: (*SignServiceRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSign",
			Handler:    _SignServiceRequest_GetSign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sign.proto",
}