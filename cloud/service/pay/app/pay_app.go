package app

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_app 支付应用信息

// PayAppDao 数据转换
func PayAppDao(item *PayAppObject) *dao.PayApp {
	daoItem := &dao.PayApp{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 应用编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 应用名称
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 开启状态
	}
	if item != nil && item.Remark != nil {
		daoItem.Remark = null.StringFrom(item.GetRemark()) // 备注
	}
	if item != nil && item.OrderNotifyUrl != nil {
		daoItem.OrderNotifyUrl = item.OrderNotifyUrl // 支付结果的回调地址
	}
	if item != nil && item.RefundNotifyUrl != nil {
		daoItem.RefundNotifyUrl = item.RefundNotifyUrl // 退款结果的回调地址
	}
	if item != nil && item.TransferNotifyUrl != nil {
		daoItem.TransferNotifyUrl = null.StringFrom(item.GetTransferNotifyUrl()) // 转账结果的回调地址
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

// PayAppProto 数据绑定
func PayAppProto(item dao.PayApp) *PayAppObject {
	res := &PayAppObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Remark.IsValid() {
		res.Remark = item.Remark.Ptr()
	}
	if item.OrderNotifyUrl != nil {
		res.OrderNotifyUrl = item.OrderNotifyUrl
	}
	if item.RefundNotifyUrl != nil {
		res.RefundNotifyUrl = item.RefundNotifyUrl
	}
	if item.TransferNotifyUrl.IsValid() {
		res.TransferNotifyUrl = item.TransferNotifyUrl.Ptr()
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
