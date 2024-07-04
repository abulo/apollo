// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: pay_channel.proto

// pay_channel 支付渠道

package channel

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
	PayChannelService_PayChannelCreate_FullMethodName  = "/channel.PayChannelService/PayChannelCreate"
	PayChannelService_PayChannelUpdate_FullMethodName  = "/channel.PayChannelService/PayChannelUpdate"
	PayChannelService_PayChannelDelete_FullMethodName  = "/channel.PayChannelService/PayChannelDelete"
	PayChannelService_PayChannel_FullMethodName        = "/channel.PayChannelService/PayChannel"
	PayChannelService_PayChannelRecover_FullMethodName = "/channel.PayChannelService/PayChannelRecover"
	PayChannelService_PayChannelCode_FullMethodName    = "/channel.PayChannelService/PayChannelCode"
	PayChannelService_PayChannelList_FullMethodName    = "/channel.PayChannelService/PayChannelList"
)

// PayChannelServiceClient is the client API for PayChannelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// PayChannelService 服务
type PayChannelServiceClient interface {
	PayChannelCreate(ctx context.Context, in *PayChannelCreateRequest, opts ...grpc.CallOption) (*PayChannelCreateResponse, error)
	PayChannelUpdate(ctx context.Context, in *PayChannelUpdateRequest, opts ...grpc.CallOption) (*PayChannelUpdateResponse, error)
	PayChannelDelete(ctx context.Context, in *PayChannelDeleteRequest, opts ...grpc.CallOption) (*PayChannelDeleteResponse, error)
	PayChannel(ctx context.Context, in *PayChannelRequest, opts ...grpc.CallOption) (*PayChannelResponse, error)
	PayChannelRecover(ctx context.Context, in *PayChannelRecoverRequest, opts ...grpc.CallOption) (*PayChannelRecoverResponse, error)
	PayChannelCode(ctx context.Context, in *PayChannelCodeRequest, opts ...grpc.CallOption) (*PayChannelCodeResponse, error)
	PayChannelList(ctx context.Context, in *PayChannelListRequest, opts ...grpc.CallOption) (*PayChannelListResponse, error)
}

type payChannelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPayChannelServiceClient(cc grpc.ClientConnInterface) PayChannelServiceClient {
	return &payChannelServiceClient{cc}
}

func (c *payChannelServiceClient) PayChannelCreate(ctx context.Context, in *PayChannelCreateRequest, opts ...grpc.CallOption) (*PayChannelCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayChannelCreateResponse)
	err := c.cc.Invoke(ctx, PayChannelService_PayChannelCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payChannelServiceClient) PayChannelUpdate(ctx context.Context, in *PayChannelUpdateRequest, opts ...grpc.CallOption) (*PayChannelUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayChannelUpdateResponse)
	err := c.cc.Invoke(ctx, PayChannelService_PayChannelUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payChannelServiceClient) PayChannelDelete(ctx context.Context, in *PayChannelDeleteRequest, opts ...grpc.CallOption) (*PayChannelDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayChannelDeleteResponse)
	err := c.cc.Invoke(ctx, PayChannelService_PayChannelDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payChannelServiceClient) PayChannel(ctx context.Context, in *PayChannelRequest, opts ...grpc.CallOption) (*PayChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayChannelResponse)
	err := c.cc.Invoke(ctx, PayChannelService_PayChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payChannelServiceClient) PayChannelRecover(ctx context.Context, in *PayChannelRecoverRequest, opts ...grpc.CallOption) (*PayChannelRecoverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayChannelRecoverResponse)
	err := c.cc.Invoke(ctx, PayChannelService_PayChannelRecover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payChannelServiceClient) PayChannelCode(ctx context.Context, in *PayChannelCodeRequest, opts ...grpc.CallOption) (*PayChannelCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayChannelCodeResponse)
	err := c.cc.Invoke(ctx, PayChannelService_PayChannelCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payChannelServiceClient) PayChannelList(ctx context.Context, in *PayChannelListRequest, opts ...grpc.CallOption) (*PayChannelListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayChannelListResponse)
	err := c.cc.Invoke(ctx, PayChannelService_PayChannelList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayChannelServiceServer is the server API for PayChannelService service.
// All implementations must embed UnimplementedPayChannelServiceServer
// for forward compatibility
//
// PayChannelService 服务
type PayChannelServiceServer interface {
	PayChannelCreate(context.Context, *PayChannelCreateRequest) (*PayChannelCreateResponse, error)
	PayChannelUpdate(context.Context, *PayChannelUpdateRequest) (*PayChannelUpdateResponse, error)
	PayChannelDelete(context.Context, *PayChannelDeleteRequest) (*PayChannelDeleteResponse, error)
	PayChannel(context.Context, *PayChannelRequest) (*PayChannelResponse, error)
	PayChannelRecover(context.Context, *PayChannelRecoverRequest) (*PayChannelRecoverResponse, error)
	PayChannelCode(context.Context, *PayChannelCodeRequest) (*PayChannelCodeResponse, error)
	PayChannelList(context.Context, *PayChannelListRequest) (*PayChannelListResponse, error)
	mustEmbedUnimplementedPayChannelServiceServer()
}

// UnimplementedPayChannelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPayChannelServiceServer struct {
}

