package wallet

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_wallet_transaction 会员钱包流水表

// PayWalletTransactionDao 数据转换
func PayWalletTransactionDao(item *PayWalletTransactionObject) *dao.PayWalletTransaction {
	daoItem := &dao.PayWalletTransaction{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.WalletId != nil {
		daoItem.WalletId = item.WalletId // 会员钱包 id
	}
	if item != nil && item.BizType != nil {
		daoItem.BizType = item.BizType // 关联类型
	}
	if item != nil && item.BizId != nil {
		daoItem.BizId = item.BizId // 关联业务编号
	}
	if item != nil && item.No != nil {
		daoItem.No = item.No // 流水号
	}
	if item != nil && item.Title != nil {
		daoItem.Title = item.Title // 流水标题
	}
	if item != nil && item.Price != nil {
		daoItem.Price = item.Price // 交易金额, 单位分
	}
	if item != nil && item.Balance != nil {
		daoItem.Balance = item.Balance // 余额, 单位分
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

// PayWalletTransactionProto 数据绑定
func PayWalletTransactionProto(item dao.PayWalletTransaction) *PayWalletTransactionObject {
	res := &PayWalletTransactionObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.WalletId != nil {
		res.WalletId = item.WalletId
	}
	if item.BizType != nil {
		res.BizType = item.BizType
	}
	if item.BizId != nil {
		res.BizId = item.BizId
	}
	if item.No != nil {
		res.No = item.No
	}
	if item.Title != nil {
		res.Title = item.Title
	}
	if item.Price != nil {
		res.Price = item.Price
	}
	if item.Balance != nil {
		res.Balance = item.Balance
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
