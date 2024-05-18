// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: notify/system_notify_message.proto

// system_notify_message 站内信消息表

package notify

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
	SystemNotifyMessageService_SystemNotifyMessageCreate_FullMethodName    = "/notify.SystemNotifyMessageService/SystemNotifyMessageCreate"
	SystemNotifyMessageService_SystemNotifyMessageUpdate_FullMethodName    = "/notify.SystemNotifyMessageService/SystemNotifyMessageUpdate"
	SystemNotifyMessageService_SystemNotifyMessageDelete_FullMethodName    = "/notify.SystemNotifyMessageService/SystemNotifyMessageDelete"
	SystemNotifyMessageService_SystemNotifyMessage_FullMethodName          = "/notify.SystemNotifyMessageService/SystemNotifyMessage"
	SystemNotifyMessageService_SystemNotifyMessageRecover_FullMethodName   = "/notify.SystemNotifyMessageService/SystemNotifyMessageRecover"
	SystemNotifyMessageService_SystemNotifyMessageList_FullMethodName      = "/notify.SystemNotifyMessageService/SystemNotifyMessageList"
	SystemNotifyMessageService_SystemNotifyMessageListTotal_FullMethodName = "/notify.SystemNotifyMessageService/SystemNotifyMessageListTotal"
)

// SystemNotifyMessageServiceClient is the client API for SystemNotifyMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemNotifyMessageServiceClient interface {
	SystemNotifyMessageCreate(ctx context.Context, in *SystemNotifyMessageCreateRequest, opts ...grpc.CallOption) (*SystemNotifyMessageCreateResponse, error)
	SystemNotifyMessageUpdate(ctx context.Context, in *SystemNotifyMessageUpdateRequest, opts ...grpc.CallOption) (*SystemNotifyMessageUpdateResponse, error)
	SystemNotifyMessageDelete(ctx context.Context, in *SystemNotifyMessageDeleteRequest, opts ...grpc.CallOption) (*SystemNotifyMessageDeleteResponse, error)
	SystemNotifyMessage(ctx context.Context, in *SystemNotifyMessageRequest, opts ...grpc.CallOption) (*SystemNotifyMessageResponse, error)
	SystemNotifyMessageRecover(ctx context.Context, in *SystemNotifyMessageRecoverRequest, opts ...grpc.CallOption) (*SystemNotifyMessageRecoverResponse, error)
	SystemNotifyMessageList(ctx context.Context, in *SystemNotifyMessageListRequest, opts ...grpc.CallOption) (*SystemNotifyMessageListResponse, error)
	SystemNotifyMessageListTotal(ctx context.Context, in *SystemNotifyMessageListTotalRequest, opts ...grpc.CallOption) (*SystemNotifyMessageTotalResponse, error)
}

type systemNotifyMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemNotifyMessageServiceClient(cc grpc.ClientConnInterface) SystemNotifyMessageServiceClient {
	return &systemNotifyMessageServiceClient{cc}
}

