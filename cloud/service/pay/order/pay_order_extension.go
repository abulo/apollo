package order

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_order_extension 支付订单拓展

// PayOrderExtensionDao 数据转换
func PayOrderExtensionDao(item *PayOrderExtensionObject) *dao.PayOrderExtension {
	daoItem := &dao.PayOrderExtension{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.No != nil {
		daoItem.No = item.No // 支付订单号
	}
	if item != nil && item.OrderId != nil {
		daoItem.OrderId = item.OrderId // 支付订单编号
	}
	if item != nil && item.ChannelId != nil {
		daoItem.ChannelId = item.ChannelId // 渠道编号
	}
	if item != nil && item.ChannelCode != nil {
		daoItem.ChannelCode = item.ChannelCode // 渠道编码
	}
	if item != nil && item.UserIp != nil {
		daoItem.UserIp = null.StringFrom(item.GetUserIp()) // ip
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 支付状态
	}
	if item != nil && item.ChannelExtras != nil {
		daoItem.ChannelExtras = null.StringFrom(item.GetChannelExtras()) // 支付渠道的额外参数
	}
	if item != nil && item.ChannelErrorCode != nil {
		daoItem.ChannelErrorCode = null.StringFrom(item.GetChannelErrorCode()) // 渠道调用报错时，错误码
	}
	if item != nil && item.ChannelErrorMsg != nil {
		daoItem.ChannelErrorMsg = null.StringFrom(item.GetChannelErrorMsg()) // 渠道调用报错时，错误信息
	}
	if item != nil && item.ChannelNotifyData != nil {
		daoItem.ChannelNotifyData = null.StringFrom(item.GetChannelNotifyData()) // 支付渠道异步通知的内容
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

// PayOrderExtensionProto 数据绑定
func PayOrderExtensionProto(item dao.PayOrderExtension) *PayOrderExtensionObject {
	res := &PayOrderExtensionObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.No != nil {
		res.No = item.No
	}
	if item.OrderId != nil {
		res.OrderId = item.OrderId
	}
	if item.ChannelId != nil {
		res.ChannelId = item.ChannelId
	}
	if item.ChannelCode != nil {
		res.ChannelCode = item.ChannelCode
	}
	if item.UserIp.IsValid() {
		res.UserIp = item.UserIp.Ptr()
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.ChannelExtras.IsValid() {
		res.ChannelExtras = item.ChannelExtras.Ptr()
	}
	if item.ChannelErrorCode.IsValid() {
		res.ChannelErrorCode = item.ChannelErrorCode.Ptr()
	}
	if item.ChannelErrorMsg.IsValid() {
		res.ChannelErrorMsg = item.ChannelErrorMsg.Ptr()
	}
	if item.ChannelNotifyData.IsValid() {
		res.ChannelNotifyData = item.ChannelNotifyData.Ptr()
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
