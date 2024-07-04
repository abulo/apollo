// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: system_tenant.proto

// system_tenant 租户

package tenant

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
	SystemTenantService_SystemTenantCreate_FullMethodName    = "/tenant.SystemTenantService/SystemTenantCreate"
	SystemTenantService_SystemTenantUpdate_FullMethodName    = "/tenant.SystemTenantService/SystemTenantUpdate"
	SystemTenantService_SystemTenantDelete_FullMethodName    = "/tenant.SystemTenantService/SystemTenantDelete"
	SystemTenantService_SystemTenant_FullMethodName          = "/tenant.SystemTenantService/SystemTenant"
	SystemTenantService_SystemTenantRecover_FullMethodName   = "/tenant.SystemTenantService/SystemTenantRecover"
	SystemTenantService_SystemTenantList_FullMethodName      = "/tenant.SystemTenantService/SystemTenantList"
	SystemTenantService_SystemTenantListTotal_FullMethodName = "/tenant.SystemTenantService/SystemTenantListTotal"
)

// SystemTenantServiceClient is the client API for SystemTenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SystemTenantService 服务
type SystemTenantServiceClient interface {
	SystemTenantCreate(ctx context.Context, in *SystemTenantCreateRequest, opts ...grpc.CallOption) (*SystemTenantCreateResponse, error)
	SystemTenantUpdate(ctx context.Context, in *SystemTenantUpdateRequest, opts ...grpc.CallOption) (*SystemTenantUpdateResponse, error)
	SystemTenantDelete(ctx context.Context, in *SystemTenantDeleteRequest, opts ...grpc.CallOption) (*SystemTenantDeleteResponse, error)
	SystemTenant(ctx context.Context, in *SystemTenantRequest, opts ...grpc.CallOption) (*SystemTenantResponse, error)
	SystemTenantRecover(ctx context.Context, in *SystemTenantRecoverRequest, opts ...grpc.CallOption) (*SystemTenantRecoverResponse, error)
	SystemTenantList(ctx context.Context, in *SystemTenantListRequest, opts ...grpc.CallOption) (*SystemTenantListResponse, error)
	SystemTenantListTotal(ctx context.Context, in *SystemTenantListTotalRequest, opts ...grpc.CallOption) (*SystemTenantTotalResponse, error)
}

type systemTenantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemTenantServiceClient(cc grpc.ClientConnInterface) SystemTenantServiceClient {
	return &systemTenantServiceClient{cc}
}

func (c *systemTenantServiceClient) SystemTenantCreate(ctx context.Context, in *SystemTenantCreateRequest, opts ...grpc.CallOption) (*SystemTenantCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemTenantCreateResponse)
	err := c.cc.Invoke(ctx, SystemTenantService_SystemTenantCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemTenantServiceClient) SystemTenantUpdate(ctx context.Context, in *SystemTenantUpdateRequest, opts ...grpc.CallOption) (*SystemTenantUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemTenantUpdateResponse)
	err := c.cc.Invoke(ctx, SystemTenantService_SystemTenantUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemTenantServiceClient) SystemTenantDelete(ctx context.Context, in *SystemTenantDeleteRequest, opts ...grpc.CallOption) (*SystemTenantDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemTenantDeleteResponse)
	err := c.cc.Invoke(ctx, SystemTenantService_SystemTenantDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemTenantServiceClient) SystemTenant(ctx context.Context, in *SystemTenantRequest, opts ...grpc.CallOption) (*SystemTenantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemTenantResponse)
	err := c.cc.Invoke(ctx, SystemTenantService_SystemTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemTenantServiceClient) SystemTenantRecover(ctx context.Context, in *SystemTenantRecoverRequest, opts ...grpc.CallOption) (*SystemTenantRecoverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemTenantRecoverResponse)
	err := c.cc.Invoke(ctx, SystemTenantService_SystemTenantRecover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemTenantServiceClient) SystemTenantList(ctx context.Context, in *SystemTenantListRequest, opts ...grpc.CallOption) (*SystemTenantListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemTenantListResponse)
	err := c.cc.Invoke(ctx, SystemTenantService_SystemTenantList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemTenantServiceClient) SystemTenantListTotal(ctx context.Context, in *SystemTenantListTotalRequest, opts ...grpc.CallOption) (*SystemTenantTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemTenantTotalResponse)
	err := c.cc.Invoke(ctx, SystemTenantService_SystemTenantListTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemTenantServiceServer is the server API for SystemTenantService service.
// All implementations must embed UnimplementedSystemTenantServiceServer
// for forward compatibility
//
// SystemTenantService 服务
type SystemTenantServiceServer interface {
	SystemTenantCreate(context.Context, *SystemTenantCreateRequest) (*SystemTenantCreateResponse, error)
	SystemTenantUpdate(context.Context, *SystemTenantUpdateRequest) (*SystemTenantUpdateResponse, error)
	SystemTenantDelete(context.Context, *SystemTenantDeleteRequest) (*SystemTenantDeleteResponse, error)
	SystemTenant(context.Context, *SystemTenantRequest) (*SystemTenantResponse, error)
	SystemTenantRecover(context.Context, *SystemTenantRecoverRequest) (*SystemTenantRecoverResponse, error)
	SystemTenantList(context.Context, *SystemTenantListRequest) (*SystemTenantListResponse, error)
	SystemTenantListTotal(context.Context, *SystemTenantListTotalRequest) (*SystemTenantTotalResponse, error)
	mustEmbedUnimplementedSystemTenantServiceServer()
}

// UnimplementedSystemTenantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemTenantServiceServer struct {
}

