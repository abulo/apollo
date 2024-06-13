// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: system_mail_template.proto

// system_mail_template 邮件模版表

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
	SystemMailTemplateService_SystemMailTemplateCreate_FullMethodName    = "/mail.SystemMailTemplateService/SystemMailTemplateCreate"
	SystemMailTemplateService_SystemMailTemplateUpdate_FullMethodName    = "/mail.SystemMailTemplateService/SystemMailTemplateUpdate"
	SystemMailTemplateService_SystemMailTemplateDelete_FullMethodName    = "/mail.SystemMailTemplateService/SystemMailTemplateDelete"
	SystemMailTemplateService_SystemMailTemplate_FullMethodName          = "/mail.SystemMailTemplateService/SystemMailTemplate"
	SystemMailTemplateService_SystemMailTemplateRecover_FullMethodName   = "/mail.SystemMailTemplateService/SystemMailTemplateRecover"
	SystemMailTemplateService_SystemMailTemplateList_FullMethodName      = "/mail.SystemMailTemplateService/SystemMailTemplateList"
	SystemMailTemplateService_SystemMailTemplateListTotal_FullMethodName = "/mail.SystemMailTemplateService/SystemMailTemplateListTotal"
)

// SystemMailTemplateServiceClient is the client API for SystemMailTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SystemMailTemplateService 服务
type SystemMailTemplateServiceClient interface {
	SystemMailTemplateCreate(ctx context.Context, in *SystemMailTemplateCreateRequest, opts ...grpc.CallOption) (*SystemMailTemplateCreateResponse, error)
	SystemMailTemplateUpdate(ctx context.Context, in *SystemMailTemplateUpdateRequest, opts ...grpc.CallOption) (*SystemMailTemplateUpdateResponse, error)
	SystemMailTemplateDelete(ctx context.Context, in *SystemMailTemplateDeleteRequest, opts ...grpc.CallOption) (*SystemMailTemplateDeleteResponse, error)
	SystemMailTemplate(ctx context.Context, in *SystemMailTemplateRequest, opts ...grpc.CallOption) (*SystemMailTemplateResponse, error)
	SystemMailTemplateRecover(ctx context.Context, in *SystemMailTemplateRecoverRequest, opts ...grpc.CallOption) (*SystemMailTemplateRecoverResponse, error)
	SystemMailTemplateList(ctx context.Context, in *SystemMailTemplateListRequest, opts ...grpc.CallOption) (*SystemMailTemplateListResponse, error)
	SystemMailTemplateListTotal(ctx context.Context, in *SystemMailTemplateListTotalRequest, opts ...grpc.CallOption) (*SystemMailTemplateTotalResponse, error)
}

type systemMailTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemMailTemplateServiceClient(cc grpc.ClientConnInterface) SystemMailTemplateServiceClient {
	return &systemMailTemplateServiceClient{cc}
}

func (c *systemMailTemplateServiceClient) SystemMailTemplateCreate(ctx context.Context, in *SystemMailTemplateCreateRequest, opts ...grpc.CallOption) (*SystemMailTemplateCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailTemplateCreateResponse)
	err := c.cc.Invoke(ctx, SystemMailTemplateService_SystemMailTemplateCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailTemplateServiceClient) SystemMailTemplateUpdate(ctx context.Context, in *SystemMailTemplateUpdateRequest, opts ...grpc.CallOption) (*SystemMailTemplateUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailTemplateUpdateResponse)
	err := c.cc.Invoke(ctx, SystemMailTemplateService_SystemMailTemplateUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailTemplateServiceClient) SystemMailTemplateDelete(ctx context.Context, in *SystemMailTemplateDeleteRequest, opts ...grpc.CallOption) (*SystemMailTemplateDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailTemplateDeleteResponse)
	err := c.cc.Invoke(ctx, SystemMailTemplateService_SystemMailTemplateDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailTemplateServiceClient) SystemMailTemplate(ctx context.Context, in *SystemMailTemplateRequest, opts ...grpc.CallOption) (*SystemMailTemplateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailTemplateResponse)
	err := c.cc.Invoke(ctx, SystemMailTemplateService_SystemMailTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailTemplateServiceClient) SystemMailTemplateRecover(ctx context.Context, in *SystemMailTemplateRecoverRequest, opts ...grpc.CallOption) (*SystemMailTemplateRecoverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailTemplateRecoverResponse)
	err := c.cc.Invoke(ctx, SystemMailTemplateService_SystemMailTemplateRecover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailTemplateServiceClient) SystemMailTemplateList(ctx context.Context, in *SystemMailTemplateListRequest, opts ...grpc.CallOption) (*SystemMailTemplateListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailTemplateListResponse)
	err := c.cc.Invoke(ctx, SystemMailTemplateService_SystemMailTemplateList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMailTemplateServiceClient) SystemMailTemplateListTotal(ctx context.Context, in *SystemMailTemplateListTotalRequest, opts ...grpc.CallOption) (*SystemMailTemplateTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemMailTemplateTotalResponse)
	err := c.cc.Invoke(ctx, SystemMailTemplateService_SystemMailTemplateListTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemMailTemplateServiceServer is the server API for SystemMailTemplateService service.
// All implementations must embed UnimplementedSystemMailTemplateServiceServer
// for forward compatibility
//
// SystemMailTemplateService 服务
type SystemMailTemplateServiceServer interface {
	SystemMailTemplateCreate(context.Context, *SystemMailTemplateCreateRequest) (*SystemMailTemplateCreateResponse, error)
	SystemMailTemplateUpdate(context.Context, *SystemMailTemplateUpdateRequest) (*SystemMailTemplateUpdateResponse, error)
	SystemMailTemplateDelete(context.Context, *SystemMailTemplateDeleteRequest) (*SystemMailTemplateDeleteResponse, error)
	SystemMailTemplate(context.Context, *SystemMailTemplateRequest) (*SystemMailTemplateResponse, error)
	SystemMailTemplateRecover(context.Context, *SystemMailTemplateRecoverRequest) (*SystemMailTemplateRecoverResponse, error)
	SystemMailTemplateList(context.Context, *SystemMailTemplateListRequest) (*SystemMailTemplateListResponse, error)
	SystemMailTemplateListTotal(context.Context, *SystemMailTemplateListTotalRequest) (*SystemMailTemplateTotalResponse, error)
	mustEmbedUnimplementedSystemMailTemplateServiceServer()
}

// UnimplementedSystemMailTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemMailTemplateServiceServer struct {
}

