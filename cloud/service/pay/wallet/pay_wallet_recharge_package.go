package wallet

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_wallet_recharge_package 充值套餐表

// PayWalletRechargePackageDao 数据转换
func PayWalletRechargePackageDao(item *PayWalletRechargePackageObject) *dao.PayWalletRechargePackage {
	daoItem := &dao.PayWalletRechargePackage{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 套餐名称
	}
	if item != nil && item.PayPrice != nil {
		daoItem.PayPrice = item.PayPrice // 支付金额
	}
	if item != nil && item.BonusPrice != nil {
		daoItem.BonusPrice = item.BonusPrice // 赠送金额
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态
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

// PayWalletRechargePackageProto 数据绑定
func PayWalletRechargePackageProto(item dao.PayWalletRechargePackage) *PayWalletRechargePackageObject {
	res := &PayWalletRechargePackageObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.PayPrice != nil {
		res.PayPrice = item.PayPrice
	}
	if item.BonusPrice != nil {
		res.BonusPrice = item.BonusPrice
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
