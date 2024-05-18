// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: order/pay_order_extension.proto

// pay_order_extension 支付订单拓展

package order

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
	PayOrderExtensionService_PayOrderExtensionCreate_FullMethodName  = "/order.PayOrderExtensionService/PayOrderExtensionCreate"
	PayOrderExtensionService_PayOrderExtensionUpdate_FullMethodName  = "/order.PayOrderExtensionService/PayOrderExtensionUpdate"
	PayOrderExtensionService_PayOrderExtensionDelete_FullMethodName  = "/order.PayOrderExtensionService/PayOrderExtensionDelete"
	PayOrderExtensionService_PayOrderExtension_FullMethodName        = "/order.PayOrderExtensionService/PayOrderExtension"
	PayOrderExtensionService_PayOrderExtensionRecover_FullMethodName = "/order.PayOrderExtensionService/PayOrderExtensionRecover"
	PayOrderExtensionService_PayOrderExtensionItem_FullMethodName    = "/order.PayOrderExtensionService/PayOrderExtensionItem"
	PayOrderExtensionService_PayOrderExtensionList_FullMethodName    = "/order.PayOrderExtensionService/PayOrderExtensionList"
)

// PayOrderExtensionServiceClient is the client API for PayOrderExtensionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayOrderExtensionServiceClient interface {
	PayOrderExtensionCreate(ctx context.Context, in *PayOrderExtensionCreateRequest, opts ...grpc.CallOption) (*PayOrderExtensionCreateResponse, error)
	PayOrderExtensionUpdate(ctx context.Context, in *PayOrderExtensionUpdateRequest, opts ...grpc.CallOption) (*PayOrderExtensionUpdateResponse, error)
	PayOrderExtensionDelete(ctx context.Context, in *PayOrderExtensionDeleteRequest, opts ...grpc.CallOption) (*PayOrderExtensionDeleteResponse, error)
	PayOrderExtension(ctx context.Context, in *PayOrderExtensionRequest, opts ...grpc.CallOption) (*PayOrderExtensionResponse, error)
	PayOrderExtensionRecover(ctx context.Context, in *PayOrderExtensionRecoverRequest, opts ...grpc.CallOption) (*PayOrderExtensionRecoverResponse, error)
	PayOrderExtensionItem(ctx context.Context, in *PayOrderExtensionItemRequest, opts ...grpc.CallOption) (*PayOrderExtensionItemResponse, error)
	PayOrderExtensionList(ctx context.Context, in *PayOrderExtensionListRequest, opts ...grpc.CallOption) (*PayOrderExtensionListResponse, error)
}

type payOrderExtensionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPayOrderExtensionServiceClient(cc grpc.ClientConnInterface) PayOrderExtensionServiceClient {
	return &payOrderExtensionServiceClient{cc}
}

