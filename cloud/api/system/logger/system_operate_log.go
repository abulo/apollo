package logger

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/system/logger"

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

// system_operate_log 操作日志

// SystemOperateLogDao 数据转换
func SystemOperateLogDao(item *logger.SystemOperateLogObject) *dao.SystemOperateLog {
	daoItem := &dao.SystemOperateLog{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 主键
	}
	if item != nil && item.Username != nil {
		daoItem.Username = item.Username // 用户账号
	}
	if item != nil && item.Module != nil {
		daoItem.Module = item.Module // 模块名称
	}
	if item != nil && item.RequestMethod != nil {
		daoItem.RequestMethod = item.RequestMethod // 请求方法名
	}
	if item != nil && item.RequestUrl != nil {
		daoItem.RequestUrl = item.RequestUrl // 请求地址
	}
	if item != nil && item.UserIp != nil {
		daoItem.UserIp = item.UserIp // 用户 ip
	}
	if item != nil && item.UserAgent != nil {
		daoItem.UserAgent = null.StringFrom(item.GetUserAgent()) // UA
	}
	if item != nil && item.GoMethod != nil {
		daoItem.GoMethod = item.GoMethod // 方法名
	}
	if item != nil && item.GoMethodArgs != nil {
		daoItem.GoMethodArgs = null.JSONFrom(item.GetGoMethodArgs()) // 方法的参数
	}
	if item != nil && item.StartTime != nil {
		daoItem.StartTime = null.DateTimeFrom(util.GrpcTime(item.StartTime)) // 操作开始时间
	}
	if item != nil && item.Duration != nil {
		daoItem.Duration = item.Duration // 执行时长
	}
	if item != nil && item.Channel != nil {
		daoItem.Channel = item.Channel // 渠道
	}
	if item != nil && item.Result != nil {
		daoItem.Result = item.Result // 结果(0 成功/1 失败)
	}
	if item != nil && item.Creator != nil {
		daoItem.Creator = null.StringFrom(item.GetCreator()) // 创建人
	}
	if item != nil && item.CreateTime != nil {
		daoItem.CreateTime = null.DateTimeFrom(util.GrpcTime(item.CreateTime)) // 创建时间
	}
	if item != nil && item.Updater != nil {
		daoItem.Updater = null.StringFrom(item.GetUpdater()) // 更新人
	}
	if item != nil && item.UpdateTime != nil {
		daoItem.UpdateTime = null.DateTimeFrom(util.GrpcTime(item.UpdateTime)) // 更新时间
	}

	return daoItem
}

// SystemOperateLogProto 数据绑定
func SystemOperateLogProto(item dao.SystemOperateLog) *logger.SystemOperateLogObject {
	res := &logger.SystemOperateLogObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Username != nil {
		res.Username = item.Username
	}
	if item.Module != nil {
		res.Module = item.Module
	}
	if item.RequestMethod != nil {
		res.RequestMethod = item.RequestMethod
	}
	if item.RequestUrl != nil {
		res.RequestUrl = item.RequestUrl
	}
	if item.UserIp != nil {
		res.UserIp = item.UserIp
	}
	if item.UserAgent.IsValid() {
		res.UserAgent = item.UserAgent.Ptr()
	}
	if item.GoMethod != nil {
		res.GoMethod = item.GoMethod
	}
	if item.GoMethodArgs.IsValid() {
		res.GoMethodArgs = *item.GoMethodArgs.Ptr()
	}
	if item.StartTime.IsValid() {
		res.StartTime = timestamppb.New(*item.StartTime.Ptr())
	}
	if item.Duration != nil {
		res.Duration = item.Duration
	}
	if item.Channel != nil {
		res.Channel = item.Channel
	}
	if item.Result != nil {
		res.Result = item.Result
	}
	if item.Creator.IsValid() {
		res.Creator = item.Creator.Ptr()
	}
	if item.CreateTime.IsValid() {
		res.CreateTime = timestamppb.New(*item.CreateTime.Ptr())
	}
	if item.Updater.IsValid() {
		res.Updater = item.Updater.Ptr()
	}
	if item.UpdateTime.IsValid() {
		res.UpdateTime = timestamppb.New(*item.UpdateTime.Ptr())
	}

	return res
}

