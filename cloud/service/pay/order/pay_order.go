package order

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_order 支付订单

// PayOrderDao 数据转换
func PayOrderDao(item *PayOrderObject) *dao.PayOrder {
	daoItem := &dao.PayOrder{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 支付订单编号
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
	if item != nil && item.MerchantOrderId != nil {
		daoItem.MerchantOrderId = item.MerchantOrderId // 商户订单编号
	}
	if item != nil && item.Subject != nil {
		daoItem.Subject = item.Subject // 商品标题
	}
	if item != nil && item.Body != nil {
		daoItem.Body = item.Body // 商品描述
	}
	if item != nil && item.NotifyUrl != nil {
		daoItem.NotifyUrl = item.NotifyUrl // 异步通知地址
	}
	if item != nil && item.Price != nil {
		daoItem.Price = item.Price // 支付金额，单位：分
	}
	if item != nil && item.ChannelFeeRate != nil {
		daoItem.ChannelFeeRate = item.ChannelFeeRate // 渠道手续费，单位：百分比
	}
	if item != nil && item.ChannelFeePrice != nil {
		daoItem.ChannelFeePrice = item.ChannelFeePrice // 渠道手续金额，单位：分
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 支付状态
	}
	if item != nil && item.UserIp != nil {
		daoItem.UserIp = null.StringFrom(item.GetUserIp()) // 用户 IP
	}
	if item != nil && item.ExpireTime != nil {
		daoItem.ExpireTime = null.DateTimeFrom(util.GrpcTime(item.ExpireTime)) // 订单失效时间
	}
	if item != nil && item.SuccessTime != nil {
		daoItem.SuccessTime = null.DateTimeFrom(util.GrpcTime(item.SuccessTime)) // 订单支付成功时间
	}
	if item != nil && item.ExtensionId != nil {
		daoItem.ExtensionId = null.Int64From(item.GetExtensionId()) // 支付成功的订单拓展单编号
	}
	if item != nil && item.No != nil {
		daoItem.No = null.StringFrom(item.GetNo()) // 支付订单号
	}
	if item != nil && item.RefundPrice != nil {
		daoItem.RefundPrice = item.RefundPrice // 退款总金额，单位：分
	}
	if item != nil && item.ChannelUserId != nil {
		daoItem.ChannelUserId = null.StringFrom(item.GetChannelUserId()) // 渠道用户编号
	}
	if item != nil && item.ChannelOrderNo != nil {
		daoItem.ChannelOrderNo = null.StringFrom(item.GetChannelOrderNo()) // 渠道订单号
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

// PayOrderProto 数据绑定
func PayOrderProto(item dao.PayOrder) *PayOrderObject {
	res := &PayOrderObject{}
	if item.Id != nil {
		res.Id = item.Id
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
	if item.MerchantOrderId != nil {
		res.MerchantOrderId = item.MerchantOrderId
	}
	if item.Subject != nil {
		res.Subject = item.Subject
	}
	if item.Body != nil {
		res.Body = item.Body
	}
	if item.NotifyUrl != nil {
		res.NotifyUrl = item.NotifyUrl
	}
	if item.Price != nil {
		res.Price = item.Price
	}
	if item.ChannelFeeRate != nil {
		res.ChannelFeeRate = item.ChannelFeeRate
	}
	if item.ChannelFeePrice != nil {
		res.ChannelFeePrice = item.ChannelFeePrice
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.UserIp.IsValid() {
		res.UserIp = item.UserIp.Ptr()
	}
	if item.ExpireTime.IsValid() {
		res.ExpireTime = timestamppb.New(*item.ExpireTime.Ptr())
	}
	if item.SuccessTime.IsValid() {
		res.SuccessTime = timestamppb.New(*item.SuccessTime.Ptr())
	}
	if item.ExtensionId.IsValid() {
		res.ExtensionId = item.ExtensionId.Ptr()
	}
	if item.No.IsValid() {
		res.No = item.No.Ptr()
	}
	if item.RefundPrice != nil {
		res.RefundPrice = item.RefundPrice
	}
	if item.ChannelUserId.IsValid() {
		res.ChannelUserId = item.ChannelUserId.Ptr()
	}
	if item.ChannelOrderNo.IsValid() {
		res.ChannelOrderNo = item.ChannelOrderNo.Ptr()
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
