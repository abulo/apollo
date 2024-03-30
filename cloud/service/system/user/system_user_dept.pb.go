// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: user/system_user_dept.proto

// system_user_dept 系统用户部门

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

// SystemUserDeptObject 数据对象
type SystemUserDeptObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"id" json:"id"
	Id *int64 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id" db:"id"` //编号
	// @inject_tag: db:"system_user_id" json:"systemUserId"
	SystemUserId *int64 `protobuf:"varint,2,opt,name=system_user_id,json=systemUserId,proto3,oneof" json:"systemUserId" db:"system_user_id"` //系统用户 id
	// @inject_tag: db:"system_dept_id" json:"systemDeptId"
	SystemDeptId *int64 `protobuf:"varint,3,opt,name=system_dept_id,json=systemDeptId,proto3,oneof" json:"systemDeptId" db:"system_dept_id"` //部门 id
	// @inject_tag: db:"deleted" json:"deleted"
	Deleted *int32 `protobuf:"varint,4,opt,name=deleted,proto3,oneof" json:"deleted" db:"deleted"` //删除
	// @inject_tag: db:"system_tenant_id" json:"systemTenantId"
	SystemTenantId *int64 `protobuf:"varint,5,opt,name=system_tenant_id,json=systemTenantId,proto3,oneof" json:"systemTenantId" db:"system_tenant_id"` //租户
	// @inject_tag: db:"creator" json:"creator"
	Creator *string `protobuf:"bytes,6,opt,name=creator,proto3,oneof" json:"creator" db:"creator"` //创建人
	// @inject_tag: db:"create_time" json:"createTime"
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"createTime" db:"create_time"` //创建时间
	// @inject_tag: db:"updater" json:"updater"
	Updater *string `protobuf:"bytes,8,opt,name=updater,proto3,oneof" json:"updater" db:"updater"` //更新人
	// @inject_tag: db:"update_time" json:"updateTime"
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"updateTime" db:"update_time"` //更新时间
}

func (x *SystemUserDeptObject) Reset() {
	*x = SystemUserDeptObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_system_user_dept_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserDeptObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserDeptObject) ProtoMessage() {}

func (x *SystemUserDeptObject) ProtoReflect() protoreflect.Message {
	mi := &file_user_system_user_dept_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserDeptObject.ProtoReflect.Descriptor instead.
func (*SystemUserDeptObject) Descriptor() ([]byte, []int) {
	return file_user_system_user_dept_proto_rawDescGZIP(), []int{0}
}

func (x *SystemUserDeptObject) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *SystemUserDeptObject) GetSystemUserId() int64 {
	if x != nil && x.SystemUserId != nil {
		return *x.SystemUserId
	}
	return 0
}

func (x *SystemUserDeptObject) GetSystemDeptId() int64 {
	if x != nil && x.SystemDeptId != nil {
		return *x.SystemDeptId
	}
	return 0
}

func (x *SystemUserDeptObject) GetDeleted() int32 {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return 0
}

func (x *SystemUserDeptObject) GetSystemTenantId() int64 {
	if x != nil && x.SystemTenantId != nil {
		return *x.SystemTenantId
	}
	return 0
}

func (x *SystemUserDeptObject) GetCreator() string {
	if x != nil && x.Creator != nil {
		return *x.Creator
	}
	return ""
}

func (x *SystemUserDeptObject) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *SystemUserDeptObject) GetUpdater() string {
	if x != nil && x.Updater != nil {
		return *x.Updater
	}
	return ""
}

