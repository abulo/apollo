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

// SystemLoginLogDelete 删除数据
func SystemLoginLogDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:登录日志:system_login_log:SystemLoginLogDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemLoginLogServiceClient(grpcClient)
	request := &logger.SystemLoginLogDeleteRequest{}
	// 数据绑定
	var reqInfo dao.SystemLoginLogDelete
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	if reqInfo.Ids.IsValid() {
		request.Ids = *reqInfo.Ids.Ptr()
	}
	// 执行服务
	res, err := client.SystemLoginLogDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:system_login_log:SystemLoginLogDelete")
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

// SystemLoginLog 查询单条数据
func SystemLoginLog(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:登录日志:system_login_log:SystemLoginLog")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemLoginLogServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &logger.SystemLoginLogRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemLoginLog(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:system_login_log:SystemLoginLog")
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
		"data": logger.SystemLoginLogDao(res.GetData()),
	})
}

// SystemLoginLogList 列表数据
func SystemLoginLogList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:登录日志:system_login_log:SystemLoginLogList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemLoginLogServiceClient(grpcClient)
	// 构造查询条件
	request := &logger.SystemLoginLogListRequest{}
	requestTotal := &logger.SystemLoginLogListTotalRequest{}
	if val, ok := newCtx.GetQuery("username"); ok {
		request.Username = proto.String(val)      // 用户账号
		requestTotal.Username = proto.String(val) // 用户账号
	}
	if val, ok := newCtx.GetQuery("beginLoginTime"); ok {
		request.BeginLoginTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 登录时间
		requestTotal.BeginLoginTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 登录时间
	}
	if val, ok := newCtx.GetQuery("finishLoginTime"); ok {
		request.FinishLoginTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 登录时间
		requestTotal.FinishLoginTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 登录时间
	}
	if val, ok := newCtx.GetQuery("channel"); ok {
		request.Channel = proto.String(val)      // 渠道
		requestTotal.Channel = proto.String(val) // 渠道
	}
	// 执行服务,获取数据量
	resTotal, err := client.SystemLoginLogListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:system_login_log:SystemLoginLogList")
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
	res, err := client.SystemLoginLogList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:system_login_log:SystemLoginLogList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemLoginLog
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, logger.SystemLoginLogDao(item))
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

// SystemLoginLogDropProto 数据绑定
func SystemLoginLogDropProto(item dao.SystemLoginLogDrop) *logger.SystemLoginLogDropRequest {
	res := &logger.SystemLoginLogDropRequest{}
	if item.Username != nil {
		res.Username = item.Username
	}
	if item.Channel != nil {
		res.Channel = item.Channel
	}
	if item.BeginLoginTime.IsValid() {
		res.BeginLoginTime = timestamppb.New(*item.BeginLoginTime.Ptr())
	}
	if item.FinishLoginTime.IsValid() {
		res.FinishLoginTime = timestamppb.New(*item.FinishLoginTime.Ptr())
	}
	return res
}

// SystemLoginLogDrop 清空数据
func SystemLoginLogDrop(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:登录日志:system_login_log:SystemLoginLogDrop")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemLoginLogServiceClient(grpcClient)
	// 构造查询条件
	var reqInfo dao.SystemLoginLogDrop
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request := SystemLoginLogDropProto(reqInfo)
	// 执行服务
	res, err := client.SystemLoginLogDrop(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:system_login_log:SystemLoginLogDrop")
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
