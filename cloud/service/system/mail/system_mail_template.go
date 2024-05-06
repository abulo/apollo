package mail

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_mail_template 邮件模版表

// SystemMailTemplateDao 数据转换
func SystemMailTemplateDao(item *SystemMailTemplateObject) *dao.SystemMailTemplate {
	daoItem := &dao.SystemMailTemplate{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 模板名称
	}
	if item != nil && item.Code != nil {
		daoItem.Code = item.Code // 模板编码
	}
	if item != nil && item.AccountId != nil {
		daoItem.AccountId = item.AccountId // 发送的邮箱账号编号
	}
	if item != nil && item.Nickname != nil {
		daoItem.Nickname = null.StringFrom(item.GetNickname()) // 发送人名称
	}
	if item != nil && item.Title != nil {
		daoItem.Title = item.Title // 模板标题
	}
	if item != nil && item.Content != nil {
		daoItem.Content = item.Content // 模板内容
	}
	if item != nil && item.Params != nil {
		daoItem.Params = null.JSONFrom(item.GetParams()) // 参数数组
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 开启状态
	}
	if item != nil && item.Remark != nil {
		daoItem.Remark = null.StringFrom(item.GetRemark()) // 备注
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

// SystemMailTemplateProto 数据绑定
func SystemMailTemplateProto(item dao.SystemMailTemplate) *SystemMailTemplateObject {
	res := &SystemMailTemplateObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.Code != nil {
		res.Code = item.Code
	}
	if item.AccountId != nil {
		res.AccountId = item.AccountId
	}
	if item.Nickname.IsValid() {
		res.Nickname = item.Nickname.Ptr()
	}
	if item.Title != nil {
		res.Title = item.Title
	}
	if item.Content != nil {
		res.Content = item.Content
	}
	if item.Params.IsValid() {
		res.Params = *item.Params.Ptr()
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Remark.IsValid() {
		res.Remark = item.Remark.Ptr()
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
