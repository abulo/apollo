package channel

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pay/channel"

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
)

// pay_channel 支付渠道
// PayChannelCreate 创建数据
func PayChannelCreate(ctx context.Context, newCtx *app.RequestContext) {
	// 数据绑定
	var reqInfo dao.PayChannel
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付渠道:pay_channel:PayChannelCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	tenantId := newCtx.GetInt64("tenantId")                   // 租户
	payAppId := cast.ToInt64(newCtx.Param("payAppId"))        // 应用 ID
	channelCode := cast.ToString(newCtx.Param("channelCode")) // 渠道编码
	//链接服务
	client := channel.NewPayChannelServiceClient(grpcClient)
	// 先去查询有没有这个渠道编码的信息
	requestChannel := &channel.PayChannelCodeRequest{}
	requestChannel.Code = proto.String(channelCode) // 渠道编码
	requestChannel.AppId = proto.Int64(payAppId)    // 应用 ID
	requestChannel.TenantId = proto.Int64(tenantId) // 租户
	channelInfo, err := client.PayChannelCode(ctx, requestChannel)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestChannel,
			"err": err,
		}).Error("GrpcCall:支付渠道:pay_channel:PayChannelCode")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	channelItem := channel.PayChannelDao(channelInfo.GetData())
	reqInfo.Deleted = proto.Int32(0)
	reqInfo.TenantId = proto.Int64(tenantId) // 租户
	var httpCode int64
	var httpMsg string
	if channelItem.Id == nil {
		reqInfo.Creator = null.StringFrom(newCtx.GetString("userName"))
		reqInfo.CreateTime = null.DateTimeFrom(util.Now())
		request := &channel.PayChannelCreateRequest{}
		request.Data = channel.PayChannelProto(reqInfo)
		// 执行服务
		res, err := client.PayChannelCreate(ctx, request)
		if err != nil {
			globalLogger.Logger.WithFields(logrus.Fields{
				"req": request,
				"err": err,
			}).Error("GrpcCall:支付渠道:pay_channel:PayChannelCreate")
			fromError := status.Convert(err)
			response.JSON(newCtx, consts.StatusOK, utils.H{
				"code": code.ConvertToHttp(fromError.Code()),
				"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
			})
			return
		}
		httpCode, httpMsg = res.GetCode(), res.GetMsg()
	} else {
		reqInfo.Creator = channelItem.Creator
		reqInfo.CreateTime = channelItem.CreateTime
		reqInfo.Updater = null.StringFrom(newCtx.GetString("userName"))
		reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
		request := &channel.PayChannelUpdateRequest{}
		request.Id = *channelItem.Id
		request.Data = channel.PayChannelProto(reqInfo)
		res, err := client.PayChannelUpdate(ctx, request)
		if err != nil {
			globalLogger.Logger.WithFields(logrus.Fields{
				"req": request,
				"err": err,
			}).Error("GrpcCall:支付渠道:pay_channel:PayChannelUpdate")
			fromError := status.Convert(err)
			response.JSON(newCtx, consts.StatusOK, utils.H{
				"code": code.ConvertToHttp(fromError.Code()),
				"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
			})
			return
		}
		httpCode, httpMsg = res.GetCode(), res.GetMsg()
	}

	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": httpCode,
		"msg":  httpMsg,
	})
}

// PayChannel 查询单条数据
func PayChannel(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付渠道:pay_channel:PayChannel")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	tenantId := newCtx.GetInt64("tenantId")                   // 租户
	payAppId := cast.ToInt64(newCtx.Param("payAppId"))        // 应用 ID
	channelCode := cast.ToString(newCtx.Param("channelCode")) // 渠道编码
	//链接服务
	client := channel.NewPayChannelServiceClient(grpcClient)
	// 先去查询有没有这个渠道编码的信息
	request := &channel.PayChannelCodeRequest{}
	request.Code = proto.String(channelCode) // 渠道编码
	request.AppId = proto.Int64(payAppId)    // 应用 ID
	request.TenantId = proto.Int64(tenantId) // 租户
	res, err := client.PayChannelCode(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付渠道:pay_channel:PayChannelCode")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付渠道:pay_channel:PayChannel")
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
		"data": channel.PayChannelDao(res.GetData()),
	})
}
