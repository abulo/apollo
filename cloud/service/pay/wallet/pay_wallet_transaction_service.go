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

// pay_wallet_transaction 会员钱包流水表

// SrvPayWalletTransactionServiceServer 会员钱包流水表
type SrvPayWalletTransactionServiceServer struct {
	UnimplementedPayWalletTransactionServiceServer
	Server *xgrpc.Server
}

// PayWalletTransactionCreate 创建数据
func (srv SrvPayWalletTransactionServiceServer) PayWalletTransactionCreate(ctx context.Context, request *PayWalletTransactionCreateRequest) (*PayWalletTransactionCreateResponse, error) {
	req := PayWalletTransactionDao(request.GetData())
	data, err := wallet.PayWalletTransactionCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionCreate")
		return &PayWalletTransactionCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletTransactionCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// PayWalletTransactionUpdate 更新数据
func (srv SrvPayWalletTransactionServiceServer) PayWalletTransactionUpdate(ctx context.Context, request *PayWalletTransactionUpdateRequest) (*PayWalletTransactionUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletTransactionUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := PayWalletTransactionDao(request.GetData())
	_, err := wallet.PayWalletTransactionUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionUpdate")
		return &PayWalletTransactionUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletTransactionUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayWalletTransactionDelete 删除数据
func (srv SrvPayWalletTransactionServiceServer) PayWalletTransactionDelete(ctx context.Context, request *PayWalletTransactionDeleteRequest) (*PayWalletTransactionDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletTransactionDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := wallet.PayWalletTransactionDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionDelete")
		return &PayWalletTransactionDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletTransactionDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayWalletTransaction 查询单条数据
func (srv SrvPayWalletTransactionServiceServer) PayWalletTransaction(ctx context.Context, request *PayWalletTransactionRequest) (*PayWalletTransactionResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletTransactionResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := wallet.PayWalletTransaction(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:会员钱包流水表:pay_wallet_transaction:PayWalletTransaction")
		return &PayWalletTransactionResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletTransactionResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayWalletTransactionProto(res),
	}, nil
}

// PayWalletTransactionRecover 恢复数据
func (srv SrvPayWalletTransactionServiceServer) PayWalletTransactionRecover(ctx context.Context, request *PayWalletTransactionRecoverRequest) (*PayWalletTransactionRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletTransactionRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := wallet.PayWalletTransactionRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionRecover")
		return &PayWalletTransactionRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletTransactionRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvPayWalletTransactionServiceServer) PayWalletTransactionList(ctx context.Context, request *PayWalletTransactionListRequest) (*PayWalletTransactionListResponse, error) {
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
	if request.BizType != nil {
		condition["bizType"] = request.GetBizType()
	}
	if request.BizId != nil {
		condition["bizId"] = request.GetBizId()
	}
	if request.No != nil {
		condition["no"] = request.GetNo()
	}
	if request.Title != nil {
		condition["title"] = request.GetTitle()
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
	list, err := wallet.PayWalletTransactionList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionList")
		return &PayWalletTransactionListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*PayWalletTransactionObject
	for _, item := range list {
		res = append(res, PayWalletTransactionProto(item))
	}
	return &PayWalletTransactionListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// PayWalletTransactionListTotal 获取总数
func (srv SrvPayWalletTransactionServiceServer) PayWalletTransactionListTotal(ctx context.Context, request *PayWalletTransactionListTotalRequest) (*PayWalletTransactionTotalResponse, error) {
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
	if request.BizType != nil {
		condition["bizType"] = request.GetBizType()
	}
	if request.BizId != nil {
		condition["bizId"] = request.GetBizId()
	}
	if request.No != nil {
		condition["no"] = request.GetNo()
	}
	if request.Title != nil {
		condition["title"] = request.GetTitle()
	}

	// 获取数据集合
	total, err := wallet.PayWalletTransactionListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionListTotal")
		return &PayWalletTransactionTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletTransactionTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
