package notify

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_notify_message 站内信消息表

// SystemNotifyMessageDao 数据转换
func SystemNotifyMessageDao(item *SystemNotifyMessageObject) *dao.SystemNotifyMessage {
	daoItem := &dao.SystemNotifyMessage{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 用户ID
	}
	if item != nil && item.UserId != nil {
		daoItem.UserId = item.UserId // 用户id
	}
	if item != nil && item.UserType != nil {
		daoItem.UserType = item.UserType // 用户类型
	}
	if item != nil && item.TemplateId != nil {
		daoItem.TemplateId = item.TemplateId // 模版编号
	}
	if item != nil && item.TemplateCode != nil {
		daoItem.TemplateCode = item.TemplateCode // 模板编码
	}
	if item != nil && item.TemplateNickname != nil {
		daoItem.TemplateNickname = item.TemplateNickname // 模版发送人名称
	}
	if item != nil && item.TemplateContent != nil {
		daoItem.TemplateContent = item.TemplateContent // 模版内容
	}
	if item != nil && item.TemplateType != nil {
		daoItem.TemplateType = item.TemplateType // 模版类型
	}
	if item != nil && item.TemplateParams != nil {
		daoItem.TemplateParams = null.JSONFrom(item.GetTemplateParams()) // 模版参数
	}
	if item != nil && item.ReadStatus != nil {
		daoItem.ReadStatus = item.ReadStatus // 是否已读
	}
	if item != nil && item.ReadTime != nil {
		daoItem.ReadTime = null.DateTimeFrom(util.GrpcTime(item.ReadTime)) // 阅读时间
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

// SystemNotifyMessageProto 数据绑定
func SystemNotifyMessageProto(item dao.SystemNotifyMessage) *SystemNotifyMessageObject {
	res := &SystemNotifyMessageObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.UserId != nil {
		res.UserId = item.UserId
	}
	if item.UserType != nil {
		res.UserType = item.UserType
	}
	if item.TemplateId != nil {
		res.TemplateId = item.TemplateId
	}
	if item.TemplateCode != nil {
		res.TemplateCode = item.TemplateCode
	}
	if item.TemplateNickname != nil {
		res.TemplateNickname = item.TemplateNickname
	}
	if item.TemplateContent != nil {
		res.TemplateContent = item.TemplateContent
	}
	if item.TemplateType != nil {
		res.TemplateType = item.TemplateType
	}
	if item.TemplateParams.IsValid() {
		res.TemplateParams = *item.TemplateParams.Ptr()
	}
	if item.ReadStatus != nil {
		res.ReadStatus = item.ReadStatus
	}
	if item.ReadTime.IsValid() {
		res.ReadTime = timestamppb.New(*item.ReadTime.Ptr())
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
