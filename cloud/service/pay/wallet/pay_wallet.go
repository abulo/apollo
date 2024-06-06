package wallet

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_wallet 会员钱包表

// PayWalletDao 数据转换
func PayWalletDao(item *PayWalletObject) *dao.PayWallet {
	daoItem := &dao.PayWallet{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.UserId != nil {
		daoItem.UserId = item.UserId // 用户编号
	}
	if item != nil && item.UserType != nil {
		daoItem.UserType = item.UserType // 用户类型
	}
	if item != nil && item.Balance != nil {
		daoItem.Balance = item.Balance // 余额，单位分
	}
	if item != nil && item.TotalExpense != nil {
		daoItem.TotalExpense = item.TotalExpense // 累计支出，单位分
	}
	if item != nil && item.TotalRecharge != nil {
		daoItem.TotalRecharge = item.TotalRecharge // 累计充值，单位分
	}
	if item != nil && item.FreezePrice != nil {
		daoItem.FreezePrice = item.FreezePrice // 冻结金额，单位分
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

// PayWalletProto 数据绑定
func PayWalletProto(item dao.PayWallet) *PayWalletObject {
	res := &PayWalletObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.UserId != nil {
		res.UserId = item.UserId
	}
	if item.UserType != nil {
		res.UserType = item.UserType
	}
	if item.Balance != nil {
		res.Balance = item.Balance
	}
	if item.TotalExpense != nil {
		res.TotalExpense = item.TotalExpense
	}
	if item.TotalRecharge != nil {
		res.TotalRecharge = item.TotalRecharge
	}
	if item.FreezePrice != nil {
		res.FreezePrice = item.FreezePrice
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