func (c *systemNotifyMessageServiceClient) SystemNotifyMessageCreate(ctx context.Context, in *SystemNotifyMessageCreateRequest, opts ...grpc.CallOption) (*SystemNotifyMessageCreateResponse, error) {
	out := new(SystemNotifyMessageCreateResponse)
	err := c.cc.Invoke(ctx, SystemNotifyMessageService_SystemNotifyMessageCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemNotifyMessageServiceClient) SystemNotifyMessageUpdate(ctx context.Context, in *SystemNotifyMessageUpdateRequest, opts ...grpc.CallOption) (*SystemNotifyMessageUpdateResponse, error) {
	out := new(SystemNotifyMessageUpdateResponse)
	err := c.cc.Invoke(ctx, SystemNotifyMessageService_SystemNotifyMessageUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemNotifyMessageServiceClient) SystemNotifyMessageDelete(ctx context.Context, in *SystemNotifyMessageDeleteRequest, opts ...grpc.CallOption) (*SystemNotifyMessageDeleteResponse, error) {
	out := new(SystemNotifyMessageDeleteResponse)
	err := c.cc.Invoke(ctx, SystemNotifyMessageService_SystemNotifyMessageDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemNotifyMessageServiceClient) SystemNotifyMessage(ctx context.Context, in *SystemNotifyMessageRequest, opts ...grpc.CallOption) (*SystemNotifyMessageResponse, error) {
	out := new(SystemNotifyMessageResponse)
	err := c.cc.Invoke(ctx, SystemNotifyMessageService_SystemNotifyMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemNotifyMessageServiceClient) SystemNotifyMessageRecover(ctx context.Context, in *SystemNotifyMessageRecoverRequest, opts ...grpc.CallOption) (*SystemNotifyMessageRecoverResponse, error) {
	out := new(SystemNotifyMessageRecoverResponse)
	err := c.cc.Invoke(ctx, SystemNotifyMessageService_SystemNotifyMessageRecover_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemNotifyMessageServiceClient) SystemNotifyMessageList(ctx context.Context, in *SystemNotifyMessageListRequest, opts ...grpc.CallOption) (*SystemNotifyMessageListResponse, error) {
	out := new(SystemNotifyMessageListResponse)
	err := c.cc.Invoke(ctx, SystemNotifyMessageService_SystemNotifyMessageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemNotifyMessageServiceClient) SystemNotifyMessageListTotal(ctx context.Context, in *SystemNotifyMessageListTotalRequest, opts ...grpc.CallOption) (*SystemNotifyMessageTotalResponse, error) {
	out := new(SystemNotifyMessageTotalResponse)
	err := c.cc.Invoke(ctx, SystemNotifyMessageService_SystemNotifyMessageListTotal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemNotifyMessageServiceServer is the server API for SystemNotifyMessageService service.
// All implementations must embed UnimplementedSystemNotifyMessageServiceServer
// for forward compatibility
type SystemNotifyMessageServiceServer interface {
	SystemNotifyMessageCreate(context.Context, *SystemNotifyMessageCreateRequest) (*SystemNotifyMessageCreateResponse, error)
	SystemNotifyMessageUpdate(context.Context, *SystemNotifyMessageUpdateRequest) (*SystemNotifyMessageUpdateResponse, error)
	SystemNotifyMessageDelete(context.Context, *SystemNotifyMessageDeleteRequest) (*SystemNotifyMessageDeleteResponse, error)
	SystemNotifyMessage(context.Context, *SystemNotifyMessageRequest) (*SystemNotifyMessageResponse, error)
	SystemNotifyMessageRecover(context.Context, *SystemNotifyMessageRecoverRequest) (*SystemNotifyMessageRecoverResponse, error)
	SystemNotifyMessageList(context.Context, *SystemNotifyMessageListRequest) (*SystemNotifyMessageListResponse, error)
	SystemNotifyMessageListTotal(context.Context, *SystemNotifyMessageListTotalRequest) (*SystemNotifyMessageTotalResponse, error)
	mustEmbedUnimplementedSystemNotifyMessageServiceServer()
}

// UnimplementedSystemNotifyMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemNotifyMessageServiceServer struct {
}

func (UnimplementedSystemNotifyMessageServiceServer) SystemNotifyMessageCreate(context.Context, *SystemNotifyMessageCreateRequest) (*SystemNotifyMessageCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemNotifyMessageCreate not implemented")
}
func (UnimplementedSystemNotifyMessageServiceServer) SystemNotifyMessageUpdate(context.Context, *SystemNotifyMessageUpdateRequest) (*SystemNotifyMessageUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemNotifyMessageUpdate not implemented")
}
func (UnimplementedSystemNotifyMessageServiceServer) SystemNotifyMessageDelete(context.Context, *SystemNotifyMessageDeleteRequest) (*SystemNotifyMessageDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemNotifyMessageDelete not implemented")
}
func (UnimplementedSystemNotifyMessageServiceServer) SystemNotifyMessage(context.Context, *SystemNotifyMessageRequest) (*SystemNotifyMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemNotifyMessage not implemented")
}
func (UnimplementedSystemNotifyMessageServiceServer) SystemNotifyMessageRecover(context.Context, *SystemNotifyMessageRecoverRequest) (*SystemNotifyMessageRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemNotifyMessageRecover not implemented")
}
func (UnimplementedSystemNotifyMessageServiceServer) SystemNotifyMessageList(context.Context, *SystemNotifyMessageListRequest) (*SystemNotifyMessageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemNotifyMessageList not implemented")
}
func (UnimplementedSystemNotifyMessageServiceServer) SystemNotifyMessageListTotal(context.Context, *SystemNotifyMessageListTotalRequest) (*SystemNotifyMessageTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemNotifyMessageListTotal not implemented")
}
func (UnimplementedSystemNotifyMessageServiceServer) mustEmbedUnimplementedSystemNotifyMessageServiceServer() {
}

// UnsafeSystemNotifyMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemNotifyMessageServiceServer will
// result in compilation errors.
type UnsafeSystemNotifyMessageServiceServer interface {
	mustEmbedUnimplementedSystemNotifyMessageServiceServer()
}

func RegisterSystemNotifyMessageServiceServer(s grpc.ServiceRegistrar, srv SystemNotifyMessageServiceServer) {
	s.RegisterService(&SystemNotifyMessageService_ServiceDesc, srv)
}

func _SystemNotifyMessageService_SystemNotifyMessageCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemNotifyMessageCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessageCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemNotifyMessageService_SystemNotifyMessageCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessageCreate(ctx, req.(*SystemNotifyMessageCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemNotifyMessageService_SystemNotifyMessageUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemNotifyMessageUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessageUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemNotifyMessageService_SystemNotifyMessageUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessageUpdate(ctx, req.(*SystemNotifyMessageUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemNotifyMessageService_SystemNotifyMessageDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemNotifyMessageDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessageDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemNotifyMessageService_SystemNotifyMessageDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessageDelete(ctx, req.(*SystemNotifyMessageDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemNotifyMessageService_SystemNotifyMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemNotifyMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemNotifyMessageService_SystemNotifyMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessage(ctx, req.(*SystemNotifyMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemNotifyMessageService_SystemNotifyMessageRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemNotifyMessageRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessageRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemNotifyMessageService_SystemNotifyMessageRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessageRecover(ctx, req.(*SystemNotifyMessageRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemNotifyMessageService_SystemNotifyMessageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemNotifyMessageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemNotifyMessageService_SystemNotifyMessageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessageList(ctx, req.(*SystemNotifyMessageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemNotifyMessageService_SystemNotifyMessageListTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemNotifyMessageListTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessageListTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemNotifyMessageService_SystemNotifyMessageListTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemNotifyMessageServiceServer).SystemNotifyMessageListTotal(ctx, req.(*SystemNotifyMessageListTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemNotifyMessageService_ServiceDesc is the grpc.ServiceDesc for SystemNotifyMessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemNotifyMessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notify.SystemNotifyMessageService",
	HandlerType: (*SystemNotifyMessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemNotifyMessageCreate",
			Handler:    _SystemNotifyMessageService_SystemNotifyMessageCreate_Handler,
		},
		{
			MethodName: "SystemNotifyMessageUpdate",
			Handler:    _SystemNotifyMessageService_SystemNotifyMessageUpdate_Handler,
		},
		{
			MethodName: "SystemNotifyMessageDelete",
			Handler:    _SystemNotifyMessageService_SystemNotifyMessageDelete_Handler,
		},
		{
			MethodName: "SystemNotifyMessage",
			Handler:    _SystemNotifyMessageService_SystemNotifyMessage_Handler,
		},
		{
			MethodName: "SystemNotifyMessageRecover",
			Handler:    _SystemNotifyMessageService_SystemNotifyMessageRecover_Handler,
		},
		{
			MethodName: "SystemNotifyMessageList",
			Handler:    _SystemNotifyMessageService_SystemNotifyMessageList_Handler,
		},
		{
			MethodName: "SystemNotifyMessageListTotal",
			Handler:    _SystemNotifyMessageService_SystemNotifyMessageListTotal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notify/system_notify_message.proto",
}