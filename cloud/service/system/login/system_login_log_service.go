package login

import (
	"cloud/code"
	"cloud/dao"
	"cloud/module/system/login"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_login_log 登录日志

// SrvSystemLoginLogServiceServer 登录日志
type SrvSystemLoginLogServiceServer struct {
	UnimplementedSystemLoginLogServiceServer
	Server *xgrpc.Server
}

func (srv SrvSystemLoginLogServiceServer) SystemLoginLogConvert(request *SystemLoginLogObject) dao.SystemLoginLog {
	var res dao.SystemLoginLog

	if request != nil && request.Id != nil {
		res.Id = request.Id // 主键
	}
	if request != nil && request.Username != nil {
		res.Username = request.Username // 用户账号
	}
	if request != nil && request.UserIp != nil {
		res.UserIp = request.UserIp // 用户ip
	}
	if request != nil && request.UserAgent != nil {
		res.UserAgent = null.StringFrom(request.GetUserAgent()) // UA
	}
	if request != nil && request.LoginTime != nil {
		res.LoginTime = null.DateTimeFrom(util.GrpcTime(request.LoginTime)) // 登录时间
	}
	if request != nil && request.Channel != nil {
		res.Channel = request.Channel // 渠道
	}
	if request != nil && request.Creator != nil {
		res.Creator = null.StringFrom(request.GetCreator()) // 创建人
	}
	if request != nil && request.CreateTime != nil {
		res.CreateTime = null.DateTimeFrom(util.GrpcTime(request.CreateTime)) // 创建时间
	}
	if request != nil && request.Updater != nil {
		res.Updater = null.StringFrom(request.GetUpdater()) // 更新人
	}
	if request != nil && request.UpdateTime != nil {
		res.UpdateTime = null.DateTimeFrom(util.GrpcTime(request.UpdateTime)) // 更新时间
	}

	return res
}

func (srv SrvSystemLoginLogServiceServer) SystemLoginLogResult(item dao.SystemLoginLog) *SystemLoginLogObject {
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

// SystemLoginLogCreate 创建数据
func (srv SrvSystemLoginLogServiceServer) SystemLoginLogCreate(ctx context.Context, request *SystemLoginLogCreateRequest) (*SystemLoginLogCreateResponse, error) {
	req := srv.SystemLoginLogConvert(request.GetData())
	data, err := login.SystemLoginLogCreate(ctx, req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:登录日志:system_login_log:SystemLoginLogCreate")
		return &SystemLoginLogCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemLoginLogCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemLoginLogUpdate 更新数据
func (srv SrvSystemLoginLogServiceServer) SystemLoginLogUpdate(ctx context.Context, request *SystemLoginLogUpdateRequest) (*SystemLoginLogUpdateResponse, error) {
	systemLoginLogId := request.GetSystemLoginLogId()
	if systemLoginLogId < 1 {
		return &SystemLoginLogUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := srv.SystemLoginLogConvert(request.GetData())
	_, err := login.SystemLoginLogUpdate(ctx, systemLoginLogId, req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:登录日志:system_login_log:SystemLoginLogUpdate")
		return &SystemLoginLogUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemLoginLogUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemLoginLogDelete 删除数据
func (srv SrvSystemLoginLogServiceServer) SystemLoginLogDelete(ctx context.Context, request *SystemLoginLogDeleteRequest) (*SystemLoginLogDeleteResponse, error) {
	systemLoginLogId := request.GetSystemLoginLogId()
	if systemLoginLogId < 1 {
		return &SystemLoginLogDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := login.SystemLoginLogDelete(ctx, systemLoginLogId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemLoginLogId,
			"err": err,
		}).Error("Sql:登录日志:system_login_log:SystemLoginLogDelete")
		return &SystemLoginLogDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemLoginLogDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemLoginLog 查询单条数据
func (srv SrvSystemLoginLogServiceServer) SystemLoginLog(ctx context.Context, request *SystemLoginLogRequest) (*SystemLoginLogResponse, error) {
	systemLoginLogId := request.GetSystemLoginLogId()
	if systemLoginLogId < 1 {
		return &SystemLoginLogResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := login.SystemLoginLog(ctx, systemLoginLogId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemLoginLogId,
			"err": err,
		}).Error("Sql:登录日志:system_login_log:SystemLoginLog")
		return &SystemLoginLogResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemLoginLogResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: srv.SystemLoginLogResult(res),
	}, nil
}
func (srv SrvSystemLoginLogServiceServer) SystemLoginLogList(ctx context.Context, request *SystemLoginLogListRequest) (*SystemLoginLogListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}
	if request.LoginTime != nil {
		condition["loginTime"] = request.GetLoginTime()
	}
	if request.Channel != nil {
		condition["channel"] = request.GetChannel()
	}

	// 当前页面
	pageNum := request.GetPageNum()
	// 每页多少数据
	pageSize := request.GetPageSize()
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	// 分页数据
	offset := pageSize * (pageNum - 1)
	condition["offset"] = offset
	condition["limit"] = pageSize
	// 获取数据集合
	list, err := login.SystemLoginLogList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:登录日志:system_login_log:SystemLoginLogList")
		return &SystemLoginLogListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemLoginLogObject
	for _, item := range list {
		res = append(res, srv.SystemLoginLogResult(item))
	}
	return &SystemLoginLogListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemLoginLogListTotal 获取总数
func (srv SrvSystemLoginLogServiceServer) SystemLoginLogListTotal(ctx context.Context, request *SystemLoginLogListTotalRequest) (*SystemLoginLogTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}
	if request.LoginTime != nil {
		condition["loginTime"] = request.GetLoginTime()
	}
	if request.Channel != nil {
		condition["channel"] = request.GetChannel()
	}

	// 获取数据集合
	total, err := login.SystemLoginLogListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:登录日志:system_login_log:SystemLoginLogListTotal")
		return &SystemLoginLogTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemLoginLogTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