func (UnimplementedSystemTenantServiceServer) SystemTenantCreate(context.Context, *SystemTenantCreateRequest) (*SystemTenantCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenantCreate not implemented")
}
func (UnimplementedSystemTenantServiceServer) SystemTenantUpdate(context.Context, *SystemTenantUpdateRequest) (*SystemTenantUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenantUpdate not implemented")
}
func (UnimplementedSystemTenantServiceServer) SystemTenantDelete(context.Context, *SystemTenantDeleteRequest) (*SystemTenantDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenantDelete not implemented")
}
func (UnimplementedSystemTenantServiceServer) SystemTenant(context.Context, *SystemTenantRequest) (*SystemTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenant not implemented")
}
func (UnimplementedSystemTenantServiceServer) SystemTenantRecover(context.Context, *SystemTenantRecoverRequest) (*SystemTenantRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenantRecover not implemented")
}
func (UnimplementedSystemTenantServiceServer) SystemTenantList(context.Context, *SystemTenantListRequest) (*SystemTenantListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenantList not implemented")
}
func (UnimplementedSystemTenantServiceServer) SystemTenantListTotal(context.Context, *SystemTenantListTotalRequest) (*SystemTenantTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenantListTotal not implemented")
}
func (UnimplementedSystemTenantServiceServer) mustEmbedUnimplementedSystemTenantServiceServer() {}

// UnsafeSystemTenantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemTenantServiceServer will
// result in compilation errors.
type UnsafeSystemTenantServiceServer interface {
	mustEmbedUnimplementedSystemTenantServiceServer()
}

func RegisterSystemTenantServiceServer(s grpc.ServiceRegistrar, srv SystemTenantServiceServer) {
	s.RegisterService(&SystemTenantService_ServiceDesc, srv)
}

func _SystemTenantService_SystemTenantCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantServiceServer).SystemTenantCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantService_SystemTenantCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantServiceServer).SystemTenantCreate(ctx, req.(*SystemTenantCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemTenantService_SystemTenantUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantServiceServer).SystemTenantUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantService_SystemTenantUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantServiceServer).SystemTenantUpdate(ctx, req.(*SystemTenantUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemTenantService_SystemTenantDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantServiceServer).SystemTenantDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantService_SystemTenantDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantServiceServer).SystemTenantDelete(ctx, req.(*SystemTenantDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemTenantService_SystemTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantServiceServer).SystemTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantService_SystemTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantServiceServer).SystemTenant(ctx, req.(*SystemTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemTenantService_SystemTenantRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantServiceServer).SystemTenantRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantService_SystemTenantRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantServiceServer).SystemTenantRecover(ctx, req.(*SystemTenantRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemTenantService_SystemTenantList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantServiceServer).SystemTenantList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantService_SystemTenantList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantServiceServer).SystemTenantList(ctx, req.(*SystemTenantListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemTenantService_SystemTenantListTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantListTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantServiceServer).SystemTenantListTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantService_SystemTenantListTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantServiceServer).SystemTenantListTotal(ctx, req.(*SystemTenantListTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemTenantService_ServiceDesc is the grpc.ServiceDesc for SystemTenantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemTenantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tenant.SystemTenantService",
	HandlerType: (*SystemTenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemTenantCreate",
			Handler:    _SystemTenantService_SystemTenantCreate_Handler,
		},
		{
			MethodName: "SystemTenantUpdate",
			Handler:    _SystemTenantService_SystemTenantUpdate_Handler,
		},
		{
			MethodName: "SystemTenantDelete",
			Handler:    _SystemTenantService_SystemTenantDelete_Handler,
		},
		{
			MethodName: "SystemTenant",
			Handler:    _SystemTenantService_SystemTenant_Handler,
		},
		{
			MethodName: "SystemTenantRecover",
			Handler:    _SystemTenantService_SystemTenantRecover_Handler,
		},
		{
			MethodName: "SystemTenantList",
			Handler:    _SystemTenantService_SystemTenantList_Handler,
		},
		{
			MethodName: "SystemTenantListTotal",
			Handler:    _SystemTenantService_SystemTenantListTotal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_tenant.proto",
}
