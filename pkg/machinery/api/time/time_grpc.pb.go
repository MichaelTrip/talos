// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: time/time.proto

package time

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TimeService_Time_FullMethodName      = "/time.TimeService/Time"
	TimeService_TimeCheck_FullMethodName = "/time.TimeService/TimeCheck"
)

// TimeServiceClient is the client API for TimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The time service definition.
type TimeServiceClient interface {
	Time(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TimeResponse, error)
	TimeCheck(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*TimeResponse, error)
}

type timeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeServiceClient(cc grpc.ClientConnInterface) TimeServiceClient {
	return &timeServiceClient{cc}
}

func (c *timeServiceClient) Time(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TimeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TimeResponse)
	err := c.cc.Invoke(ctx, TimeService_Time_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeServiceClient) TimeCheck(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*TimeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TimeResponse)
	err := c.cc.Invoke(ctx, TimeService_TimeCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeServiceServer is the server API for TimeService service.
// All implementations must embed UnimplementedTimeServiceServer
// for forward compatibility.
//
// The time service definition.
type TimeServiceServer interface {
	Time(context.Context, *emptypb.Empty) (*TimeResponse, error)
	TimeCheck(context.Context, *TimeRequest) (*TimeResponse, error)
	mustEmbedUnimplementedTimeServiceServer()
}

// UnimplementedTimeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTimeServiceServer struct{}

func (UnimplementedTimeServiceServer) Time(context.Context, *emptypb.Empty) (*TimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Time not implemented")
}
func (UnimplementedTimeServiceServer) TimeCheck(context.Context, *TimeRequest) (*TimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimeCheck not implemented")
}
func (UnimplementedTimeServiceServer) mustEmbedUnimplementedTimeServiceServer() {}
func (UnimplementedTimeServiceServer) testEmbeddedByValue()                     {}

// UnsafeTimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeServiceServer will
// result in compilation errors.
type UnsafeTimeServiceServer interface {
	mustEmbedUnimplementedTimeServiceServer()
}

func RegisterTimeServiceServer(s grpc.ServiceRegistrar, srv TimeServiceServer) {
	// If the following call panics, it indicates UnimplementedTimeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TimeService_ServiceDesc, srv)
}

func _TimeService_Time_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).Time(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeService_Time_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).Time(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeService_TimeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).TimeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeService_TimeCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).TimeCheck(ctx, req.(*TimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeService_ServiceDesc is the grpc.ServiceDesc for TimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "time.TimeService",
	HandlerType: (*TimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Time",
			Handler:    _TimeService_Time_Handler,
		},
		{
			MethodName: "TimeCheck",
			Handler:    _TimeService_TimeCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "time/time.proto",
}
