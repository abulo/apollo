// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: role/system_role_menu.proto

// system_role_menu 系统角色和系统菜单关联表

package role

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

// SystemRoleMenuObject 数据对象
type SystemRoleMenuObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"id" json:"id"
	Id *int64 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id" db:"id"` //自增编号
	// @inject_tag: db:"system_role_id" json:"systemRoleId"
	SystemRoleId *int64 `protobuf:"varint,2,opt,name=system_role_id,json=systemRoleId,proto3,oneof" json:"systemRoleId" db:"system_role_id"` //角色编号
	// @inject_tag: db:"system_menu_id" json:"systemMenuId"
	SystemMenuId *int64 `protobuf:"varint,3,opt,name=system_menu_id,json=systemMenuId,proto3,oneof" json:"systemMenuId" db:"system_menu_id"` //菜单编号
	// @inject_tag: db:"creator" json:"creator"
	Creator *string `protobuf:"bytes,4,opt,name=creator,proto3,oneof" json:"creator" db:"creator"` //创建者
	// @inject_tag: db:"create_time" json:"createTime"
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"createTime" db:"create_time"` //创建时间
	// @inject_tag: db:"updater" json:"updater"
	Updater *string `protobuf:"bytes,6,opt,name=updater,proto3,oneof" json:"updater" db:"updater"` //更新者
	// @inject_tag: db:"update_time" json:"updateTime"
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=update_time,json=updateTime,proto3" json:"updateTime" db:"update_time"` //更新时间
}

func (x *SystemRoleMenuObject) Reset() {
	*x = SystemRoleMenuObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_system_role_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemRoleMenuObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemRoleMenuObject) ProtoMessage() {}

func (x *SystemRoleMenuObject) ProtoReflect() protoreflect.Message {
	mi := &file_role_system_role_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemRoleMenuObject.ProtoReflect.Descriptor instead.
func (*SystemRoleMenuObject) Descriptor() ([]byte, []int) {
	return file_role_system_role_menu_proto_rawDescGZIP(), []int{0}
}

func (x *SystemRoleMenuObject) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *SystemRoleMenuObject) GetSystemRoleId() int64 {
	if x != nil && x.SystemRoleId != nil {
		return *x.SystemRoleId
	}
	return 0
}

func (x *SystemRoleMenuObject) GetSystemMenuId() int64 {
	if x != nil && x.SystemMenuId != nil {
		return *x.SystemMenuId
	}
	return 0
}

func (x *SystemRoleMenuObject) GetCreator() string {
	if x != nil && x.Creator != nil {
		return *x.Creator
	}
	return ""
}

func (x *SystemRoleMenuObject) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *SystemRoleMenuObject) GetUpdater() string {
	if x != nil && x.Updater != nil {
		return *x.Updater
	}
	return ""
}

func (x *SystemRoleMenuObject) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// SystemRoleMenuCreateRequest 创建数据请求
type SystemRoleMenuCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"system_role_id" json:"systemRoleId"
	SystemRoleId *int64 `protobuf:"varint,1,opt,name=system_role_id,json=systemRoleId,proto3,oneof" json:"systemRoleId" db:"system_role_id"` //角色ID
	// @inject_tag: db:"system_menu_ids" json:"systemMenuIds"
	SystemMenuIds []byte `protobuf:"bytes,2,opt,name=system_menu_ids,json=systemMenuIds,proto3,oneof" json:"systemMenuIds" db:"system_menu_ids"` //菜单ID
	// @inject_tag: db:"creator" json:"creator"
	Creator *string `protobuf:"bytes,3,opt,name=creator,proto3,oneof" json:"creator" db:"creator"` //创建者
	// @inject_tag: db:"create_time" json:"createTime"
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"createTime" db:"create_time"` //创建时间
	// @inject_tag: db:"updater" json:"updater"
	Updater *string `protobuf:"bytes,5,opt,name=updater,proto3,oneof" json:"updater" db:"updater"` //更新者
	// @inject_tag: db:"update_time" json:"updateTime"
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"updateTime" db:"update_time"` //更新时间
}

func (x *SystemRoleMenuCreateRequest) Reset() {
	*x = SystemRoleMenuCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_system_role_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemRoleMenuCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemRoleMenuCreateRequest) ProtoMessage() {}

