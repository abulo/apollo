package logger

import (
	"context"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/system/logger"

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
		"data": logger.SystemOperateLogDao(res.GetData()),
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
		request.BeginStartTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 登录时间
		requestTotal.BeginStartTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 登录时间
	}
	if val, ok := newCtx.GetQuery("finishStartTime"); ok {
		request.FinishStartTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 登录时间
		requestTotal.FinishStartTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 登录时间
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
			list = append(list, logger.SystemOperateLogDao(item))
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
