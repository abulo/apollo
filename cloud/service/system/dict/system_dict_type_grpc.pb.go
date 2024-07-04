// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: system_dict_type.proto

// system_dict_type 字典类型

package dict

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
	SystemDictTypeService_SystemDictTypeCreate_FullMethodName    = "/dict.SystemDictTypeService/SystemDictTypeCreate"
	SystemDictTypeService_SystemDictTypeUpdate_FullMethodName    = "/dict.SystemDictTypeService/SystemDictTypeUpdate"
	SystemDictTypeService_SystemDictTypeDelete_FullMethodName    = "/dict.SystemDictTypeService/SystemDictTypeDelete"
	SystemDictTypeService_SystemDictType_FullMethodName          = "/dict.SystemDictTypeService/SystemDictType"
	SystemDictTypeService_SystemDictTypeList_FullMethodName      = "/dict.SystemDictTypeService/SystemDictTypeList"
	SystemDictTypeService_SystemDictTypeListTotal_FullMethodName = "/dict.SystemDictTypeService/SystemDictTypeListTotal"
)

// SystemDictTypeServiceClient is the client API for SystemDictTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SystemDictTypeService 服务
type SystemDictTypeServiceClient interface {
	SystemDictTypeCreate(ctx context.Context, in *SystemDictTypeCreateRequest, opts ...grpc.CallOption) (*SystemDictTypeCreateResponse, error)
	SystemDictTypeUpdate(ctx context.Context, in *SystemDictTypeUpdateRequest, opts ...grpc.CallOption) (*SystemDictTypeUpdateResponse, error)
	SystemDictTypeDelete(ctx context.Context, in *SystemDictTypeDeleteRequest, opts ...grpc.CallOption) (*SystemDictTypeDeleteResponse, error)
	SystemDictType(ctx context.Context, in *SystemDictTypeRequest, opts ...grpc.CallOption) (*SystemDictTypeResponse, error)
	SystemDictTypeList(ctx context.Context, in *SystemDictTypeListRequest, opts ...grpc.CallOption) (*SystemDictTypeListResponse, error)
	SystemDictTypeListTotal(ctx context.Context, in *SystemDictTypeListTotalRequest, opts ...grpc.CallOption) (*SystemDictTypeTotalResponse, error)
}

type systemDictTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemDictTypeServiceClient(cc grpc.ClientConnInterface) SystemDictTypeServiceClient {
	return &systemDictTypeServiceClient{cc}
}

func (c *systemDictTypeServiceClient) SystemDictTypeCreate(ctx context.Context, in *SystemDictTypeCreateRequest, opts ...grpc.CallOption) (*SystemDictTypeCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemDictTypeCreateResponse)
	err := c.cc.Invoke(ctx, SystemDictTypeService_SystemDictTypeCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemDictTypeServiceClient) SystemDictTypeUpdate(ctx context.Context, in *SystemDictTypeUpdateRequest, opts ...grpc.CallOption) (*SystemDictTypeUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemDictTypeUpdateResponse)
	err := c.cc.Invoke(ctx, SystemDictTypeService_SystemDictTypeUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemDictTypeServiceClient) SystemDictTypeDelete(ctx context.Context, in *SystemDictTypeDeleteRequest, opts ...grpc.CallOption) (*SystemDictTypeDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemDictTypeDeleteResponse)
	err := c.cc.Invoke(ctx, SystemDictTypeService_SystemDictTypeDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemDictTypeServiceClient) SystemDictType(ctx context.Context, in *SystemDictTypeRequest, opts ...grpc.CallOption) (*SystemDictTypeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemDictTypeResponse)
	err := c.cc.Invoke(ctx, SystemDictTypeService_SystemDictType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemDictTypeServiceClient) SystemDictTypeList(ctx context.Context, in *SystemDictTypeListRequest, opts ...grpc.CallOption) (*SystemDictTypeListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemDictTypeListResponse)
	err := c.cc.Invoke(ctx, SystemDictTypeService_SystemDictTypeList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemDictTypeServiceClient) SystemDictTypeListTotal(ctx context.Context, in *SystemDictTypeListTotalRequest, opts ...grpc.CallOption) (*SystemDictTypeTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemDictTypeTotalResponse)
	err := c.cc.Invoke(ctx, SystemDictTypeService_SystemDictTypeListTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemDictTypeServiceServer is the server API for SystemDictTypeService service.
// All implementations must embed UnimplementedSystemDictTypeServiceServer
// for forward compatibility
//
// SystemDictTypeService 服务
type SystemDictTypeServiceServer interface {
	SystemDictTypeCreate(context.Context, *SystemDictTypeCreateRequest) (*SystemDictTypeCreateResponse, error)
	SystemDictTypeUpdate(context.Context, *SystemDictTypeUpdateRequest) (*SystemDictTypeUpdateResponse, error)
	SystemDictTypeDelete(context.Context, *SystemDictTypeDeleteRequest) (*SystemDictTypeDeleteResponse, error)
	SystemDictType(context.Context, *SystemDictTypeRequest) (*SystemDictTypeResponse, error)
	SystemDictTypeList(context.Context, *SystemDictTypeListRequest) (*SystemDictTypeListResponse, error)
	SystemDictTypeListTotal(context.Context, *SystemDictTypeListTotalRequest) (*SystemDictTypeTotalResponse, error)
	mustEmbedUnimplementedSystemDictTypeServiceServer()
}

// UnimplementedSystemDictTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemDictTypeServiceServer struct {
}

