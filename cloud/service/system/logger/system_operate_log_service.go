package logger

import (
	"cloud/code"
	"cloud/dao"
	"cloud/module/system/logger"
	"context"
	"encoding/json"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_operate_log 操作日志

// SrvSystemOperateLogServiceServer 操作日志
type SrvSystemOperateLogServiceServer struct {
	UnimplementedSystemOperateLogServiceServer
	Server *xgrpc.Server
}

func (srv SrvSystemOperateLogServiceServer) SystemOperateLogConvert(request *SystemOperateLogObject) dao.SystemOperateLog {
	var res dao.SystemOperateLog

	if request != nil && request.Id != nil {
		res.Id = request.Id // 主键
	}
	if request != nil && request.Username != nil {
		res.Username = request.Username // 用户账号
	}
	if request != nil && request.Module != nil {
		res.Module = request.Module // 模块名称
	}
	if request != nil && request.RequestMethod != nil {
		res.RequestMethod = request.RequestMethod // 请求方法名
	}
	if request != nil && request.RequestUrl != nil {
		res.RequestUrl = request.RequestUrl // 请求地址
	}
	if request != nil && request.UserIp != nil {
		res.UserIp = request.UserIp // 用户 ip
	}
	if request != nil && request.UserAgent != nil {
		res.UserAgent = null.StringFrom(request.GetUserAgent()) // UA
	}
	if request != nil && request.GoMethod != nil {
		res.GoMethod = request.GoMethod // 方法名
	}
	if request != nil && request.GoMethodArgs != nil {
		res.GoMethodArgs = null.JSONFrom(request.GetGoMethodArgs()) // 方法的参数
	}
	if request != nil && request.StartTime != nil {
		res.StartTime = null.DateTimeFrom(util.GrpcTime(request.StartTime)) // 操作开始时间
	}
	if request != nil && request.Duration != nil {
		res.Duration = request.Duration // 执行时长
	}
	if request != nil && request.Channel != nil {
		res.Channel = request.Channel // 渠道
	}
	if request != nil && request.Result != nil {
		res.Result = request.Result // 结果(0 成功/1 失败)
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

func (srv SrvSystemOperateLogServiceServer) SystemOperateLogResult(item dao.SystemOperateLog) *SystemOperateLogObject {
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

// SystemOperateLogCreate 创建数据
func (srv SrvSystemOperateLogServiceServer) SystemOperateLogCreate(ctx context.Context, request *SystemOperateLogCreateRequest) (*SystemOperateLogCreateResponse, error) {
	req := srv.SystemOperateLogConvert(request.GetData())
	data, err := logger.SystemOperateLogCreate(ctx, req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:操作日志:system_operate_log:SystemOperateLogCreate")
		return &SystemOperateLogCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemOperateLogCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemOperateLogDelete 删除数据
func (srv SrvSystemOperateLogServiceServer) SystemOperateLogDelete(ctx context.Context, request *SystemOperateLogDeleteRequest) (*SystemOperateLogDeleteResponse, error) {
	// systemOperateLogId := request.GetSystemOperateLogId()
	req := request.GetSystemOperateLogIds()
	var systemOperateLogIds []int64
	if err := json.Unmarshal(req, &systemOperateLogIds); err != nil {
		return &SystemOperateLogDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := logger.SystemOperateLogDelete(ctx, systemOperateLogIds)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:操作日志:system_operate_log:SystemOperateLogDelete")
		return &SystemOperateLogDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemOperateLogDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemOperateLog 查询单条数据
func (srv SrvSystemOperateLogServiceServer) SystemOperateLog(ctx context.Context, request *SystemOperateLogRequest) (*SystemOperateLogResponse, error) {
	systemOperateLogId := request.GetSystemOperateLogId()
	if systemOperateLogId < 1 {
		return &SystemOperateLogResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := logger.SystemOperateLog(ctx, systemOperateLogId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemOperateLogId,
			"err": err,
		}).Error("Sql:操作日志:system_operate_log:SystemOperateLog")
		return &SystemOperateLogResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemOperateLogResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: srv.SystemOperateLogResult(res),
	}, nil
}
func (srv SrvSystemOperateLogServiceServer) SystemOperateLogList(ctx context.Context, request *SystemOperateLogListRequest) (*SystemOperateLogListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}
	if request.Module != nil {
		condition["module"] = request.GetModule()
	}
	if request.BeginStartTime != nil {
		condition["beginStartTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetBeginStartTime()))
	}
	if request.FinishStartTime != nil {
		condition["finishStartTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetFinishStartTime()))
	}
	if request.Result != nil {
		condition["result"] = request.GetResult()
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
	list, err := logger.SystemOperateLogList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:操作日志:system_operate_log:SystemOperateLogList")
		return &SystemOperateLogListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemOperateLogObject
	for _, item := range list {
		res = append(res, srv.SystemOperateLogResult(item))
	}
	return &SystemOperateLogListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemOperateLogListTotal 获取总数
func (srv SrvSystemOperateLogServiceServer) SystemOperateLogListTotal(ctx context.Context, request *SystemOperateLogListTotalRequest) (*SystemOperateLogTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}
	if request.Module != nil {
		condition["module"] = request.GetModule()
	}
	if request.BeginStartTime != nil {
		condition["beginStartTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetBeginStartTime()))
	}
	if request.FinishStartTime != nil {
		condition["finishStartTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetFinishStartTime()))
	}
	if request.Result != nil {
		condition["result"] = request.GetResult()
	}

	// 获取数据集合
	total, err := logger.SystemOperateLogListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:操作日志:system_operate_log:SystemOperateLogListTotal")
		return &SystemOperateLogTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemOperateLogTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}

func (srv SrvSystemOperateLogServiceServer) SystemOperateLogDrop(ctx context.Context, request *SystemOperateLogDropRequest) (*SystemOperateLogDropResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}
	if request.Module != nil {
		condition["module"] = request.GetModule()
	}
	if request.BeginStartTime != nil {
		condition["beginStartTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetBeginStartTime()))
	}
	if request.FinishStartTime != nil {
		condition["finishStartTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetFinishStartTime()))
	}
	if request.Result != nil {
		condition["result"] = request.GetResult()
	}

	// 获取数据集合
	_, err := logger.SystemOperateLogDrop(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:操作日志:system_operate_log:SystemOperateLogList")
		return &SystemOperateLogDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}

	return &SystemOperateLogDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
