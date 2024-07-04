package logger

import (
	"context"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/pagination"
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

// SystemOperateLogItem 查询单条数据
func SystemOperateLogItem(ctx context.Context, newCtx *app.RequestContext) (*logger.SystemOperateLogResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:操作日志:system_operate_log:SystemOperateLog")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := logger.NewSystemOperateLogServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &logger.SystemOperateLogRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemOperateLog(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:system_operate_log:SystemOperateLog")
		return nil, err
	}
	tenantId := cast.ToInt64(newCtx.GetInt64("tenantId"))
	data := logger.SystemOperateLogDao(res.GetData())
	if cast.ToInt64(data.TenantId) != tenantId {
		return nil, status.Error(code.ConvertToGrpc(code.NoPermission), code.StatusText(code.NoPermission))
	}
	return res, nil
}

// SystemOperateLogCreate 创建数据
func SystemOperateLogCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:操作日志:system_operate_log:SystemOperateLogCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemOperateLogServiceClient(grpcClient)
	request := &logger.SystemOperateLogCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemOperateLog
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
	request.Data = logger.SystemOperateLogProto(reqInfo)
	// 执行服务
	res, err := client.SystemOperateLogCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:system_operate_log:SystemOperateLogCreate")
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

// SystemOperateLogUpdate 更新数据
func SystemOperateLogUpdate(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := SystemOperateLogItem(ctx, newCtx); err != nil {
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
		}).Error("Grpc:操作日志:system_operate_log:SystemOperateLogUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemOperateLogServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &logger.SystemOperateLogUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SystemOperateLog
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
	request.Data = logger.SystemOperateLogProto(reqInfo)
	// 执行服务
	res, err := client.SystemOperateLogUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:system_operate_log:SystemOperateLogUpdate")
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

// SystemOperateLogDelete 删除数据
func SystemOperateLogDelete(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := SystemOperateLogItem(ctx, newCtx); err != nil {
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
		}).Error("Grpc:操作日志:system_operate_log:SystemOperateLogDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemOperateLogServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &logger.SystemOperateLogDeleteRequest{}
	request.Id = id
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
	res, err := SystemOperateLogItem(ctx, newCtx)
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
		"data": logger.SystemOperateLogDao(res.GetData()),
	})
}

// SystemOperateLogRecover 恢复数据
func SystemOperateLogRecover(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := SystemOperateLogItem(ctx, newCtx); err != nil {
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
		}).Error("Grpc:操作日志:system_operate_log:SystemOperateLogRecover")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemOperateLogServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &logger.SystemOperateLogRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemOperateLogRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:system_operate_log:SystemOperateLogRecover")
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

	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	requestTotal.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))
	request.Deleted = proto.Int32(0)      // 删除状态
	requestTotal.Deleted = proto.Int32(0) // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
			requestTotal.Deleted = nil
		}
	}

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
	paginationRequest := &pagination.PaginationRequest{}
	paginationRequest.PageNum = proto.Int64(cast.ToInt64(newCtx.Query("pageNum")))
	paginationRequest.PageSize = proto.Int64(cast.ToInt64(newCtx.Query("pageSize")))
	request.Pagination = paginationRequest
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
			"pageNum":  paginationRequest.PageNum,
			"pageSize": paginationRequest.PageSize,
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
	if item.Deleted != nil {
		res.Deleted = item.Deleted
	}
	if item.Ids.IsValid() {
		res.Ids = *item.Ids.Ptr()
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
	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户
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
