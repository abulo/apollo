package tenant

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_tenant_package 租户套餐包

// SystemTenantPackageDao 数据转换
func SystemTenantPackageDao(item *SystemTenantPackageObject) *dao.SystemTenantPackage {
	daoItem := &dao.SystemTenantPackage{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 套餐编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 套餐名称
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态（0正常 1停用）
	}
	if item != nil && item.SystemMenuIds != nil {
		daoItem.SystemMenuIds = null.JSONFrom(item.GetSystemMenuIds()) // 目录编号
	}
	if item != nil && item.Remark != nil {
		daoItem.Remark = null.StringFrom(item.GetRemark()) // 备注
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 是否删除(0否 1是)
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

// SystemTenantPackageProto 数据绑定
func SystemTenantPackageProto(item dao.SystemTenantPackage) *SystemTenantPackageObject {
	res := &SystemTenantPackageObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.SystemMenuIds.IsValid() {
		res.SystemMenuIds = *item.SystemMenuIds.Ptr()
	}
	if item.Remark.IsValid() {
		res.Remark = item.Remark.Ptr()
	}
	if item.Deleted != nil {
		res.Deleted = item.Deleted
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
