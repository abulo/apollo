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

// SystemLoginLogUpdate 更新数据
func (srv SrvSystemLoginLogServiceServer) SystemLoginLogUpdate(ctx context.Context, request *SystemLoginLogUpdateRequest) (*SystemLoginLogUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemLoginLogUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemLoginLogDao(request.GetData())
	_, err := logger.SystemLoginLogUpdate(ctx, id, *req)
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
	id := request.GetId()
	if id < 1 {
		return &SystemLoginLogDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := logger.SystemLoginLogDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
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

// SystemLoginLogRecover 恢复数据
func (srv SrvSystemLoginLogServiceServer) SystemLoginLogRecover(ctx context.Context, request *SystemLoginLogRecoverRequest) (*SystemLoginLogRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemLoginLogRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := logger.SystemLoginLogRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:登录日志:system_login_log:SystemLoginLogRecover")
		return &SystemLoginLogRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemLoginLogRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemLoginLogServiceServer) SystemLoginLogList(ctx context.Context, request *SystemLoginLogListRequest) (*SystemLoginLogListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
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

	paginationRequest := request.GetPagination()
	if paginationRequest != nil {
		// 当前页面
		pageNum := paginationRequest.GetPageNum()
		// 每页多少数据
		pageSize := paginationRequest.GetPageSize()
		if pageNum < 1 {
			pageNum = 1
		}
		if pageSize < 1 {
			pageSize = 10
		}
		// 分页数据
		offset := pageSize * (pageNum - 1)
		pagination := &sql.Pagination{
			Offset: &offset,
			Limit:  &pageSize,
		}
		condition["pagination"] = pagination
	}
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
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
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
	condition := make(map[string]any)
	// 构造查询条件
	if request.Ids != nil {
		req := request.GetIds()
		var ids []int64
		if err := json.Unmarshal(req, &ids); err != nil {
			return &SystemLoginLogDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
		}
		condition["ids"] = ids
	}
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
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
