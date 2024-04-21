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

// system_login_log 登录日志

// SrvSystemLoginLogServiceServer 登录日志
type SrvSystemLoginLogServiceServer struct {
	UnimplementedSystemLoginLogServiceServer
	Server *xgrpc.Server
}

// SystemLoginLogCreate 创建数据
func (srv SrvSystemLoginLogServiceServer) SystemLoginLogCreate(ctx context.Context, request *SystemLoginLogCreateRequest) (*SystemLoginLogCreateResponse, error) {
	req := SystemLoginLogDao(request.GetData())
	data, err := logger.SystemLoginLogCreate(ctx, *req)
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

// SystemLoginLogDelete 删除数据
func (srv SrvSystemLoginLogServiceServer) SystemLoginLogDelete(ctx context.Context, request *SystemLoginLogDeleteRequest) (*SystemLoginLogDeleteResponse, error) {
	req := request.GetIds()
	var ids []int64
	if err := json.Unmarshal(req, &ids); err != nil {
		return &SystemLoginLogDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := logger.SystemLoginLogDelete(ctx, ids)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
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
	id := request.GetId()
	if id < 1 {
		return &SystemLoginLogResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := logger.SystemLoginLog(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:登录日志:system_login_log:SystemLoginLog")
		return &SystemLoginLogResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemLoginLogResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemLoginLogProto(res),
	}, nil
}
func (srv SrvSystemLoginLogServiceServer) SystemLoginLogList(ctx context.Context, request *SystemLoginLogListRequest) (*SystemLoginLogListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}
	if request.BeginLoginTime != nil {
		condition["beginLoginTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetBeginLoginTime()))
	}
	if request.FinishLoginTime != nil {
		condition["finishLoginTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetFinishLoginTime()))
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
	list, err := logger.SystemLoginLogList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:登录日志:system_login_log:SystemLoginLogList")
		return &SystemLoginLogListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemLoginLogObject
	for _, item := range list {
		res = append(res, SystemLoginLogProto(item))
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
	if request.BeginLoginTime != nil {
		condition["beginLoginTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetBeginLoginTime()))
	}
	if request.FinishLoginTime != nil {
		condition["finishLoginTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetFinishLoginTime()))
	}
	if request.Channel != nil {
		condition["channel"] = request.GetChannel()
	}

	// 获取数据集合
	total, err := logger.SystemLoginLogListTotal(ctx, condition)
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

func (srv SrvSystemLoginLogServiceServer) SystemLoginLogDrop(ctx context.Context, request *SystemLoginLogDropRequest) (*SystemLoginLogDropResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}
	if request.BeginLoginTime != nil {
		condition["beginLoginTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetBeginLoginTime()))
	}
	if request.FinishLoginTime != nil {
		condition["finishLoginTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetFinishLoginTime()))
	}
	if request.Channel != nil {
		condition["channel"] = request.GetChannel()
	}

	// 获取数据集合
	_, err := logger.SystemLoginLogDrop(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:登录日志:system_login_log:SystemLoginLogList")
		return &SystemLoginLogDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemLoginLogDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
