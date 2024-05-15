package app

import (
	"cloud/code"
	"cloud/module/pay/app"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// pay_app 支付应用信息

// SrvPayAppServiceServer 支付应用信息
type SrvPayAppServiceServer struct {
	UnimplementedPayAppServiceServer
	Server *xgrpc.Server
}

// PayAppCreate 创建数据
func (srv SrvPayAppServiceServer) PayAppCreate(ctx context.Context, request *PayAppCreateRequest) (*PayAppCreateResponse, error) {
	req := PayAppDao(request.GetData())
	data, err := app.PayAppCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:支付应用信息:pay_app:PayAppCreate")
		return &PayAppCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayAppCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// PayAppUpdate 更新数据
func (srv SrvPayAppServiceServer) PayAppUpdate(ctx context.Context, request *PayAppUpdateRequest) (*PayAppUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayAppUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := PayAppDao(request.GetData())
	_, err := app.PayAppUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:支付应用信息:pay_app:PayAppUpdate")
		return &PayAppUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayAppUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayAppDelete 删除数据
func (srv SrvPayAppServiceServer) PayAppDelete(ctx context.Context, request *PayAppDeleteRequest) (*PayAppDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayAppDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := app.PayAppDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付应用信息:pay_app:PayAppDelete")
		return &PayAppDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayAppDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayApp 查询单条数据
func (srv SrvPayAppServiceServer) PayApp(ctx context.Context, request *PayAppRequest) (*PayAppResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayAppResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := app.PayApp(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付应用信息:pay_app:PayApp")
		return &PayAppResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayAppResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayAppProto(res),
	}, nil
}

// PayAppRecover 恢复数据
func (srv SrvPayAppServiceServer) PayAppRecover(ctx context.Context, request *PayAppRecoverRequest) (*PayAppRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayAppRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := app.PayAppRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付应用信息:pay_app:PayAppRecover")
		return &PayAppRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayAppRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvPayAppServiceServer) PayAppList(ctx context.Context, request *PayAppListRequest) (*PayAppListResponse, error) {
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
	if request.Name != nil {
		condition["name"] = request.GetName()
	}

	// 获取数据集合
	list, err := app.PayAppList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:支付应用信息:pay_app:PayAppList")
		return &PayAppListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*PayAppObject
	for _, item := range list {
		res = append(res, PayAppProto(item))
	}
	return &PayAppListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
