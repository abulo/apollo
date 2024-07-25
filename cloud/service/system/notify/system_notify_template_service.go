package notify

import (
	"cloud/code"
	"cloud/module/system/notify"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_notify_template 站内信模板表

// SrvSystemNotifyTemplateServiceServer 站内信模板表
type SrvSystemNotifyTemplateServiceServer struct {
	UnimplementedSystemNotifyTemplateServiceServer
	Server *xgrpc.Server
}

// SystemNotifyTemplateCreate 创建数据
func (srv SrvSystemNotifyTemplateServiceServer) SystemNotifyTemplateCreate(ctx context.Context, request *SystemNotifyTemplateCreateRequest) (*SystemNotifyTemplateCreateResponse, error) {
	req := SystemNotifyTemplateDao(request.GetData())
	data, err := notify.SystemNotifyTemplateCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:站内信模板表:system_notify_template:SystemNotifyTemplateCreate")
		return &SystemNotifyTemplateCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyTemplateCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemNotifyTemplateUpdate 更新数据
func (srv SrvSystemNotifyTemplateServiceServer) SystemNotifyTemplateUpdate(ctx context.Context, request *SystemNotifyTemplateUpdateRequest) (*SystemNotifyTemplateUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNotifyTemplateUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemNotifyTemplateDao(request.GetData())
	_, err := notify.SystemNotifyTemplateUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:站内信模板表:system_notify_template:SystemNotifyTemplateUpdate")
		return &SystemNotifyTemplateUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyTemplateUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemNotifyTemplateDelete 删除数据
func (srv SrvSystemNotifyTemplateServiceServer) SystemNotifyTemplateDelete(ctx context.Context, request *SystemNotifyTemplateDeleteRequest) (*SystemNotifyTemplateDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNotifyTemplateDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := notify.SystemNotifyTemplateDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:站内信模板表:system_notify_template:SystemNotifyTemplateDelete")
		return &SystemNotifyTemplateDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyTemplateDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemNotifyTemplate 查询单条数据
func (srv SrvSystemNotifyTemplateServiceServer) SystemNotifyTemplate(ctx context.Context, request *SystemNotifyTemplateRequest) (*SystemNotifyTemplateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNotifyTemplateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := notify.SystemNotifyTemplate(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:站内信模板表:system_notify_template:SystemNotifyTemplate")
		return &SystemNotifyTemplateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyTemplateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemNotifyTemplateProto(res),
	}, nil
}

// SystemNotifyTemplateRecover 恢复数据
func (srv SrvSystemNotifyTemplateServiceServer) SystemNotifyTemplateRecover(ctx context.Context, request *SystemNotifyTemplateRecoverRequest) (*SystemNotifyTemplateRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNotifyTemplateRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := notify.SystemNotifyTemplateRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:站内信模板表:system_notify_template:SystemNotifyTemplateRecover")
		return &SystemNotifyTemplateRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyTemplateRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemNotifyTemplateDrop 清理数据
func (srv SrvSystemNotifyTemplateServiceServer) SystemNotifyTemplateDrop(ctx context.Context, request *SystemNotifyTemplateDropRequest) (*SystemNotifyTemplateDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNotifyTemplateDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := notify.SystemNotifyTemplateDrop(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:站内信模板表:system_notify_template:SystemNotifyTemplateDrop")
		return &SystemNotifyTemplateDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyTemplateDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemNotifyTemplateServiceServer) SystemNotifyTemplateList(ctx context.Context, request *SystemNotifyTemplateListRequest) (*SystemNotifyTemplateListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
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
	list, err := notify.SystemNotifyTemplateList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:站内信模板表:system_notify_template:SystemNotifyTemplateList")
		return &SystemNotifyTemplateListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemNotifyTemplateObject
	for _, item := range list {
		res = append(res, SystemNotifyTemplateProto(item))
	}
	return &SystemNotifyTemplateListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemNotifyTemplateListTotal 获取总数
func (srv SrvSystemNotifyTemplateServiceServer) SystemNotifyTemplateListTotal(ctx context.Context, request *SystemNotifyTemplateListTotalRequest) (*SystemNotifyTemplateTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}

	// 获取数据集合
	total, err := notify.SystemNotifyTemplateListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:站内信模板表:system_notify_template:SystemNotifyTemplateListTotal")
		return &SystemNotifyTemplateTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNotifyTemplateTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
