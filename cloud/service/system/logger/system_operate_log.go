package logger

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_operate_log 操作日志

// SystemOperateLogDao 数据转换
func SystemOperateLogDao(item *SystemOperateLogObject) *dao.SystemOperateLog {
	daoItem := &dao.SystemOperateLog{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 主键
	}
	if item != nil && item.Username != nil {
		daoItem.Username = item.Username // 用户账号
	}
	if item != nil && item.Module != nil {
		daoItem.Module = item.Module // 模块名称
	}
	if item != nil && item.RequestMethod != nil {
		daoItem.RequestMethod = item.RequestMethod // 请求方法名
	}
	if item != nil && item.RequestUrl != nil {
		daoItem.RequestUrl = item.RequestUrl // 请求地址
	}
	if item != nil && item.UserIp != nil {
		daoItem.UserIp = item.UserIp // 用户 ip
	}
	if item != nil && item.UserAgent != nil {
		daoItem.UserAgent = null.StringFrom(item.GetUserAgent()) // UA
	}
	if item != nil && item.GoMethod != nil {
		daoItem.GoMethod = item.GoMethod // 方法名
	}
	if item != nil && item.GoMethodArgs != nil {
		daoItem.GoMethodArgs = null.JSONFrom(item.GetGoMethodArgs()) // 方法的参数
	}
	if item != nil && item.StartTime != nil {
		daoItem.StartTime = null.DateTimeFrom(util.GrpcTime(item.StartTime)) // 操作开始时间
	}
	if item != nil && item.Duration != nil {
		daoItem.Duration = item.Duration // 执行时长
	}
	if item != nil && item.Channel != nil {
		daoItem.Channel = item.Channel // 渠道
	}
	if item != nil && item.Result != nil {
		daoItem.Result = item.Result // 结果(0 成功/1 失败)
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

// SystemOperateLogProto 数据绑定
func SystemOperateLogProto(item dao.SystemOperateLog) *SystemOperateLogObject {
	res := &SystemOperateLogObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Username != nil {
		res.Username = item.Username
	}
	if item.Module != nil {
		res.Module = item.Module
	}
	if item.RequestMethod != nil {
		res.RequestMethod = item.RequestMethod
	}
	if item.RequestUrl != nil {
		res.RequestUrl = item.RequestUrl
	}
	if item.UserIp != nil {
		res.UserIp = item.UserIp
	}
	if item.UserAgent.IsValid() {
		res.UserAgent = item.UserAgent.Ptr()
	}
	if item.GoMethod != nil {
		res.GoMethod = item.GoMethod
	}
	if item.GoMethodArgs.IsValid() {
		res.GoMethodArgs = *item.GoMethodArgs.Ptr()
	}
	if item.StartTime.IsValid() {
		res.StartTime = timestamppb.New(*item.StartTime.Ptr())
	}
	if item.Duration != nil {
		res.Duration = item.Duration
	}
	if item.Channel != nil {
		res.Channel = item.Channel
	}
	if item.Result != nil {
		res.Result = item.Result
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
