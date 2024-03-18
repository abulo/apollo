package user

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_user_post 系统用户职位

// SystemUserPostCustomDao 数据转换
func SystemUserPostCustomDao(item *SystemUserPostCustomObject) *dao.SystemUserPostCustom {
	daoItem := &dao.SystemUserPostCustom{}

	if item != nil && item.SystemUserId != nil {
		daoItem.SystemUserId = item.SystemUserId // 系统用户 id
	}
	if item != nil && item.SystemPostIds != nil {
		daoItem.SystemPostIds = null.JSONFrom(item.GetSystemPostIds()) // 职位 id
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

// SystemUserPostCustomProto 数据绑定
func SystemUserPostCustomProto(item dao.SystemUserPostCustom) *SystemUserPostCustomObject {
	res := &SystemUserPostCustomObject{}
	if item.SystemUserId != nil {
		res.SystemUserId = item.SystemUserId
	}
	if item.SystemPostIds.IsValid() {
		res.SystemPostIds = *item.SystemPostIds.Ptr()
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

// SystemUserPostDao 数据转换
func SystemUserPostDao(item *SystemUserPostObject) *dao.SystemUserPost {
	daoItem := &dao.SystemUserPost{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.SystemUserId != nil {
		daoItem.SystemUserId = item.SystemUserId // 系统用户 ID
	}
	if item != nil && item.SystemPostId != nil {
		daoItem.SystemPostId = item.SystemPostId // 职位 id
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

// SystemUserPostProto 数据绑定
func SystemUserPostProto(item dao.SystemUserPost) *SystemUserPostObject {
	res := &SystemUserPostObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.SystemUserId != nil {
		res.SystemUserId = item.SystemUserId
	}
	if item.SystemPostId != nil {
		res.SystemPostId = item.SystemPostId
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
