// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: system_mail_account.proto

// system_mail_account 邮箱账号表

package mail

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
	SystemMailAccountService_SystemMailAccountCreate_FullMethodName    = "/mail.SystemMailAccountService/SystemMailAccountCreate"
	SystemMailAccountService_SystemMailAccountUpdate_FullMethodName    = "/mail.SystemMailAccountService/SystemMailAccountUpdate"
	SystemMailAccountService_SystemMailAccountDelete_FullMethodName    = "/mail.SystemMailAccountService/SystemMailAccountDelete"
	SystemMailAccountService_SystemMailAccount_FullMethodName          = "/mail.SystemMailAccountService/SystemMailAccount"
	SystemMailAccountService_SystemMailAccountRecover_FullMethodName   = "/mail.SystemMailAccountService/SystemMailAccountRecover"
	SystemMailAccountService_SystemMailAccountDrop_FullMethodName      = "/mail.SystemMailAccountService/SystemMailAccountDrop"
	SystemMailAccountService_SystemMailAccountList_FullMethodName      = "/mail.SystemMailAccountService/SystemMailAccountList"
	SystemMailAccountService_SystemMailAccountListTotal_FullMethodName = "/mail.SystemMailAccountService/SystemMailAccountListTotal"
)

// SystemMailAccountServiceClient is the client API for SystemMailAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SystemMailAccountService 服务
type SystemMailAccountServiceClient interface {
	SystemMailAccountCreate(ctx context.Context, in *SystemMailAccountCreateRequest, opts ...grpc.CallOption) (*SystemMailAccountCreateResponse, error)
	SystemMailAccountUpdate(ctx context.Context, in *SystemMailAccountUpdateRequest, opts ...grpc.CallOption) (*SystemMailAccountUpdateResponse, error)
	SystemMailAccountDelete(ctx context.Context, in *SystemMailAccountDeleteRequest, opts ...grpc.CallOption) (*SystemMailAccountDeleteResponse, error)
	SystemMailAccount(ctx context.Context, in *SystemMailAccountRequest, opts ...grpc.CallOption) (*SystemMailAccountResponse, error)
	SystemMailAccountRecover(ctx context.Context, in *SystemMailAccountRecoverRequest, opts ...grpc.CallOption) (*SystemMailAccountRecoverResponse, error)
	SystemMailAccountDrop(ctx context.Context, in *SystemMailAccountDropRequest, opts ...grpc.CallOption) (*SystemMailAccountDropResponse, error)
	SystemMailAccountList(ctx context.Context, in *SystemMailAccountListRequest, opts ...grpc.CallOption) (*SystemMailAccountListResponse, error)
	SystemMailAccountListTotal(ctx context.Context, in *SystemMailAccountListTotalRequest, opts ...grpc.CallOption) (*SystemMailAccountTotalResponse, error)
}

type systemMailAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemMailAccountServiceClient(cc grpc.ClientConnInterface) SystemMailAccountServiceClient {
	return &systemMailAccountServiceClient{cc}
}

