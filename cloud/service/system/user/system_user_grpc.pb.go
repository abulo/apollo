// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: system_user.proto

// system_user 系统用户

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
	SystemUserService_SystemUserCreate_FullMethodName        = "/user.SystemUserService/SystemUserCreate"
	SystemUserService_SystemUserUpdate_FullMethodName        = "/user.SystemUserService/SystemUserUpdate"
	SystemUserService_SystemUserDelete_FullMethodName        = "/user.SystemUserService/SystemUserDelete"
	SystemUserService_SystemUser_FullMethodName              = "/user.SystemUserService/SystemUser"
	SystemUserService_SystemUserRecover_FullMethodName       = "/user.SystemUserService/SystemUserRecover"
	SystemUserService_SystemUserLogin_FullMethodName         = "/user.SystemUserService/SystemUserLogin"
	SystemUserService_SystemUserList_FullMethodName          = "/user.SystemUserService/SystemUserList"
	SystemUserService_SystemUserListTotal_FullMethodName     = "/user.SystemUserService/SystemUserListTotal"
	SystemUserService_SystemUserPassword_FullMethodName      = "/user.SystemUserService/SystemUserPassword"
	SystemUserService_SystemUserRoleDataScope_FullMethodName = "/user.SystemUserService/SystemUserRoleDataScope"
	SystemUserService_SystemRoleDrop_FullMethodName          = "/user.SystemUserService/SystemRoleDrop"
)

// SystemUserServiceClient is the client API for SystemUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SystemUserService 服务
type SystemUserServiceClient interface {
	SystemUserCreate(ctx context.Context, in *SystemUserCreateRequest, opts ...grpc.CallOption) (*SystemUserCreateResponse, error)
	SystemUserUpdate(ctx context.Context, in *SystemUserUpdateRequest, opts ...grpc.CallOption) (*SystemUserUpdateResponse, error)
	SystemUserDelete(ctx context.Context, in *SystemUserDeleteRequest, opts ...grpc.CallOption) (*SystemUserDeleteResponse, error)
	SystemUser(ctx context.Context, in *SystemUserRequest, opts ...grpc.CallOption) (*SystemUserResponse, error)
	SystemUserRecover(ctx context.Context, in *SystemUserRecoverRequest, opts ...grpc.CallOption) (*SystemUserRecoverResponse, error)
	SystemUserLogin(ctx context.Context, in *SystemUserLoginRequest, opts ...grpc.CallOption) (*SystemUserLoginResponse, error)
	SystemUserList(ctx context.Context, in *SystemUserListRequest, opts ...grpc.CallOption) (*SystemUserListResponse, error)
	SystemUserListTotal(ctx context.Context, in *SystemUserListTotalRequest, opts ...grpc.CallOption) (*SystemUserTotalResponse, error)
	SystemUserPassword(ctx context.Context, in *SystemUserPasswordRequest, opts ...grpc.CallOption) (*SystemUserPasswordResponse, error)
	SystemUserRoleDataScope(ctx context.Context, in *SystemUserRoleDataScopeRequest, opts ...grpc.CallOption) (*SystemUserRoleDataScopeResponse, error)
	SystemRoleDrop(ctx context.Context, in *SystemUserDropRequest, opts ...grpc.CallOption) (*SystemUserDropResponse, error)
}

type systemUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemUserServiceClient(cc grpc.ClientConnInterface) SystemUserServiceClient {
	return &systemUserServiceClient{cc}
}

