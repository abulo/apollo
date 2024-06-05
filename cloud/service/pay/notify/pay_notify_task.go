package notify

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_notify_task 商户支付-任务通知

// PayNotifyTaskDao 数据转换
func PayNotifyTaskDao(item *PayNotifyTaskObject) *dao.PayNotifyTask {
	daoItem := &dao.PayNotifyTask{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 任务编号
	}
	if item != nil && item.AppId != nil {
		daoItem.AppId = item.AppId // 应用编号
	}
	if item != nil && item.Type != nil {
		daoItem.Type = item.Type // 通知类型
	}
	if item != nil && item.DataId != nil {
		daoItem.DataId = item.DataId // 数据编号
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 通知状态
	}
	if item != nil && item.MerchantOrderId != nil {
		daoItem.MerchantOrderId = item.MerchantOrderId // 商户订单编号
	}
	if item != nil && item.MerchantTransferId != nil {
		daoItem.MerchantTransferId = item.MerchantTransferId // 商户转账单编号
	}
	if item != nil && item.NextNotifyTime != nil {
		daoItem.NextNotifyTime = null.DateTimeFrom(util.GrpcTime(item.NextNotifyTime)) // 下一次通知时间
	}
	if item != nil && item.LastExecuteTime != nil {
		daoItem.LastExecuteTime = null.DateTimeFrom(util.GrpcTime(item.LastExecuteTime)) // 最后一次执行时间
	}
	if item != nil && item.NotifyTimes != nil {
		daoItem.NotifyTimes = item.NotifyTimes // 当前通知次数
	}
	if item != nil && item.MaxNotifyTimes != nil {
		daoItem.MaxNotifyTimes = null.Int32From(item.GetMaxNotifyTimes()) // 最大可通知次数
	}
	if item != nil && item.NotifyUrl != nil {
		daoItem.NotifyUrl = item.NotifyUrl // 异步通知地址
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

// PayNotifyTaskProto 数据绑定
func PayNotifyTaskProto(item dao.PayNotifyTask) *PayNotifyTaskObject {
	res := &PayNotifyTaskObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.AppId != nil {
		res.AppId = item.AppId
	}
	if item.Type != nil {
		res.Type = item.Type
	}
	if item.DataId != nil {
		res.DataId = item.DataId
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.MerchantOrderId != nil {
		res.MerchantOrderId = item.MerchantOrderId
	}
	if item.MerchantTransferId != nil {
		res.MerchantTransferId = item.MerchantTransferId
	}
	if item.NextNotifyTime.IsValid() {
		res.NextNotifyTime = timestamppb.New(*item.NextNotifyTime.Ptr())
	}
	if item.LastExecuteTime.IsValid() {
		res.LastExecuteTime = timestamppb.New(*item.LastExecuteTime.Ptr())
	}
	if item.NotifyTimes != nil {
		res.NotifyTimes = item.NotifyTimes
	}
	if item.MaxNotifyTimes.IsValid() {
		res.MaxNotifyTimes = item.MaxNotifyTimes.Ptr()
	}
	if item.NotifyUrl != nil {
		res.NotifyUrl = item.NotifyUrl
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
