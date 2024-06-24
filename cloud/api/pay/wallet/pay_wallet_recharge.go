package wallet

import (
	"context"
	"time"

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
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_wallet_recharge 会员钱包充值
// PayWalletRechargeDelete 删除数据
func PayWalletRechargeDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:会员钱包充值:pay_wallet_recharge:PayWalletRechargeDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}

	//链接服务
	clientWallet := wallet.NewPayWalletServiceClient(grpcClient)
	walletId := cast.ToInt64(newCtx.Param("walletId"))
	requestWallet := &wallet.PayWalletRequest{}
	requestWallet.Id = walletId
	// 执行服务
	_, err = clientWallet.PayWallet(ctx, requestWallet)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestWallet,
			"err": err,
		}).Error("GrpcCall:会员钱包表:pay_wallet:PayWallet")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}

	//链接服务
	client := wallet.NewPayWalletRechargeServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))

	requestItem := &wallet.PayWalletRechargeRequest{}
	requestItem.Id = id
	// 执行服务
	rechargeItem, err := client.PayWalletRecharge(ctx, requestItem)
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
	item := wallet.PayWalletRechargeDao(rechargeItem.GetData())
	if item.WalletId != proto.Int64(walletId) {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}

	request := &wallet.PayWalletRechargeDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayWalletRechargeDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:会员钱包充值:pay_wallet_recharge:PayWalletRechargeDelete")
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

// PayWalletRecharge 查询单条数据
func PayWalletRecharge(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:会员钱包充值:pay_wallet_recharge:PayWalletRecharge")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}

	//链接服务
	clientWallet := wallet.NewPayWalletServiceClient(grpcClient)
	walletId := cast.ToInt64(newCtx.Param("walletId"))
	requestWallet := &wallet.PayWalletRequest{}
	requestWallet.Id = walletId
	// 执行服务
	_, err = clientWallet.PayWallet(ctx, requestWallet)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestWallet,
			"err": err,
		}).Error("GrpcCall:会员钱包表:pay_wallet:PayWallet")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}

	//链接服务
	client := wallet.NewPayWalletRechargeServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &wallet.PayWalletRechargeRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayWalletRecharge(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:会员钱包充值:pay_wallet_recharge:PayWalletRecharge")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}

	item := wallet.PayWalletRechargeDao(res.GetData())
	if item.WalletId != proto.Int64(walletId) {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": item,
	})
}

// PayWalletRechargeRecover 恢复数据
func PayWalletRechargeRecover(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:会员钱包充值:pay_wallet_recharge:PayWalletRechargeRecover")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	clientWallet := wallet.NewPayWalletServiceClient(grpcClient)
	walletId := cast.ToInt64(newCtx.Param("walletId"))
	requestWallet := &wallet.PayWalletRequest{}
	requestWallet.Id = walletId
	// 执行服务
	_, err = clientWallet.PayWallet(ctx, requestWallet)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestWallet,
			"err": err,
		}).Error("GrpcCall:会员钱包表:pay_wallet:PayWallet")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}

	//链接服务
	client := wallet.NewPayWalletRechargeServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))

	requestItem := &wallet.PayWalletRechargeRequest{}
	requestItem.Id = id
	// 执行服务
	rechargeItem, err := client.PayWalletRecharge(ctx, requestItem)
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
	item := wallet.PayWalletRechargeDao(rechargeItem.GetData())
	if item.WalletId != proto.Int64(walletId) {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}

	request := &wallet.PayWalletRechargeRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayWalletRechargeRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:会员钱包充值:pay_wallet_recharge:PayWalletRechargeRecover")
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

// PayWalletRechargeList 列表数据
func PayWalletRechargeList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:会员钱包充值:pay_wallet_recharge:PayWalletRechargeList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	clientWallet := wallet.NewPayWalletServiceClient(grpcClient)
	walletId := cast.ToInt64(newCtx.Param("walletId"))
	requestWallet := &wallet.PayWalletRequest{}
	requestWallet.Id = walletId
	// 执行服务
	_, err = clientWallet.PayWallet(ctx, requestWallet)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestWallet,
			"err": err,
		}).Error("GrpcCall:会员钱包表:pay_wallet:PayWallet")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	//链接服务
	client := wallet.NewPayWalletRechargeServiceClient(grpcClient)
	// 构造查询条件
	request := &wallet.PayWalletRechargeListRequest{}
	requestTotal := &wallet.PayWalletRechargeListTotalRequest{}
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
	if val, ok := newCtx.GetQuery("packageId"); ok {
		request.PackageId = proto.Int64(cast.ToInt64(val))      // 充值套餐编号
		requestTotal.PackageId = proto.Int64(cast.ToInt64(val)) // 充值套餐编号
	}
	if val, ok := newCtx.GetQuery("payStatus"); ok {
		request.PayStatus = proto.Int32(cast.ToInt32(val))      // 是否已支付：[0:未支付 1:已经支付过]
		requestTotal.PayStatus = proto.Int32(cast.ToInt32(val)) // 是否已支付：[0:未支付 1:已经支付过]
	}
	if val, ok := newCtx.GetQuery("payOrderId"); ok {
		request.PayOrderId = proto.Int64(cast.ToInt64(val))      // 支付订单编号
		requestTotal.PayOrderId = proto.Int64(cast.ToInt64(val)) // 支付订单编号
	}
	if val, ok := newCtx.GetQuery("payChannelCode"); ok {
		request.PayChannelCode = proto.String(val)      // 支付成功的支付渠道
		requestTotal.PayChannelCode = proto.String(val) // 支付成功的支付渠道
	}
	if val, ok := newCtx.GetQuery("beginPayTime"); ok {
		request.BeginPayTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 订单支付时间
		requestTotal.BeginPayTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 订单支付时间
	}
	if val, ok := newCtx.GetQuery("finishPayTime"); ok {
		request.FinishPayTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 订单支付完成时间
		requestTotal.FinishPayTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 订单支付完成时间
	}
	if val, ok := newCtx.GetQuery("refundStatus"); ok {
		request.RefundStatus = proto.Int32(cast.ToInt32(val))      // 退款状态
		requestTotal.RefundStatus = proto.Int32(cast.ToInt32(val)) // 退款状态
	}
	if val, ok := newCtx.GetQuery("payRefundId"); ok {
		request.PayRefundId = proto.Int64(cast.ToInt64(val))      // 支付退款单编号
		requestTotal.PayRefundId = proto.Int64(cast.ToInt64(val)) // 支付退款单编号
	}
	if val, ok := newCtx.GetQuery("beginRefundTime"); ok {
		request.BeginRefundTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 退款时间
		requestTotal.BeginRefundTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 退款时间
	}
	if val, ok := newCtx.GetQuery("finishRefundTime"); ok {
		request.FinishRefundTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 退款完成时间
		requestTotal.FinishRefundTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 退款完成时间
	}

	// 执行服务,获取数据量
	resTotal, err := client.PayWalletRechargeListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:会员钱包充值:pay_wallet_recharge:PayWalletRechargeList")
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
	res, err := client.PayWalletRechargeList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:会员钱包充值:pay_wallet_recharge:PayWalletRechargeList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.PayWalletRecharge
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, wallet.PayWalletRechargeDao(item))
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