func (x *SystemUserDeptObject) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// SystemUserDeptCustomObject 数据对象
type SystemUserDeptCustomObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"system_user_id" json:"systemUserId"
	SystemUserId *int64 `protobuf:"varint,1,opt,name=system_user_id,json=systemUserId,proto3,oneof" json:"systemUserId" db:"system_user_id"` //系统用户 id
	// @inject_tag: db:"system_dept_ids" json:"systemDeptIds"
	SystemDeptIds []byte `protobuf:"bytes,2,opt,name=system_dept_ids,json=systemDeptIds,proto3,oneof" json:"systemDeptIds" db:"system_dept_ids"` //部门 id
	// @inject_tag: db:"deleted" json:"deleted"
	Deleted *int32 `protobuf:"varint,3,opt,name=deleted,proto3,oneof" json:"deleted" db:"deleted"` //删除
	// @inject_tag: db:"system_tenant_id" json:"systemTenantId"
	SystemTenantId *int64 `protobuf:"varint,4,opt,name=system_tenant_id,json=systemTenantId,proto3,oneof" json:"systemTenantId" db:"system_tenant_id"` //租户
	// @inject_tag: db:"creator" json:"creator"
	Creator *string `protobuf:"bytes,5,opt,name=creator,proto3,oneof" json:"creator" db:"creator"` //创建人
	// @inject_tag: db:"create_time" json:"createTime"
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"createTime" db:"create_time"` //创建时间
	// @inject_tag: db:"updater" json:"updater"
	Updater *string `protobuf:"bytes,7,opt,name=updater,proto3,oneof" json:"updater" db:"updater"` //更新人
	// @inject_tag: db:"update_time" json:"updateTime"
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"updateTime" db:"update_time"` //更新时间
}

func (x *SystemUserDeptCustomObject) Reset() {
	*x = SystemUserDeptCustomObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_system_user_dept_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserDeptCustomObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserDeptCustomObject) ProtoMessage() {}

func (x *SystemUserDeptCustomObject) ProtoReflect() protoreflect.Message {
	mi := &file_user_system_user_dept_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserDeptCustomObject.ProtoReflect.Descriptor instead.
func (*SystemUserDeptCustomObject) Descriptor() ([]byte, []int) {
	return file_user_system_user_dept_proto_rawDescGZIP(), []int{1}
}

func (x *SystemUserDeptCustomObject) GetSystemUserId() int64 {
	if x != nil && x.SystemUserId != nil {
		return *x.SystemUserId
	}
	return 0
}

func (x *SystemUserDeptCustomObject) GetSystemDeptIds() []byte {
	if x != nil {
		return x.SystemDeptIds
	}
	return nil
}

func (x *SystemUserDeptCustomObject) GetDeleted() int32 {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return 0
}

func (x *SystemUserDeptCustomObject) GetSystemTenantId() int64 {
	if x != nil && x.SystemTenantId != nil {
		return *x.SystemTenantId
	}
	return 0
}

func (x *SystemUserDeptCustomObject) GetCreator() string {
	if x != nil && x.Creator != nil {
		return *x.Creator
	}
	return ""
}

func (x *SystemUserDeptCustomObject) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *SystemUserDeptCustomObject) GetUpdater() string {
	if x != nil && x.Updater != nil {
		return *x.Updater
	}
	return ""
}

func (x *SystemUserDeptCustomObject) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// SystemUserDeptCreateRequest 创建数据请求
type SystemUserDeptCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SystemUserDeptCustomObject `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemUserDeptCreateRequest) Reset() {
	*x = SystemUserDeptCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_system_user_dept_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserDeptCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserDeptCreateRequest) ProtoMessage() {}

func (x *SystemUserDeptCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_system_user_dept_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserDeptCreateRequest.ProtoReflect.Descriptor instead.
func (*SystemUserDeptCreateRequest) Descriptor() ([]byte, []int) {
	return file_user_system_user_dept_proto_rawDescGZIP(), []int{2}
}

func (x *SystemUserDeptCreateRequest) GetData() *SystemUserDeptCustomObject {
	if x != nil {
		return x.Data
	}
	return nil
}

// SystemUserDeptCreateResponse 创建数据响应
type SystemUserDeptCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data int64  `protobuf:"varint,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemUserDeptCreateResponse) Reset() {
	*x = SystemUserDeptCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_system_user_dept_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserDeptCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserDeptCreateResponse) ProtoMessage() {}