func (x *SystemRoleMenuCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_system_role_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemRoleMenuCreateRequest.ProtoReflect.Descriptor instead.
func (*SystemRoleMenuCreateRequest) Descriptor() ([]byte, []int) {
	return file_role_system_role_menu_proto_rawDescGZIP(), []int{1}
}

func (x *SystemRoleMenuCreateRequest) GetSystemRoleId() int64 {
	if x != nil && x.SystemRoleId != nil {
		return *x.SystemRoleId
	}
	return 0
}

func (x *SystemRoleMenuCreateRequest) GetSystemMenuIds() []byte {
	if x != nil {
		return x.SystemMenuIds
	}
	return nil
}

func (x *SystemRoleMenuCreateRequest) GetCreator() string {
	if x != nil && x.Creator != nil {
		return *x.Creator
	}
	return ""
}

func (x *SystemRoleMenuCreateRequest) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *SystemRoleMenuCreateRequest) GetUpdater() string {
	if x != nil && x.Updater != nil {
		return *x.Updater
	}
	return ""
}

func (x *SystemRoleMenuCreateRequest) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// SystemRoleMenuCreateResponse 创建数据响应
type SystemRoleMenuCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data int64  `protobuf:"varint,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemRoleMenuCreateResponse) Reset() {
	*x = SystemRoleMenuCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_system_role_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemRoleMenuCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemRoleMenuCreateResponse) ProtoMessage() {}

func (x *SystemRoleMenuCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_system_role_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemRoleMenuCreateResponse.ProtoReflect.Descriptor instead.
func (*SystemRoleMenuCreateResponse) Descriptor() ([]byte, []int) {
	return file_role_system_role_menu_proto_rawDescGZIP(), []int{2}
}

func (x *SystemRoleMenuCreateResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SystemRoleMenuCreateResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SystemRoleMenuCreateResponse) GetData() int64 {
	if x != nil {
		return x.Data
	}
	return 0
}

// SystemRoleMenuListRequest 列表数据
type SystemRoleMenuListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"system_menu_id" json:"systemMenuId"
	SystemMenuId *int64 `protobuf:"varint,1,opt,name=system_menu_id,json=systemMenuId,proto3,oneof" json:"systemMenuId" db:"system_menu_id"` // 菜单编号
	// @inject_tag: db:"system_role_id" json:"systemRoleId"
	SystemRoleId *int64 `protobuf:"varint,2,opt,name=system_role_id,json=systemRoleId,proto3,oneof" json:"systemRoleId" db:"system_role_id"` // 角色编号
}

func (x *SystemRoleMenuListRequest) Reset() {
	*x = SystemRoleMenuListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_system_role_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemRoleMenuListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemRoleMenuListRequest) ProtoMessage() {}

func (x *SystemRoleMenuListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_system_role_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemRoleMenuListRequest.ProtoReflect.Descriptor instead.
func (*SystemRoleMenuListRequest) Descriptor() ([]byte, []int) {
	return file_role_system_role_menu_proto_rawDescGZIP(), []int{3}
}

func (x *SystemRoleMenuListRequest) GetSystemMenuId() int64 {
	if x != nil && x.SystemMenuId != nil {
		return *x.SystemMenuId
	}
	return 0
}

func (x *SystemRoleMenuListRequest) GetSystemRoleId() int64 {
	if x != nil && x.SystemRoleId != nil {
		return *x.SystemRoleId
	}
	return 0
}

// SystemRoleMenuListResponse 数据响应
type SystemRoleMenuListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64                   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []*SystemRoleMenuObject `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemRoleMenuListResponse) Reset() {
	*x = SystemRoleMenuListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_system_role_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemRoleMenuListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemRoleMenuListResponse) ProtoMessage() {}

