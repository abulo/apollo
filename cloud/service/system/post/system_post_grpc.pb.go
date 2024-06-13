// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: system_post.proto

// system_post 职位

package post

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
	SystemPostService_SystemPostCreate_FullMethodName    = "/post.SystemPostService/SystemPostCreate"
	SystemPostService_SystemPostUpdate_FullMethodName    = "/post.SystemPostService/SystemPostUpdate"
	SystemPostService_SystemPostDelete_FullMethodName    = "/post.SystemPostService/SystemPostDelete"
	SystemPostService_SystemPost_FullMethodName          = "/post.SystemPostService/SystemPost"
	SystemPostService_SystemPostRecover_FullMethodName   = "/post.SystemPostService/SystemPostRecover"
	SystemPostService_SystemPostList_FullMethodName      = "/post.SystemPostService/SystemPostList"
	SystemPostService_SystemPostListTotal_FullMethodName = "/post.SystemPostService/SystemPostListTotal"
)

// SystemPostServiceClient is the client API for SystemPostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SystemPostService 服务
type SystemPostServiceClient interface {
	SystemPostCreate(ctx context.Context, in *SystemPostCreateRequest, opts ...grpc.CallOption) (*SystemPostCreateResponse, error)
	SystemPostUpdate(ctx context.Context, in *SystemPostUpdateRequest, opts ...grpc.CallOption) (*SystemPostUpdateResponse, error)
	SystemPostDelete(ctx context.Context, in *SystemPostDeleteRequest, opts ...grpc.CallOption) (*SystemPostDeleteResponse, error)
	SystemPost(ctx context.Context, in *SystemPostRequest, opts ...grpc.CallOption) (*SystemPostResponse, error)
	SystemPostRecover(ctx context.Context, in *SystemPostRecoverRequest, opts ...grpc.CallOption) (*SystemPostRecoverResponse, error)
	SystemPostList(ctx context.Context, in *SystemPostListRequest, opts ...grpc.CallOption) (*SystemPostListResponse, error)
	SystemPostListTotal(ctx context.Context, in *SystemPostListTotalRequest, opts ...grpc.CallOption) (*SystemPostTotalResponse, error)
}

type systemPostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemPostServiceClient(cc grpc.ClientConnInterface) SystemPostServiceClient {
	return &systemPostServiceClient{cc}
}

func (c *systemPostServiceClient) SystemPostCreate(ctx context.Context, in *SystemPostCreateRequest, opts ...grpc.CallOption) (*SystemPostCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemPostCreateResponse)
	err := c.cc.Invoke(ctx, SystemPostService_SystemPostCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemPostServiceClient) SystemPostUpdate(ctx context.Context, in *SystemPostUpdateRequest, opts ...grpc.CallOption) (*SystemPostUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemPostUpdateResponse)
	err := c.cc.Invoke(ctx, SystemPostService_SystemPostUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemPostServiceClient) SystemPostDelete(ctx context.Context, in *SystemPostDeleteRequest, opts ...grpc.CallOption) (*SystemPostDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemPostDeleteResponse)
	err := c.cc.Invoke(ctx, SystemPostService_SystemPostDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemPostServiceClient) SystemPost(ctx context.Context, in *SystemPostRequest, opts ...grpc.CallOption) (*SystemPostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemPostResponse)
	err := c.cc.Invoke(ctx, SystemPostService_SystemPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemPostServiceClient) SystemPostRecover(ctx context.Context, in *SystemPostRecoverRequest, opts ...grpc.CallOption) (*SystemPostRecoverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemPostRecoverResponse)
	err := c.cc.Invoke(ctx, SystemPostService_SystemPostRecover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemPostServiceClient) SystemPostList(ctx context.Context, in *SystemPostListRequest, opts ...grpc.CallOption) (*SystemPostListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemPostListResponse)
	err := c.cc.Invoke(ctx, SystemPostService_SystemPostList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemPostServiceClient) SystemPostListTotal(ctx context.Context, in *SystemPostListTotalRequest, opts ...grpc.CallOption) (*SystemPostTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemPostTotalResponse)
	err := c.cc.Invoke(ctx, SystemPostService_SystemPostListTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemPostServiceServer is the server API for SystemPostService service.
// All implementations must embed UnimplementedSystemPostServiceServer
// for forward compatibility
//
// SystemPostService 服务
type SystemPostServiceServer interface {
	SystemPostCreate(context.Context, *SystemPostCreateRequest) (*SystemPostCreateResponse, error)
	SystemPostUpdate(context.Context, *SystemPostUpdateRequest) (*SystemPostUpdateResponse, error)
	SystemPostDelete(context.Context, *SystemPostDeleteRequest) (*SystemPostDeleteResponse, error)
	SystemPost(context.Context, *SystemPostRequest) (*SystemPostResponse, error)
	SystemPostRecover(context.Context, *SystemPostRecoverRequest) (*SystemPostRecoverResponse, error)
	SystemPostList(context.Context, *SystemPostListRequest) (*SystemPostListResponse, error)
	SystemPostListTotal(context.Context, *SystemPostListTotalRequest) (*SystemPostTotalResponse, error)
	mustEmbedUnimplementedSystemPostServiceServer()
}

// UnimplementedSystemPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemPostServiceServer struct {
}

