package wallet

import (
	"cloud/code"
	"cloud/module/pay/wallet"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// pay_wallet_recharge_package 充值套餐表

// SrvPayWalletRechargePackageServiceServer 充值套餐表
type SrvPayWalletRechargePackageServiceServer struct {
	UnimplementedPayWalletRechargePackageServiceServer
	Server *xgrpc.Server
}

// PayWalletRechargePackageCreate 创建数据
func (srv SrvPayWalletRechargePackageServiceServer) PayWalletRechargePackageCreate(ctx context.Context, request *PayWalletRechargePackageCreateRequest) (*PayWalletRechargePackageCreateResponse, error) {
	req := PayWalletRechargePackageDao(request.GetData())
	data, err := wallet.PayWalletRechargePackageCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageCreate")
		return &PayWalletRechargePackageCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRechargePackageCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// PayWalletRechargePackageUpdate 更新数据
func (srv SrvPayWalletRechargePackageServiceServer) PayWalletRechargePackageUpdate(ctx context.Context, request *PayWalletRechargePackageUpdateRequest) (*PayWalletRechargePackageUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletRechargePackageUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := PayWalletRechargePackageDao(request.GetData())
	_, err := wallet.PayWalletRechargePackageUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageUpdate")
		return &PayWalletRechargePackageUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRechargePackageUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayWalletRechargePackageDelete 删除数据
func (srv SrvPayWalletRechargePackageServiceServer) PayWalletRechargePackageDelete(ctx context.Context, request *PayWalletRechargePackageDeleteRequest) (*PayWalletRechargePackageDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletRechargePackageDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := wallet.PayWalletRechargePackageDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageDelete")
		return &PayWalletRechargePackageDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRechargePackageDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayWalletRechargePackage 查询单条数据
func (srv SrvPayWalletRechargePackageServiceServer) PayWalletRechargePackage(ctx context.Context, request *PayWalletRechargePackageRequest) (*PayWalletRechargePackageResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletRechargePackageResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := wallet.PayWalletRechargePackage(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackage")
		return &PayWalletRechargePackageResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRechargePackageResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayWalletRechargePackageProto(res),
	}, nil
}

// PayWalletRechargePackageRecover 恢复数据
func (srv SrvPayWalletRechargePackageServiceServer) PayWalletRechargePackageRecover(ctx context.Context, request *PayWalletRechargePackageRecoverRequest) (*PayWalletRechargePackageRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletRechargePackageRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := wallet.PayWalletRechargePackageRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageRecover")
		return &PayWalletRechargePackageRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRechargePackageRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvPayWalletRechargePackageServiceServer) PayWalletRechargePackageList(ctx context.Context, request *PayWalletRechargePackageListRequest) (*PayWalletRechargePackageListResponse, error) {
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
	list, err := wallet.PayWalletRechargePackageList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageList")
		return &PayWalletRechargePackageListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*PayWalletRechargePackageObject
	for _, item := range list {
		res = append(res, PayWalletRechargePackageProto(item))
	}
	return &PayWalletRechargePackageListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// PayWalletRechargePackageListTotal 获取总数
func (srv SrvPayWalletRechargePackageServiceServer) PayWalletRechargePackageListTotal(ctx context.Context, request *PayWalletRechargePackageListTotalRequest) (*PayWalletRechargePackageTotalResponse, error) {
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
	total, err := wallet.PayWalletRechargePackageListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageListTotal")
		return &PayWalletRechargePackageTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRechargePackageTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