func (c *systemUserServiceClient) SystemUserCreate(ctx context.Context, in *SystemUserCreateRequest, opts ...grpc.CallOption) (*SystemUserCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserCreateResponse)
	err := c.cc.Invoke(ctx, SystemUserService_SystemUserCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) SystemUserUpdate(ctx context.Context, in *SystemUserUpdateRequest, opts ...grpc.CallOption) (*SystemUserUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserUpdateResponse)
	err := c.cc.Invoke(ctx, SystemUserService_SystemUserUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) SystemUserDelete(ctx context.Context, in *SystemUserDeleteRequest, opts ...grpc.CallOption) (*SystemUserDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserDeleteResponse)
	err := c.cc.Invoke(ctx, SystemUserService_SystemUserDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) SystemUser(ctx context.Context, in *SystemUserRequest, opts ...grpc.CallOption) (*SystemUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserResponse)
	err := c.cc.Invoke(ctx, SystemUserService_SystemUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) SystemUserRecover(ctx context.Context, in *SystemUserRecoverRequest, opts ...grpc.CallOption) (*SystemUserRecoverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserRecoverResponse)
	err := c.cc.Invoke(ctx, SystemUserService_SystemUserRecover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) SystemUserLogin(ctx context.Context, in *SystemUserLoginRequest, opts ...grpc.CallOption) (*SystemUserLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserLoginResponse)
	err := c.cc.Invoke(ctx, SystemUserService_SystemUserLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) SystemUserList(ctx context.Context, in *SystemUserListRequest, opts ...grpc.CallOption) (*SystemUserListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserListResponse)
	err := c.cc.Invoke(ctx, SystemUserService_SystemUserList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) SystemUserListTotal(ctx context.Context, in *SystemUserListTotalRequest, opts ...grpc.CallOption) (*SystemUserTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserTotalResponse)
	err := c.cc.Invoke(ctx, SystemUserService_SystemUserListTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) SystemUserPassword(ctx context.Context, in *SystemUserPasswordRequest, opts ...grpc.CallOption) (*SystemUserPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserPasswordResponse)
	err := c.cc.Invoke(ctx, SystemUserService_SystemUserPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) SystemUserRoleDataScope(ctx context.Context, in *SystemUserRoleDataScopeRequest, opts ...grpc.CallOption) (*SystemUserRoleDataScopeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserRoleDataScopeResponse)
	err := c.cc.Invoke(ctx, SystemUserService_SystemUserRoleDataScope_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) SystemRoleDrop(ctx context.Context, in *SystemUserDropRequest, opts ...grpc.CallOption) (*SystemUserDropResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemUserDropResponse)
	err := c.cc.Invoke(ctx, SystemUserService_SystemRoleDrop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemUserServiceServer is the server API for SystemUserService service.
// All implementations must embed UnimplementedSystemUserServiceServer
// for forward compatibility
//
// SystemUserService 服务
type SystemUserServiceServer interface {
	SystemUserCreate(context.Context, *SystemUserCreateRequest) (*SystemUserCreateResponse, error)
	SystemUserUpdate(context.Context, *SystemUserUpdateRequest) (*SystemUserUpdateResponse, error)
	SystemUserDelete(context.Context, *SystemUserDeleteRequest) (*SystemUserDeleteResponse, error)
	SystemUser(context.Context, *SystemUserRequest) (*SystemUserResponse, error)
	SystemUserRecover(context.Context, *SystemUserRecoverRequest) (*SystemUserRecoverResponse, error)
	SystemUserLogin(context.Context, *SystemUserLoginRequest) (*SystemUserLoginResponse, error)
	SystemUserList(context.Context, *SystemUserListRequest) (*SystemUserListResponse, error)
	SystemUserListTotal(context.Context, *SystemUserListTotalRequest) (*SystemUserTotalResponse, error)
	SystemUserPassword(context.Context, *SystemUserPasswordRequest) (*SystemUserPasswordResponse, error)
	SystemUserRoleDataScope(context.Context, *SystemUserRoleDataScopeRequest) (*SystemUserRoleDataScopeResponse, error)
	SystemRoleDrop(context.Context, *SystemUserDropRequest) (*SystemUserDropResponse, error)
	mustEmbedUnimplementedSystemUserServiceServer()
}

// UnimplementedSystemUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemUserServiceServer struct {
}

