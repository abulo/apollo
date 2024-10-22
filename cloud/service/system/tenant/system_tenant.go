package tenant

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_tenant 租户

// SystemTenantDao 数据转换
func SystemTenantDao(item *SystemTenantObject) *dao.SystemTenant {
	daoItem := &dao.SystemTenant{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 租户编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 租户名称
	}
	if item != nil && item.UserId != nil {
		daoItem.UserId = null.Int64From(item.GetUserId()) // 联系人ID
	}
	if item != nil && item.ContactName != nil {
		daoItem.ContactName = item.ContactName // 联系人
	}
	if item != nil && item.ContactMobile != nil {
		daoItem.ContactMobile = item.ContactMobile // 租户联系电话
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态（0正常 1停用）
	}
	if item != nil && item.Domain != nil {
		daoItem.Domain = null.StringFrom(item.GetDomain()) // 域名
	}
	if item != nil && item.ExpireDate != nil {
		daoItem.ExpireDate = null.DateFrom(util.GrpcTime(item.ExpireDate)) // 过期时间
	}
	if item != nil && item.AccountCount != nil {
		daoItem.AccountCount = item.AccountCount // 账号数量
	}
	if item != nil && item.TenantPackageId != nil {
		daoItem.TenantPackageId = item.TenantPackageId // 套餐编号
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
	if item != nil && item.Username != nil {
		daoItem.Username = item.Username
	}
	if item != nil && item.Password != nil {
		daoItem.Password = item.Password
	}

	return daoItem
}

// SystemTenantSafetyDao 数据转换
func SystemTenantSafetyDao(item *SystemTenantObject) *dao.SystemTenant {
	daoItem := &dao.SystemTenant{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 租户编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 租户名称
	}
	if item != nil && item.UserId != nil {
		daoItem.UserId = null.Int64From(item.GetUserId()) // 联系人ID
	}
	if item != nil && item.ContactName != nil {
		daoItem.ContactName = item.ContactName // 联系人
	}
	if item != nil && item.ContactMobile != nil {
		daoItem.ContactMobile = item.ContactMobile // 租户联系电话
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态（0正常 1停用）
	}
	if item != nil && item.Domain != nil {
		daoItem.Domain = null.StringFrom(item.GetDomain()) // 域名
	}
	if item != nil && item.ExpireDate != nil {
		daoItem.ExpireDate = null.DateFrom(util.GrpcTime(item.ExpireDate)) // 过期时间
	}
	if item != nil && item.AccountCount != nil {
		daoItem.AccountCount = item.AccountCount // 账号数量
	}
	if item != nil && item.TenantPackageId != nil {
		daoItem.TenantPackageId = item.TenantPackageId // 套餐编号
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
	if item != nil && item.Username != nil {
		daoItem.Username = item.Username
	}
	daoItem.Password = nil

	return daoItem
}

// SystemTenantProto 数据绑定
func SystemTenantProto(item dao.SystemTenant) *SystemTenantObject {
	res := &SystemTenantObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.UserId.IsValid() {
		res.UserId = item.UserId.Ptr()
	}
	if item.ContactName != nil {
		res.ContactName = item.ContactName
	}
	if item.ContactMobile != nil {
		res.ContactMobile = item.ContactMobile
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Domain.IsValid() {
		res.Domain = item.Domain.Ptr()
	}
	if item.ExpireDate.IsValid() {
		res.ExpireDate = timestamppb.New(*item.ExpireDate.Ptr())
	}
	if item.AccountCount != nil {
		res.AccountCount = item.AccountCount
	}
	if item.TenantPackageId != nil {
		res.TenantPackageId = item.TenantPackageId
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
	if item.Username != nil {
		res.Username = item.Username
	}
	if item.Password != nil {
		res.Password = item.Password
	}

	return res
}
