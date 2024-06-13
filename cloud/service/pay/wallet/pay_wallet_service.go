package wallet

import (
	"cloud/code"
	"cloud/module/pay/wallet"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// pay_wallet 会员钱包表

// SrvPayWalletServiceServer 会员钱包表
type SrvPayWalletServiceServer struct {
	UnimplementedPayWalletServiceServer
	Server *xgrpc.Server
}

// PayWalletCreate 创建数据
func (srv SrvPayWalletServiceServer) PayWalletCreate(ctx context.Context, request *PayWalletCreateRequest) (*PayWalletCreateResponse, error) {
	req := PayWalletDao(request.GetData())
	data, err := wallet.PayWalletCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:会员钱包表:pay_wallet:PayWalletCreate")
		return &PayWalletCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// PayWalletUpdate 更新数据
func (srv SrvPayWalletServiceServer) PayWalletUpdate(ctx context.Context, request *PayWalletUpdateRequest) (*PayWalletUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := PayWalletDao(request.GetData())
	_, err := wallet.PayWalletUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:会员钱包表:pay_wallet:PayWalletUpdate")
		return &PayWalletUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayWalletDelete 删除数据
func (srv SrvPayWalletServiceServer) PayWalletDelete(ctx context.Context, request *PayWalletDeleteRequest) (*PayWalletDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := wallet.PayWalletDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:会员钱包表:pay_wallet:PayWalletDelete")
		return &PayWalletDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayWallet 查询单条数据
func (srv SrvPayWalletServiceServer) PayWallet(ctx context.Context, request *PayWalletRequest) (*PayWalletResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := wallet.PayWallet(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:会员钱包表:pay_wallet:PayWallet")
		return &PayWalletResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayWalletProto(res),
	}, nil
}

// PayWalletRecover 恢复数据
func (srv SrvPayWalletServiceServer) PayWalletRecover(ctx context.Context, request *PayWalletRecoverRequest) (*PayWalletRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayWalletRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := wallet.PayWalletRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:会员钱包表:pay_wallet:PayWalletRecover")
		return &PayWalletRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayWalletUser 查询单条数据
func (srv SrvPayWalletServiceServer) PayWalletUser(ctx context.Context, request *PayWalletUserRequest) (*PayWalletUserResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserType != nil {
		condition["userType"] = request.GetUserType()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}

	if util.Empty(condition) {
		err := errors.New("condition is empty")
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:会员钱包表:pay_wallet:PayWalletUser")
		return &PayWalletUserResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := wallet.PayWalletUser(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:会员钱包表:pay_wallet:PayWalletUser")
		return &PayWalletUserResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletUserResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayWalletProto(res),
	}, nil
}
func (srv SrvPayWalletServiceServer) PayWalletList(ctx context.Context, request *PayWalletListRequest) (*PayWalletListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserType != nil {
		condition["userType"] = request.GetUserType()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
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
	list, err := wallet.PayWalletList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:会员钱包表:pay_wallet:PayWalletList")
		return &PayWalletListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*PayWalletObject
	for _, item := range list {
		res = append(res, PayWalletProto(item))
	}
	return &PayWalletListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// PayWalletListTotal 获取总数
func (srv SrvPayWalletServiceServer) PayWalletListTotal(ctx context.Context, request *PayWalletListTotalRequest) (*PayWalletTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserType != nil {
		condition["userType"] = request.GetUserType()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}

	// 获取数据集合
	total, err := wallet.PayWalletListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:会员钱包表:pay_wallet:PayWalletListTotal")
		return &PayWalletTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayWalletTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
