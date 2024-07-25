// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: system_operate_log.proto

// system_operate_log 操作日志

package logger

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
	SystemOperateLogService_SystemOperateLogCreate_FullMethodName          = "/logger.SystemOperateLogService/SystemOperateLogCreate"
	SystemOperateLogService_SystemOperateLogUpdate_FullMethodName          = "/logger.SystemOperateLogService/SystemOperateLogUpdate"
	SystemOperateLogService_SystemOperateLogDelete_FullMethodName          = "/logger.SystemOperateLogService/SystemOperateLogDelete"
	SystemOperateLogService_SystemOperateLog_FullMethodName                = "/logger.SystemOperateLogService/SystemOperateLog"
	SystemOperateLogService_SystemOperateLogRecover_FullMethodName         = "/logger.SystemOperateLogService/SystemOperateLogRecover"
	SystemOperateLogService_SystemOperateLogDrop_FullMethodName            = "/logger.SystemOperateLogService/SystemOperateLogDrop"
	SystemOperateLogService_SystemOperateLogList_FullMethodName            = "/logger.SystemOperateLogService/SystemOperateLogList"
	SystemOperateLogService_SystemOperateLogListTotal_FullMethodName       = "/logger.SystemOperateLogService/SystemOperateLogListTotal"
	SystemOperateLogService_SystemOperateLogMultipleDelete_FullMethodName  = "/logger.SystemOperateLogService/SystemOperateLogMultipleDelete"
	SystemOperateLogService_SystemOperateLogMultipleRecover_FullMethodName = "/logger.SystemOperateLogService/SystemOperateLogMultipleRecover"
	SystemOperateLogService_SystemOperateLogMultipleDrop_FullMethodName    = "/logger.SystemOperateLogService/SystemOperateLogMultipleDrop"
)

// SystemOperateLogServiceClient is the client API for SystemOperateLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SystemOperateLogService 服务
type SystemOperateLogServiceClient interface {
	SystemOperateLogCreate(ctx context.Context, in *SystemOperateLogCreateRequest, opts ...grpc.CallOption) (*SystemOperateLogCreateResponse, error)
	SystemOperateLogUpdate(ctx context.Context, in *SystemOperateLogUpdateRequest, opts ...grpc.CallOption) (*SystemOperateLogUpdateResponse, error)
	SystemOperateLogDelete(ctx context.Context, in *SystemOperateLogDeleteRequest, opts ...grpc.CallOption) (*SystemOperateLogDeleteResponse, error)
	SystemOperateLog(ctx context.Context, in *SystemOperateLogRequest, opts ...grpc.CallOption) (*SystemOperateLogResponse, error)
	SystemOperateLogRecover(ctx context.Context, in *SystemOperateLogRecoverRequest, opts ...grpc.CallOption) (*SystemOperateLogRecoverResponse, error)
	SystemOperateLogDrop(ctx context.Context, in *SystemOperateLogDropRequest, opts ...grpc.CallOption) (*SystemOperateLogDropResponse, error)
	SystemOperateLogList(ctx context.Context, in *SystemOperateLogListRequest, opts ...grpc.CallOption) (*SystemOperateLogListResponse, error)
	SystemOperateLogListTotal(ctx context.Context, in *SystemOperateLogListTotalRequest, opts ...grpc.CallOption) (*SystemOperateLogTotalResponse, error)
	SystemOperateLogMultipleDelete(ctx context.Context, in *SystemOperateLogMultipleRequest, opts ...grpc.CallOption) (*SystemOperateLogMultipleResponse, error)
	SystemOperateLogMultipleRecover(ctx context.Context, in *SystemOperateLogMultipleRequest, opts ...grpc.CallOption) (*SystemOperateLogMultipleResponse, error)
	SystemOperateLogMultipleDrop(ctx context.Context, in *SystemOperateLogMultipleRequest, opts ...grpc.CallOption) (*SystemOperateLogMultipleResponse, error)
}

type systemOperateLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemOperateLogServiceClient(cc grpc.ClientConnInterface) SystemOperateLogServiceClient {
	return &systemOperateLogServiceClient{cc}
}

