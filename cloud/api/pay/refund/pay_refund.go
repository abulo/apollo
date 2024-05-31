package refund

import (
	"context"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/pay/refund"

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

// pay_refund 退款订单
// PayRefundCreate 创建数据
func PayRefundCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:退款订单:pay_refund:PayRefundCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := refund.NewPayRefundServiceClient(grpcClient)
	request := &refund.PayRefundCreateRequest{}
	// 数据绑定
	var reqInfo dao.PayRefund
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Deleted = proto.Int32(0)
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户
	reqInfo.Creator = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	request.Data = refund.PayRefundProto(reqInfo)
	// 执行服务
	res, err := client.PayRefundCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:退款订单:pay_refund:PayRefundCreate")
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

// PayRefundUpdate 更新数据
func PayRefundUpdate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:退款订单:pay_refund:PayRefundUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := refund.NewPayRefundServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &refund.PayRefundUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.PayRefund
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
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
	request.Data = refund.PayRefundProto(reqInfo)
	// 执行服务
	res, err := client.PayRefundUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:退款订单:pay_refund:PayRefundUpdate")
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

// PayRefundDelete 删除数据
func PayRefundDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:退款订单:pay_refund:PayRefundDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := refund.NewPayRefundServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &refund.PayRefundDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayRefundDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:退款订单:pay_refund:PayRefundDelete")
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

// PayRefund 查询单条数据
func PayRefund(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:退款订单:pay_refund:PayRefund")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := refund.NewPayRefundServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &refund.PayRefundRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayRefund(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:退款订单:pay_refund:PayRefund")
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
		"data": refund.PayRefundDao(res.GetData()),
	})
}

// PayRefundRecover 恢复数据
func PayRefundRecover(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:退款订单:pay_refund:PayRefundRecover")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := refund.NewPayRefundServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &refund.PayRefundRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayRefundRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:退款订单:pay_refund:PayRefundRecover")
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

// PayRefundList 列表数据
func PayRefundList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:退款订单:pay_refund:PayRefundList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := refund.NewPayRefundServiceClient(grpcClient)
	// 构造查询条件
	request := &refund.PayRefundListRequest{}
	requestTotal := &refund.PayRefundListTotalRequest{}

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
		request.MerchantOrderId = proto.String(val)      // 商户订单编号（商户系统生成）
		requestTotal.MerchantOrderId = proto.String(val) // 商户订单编号（商户系统生成）
	}
	if val, ok := newCtx.GetQuery("orderId"); ok {
		request.OrderId = proto.Int64(cast.ToInt64(val))      // 支付订单编号 pay_order 表id
		requestTotal.OrderId = proto.Int64(cast.ToInt64(val)) // 支付订单编号 pay_order 表id
	}
	if val, ok := newCtx.GetQuery("orderNo"); ok {
		request.OrderNo = proto.String(val)      // 支付订单 no
		requestTotal.OrderNo = proto.String(val) // 支付订单 no
	}
	if val, ok := newCtx.GetQuery("channelOrderNo"); ok {
		request.ChannelOrderNo = proto.String(val)      // 渠道订单号，pay_order 中的 channel_order_no 对应
		requestTotal.ChannelOrderNo = proto.String(val) // 渠道订单号，pay_order 中的 channel_order_no 对应
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 退款状态
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 退款状态
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
	resTotal, err := client.PayRefundListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:退款订单:pay_refund:PayRefundList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var total int64
	request.PageNum = proto.Int64(cast.ToInt64(newCtx.Query("pageNum")))
	request.PageSize = proto.Int64(cast.ToInt64(newCtx.Query("pageSize")))
	if resTotal.GetCode() == code.Success {
		total = resTotal.GetData()
	}
	// 执行服务
	res, err := client.PayRefundList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:退款订单:pay_refund:PayRefundList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.PayRefund
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, refund.PayRefundDao(item))
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"total":    total,
			"list":     list,
			"pageNum":  request.PageNum,
			"pageSize": request.PageSize,
		},
	})
}
