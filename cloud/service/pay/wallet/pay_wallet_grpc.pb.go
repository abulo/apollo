// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: wallet/pay_wallet.proto

// pay_wallet 会员钱包表

package wallet

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
	PayWalletService_PayWalletCreate_FullMethodName    = "/wallet.PayWalletService/PayWalletCreate"
	PayWalletService_PayWalletUpdate_FullMethodName    = "/wallet.PayWalletService/PayWalletUpdate"
	PayWalletService_PayWalletDelete_FullMethodName    = "/wallet.PayWalletService/PayWalletDelete"
	PayWalletService_PayWallet_FullMethodName          = "/wallet.PayWalletService/PayWallet"
	PayWalletService_PayWalletRecover_FullMethodName   = "/wallet.PayWalletService/PayWalletRecover"
	PayWalletService_PayWalletList_FullMethodName      = "/wallet.PayWalletService/PayWalletList"
	PayWalletService_PayWalletListTotal_FullMethodName = "/wallet.PayWalletService/PayWalletListTotal"
)

// PayWalletServiceClient is the client API for PayWalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// PayWalletService 服务
type PayWalletServiceClient interface {
	PayWalletCreate(ctx context.Context, in *PayWalletCreateRequest, opts ...grpc.CallOption) (*PayWalletCreateResponse, error)
	PayWalletUpdate(ctx context.Context, in *PayWalletUpdateRequest, opts ...grpc.CallOption) (*PayWalletUpdateResponse, error)
	PayWalletDelete(ctx context.Context, in *PayWalletDeleteRequest, opts ...grpc.CallOption) (*PayWalletDeleteResponse, error)
	PayWallet(ctx context.Context, in *PayWalletRequest, opts ...grpc.CallOption) (*PayWalletResponse, error)
	PayWalletRecover(ctx context.Context, in *PayWalletRecoverRequest, opts ...grpc.CallOption) (*PayWalletRecoverResponse, error)
	PayWalletList(ctx context.Context, in *PayWalletListRequest, opts ...grpc.CallOption) (*PayWalletListResponse, error)
	PayWalletListTotal(ctx context.Context, in *PayWalletListTotalRequest, opts ...grpc.CallOption) (*PayWalletTotalResponse, error)
}

type payWalletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPayWalletServiceClient(cc grpc.ClientConnInterface) PayWalletServiceClient {
	return &payWalletServiceClient{cc}
}

func (c *payWalletServiceClient) PayWalletCreate(ctx context.Context, in *PayWalletCreateRequest, opts ...grpc.CallOption) (*PayWalletCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletCreateResponse)
	err := c.cc.Invoke(ctx, PayWalletService_PayWalletCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payWalletServiceClient) PayWalletUpdate(ctx context.Context, in *PayWalletUpdateRequest, opts ...grpc.CallOption) (*PayWalletUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletUpdateResponse)
	err := c.cc.Invoke(ctx, PayWalletService_PayWalletUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payWalletServiceClient) PayWalletDelete(ctx context.Context, in *PayWalletDeleteRequest, opts ...grpc.CallOption) (*PayWalletDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletDeleteResponse)
	err := c.cc.Invoke(ctx, PayWalletService_PayWalletDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payWalletServiceClient) PayWallet(ctx context.Context, in *PayWalletRequest, opts ...grpc.CallOption) (*PayWalletResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletResponse)
	err := c.cc.Invoke(ctx, PayWalletService_PayWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payWalletServiceClient) PayWalletRecover(ctx context.Context, in *PayWalletRecoverRequest, opts ...grpc.CallOption) (*PayWalletRecoverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletRecoverResponse)
	err := c.cc.Invoke(ctx, PayWalletService_PayWalletRecover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payWalletServiceClient) PayWalletList(ctx context.Context, in *PayWalletListRequest, opts ...grpc.CallOption) (*PayWalletListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletListResponse)
	err := c.cc.Invoke(ctx, PayWalletService_PayWalletList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payWalletServiceClient) PayWalletListTotal(ctx context.Context, in *PayWalletListTotalRequest, opts ...grpc.CallOption) (*PayWalletTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletTotalResponse)
	err := c.cc.Invoke(ctx, PayWalletService_PayWalletListTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayWalletServiceServer is the server API for PayWalletService service.
// All implementations must embed UnimplementedPayWalletServiceServer
// for forward compatibility
//
// PayWalletService 服务
type PayWalletServiceServer interface {
	PayWalletCreate(context.Context, *PayWalletCreateRequest) (*PayWalletCreateResponse, error)
	PayWalletUpdate(context.Context, *PayWalletUpdateRequest) (*PayWalletUpdateResponse, error)
	PayWalletDelete(context.Context, *PayWalletDeleteRequest) (*PayWalletDeleteResponse, error)
	PayWallet(context.Context, *PayWalletRequest) (*PayWalletResponse, error)
	PayWalletRecover(context.Context, *PayWalletRecoverRequest) (*PayWalletRecoverResponse, error)
	PayWalletList(context.Context, *PayWalletListRequest) (*PayWalletListResponse, error)
	PayWalletListTotal(context.Context, *PayWalletListTotalRequest) (*PayWalletTotalResponse, error)
	mustEmbedUnimplementedPayWalletServiceServer()
}

// UnimplementedPayWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPayWalletServiceServer struct {
}

