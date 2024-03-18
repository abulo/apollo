package user

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_user_role 系统用户和系统角色关联表

// SystemUserRoleCustomDao 数据转换
func SystemUserRoleCustomDao(item *SystemUserRoleCustomObject) *dao.SystemUserRoleCustom {
	daoItem := &dao.SystemUserRoleCustom{}

	if item != nil && item.SystemUserId != nil {
		daoItem.SystemUserId = item.SystemUserId // 系统用户 id
	}
	if item != nil && item.SystemRoleIds != nil {
		daoItem.SystemRoleIds = null.JSONFrom(item.GetSystemRoleIds()) // 角色 id
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

// SystemUserRoleCustomProto 数据绑定
func SystemUserRoleCustomProto(item dao.SystemUserRoleCustom) *SystemUserRoleCustomObject {
	res := &SystemUserRoleCustomObject{}
	if item.SystemUserId != nil {
		res.SystemUserId = item.SystemUserId
	}
	if item.SystemRoleIds.IsValid() {
		res.SystemRoleIds = *item.SystemRoleIds.Ptr()
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

// SystemUserRoleDao 数据转换
func SystemUserRoleDao(item *SystemUserRoleObject) *dao.SystemUserRole {
	daoItem := &dao.SystemUserRole{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 自增编号
	}
	if item != nil && item.SystemUserId != nil {
		daoItem.SystemUserId = item.SystemUserId // 用户编号
	}
	if item != nil && item.SystemRoleId != nil {
		daoItem.SystemRoleId = item.SystemRoleId // 角色编号
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 删除
	}
	if item != nil && item.SystemTenantId != nil {
		daoItem.SystemTenantId = item.SystemTenantId // 租户
	}
	if item != nil && item.Creator != nil {
		daoItem.Creator = null.StringFrom(item.GetCreator()) // 创建者
	}
	if item != nil && item.CreateTime != nil {
		daoItem.CreateTime = null.DateTimeFrom(util.GrpcTime(item.CreateTime)) // 创建时间
	}
	if item != nil && item.Updater != nil {
		daoItem.Updater = null.StringFrom(item.GetUpdater()) // 更新者
	}
	if item != nil && item.UpdateTime != nil {
		daoItem.UpdateTime = null.DateTimeFrom(util.GrpcTime(item.UpdateTime)) // 更新时间
	}

	return daoItem
}

// SystemUserRoleProto 数据绑定
func SystemUserRoleProto(item dao.SystemUserRole) *SystemUserRoleObject {
	res := &SystemUserRoleObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.SystemUserId != nil {
		res.SystemUserId = item.SystemUserId
	}
	if item.SystemRoleId != nil {
		res.SystemRoleId = item.SystemRoleId
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
