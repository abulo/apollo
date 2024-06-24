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

// pay_wallet_recharge 会员钱包充值

// SrvPayWalletRechargeServiceServer 会员钱包充值
type SrvPayWalletRechargeServiceServer struct {
	UnimplementedPayWalletRechargeServiceServer
	Server *xgrpc.Server
}

// PayWalletRechargeCreate 创建数据
func (srv SrvPayWalletRechargeServiceServer) PayWalletRechargeCreate(ctx context.Context, request *PayWalletRechargeCreateRequest) (*PayWalletRechargeCreateResponse, error) {
	req := PayWalletRechargeDao(request.GetData())
	data, err := wallet.PayWalletRechargeCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:会员钱包充值:pay_wallet_recharge:PayWalletRechargeCreate")
		return &PayWalletRechargeCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRechargeCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// PayWalletRechargeUpdate 更新数据
func (srv SrvPayWalletRechargeServiceServer) PayWalletRechargeUpdate(ctx context.Context, request *PayWalletRechargeUpdateRequest) (*PayWalletRechargeUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletRechargeUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := PayWalletRechargeDao(request.GetData())
	_, err := wallet.PayWalletRechargeUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:会员钱包充值:pay_wallet_recharge:PayWalletRechargeUpdate")
		return &PayWalletRechargeUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRechargeUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayWalletRechargeDelete 删除数据
func (srv SrvPayWalletRechargeServiceServer) PayWalletRechargeDelete(ctx context.Context, request *PayWalletRechargeDeleteRequest) (*PayWalletRechargeDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletRechargeDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := wallet.PayWalletRechargeDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:会员钱包充值:pay_wallet_recharge:PayWalletRechargeDelete")
		return &PayWalletRechargeDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRechargeDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayWalletRecharge 查询单条数据
func (srv SrvPayWalletRechargeServiceServer) PayWalletRecharge(ctx context.Context, request *PayWalletRechargeRequest) (*PayWalletRechargeResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletRechargeResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := wallet.PayWalletRecharge(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:会员钱包充值:pay_wallet_recharge:PayWalletRecharge")
		return &PayWalletRechargeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRechargeResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayWalletRechargeProto(res),
	}, nil
}

// PayWalletRechargeRecover 恢复数据
func (srv SrvPayWalletRechargeServiceServer) PayWalletRechargeRecover(ctx context.Context, request *PayWalletRechargeRecoverRequest) (*PayWalletRechargeRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletRechargeRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := wallet.PayWalletRechargeRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:会员钱包充值:pay_wallet_recharge:PayWalletRechargeRecover")
		return &PayWalletRechargeRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRechargeRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvPayWalletRechargeServiceServer) PayWalletRechargeList(ctx context.Context, request *PayWalletRechargeListRequest) (*PayWalletRechargeListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.WalletId != nil {
		condition["walletId"] = request.GetWalletId()
	}
	if request.PackageId != nil {
		condition["packageId"] = request.GetPackageId()
	}
	if request.PayStatus != nil {
		condition["payStatus"] = request.GetPayStatus()
	}
	if request.PayOrderId != nil {
		condition["payOrderId"] = request.GetPayOrderId()
	}
	if request.PayChannelCode != nil {
		condition["payChannelCode"] = request.GetPayChannelCode()
	}
	if request.BeginPayTime != nil {
		condition["beginPayTime"] = request.GetBeginPayTime()
	}
	if request.FinishPayTime != nil {
		condition["finishPayTime"] = request.GetFinishPayTime()
	}
	if request.RefundStatus != nil {
		condition["refundStatus"] = request.GetRefundStatus()
	}
	if request.PayRefundId != nil {
		condition["payRefundId"] = request.GetPayRefundId()
	}
	if request.BeginRefundTime != nil {
		condition["beginRefundTime"] = request.GetBeginRefundTime()
	}
	if request.FinishRefundTime != nil {
		condition["finishRefundTime"] = request.GetFinishRefundTime()
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
	list, err := wallet.PayWalletRechargeList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:会员钱包充值:pay_wallet_recharge:PayWalletRechargeList")
		return &PayWalletRechargeListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*PayWalletRechargeObject
	for _, item := range list {
		res = append(res, PayWalletRechargeProto(item))
	}
	return &PayWalletRechargeListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// PayWalletRechargeListTotal 获取总数
func (srv SrvPayWalletRechargeServiceServer) PayWalletRechargeListTotal(ctx context.Context, request *PayWalletRechargeListTotalRequest) (*PayWalletRechargeTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.WalletId != nil {
		condition["walletId"] = request.GetWalletId()
	}
	if request.PackageId != nil {
		condition["packageId"] = request.GetPackageId()
	}
	if request.PayStatus != nil {
		condition["payStatus"] = request.GetPayStatus()
	}
	if request.PayOrderId != nil {
		condition["payOrderId"] = request.GetPayOrderId()
	}
	if request.PayChannelCode != nil {
		condition["payChannelCode"] = request.GetPayChannelCode()
	}
	if request.BeginPayTime != nil {
		condition["beginPayTime"] = request.GetBeginPayTime()
	}
	if request.FinishPayTime != nil {
		condition["finishPayTime"] = request.GetFinishPayTime()
	}
	if request.RefundStatus != nil {
		condition["refundStatus"] = request.GetRefundStatus()
	}
	if request.PayRefundId != nil {
		condition["payRefundId"] = request.GetPayRefundId()
	}
	if request.BeginRefundTime != nil {
		condition["beginRefundTime"] = request.GetBeginRefundTime()
	}
	if request.FinishRefundTime != nil {
		condition["finishRefundTime"] = request.GetFinishRefundTime()
	}

	// 获取数据集合
	total, err := wallet.PayWalletRechargeListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:会员钱包充值:pay_wallet_recharge:PayWalletRechargeListTotal")
		return &PayWalletRechargeTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRechargeTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
