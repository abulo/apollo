package mail

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_mail_account 邮箱账号表

// SystemMailAccountDao 数据转换
func SystemMailAccountDao(item *SystemMailAccountObject) *dao.SystemMailAccount {
	daoItem := &dao.SystemMailAccount{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 主键
	}
	if item != nil && item.Mail != nil {
		daoItem.Mail = item.Mail // 邮箱
	}
	if item != nil && item.Username != nil {
		daoItem.Username = item.Username // 用户名
	}
	if item != nil && item.Password != nil {
		daoItem.Password = item.Password // 密码
	}
	if item != nil && item.Host != nil {
		daoItem.Host = item.Host // SMTP 服务器域名
	}
	if item != nil && item.Port != nil {
		daoItem.Port = item.Port // SMTP 服务器端口
	}
	if item != nil && item.SslEnable != nil {
		daoItem.SslEnable = item.SslEnable // 是否开启 SSL
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

// SystemMailAccountProto 数据绑定
func SystemMailAccountProto(item dao.SystemMailAccount) *SystemMailAccountObject {
	res := &SystemMailAccountObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Mail != nil {
		res.Mail = item.Mail
	}
	if item.Username != nil {
		res.Username = item.Username
	}
	if item.Password != nil {
		res.Password = item.Password
	}
	if item.Host != nil {
		res.Host = item.Host
	}
	if item.Port != nil {
		res.Port = item.Port
	}
	if item.SslEnable != nil {
		res.SslEnable = item.SslEnable
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
