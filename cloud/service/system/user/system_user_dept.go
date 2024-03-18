package user

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_user_dept 系统用户部门

// SystemUserDeptCustomDao 数据转换
func SystemUserDeptCustomDao(item *SystemUserDeptCustomObject) *dao.SystemUserDeptCustom {
	daoItem := &dao.SystemUserDeptCustom{}

	if item != nil && item.SystemUserId != nil {
		daoItem.SystemUserId = item.SystemUserId // 系统用户 id
	}
	if item != nil && item.SystemDeptIds != nil {
		daoItem.SystemDeptIds = null.JSONFrom(item.GetSystemDeptIds()) // 部门 id
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 删除
	}
	if item != nil && item.SystemTenantId != nil {
		daoItem.SystemTenantId = item.SystemTenantId // 租户
	}
	if item != nil && item.Creator != nil {
		daoItem.Creator = null.StringFrom(item.GetCreator()) // 创建人
	}
	if item != nil && item.CreateTime != nil {
		daoItem.CreateTime = null.DateTimeFrom(util.GrpcTime(item.CreateTime)) // 创建时间
	}
	if item != nil && item.Updater != nil {
		daoItem.Updater = null.StringFrom(item.GetUpdater()) // 更新人
	}
	if item != nil && item.UpdateTime != nil {
		daoItem.UpdateTime = null.DateTimeFrom(util.GrpcTime(item.UpdateTime)) // 更新时间
	}
	return daoItem
}

// SystemUserDeptCustomProto 数据绑定
func SystemUserDeptCustomProto(item dao.SystemUserDeptCustom) *SystemUserDeptCustomObject {
	res := &SystemUserDeptCustomObject{}
	if item.SystemUserId != nil {
		res.SystemUserId = item.SystemUserId
	}
	if item.SystemDeptIds.IsValid() {
		res.SystemDeptIds = *item.SystemDeptIds.Ptr()
	}
	if item.Deleted != nil {
		res.Deleted = item.Deleted
	}
	if item.SystemTenantId != nil {
		res.SystemTenantId = item.SystemTenantId
	}
	if item.Creator.IsValid() {
		res.Creator = item.Creator.Ptr()
	}
	if item.CreateTime.IsValid() {
		res.CreateTime = timestamppb.New(*item.CreateTime.Ptr())
	}
	if item.Updater.IsValid() {
		res.Updater = item.Updater.Ptr()
	}
	if item.UpdateTime.IsValid() {
		res.UpdateTime = timestamppb.New(*item.UpdateTime.Ptr())
	}
	return res
}

// SystemUserDeptDao 数据转换
func SystemUserDeptDao(item *SystemUserDeptObject) *dao.SystemUserDept {
	daoItem := &dao.SystemUserDept{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.SystemUserId != nil {
		daoItem.SystemUserId = item.SystemUserId // 系统用户 id
	}
	if item != nil && item.SystemDeptId != nil {
		daoItem.SystemDeptId = item.SystemDeptId // 部门 id
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 删除
	}
	if item != nil && item.SystemTenantId != nil {
		daoItem.SystemTenantId = item.SystemTenantId // 租户
	}
	if item != nil && item.Creator != nil {
		daoItem.Creator = null.StringFrom(item.GetCreator()) // 创建人
	}
	if item != nil && item.CreateTime != nil {
		daoItem.CreateTime = null.DateTimeFrom(util.GrpcTime(item.CreateTime)) // 创建时间
	}
	if item != nil && item.Updater != nil {
		daoItem.Updater = null.StringFrom(item.GetUpdater()) // 更新人
	}
	if item != nil && item.UpdateTime != nil {
		daoItem.UpdateTime = null.DateTimeFrom(util.GrpcTime(item.UpdateTime)) // 更新时间
	}

	return daoItem
}

// SystemUserDeptProto 数据绑定
func SystemUserDeptProto(item dao.SystemUserDept) *SystemUserDeptObject {
	res := &SystemUserDeptObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.SystemUserId != nil {
		res.SystemUserId = item.SystemUserId
	}
	if item.SystemDeptId != nil {
		res.SystemDeptId = item.SystemDeptId
	}
	if item.Deleted != nil {
		res.Deleted = item.Deleted
	}
	if item.SystemTenantId != nil {
		res.SystemTenantId = item.SystemTenantId
	}
	if item.Creator.IsValid() {
		res.Creator = item.Creator.Ptr()
	}
	if item.CreateTime.IsValid() {
		res.CreateTime = timestamppb.New(*item.CreateTime.Ptr())
	}
	if item.Updater.IsValid() {
		res.Updater = item.Updater.Ptr()
	}
	if item.UpdateTime.IsValid() {
		res.UpdateTime = timestamppb.New(*item.UpdateTime.Ptr())
	}

	return res
}