func (x *SystemRoleMenuListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_system_role_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemRoleMenuListResponse.ProtoReflect.Descriptor instead.
func (*SystemRoleMenuListResponse) Descriptor() ([]byte, []int) {
	return file_role_system_role_menu_proto_rawDescGZIP(), []int{4}
}

func (x *SystemRoleMenuListResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SystemRoleMenuListResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SystemRoleMenuListResponse) GetData() []*SystemRoleMenuObject {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_role_system_role_menu_proto protoreflect.FileDescriptor

var file_role_system_role_menu_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x14, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x13, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0c, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a,
	0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4d,
	0x65, 0x6e, 0x75, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x72, 0x22, 0xec, 0x02, 0x0a, 0x1b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x2b, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x0d, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x73, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x72, 0x22, 0x58, 0x0a, 0x1c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f,
	0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x97,
	0x01, 0x0a, 0x19, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e,
	0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4d, 0x65,
	0x6e, 0x75, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x01, 0x52, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x65,
	0x6e, 0x75, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x72, 0x0a, 0x1a, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xcf, 0x01, 0x0a,
	0x15, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x14, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65,
	0x4d, 0x65, 0x6e, 0x75, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e,
	0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65,
	0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_system_role_menu_proto_rawDescOnce sync.Once
	file_role_system_role_menu_proto_rawDescData = file_role_system_role_menu_proto_rawDesc
)

func file_role_system_role_menu_proto_rawDescGZIP() []byte {
	file_role_system_role_menu_proto_rawDescOnce.Do(func() {
		file_role_system_role_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_system_role_menu_proto_rawDescData)
	})
	return file_role_system_role_menu_proto_rawDescData
}

var file_role_system_role_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_role_system_role_menu_proto_goTypes = []interface{}{
	(*SystemRoleMenuObject)(nil),         // 0: role.SystemRoleMenuObject
	(*SystemRoleMenuCreateRequest)(nil),  // 1: role.SystemRoleMenuCreateRequest
	(*SystemRoleMenuCreateResponse)(nil), // 2: role.SystemRoleMenuCreateResponse
	(*SystemRoleMenuListRequest)(nil),    // 3: role.SystemRoleMenuListRequest
	(*SystemRoleMenuListResponse)(nil),   // 4: role.SystemRoleMenuListResponse
	(*timestamppb.Timestamp)(nil),        // 5: google.protobuf.Timestamp
}
var file_role_system_role_menu_proto_depIdxs = []int32{
	5, // 0: role.SystemRoleMenuObject.create_time:type_name -> google.protobuf.Timestamp
	5, // 1: role.SystemRoleMenuObject.update_time:type_name -> google.protobuf.Timestamp
	5, // 2: role.SystemRoleMenuCreateRequest.create_time:type_name -> google.protobuf.Timestamp
	5, // 3: role.SystemRoleMenuCreateRequest.update_time:type_name -> google.protobuf.Timestamp
	0, // 4: role.SystemRoleMenuListResponse.data:type_name -> role.SystemRoleMenuObject
	1, // 5: role.SystemRoleMenuService.SystemRoleMenuCreate:input_type -> role.SystemRoleMenuCreateRequest
	3, // 6: role.SystemRoleMenuService.SystemRoleMenuList:input_type -> role.SystemRoleMenuListRequest
	2, // 7: role.SystemRoleMenuService.SystemRoleMenuCreate:output_type -> role.SystemRoleMenuCreateResponse
	4, // 8: role.SystemRoleMenuService.SystemRoleMenuList:output_type -> role.SystemRoleMenuListResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_role_system_role_menu_proto_init() }
func file_role_system_role_menu_proto_init() {
	if File_role_system_role_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_system_role_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemRoleMenuObject); i {
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
		file_role_system_role_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemRoleMenuCreateRequest); i {
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
		file_role_system_role_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemRoleMenuCreateResponse); i {
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
		file_role_system_role_menu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemRoleMenuListRequest); i {
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
		file_role_system_role_menu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemRoleMenuListResponse); i {
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
	file_role_system_role_menu_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_role_system_role_menu_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_role_system_role_menu_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_role_system_role_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_system_role_menu_proto_goTypes,
		DependencyIndexes: file_role_system_role_menu_proto_depIdxs,
		MessageInfos:      file_role_system_role_menu_proto_msgTypes,
	}.Build()
	File_role_system_role_menu_proto = out.File
	file_role_system_role_menu_proto_rawDesc = nil
	file_role_system_role_menu_proto_goTypes = nil
	file_role_system_role_menu_proto_depIdxs = nil
}
