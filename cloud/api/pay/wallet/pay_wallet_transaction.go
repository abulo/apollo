package wallet

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/pagination"
	"cloud/service/pay/wallet"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// pay_wallet_transaction 会员钱包流水表
// PayWalletTransactionDelete 删除数据
func PayWalletTransactionDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	walletId := cast.ToInt64(newCtx.Param("walletId"))
	if _, err := PayWalletItem(ctx, newCtx, walletId); err != nil {
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	//链接服务
	client := wallet.NewPayWalletTransactionServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	requestItem := &wallet.PayWalletTransactionRequest{}
	requestItem.Id = id
	// 执行服务
	transactionItem, err := client.PayWalletTransaction(ctx, requestItem)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestItem,
			"err": err,
		}).Error("GrpcCall:会员钱包充值:pay_wallet_recharge:PayWalletRecharge")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	item := wallet.PayWalletTransactionDao(transactionItem.GetData())
	if cast.ToInt64(item.WalletId) != walletId {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request := &wallet.PayWalletTransactionDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayWalletTransactionDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionDelete")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// PayWalletTransaction 查询单条数据
func PayWalletTransaction(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:会员钱包流水表:pay_wallet_transaction:PayWalletTransaction")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	walletId := cast.ToInt64(newCtx.Param("walletId"))
	if _, err := PayWalletItem(ctx, newCtx, walletId); err != nil {
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	//链接服务
	client := wallet.NewPayWalletTransactionServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))

	requestItem := &wallet.PayWalletTransactionRequest{}
	requestItem.Id = id
	// 执行服务
	transactionItem, err := client.PayWalletTransaction(ctx, requestItem)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestItem,
			"err": err,
		}).Error("GrpcCall:会员钱包充值:pay_wallet_recharge:PayWalletRecharge")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	item := wallet.PayWalletTransactionDao(transactionItem.GetData())
	if cast.ToInt64(item.WalletId) != walletId {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}

	request := &wallet.PayWalletTransactionRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayWalletTransaction(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:会员钱包流水表:pay_wallet_transaction:PayWalletTransaction")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": wallet.PayWalletTransactionDao(res.GetData()),
	})
}

// PayWalletTransactionRecover 恢复数据
func PayWalletTransactionRecover(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionRecover")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	walletId := cast.ToInt64(newCtx.Param("walletId"))
	if _, err := PayWalletItem(ctx, newCtx, walletId); err != nil {
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	//链接服务
	client := wallet.NewPayWalletTransactionServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))

	requestItem := &wallet.PayWalletTransactionRequest{}
	requestItem.Id = id
	// 执行服务
	transactionItem, err := client.PayWalletTransaction(ctx, requestItem)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestItem,
			"err": err,
		}).Error("GrpcCall:会员钱包充值:pay_wallet_recharge:PayWalletRecharge")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	item := wallet.PayWalletTransactionDao(transactionItem.GetData())
	if cast.ToInt64(item.WalletId) != walletId {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}

	request := &wallet.PayWalletTransactionRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayWalletTransactionRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionRecover")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// PayWalletTransactionList 列表数据
func PayWalletTransactionList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	//链接服务
	walletId := cast.ToInt64(newCtx.Param("walletId"))
	if _, err := PayWalletItem(ctx, newCtx, walletId); err != nil {
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	//链接服务
	client := wallet.NewPayWalletTransactionServiceClient(grpcClient)
	// 构造查询条件
	request := &wallet.PayWalletTransactionListRequest{}
	requestTotal := &wallet.PayWalletTransactionListTotalRequest{}

	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))      // 租户ID
	requestTotal.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                                 // 删除状态
	requestTotal.Deleted = proto.Int32(0)                            // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
			requestTotal.Deleted = nil
		}
	}
	request.WalletId = proto.Int64(walletId)      // 会员钱包id
	requestTotal.WalletId = proto.Int64(walletId) // 会员钱包id
	if val, ok := newCtx.GetQuery("bizType"); ok {
		request.BizType = proto.Int32(cast.ToInt32(val))      // 关联类型
		requestTotal.BizType = proto.Int32(cast.ToInt32(val)) // 关联类型
	}
	if val, ok := newCtx.GetQuery("bizId"); ok {
		request.BizId = proto.String(val)      // 关联业务编号
		requestTotal.BizId = proto.String(val) // 关联业务编号
	}
	if val, ok := newCtx.GetQuery("no"); ok {
		request.No = proto.String(val)      // 流水号
		requestTotal.No = proto.String(val) // 流水号
	}
	if val, ok := newCtx.GetQuery("title"); ok {
		request.Title = proto.String(val)      // 流水标题
		requestTotal.Title = proto.String(val) // 流水标题
	}

	// 执行服务,获取数据量
	resTotal, err := client.PayWalletTransactionListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var total int64
	paginationRequest := &pagination.PaginationRequest{}
	paginationRequest.PageNum = proto.Int64(cast.ToInt64(newCtx.Query("pageNum")))
	paginationRequest.PageSize = proto.Int64(cast.ToInt64(newCtx.Query("pageSize")))
	request.Pagination = paginationRequest
	if resTotal.GetCode() == code.Success {
		total = resTotal.GetData()
	}
	// 执行服务
	res, err := client.PayWalletTransactionList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:会员钱包流水表:pay_wallet_transaction:PayWalletTransactionList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.PayWalletTransaction
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, wallet.PayWalletTransactionDao(item))
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"total":    total,
			"list":     list,
			"pageNum":  paginationRequest.PageNum,
			"pageSize": paginationRequest.PageSize,
		},
	})
}