func (x *SystemUserDeptCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_system_user_dept_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserDeptCreateResponse.ProtoReflect.Descriptor instead.
func (*SystemUserDeptCreateResponse) Descriptor() ([]byte, []int) {
	return file_user_system_user_dept_proto_rawDescGZIP(), []int{3}
}

func (x *SystemUserDeptCreateResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SystemUserDeptCreateResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SystemUserDeptCreateResponse) GetData() int64 {
	if x != nil {
		return x.Data
	}
	return 0
}

// SystemUserDeptListRequest 列表数据
type SystemUserDeptListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"system_tenant_id" json:"systemTenantId"
	SystemTenantId *int64 `protobuf:"varint,1,opt,name=system_tenant_id,json=systemTenantId,proto3,oneof" json:"systemTenantId" db:"system_tenant_id"` // 租户
	// @inject_tag: db:"deleted" json:"deleted"
	Deleted *int32 `protobuf:"varint,2,opt,name=deleted,proto3,oneof" json:"deleted" db:"deleted"` // 删除
	// @inject_tag: db:"system_dept_id" json:"systemDeptId"
	SystemDeptId *int64 `protobuf:"varint,3,opt,name=system_dept_id,json=systemDeptId,proto3,oneof" json:"systemDeptId" db:"system_dept_id"` // 部门 id
	// @inject_tag: db:"system_user_id" json:"systemUserId"
	SystemUserId *int64 `protobuf:"varint,4,opt,name=system_user_id,json=systemUserId,proto3,oneof" json:"systemUserId" db:"system_user_id"` // 系统用户 id
}

func (x *SystemUserDeptListRequest) Reset() {
	*x = SystemUserDeptListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_system_user_dept_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserDeptListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserDeptListRequest) ProtoMessage() {}