func (c *systemMailAccountServiceClient) SystemMailAccountCreate(ctx context.Context, in *SystemMailAccountCreateRequest, opts ...grpc.CallOption) (*SystemMailAccountCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailAccountCreateResponse)
	err := c.cc.Invoke(ctx, SystemMailAccountService_SystemMailAccountCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailAccountServiceClient) SystemMailAccountUpdate(ctx context.Context, in *SystemMailAccountUpdateRequest, opts ...grpc.CallOption) (*SystemMailAccountUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailAccountUpdateResponse)
	err := c.cc.Invoke(ctx, SystemMailAccountService_SystemMailAccountUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailAccountServiceClient) SystemMailAccountDelete(ctx context.Context, in *SystemMailAccountDeleteRequest, opts ...grpc.CallOption) (*SystemMailAccountDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailAccountDeleteResponse)
	err := c.cc.Invoke(ctx, SystemMailAccountService_SystemMailAccountDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailAccountServiceClient) SystemMailAccount(ctx context.Context, in *SystemMailAccountRequest, opts ...grpc.CallOption) (*SystemMailAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailAccountResponse)
	err := c.cc.Invoke(ctx, SystemMailAccountService_SystemMailAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailAccountServiceClient) SystemMailAccountRecover(ctx context.Context, in *SystemMailAccountRecoverRequest, opts ...grpc.CallOption) (*SystemMailAccountRecoverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailAccountRecoverResponse)
	err := c.cc.Invoke(ctx, SystemMailAccountService_SystemMailAccountRecover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailAccountServiceClient) SystemMailAccountDrop(ctx context.Context, in *SystemMailAccountDropRequest, opts ...grpc.CallOption) (*SystemMailAccountDropResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailAccountDropResponse)
	err := c.cc.Invoke(ctx, SystemMailAccountService_SystemMailAccountDrop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailAccountServiceClient) SystemMailAccountList(ctx context.Context, in *SystemMailAccountListRequest, opts ...grpc.CallOption) (*SystemMailAccountListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailAccountListResponse)
	err := c.cc.Invoke(ctx, SystemMailAccountService_SystemMailAccountList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailAccountServiceClient) SystemMailAccountListTotal(ctx context.Context, in *SystemMailAccountListTotalRequest, opts ...grpc.CallOption) (*SystemMailAccountTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailAccountTotalResponse)
	err := c.cc.Invoke(ctx, SystemMailAccountService_SystemMailAccountListTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemMailAccountServiceServer is the server API for SystemMailAccountService service.
// All implementations must embed UnimplementedSystemMailAccountServiceServer
// for forward compatibility
//
// SystemMailAccountService 服务
type SystemMailAccountServiceServer interface {
	SystemMailAccountCreate(context.Context, *SystemMailAccountCreateRequest) (*SystemMailAccountCreateResponse, error)
	SystemMailAccountUpdate(context.Context, *SystemMailAccountUpdateRequest) (*SystemMailAccountUpdateResponse, error)
	SystemMailAccountDelete(context.Context, *SystemMailAccountDeleteRequest) (*SystemMailAccountDeleteResponse, error)
	SystemMailAccount(context.Context, *SystemMailAccountRequest) (*SystemMailAccountResponse, error)
	SystemMailAccountRecover(context.Context, *SystemMailAccountRecoverRequest) (*SystemMailAccountRecoverResponse, error)
	SystemMailAccountDrop(context.Context, *SystemMailAccountDropRequest) (*SystemMailAccountDropResponse, error)
	SystemMailAccountList(context.Context, *SystemMailAccountListRequest) (*SystemMailAccountListResponse, error)
	SystemMailAccountListTotal(context.Context, *SystemMailAccountListTotalRequest) (*SystemMailAccountTotalResponse, error)
	mustEmbedUnimplementedSystemMailAccountServiceServer()
}

// UnimplementedSystemMailAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemMailAccountServiceServer struct {
}

func (UnimplementedSystemMailAccountServiceServer) SystemMailAccountCreate(context.Context, *SystemMailAccountCreateRequest) (*SystemMailAccountCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailAccountCreate not implemented")
}
func (UnimplementedSystemMailAccountServiceServer) SystemMailAccountUpdate(context.Context, *SystemMailAccountUpdateRequest) (*SystemMailAccountUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailAccountUpdate not implemented")
}
func (UnimplementedSystemMailAccountServiceServer) SystemMailAccountDelete(context.Context, *SystemMailAccountDeleteRequest) (*SystemMailAccountDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailAccountDelete not implemented")
}
func (UnimplementedSystemMailAccountServiceServer) SystemMailAccount(context.Context, *SystemMailAccountRequest) (*SystemMailAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailAccount not implemented")
}
func (UnimplementedSystemMailAccountServiceServer) SystemMailAccountRecover(context.Context, *SystemMailAccountRecoverRequest) (*SystemMailAccountRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailAccountRecover not implemented")
}
func (UnimplementedSystemMailAccountServiceServer) SystemMailAccountDrop(context.Context, *SystemMailAccountDropRequest) (*SystemMailAccountDropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailAccountDrop not implemented")
}
func (UnimplementedSystemMailAccountServiceServer) SystemMailAccountList(context.Context, *SystemMailAccountListRequest) (*SystemMailAccountListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailAccountList not implemented")
}
func (UnimplementedSystemMailAccountServiceServer) SystemMailAccountListTotal(context.Context, *SystemMailAccountListTotalRequest) (*SystemMailAccountTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailAccountListTotal not implemented")
}
func (UnimplementedSystemMailAccountServiceServer) mustEmbedUnimplementedSystemMailAccountServiceServer() {
}

// UnsafeSystemMailAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemMailAccountServiceServer will
// result in compilation errors.
type UnsafeSystemMailAccountServiceServer interface {
	mustEmbedUnimplementedSystemMailAccountServiceServer()
}

func RegisterSystemMailAccountServiceServer(s grpc.ServiceRegistrar, srv SystemMailAccountServiceServer) {
	s.RegisterService(&SystemMailAccountService_ServiceDesc, srv)
}

func _SystemMailAccountService_SystemMailAccountCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailAccountCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailAccountService_SystemMailAccountCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountCreate(ctx, req.(*SystemMailAccountCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailAccountService_SystemMailAccountUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailAccountUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailAccountService_SystemMailAccountUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountUpdate(ctx, req.(*SystemMailAccountUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailAccountService_SystemMailAccountDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailAccountDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailAccountService_SystemMailAccountDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountDelete(ctx, req.(*SystemMailAccountDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailAccountService_SystemMailAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailAccountServiceServer).SystemMailAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailAccountService_SystemMailAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailAccountServiceServer).SystemMailAccount(ctx, req.(*SystemMailAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailAccountService_SystemMailAccountRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailAccountRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailAccountService_SystemMailAccountRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountRecover(ctx, req.(*SystemMailAccountRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailAccountService_SystemMailAccountDrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailAccountDropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountDrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailAccountService_SystemMailAccountDrop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountDrop(ctx, req.(*SystemMailAccountDropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailAccountService_SystemMailAccountList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailAccountListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailAccountService_SystemMailAccountList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountList(ctx, req.(*SystemMailAccountListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailAccountService_SystemMailAccountListTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailAccountListTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountListTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailAccountService_SystemMailAccountListTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailAccountServiceServer).SystemMailAccountListTotal(ctx, req.(*SystemMailAccountListTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemMailAccountService_ServiceDesc is the grpc.ServiceDesc for SystemMailAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemMailAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mail.SystemMailAccountService",
	HandlerType: (*SystemMailAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemMailAccountCreate",
			Handler:    _SystemMailAccountService_SystemMailAccountCreate_Handler,
		},
		{
			MethodName: "SystemMailAccountUpdate",
			Handler:    _SystemMailAccountService_SystemMailAccountUpdate_Handler,
		},
		{
			MethodName: "SystemMailAccountDelete",
			Handler:    _SystemMailAccountService_SystemMailAccountDelete_Handler,
		},
		{
			MethodName: "SystemMailAccount",
			Handler:    _SystemMailAccountService_SystemMailAccount_Handler,
		},
		{
			MethodName: "SystemMailAccountRecover",
			Handler:    _SystemMailAccountService_SystemMailAccountRecover_Handler,
		},
		{
			MethodName: "SystemMailAccountDrop",
			Handler:    _SystemMailAccountService_SystemMailAccountDrop_Handler,
		},
		{
			MethodName: "SystemMailAccountList",
			Handler:    _SystemMailAccountService_SystemMailAccountList_Handler,
		},
		{
			MethodName: "SystemMailAccountListTotal",
			Handler:    _SystemMailAccountService_SystemMailAccountListTotal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_mail_account.proto",
}
