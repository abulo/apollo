package mail

import (
	"cloud/code"
	"cloud/module/system/mail"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_mail_template 邮件模版表

// SrvSystemMailTemplateServiceServer 邮件模版表
type SrvSystemMailTemplateServiceServer struct {
	UnimplementedSystemMailTemplateServiceServer
	Server *xgrpc.Server
}

// SystemMailTemplateCreate 创建数据
func (srv SrvSystemMailTemplateServiceServer) SystemMailTemplateCreate(ctx context.Context, request *SystemMailTemplateCreateRequest) (*SystemMailTemplateCreateResponse, error) {
	req := SystemMailTemplateDao(request.GetData())
	data, err := mail.SystemMailTemplateCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:邮件模版表:system_mail_template:SystemMailTemplateCreate")
		return &SystemMailTemplateCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailTemplateCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemMailTemplateUpdate 更新数据
func (srv SrvSystemMailTemplateServiceServer) SystemMailTemplateUpdate(ctx context.Context, request *SystemMailTemplateUpdateRequest) (*SystemMailTemplateUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailTemplateUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemMailTemplateDao(request.GetData())
	_, err := mail.SystemMailTemplateUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:邮件模版表:system_mail_template:SystemMailTemplateUpdate")
		return &SystemMailTemplateUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailTemplateUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemMailTemplateDelete 删除数据
func (srv SrvSystemMailTemplateServiceServer) SystemMailTemplateDelete(ctx context.Context, request *SystemMailTemplateDeleteRequest) (*SystemMailTemplateDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailTemplateDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := mail.SystemMailTemplateDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:邮件模版表:system_mail_template:SystemMailTemplateDelete")
		return &SystemMailTemplateDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailTemplateDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemMailTemplate 查询单条数据
func (srv SrvSystemMailTemplateServiceServer) SystemMailTemplate(ctx context.Context, request *SystemMailTemplateRequest) (*SystemMailTemplateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailTemplateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := mail.SystemMailTemplate(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:邮件模版表:system_mail_template:SystemMailTemplate")
		return &SystemMailTemplateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailTemplateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemMailTemplateProto(res),
	}, nil
}

// SystemMailTemplateRecover 恢复数据
func (srv SrvSystemMailTemplateServiceServer) SystemMailTemplateRecover(ctx context.Context, request *SystemMailTemplateRecoverRequest) (*SystemMailTemplateRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailTemplateRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := mail.SystemMailTemplateRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:邮件模版表:system_mail_template:SystemMailTemplateRecover")
		return &SystemMailTemplateRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailTemplateRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemMailTemplateServiceServer) SystemMailTemplateList(ctx context.Context, request *SystemMailTemplateListRequest) (*SystemMailTemplateListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Title != nil {
		condition["title"] = request.GetTitle()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}
	if request.Code != nil {
		condition["code"] = request.GetCode()
	}
	if request.AccountId != nil {
		condition["accountId"] = request.GetAccountId()
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
	list, err := mail.SystemMailTemplateList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:邮件模版表:system_mail_template:SystemMailTemplateList")
		return &SystemMailTemplateListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemMailTemplateObject
	for _, item := range list {
		res = append(res, SystemMailTemplateProto(item))
	}
	return &SystemMailTemplateListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemMailTemplateListTotal 获取总数
func (srv SrvSystemMailTemplateServiceServer) SystemMailTemplateListTotal(ctx context.Context, request *SystemMailTemplateListTotalRequest) (*SystemMailTemplateTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Title != nil {
		condition["title"] = request.GetTitle()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}
	if request.Code != nil {
		condition["code"] = request.GetCode()
	}
	if request.AccountId != nil {
		condition["accountId"] = request.GetAccountId()
	}

	// 获取数据集合
	total, err := mail.SystemMailTemplateListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:邮件模版表:system_mail_template:SystemMailTemplateListTotal")
		return &SystemMailTemplateTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailTemplateTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