func (UnimplementedPayWalletServiceServer) PayWalletCreate(context.Context, *PayWalletCreateRequest) (*PayWalletCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletCreate not implemented")
}
func (UnimplementedPayWalletServiceServer) PayWalletUpdate(context.Context, *PayWalletUpdateRequest) (*PayWalletUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletUpdate not implemented")
}
func (UnimplementedPayWalletServiceServer) PayWalletDelete(context.Context, *PayWalletDeleteRequest) (*PayWalletDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletDelete not implemented")
}
func (UnimplementedPayWalletServiceServer) PayWallet(context.Context, *PayWalletRequest) (*PayWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWallet not implemented")
}
func (UnimplementedPayWalletServiceServer) PayWalletRecover(context.Context, *PayWalletRecoverRequest) (*PayWalletRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletRecover not implemented")
}
func (UnimplementedPayWalletServiceServer) PayWalletList(context.Context, *PayWalletListRequest) (*PayWalletListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletList not implemented")
}
func (UnimplementedPayWalletServiceServer) PayWalletListTotal(context.Context, *PayWalletListTotalRequest) (*PayWalletTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletListTotal not implemented")
}
func (UnimplementedPayWalletServiceServer) mustEmbedUnimplementedPayWalletServiceServer() {}

// UnsafePayWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayWalletServiceServer will
// result in compilation errors.
type UnsafePayWalletServiceServer interface {
	mustEmbedUnimplementedPayWalletServiceServer()
}

func RegisterPayWalletServiceServer(s grpc.ServiceRegistrar, srv PayWalletServiceServer) {
	s.RegisterService(&PayWalletService_ServiceDesc, srv)
}

func _PayWalletService_PayWalletCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletServiceServer).PayWalletCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletService_PayWalletCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletServiceServer).PayWalletCreate(ctx, req.(*PayWalletCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayWalletService_PayWalletUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletServiceServer).PayWalletUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletService_PayWalletUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletServiceServer).PayWalletUpdate(ctx, req.(*PayWalletUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayWalletService_PayWalletDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletServiceServer).PayWalletDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletService_PayWalletDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletServiceServer).PayWalletDelete(ctx, req.(*PayWalletDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayWalletService_PayWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletServiceServer).PayWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletService_PayWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletServiceServer).PayWallet(ctx, req.(*PayWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayWalletService_PayWalletRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletServiceServer).PayWalletRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletService_PayWalletRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletServiceServer).PayWalletRecover(ctx, req.(*PayWalletRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayWalletService_PayWalletList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletServiceServer).PayWalletList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletService_PayWalletList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletServiceServer).PayWalletList(ctx, req.(*PayWalletListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayWalletService_PayWalletListTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletListTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletServiceServer).PayWalletListTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletService_PayWalletListTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletServiceServer).PayWalletListTotal(ctx, req.(*PayWalletListTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PayWalletService_ServiceDesc is the grpc.ServiceDesc for PayWalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayWalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.PayWalletService",
	HandlerType: (*PayWalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PayWalletCreate",
			Handler:    _PayWalletService_PayWalletCreate_Handler,
		},
		{
			MethodName: "PayWalletUpdate",
			Handler:    _PayWalletService_PayWalletUpdate_Handler,
		},
		{
			MethodName: "PayWalletDelete",
			Handler:    _PayWalletService_PayWalletDelete_Handler,
		},
		{
			MethodName: "PayWallet",
			Handler:    _PayWalletService_PayWallet_Handler,
		},
		{
			MethodName: "PayWalletRecover",
			Handler:    _PayWalletService_PayWalletRecover_Handler,
		},
		{
			MethodName: "PayWalletList",
			Handler:    _PayWalletService_PayWalletList_Handler,
		},
		{
			MethodName: "PayWalletListTotal",
			Handler:    _PayWalletService_PayWalletListTotal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet/pay_wallet.proto",
}
