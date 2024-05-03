package notify

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_notify_template 站内信模板表

// SystemNotifyTemplateDao 数据转换
func SystemNotifyTemplateDao(item *SystemNotifyTemplateObject) *dao.SystemNotifyTemplate {
	daoItem := &dao.SystemNotifyTemplate{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 主键
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 模板名称
	}
	if item != nil && item.Code != nil {
		daoItem.Code = item.Code // 模版编码
	}
	if item != nil && item.Nickname != nil {
		daoItem.Nickname = item.Nickname // 发送人名称
	}
	if item != nil && item.Content != nil {
		daoItem.Content = item.Content // 模版内容
	}
	if item != nil && item.Type != nil {
		daoItem.Type = item.Type // 类型
	}
	if item != nil && item.Params != nil {
		daoItem.Params = null.JSONFrom(item.GetParams()) // 参数数组
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态
	}
	if item != nil && item.Remark != nil {
		daoItem.Remark = null.StringFrom(item.GetRemark()) // 备注
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 删除
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

// SystemNotifyTemplateProto 数据绑定
func SystemNotifyTemplateProto(item dao.SystemNotifyTemplate) *SystemNotifyTemplateObject {
	res := &SystemNotifyTemplateObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.Code != nil {
		res.Code = item.Code
	}
	if item.Nickname != nil {
		res.Nickname = item.Nickname
	}
	if item.Content != nil {
		res.Content = item.Content
	}
	if item.Type != nil {
		res.Type = item.Type
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
	if item.Deleted != nil {
		res.Deleted = item.Deleted
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
