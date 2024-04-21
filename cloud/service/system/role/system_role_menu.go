package role

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_role_menu 系统角色和系统菜单关联表

func SystemRoleMenuCustomDao(item *SystemRoleMenuCustomObject) *dao.SystemRoleMenuCustom {
	daoItem := &dao.SystemRoleMenuCustom{}
	if item != nil && item.RoleId != nil {
		daoItem.RoleId = item.RoleId // 角色编号
	}
	if item != nil && item.MenuIds != nil {
		daoItem.MenuIds = null.JSONFrom(item.GetMenuIds()) // 数据范围(指定部门数组)
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 删除
	}
	if item != nil && item.TenantId != nil {
		daoItem.TenantId = item.TenantId // 租户
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

// SystemRoleMenuCustomProto 数据绑定
func SystemRoleMenuCustomProto(item dao.SystemRoleMenuCustom) *SystemRoleMenuCustomObject {
	res := &SystemRoleMenuCustomObject{}
	if item.RoleId != nil {
		res.RoleId = item.RoleId
	}
	if item.MenuIds.IsValid() {
		res.MenuIds = *item.MenuIds.Ptr()
	}
	if item.Deleted != nil {
		res.Deleted = item.Deleted
	}
	if item.TenantId != nil {
		res.TenantId = item.TenantId
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

// SystemRoleMenuDao 数据转换
func SystemRoleMenuDao(item *SystemRoleMenuObject) *dao.SystemRoleMenu {
	daoItem := &dao.SystemRoleMenu{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 自增编号
	}
	if item != nil && item.RoleId != nil {
		daoItem.RoleId = item.RoleId // 角色编号
	}
	if item != nil && item.MenuId != nil {
		daoItem.MenuId = item.MenuId // 菜单编号
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 删除
	}
	if item != nil && item.TenantId != nil {
		daoItem.TenantId = item.TenantId // 租户
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

// SystemRoleMenuProto 数据绑定
func SystemRoleMenuProto(item dao.SystemRoleMenu) *SystemRoleMenuObject {
	res := &SystemRoleMenuObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.RoleId != nil {
		res.RoleId = item.RoleId
	}
	if item.MenuId != nil {
		res.MenuId = item.MenuId
	}
	if item.Deleted != nil {
		res.Deleted = item.Deleted
	}
	if item.TenantId != nil {
		res.TenantId = item.TenantId
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
