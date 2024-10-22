// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: pay_wallet_recharge_package.proto

// pay_wallet_recharge_package 充值套餐表

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
	PayWalletRechargePackageService_PayWalletRechargePackageCreate_FullMethodName    = "/wallet.PayWalletRechargePackageService/PayWalletRechargePackageCreate"
	PayWalletRechargePackageService_PayWalletRechargePackageUpdate_FullMethodName    = "/wallet.PayWalletRechargePackageService/PayWalletRechargePackageUpdate"
	PayWalletRechargePackageService_PayWalletRechargePackageDelete_FullMethodName    = "/wallet.PayWalletRechargePackageService/PayWalletRechargePackageDelete"
	PayWalletRechargePackageService_PayWalletRechargePackage_FullMethodName          = "/wallet.PayWalletRechargePackageService/PayWalletRechargePackage"
	PayWalletRechargePackageService_PayWalletRechargePackageRecover_FullMethodName   = "/wallet.PayWalletRechargePackageService/PayWalletRechargePackageRecover"
	PayWalletRechargePackageService_PayWalletRechargePackageList_FullMethodName      = "/wallet.PayWalletRechargePackageService/PayWalletRechargePackageList"
	PayWalletRechargePackageService_PayWalletRechargePackageListTotal_FullMethodName = "/wallet.PayWalletRechargePackageService/PayWalletRechargePackageListTotal"
)

// PayWalletRechargePackageServiceClient is the client API for PayWalletRechargePackageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// PayWalletRechargePackageService 服务
type PayWalletRechargePackageServiceClient interface {
	PayWalletRechargePackageCreate(ctx context.Context, in *PayWalletRechargePackageCreateRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageCreateResponse, error)
	PayWalletRechargePackageUpdate(ctx context.Context, in *PayWalletRechargePackageUpdateRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageUpdateResponse, error)
	PayWalletRechargePackageDelete(ctx context.Context, in *PayWalletRechargePackageDeleteRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageDeleteResponse, error)
	PayWalletRechargePackage(ctx context.Context, in *PayWalletRechargePackageRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageResponse, error)
	PayWalletRechargePackageRecover(ctx context.Context, in *PayWalletRechargePackageRecoverRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageRecoverResponse, error)
	PayWalletRechargePackageList(ctx context.Context, in *PayWalletRechargePackageListRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageListResponse, error)
	PayWalletRechargePackageListTotal(ctx context.Context, in *PayWalletRechargePackageListTotalRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageTotalResponse, error)
}

type payWalletRechargePackageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPayWalletRechargePackageServiceClient(cc grpc.ClientConnInterface) PayWalletRechargePackageServiceClient {
	return &payWalletRechargePackageServiceClient{cc}
}

