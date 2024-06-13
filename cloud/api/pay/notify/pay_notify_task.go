package notify

import (
	"context"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/pagination"
	"cloud/service/pay/notify"

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

// pay_notify_task 商户支付-任务通知
// PayNotifyTaskCreate 创建数据
func PayNotifyTaskCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:商户支付-任务通知:pay_notify_task:PayNotifyTaskCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := notify.NewPayNotifyTaskServiceClient(grpcClient)
	request := &notify.PayNotifyTaskCreateRequest{}
	// 数据绑定
	var reqInfo dao.PayNotifyTask
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
	request.Data = notify.PayNotifyTaskProto(reqInfo)
	// 执行服务
	res, err := client.PayNotifyTaskCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:商户支付-任务通知:pay_notify_task:PayNotifyTaskCreate")
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

// PayNotifyTaskUpdate 更新数据
func PayNotifyTaskUpdate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:商户支付-任务通知:pay_notify_task:PayNotifyTaskUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := notify.NewPayNotifyTaskServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &notify.PayNotifyTaskUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.PayNotifyTask
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
	request.Data = notify.PayNotifyTaskProto(reqInfo)
	// 执行服务
	res, err := client.PayNotifyTaskUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:商户支付-任务通知:pay_notify_task:PayNotifyTaskUpdate")
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

// PayNotifyTaskDelete 删除数据
func PayNotifyTaskDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:商户支付-任务通知:pay_notify_task:PayNotifyTaskDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := notify.NewPayNotifyTaskServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &notify.PayNotifyTaskDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayNotifyTaskDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:商户支付-任务通知:pay_notify_task:PayNotifyTaskDelete")
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

// PayNotifyTask 查询单条数据
func PayNotifyTask(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:商户支付-任务通知:pay_notify_task:PayNotifyTask")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := notify.NewPayNotifyTaskServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &notify.PayNotifyTaskRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayNotifyTask(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:商户支付-任务通知:pay_notify_task:PayNotifyTask")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	notifyItem := notify.PayNotifyTaskDao(res.GetData())
	clientLog := notify.NewPayNotifyLogServiceClient(grpcClient)
	requestLog := &notify.PayNotifyLogListRequest{}
	requestLog.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户
	requestLog.TaskId = proto.Int64(id)
	logs := make([]*dao.PayNotifyLog, 0)
	if res, err := clientLog.PayNotifyLogList(ctx, requestLog); err == nil {
		if res.GetCode() == code.Success {
			rpcList := res.GetData()
			for _, item := range rpcList {
				logs = append(logs, notify.PayNotifyLogDao(item))
			}
		}
	}
	notifyItem.Logs = logs
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": notifyItem,
	})
}

// PayNotifyTaskRecover 恢复数据
func PayNotifyTaskRecover(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:商户支付-任务通知:pay_notify_task:PayNotifyTaskRecover")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := notify.NewPayNotifyTaskServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &notify.PayNotifyTaskRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayNotifyTaskRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:商户支付-任务通知:pay_notify_task:PayNotifyTaskRecover")
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

// PayNotifyTaskList 列表数据
func PayNotifyTaskList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:商户支付-任务通知:pay_notify_task:PayNotifyTaskList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := notify.NewPayNotifyTaskServiceClient(grpcClient)
	// 构造查询条件
	request := &notify.PayNotifyTaskListRequest{}
	requestTotal := &notify.PayNotifyTaskListTotalRequest{}

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
	if val, ok := newCtx.GetQuery("type"); ok {
		request.Type = proto.Int32(cast.ToInt32(val))      // 通知类型
		requestTotal.Type = proto.Int32(cast.ToInt32(val)) // 通知类型
	}
	if val, ok := newCtx.GetQuery("dataId"); ok {
		request.DataId = proto.Int64(cast.ToInt64(val))      // 数据编号
		requestTotal.DataId = proto.Int64(cast.ToInt64(val)) // 数据编号
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 通知状态
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 通知状态
	}
	if val, ok := newCtx.GetQuery("merchantOrderId"); ok {
		request.MerchantOrderId = proto.String(val)      // 商户订单编号
		requestTotal.MerchantOrderId = proto.String(val) // 商户订单编号
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
	resTotal, err := client.PayNotifyTaskListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:商户支付-任务通知:pay_notify_task:PayNotifyTaskList")
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
	res, err := client.PayNotifyTaskList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:商户支付-任务通知:pay_notify_task:PayNotifyTaskList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.PayNotifyTask
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, notify.PayNotifyTaskDao(item))
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
