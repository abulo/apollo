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

// pay_notify_task 商户支付-任务通知

// SrvPayNotifyTaskServiceServer 商户支付-任务通知
type SrvPayNotifyTaskServiceServer struct {
	UnimplementedPayNotifyTaskServiceServer
	Server *xgrpc.Server
}

// PayNotifyTaskCreate 创建数据
func (srv SrvPayNotifyTaskServiceServer) PayNotifyTaskCreate(ctx context.Context, request *PayNotifyTaskCreateRequest) (*PayNotifyTaskCreateResponse, error) {
	req := PayNotifyTaskDao(request.GetData())
	data, err := notify.PayNotifyTaskCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:商户支付-任务通知:pay_notify_task:PayNotifyTaskCreate")
		return &PayNotifyTaskCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayNotifyTaskCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// PayNotifyTaskUpdate 更新数据
func (srv SrvPayNotifyTaskServiceServer) PayNotifyTaskUpdate(ctx context.Context, request *PayNotifyTaskUpdateRequest) (*PayNotifyTaskUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayNotifyTaskUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := PayNotifyTaskDao(request.GetData())
	_, err := notify.PayNotifyTaskUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:商户支付-任务通知:pay_notify_task:PayNotifyTaskUpdate")
		return &PayNotifyTaskUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayNotifyTaskUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayNotifyTaskDelete 删除数据
func (srv SrvPayNotifyTaskServiceServer) PayNotifyTaskDelete(ctx context.Context, request *PayNotifyTaskDeleteRequest) (*PayNotifyTaskDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayNotifyTaskDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := notify.PayNotifyTaskDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:商户支付-任务通知:pay_notify_task:PayNotifyTaskDelete")
		return &PayNotifyTaskDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayNotifyTaskDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayNotifyTask 查询单条数据
func (srv SrvPayNotifyTaskServiceServer) PayNotifyTask(ctx context.Context, request *PayNotifyTaskRequest) (*PayNotifyTaskResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayNotifyTaskResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := notify.PayNotifyTask(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:商户支付-任务通知:pay_notify_task:PayNotifyTask")
		return &PayNotifyTaskResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayNotifyTaskResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayNotifyTaskProto(res),
	}, nil
}

// PayNotifyTaskRecover 恢复数据
func (srv SrvPayNotifyTaskServiceServer) PayNotifyTaskRecover(ctx context.Context, request *PayNotifyTaskRecoverRequest) (*PayNotifyTaskRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayNotifyTaskRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := notify.PayNotifyTaskRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:商户支付-任务通知:pay_notify_task:PayNotifyTaskRecover")
		return &PayNotifyTaskRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayNotifyTaskRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvPayNotifyTaskServiceServer) PayNotifyTaskList(ctx context.Context, request *PayNotifyTaskListRequest) (*PayNotifyTaskListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.AppId != nil {
		condition["appId"] = request.GetAppId()
	}
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.DataId != nil {
		condition["dataId"] = request.GetDataId()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.MerchantOrderId != nil {
		condition["merchantOrderId"] = request.GetMerchantOrderId()
	}
	if request.BeginCreateTime != nil {
		condition["beginCreateTime"] = request.GetBeginCreateTime()
	}
	if request.FinishCreateTime != nil {
		condition["finishCreateTime"] = request.GetFinishCreateTime()
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
	list, err := notify.PayNotifyTaskList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:商户支付-任务通知:pay_notify_task:PayNotifyTaskList")
		return &PayNotifyTaskListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*PayNotifyTaskObject
	for _, item := range list {
		res = append(res, PayNotifyTaskProto(item))
	}
	return &PayNotifyTaskListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// PayNotifyTaskListTotal 获取总数
func (srv SrvPayNotifyTaskServiceServer) PayNotifyTaskListTotal(ctx context.Context, request *PayNotifyTaskListTotalRequest) (*PayNotifyTaskTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.AppId != nil {
		condition["appId"] = request.GetAppId()
	}
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.DataId != nil {
		condition["dataId"] = request.GetDataId()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.MerchantOrderId != nil {
		condition["merchantOrderId"] = request.GetMerchantOrderId()
	}
	if request.BeginCreateTime != nil {
		condition["beginCreateTime"] = request.GetBeginCreateTime()
	}
	if request.FinishCreateTime != nil {
		condition["finishCreateTime"] = request.GetFinishCreateTime()
	}

	// 获取数据集合
	total, err := notify.PayNotifyTaskListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:商户支付-任务通知:pay_notify_task:PayNotifyTaskListTotal")
		return &PayNotifyTaskTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayNotifyTaskTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