func (UnimplementedSystemDictTypeServiceServer) SystemDictTypeCreate(context.Context, *SystemDictTypeCreateRequest) (*SystemDictTypeCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemDictTypeCreate not implemented")
}
func (UnimplementedSystemDictTypeServiceServer) SystemDictTypeUpdate(context.Context, *SystemDictTypeUpdateRequest) (*SystemDictTypeUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemDictTypeUpdate not implemented")
}
func (UnimplementedSystemDictTypeServiceServer) SystemDictTypeDelete(context.Context, *SystemDictTypeDeleteRequest) (*SystemDictTypeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemDictTypeDelete not implemented")
}
func (UnimplementedSystemDictTypeServiceServer) SystemDictType(context.Context, *SystemDictTypeRequest) (*SystemDictTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemDictType not implemented")
}
func (UnimplementedSystemDictTypeServiceServer) SystemDictTypeList(context.Context, *SystemDictTypeListRequest) (*SystemDictTypeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemDictTypeList not implemented")
}
func (UnimplementedSystemDictTypeServiceServer) SystemDictTypeListTotal(context.Context, *SystemDictTypeListTotalRequest) (*SystemDictTypeTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemDictTypeListTotal not implemented")
}
func (UnimplementedSystemDictTypeServiceServer) mustEmbedUnimplementedSystemDictTypeServiceServer() {}

// UnsafeSystemDictTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemDictTypeServiceServer will
// result in compilation errors.
type UnsafeSystemDictTypeServiceServer interface {
	mustEmbedUnimplementedSystemDictTypeServiceServer()
}

func RegisterSystemDictTypeServiceServer(s grpc.ServiceRegistrar, srv SystemDictTypeServiceServer) {
	s.RegisterService(&SystemDictTypeService_ServiceDesc, srv)
}

func _SystemDictTypeService_SystemDictTypeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemDictTypeCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemDictTypeServiceServer).SystemDictTypeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemDictTypeService_SystemDictTypeCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemDictTypeServiceServer).SystemDictTypeCreate(ctx, req.(*SystemDictTypeCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemDictTypeService_SystemDictTypeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemDictTypeUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemDictTypeServiceServer).SystemDictTypeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemDictTypeService_SystemDictTypeUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemDictTypeServiceServer).SystemDictTypeUpdate(ctx, req.(*SystemDictTypeUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemDictTypeService_SystemDictTypeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemDictTypeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemDictTypeServiceServer).SystemDictTypeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemDictTypeService_SystemDictTypeDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemDictTypeServiceServer).SystemDictTypeDelete(ctx, req.(*SystemDictTypeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemDictTypeService_SystemDictType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemDictTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemDictTypeServiceServer).SystemDictType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemDictTypeService_SystemDictType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemDictTypeServiceServer).SystemDictType(ctx, req.(*SystemDictTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemDictTypeService_SystemDictTypeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemDictTypeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemDictTypeServiceServer).SystemDictTypeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemDictTypeService_SystemDictTypeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemDictTypeServiceServer).SystemDictTypeList(ctx, req.(*SystemDictTypeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemDictTypeService_SystemDictTypeListTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemDictTypeListTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemDictTypeServiceServer).SystemDictTypeListTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemDictTypeService_SystemDictTypeListTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemDictTypeServiceServer).SystemDictTypeListTotal(ctx, req.(*SystemDictTypeListTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemDictTypeService_ServiceDesc is the grpc.ServiceDesc for SystemDictTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemDictTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dict.SystemDictTypeService",
	HandlerType: (*SystemDictTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemDictTypeCreate",
			Handler:    _SystemDictTypeService_SystemDictTypeCreate_Handler,
		},
		{
			MethodName: "SystemDictTypeUpdate",
			Handler:    _SystemDictTypeService_SystemDictTypeUpdate_Handler,
		},
		{
			MethodName: "SystemDictTypeDelete",
			Handler:    _SystemDictTypeService_SystemDictTypeDelete_Handler,
		},
		{
			MethodName: "SystemDictType",
			Handler:    _SystemDictTypeService_SystemDictType_Handler,
		},
		{
			MethodName: "SystemDictTypeList",
			Handler:    _SystemDictTypeService_SystemDictTypeList_Handler,
		},
		{
			MethodName: "SystemDictTypeListTotal",
			Handler:    _SystemDictTypeService_SystemDictTypeListTotal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_dict_type.proto",
}