func (c *systemOperateLogServiceClient) SystemOperateLogCreate(ctx context.Context, in *SystemOperateLogCreateRequest, opts ...grpc.CallOption) (*SystemOperateLogCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemOperateLogCreateResponse)
	err := c.cc.Invoke(ctx, SystemOperateLogService_SystemOperateLogCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemOperateLogServiceClient) SystemOperateLogUpdate(ctx context.Context, in *SystemOperateLogUpdateRequest, opts ...grpc.CallOption) (*SystemOperateLogUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemOperateLogUpdateResponse)
	err := c.cc.Invoke(ctx, SystemOperateLogService_SystemOperateLogUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemOperateLogServiceClient) SystemOperateLogDelete(ctx context.Context, in *SystemOperateLogDeleteRequest, opts ...grpc.CallOption) (*SystemOperateLogDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemOperateLogDeleteResponse)
	err := c.cc.Invoke(ctx, SystemOperateLogService_SystemOperateLogDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemOperateLogServiceClient) SystemOperateLog(ctx context.Context, in *SystemOperateLogRequest, opts ...grpc.CallOption) (*SystemOperateLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemOperateLogResponse)
	err := c.cc.Invoke(ctx, SystemOperateLogService_SystemOperateLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemOperateLogServiceClient) SystemOperateLogRecover(ctx context.Context, in *SystemOperateLogRecoverRequest, opts ...grpc.CallOption) (*SystemOperateLogRecoverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemOperateLogRecoverResponse)
	err := c.cc.Invoke(ctx, SystemOperateLogService_SystemOperateLogRecover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemOperateLogServiceClient) SystemOperateLogDrop(ctx context.Context, in *SystemOperateLogDropRequest, opts ...grpc.CallOption) (*SystemOperateLogDropResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemOperateLogDropResponse)
	err := c.cc.Invoke(ctx, SystemOperateLogService_SystemOperateLogDrop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemOperateLogServiceClient) SystemOperateLogList(ctx context.Context, in *SystemOperateLogListRequest, opts ...grpc.CallOption) (*SystemOperateLogListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemOperateLogListResponse)
	err := c.cc.Invoke(ctx, SystemOperateLogService_SystemOperateLogList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemOperateLogServiceClient) SystemOperateLogListTotal(ctx context.Context, in *SystemOperateLogListTotalRequest, opts ...grpc.CallOption) (*SystemOperateLogTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemOperateLogTotalResponse)
	err := c.cc.Invoke(ctx, SystemOperateLogService_SystemOperateLogListTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemOperateLogServiceClient) SystemOperateLogMultipleDelete(ctx context.Context, in *SystemOperateLogMultipleRequest, opts ...grpc.CallOption) (*SystemOperateLogMultipleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemOperateLogMultipleResponse)
	err := c.cc.Invoke(ctx, SystemOperateLogService_SystemOperateLogMultipleDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemOperateLogServiceClient) SystemOperateLogMultipleRecover(ctx context.Context, in *SystemOperateLogMultipleRequest, opts ...grpc.CallOption) (*SystemOperateLogMultipleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemOperateLogMultipleResponse)
	err := c.cc.Invoke(ctx, SystemOperateLogService_SystemOperateLogMultipleRecover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemOperateLogServiceClient) SystemOperateLogMultipleDrop(ctx context.Context, in *SystemOperateLogMultipleRequest, opts ...grpc.CallOption) (*SystemOperateLogMultipleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemOperateLogMultipleResponse)
	err := c.cc.Invoke(ctx, SystemOperateLogService_SystemOperateLogMultipleDrop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemOperateLogServiceServer is the server API for SystemOperateLogService service.
// All implementations must embed UnimplementedSystemOperateLogServiceServer
// for forward compatibility
//
// SystemOperateLogService 服务
type SystemOperateLogServiceServer interface {
	SystemOperateLogCreate(context.Context, *SystemOperateLogCreateRequest) (*SystemOperateLogCreateResponse, error)
	SystemOperateLogUpdate(context.Context, *SystemOperateLogUpdateRequest) (*SystemOperateLogUpdateResponse, error)
	SystemOperateLogDelete(context.Context, *SystemOperateLogDeleteRequest) (*SystemOperateLogDeleteResponse, error)
	SystemOperateLog(context.Context, *SystemOperateLogRequest) (*SystemOperateLogResponse, error)
	SystemOperateLogRecover(context.Context, *SystemOperateLogRecoverRequest) (*SystemOperateLogRecoverResponse, error)
	SystemOperateLogDrop(context.Context, *SystemOperateLogDropRequest) (*SystemOperateLogDropResponse, error)
	SystemOperateLogList(context.Context, *SystemOperateLogListRequest) (*SystemOperateLogListResponse, error)
	SystemOperateLogListTotal(context.Context, *SystemOperateLogListTotalRequest) (*SystemOperateLogTotalResponse, error)
	SystemOperateLogMultipleDelete(context.Context, *SystemOperateLogMultipleRequest) (*SystemOperateLogMultipleResponse, error)
	SystemOperateLogMultipleRecover(context.Context, *SystemOperateLogMultipleRequest) (*SystemOperateLogMultipleResponse, error)
	SystemOperateLogMultipleDrop(context.Context, *SystemOperateLogMultipleRequest) (*SystemOperateLogMultipleResponse, error)
	mustEmbedUnimplementedSystemOperateLogServiceServer()
}

// UnimplementedSystemOperateLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemOperateLogServiceServer struct {
}