func (c *payWalletRechargePackageServiceClient) PayWalletRechargePackageCreate(ctx context.Context, in *PayWalletRechargePackageCreateRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletRechargePackageCreateResponse)
	err := c.cc.Invoke(ctx, PayWalletRechargePackageService_PayWalletRechargePackageCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payWalletRechargePackageServiceClient) PayWalletRechargePackageUpdate(ctx context.Context, in *PayWalletRechargePackageUpdateRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletRechargePackageUpdateResponse)
	err := c.cc.Invoke(ctx, PayWalletRechargePackageService_PayWalletRechargePackageUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payWalletRechargePackageServiceClient) PayWalletRechargePackageDelete(ctx context.Context, in *PayWalletRechargePackageDeleteRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletRechargePackageDeleteResponse)
	err := c.cc.Invoke(ctx, PayWalletRechargePackageService_PayWalletRechargePackageDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payWalletRechargePackageServiceClient) PayWalletRechargePackage(ctx context.Context, in *PayWalletRechargePackageRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletRechargePackageResponse)
	err := c.cc.Invoke(ctx, PayWalletRechargePackageService_PayWalletRechargePackage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payWalletRechargePackageServiceClient) PayWalletRechargePackageRecover(ctx context.Context, in *PayWalletRechargePackageRecoverRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageRecoverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletRechargePackageRecoverResponse)
	err := c.cc.Invoke(ctx, PayWalletRechargePackageService_PayWalletRechargePackageRecover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payWalletRechargePackageServiceClient) PayWalletRechargePackageList(ctx context.Context, in *PayWalletRechargePackageListRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletRechargePackageListResponse)
	err := c.cc.Invoke(ctx, PayWalletRechargePackageService_PayWalletRechargePackageList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payWalletRechargePackageServiceClient) PayWalletRechargePackageListTotal(ctx context.Context, in *PayWalletRechargePackageListTotalRequest, opts ...grpc.CallOption) (*PayWalletRechargePackageTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayWalletRechargePackageTotalResponse)
	err := c.cc.Invoke(ctx, PayWalletRechargePackageService_PayWalletRechargePackageListTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayWalletRechargePackageServiceServer is the server API for PayWalletRechargePackageService service.
// All implementations must embed UnimplementedPayWalletRechargePackageServiceServer
// for forward compatibility
//
// PayWalletRechargePackageService 服务
type PayWalletRechargePackageServiceServer interface {
	PayWalletRechargePackageCreate(context.Context, *PayWalletRechargePackageCreateRequest) (*PayWalletRechargePackageCreateResponse, error)
	PayWalletRechargePackageUpdate(context.Context, *PayWalletRechargePackageUpdateRequest) (*PayWalletRechargePackageUpdateResponse, error)
	PayWalletRechargePackageDelete(context.Context, *PayWalletRechargePackageDeleteRequest) (*PayWalletRechargePackageDeleteResponse, error)
	PayWalletRechargePackage(context.Context, *PayWalletRechargePackageRequest) (*PayWalletRechargePackageResponse, error)
	PayWalletRechargePackageRecover(context.Context, *PayWalletRechargePackageRecoverRequest) (*PayWalletRechargePackageRecoverResponse, error)
	PayWalletRechargePackageList(context.Context, *PayWalletRechargePackageListRequest) (*PayWalletRechargePackageListResponse, error)
	PayWalletRechargePackageListTotal(context.Context, *PayWalletRechargePackageListTotalRequest) (*PayWalletRechargePackageTotalResponse, error)
	mustEmbedUnimplementedPayWalletRechargePackageServiceServer()
}

// UnimplementedPayWalletRechargePackageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPayWalletRechargePackageServiceServer struct {
}

func (UnimplementedPayWalletRechargePackageServiceServer) PayWalletRechargePackageCreate(context.Context, *PayWalletRechargePackageCreateRequest) (*PayWalletRechargePackageCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletRechargePackageCreate not implemented")
}
func (UnimplementedPayWalletRechargePackageServiceServer) PayWalletRechargePackageUpdate(context.Context, *PayWalletRechargePackageUpdateRequest) (*PayWalletRechargePackageUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletRechargePackageUpdate not implemented")
}
func (UnimplementedPayWalletRechargePackageServiceServer) PayWalletRechargePackageDelete(context.Context, *PayWalletRechargePackageDeleteRequest) (*PayWalletRechargePackageDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletRechargePackageDelete not implemented")
}
func (UnimplementedPayWalletRechargePackageServiceServer) PayWalletRechargePackage(context.Context, *PayWalletRechargePackageRequest) (*PayWalletRechargePackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletRechargePackage not implemented")
}
func (UnimplementedPayWalletRechargePackageServiceServer) PayWalletRechargePackageRecover(context.Context, *PayWalletRechargePackageRecoverRequest) (*PayWalletRechargePackageRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletRechargePackageRecover not implemented")
}
func (UnimplementedPayWalletRechargePackageServiceServer) PayWalletRechargePackageList(context.Context, *PayWalletRechargePackageListRequest) (*PayWalletRechargePackageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletRechargePackageList not implemented")
}
func (UnimplementedPayWalletRechargePackageServiceServer) PayWalletRechargePackageListTotal(context.Context, *PayWalletRechargePackageListTotalRequest) (*PayWalletRechargePackageTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayWalletRechargePackageListTotal not implemented")
}
func (UnimplementedPayWalletRechargePackageServiceServer) mustEmbedUnimplementedPayWalletRechargePackageServiceServer() {
}

// UnsafePayWalletRechargePackageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayWalletRechargePackageServiceServer will
// result in compilation errors.
type UnsafePayWalletRechargePackageServiceServer interface {
	mustEmbedUnimplementedPayWalletRechargePackageServiceServer()
}

func RegisterPayWalletRechargePackageServiceServer(s grpc.ServiceRegistrar, srv PayWalletRechargePackageServiceServer) {
	s.RegisterService(&PayWalletRechargePackageService_ServiceDesc, srv)
}

func _PayWalletRechargePackageService_PayWalletRechargePackageCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletRechargePackageCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackageCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletRechargePackageService_PayWalletRechargePackageCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackageCreate(ctx, req.(*PayWalletRechargePackageCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayWalletRechargePackageService_PayWalletRechargePackageUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletRechargePackageUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackageUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletRechargePackageService_PayWalletRechargePackageUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackageUpdate(ctx, req.(*PayWalletRechargePackageUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayWalletRechargePackageService_PayWalletRechargePackageDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletRechargePackageDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackageDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletRechargePackageService_PayWalletRechargePackageDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackageDelete(ctx, req.(*PayWalletRechargePackageDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayWalletRechargePackageService_PayWalletRechargePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletRechargePackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletRechargePackageService_PayWalletRechargePackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackage(ctx, req.(*PayWalletRechargePackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayWalletRechargePackageService_PayWalletRechargePackageRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletRechargePackageRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackageRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletRechargePackageService_PayWalletRechargePackageRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackageRecover(ctx, req.(*PayWalletRechargePackageRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayWalletRechargePackageService_PayWalletRechargePackageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletRechargePackageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletRechargePackageService_PayWalletRechargePackageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackageList(ctx, req.(*PayWalletRechargePackageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayWalletRechargePackageService_PayWalletRechargePackageListTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayWalletRechargePackageListTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackageListTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayWalletRechargePackageService_PayWalletRechargePackageListTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayWalletRechargePackageServiceServer).PayWalletRechargePackageListTotal(ctx, req.(*PayWalletRechargePackageListTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PayWalletRechargePackageService_ServiceDesc is the grpc.ServiceDesc for PayWalletRechargePackageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayWalletRechargePackageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.PayWalletRechargePackageService",
	HandlerType: (*PayWalletRechargePackageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PayWalletRechargePackageCreate",
			Handler:    _PayWalletRechargePackageService_PayWalletRechargePackageCreate_Handler,
		},
		{
			MethodName: "PayWalletRechargePackageUpdate",
			Handler:    _PayWalletRechargePackageService_PayWalletRechargePackageUpdate_Handler,
		},
		{
			MethodName: "PayWalletRechargePackageDelete",
			Handler:    _PayWalletRechargePackageService_PayWalletRechargePackageDelete_Handler,
		},
		{
			MethodName: "PayWalletRechargePackage",
			Handler:    _PayWalletRechargePackageService_PayWalletRechargePackage_Handler,
		},
		{
			MethodName: "PayWalletRechargePackageRecover",
			Handler:    _PayWalletRechargePackageService_PayWalletRechargePackageRecover_Handler,
		},
		{
			MethodName: "PayWalletRechargePackageList",
			Handler:    _PayWalletRechargePackageService_PayWalletRechargePackageList_Handler,
		},
		{
			MethodName: "PayWalletRechargePackageListTotal",
			Handler:    _PayWalletRechargePackageService_PayWalletRechargePackageListTotal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pay_wallet_recharge_package.proto",
}
