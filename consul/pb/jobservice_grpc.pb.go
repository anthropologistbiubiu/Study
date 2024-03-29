// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: pb/jobservice.proto

package pb

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
	JobServicevRequest_GetJobService_FullMethodName = "/JobServicevRequest/GetJobService"
)

// JobServicevRequestClient is the client API for JobServicevRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobServicevRequestClient interface {
	GetJobService(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type jobServicevRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServicevRequestClient(cc grpc.ClientConnInterface) JobServicevRequestClient {
	return &jobServicevRequestClient{cc}
}

func (c *jobServicevRequestClient) GetJobService(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, JobServicevRequest_GetJobService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServicevRequestServer is the server API for JobServicevRequest service.
// All implementations must embed UnimplementedJobServicevRequestServer
// for forward compatibility
type JobServicevRequestServer interface {
	GetJobService(context.Context, *Request) (*Response, error)
	//mustEmbedUnimplementedJobServicevRequestServer()
}

// UnimplementedJobServicevRequestServer must be embedded to have forward compatible implementations.
type UnimplementedJobServicevRequestServer struct {
}

func (UnimplementedJobServicevRequestServer) GetJobService(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobService not implemented")
}
func (UnimplementedJobServicevRequestServer) mustEmbedUnimplementedJobServicevRequestServer() {}

// UnsafeJobServicevRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServicevRequestServer will
// result in compilation errors.
type UnsafeJobServicevRequestServer interface {
	mustEmbedUnimplementedJobServicevRequestServer()
}

func RegisterJobServicevRequestServer(s grpc.ServiceRegistrar, srv JobServicevRequestServer) {
	s.RegisterService(&JobServicevRequest_ServiceDesc, srv)
}

func _JobServicevRequest_GetJobService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServicevRequestServer).GetJobService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobServicevRequest_GetJobService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServicevRequestServer).GetJobService(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// JobServicevRequest_ServiceDesc is the grpc.ServiceDesc for JobServicevRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobServicevRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "JobServicevRequest",
	HandlerType: (*JobServicevRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJobService",
			Handler:    _JobServicevRequest_GetJobService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/jobservice.proto",
}
