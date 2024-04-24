package notice

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_notice 通知公告表

// SystemNoticeDao 数据转换
func SystemNoticeDao(item *SystemNoticeObject) *dao.SystemNotice {
	daoItem := &dao.SystemNotice{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 公告ID
	}
	if item != nil && item.Title != nil {
		daoItem.Title = item.Title // 公告标题
	}
	if item != nil && item.Content != nil {
		daoItem.Content = item.Content // 公告内容
	}
	if item != nil && item.Type != nil {
		daoItem.Type = item.Type // 公告类型（1通知 2公告）
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 公告状态（0正常 1关闭）
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 删除
	}
	if item != nil && item.TenantId != nil {
		daoItem.TenantId = item.TenantId // 租户
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

// SystemNoticeProto 数据绑定
func SystemNoticeProto(item dao.SystemNotice) *SystemNoticeObject {
	res := &SystemNoticeObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Title != nil {
		res.Title = item.Title
	}
	if item.Content != nil {
		res.Content = item.Content
	}
	if item.Type != nil {
		res.Type = item.Type
	}
	if item.Status != nil {
		res.Status = item.Status
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
