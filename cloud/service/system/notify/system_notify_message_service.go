package notify

import (
	"cloud/code"
	"cloud/module/system/notify"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_notify_message 站内信消息表

// SrvSystemNotifyMessageServiceServer 站内信消息表
type SrvSystemNotifyMessageServiceServer struct {
	UnimplementedSystemNotifyMessageServiceServer
	Server *xgrpc.Server
}

// SystemNotifyMessageCreate 创建数据
func (srv SrvSystemNotifyMessageServiceServer) SystemNotifyMessageCreate(ctx context.Context, request *SystemNotifyMessageCreateRequest) (*SystemNotifyMessageCreateResponse, error) {
	req := SystemNotifyMessageDao(request.GetData())
	data, err := notify.SystemNotifyMessageCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:站内信消息表:system_notify_message:SystemNotifyMessageCreate")
		return &SystemNotifyMessageCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyMessageCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemNotifyMessageUpdate 更新数据
func (srv SrvSystemNotifyMessageServiceServer) SystemNotifyMessageUpdate(ctx context.Context, request *SystemNotifyMessageUpdateRequest) (*SystemNotifyMessageUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNotifyMessageUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemNotifyMessageDao(request.GetData())
	_, err := notify.SystemNotifyMessageUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:站内信消息表:system_notify_message:SystemNotifyMessageUpdate")
		return &SystemNotifyMessageUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyMessageUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemNotifyMessageDelete 删除数据
func (srv SrvSystemNotifyMessageServiceServer) SystemNotifyMessageDelete(ctx context.Context, request *SystemNotifyMessageDeleteRequest) (*SystemNotifyMessageDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNotifyMessageDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := notify.SystemNotifyMessageDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:站内信消息表:system_notify_message:SystemNotifyMessageDelete")
		return &SystemNotifyMessageDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyMessageDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemNotifyMessage 查询单条数据
func (srv SrvSystemNotifyMessageServiceServer) SystemNotifyMessage(ctx context.Context, request *SystemNotifyMessageRequest) (*SystemNotifyMessageResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNotifyMessageResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := notify.SystemNotifyMessage(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:站内信消息表:system_notify_message:SystemNotifyMessage")
		return &SystemNotifyMessageResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyMessageResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemNotifyMessageProto(res),
	}, nil
}

// SystemNotifyMessageRecover 恢复数据
func (srv SrvSystemNotifyMessageServiceServer) SystemNotifyMessageRecover(ctx context.Context, request *SystemNotifyMessageRecoverRequest) (*SystemNotifyMessageRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNotifyMessageRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := notify.SystemNotifyMessageRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:站内信消息表:system_notify_message:SystemNotifyMessageRecover")
		return &SystemNotifyMessageRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyMessageRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemNotifyMessageServiceServer) SystemNotifyMessageList(ctx context.Context, request *SystemNotifyMessageListRequest) (*SystemNotifyMessageListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.TemplateType != nil {
		condition["templateType"] = request.GetTemplateType()
	}
	if request.ReadStatus != nil {
		condition["readStatus"] = request.GetReadStatus()
	}
	if request.BeginReadTime != nil {
		condition["beginReadTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetBeginReadTime()))
	}
	if request.FinishReadTime != nil {
		condition["finishReadTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetFinishReadTime()))
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
	list, err := notify.SystemNotifyMessageList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:站内信消息表:system_notify_message:SystemNotifyMessageList")
		return &SystemNotifyMessageListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemNotifyMessageObject
	for _, item := range list {
		res = append(res, SystemNotifyMessageProto(item))
	}
	return &SystemNotifyMessageListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemNotifyMessageListTotal 获取总数
func (srv SrvSystemNotifyMessageServiceServer) SystemNotifyMessageListTotal(ctx context.Context, request *SystemNotifyMessageListTotalRequest) (*SystemNotifyMessageTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.TemplateType != nil {
		condition["templateType"] = request.GetTemplateType()
	}
	if request.ReadStatus != nil {
		condition["readStatus"] = request.GetReadStatus()
	}
	if request.BeginReadTime != nil {
		condition["beginReadTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetBeginReadTime()))
	}
	if request.FinishReadTime != nil {
		condition["finishReadTime"] = util.Date("Y-m-d H:i:s", util.GrpcTime(request.GetFinishReadTime()))
	}

	// 获取数据集合
	total, err := notify.SystemNotifyMessageListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:站内信消息表:system_notify_message:SystemNotifyMessageListTotal")
		return &SystemNotifyMessageTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyMessageTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