func (x *SystemUserDeptListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_system_user_dept_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserDeptListRequest.ProtoReflect.Descriptor instead.
func (*SystemUserDeptListRequest) Descriptor() ([]byte, []int) {
	return file_user_system_user_dept_proto_rawDescGZIP(), []int{4}
}

func (x *SystemUserDeptListRequest) GetSystemTenantId() int64 {
	if x != nil && x.SystemTenantId != nil {
		return *x.SystemTenantId
	}
	return 0
}

func (x *SystemUserDeptListRequest) GetDeleted() int32 {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return 0
}

func (x *SystemUserDeptListRequest) GetSystemDeptId() int64 {
	if x != nil && x.SystemDeptId != nil {
		return *x.SystemDeptId
	}
	return 0
}

func (x *SystemUserDeptListRequest) GetSystemUserId() int64 {
	if x != nil && x.SystemUserId != nil {
		return *x.SystemUserId
	}
	return 0
}

// SystemUserDeptListResponse 数据响应
type SystemUserDeptListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64                   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []*SystemUserDeptObject `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemUserDeptListResponse) Reset() {
	*x = SystemUserDeptListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_system_user_dept_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserDeptListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserDeptListResponse) ProtoMessage() {}

func (x *SystemUserDeptListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_system_user_dept_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserDeptListResponse.ProtoReflect.Descriptor instead.
func (*SystemUserDeptListResponse) Descriptor() ([]byte, []int) {
	return file_user_system_user_dept_proto_rawDescGZIP(), []int{5}
}

func (x *SystemUserDeptListResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SystemUserDeptListResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SystemUserDeptListResponse) GetData() []*SystemUserDeptObject {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_user_system_user_dept_proto protoreflect.FileDescriptor

var file_user_system_user_dept_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x03, 0x0a, 0x14, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x70, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x13, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0c, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a,
	0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44,
	0x65, 0x70, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x04, 0x52, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x88, 0x01,
	0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x72, 0x22, 0xda, 0x03, 0x0a, 0x1a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x70, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x29, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0c, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2b,
	0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x44, 0x65, 0x70, 0x74, 0x49, 0x64, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x64, 0x65, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x72, 0x22, 0x53, 0x0a, 0x1b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x70, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x70, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x58, 0x0a, 0x1c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x70, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x86, 0x02, 0x0a, 0x19, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x70, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x10, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01,
	0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x65,
	0x70, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x03, 0x52, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x64,
	0x65, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x72, 0x0a, 0x1a, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x70, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2e, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x70,
	0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xcf, 0x01,
	0x0a, 0x15, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x70, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x14, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x70, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x70, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x70, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x70, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x70, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x70, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_user_system_user_dept_proto_rawDescOnce sync.Once
	file_user_system_user_dept_proto_rawDescData = file_user_system_user_dept_proto_rawDesc
)

func file_user_system_user_dept_proto_rawDescGZIP() []byte {
	file_user_system_user_dept_proto_rawDescOnce.Do(func() {
		file_user_system_user_dept_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_system_user_dept_proto_rawDescData)
	})
	return file_user_system_user_dept_proto_rawDescData
}

var file_user_system_user_dept_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_system_user_dept_proto_goTypes = []interface{}{
	(*SystemUserDeptObject)(nil),         // 0: user.SystemUserDeptObject
	(*SystemUserDeptCustomObject)(nil),   // 1: user.SystemUserDeptCustomObject
	(*SystemUserDeptCreateRequest)(nil),  // 2: user.SystemUserDeptCreateRequest
	(*SystemUserDeptCreateResponse)(nil), // 3: user.SystemUserDeptCreateResponse
	(*SystemUserDeptListRequest)(nil),    // 4: user.SystemUserDeptListRequest
	(*SystemUserDeptListResponse)(nil),   // 5: user.SystemUserDeptListResponse
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
}
var file_user_system_user_dept_proto_depIdxs = []int32{
	6, // 0: user.SystemUserDeptObject.create_time:type_name -> google.protobuf.Timestamp
	6, // 1: user.SystemUserDeptObject.update_time:type_name -> google.protobuf.Timestamp
	6, // 2: user.SystemUserDeptCustomObject.create_time:type_name -> google.protobuf.Timestamp
	6, // 3: user.SystemUserDeptCustomObject.update_time:type_name -> google.protobuf.Timestamp
	1, // 4: user.SystemUserDeptCreateRequest.data:type_name -> user.SystemUserDeptCustomObject
	0, // 5: user.SystemUserDeptListResponse.data:type_name -> user.SystemUserDeptObject
	2, // 6: user.SystemUserDeptService.SystemUserDeptCreate:input_type -> user.SystemUserDeptCreateRequest
	4, // 7: user.SystemUserDeptService.SystemUserDeptList:input_type -> user.SystemUserDeptListRequest
	3, // 8: user.SystemUserDeptService.SystemUserDeptCreate:output_type -> user.SystemUserDeptCreateResponse
	5, // 9: user.SystemUserDeptService.SystemUserDeptList:output_type -> user.SystemUserDeptListResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_user_system_user_dept_proto_init() }
func file_user_system_user_dept_proto_init() {
	if File_user_system_user_dept_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_system_user_dept_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemUserDeptObject); i {
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
		file_user_system_user_dept_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemUserDeptCustomObject); i {
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
		file_user_system_user_dept_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemUserDeptCreateRequest); i {
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
		file_user_system_user_dept_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemUserDeptCreateResponse); i {
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
		file_user_system_user_dept_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemUserDeptListRequest); i {
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
		file_user_system_user_dept_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemUserDeptListResponse); i {
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
	file_user_system_user_dept_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_user_system_user_dept_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_user_system_user_dept_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_system_user_dept_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_system_user_dept_proto_goTypes,
		DependencyIndexes: file_user_system_user_dept_proto_depIdxs,
		MessageInfos:      file_user_system_user_dept_proto_msgTypes,
	}.Build()
	File_user_system_user_dept_proto = out.File
	file_user_system_user_dept_proto_rawDesc = nil
	file_user_system_user_dept_proto_goTypes = nil
	file_user_system_user_dept_proto_depIdxs = nil
}