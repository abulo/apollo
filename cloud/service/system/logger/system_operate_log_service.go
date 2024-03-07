package logger

import (
	"cloud/code"
	"cloud/module/system/logger"
	"context"
	"encoding/json"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_operate_log 操作日志

// SrvSystemOperateLogServiceServer 操作日志
type SrvSystemOperateLogServiceServer struct {
	UnimplementedSystemOperateLogServiceServer
	Server *xgrpc.Server
}

// SystemOperateLogCreate 创建数据
func (srv SrvSystemOperateLogServiceServer) SystemOperateLogCreate(ctx context.Context, request *SystemOperateLogCreateRequest) (*SystemOperateLogCreateResponse, error) {
	req := SystemOperateLogDao(request.GetData())
	data, err := logger.SystemOperateLogCreate(ctx, *req)
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
		Data: SystemOperateLogProto(res),
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
		res = append(res, SystemOperateLogProto(item))
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
