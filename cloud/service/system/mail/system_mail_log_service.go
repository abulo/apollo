package mail

import (
	"cloud/code"
	"cloud/module/system/mail"
	"context"
	"encoding/json"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_mail_log 邮件日志表

// SrvSystemMailLogServiceServer 邮件日志表
type SrvSystemMailLogServiceServer struct {
	UnimplementedSystemMailLogServiceServer
	Server *xgrpc.Server
}

// SystemMailLogCreate 创建数据
func (srv SrvSystemMailLogServiceServer) SystemMailLogCreate(ctx context.Context, request *SystemMailLogCreateRequest) (*SystemMailLogCreateResponse, error) {
	req := SystemMailLogDao(request.GetData())
	data, err := mail.SystemMailLogCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:邮件日志表:system_mail_log:SystemMailLogCreate")
		return &SystemMailLogCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailLogCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemMailLogUpdate 更新数据
func (srv SrvSystemMailLogServiceServer) SystemMailLogUpdate(ctx context.Context, request *SystemMailLogUpdateRequest) (*SystemMailLogUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailLogUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemMailLogDao(request.GetData())
	_, err := mail.SystemMailLogUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:邮件日志表:system_mail_log:SystemMailLogUpdate")
		return &SystemMailLogUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailLogUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemMailLogDelete 删除数据
func (srv SrvSystemMailLogServiceServer) SystemMailLogDelete(ctx context.Context, request *SystemMailLogDeleteRequest) (*SystemMailLogDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailLogDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := mail.SystemMailLogDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:邮件日志表:system_mail_log:SystemMailLogDelete")
		return &SystemMailLogDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailLogDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemMailLog 查询单条数据
func (srv SrvSystemMailLogServiceServer) SystemMailLog(ctx context.Context, request *SystemMailLogRequest) (*SystemMailLogResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailLogResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := mail.SystemMailLog(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:邮件日志表:system_mail_log:SystemMailLog")
		return &SystemMailLogResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailLogResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemMailLogProto(res),
	}, nil
}

// SystemMailLogRecover 恢复数据
func (srv SrvSystemMailLogServiceServer) SystemMailLogRecover(ctx context.Context, request *SystemMailLogRecoverRequest) (*SystemMailLogRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailLogRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := mail.SystemMailLogRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:邮件日志表:system_mail_log:SystemMailLogRecover")
		return &SystemMailLogRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailLogRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemMailLogDrop 清理数据
func (srv SrvSystemMailLogServiceServer) SystemMailLogDrop(ctx context.Context, request *SystemMailLogDropRequest) (*SystemMailLogDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailLogDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := mail.SystemMailLogDrop(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:邮件日志表:system_mail_log:SystemMailLogDrop")
		return &SystemMailLogDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailLogDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemMailLogServiceServer) SystemMailLogList(ctx context.Context, request *SystemMailLogListRequest) (*SystemMailLogListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.SendStatus != nil {
		condition["sendStatus"] = request.GetSendStatus()
	}
	if request.BeginSendTime != nil {
		condition["beginSendTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetBeginSendTime()))
	}
	if request.FinishSendTime != nil {
		condition["finishSendTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetFinishSendTime()))
	}
	if request.TemplateTitle != nil {
		condition["templateTitle"] = request.GetTemplateTitle()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.UserType != nil {
		condition["userType"] = request.GetUserType()
	}
	if request.AccountId != nil {
		condition["accountId"] = request.GetAccountId()
	}
	if request.TemplateCode != nil {
		condition["templateCode"] = request.GetTemplateCode()
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
	list, err := mail.SystemMailLogList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:邮件日志表:system_mail_log:SystemMailLogList")
		return &SystemMailLogListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemMailLogObject
	for _, item := range list {
		res = append(res, SystemMailLogProto(item))
	}
	return &SystemMailLogListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemMailLogListTotal 获取总数
func (srv SrvSystemMailLogServiceServer) SystemMailLogListTotal(ctx context.Context, request *SystemMailLogListTotalRequest) (*SystemMailLogTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.SendStatus != nil {
		condition["sendStatus"] = request.GetSendStatus()
	}
	if request.BeginSendTime != nil {
		condition["beginSendTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetBeginSendTime()))
	}
	if request.FinishSendTime != nil {
		condition["finishSendTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetFinishSendTime()))
	}
	if request.TemplateTitle != nil {
		condition["templateTitle"] = request.GetTemplateTitle()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.UserType != nil {
		condition["userType"] = request.GetUserType()
	}
	if request.AccountId != nil {
		condition["accountId"] = request.GetAccountId()
	}
	if request.TemplateCode != nil {
		condition["templateCode"] = request.GetTemplateCode()
	}

	// 获取数据集合
	total, err := mail.SystemMailLogListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:邮件日志表:system_mail_log:SystemMailLogListTotal")
		return &SystemMailLogTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailLogTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}

// SystemMailLogMultipleDelete 批量删除
func (srv SrvSystemMailLogServiceServer) SystemMailLogMultipleDelete(ctx context.Context, request *SystemMailLogMultipleRequest) (*SystemMailLogMultipleResponse, error) {
	condition := make(map[string]any)
	if request.Ids != nil {
		var ids []int64
		json.Unmarshal(request.GetIds(), &ids)
		condition["ids"] = ids
	}
	// 获取数据集合
	_, err := mail.SystemMailLogMultipleDelete(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:邮件日志表:system_mail_log:SystemMailLogMultipleDelete")
		return &SystemMailLogMultipleResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailLogMultipleResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemMailLogMultipleRecover 批量恢复
func (srv SrvSystemMailLogServiceServer) SystemMailLogMultipleRecover(ctx context.Context, request *SystemMailLogMultipleRequest) (*SystemMailLogMultipleResponse, error) {
	condition := make(map[string]any)
	if request.Ids != nil {
		var ids []int64
		json.Unmarshal(request.GetIds(), &ids)
		condition["ids"] = ids
	}
	// 获取数据集合
	_, err := mail.SystemMailLogMultipleRecover(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:邮件日志表:system_mail_log:SystemMailLogMultipleRecover")
		return &SystemMailLogMultipleResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailLogMultipleResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemMailLogMultipleDrop 批量清理
func (srv SrvSystemMailLogServiceServer) SystemMailLogMultipleDrop(ctx context.Context, request *SystemMailLogMultipleRequest) (*SystemMailLogMultipleResponse, error) {
	condition := make(map[string]any)
	if request.Ids != nil {
		var ids []int64
		json.Unmarshal(request.GetIds(), &ids)
		condition["ids"] = ids
	}
	// 获取数据集合
	_, err := mail.SystemMailLogMultipleDrop(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:邮件日志表:system_mail_log:SystemMailLogMultipleDrop")
		return &SystemMailLogMultipleResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailLogMultipleResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
