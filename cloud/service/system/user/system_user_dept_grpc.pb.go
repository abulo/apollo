// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: system_user_dept.proto

// system_user_dept 系统用户部门

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SystemUserDeptService_SystemUserDeptCreate_FullMethodName = "/user.SystemUserDeptService/SystemUserDeptCreate"
	SystemUserDeptService_SystemUserDeptList_FullMethodName   = "/user.SystemUserDeptService/SystemUserDeptList"
)

// SystemUserDeptServiceClient is the client API for SystemUserDeptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SystemUserDeptService 服务
type SystemUserDeptServiceClient interface {
	SystemUserDeptCreate(ctx context.Context, in *SystemUserDeptCreateRequest, opts ...grpc.CallOption) (*SystemUserDeptCreateResponse, error)
	SystemUserDeptList(ctx context.Context, in *SystemUserDeptListRequest, opts ...grpc.CallOption) (*SystemUserDeptListResponse, error)
}

type systemUserDeptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemUserDeptServiceClient(cc grpc.ClientConnInterface) SystemUserDeptServiceClient {
	return &systemUserDeptServiceClient{cc}
}

func (c *systemUserDeptServiceClient) SystemUserDeptCreate(ctx context.Context, in *SystemUserDeptCreateRequest, opts ...grpc.CallOption) (*SystemUserDeptCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserDeptCreateResponse)
	err := c.cc.Invoke(ctx, SystemUserDeptService_SystemUserDeptCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserDeptServiceClient) SystemUserDeptList(ctx context.Context, in *SystemUserDeptListRequest, opts ...grpc.CallOption) (*SystemUserDeptListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserDeptListResponse)
	err := c.cc.Invoke(ctx, SystemUserDeptService_SystemUserDeptList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemUserDeptServiceServer is the server API for SystemUserDeptService service.
// All implementations must embed UnimplementedSystemUserDeptServiceServer
// for forward compatibility
//
// SystemUserDeptService 服务
type SystemUserDeptServiceServer interface {
	SystemUserDeptCreate(context.Context, *SystemUserDeptCreateRequest) (*SystemUserDeptCreateResponse, error)
	SystemUserDeptList(context.Context, *SystemUserDeptListRequest) (*SystemUserDeptListResponse, error)
	mustEmbedUnimplementedSystemUserDeptServiceServer()
}

// UnimplementedSystemUserDeptServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemUserDeptServiceServer struct {
}

func (UnimplementedSystemUserDeptServiceServer) SystemUserDeptCreate(context.Context, *SystemUserDeptCreateRequest) (*SystemUserDeptCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserDeptCreate not implemented")
}
func (UnimplementedSystemUserDeptServiceServer) SystemUserDeptList(context.Context, *SystemUserDeptListRequest) (*SystemUserDeptListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserDeptList not implemented")
}
func (UnimplementedSystemUserDeptServiceServer) mustEmbedUnimplementedSystemUserDeptServiceServer() {}

// UnsafeSystemUserDeptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemUserDeptServiceServer will
// result in compilation errors.
type UnsafeSystemUserDeptServiceServer interface {
	mustEmbedUnimplementedSystemUserDeptServiceServer()
}

func RegisterSystemUserDeptServiceServer(s grpc.ServiceRegistrar, srv SystemUserDeptServiceServer) {
	s.RegisterService(&SystemUserDeptService_ServiceDesc, srv)
}

func _SystemUserDeptService_SystemUserDeptCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserDeptCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserDeptServiceServer).SystemUserDeptCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserDeptService_SystemUserDeptCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserDeptServiceServer).SystemUserDeptCreate(ctx, req.(*SystemUserDeptCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserDeptService_SystemUserDeptList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserDeptListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserDeptServiceServer).SystemUserDeptList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserDeptService_SystemUserDeptList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserDeptServiceServer).SystemUserDeptList(ctx, req.(*SystemUserDeptListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemUserDeptService_ServiceDesc is the grpc.ServiceDesc for SystemUserDeptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemUserDeptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.SystemUserDeptService",
	HandlerType: (*SystemUserDeptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemUserDeptCreate",
			Handler:    _SystemUserDeptService_SystemUserDeptCreate_Handler,
		},
		{
			MethodName: "SystemUserDeptList",
			Handler:    _SystemUserDeptService_SystemUserDeptList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_user_dept.proto",
}