func (UnimplementedSystemPostServiceServer) SystemPostCreate(context.Context, *SystemPostCreateRequest) (*SystemPostCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemPostCreate not implemented")
}
func (UnimplementedSystemPostServiceServer) SystemPostUpdate(context.Context, *SystemPostUpdateRequest) (*SystemPostUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemPostUpdate not implemented")
}
func (UnimplementedSystemPostServiceServer) SystemPostDelete(context.Context, *SystemPostDeleteRequest) (*SystemPostDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemPostDelete not implemented")
}
func (UnimplementedSystemPostServiceServer) SystemPost(context.Context, *SystemPostRequest) (*SystemPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemPost not implemented")
}
func (UnimplementedSystemPostServiceServer) SystemPostRecover(context.Context, *SystemPostRecoverRequest) (*SystemPostRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemPostRecover not implemented")
}
func (UnimplementedSystemPostServiceServer) SystemPostList(context.Context, *SystemPostListRequest) (*SystemPostListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemPostList not implemented")
}
func (UnimplementedSystemPostServiceServer) SystemPostListTotal(context.Context, *SystemPostListTotalRequest) (*SystemPostTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemPostListTotal not implemented")
}
func (UnimplementedSystemPostServiceServer) mustEmbedUnimplementedSystemPostServiceServer() {}

// UnsafeSystemPostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemPostServiceServer will
// result in compilation errors.
type UnsafeSystemPostServiceServer interface {
	mustEmbedUnimplementedSystemPostServiceServer()
}

func RegisterSystemPostServiceServer(s grpc.ServiceRegistrar, srv SystemPostServiceServer) {
	s.RegisterService(&SystemPostService_ServiceDesc, srv)
}

func _SystemPostService_SystemPostCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemPostCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemPostServiceServer).SystemPostCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemPostService_SystemPostCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemPostServiceServer).SystemPostCreate(ctx, req.(*SystemPostCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemPostService_SystemPostUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemPostUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemPostServiceServer).SystemPostUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemPostService_SystemPostUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemPostServiceServer).SystemPostUpdate(ctx, req.(*SystemPostUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemPostService_SystemPostDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemPostDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemPostServiceServer).SystemPostDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemPostService_SystemPostDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemPostServiceServer).SystemPostDelete(ctx, req.(*SystemPostDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemPostService_SystemPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemPostServiceServer).SystemPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemPostService_SystemPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemPostServiceServer).SystemPost(ctx, req.(*SystemPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemPostService_SystemPostRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemPostRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemPostServiceServer).SystemPostRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemPostService_SystemPostRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemPostServiceServer).SystemPostRecover(ctx, req.(*SystemPostRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemPostService_SystemPostList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemPostListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemPostServiceServer).SystemPostList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemPostService_SystemPostList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemPostServiceServer).SystemPostList(ctx, req.(*SystemPostListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemPostService_SystemPostListTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemPostListTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemPostServiceServer).SystemPostListTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemPostService_SystemPostListTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemPostServiceServer).SystemPostListTotal(ctx, req.(*SystemPostListTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemPostService_ServiceDesc is the grpc.ServiceDesc for SystemPostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemPostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post.SystemPostService",
	HandlerType: (*SystemPostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemPostCreate",
			Handler:    _SystemPostService_SystemPostCreate_Handler,
		},
		{
			MethodName: "SystemPostUpdate",
			Handler:    _SystemPostService_SystemPostUpdate_Handler,
		},
		{
			MethodName: "SystemPostDelete",
			Handler:    _SystemPostService_SystemPostDelete_Handler,
		},
		{
			MethodName: "SystemPost",
			Handler:    _SystemPostService_SystemPost_Handler,
		},
		{
			MethodName: "SystemPostRecover",
			Handler:    _SystemPostService_SystemPostRecover_Handler,
		},
		{
			MethodName: "SystemPostList",
			Handler:    _SystemPostService_SystemPostList_Handler,
		},
		{
			MethodName: "SystemPostListTotal",
			Handler:    _SystemPostService_SystemPostListTotal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_post.proto",
}
