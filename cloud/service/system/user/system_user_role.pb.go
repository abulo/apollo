// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: system_user_role.proto

// system_user_role 系统用户和系统角色关联表

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SystemUserRoleObject 数据对象
type SystemUserRoleObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"id" json:"id"
	Id *int64 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id" db:"id"` //自增编号
	// @inject_tag: db:"user_id" json:"userId"
	UserId *int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3,oneof" json:"userId" db:"user_id"` //用户编号
	// @inject_tag: db:"role_id" json:"roleId"
	RoleId *int64 `protobuf:"varint,3,opt,name=role_id,json=roleId,proto3,oneof" json:"roleId" db:"role_id"` //角色编号
	// @inject_tag: db:"deleted" json:"deleted"
	Deleted *int32 `protobuf:"varint,4,opt,name=deleted,proto3,oneof" json:"deleted" db:"deleted"` //删除
	// @inject_tag: db:"tenant_id" json:"tenantId"
	TenantId *int64 `protobuf:"varint,5,opt,name=tenant_id,json=tenantId,proto3,oneof" json:"tenantId" db:"tenant_id"` //租户
	// @inject_tag: db:"creator" json:"creator"
	Creator *string `protobuf:"bytes,6,opt,name=creator,proto3,oneof" json:"creator" db:"creator"` //创建者
	// @inject_tag: db:"create_time" json:"createTime"
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"createTime" db:"create_time"` //创建时间
	// @inject_tag: db:"updater" json:"updater"
	Updater *string `protobuf:"bytes,8,opt,name=updater,proto3,oneof" json:"updater" db:"updater"` //更新者
	// @inject_tag: db:"update_time" json:"updateTime"
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"updateTime" db:"update_time"` //更新时间
}

func (x *SystemUserRoleObject) Reset() {
	*x = SystemUserRoleObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserRoleObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserRoleObject) ProtoMessage() {}

func (x *SystemUserRoleObject) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserRoleObject.ProtoReflect.Descriptor instead.
func (*SystemUserRoleObject) Descriptor() ([]byte, []int) {
	return file_system_user_role_proto_rawDescGZIP(), []int{0}
}

func (x *SystemUserRoleObject) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *SystemUserRoleObject) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *SystemUserRoleObject) GetRoleId() int64 {
	if x != nil && x.RoleId != nil {
		return *x.RoleId
	}
	return 0
}

func (x *SystemUserRoleObject) GetDeleted() int32 {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return 0
}

func (x *SystemUserRoleObject) GetTenantId() int64 {
	if x != nil && x.TenantId != nil {
		return *x.TenantId
	}
	return 0
}

func (x *SystemUserRoleObject) GetCreator() string {
	if x != nil && x.Creator != nil {
		return *x.Creator
	}
	return ""
}

func (x *SystemUserRoleObject) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *SystemUserRoleObject) GetUpdater() string {
	if x != nil && x.Updater != nil {
		return *x.Updater
	}
	return ""
}

func (x *SystemUserRoleObject) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// SystemUserRoleCustomObject 数据对象
type SystemUserRoleCustomObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"user_id" json:"userId"
	UserId *int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3,oneof" json:"userId" db:"user_id"` //系统用户 id
	// @inject_tag: db:"role_ids" json:"roleIds"
	RoleIds []byte `protobuf:"bytes,2,opt,name=role_ids,json=roleIds,proto3,oneof" json:"roleIds" db:"role_ids"` //角色编号 id
	// @inject_tag: db:"deleted" json:"deleted"
	Deleted *int32 `protobuf:"varint,3,opt,name=deleted,proto3,oneof" json:"deleted" db:"deleted"` //删除
	// @inject_tag: db:"tenant_id" json:"tenantId"
	TenantId *int64 `protobuf:"varint,4,opt,name=tenant_id,json=tenantId,proto3,oneof" json:"tenantId" db:"tenant_id"` //租户
	// @inject_tag: db:"creator" json:"creator"
	Creator *string `protobuf:"bytes,5,opt,name=creator,proto3,oneof" json:"creator" db:"creator"` //创建人
	// @inject_tag: db:"create_time" json:"createTime"
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"createTime" db:"create_time"` //创建时间
	// @inject_tag: db:"updater" json:"updater"
	Updater *string `protobuf:"bytes,7,opt,name=updater,proto3,oneof" json:"updater" db:"updater"` //更新人
	// @inject_tag: db:"update_time" json:"updateTime"
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"updateTime" db:"update_time"` //更新时间
}

