package notify

import (
	"cloud/code"
	"cloud/module/pay/notify"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// pay_notify_log 支付通知日志

// SrvPayNotifyLogServiceServer 支付通知日志
type SrvPayNotifyLogServiceServer struct {
	UnimplementedPayNotifyLogServiceServer
	Server *xgrpc.Server
}

// PayNotifyLogCreate 创建数据
func (srv SrvPayNotifyLogServiceServer) PayNotifyLogCreate(ctx context.Context, request *PayNotifyLogCreateRequest) (*PayNotifyLogCreateResponse, error) {
	req := PayNotifyLogDao(request.GetData())
	data, err := notify.PayNotifyLogCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:支付通知日志:pay_notify_log:PayNotifyLogCreate")
		return &PayNotifyLogCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayNotifyLogCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// PayNotifyLogUpdate 更新数据
func (srv SrvPayNotifyLogServiceServer) PayNotifyLogUpdate(ctx context.Context, request *PayNotifyLogUpdateRequest) (*PayNotifyLogUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayNotifyLogUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := PayNotifyLogDao(request.GetData())
	_, err := notify.PayNotifyLogUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:支付通知日志:pay_notify_log:PayNotifyLogUpdate")
		return &PayNotifyLogUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayNotifyLogUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayNotifyLogDelete 删除数据
func (srv SrvPayNotifyLogServiceServer) PayNotifyLogDelete(ctx context.Context, request *PayNotifyLogDeleteRequest) (*PayNotifyLogDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayNotifyLogDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := notify.PayNotifyLogDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付通知日志:pay_notify_log:PayNotifyLogDelete")
		return &PayNotifyLogDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayNotifyLogDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayNotifyLog 查询单条数据
func (srv SrvPayNotifyLogServiceServer) PayNotifyLog(ctx context.Context, request *PayNotifyLogRequest) (*PayNotifyLogResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayNotifyLogResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := notify.PayNotifyLog(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付通知日志:pay_notify_log:PayNotifyLog")
		return &PayNotifyLogResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayNotifyLogResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayNotifyLogProto(res),
	}, nil
}

// PayNotifyLogRecover 恢复数据
func (srv SrvPayNotifyLogServiceServer) PayNotifyLogRecover(ctx context.Context, request *PayNotifyLogRecoverRequest) (*PayNotifyLogRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayNotifyLogRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := notify.PayNotifyLogRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付通知日志:pay_notify_log:PayNotifyLogRecover")
		return &PayNotifyLogRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayNotifyLogRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvPayNotifyLogServiceServer) PayNotifyLogList(ctx context.Context, request *PayNotifyLogListRequest) (*PayNotifyLogListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.TaskId != nil {
		condition["taskId"] = request.GetTaskId()
	}

	// 获取数据集合
	list, err := notify.PayNotifyLogList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:支付通知日志:pay_notify_log:PayNotifyLogList")
		return &PayNotifyLogListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*PayNotifyLogObject
	for _, item := range list {
		res = append(res, PayNotifyLogProto(item))
	}
	return &PayNotifyLogListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
