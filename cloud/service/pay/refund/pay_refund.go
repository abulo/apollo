package refund

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_refund 退款订单

// PayRefundDao 数据转换
func PayRefundDao(item *PayRefundObject) *dao.PayRefund {
	daoItem := &dao.PayRefund{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 支付退款编号
	}
	if item != nil && item.No != nil {
		daoItem.No = item.No // 退款单号
	}
	if item != nil && item.AppId != nil {
		daoItem.AppId = item.AppId // 应用编号
	}
	if item != nil && item.ChannelId != nil {
		daoItem.ChannelId = item.ChannelId // 渠道编号
	}
	if item != nil && item.ChannelCode != nil {
		daoItem.ChannelCode = item.ChannelCode // 渠道编码
	}
	if item != nil && item.OrderId != nil {
		daoItem.OrderId = item.OrderId // 支付订单编号 pay_order 表id
	}
	if item != nil && item.OrderNo != nil {
		daoItem.OrderNo = item.OrderNo // 支付订单 no
	}
	if item != nil && item.MerchantOrderId != nil {
		daoItem.MerchantOrderId = item.MerchantOrderId // 商户订单编号（商户系统生成）
	}
	if item != nil && item.MerchantRefundId != nil {
		daoItem.MerchantRefundId = item.MerchantRefundId // 商户退款订单号（商户系统生成）
	}
	if item != nil && item.NotifyUrl != nil {
		daoItem.NotifyUrl = item.NotifyUrl // 异步通知商户地址
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 退款状态
	}
	if item != nil && item.PayPrice != nil {
		daoItem.PayPrice = item.PayPrice // 支付金额,单位分
	}
	if item != nil && item.RefundPrice != nil {
		daoItem.RefundPrice = item.RefundPrice // 退款金额,单位分
	}
	if item != nil && item.Reason != nil {
		daoItem.Reason = item.Reason // 退款原因
	}
	if item != nil && item.UserIp != nil {
		daoItem.UserIp = null.StringFrom(item.GetUserIp()) // ip
	}
	if item != nil && item.ChannelOrderNo != nil {
		daoItem.ChannelOrderNo = item.ChannelOrderNo // 渠道订单号，pay_order 中的 channel_order_no 对应
	}
	if item != nil && item.ChannelRefundNo != nil {
		daoItem.ChannelRefundNo = null.StringFrom(item.GetChannelRefundNo()) // 渠道退款单号，渠道返
	}
	if item != nil && item.SuccessTime != nil {
		daoItem.SuccessTime = null.DateTimeFrom(util.GrpcTime(item.SuccessTime)) // 退款成功时间
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

// PayRefundProto 数据绑定
func PayRefundProto(item dao.PayRefund) *PayRefundObject {
	res := &PayRefundObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.No != nil {
		res.No = item.No
	}
	if item.AppId != nil {
		res.AppId = item.AppId
	}
	if item.ChannelId != nil {
		res.ChannelId = item.ChannelId
	}
	if item.ChannelCode != nil {
		res.ChannelCode = item.ChannelCode
	}
	if item.OrderId != nil {
		res.OrderId = item.OrderId
	}
	if item.OrderNo != nil {
		res.OrderNo = item.OrderNo
	}
	if item.MerchantOrderId != nil {
		res.MerchantOrderId = item.MerchantOrderId
	}
	if item.MerchantRefundId != nil {
		res.MerchantRefundId = item.MerchantRefundId
	}
	if item.NotifyUrl != nil {
		res.NotifyUrl = item.NotifyUrl
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.PayPrice != nil {
		res.PayPrice = item.PayPrice
	}
	if item.RefundPrice != nil {
		res.RefundPrice = item.RefundPrice
	}
	if item.Reason != nil {
		res.Reason = item.Reason
	}
	if item.UserIp.IsValid() {
		res.UserIp = item.UserIp.Ptr()
	}
	if item.ChannelOrderNo != nil {
		res.ChannelOrderNo = item.ChannelOrderNo
	}
	if item.ChannelRefundNo.IsValid() {
		res.ChannelRefundNo = item.ChannelRefundNo.Ptr()
	}
	if item.SuccessTime.IsValid() {
		res.SuccessTime = timestamppb.New(*item.SuccessTime.Ptr())
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
