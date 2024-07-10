package order

import (
	"context"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/pay/order"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// pay_order 支付订单

// PayOrder 查询单条数据
func PayOrderItem(ctx context.Context, newCtx *app.RequestContext) (*order.PayOrderResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付订单:pay_order:PayOrder")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := order.NewPayOrderServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &order.PayOrderRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayOrder(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付订单:pay_order:PayOrder")
		return nil, err
	}
	tenantId := cast.ToInt64(newCtx.GetInt64("tenantId"))
	data := order.PayOrderDao(res.GetData())
	if cast.ToInt64(data.TenantId) != tenantId {
		return nil, status.Error(code.ConvertToGrpc(code.NoPermission), code.StatusText(code.NoPermission))
	}
	return res, nil
}

// PayOrderCreate 创建数据
func PayOrderCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付订单:pay_order:PayOrderCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := order.NewPayOrderServiceClient(grpcClient)
	request := &order.PayOrderCreateRequest{}
	// 数据绑定
	var reqInfo dao.PayOrder
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Deleted = proto.Int32(0)
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户
	reqInfo.Creator = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	request.Data = order.PayOrderProto(reqInfo)
	// 执行服务
	res, err := client.PayOrderCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付订单:pay_order:PayOrderCreate")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// PayOrderUpdate 更新数据
func PayOrderUpdate(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := PayOrderItem(ctx, newCtx); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付订单:pay_order:PayOrderUpdate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := order.NewPayOrderServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &order.PayOrderUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.PayOrder
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户
	reqInfo.Updater = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	reqInfo.Creator = null.StringFromPtr(nil)
	reqInfo.CreateTime = null.DateTimeFromPtr(nil)
	request.Data = order.PayOrderProto(reqInfo)
	// 执行服务
	res, err := client.PayOrderUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付订单:pay_order:PayOrderUpdate")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// PayOrderDelete 删除数据
func PayOrderDelete(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := PayOrderItem(ctx, newCtx); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付订单:pay_order:PayOrderDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := order.NewPayOrderServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &order.PayOrderDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayOrderDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付订单:pay_order:PayOrderDelete")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// PayOrder 查询单条数据
func PayOrder(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付订单:pay_order:PayOrder")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	res, err := PayOrderItem(ctx, newCtx)
	if err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	orderItem := order.PayOrderDao(res.GetData())
	clientExtension := order.NewPayOrderExtensionServiceClient(grpcClient)
	requestExtension := &order.PayOrderExtensionRequest{}
	requestExtension.Id = cast.ToInt64(orderItem.ExtensionId)
	// 执行服务
	if res, err := clientExtension.PayOrderExtension(ctx, requestExtension); err == nil {
		orderItem.Extension = order.PayOrderExtensionDao(res.GetData())
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": orderItem,
	})
}

// PayOrderRecover 恢复数据
func PayOrderRecover(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := PayOrderItem(ctx, newCtx); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付订单:pay_order:PayOrderRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := order.NewPayOrderServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &order.PayOrderRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayOrderRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付订单:pay_order:PayOrderRecover")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// PayOrderList 列表数据
func PayOrderList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付订单:pay_order:PayOrderList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := order.NewPayOrderServiceClient(grpcClient)
	// 构造查询条件
	request := &order.PayOrderListRequest{}
	requestTotal := &order.PayOrderListTotalRequest{}

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
	if val, ok := newCtx.GetQuery("appId"); ok {
		request.AppId = proto.Int64(cast.ToInt64(val))      // 应用编号
		requestTotal.AppId = proto.Int64(cast.ToInt64(val)) // 应用编号
	}
	if val, ok := newCtx.GetQuery("channelId"); ok {
		request.ChannelId = proto.Int64(cast.ToInt64(val))      // 渠道编号
		requestTotal.ChannelId = proto.Int64(cast.ToInt64(val)) // 渠道编号
	}
	if val, ok := newCtx.GetQuery("channelCode"); ok {
		request.ChannelCode = proto.String(val)      // 渠道编码
		requestTotal.ChannelCode = proto.String(val) // 渠道编码
	}
	if val, ok := newCtx.GetQuery("merchantOrderId"); ok {
		request.MerchantOrderId = proto.String(val)      // 商户订单编号
		requestTotal.MerchantOrderId = proto.String(val) // 商户订单编号
	}
	if val, ok := newCtx.GetQuery("no"); ok {
		request.No = proto.String(val)      // 支付订单号
		requestTotal.No = proto.String(val) // 支付订单号
	}
	if val, ok := newCtx.GetQuery("channelOrderNo"); ok {
		request.ChannelOrderNo = proto.String(val)      // 渠道订单号
		requestTotal.ChannelOrderNo = proto.String(val) // 渠道订单号
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 支付状态
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 支付状态
	}

	if val, ok := newCtx.GetQuery("beginCreateTime"); ok {
		request.BeginCreateTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 登录时间
		requestTotal.BeginCreateTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 登录时间
	}
	if val, ok := newCtx.GetQuery("finishCreateTime"); ok {
		request.FinishCreateTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 登录时间
		requestTotal.FinishCreateTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 登录时间
	}

	// 执行服务,获取数据量
	resTotal, err := client.PayOrderListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付订单:pay_order:PayOrderList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
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
	res, err := client.PayOrderList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付订单:pay_order:PayOrderList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.PayOrder
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, order.PayOrderDao(item))
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
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