func (x *SystemUserRoleCustomObject) Reset() {
	*x = SystemUserRoleCustomObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserRoleCustomObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserRoleCustomObject) ProtoMessage() {}

func (x *SystemUserRoleCustomObject) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserRoleCustomObject.ProtoReflect.Descriptor instead.
func (*SystemUserRoleCustomObject) Descriptor() ([]byte, []int) {
	return file_system_user_role_proto_rawDescGZIP(), []int{1}
}

func (x *SystemUserRoleCustomObject) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *SystemUserRoleCustomObject) GetRoleIds() []byte {
	if x != nil {
		return x.RoleIds
	}
	return nil
}

func (x *SystemUserRoleCustomObject) GetDeleted() int32 {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return 0
}

func (x *SystemUserRoleCustomObject) GetTenantId() int64 {
	if x != nil && x.TenantId != nil {
		return *x.TenantId
	}
	return 0
}

func (x *SystemUserRoleCustomObject) GetCreator() string {
	if x != nil && x.Creator != nil {
		return *x.Creator
	}
	return ""
}

func (x *SystemUserRoleCustomObject) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *SystemUserRoleCustomObject) GetUpdater() string {
	if x != nil && x.Updater != nil {
		return *x.Updater
	}
	return ""
}

func (x *SystemUserRoleCustomObject) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// SystemUserRoleCreateRequest 创建数据请求
type SystemUserRoleCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SystemUserRoleCustomObject `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemUserRoleCreateRequest) Reset() {
	*x = SystemUserRoleCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserRoleCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserRoleCreateRequest) ProtoMessage() {}

func (x *SystemUserRoleCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserRoleCreateRequest.ProtoReflect.Descriptor instead.
func (*SystemUserRoleCreateRequest) Descriptor() ([]byte, []int) {
	return file_system_user_role_proto_rawDescGZIP(), []int{2}
}

func (x *SystemUserRoleCreateRequest) GetData() *SystemUserRoleCustomObject {
	if x != nil {
		return x.Data
	}
	return nil
}

// SystemUserRoleCreateResponse 创建数据响应
type SystemUserRoleCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data int64  `protobuf:"varint,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemUserRoleCreateResponse) Reset() {
	*x = SystemUserRoleCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserRoleCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserRoleCreateResponse) ProtoMessage() {}

func (x *SystemUserRoleCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserRoleCreateResponse.ProtoReflect.Descriptor instead.
func (*SystemUserRoleCreateResponse) Descriptor() ([]byte, []int) {
	return file_system_user_role_proto_rawDescGZIP(), []int{3}
}

func (x *SystemUserRoleCreateResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SystemUserRoleCreateResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SystemUserRoleCreateResponse) GetData() int64 {
	if x != nil {
		return x.Data
	}
	return 0
}

// SystemUserRoleListRequest 列表数据
type SystemUserRoleListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"user_id" json:"userId"
	UserId *int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3,oneof" json:"userId" db:"user_id"` // 用户编号
	// @inject_tag: db:"role_id" json:"roleId"
	RoleId *int64 `protobuf:"varint,2,opt,name=role_id,json=roleId,proto3,oneof" json:"roleId" db:"role_id"` // 角色编号
	// @inject_tag: db:"tenant_id" json:"tenantId"
	TenantId *int64 `protobuf:"varint,3,opt,name=tenant_id,json=tenantId,proto3,oneof" json:"tenantId" db:"tenant_id"` // 租户
	// @inject_tag: db:"deleted" json:"deleted"
	Deleted *int32 `protobuf:"varint,4,opt,name=deleted,proto3,oneof" json:"deleted" db:"deleted"` // 删除
}

func (x *SystemUserRoleListRequest) Reset() {
	*x = SystemUserRoleListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserRoleListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserRoleListRequest) ProtoMessage() {}

func (x *SystemUserRoleListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserRoleListRequest.ProtoReflect.Descriptor instead.
func (*SystemUserRoleListRequest) Descriptor() ([]byte, []int) {
	return file_system_user_role_proto_rawDescGZIP(), []int{4}
}

func (x *SystemUserRoleListRequest) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *SystemUserRoleListRequest) GetRoleId() int64 {
	if x != nil && x.RoleId != nil {
		return *x.RoleId
	}
	return 0
}

func (x *SystemUserRoleListRequest) GetTenantId() int64 {
	if x != nil && x.TenantId != nil {
		return *x.TenantId
	}
	return 0
}

func (x *SystemUserRoleListRequest) GetDeleted() int32 {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return 0
}

