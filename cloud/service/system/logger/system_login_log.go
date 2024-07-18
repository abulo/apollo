package logger

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_login_log 登录日志

// SystemLoginLogDao 数据转换
func SystemLoginLogDao(item *SystemLoginLogObject) *dao.SystemLoginLog {
	daoItem := &dao.SystemLoginLog{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 主键
	}
	if item != nil && item.Username != nil {
		daoItem.Username = item.Username // 用户账号
	}
	if item != nil && item.UserIp != nil {
		daoItem.UserIp = item.UserIp // 用户ip
	}
	if item != nil && item.UserAgent != nil {
		daoItem.UserAgent = null.StringFrom(item.GetUserAgent()) // UA
	}
	if item != nil && item.LoginTime != nil {
		daoItem.LoginTime = null.DateTimeFrom(util.GrpcTime(item.LoginTime)) // 登录时间
	}
	if item != nil && item.Channel != nil {
		daoItem.Channel = item.Channel // 渠道
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
	if item != nil && item.UserId != nil {
		daoItem.UserId = item.UserId // 用户id
	}

	return daoItem
}

// SystemLoginLogProto 数据绑定
func SystemLoginLogProto(item dao.SystemLoginLog) *SystemLoginLogObject {
	res := &SystemLoginLogObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Username != nil {
		res.Username = item.Username
	}
	if item.UserIp != nil {
		res.UserIp = item.UserIp
	}
	if item.UserAgent.IsValid() {
		res.UserAgent = item.UserAgent.Ptr()
	}
	if item.LoginTime.IsValid() {
		res.LoginTime = timestamppb.New(*item.LoginTime.Ptr())
	}
	if item.Channel != nil {
		res.Channel = item.Channel
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
	if item.UserId != nil {
		res.UserId = item.UserId
	}

	return res
}