func (UnimplementedSystemOperateLogServiceServer) SystemOperateLogCreate(context.Context, *SystemOperateLogCreateRequest) (*SystemOperateLogCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemOperateLogCreate not implemented")
}
func (UnimplementedSystemOperateLogServiceServer) SystemOperateLogUpdate(context.Context, *SystemOperateLogUpdateRequest) (*SystemOperateLogUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemOperateLogUpdate not implemented")
}
func (UnimplementedSystemOperateLogServiceServer) SystemOperateLogDelete(context.Context, *SystemOperateLogDeleteRequest) (*SystemOperateLogDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemOperateLogDelete not implemented")
}
func (UnimplementedSystemOperateLogServiceServer) SystemOperateLog(context.Context, *SystemOperateLogRequest) (*SystemOperateLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemOperateLog not implemented")
}
func (UnimplementedSystemOperateLogServiceServer) SystemOperateLogRecover(context.Context, *SystemOperateLogRecoverRequest) (*SystemOperateLogRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemOperateLogRecover not implemented")
}
func (UnimplementedSystemOperateLogServiceServer) SystemOperateLogDrop(context.Context, *SystemOperateLogDropRequest) (*SystemOperateLogDropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemOperateLogDrop not implemented")
}
func (UnimplementedSystemOperateLogServiceServer) SystemOperateLogList(context.Context, *SystemOperateLogListRequest) (*SystemOperateLogListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemOperateLogList not implemented")
}
func (UnimplementedSystemOperateLogServiceServer) SystemOperateLogListTotal(context.Context, *SystemOperateLogListTotalRequest) (*SystemOperateLogTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemOperateLogListTotal not implemented")
}
func (UnimplementedSystemOperateLogServiceServer) SystemOperateLogMultipleDelete(context.Context, *SystemOperateLogMultipleRequest) (*SystemOperateLogMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemOperateLogMultipleDelete not implemented")
}
func (UnimplementedSystemOperateLogServiceServer) SystemOperateLogMultipleRecover(context.Context, *SystemOperateLogMultipleRequest) (*SystemOperateLogMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemOperateLogMultipleRecover not implemented")
}
func (UnimplementedSystemOperateLogServiceServer) SystemOperateLogMultipleDrop(context.Context, *SystemOperateLogMultipleRequest) (*SystemOperateLogMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemOperateLogMultipleDrop not implemented")
}
func (UnimplementedSystemOperateLogServiceServer) mustEmbedUnimplementedSystemOperateLogServiceServer() {
}

// UnsafeSystemOperateLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemOperateLogServiceServer will
// result in compilation errors.
type UnsafeSystemOperateLogServiceServer interface {
	mustEmbedUnimplementedSystemOperateLogServiceServer()
}

func RegisterSystemOperateLogServiceServer(s grpc.ServiceRegistrar, srv SystemOperateLogServiceServer) {
	s.RegisterService(&SystemOperateLogService_ServiceDesc, srv)
}

