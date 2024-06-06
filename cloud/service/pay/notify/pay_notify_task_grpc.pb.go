// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: notify/pay_notify_task.proto

// pay_notify_task 商户支付-任务通知

package notify

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
	PayNotifyTaskService_PayNotifyTaskCreate_FullMethodName    = "/notify.PayNotifyTaskService/PayNotifyTaskCreate"
	PayNotifyTaskService_PayNotifyTaskUpdate_FullMethodName    = "/notify.PayNotifyTaskService/PayNotifyTaskUpdate"
	PayNotifyTaskService_PayNotifyTaskDelete_FullMethodName    = "/notify.PayNotifyTaskService/PayNotifyTaskDelete"
	PayNotifyTaskService_PayNotifyTask_FullMethodName          = "/notify.PayNotifyTaskService/PayNotifyTask"
	PayNotifyTaskService_PayNotifyTaskRecover_FullMethodName   = "/notify.PayNotifyTaskService/PayNotifyTaskRecover"
	PayNotifyTaskService_PayNotifyTaskList_FullMethodName      = "/notify.PayNotifyTaskService/PayNotifyTaskList"
	PayNotifyTaskService_PayNotifyTaskListTotal_FullMethodName = "/notify.PayNotifyTaskService/PayNotifyTaskListTotal"
)

// PayNotifyTaskServiceClient is the client API for PayNotifyTaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// PayNotifyTaskService 服务
type PayNotifyTaskServiceClient interface {
	PayNotifyTaskCreate(ctx context.Context, in *PayNotifyTaskCreateRequest, opts ...grpc.CallOption) (*PayNotifyTaskCreateResponse, error)
	PayNotifyTaskUpdate(ctx context.Context, in *PayNotifyTaskUpdateRequest, opts ...grpc.CallOption) (*PayNotifyTaskUpdateResponse, error)
	PayNotifyTaskDelete(ctx context.Context, in *PayNotifyTaskDeleteRequest, opts ...grpc.CallOption) (*PayNotifyTaskDeleteResponse, error)
	PayNotifyTask(ctx context.Context, in *PayNotifyTaskRequest, opts ...grpc.CallOption) (*PayNotifyTaskResponse, error)
	PayNotifyTaskRecover(ctx context.Context, in *PayNotifyTaskRecoverRequest, opts ...grpc.CallOption) (*PayNotifyTaskRecoverResponse, error)
	PayNotifyTaskList(ctx context.Context, in *PayNotifyTaskListRequest, opts ...grpc.CallOption) (*PayNotifyTaskListResponse, error)
	PayNotifyTaskListTotal(ctx context.Context, in *PayNotifyTaskListTotalRequest, opts ...grpc.CallOption) (*PayNotifyTaskTotalResponse, error)
}

type payNotifyTaskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPayNotifyTaskServiceClient(cc grpc.ClientConnInterface) PayNotifyTaskServiceClient {
	return &payNotifyTaskServiceClient{cc}
}

func (c *payNotifyTaskServiceClient) PayNotifyTaskCreate(ctx context.Context, in *PayNotifyTaskCreateRequest, opts ...grpc.CallOption) (*PayNotifyTaskCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayNotifyTaskCreateResponse)
	err := c.cc.Invoke(ctx, PayNotifyTaskService_PayNotifyTaskCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payNotifyTaskServiceClient) PayNotifyTaskUpdate(ctx context.Context, in *PayNotifyTaskUpdateRequest, opts ...grpc.CallOption) (*PayNotifyTaskUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayNotifyTaskUpdateResponse)
	err := c.cc.Invoke(ctx, PayNotifyTaskService_PayNotifyTaskUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payNotifyTaskServiceClient) PayNotifyTaskDelete(ctx context.Context, in *PayNotifyTaskDeleteRequest, opts ...grpc.CallOption) (*PayNotifyTaskDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayNotifyTaskDeleteResponse)
	err := c.cc.Invoke(ctx, PayNotifyTaskService_PayNotifyTaskDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payNotifyTaskServiceClient) PayNotifyTask(ctx context.Context, in *PayNotifyTaskRequest, opts ...grpc.CallOption) (*PayNotifyTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayNotifyTaskResponse)
	err := c.cc.Invoke(ctx, PayNotifyTaskService_PayNotifyTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payNotifyTaskServiceClient) PayNotifyTaskRecover(ctx context.Context, in *PayNotifyTaskRecoverRequest, opts ...grpc.CallOption) (*PayNotifyTaskRecoverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayNotifyTaskRecoverResponse)
	err := c.cc.Invoke(ctx, PayNotifyTaskService_PayNotifyTaskRecover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payNotifyTaskServiceClient) PayNotifyTaskList(ctx context.Context, in *PayNotifyTaskListRequest, opts ...grpc.CallOption) (*PayNotifyTaskListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayNotifyTaskListResponse)
	err := c.cc.Invoke(ctx, PayNotifyTaskService_PayNotifyTaskList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payNotifyTaskServiceClient) PayNotifyTaskListTotal(ctx context.Context, in *PayNotifyTaskListTotalRequest, opts ...grpc.CallOption) (*PayNotifyTaskTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayNotifyTaskTotalResponse)
	err := c.cc.Invoke(ctx, PayNotifyTaskService_PayNotifyTaskListTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayNotifyTaskServiceServer is the server API for PayNotifyTaskService service.
// All implementations must embed UnimplementedPayNotifyTaskServiceServer
// for forward compatibility
//
// PayNotifyTaskService 服务
type PayNotifyTaskServiceServer interface {
	PayNotifyTaskCreate(context.Context, *PayNotifyTaskCreateRequest) (*PayNotifyTaskCreateResponse, error)
	PayNotifyTaskUpdate(context.Context, *PayNotifyTaskUpdateRequest) (*PayNotifyTaskUpdateResponse, error)
	PayNotifyTaskDelete(context.Context, *PayNotifyTaskDeleteRequest) (*PayNotifyTaskDeleteResponse, error)
	PayNotifyTask(context.Context, *PayNotifyTaskRequest) (*PayNotifyTaskResponse, error)
	PayNotifyTaskRecover(context.Context, *PayNotifyTaskRecoverRequest) (*PayNotifyTaskRecoverResponse, error)
	PayNotifyTaskList(context.Context, *PayNotifyTaskListRequest) (*PayNotifyTaskListResponse, error)
	PayNotifyTaskListTotal(context.Context, *PayNotifyTaskListTotalRequest) (*PayNotifyTaskTotalResponse, error)
	mustEmbedUnimplementedPayNotifyTaskServiceServer()
}

// UnimplementedPayNotifyTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPayNotifyTaskServiceServer struct {
}

