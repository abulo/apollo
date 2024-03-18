package role

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_role 系统角色

func SystemRoleDataScopeDao(item *SystemRoleDataScopeObject) *dao.SystemRoleDataScope {
	daoItem := &dao.SystemRoleDataScope{}
	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 角色编号
	}
	if item != nil && item.DataScope != nil {
		daoItem.DataScope = null.Int32From(item.GetDataScope()) // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	}
	if item != nil && item.DataScopeDept != nil {
		daoItem.DataScopeDept = null.JSONFrom(item.GetDataScopeDept()) // 数据范围(指定部门数组)
	}
	if item != nil && item.Updater != nil {
		daoItem.Updater = null.StringFrom(item.GetUpdater()) // 更新者
	}
	if item != nil && item.UpdateTime != nil {
		daoItem.UpdateTime = null.DateTimeFrom(util.GrpcTime(item.UpdateTime)) // 更新时间
	}
	return daoItem
}

// SystemRoleDataScopeProto 数据绑定
func SystemRoleDataScopeProto(item dao.SystemRoleDataScope) *SystemRoleDataScopeObject {
	res := &SystemRoleDataScopeObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.DataScope.IsValid() {
		res.DataScope = item.DataScope.Ptr()
	}
	if item.DataScopeDept.IsValid() {
		res.DataScopeDept = *item.DataScopeDept.Ptr()
	}
	if item.Updater.IsValid() {
		res.Updater = item.Updater.Ptr()
	}
	if item.UpdateTime.IsValid() {
		res.UpdateTime = timestamppb.New(*item.UpdateTime.Ptr())
	}
	return res
}

// SystemRoleDao 数据转换
func SystemRoleDao(item *SystemRoleObject) *dao.SystemRole {
	daoItem := &dao.SystemRole{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 角色编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 角色名称
	}
	if item != nil && item.Code != nil {
		daoItem.Code = item.Code // 角色权限字符串
	}
	if item != nil && item.Sort != nil {
		daoItem.Sort = item.Sort // 显示顺序
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 角色状态（0正常 1停用）
	}
	if item != nil && item.Type != nil {
		daoItem.Type = item.Type // 角色类型(1内置/2定义)
	}
	if item != nil && item.Remark != nil {
		daoItem.Remark = null.StringFrom(item.GetRemark()) // 备注
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

// SystemRoleProto 数据绑定
func SystemRoleProto(item dao.SystemRole) *SystemRoleObject {
	res := &SystemRoleObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.Code != nil {
		res.Code = item.Code
	}
	if item.Sort != nil {
		res.Sort = item.Sort
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Type != nil {
		res.Type = item.Type
	}
	if item.Remark.IsValid() {
		res.Remark = item.Remark.Ptr()
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
