package notify

import (
	"context"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/system/notify"

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

// system_notify_message 站内信消息表

// SystemNotifyMessageItem 查询单条数据
func SystemNotifyMessageItem(ctx context.Context, newCtx *app.RequestContext) (*notify.SystemNotifyMessageResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:站内信消息表:system_notify_message:SystemNotifyMessage")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := notify.NewSystemNotifyMessageServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &notify.SystemNotifyMessageRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemNotifyMessage(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:站内信消息表:system_notify_message:SystemNotifyMessage")
		return nil, err
	}
	tenantId := cast.ToInt64(newCtx.GetInt64("tenantId"))
	data := notify.SystemNotifyMessageDao(res.GetData())
	if cast.ToInt64(data.TenantId) != tenantId {
		return nil, status.Error(code.ConvertToGrpc(code.NoPermission), code.StatusText(code.NoPermission))
	}
	return res, nil
}

// SystemNotifyMessageCreate 创建数据
func SystemNotifyMessageCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:站内信消息表:system_notify_message:SystemNotifyMessageCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := notify.NewSystemNotifyMessageServiceClient(grpcClient)
	request := &notify.SystemNotifyMessageCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemNotifyMessage
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
	request.Data = notify.SystemNotifyMessageProto(reqInfo)
	// 执行服务
	res, err := client.SystemNotifyMessageCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:站内信消息表:system_notify_message:SystemNotifyMessageCreate")
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

// SystemNotifyMessageUpdate 更新数据
func SystemNotifyMessageUpdate(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := SystemNotifyMessageItem(ctx, newCtx); err != nil {
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
		}).Error("Grpc:站内信消息表:system_notify_message:SystemNotifyMessageUpdate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := notify.NewSystemNotifyMessageServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &notify.SystemNotifyMessageUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SystemNotifyMessage
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
	request.Data = notify.SystemNotifyMessageProto(reqInfo)
	// 执行服务
	res, err := client.SystemNotifyMessageUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:站内信消息表:system_notify_message:SystemNotifyMessageUpdate")
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

// SystemNotifyMessageDelete 删除数据
func SystemNotifyMessageDelete(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := SystemNotifyMessageItem(ctx, newCtx); err != nil {
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
		}).Error("Grpc:站内信消息表:system_notify_message:SystemNotifyMessageDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := notify.NewSystemNotifyMessageServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &notify.SystemNotifyMessageDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemNotifyMessageDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:站内信消息表:system_notify_message:SystemNotifyMessageDelete")
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

// SystemNotifyMessage 查询单条数据
func SystemNotifyMessage(ctx context.Context, newCtx *app.RequestContext) {
	res, err := SystemNotifyMessageItem(ctx, newCtx)
	if err != nil {
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
		"data": notify.SystemNotifyMessageDao(res.GetData()),
	})
}

// SystemNotifyMessageRecover 恢复数据
func SystemNotifyMessageRecover(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := SystemNotifyMessageItem(ctx, newCtx); err != nil {
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
		}).Error("Grpc:站内信消息表:system_notify_message:SystemNotifyMessageRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := notify.NewSystemNotifyMessageServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &notify.SystemNotifyMessageRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemNotifyMessageRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:站内信消息表:system_notify_message:SystemNotifyMessageRecover")
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

// SystemNotifyMessageList 列表数据
func SystemNotifyMessageList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:站内信消息表:system_notify_message:SystemNotifyMessageList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := notify.NewSystemNotifyMessageServiceClient(grpcClient)
	// 构造查询条件
	request := &notify.SystemNotifyMessageListRequest{}
	requestTotal := &notify.SystemNotifyMessageListTotalRequest{}
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
	if val, ok := newCtx.GetQuery("templateType"); ok {
		request.TemplateType = proto.Int32(cast.ToInt32(val))      // 模版类型
		requestTotal.TemplateType = proto.Int32(cast.ToInt32(val)) // 模版类型
	}
	if val, ok := newCtx.GetQuery("readStatus"); ok {
		request.ReadStatus = proto.Int32(cast.ToInt32(val))      // 是否已读
		requestTotal.ReadStatus = proto.Int32(cast.ToInt32(val)) // 是否已读
	}
	if val, ok := newCtx.GetQuery("beginReadTime"); ok {
		request.BeginReadTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 登录时间
		requestTotal.BeginReadTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 登录时间
	}
	if val, ok := newCtx.GetQuery("finishReadTime"); ok {
		request.FinishReadTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 登录时间
		requestTotal.FinishReadTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 登录时间
	}

	// 执行服务,获取数据量
	resTotal, err := client.SystemNotifyMessageListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:站内信消息表:system_notify_message:SystemNotifyMessageList")
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
	res, err := client.SystemNotifyMessageList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:站内信消息表:system_notify_message:SystemNotifyMessageList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemNotifyMessage
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, notify.SystemNotifyMessageDao(item))
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
