// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: system_menu.proto

// system_menu 系统菜单

package menu

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
	SystemMenuService_SystemMenuCreate_FullMethodName    = "/menu.SystemMenuService/SystemMenuCreate"
	SystemMenuService_SystemMenuUpdate_FullMethodName    = "/menu.SystemMenuService/SystemMenuUpdate"
	SystemMenuService_SystemMenuDelete_FullMethodName    = "/menu.SystemMenuService/SystemMenuDelete"
	SystemMenuService_SystemMenu_FullMethodName          = "/menu.SystemMenuService/SystemMenu"
	SystemMenuService_SystemMenuRecover_FullMethodName   = "/menu.SystemMenuService/SystemMenuRecover"
	SystemMenuService_SystemMenuDrop_FullMethodName      = "/menu.SystemMenuService/SystemMenuDrop"
	SystemMenuService_SystemMenuList_FullMethodName      = "/menu.SystemMenuService/SystemMenuList"
	SystemMenuService_SystemMenuListTotal_FullMethodName = "/menu.SystemMenuService/SystemMenuListTotal"
)

// SystemMenuServiceClient is the client API for SystemMenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SystemMenuService 服务
type SystemMenuServiceClient interface {
	SystemMenuCreate(ctx context.Context, in *SystemMenuCreateRequest, opts ...grpc.CallOption) (*SystemMenuCreateResponse, error)
	SystemMenuUpdate(ctx context.Context, in *SystemMenuUpdateRequest, opts ...grpc.CallOption) (*SystemMenuUpdateResponse, error)
	SystemMenuDelete(ctx context.Context, in *SystemMenuDeleteRequest, opts ...grpc.CallOption) (*SystemMenuDeleteResponse, error)
	SystemMenu(ctx context.Context, in *SystemMenuRequest, opts ...grpc.CallOption) (*SystemMenuResponse, error)
	SystemMenuRecover(ctx context.Context, in *SystemMenuRecoverRequest, opts ...grpc.CallOption) (*SystemMenuRecoverResponse, error)
	SystemMenuDrop(ctx context.Context, in *SystemMenuDropRequest, opts ...grpc.CallOption) (*SystemMenuDropResponse, error)
	SystemMenuList(ctx context.Context, in *SystemMenuListRequest, opts ...grpc.CallOption) (*SystemMenuListResponse, error)
	SystemMenuListTotal(ctx context.Context, in *SystemMenuListTotalRequest, opts ...grpc.CallOption) (*SystemMenuTotalResponse, error)
}

type systemMenuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemMenuServiceClient(cc grpc.ClientConnInterface) SystemMenuServiceClient {
	return &systemMenuServiceClient{cc}
}

func (c *systemMenuServiceClient) SystemMenuCreate(ctx context.Context, in *SystemMenuCreateRequest, opts ...grpc.CallOption) (*SystemMenuCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMenuCreateResponse)
	err := c.cc.Invoke(ctx, SystemMenuService_SystemMenuCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMenuServiceClient) SystemMenuUpdate(ctx context.Context, in *SystemMenuUpdateRequest, opts ...grpc.CallOption) (*SystemMenuUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMenuUpdateResponse)
	err := c.cc.Invoke(ctx, SystemMenuService_SystemMenuUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMenuServiceClient) SystemMenuDelete(ctx context.Context, in *SystemMenuDeleteRequest, opts ...grpc.CallOption) (*SystemMenuDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMenuDeleteResponse)
	err := c.cc.Invoke(ctx, SystemMenuService_SystemMenuDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMenuServiceClient) SystemMenu(ctx context.Context, in *SystemMenuRequest, opts ...grpc.CallOption) (*SystemMenuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMenuResponse)
	err := c.cc.Invoke(ctx, SystemMenuService_SystemMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMenuServiceClient) SystemMenuRecover(ctx context.Context, in *SystemMenuRecoverRequest, opts ...grpc.CallOption) (*SystemMenuRecoverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMenuRecoverResponse)
	err := c.cc.Invoke(ctx, SystemMenuService_SystemMenuRecover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMenuServiceClient) SystemMenuDrop(ctx context.Context, in *SystemMenuDropRequest, opts ...grpc.CallOption) (*SystemMenuDropResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMenuDropResponse)
	err := c.cc.Invoke(ctx, SystemMenuService_SystemMenuDrop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMenuServiceClient) SystemMenuList(ctx context.Context, in *SystemMenuListRequest, opts ...grpc.CallOption) (*SystemMenuListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMenuListResponse)
	err := c.cc.Invoke(ctx, SystemMenuService_SystemMenuList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMenuServiceClient) SystemMenuListTotal(ctx context.Context, in *SystemMenuListTotalRequest, opts ...grpc.CallOption) (*SystemMenuTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMenuTotalResponse)
	err := c.cc.Invoke(ctx, SystemMenuService_SystemMenuListTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemMenuServiceServer is the server API for SystemMenuService service.
// All implementations must embed UnimplementedSystemMenuServiceServer
// for forward compatibility
//
// SystemMenuService 服务
type SystemMenuServiceServer interface {
	SystemMenuCreate(context.Context, *SystemMenuCreateRequest) (*SystemMenuCreateResponse, error)
	SystemMenuUpdate(context.Context, *SystemMenuUpdateRequest) (*SystemMenuUpdateResponse, error)
	SystemMenuDelete(context.Context, *SystemMenuDeleteRequest) (*SystemMenuDeleteResponse, error)
	SystemMenu(context.Context, *SystemMenuRequest) (*SystemMenuResponse, error)
	SystemMenuRecover(context.Context, *SystemMenuRecoverRequest) (*SystemMenuRecoverResponse, error)
	SystemMenuDrop(context.Context, *SystemMenuDropRequest) (*SystemMenuDropResponse, error)
	SystemMenuList(context.Context, *SystemMenuListRequest) (*SystemMenuListResponse, error)
	SystemMenuListTotal(context.Context, *SystemMenuListTotalRequest) (*SystemMenuTotalResponse, error)
	mustEmbedUnimplementedSystemMenuServiceServer()
}

// UnimplementedSystemMenuServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemMenuServiceServer struct {
}

