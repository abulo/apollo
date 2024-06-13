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

// system_mail_account 邮箱账号表

// SrvSystemMailAccountServiceServer 邮箱账号表
type SrvSystemMailAccountServiceServer struct {
	UnimplementedSystemMailAccountServiceServer
	Server *xgrpc.Server
}

// SystemMailAccountCreate 创建数据
func (srv SrvSystemMailAccountServiceServer) SystemMailAccountCreate(ctx context.Context, request *SystemMailAccountCreateRequest) (*SystemMailAccountCreateResponse, error) {
	req := SystemMailAccountDao(request.GetData())
	data, err := mail.SystemMailAccountCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:邮箱账号表:system_mail_account:SystemMailAccountCreate")
		return &SystemMailAccountCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailAccountCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemMailAccountUpdate 更新数据
func (srv SrvSystemMailAccountServiceServer) SystemMailAccountUpdate(ctx context.Context, request *SystemMailAccountUpdateRequest) (*SystemMailAccountUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailAccountUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemMailAccountDao(request.GetData())
	_, err := mail.SystemMailAccountUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:邮箱账号表:system_mail_account:SystemMailAccountUpdate")
		return &SystemMailAccountUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailAccountUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemMailAccountDelete 删除数据
func (srv SrvSystemMailAccountServiceServer) SystemMailAccountDelete(ctx context.Context, request *SystemMailAccountDeleteRequest) (*SystemMailAccountDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailAccountDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := mail.SystemMailAccountDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:邮箱账号表:system_mail_account:SystemMailAccountDelete")
		return &SystemMailAccountDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailAccountDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemMailAccount 查询单条数据
func (srv SrvSystemMailAccountServiceServer) SystemMailAccount(ctx context.Context, request *SystemMailAccountRequest) (*SystemMailAccountResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailAccountResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := mail.SystemMailAccount(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:邮箱账号表:system_mail_account:SystemMailAccount")
		return &SystemMailAccountResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailAccountResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemMailAccountProto(res),
	}, nil
}

// SystemMailAccountRecover 恢复数据
func (srv SrvSystemMailAccountServiceServer) SystemMailAccountRecover(ctx context.Context, request *SystemMailAccountRecoverRequest) (*SystemMailAccountRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemMailAccountRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := mail.SystemMailAccountRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:邮箱账号表:system_mail_account:SystemMailAccountRecover")
		return &SystemMailAccountRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailAccountRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemMailAccountServiceServer) SystemMailAccountList(ctx context.Context, request *SystemMailAccountListRequest) (*SystemMailAccountListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Mail != nil {
		condition["mail"] = request.GetMail()
	}
	if request.Username != nil {
		condition["username"] = request.GetUsername()
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
	list, err := mail.SystemMailAccountList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:邮箱账号表:system_mail_account:SystemMailAccountList")
		return &SystemMailAccountListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemMailAccountObject
	for _, item := range list {
		res = append(res, SystemMailAccountProto(item))
	}
	return &SystemMailAccountListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemMailAccountListTotal 获取总数
func (srv SrvSystemMailAccountServiceServer) SystemMailAccountListTotal(ctx context.Context, request *SystemMailAccountListTotalRequest) (*SystemMailAccountTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Mail != nil {
		condition["mail"] = request.GetMail()
	}
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}

	// 获取数据集合
	total, err := mail.SystemMailAccountListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:邮箱账号表:system_mail_account:SystemMailAccountListTotal")
		return &SystemMailAccountTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMailAccountTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