func (UnimplementedSystemUserServiceServer) SystemUserCreate(context.Context, *SystemUserCreateRequest) (*SystemUserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserCreate not implemented")
}
func (UnimplementedSystemUserServiceServer) SystemUserUpdate(context.Context, *SystemUserUpdateRequest) (*SystemUserUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserUpdate not implemented")
}
func (UnimplementedSystemUserServiceServer) SystemUserDelete(context.Context, *SystemUserDeleteRequest) (*SystemUserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserDelete not implemented")
}
func (UnimplementedSystemUserServiceServer) SystemUser(context.Context, *SystemUserRequest) (*SystemUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUser not implemented")
}
func (UnimplementedSystemUserServiceServer) SystemUserRecover(context.Context, *SystemUserRecoverRequest) (*SystemUserRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserRecover not implemented")
}
func (UnimplementedSystemUserServiceServer) SystemUserLogin(context.Context, *SystemUserLoginRequest) (*SystemUserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserLogin not implemented")
}
func (UnimplementedSystemUserServiceServer) SystemUserList(context.Context, *SystemUserListRequest) (*SystemUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserList not implemented")
}
func (UnimplementedSystemUserServiceServer) SystemUserListTotal(context.Context, *SystemUserListTotalRequest) (*SystemUserTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserListTotal not implemented")
}
func (UnimplementedSystemUserServiceServer) SystemUserPassword(context.Context, *SystemUserPasswordRequest) (*SystemUserPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserPassword not implemented")
}
func (UnimplementedSystemUserServiceServer) SystemUserRoleDataScope(context.Context, *SystemUserRoleDataScopeRequest) (*SystemUserRoleDataScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserRoleDataScope not implemented")
}
func (UnimplementedSystemUserServiceServer) SystemRoleDrop(context.Context, *SystemUserDropRequest) (*SystemUserDropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemRoleDrop not implemented")
}
func (UnimplementedSystemUserServiceServer) mustEmbedUnimplementedSystemUserServiceServer() {}

// UnsafeSystemUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemUserServiceServer will
// result in compilation errors.
type UnsafeSystemUserServiceServer interface {
	mustEmbedUnimplementedSystemUserServiceServer()
}

func RegisterSystemUserServiceServer(s grpc.ServiceRegistrar, srv SystemUserServiceServer) {
	s.RegisterService(&SystemUserService_ServiceDesc, srv)
}

func _SystemUserService_SystemUserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).SystemUserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_SystemUserCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).SystemUserCreate(ctx, req.(*SystemUserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_SystemUserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).SystemUserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_SystemUserUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).SystemUserUpdate(ctx, req.(*SystemUserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_SystemUserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).SystemUserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_SystemUserDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).SystemUserDelete(ctx, req.(*SystemUserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_SystemUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).SystemUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_SystemUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).SystemUser(ctx, req.(*SystemUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_SystemUserRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).SystemUserRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_SystemUserRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).SystemUserRecover(ctx, req.(*SystemUserRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_SystemUserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).SystemUserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_SystemUserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).SystemUserLogin(ctx, req.(*SystemUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_SystemUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).SystemUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_SystemUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).SystemUserList(ctx, req.(*SystemUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_SystemUserListTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserListTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).SystemUserListTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_SystemUserListTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).SystemUserListTotal(ctx, req.(*SystemUserListTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_SystemUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).SystemUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_SystemUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).SystemUserPassword(ctx, req.(*SystemUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_SystemUserRoleDataScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserRoleDataScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).SystemUserRoleDataScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_SystemUserRoleDataScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).SystemUserRoleDataScope(ctx, req.(*SystemUserRoleDataScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_SystemRoleDrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserDropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).SystemRoleDrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserService_SystemRoleDrop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).SystemRoleDrop(ctx, req.(*SystemUserDropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemUserService_ServiceDesc is the grpc.ServiceDesc for SystemUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.SystemUserService",
	HandlerType: (*SystemUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemUserCreate",
			Handler:    _SystemUserService_SystemUserCreate_Handler,
		},
		{
			MethodName: "SystemUserUpdate",
			Handler:    _SystemUserService_SystemUserUpdate_Handler,
		},
		{
			MethodName: "SystemUserDelete",
			Handler:    _SystemUserService_SystemUserDelete_Handler,
		},
		{
			MethodName: "SystemUser",
			Handler:    _SystemUserService_SystemUser_Handler,
		},
		{
			MethodName: "SystemUserRecover",
			Handler:    _SystemUserService_SystemUserRecover_Handler,
		},
		{
			MethodName: "SystemUserLogin",
			Handler:    _SystemUserService_SystemUserLogin_Handler,
		},
		{
			MethodName: "SystemUserList",
			Handler:    _SystemUserService_SystemUserList_Handler,
		},
		{
			MethodName: "SystemUserListTotal",
			Handler:    _SystemUserService_SystemUserListTotal_Handler,
		},
		{
			MethodName: "SystemUserPassword",
			Handler:    _SystemUserService_SystemUserPassword_Handler,
		},
		{
			MethodName: "SystemUserRoleDataScope",
			Handler:    _SystemUserService_SystemUserRoleDataScope_Handler,
		},
		{
			MethodName: "SystemRoleDrop",
			Handler:    _SystemUserService_SystemRoleDrop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_user.proto",
}
