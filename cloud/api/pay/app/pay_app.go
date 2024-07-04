package app

import (
	"context"
	"encoding/json"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	serviceApp "cloud/service/pay/app"
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

// pay_app 支付应用信息

// PayApp 查询单条数据
func PayAppItem(ctx context.Context, newCtx *app.RequestContext) (*serviceApp.PayAppResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付应用信息:pay_app:PayApp")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := serviceApp.NewPayAppServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &serviceApp.PayAppRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayApp(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付应用信息:pay_app:PayApp")
		return nil, err
	}
	tenantId := cast.ToInt64(newCtx.GetInt64("tenantId"))
	data := serviceApp.PayAppDao(res.GetData())
	if cast.ToInt64(data.TenantId) != tenantId {
		return nil, status.Error(code.ConvertToGrpc(code.NoPermission), code.StatusText(code.NoPermission))
	}
	return res, nil
}

// PayAppCreate 创建数据
func PayAppCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付应用信息:pay_app:PayAppCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := serviceApp.NewPayAppServiceClient(grpcClient)
	request := &serviceApp.PayAppCreateRequest{}
	// 数据绑定
	var reqInfo dao.PayApp
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
	request.Data = serviceApp.PayAppProto(reqInfo)
	// 执行服务
	res, err := client.PayAppCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付应用信息:pay_app:PayAppCreate")
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

// PayAppUpdate 更新数据
func PayAppUpdate(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := PayAppItem(ctx, newCtx); err != nil {
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
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
		}).Error("Grpc:支付应用信息:pay_app:PayAppUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := serviceApp.NewPayAppServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &serviceApp.PayAppUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.PayApp
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
	request.Data = serviceApp.PayAppProto(reqInfo)
	// 执行服务
	res, err := client.PayAppUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付应用信息:pay_app:PayAppUpdate")
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

// PayAppDelete 删除数据
func PayAppDelete(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := PayAppItem(ctx, newCtx); err != nil {
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付应用信息:pay_app:PayAppDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := serviceApp.NewPayAppServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &serviceApp.PayAppDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayAppDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付应用信息:pay_app:PayAppDelete")
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

// PayApp 查询单条数据
func PayApp(ctx context.Context, newCtx *app.RequestContext) {
	res, err := PayAppItem(ctx, newCtx)
	if err != nil {
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
		"data": serviceApp.PayAppDao(res.GetData()),
	})
}

// PayAppRecover 恢复数据
func PayAppRecover(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := PayAppItem(ctx, newCtx); err != nil {
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付应用信息:pay_app:PayAppRecover")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := serviceApp.NewPayAppServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &serviceApp.PayAppRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayAppRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付应用信息:pay_app:PayAppRecover")
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

// PayAppList 列表数据
func PayAppList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:支付应用信息:pay_app:PayAppList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := serviceApp.NewPayAppServiceClient(grpcClient)
	// 构造查询条件
	request := &serviceApp.PayAppListRequest{}
	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                            // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val)) // 开启状态
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val) // 应用名称
	}

	// 执行服务
	res, err := client.PayAppList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:支付应用信息:pay_app:PayAppList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.PayApp
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		clientChannel := channel.NewPayChannelServiceClient(grpcClient)
		for _, item := range rpcList {
			appItem := serviceApp.PayAppDao(item)
			requestChannel := &channel.PayChannelListRequest{}
			requestChannel.TenantId = appItem.TenantId
			requestChannel.AppId = appItem.Id
			channelMap := make([]string, 0)
			if resChannel, err := clientChannel.PayChannelList(ctx, requestChannel); err == nil {
				if res.GetCode() == code.Success {
					rpcChannelList := resChannel.GetData()
					for _, channelItem := range rpcChannelList {
						if channelItem.GetStatus() == 0 {
							channelMap = append(channelMap, channelItem.GetCode())
						}
					}
				}
			}
			if jsChannel, err := json.Marshal(channelMap); err == nil {
				appItem.ChannelList = null.JSONFrom(jsChannel)
			}
			list = append(list, appItem)
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": list,
	})
}
