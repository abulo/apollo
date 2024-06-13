package mail

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_mail_log 邮件日志表

// SystemMailLogDao 数据转换
func SystemMailLogDao(item *SystemMailLogObject) *dao.SystemMailLog {
	daoItem := &dao.SystemMailLog{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.UserId != nil {
		daoItem.UserId = null.Int64From(item.GetUserId()) // 用户编号
	}
	if item != nil && item.UserType != nil {
		daoItem.UserType = null.Int32From(item.GetUserType()) // 用户类型
	}
	if item != nil && item.ToMail != nil {
		daoItem.ToMail = item.ToMail // 接收邮箱地址
	}
	if item != nil && item.AccountId != nil {
		daoItem.AccountId = item.AccountId // 邮箱账号编号
	}
	if item != nil && item.FromMail != nil {
		daoItem.FromMail = item.FromMail // 发送邮箱地址
	}
	if item != nil && item.TemplateId != nil {
		daoItem.TemplateId = item.TemplateId // 模板编号
	}
	if item != nil && item.TemplateCode != nil {
		daoItem.TemplateCode = item.TemplateCode // 模板编码
	}
	if item != nil && item.TemplateNickname != nil {
		daoItem.TemplateNickname = null.StringFrom(item.GetTemplateNickname()) // 模版发送人名称
	}
	if item != nil && item.TemplateTitle != nil {
		daoItem.TemplateTitle = item.TemplateTitle // 邮件标题
	}
	if item != nil && item.TemplateContent != nil {
		daoItem.TemplateContent = item.TemplateContent // 邮件内容
	}
	if item != nil && item.TemplateParams != nil {
		daoItem.TemplateParams = null.JSONFrom(item.GetTemplateParams()) // 邮件参数
	}
	if item != nil && item.SendStatus != nil {
		daoItem.SendStatus = item.SendStatus // 发送状态
	}
	if item != nil && item.SendTime != nil {
		daoItem.SendTime = null.DateTimeFrom(util.GrpcTime(item.SendTime)) // 发送时间
	}
	if item != nil && item.SendMessageId != nil {
		daoItem.SendMessageId = null.StringFrom(item.GetSendMessageId()) // 发送返回的消息 ID
	}
	if item != nil && item.SendException != nil {
		daoItem.SendException = null.StringFrom(item.GetSendException()) // 发送异常
	}
	if item != nil && item.Creator != nil {
		daoItem.Creator = null.StringFrom(item.GetCreator()) // 创建者
	}
	if item != nil && item.CreateTime != nil {
		daoItem.CreateTime = null.DateTimeFrom(util.GrpcTime(item.CreateTime)) // 创建时间
	}
	if item != nil && item.Updater != nil {
		daoItem.Updater = null.StringFrom(item.GetUpdater()) // 更新者
	}
	if item != nil && item.UpdateTime != nil {
		daoItem.UpdateTime = null.DateTimeFrom(util.GrpcTime(item.UpdateTime)) // 更新时间
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 是否删除
	}

	return daoItem
}

// SystemMailLogProto 数据绑定
func SystemMailLogProto(item dao.SystemMailLog) *SystemMailLogObject {
	res := &SystemMailLogObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.UserId.IsValid() {
		res.UserId = item.UserId.Ptr()
	}
	if item.UserType.IsValid() {
		res.UserType = item.UserType.Ptr()
	}
	if item.ToMail != nil {
		res.ToMail = item.ToMail
	}
	if item.AccountId != nil {
		res.AccountId = item.AccountId
	}
	if item.FromMail != nil {
		res.FromMail = item.FromMail
	}
	if item.TemplateId != nil {
		res.TemplateId = item.TemplateId
	}
	if item.TemplateCode != nil {
		res.TemplateCode = item.TemplateCode
	}
	if item.TemplateNickname.IsValid() {
		res.TemplateNickname = item.TemplateNickname.Ptr()
	}
	if item.TemplateTitle != nil {
		res.TemplateTitle = item.TemplateTitle
	}
	if item.TemplateContent != nil {
		res.TemplateContent = item.TemplateContent
	}
	if item.TemplateParams.IsValid() {
		res.TemplateParams = *item.TemplateParams.Ptr()
	}
	if item.SendStatus != nil {
		res.SendStatus = item.SendStatus
	}
	if item.SendTime.IsValid() {
		res.SendTime = timestamppb.New(*item.SendTime.Ptr())
	}
	if item.SendMessageId.IsValid() {
		res.SendMessageId = item.SendMessageId.Ptr()
	}
	if item.SendException.IsValid() {
		res.SendException = item.SendException.Ptr()
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
	if item.Deleted != nil {
		res.Deleted = item.Deleted
	}

	return res
}
