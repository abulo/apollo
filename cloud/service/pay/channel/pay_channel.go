package channel

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_channel 支付渠道

// PayChannelDao 数据转换
func PayChannelDao(item *PayChannelObject) *dao.PayChannel {
	daoItem := &dao.PayChannel{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 商户编号
	}
	if item != nil && item.Code != nil {
		daoItem.Code = item.Code // 渠道编码
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 开启状态
	}
	if item != nil && item.Remark != nil {
		daoItem.Remark = null.StringFrom(item.GetRemark()) // 备注
	}
	if item != nil && item.FeeRate != nil {
		daoItem.FeeRate = item.FeeRate // 渠道费率，单位：百分比
	}
	if item != nil && item.AppId != nil {
		daoItem.AppId = item.AppId // 应用编号
	}
	if item != nil && item.Config != nil {
		daoItem.Config = null.JSONFrom(item.GetConfig()) // 支付渠道配置
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

// PayChannelProto 数据绑定
func PayChannelProto(item dao.PayChannel) *PayChannelObject {
	res := &PayChannelObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Code != nil {
		res.Code = item.Code
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Remark.IsValid() {
		res.Remark = item.Remark.Ptr()
	}
	if item.FeeRate != nil {
		res.FeeRate = item.FeeRate
	}
	if item.AppId != nil {
		res.AppId = item.AppId
	}
	if item.Config.IsValid() {
		res.Config = *item.Config.Ptr()
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