func (UnimplementedSystemMailTemplateServiceServer) SystemMailTemplateCreate(context.Context, *SystemMailTemplateCreateRequest) (*SystemMailTemplateCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailTemplateCreate not implemented")
}
func (UnimplementedSystemMailTemplateServiceServer) SystemMailTemplateUpdate(context.Context, *SystemMailTemplateUpdateRequest) (*SystemMailTemplateUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailTemplateUpdate not implemented")
}
func (UnimplementedSystemMailTemplateServiceServer) SystemMailTemplateDelete(context.Context, *SystemMailTemplateDeleteRequest) (*SystemMailTemplateDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailTemplateDelete not implemented")
}
func (UnimplementedSystemMailTemplateServiceServer) SystemMailTemplate(context.Context, *SystemMailTemplateRequest) (*SystemMailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailTemplate not implemented")
}
func (UnimplementedSystemMailTemplateServiceServer) SystemMailTemplateRecover(context.Context, *SystemMailTemplateRecoverRequest) (*SystemMailTemplateRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailTemplateRecover not implemented")
}
func (UnimplementedSystemMailTemplateServiceServer) SystemMailTemplateList(context.Context, *SystemMailTemplateListRequest) (*SystemMailTemplateListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailTemplateList not implemented")
}
func (UnimplementedSystemMailTemplateServiceServer) SystemMailTemplateListTotal(context.Context, *SystemMailTemplateListTotalRequest) (*SystemMailTemplateTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemMailTemplateListTotal not implemented")
}
func (UnimplementedSystemMailTemplateServiceServer) mustEmbedUnimplementedSystemMailTemplateServiceServer() {
}

// UnsafeSystemMailTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemMailTemplateServiceServer will
// result in compilation errors.
type UnsafeSystemMailTemplateServiceServer interface {
	mustEmbedUnimplementedSystemMailTemplateServiceServer()
}

func RegisterSystemMailTemplateServiceServer(s grpc.ServiceRegistrar, srv SystemMailTemplateServiceServer) {
	s.RegisterService(&SystemMailTemplateService_ServiceDesc, srv)
}

func _SystemMailTemplateService_SystemMailTemplateCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailTemplateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplateCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailTemplateService_SystemMailTemplateCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplateCreate(ctx, req.(*SystemMailTemplateCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailTemplateService_SystemMailTemplateUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailTemplateUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplateUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailTemplateService_SystemMailTemplateUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplateUpdate(ctx, req.(*SystemMailTemplateUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailTemplateService_SystemMailTemplateDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailTemplateDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplateDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailTemplateService_SystemMailTemplateDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplateDelete(ctx, req.(*SystemMailTemplateDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailTemplateService_SystemMailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailTemplateService_SystemMailTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplate(ctx, req.(*SystemMailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailTemplateService_SystemMailTemplateRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailTemplateRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplateRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailTemplateService_SystemMailTemplateRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplateRecover(ctx, req.(*SystemMailTemplateRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailTemplateService_SystemMailTemplateList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailTemplateListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplateList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailTemplateService_SystemMailTemplateList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplateList(ctx, req.(*SystemMailTemplateListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMailTemplateService_SystemMailTemplateListTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMailTemplateListTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplateListTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMailTemplateService_SystemMailTemplateListTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMailTemplateServiceServer).SystemMailTemplateListTotal(ctx, req.(*SystemMailTemplateListTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemMailTemplateService_ServiceDesc is the grpc.ServiceDesc for SystemMailTemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemMailTemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mail.SystemMailTemplateService",
	HandlerType: (*SystemMailTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemMailTemplateCreate",
			Handler:    _SystemMailTemplateService_SystemMailTemplateCreate_Handler,
		},
		{
			MethodName: "SystemMailTemplateUpdate",
			Handler:    _SystemMailTemplateService_SystemMailTemplateUpdate_Handler,
		},
		{
			MethodName: "SystemMailTemplateDelete",
			Handler:    _SystemMailTemplateService_SystemMailTemplateDelete_Handler,
		},
		{
			MethodName: "SystemMailTemplate",
			Handler:    _SystemMailTemplateService_SystemMailTemplate_Handler,
		},
		{
			MethodName: "SystemMailTemplateRecover",
			Handler:    _SystemMailTemplateService_SystemMailTemplateRecover_Handler,
		},
		{
			MethodName: "SystemMailTemplateList",
			Handler:    _SystemMailTemplateService_SystemMailTemplateList_Handler,
		},
		{
			MethodName: "SystemMailTemplateListTotal",
			Handler:    _SystemMailTemplateService_SystemMailTemplateListTotal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_mail_template.proto",
}