func (UnimplementedPayNotifyTaskServiceServer) PayNotifyTaskCreate(context.Context, *PayNotifyTaskCreateRequest) (*PayNotifyTaskCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayNotifyTaskCreate not implemented")
}
func (UnimplementedPayNotifyTaskServiceServer) PayNotifyTaskUpdate(context.Context, *PayNotifyTaskUpdateRequest) (*PayNotifyTaskUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayNotifyTaskUpdate not implemented")
}
func (UnimplementedPayNotifyTaskServiceServer) PayNotifyTaskDelete(context.Context, *PayNotifyTaskDeleteRequest) (*PayNotifyTaskDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayNotifyTaskDelete not implemented")
}
func (UnimplementedPayNotifyTaskServiceServer) PayNotifyTask(context.Context, *PayNotifyTaskRequest) (*PayNotifyTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayNotifyTask not implemented")
}
func (UnimplementedPayNotifyTaskServiceServer) PayNotifyTaskRecover(context.Context, *PayNotifyTaskRecoverRequest) (*PayNotifyTaskRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayNotifyTaskRecover not implemented")
}
func (UnimplementedPayNotifyTaskServiceServer) PayNotifyTaskList(context.Context, *PayNotifyTaskListRequest) (*PayNotifyTaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayNotifyTaskList not implemented")
}
func (UnimplementedPayNotifyTaskServiceServer) PayNotifyTaskListTotal(context.Context, *PayNotifyTaskListTotalRequest) (*PayNotifyTaskTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayNotifyTaskListTotal not implemented")
}
func (UnimplementedPayNotifyTaskServiceServer) mustEmbedUnimplementedPayNotifyTaskServiceServer() {}

// UnsafePayNotifyTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayNotifyTaskServiceServer will
// result in compilation errors.
type UnsafePayNotifyTaskServiceServer interface {
	mustEmbedUnimplementedPayNotifyTaskServiceServer()
}

func RegisterPayNotifyTaskServiceServer(s grpc.ServiceRegistrar, srv PayNotifyTaskServiceServer) {
	s.RegisterService(&PayNotifyTaskService_ServiceDesc, srv)
}

func _PayNotifyTaskService_PayNotifyTaskCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayNotifyTaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTaskCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayNotifyTaskService_PayNotifyTaskCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTaskCreate(ctx, req.(*PayNotifyTaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayNotifyTaskService_PayNotifyTaskUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayNotifyTaskUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTaskUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayNotifyTaskService_PayNotifyTaskUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTaskUpdate(ctx, req.(*PayNotifyTaskUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayNotifyTaskService_PayNotifyTaskDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayNotifyTaskDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTaskDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayNotifyTaskService_PayNotifyTaskDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTaskDelete(ctx, req.(*PayNotifyTaskDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayNotifyTaskService_PayNotifyTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayNotifyTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayNotifyTaskService_PayNotifyTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTask(ctx, req.(*PayNotifyTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayNotifyTaskService_PayNotifyTaskRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayNotifyTaskRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTaskRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayNotifyTaskService_PayNotifyTaskRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTaskRecover(ctx, req.(*PayNotifyTaskRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayNotifyTaskService_PayNotifyTaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayNotifyTaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayNotifyTaskService_PayNotifyTaskList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTaskList(ctx, req.(*PayNotifyTaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayNotifyTaskService_PayNotifyTaskListTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayNotifyTaskListTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTaskListTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayNotifyTaskService_PayNotifyTaskListTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayNotifyTaskServiceServer).PayNotifyTaskListTotal(ctx, req.(*PayNotifyTaskListTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PayNotifyTaskService_ServiceDesc is the grpc.ServiceDesc for PayNotifyTaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayNotifyTaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notify.PayNotifyTaskService",
	HandlerType: (*PayNotifyTaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PayNotifyTaskCreate",
			Handler:    _PayNotifyTaskService_PayNotifyTaskCreate_Handler,
		},
		{
			MethodName: "PayNotifyTaskUpdate",
			Handler:    _PayNotifyTaskService_PayNotifyTaskUpdate_Handler,
		},
		{
			MethodName: "PayNotifyTaskDelete",
			Handler:    _PayNotifyTaskService_PayNotifyTaskDelete_Handler,
		},
		{
			MethodName: "PayNotifyTask",
			Handler:    _PayNotifyTaskService_PayNotifyTask_Handler,
		},
		{
			MethodName: "PayNotifyTaskRecover",
			Handler:    _PayNotifyTaskService_PayNotifyTaskRecover_Handler,
		},
		{
			MethodName: "PayNotifyTaskList",
			Handler:    _PayNotifyTaskService_PayNotifyTaskList_Handler,
		},
		{
			MethodName: "PayNotifyTaskListTotal",
			Handler:    _PayNotifyTaskService_PayNotifyTaskListTotal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notify/pay_notify_task.proto",
}