func (c *payOrderExtensionServiceClient) PayOrderExtensionCreate(ctx context.Context, in *PayOrderExtensionCreateRequest, opts ...grpc.CallOption) (*PayOrderExtensionCreateResponse, error) {
	out := new(PayOrderExtensionCreateResponse)
	err := c.cc.Invoke(ctx, PayOrderExtensionService_PayOrderExtensionCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payOrderExtensionServiceClient) PayOrderExtensionUpdate(ctx context.Context, in *PayOrderExtensionUpdateRequest, opts ...grpc.CallOption) (*PayOrderExtensionUpdateResponse, error) {
	out := new(PayOrderExtensionUpdateResponse)
	err := c.cc.Invoke(ctx, PayOrderExtensionService_PayOrderExtensionUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payOrderExtensionServiceClient) PayOrderExtensionDelete(ctx context.Context, in *PayOrderExtensionDeleteRequest, opts ...grpc.CallOption) (*PayOrderExtensionDeleteResponse, error) {
	out := new(PayOrderExtensionDeleteResponse)
	err := c.cc.Invoke(ctx, PayOrderExtensionService_PayOrderExtensionDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payOrderExtensionServiceClient) PayOrderExtension(ctx context.Context, in *PayOrderExtensionRequest, opts ...grpc.CallOption) (*PayOrderExtensionResponse, error) {
	out := new(PayOrderExtensionResponse)
	err := c.cc.Invoke(ctx, PayOrderExtensionService_PayOrderExtension_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payOrderExtensionServiceClient) PayOrderExtensionRecover(ctx context.Context, in *PayOrderExtensionRecoverRequest, opts ...grpc.CallOption) (*PayOrderExtensionRecoverResponse, error) {
	out := new(PayOrderExtensionRecoverResponse)
	err := c.cc.Invoke(ctx, PayOrderExtensionService_PayOrderExtensionRecover_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payOrderExtensionServiceClient) PayOrderExtensionItem(ctx context.Context, in *PayOrderExtensionItemRequest, opts ...grpc.CallOption) (*PayOrderExtensionItemResponse, error) {
	out := new(PayOrderExtensionItemResponse)
	err := c.cc.Invoke(ctx, PayOrderExtensionService_PayOrderExtensionItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payOrderExtensionServiceClient) PayOrderExtensionList(ctx context.Context, in *PayOrderExtensionListRequest, opts ...grpc.CallOption) (*PayOrderExtensionListResponse, error) {
	out := new(PayOrderExtensionListResponse)
	err := c.cc.Invoke(ctx, PayOrderExtensionService_PayOrderExtensionList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayOrderExtensionServiceServer is the server API for PayOrderExtensionService service.
// All implementations must embed UnimplementedPayOrderExtensionServiceServer
// for forward compatibility
type PayOrderExtensionServiceServer interface {
	PayOrderExtensionCreate(context.Context, *PayOrderExtensionCreateRequest) (*PayOrderExtensionCreateResponse, error)
	PayOrderExtensionUpdate(context.Context, *PayOrderExtensionUpdateRequest) (*PayOrderExtensionUpdateResponse, error)
	PayOrderExtensionDelete(context.Context, *PayOrderExtensionDeleteRequest) (*PayOrderExtensionDeleteResponse, error)
	PayOrderExtension(context.Context, *PayOrderExtensionRequest) (*PayOrderExtensionResponse, error)
	PayOrderExtensionRecover(context.Context, *PayOrderExtensionRecoverRequest) (*PayOrderExtensionRecoverResponse, error)
	PayOrderExtensionItem(context.Context, *PayOrderExtensionItemRequest) (*PayOrderExtensionItemResponse, error)
	PayOrderExtensionList(context.Context, *PayOrderExtensionListRequest) (*PayOrderExtensionListResponse, error)
	mustEmbedUnimplementedPayOrderExtensionServiceServer()
}

// UnimplementedPayOrderExtensionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPayOrderExtensionServiceServer struct {
}

func (UnimplementedPayOrderExtensionServiceServer) PayOrderExtensionCreate(context.Context, *PayOrderExtensionCreateRequest) (*PayOrderExtensionCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayOrderExtensionCreate not implemented")
}
func (UnimplementedPayOrderExtensionServiceServer) PayOrderExtensionUpdate(context.Context, *PayOrderExtensionUpdateRequest) (*PayOrderExtensionUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayOrderExtensionUpdate not implemented")
}
func (UnimplementedPayOrderExtensionServiceServer) PayOrderExtensionDelete(context.Context, *PayOrderExtensionDeleteRequest) (*PayOrderExtensionDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayOrderExtensionDelete not implemented")
}
func (UnimplementedPayOrderExtensionServiceServer) PayOrderExtension(context.Context, *PayOrderExtensionRequest) (*PayOrderExtensionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayOrderExtension not implemented")
}
func (UnimplementedPayOrderExtensionServiceServer) PayOrderExtensionRecover(context.Context, *PayOrderExtensionRecoverRequest) (*PayOrderExtensionRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayOrderExtensionRecover not implemented")
}
func (UnimplementedPayOrderExtensionServiceServer) PayOrderExtensionItem(context.Context, *PayOrderExtensionItemRequest) (*PayOrderExtensionItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayOrderExtensionItem not implemented")
}
func (UnimplementedPayOrderExtensionServiceServer) PayOrderExtensionList(context.Context, *PayOrderExtensionListRequest) (*PayOrderExtensionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayOrderExtensionList not implemented")
}
func (UnimplementedPayOrderExtensionServiceServer) mustEmbedUnimplementedPayOrderExtensionServiceServer() {
}

// UnsafePayOrderExtensionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayOrderExtensionServiceServer will
// result in compilation errors.
type UnsafePayOrderExtensionServiceServer interface {
	mustEmbedUnimplementedPayOrderExtensionServiceServer()
}

func RegisterPayOrderExtensionServiceServer(s grpc.ServiceRegistrar, srv PayOrderExtensionServiceServer) {
	s.RegisterService(&PayOrderExtensionService_ServiceDesc, srv)
}

func _PayOrderExtensionService_PayOrderExtensionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayOrderExtensionCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtensionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayOrderExtensionService_PayOrderExtensionCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtensionCreate(ctx, req.(*PayOrderExtensionCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayOrderExtensionService_PayOrderExtensionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayOrderExtensionUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtensionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayOrderExtensionService_PayOrderExtensionUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtensionUpdate(ctx, req.(*PayOrderExtensionUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayOrderExtensionService_PayOrderExtensionDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayOrderExtensionDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtensionDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayOrderExtensionService_PayOrderExtensionDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtensionDelete(ctx, req.(*PayOrderExtensionDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayOrderExtensionService_PayOrderExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayOrderExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayOrderExtensionService_PayOrderExtension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtension(ctx, req.(*PayOrderExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayOrderExtensionService_PayOrderExtensionRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayOrderExtensionRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtensionRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayOrderExtensionService_PayOrderExtensionRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtensionRecover(ctx, req.(*PayOrderExtensionRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayOrderExtensionService_PayOrderExtensionItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayOrderExtensionItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtensionItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayOrderExtensionService_PayOrderExtensionItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtensionItem(ctx, req.(*PayOrderExtensionItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayOrderExtensionService_PayOrderExtensionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayOrderExtensionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtensionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayOrderExtensionService_PayOrderExtensionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayOrderExtensionServiceServer).PayOrderExtensionList(ctx, req.(*PayOrderExtensionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PayOrderExtensionService_ServiceDesc is the grpc.ServiceDesc for PayOrderExtensionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayOrderExtensionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.PayOrderExtensionService",
	HandlerType: (*PayOrderExtensionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PayOrderExtensionCreate",
			Handler:    _PayOrderExtensionService_PayOrderExtensionCreate_Handler,
		},
		{
			MethodName: "PayOrderExtensionUpdate",
			Handler:    _PayOrderExtensionService_PayOrderExtensionUpdate_Handler,
		},
		{
			MethodName: "PayOrderExtensionDelete",
			Handler:    _PayOrderExtensionService_PayOrderExtensionDelete_Handler,
		},
		{
			MethodName: "PayOrderExtension",
			Handler:    _PayOrderExtensionService_PayOrderExtension_Handler,
		},
		{
			MethodName: "PayOrderExtensionRecover",
			Handler:    _PayOrderExtensionService_PayOrderExtensionRecover_Handler,
		},
		{
			MethodName: "PayOrderExtensionItem",
			Handler:    _PayOrderExtensionService_PayOrderExtensionItem_Handler,
		},
		{
			MethodName: "PayOrderExtensionList",
			Handler:    _PayOrderExtensionService_PayOrderExtensionList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/pay_order_extension.proto",
}