func (UnimplementedPayChannelServiceServer) PayChannelCreate(context.Context, *PayChannelCreateRequest) (*PayChannelCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayChannelCreate not implemented")
}
func (UnimplementedPayChannelServiceServer) PayChannelUpdate(context.Context, *PayChannelUpdateRequest) (*PayChannelUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayChannelUpdate not implemented")
}
func (UnimplementedPayChannelServiceServer) PayChannelDelete(context.Context, *PayChannelDeleteRequest) (*PayChannelDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayChannelDelete not implemented")
}
func (UnimplementedPayChannelServiceServer) PayChannel(context.Context, *PayChannelRequest) (*PayChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayChannel not implemented")
}
func (UnimplementedPayChannelServiceServer) PayChannelRecover(context.Context, *PayChannelRecoverRequest) (*PayChannelRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayChannelRecover not implemented")
}
func (UnimplementedPayChannelServiceServer) PayChannelCode(context.Context, *PayChannelCodeRequest) (*PayChannelCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayChannelCode not implemented")
}
func (UnimplementedPayChannelServiceServer) PayChannelList(context.Context, *PayChannelListRequest) (*PayChannelListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayChannelList not implemented")
}
func (UnimplementedPayChannelServiceServer) mustEmbedUnimplementedPayChannelServiceServer() {}

// UnsafePayChannelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayChannelServiceServer will
// result in compilation errors.
type UnsafePayChannelServiceServer interface {
	mustEmbedUnimplementedPayChannelServiceServer()
}

func RegisterPayChannelServiceServer(s grpc.ServiceRegistrar, srv PayChannelServiceServer) {
	s.RegisterService(&PayChannelService_ServiceDesc, srv)
}

func _PayChannelService_PayChannelCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayChannelCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayChannelServiceServer).PayChannelCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayChannelService_PayChannelCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayChannelServiceServer).PayChannelCreate(ctx, req.(*PayChannelCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayChannelService_PayChannelUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayChannelUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayChannelServiceServer).PayChannelUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayChannelService_PayChannelUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayChannelServiceServer).PayChannelUpdate(ctx, req.(*PayChannelUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayChannelService_PayChannelDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayChannelDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayChannelServiceServer).PayChannelDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayChannelService_PayChannelDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayChannelServiceServer).PayChannelDelete(ctx, req.(*PayChannelDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayChannelService_PayChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayChannelServiceServer).PayChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayChannelService_PayChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayChannelServiceServer).PayChannel(ctx, req.(*PayChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayChannelService_PayChannelRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayChannelRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayChannelServiceServer).PayChannelRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayChannelService_PayChannelRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayChannelServiceServer).PayChannelRecover(ctx, req.(*PayChannelRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayChannelService_PayChannelCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayChannelCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayChannelServiceServer).PayChannelCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayChannelService_PayChannelCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayChannelServiceServer).PayChannelCode(ctx, req.(*PayChannelCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayChannelService_PayChannelList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayChannelServiceServer).PayChannelList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayChannelService_PayChannelList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayChannelServiceServer).PayChannelList(ctx, req.(*PayChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PayChannelService_ServiceDesc is the grpc.ServiceDesc for PayChannelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayChannelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "channel.PayChannelService",
	HandlerType: (*PayChannelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PayChannelCreate",
			Handler:    _PayChannelService_PayChannelCreate_Handler,
		},
		{
			MethodName: "PayChannelUpdate",
			Handler:    _PayChannelService_PayChannelUpdate_Handler,
		},
		{
			MethodName: "PayChannelDelete",
			Handler:    _PayChannelService_PayChannelDelete_Handler,
		},
		{
			MethodName: "PayChannel",
			Handler:    _PayChannelService_PayChannel_Handler,
		},
		{
			MethodName: "PayChannelRecover",
			Handler:    _PayChannelService_PayChannelRecover_Handler,
		},
		{
			MethodName: "PayChannelCode",
			Handler:    _PayChannelService_PayChannelCode_Handler,
		},
		{
			MethodName: "PayChannelList",
			Handler:    _PayChannelService_PayChannelList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pay_channel.proto",
}