func (UnimplementedSystemMenuServiceServer) SystemMenuCreate(context.Context, *SystemMenuCreateRequest) (*SystemMenuCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMenuCreate not implemented")
}
func (UnimplementedSystemMenuServiceServer) SystemMenuUpdate(context.Context, *SystemMenuUpdateRequest) (*SystemMenuUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMenuUpdate not implemented")
}
func (UnimplementedSystemMenuServiceServer) SystemMenuDelete(context.Context, *SystemMenuDeleteRequest) (*SystemMenuDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMenuDelete not implemented")
}
func (UnimplementedSystemMenuServiceServer) SystemMenu(context.Context, *SystemMenuRequest) (*SystemMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMenu not implemented")
}
func (UnimplementedSystemMenuServiceServer) SystemMenuRecover(context.Context, *SystemMenuRecoverRequest) (*SystemMenuRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMenuRecover not implemented")
}
func (UnimplementedSystemMenuServiceServer) SystemMenuDrop(context.Context, *SystemMenuDropRequest) (*SystemMenuDropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMenuDrop not implemented")
}
func (UnimplementedSystemMenuServiceServer) SystemMenuList(context.Context, *SystemMenuListRequest) (*SystemMenuListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMenuList not implemented")
}
func (UnimplementedSystemMenuServiceServer) SystemMenuListTotal(context.Context, *SystemMenuListTotalRequest) (*SystemMenuTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMenuListTotal not implemented")
}
func (UnimplementedSystemMenuServiceServer) mustEmbedUnimplementedSystemMenuServiceServer() {}

// UnsafeSystemMenuServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemMenuServiceServer will
// result in compilation errors.
type UnsafeSystemMenuServiceServer interface {
	mustEmbedUnimplementedSystemMenuServiceServer()
}

func RegisterSystemMenuServiceServer(s grpc.ServiceRegistrar, srv SystemMenuServiceServer) {
	s.RegisterService(&SystemMenuService_ServiceDesc, srv)
}

func _SystemMenuService_SystemMenuCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMenuCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMenuServiceServer).SystemMenuCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMenuService_SystemMenuCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMenuServiceServer).SystemMenuCreate(ctx, req.(*SystemMenuCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMenuService_SystemMenuUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMenuUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMenuServiceServer).SystemMenuUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMenuService_SystemMenuUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMenuServiceServer).SystemMenuUpdate(ctx, req.(*SystemMenuUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMenuService_SystemMenuDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMenuDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMenuServiceServer).SystemMenuDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMenuService_SystemMenuDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMenuServiceServer).SystemMenuDelete(ctx, req.(*SystemMenuDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMenuService_SystemMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMenuServiceServer).SystemMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMenuService_SystemMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMenuServiceServer).SystemMenu(ctx, req.(*SystemMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMenuService_SystemMenuRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMenuRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMenuServiceServer).SystemMenuRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMenuService_SystemMenuRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMenuServiceServer).SystemMenuRecover(ctx, req.(*SystemMenuRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMenuService_SystemMenuDrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMenuDropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMenuServiceServer).SystemMenuDrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMenuService_SystemMenuDrop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMenuServiceServer).SystemMenuDrop(ctx, req.(*SystemMenuDropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMenuService_SystemMenuList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMenuListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMenuServiceServer).SystemMenuList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMenuService_SystemMenuList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMenuServiceServer).SystemMenuList(ctx, req.(*SystemMenuListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMenuService_SystemMenuListTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMenuListTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMenuServiceServer).SystemMenuListTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMenuService_SystemMenuListTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMenuServiceServer).SystemMenuListTotal(ctx, req.(*SystemMenuListTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemMenuService_ServiceDesc is the grpc.ServiceDesc for SystemMenuService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemMenuService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "menu.SystemMenuService",
	HandlerType: (*SystemMenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemMenuCreate",
			Handler:    _SystemMenuService_SystemMenuCreate_Handler,
		},
		{
			MethodName: "SystemMenuUpdate",
			Handler:    _SystemMenuService_SystemMenuUpdate_Handler,
		},
		{
			MethodName: "SystemMenuDelete",
			Handler:    _SystemMenuService_SystemMenuDelete_Handler,
		},
		{
			MethodName: "SystemMenu",
			Handler:    _SystemMenuService_SystemMenu_Handler,
		},
		{
			MethodName: "SystemMenuRecover",
			Handler:    _SystemMenuService_SystemMenuRecover_Handler,
		},
		{
			MethodName: "SystemMenuDrop",
			Handler:    _SystemMenuService_SystemMenuDrop_Handler,
		},
		{
			MethodName: "SystemMenuList",
			Handler:    _SystemMenuService_SystemMenuList_Handler,
		},
		{
			MethodName: "SystemMenuListTotal",
			Handler:    _SystemMenuService_SystemMenuListTotal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_menu.proto",
}