// SystemUserRoleListResponse 数据响应
type SystemUserRoleListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64                   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []*SystemUserRoleObject `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemUserRoleListResponse) Reset() {
	*x = SystemUserRoleListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserRoleListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserRoleListResponse) ProtoMessage() {}

func (x *SystemUserRoleListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserRoleListResponse.ProtoReflect.Descriptor instead.
func (*SystemUserRoleListResponse) Descriptor() ([]byte, []int) {
	return file_system_user_role_proto_rawDescGZIP(), []int{5}
}

func (x *SystemUserRoleListResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SystemUserRoleListResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SystemUserRoleListResponse) GetData() []*SystemUserRoleObject {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_system_user_role_proto protoreflect.FileDescriptor

var file_system_user_role_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb1, 0x03, 0x0a, 0x14, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52, 0x08, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x72, 0x22, 0x9e, 0x03, 0x0a, 0x1a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x1e, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x01, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x02, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x20, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x03, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x04, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x72, 0x22, 0x53, 0x0a, 0x1b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x58, 0x0a, 0x1c, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xca, 0x01, 0x0a, 0x19, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1c, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x01, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a,
	0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x02, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x22, 0x72, 0x0a, 0x1a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0xcf, 0x01, 0x0a, 0x15, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d,
	0x0a, 0x14, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a,
	0x12, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_user_role_proto_rawDescOnce sync.Once
	file_system_user_role_proto_rawDescData = file_system_user_role_proto_rawDesc
)

func file_system_user_role_proto_rawDescGZIP() []byte {
	file_system_user_role_proto_rawDescOnce.Do(func() {
		file_system_user_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_user_role_proto_rawDescData)
	})
	return file_system_user_role_proto_rawDescData
}

var file_system_user_role_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_system_user_role_proto_goTypes = []any{
	(*SystemUserRoleObject)(nil),         // 0: user.SystemUserRoleObject
	(*SystemUserRoleCustomObject)(nil),   // 1: user.SystemUserRoleCustomObject
	(*SystemUserRoleCreateRequest)(nil),  // 2: user.SystemUserRoleCreateRequest
	(*SystemUserRoleCreateResponse)(nil), // 3: user.SystemUserRoleCreateResponse
	(*SystemUserRoleListRequest)(nil),    // 4: user.SystemUserRoleListRequest
	(*SystemUserRoleListResponse)(nil),   // 5: user.SystemUserRoleListResponse
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
}
var file_system_user_role_proto_depIdxs = []int32{
	6, // 0: user.SystemUserRoleObject.create_time:type_name -> google.protobuf.Timestamp
	6, // 1: user.SystemUserRoleObject.update_time:type_name -> google.protobuf.Timestamp
	6, // 2: user.SystemUserRoleCustomObject.create_time:type_name -> google.protobuf.Timestamp
	6, // 3: user.SystemUserRoleCustomObject.update_time:type_name -> google.protobuf.Timestamp
	1, // 4: user.SystemUserRoleCreateRequest.data:type_name -> user.SystemUserRoleCustomObject
	0, // 5: user.SystemUserRoleListResponse.data:type_name -> user.SystemUserRoleObject
	2, // 6: user.SystemUserRoleService.SystemUserRoleCreate:input_type -> user.SystemUserRoleCreateRequest
	4, // 7: user.SystemUserRoleService.SystemUserRoleList:input_type -> user.SystemUserRoleListRequest
	3, // 8: user.SystemUserRoleService.SystemUserRoleCreate:output_type -> user.SystemUserRoleCreateResponse
	5, // 9: user.SystemUserRoleService.SystemUserRoleList:output_type -> user.SystemUserRoleListResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_system_user_role_proto_init() }
func file_system_user_role_proto_init() {
	if File_system_user_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_user_role_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SystemUserRoleObject); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_user_role_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SystemUserRoleCustomObject); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_user_role_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SystemUserRoleCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_user_role_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SystemUserRoleCreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_user_role_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SystemUserRoleListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_user_role_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SystemUserRoleListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_system_user_role_proto_msgTypes[0].OneofWrappers = []any{}
	file_system_user_role_proto_msgTypes[1].OneofWrappers = []any{}
	file_system_user_role_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_user_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_user_role_proto_goTypes,
		DependencyIndexes: file_system_user_role_proto_depIdxs,
		MessageInfos:      file_system_user_role_proto_msgTypes,
	}.Build()
	File_system_user_role_proto = out.File
	file_system_user_role_proto_rawDesc = nil
	file_system_user_role_proto_goTypes = nil
	file_system_user_role_proto_depIdxs = nil
}
