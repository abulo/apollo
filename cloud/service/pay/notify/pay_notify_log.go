package notify

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_notify_log 支付通知日志

// PayNotifyLogDao 数据转换
func PayNotifyLogDao(item *PayNotifyLogObject) *dao.PayNotifyLog {
	daoItem := &dao.PayNotifyLog{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 日志编号
	}
	if item != nil && item.TaskId != nil {
		daoItem.TaskId = item.TaskId // 通知任务编号
	}
	if item != nil && item.NotifyTimes != nil {
		daoItem.NotifyTimes = item.NotifyTimes // 第几次被通知
	}
	if item != nil && item.Response != nil {
		daoItem.Response = null.JSONFrom(item.GetResponse()) // 请求参数
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 通知状态
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

// PayNotifyLogProto 数据绑定
func PayNotifyLogProto(item dao.PayNotifyLog) *PayNotifyLogObject {
	res := &PayNotifyLogObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.TaskId != nil {
		res.TaskId = item.TaskId
	}
	if item.NotifyTimes != nil {
		res.NotifyTimes = item.NotifyTimes
	}
	if item.Response.IsValid() {
		res.Response = *item.Response.Ptr()
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