// SystemOperateLogDelete 删除数据
func SystemOperateLogDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:操作日志:system_operate_log:SystemOperateLogDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemOperateLogServiceClient(grpcClient)
	request := &logger.SystemOperateLogDeleteRequest{}
	var reqInfo dao.SystemOperateLogDelete
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	if reqInfo.SystemOperateLogIds.IsValid() {
		request.SystemOperateLogIds = *reqInfo.SystemOperateLogIds.Ptr()
	}
	// 执行服务
	res, err := client.SystemOperateLogDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:system_operate_log:SystemOperateLogDelete")
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

// SystemOperateLog 查询单条数据
func SystemOperateLog(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:操作日志:system_operate_log:SystemOperateLog")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemOperateLogServiceClient(grpcClient)
	systemOperateLogId := cast.ToInt64(newCtx.Param("systemOperateLogId"))
	request := &logger.SystemOperateLogRequest{}
	request.SystemOperateLogId = systemOperateLogId
	// 执行服务
	res, err := client.SystemOperateLog(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:system_operate_log:SystemOperateLog")
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
		"data": SystemOperateLogDao(res.GetData()),
	})
}

// SystemOperateLogList 列表数据
func SystemOperateLogList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:操作日志:system_operate_log:SystemOperateLogList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemOperateLogServiceClient(grpcClient)
	// 构造查询条件
	request := &logger.SystemOperateLogListRequest{}
	requestTotal := &logger.SystemOperateLogListTotalRequest{}

	if val, ok := newCtx.GetQuery("username"); ok {
		request.Username = proto.String(val)      // 用户账号
		requestTotal.Username = proto.String(val) // 用户账号
	}
	if val, ok := newCtx.GetQuery("module"); ok {
		request.Module = proto.String(val)      // 模块名称
		requestTotal.Module = proto.String(val) // 模块名称
	}
	if val, ok := newCtx.GetQuery("beginStartTime"); ok {
		request.BeginStartTime = timestamppb.New(cast.ToTime(val))      // 登录时间
		requestTotal.BeginStartTime = timestamppb.New(cast.ToTime(val)) // 登录时间
	}
	if val, ok := newCtx.GetQuery("finishStartTime"); ok {
		request.FinishStartTime = timestamppb.New(cast.ToTime(val))      // 登录时间
		requestTotal.FinishStartTime = timestamppb.New(cast.ToTime(val)) // 登录时间
	}

	if val, ok := newCtx.GetQuery("result"); ok {
		request.Result = proto.Int32(cast.ToInt32(val))      // 结果(0 成功/1 失败)
		requestTotal.Result = proto.Int32(cast.ToInt32(val)) // 结果(0 成功/1 失败)
	}

	// 执行服务,获取数据量
	resTotal, err := client.SystemOperateLogListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:system_operate_log:SystemOperateLogList")
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
	res, err := client.SystemOperateLogList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:system_operate_log:SystemOperateLogList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemOperateLog
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, SystemOperateLogDao(item))
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

// SystemOperateLogProto 数据绑定
func SystemOperateLogDropProto(item dao.SystemOperateLogDrop) *logger.SystemOperateLogDropRequest {
	res := &logger.SystemOperateLogDropRequest{}

	if item.Username != nil {
		res.Username = item.Username
	}
	if item.Module != nil {
		res.Module = item.Module
	}
	if item.Result != nil {
		res.Result = item.Result
	}
	if item.BeginStartTime.IsValid() {
		res.BeginStartTime = timestamppb.New(*item.BeginStartTime.Ptr())
	}
	if item.FinishStartTime.IsValid() {
		res.FinishStartTime = timestamppb.New(*item.FinishStartTime.Ptr())
	}
	return res
}

// SystemOperateLogDrop 列表数据
func SystemOperateLogDrop(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:操作日志:system_operate_log:SystemOperateLogDrop")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemOperateLogServiceClient(grpcClient)
	// 构造查询条件
	var reqInfo dao.SystemOperateLogDrop
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request := SystemOperateLogDropProto(reqInfo)
	// 构造查询条件
	// 执行服务
	res, err := client.SystemOperateLogDrop(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:system_operate_log:SystemOperateLogDrop")
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