func _SystemOperateLogService_SystemOperateLogCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemOperateLogCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemOperateLogService_SystemOperateLogCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogCreate(ctx, req.(*SystemOperateLogCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemOperateLogService_SystemOperateLogUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemOperateLogUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemOperateLogService_SystemOperateLogUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogUpdate(ctx, req.(*SystemOperateLogUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemOperateLogService_SystemOperateLogDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemOperateLogDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemOperateLogService_SystemOperateLogDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogDelete(ctx, req.(*SystemOperateLogDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemOperateLogService_SystemOperateLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemOperateLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemOperateLogServiceServer).SystemOperateLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemOperateLogService_SystemOperateLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemOperateLogServiceServer).SystemOperateLog(ctx, req.(*SystemOperateLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemOperateLogService_SystemOperateLogRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemOperateLogRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemOperateLogService_SystemOperateLogRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogRecover(ctx, req.(*SystemOperateLogRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemOperateLogService_SystemOperateLogDrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemOperateLogDropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogDrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemOperateLogService_SystemOperateLogDrop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogDrop(ctx, req.(*SystemOperateLogDropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemOperateLogService_SystemOperateLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemOperateLogListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemOperateLogService_SystemOperateLogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogList(ctx, req.(*SystemOperateLogListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemOperateLogService_SystemOperateLogListTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemOperateLogListTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogListTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemOperateLogService_SystemOperateLogListTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogListTotal(ctx, req.(*SystemOperateLogListTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemOperateLogService_SystemOperateLogMultipleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemOperateLogMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogMultipleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemOperateLogService_SystemOperateLogMultipleDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogMultipleDelete(ctx, req.(*SystemOperateLogMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemOperateLogService_SystemOperateLogMultipleRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemOperateLogMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogMultipleRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemOperateLogService_SystemOperateLogMultipleRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogMultipleRecover(ctx, req.(*SystemOperateLogMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemOperateLogService_SystemOperateLogMultipleDrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemOperateLogMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogMultipleDrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemOperateLogService_SystemOperateLogMultipleDrop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemOperateLogServiceServer).SystemOperateLogMultipleDrop(ctx, req.(*SystemOperateLogMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemOperateLogService_ServiceDesc is the grpc.ServiceDesc for SystemOperateLogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemOperateLogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logger.SystemOperateLogService",
	HandlerType: (*SystemOperateLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemOperateLogCreate",
			Handler:    _SystemOperateLogService_SystemOperateLogCreate_Handler,
		},
		{
			MethodName: "SystemOperateLogUpdate",
			Handler:    _SystemOperateLogService_SystemOperateLogUpdate_Handler,
		},
		{
			MethodName: "SystemOperateLogDelete",
			Handler:    _SystemOperateLogService_SystemOperateLogDelete_Handler,
		},
		{
			MethodName: "SystemOperateLog",
			Handler:    _SystemOperateLogService_SystemOperateLog_Handler,
		},
		{
			MethodName: "SystemOperateLogRecover",
			Handler:    _SystemOperateLogService_SystemOperateLogRecover_Handler,
		},
		{
			MethodName: "SystemOperateLogDrop",
			Handler:    _SystemOperateLogService_SystemOperateLogDrop_Handler,
		},
		{
			MethodName: "SystemOperateLogList",
			Handler:    _SystemOperateLogService_SystemOperateLogList_Handler,
		},
		{
			MethodName: "SystemOperateLogListTotal",
			Handler:    _SystemOperateLogService_SystemOperateLogListTotal_Handler,
		},
		{
			MethodName: "SystemOperateLogMultipleDelete",
			Handler:    _SystemOperateLogService_SystemOperateLogMultipleDelete_Handler,
		},
		{
			MethodName: "SystemOperateLogMultipleRecover",
			Handler:    _SystemOperateLogService_SystemOperateLogMultipleRecover_Handler,
		},
		{
			MethodName: "SystemOperateLogMultipleDrop",
			Handler:    _SystemOperateLogService_SystemOperateLogMultipleDrop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_operate_log.proto",
}
