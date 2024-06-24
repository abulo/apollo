package wallet

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_wallet_recharge 会员钱包充值

// PayWalletRechargeDao 数据转换
func PayWalletRechargeDao(item *PayWalletRechargeObject) *dao.PayWalletRecharge {
	daoItem := &dao.PayWalletRecharge{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.WalletId != nil {
		daoItem.WalletId = item.WalletId // 会员钱包id
	}
	if item != nil && item.TotalPrice != nil {
		daoItem.TotalPrice = item.TotalPrice // 用户实际到账余额，例如充 100 送 20，则该值是 120
	}
	if item != nil && item.PayPrice != nil {
		daoItem.PayPrice = item.PayPrice // 实际支付金额
	}
	if item != nil && item.BonusPrice != nil {
		daoItem.BonusPrice = item.BonusPrice // 钱包赠送金额
	}
	if item != nil && item.PackageId != nil {
		daoItem.PackageId = null.Int64From(item.GetPackageId()) // 充值套餐编号
	}
	if item != nil && item.PayStatus != nil {
		daoItem.PayStatus = item.PayStatus // 是否已支付：[0:未支付 1:已经支付过]
	}
	if item != nil && item.PayOrderId != nil {
		daoItem.PayOrderId = item.PayOrderId // 支付订单编号
	}
	if item != nil && item.PayChannelCode != nil {
		daoItem.PayChannelCode = item.PayChannelCode // 支付成功的支付渠道
	}
	if item != nil && item.PayTime != nil {
		daoItem.PayTime = null.DateTimeFrom(util.GrpcTime(item.PayTime)) // 订单支付时间
	}
	if item != nil && item.PayRefundId != nil {
		daoItem.PayRefundId = null.Int64From(item.GetPayRefundId()) // 支付退款单编号
	}
	if item != nil && item.RefundTotalPrice != nil {
		daoItem.RefundTotalPrice = null.Int64From(item.GetRefundTotalPrice()) // 退款金额，包含赠送金额
	}
	if item != nil && item.RefundPayPrice != nil {
		daoItem.RefundPayPrice = null.Int64From(item.GetRefundPayPrice()) // 退款支付金额
	}
	if item != nil && item.RefundBonusPrice != nil {
		daoItem.RefundBonusPrice = null.Int64From(item.GetRefundBonusPrice()) // 退款钱包赠送金额
	}
	if item != nil && item.RefundTime != nil {
		daoItem.RefundTime = null.DateTimeFrom(util.GrpcTime(item.RefundTime)) // 退款时间
	}
	if item != nil && item.RefundStatus != nil {
		daoItem.RefundStatus = null.Int32From(item.GetRefundStatus()) // 退款状态
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

// PayWalletRechargeProto 数据绑定
func PayWalletRechargeProto(item dao.PayWalletRecharge) *PayWalletRechargeObject {
	res := &PayWalletRechargeObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.WalletId != nil {
		res.WalletId = item.WalletId
	}
	if item.TotalPrice != nil {
		res.TotalPrice = item.TotalPrice
	}
	if item.PayPrice != nil {
		res.PayPrice = item.PayPrice
	}
	if item.BonusPrice != nil {
		res.BonusPrice = item.BonusPrice
	}
	if item.PackageId.IsValid() {
		res.PackageId = item.PackageId.Ptr()
	}
	if item.PayStatus != nil {
		res.PayStatus = item.PayStatus
	}
	if item.PayOrderId != nil {
		res.PayOrderId = item.PayOrderId
	}
	if item.PayChannelCode != nil {
		res.PayChannelCode = item.PayChannelCode
	}
	if item.PayTime.IsValid() {
		res.PayTime = timestamppb.New(*item.PayTime.Ptr())
	}
	if item.PayRefundId.IsValid() {
		res.PayRefundId = item.PayRefundId.Ptr()
	}
	if item.RefundTotalPrice.IsValid() {
		res.RefundTotalPrice = item.RefundTotalPrice.Ptr()
	}
	if item.RefundPayPrice.IsValid() {
		res.RefundPayPrice = item.RefundPayPrice.Ptr()
	}
	if item.RefundBonusPrice.IsValid() {
		res.RefundBonusPrice = item.RefundBonusPrice.Ptr()
	}
	if item.RefundTime.IsValid() {
		res.RefundTime = timestamppb.New(*item.RefundTime.Ptr())
	}
	if item.RefundStatus.IsValid() {
		res.RefundStatus = item.RefundStatus.Ptr()
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
