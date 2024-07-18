package logger

import (
	"context"
	"encoding/json"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/system/logger"
	"cloud/service/system/user"

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

// system_login_log 登录日志
// SystemLoginLogItem 查询单条数据
func SystemLoginLogItem(ctx context.Context, newCtx *app.RequestContext) (*logger.SystemLoginLogResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:登录日志:system_login_log:SystemLoginLog")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
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
		return nil, err
	}
	tenantId := cast.ToInt64(newCtx.GetInt64("tenantId"))
	data := logger.SystemLoginLogDao(res.GetData())
	if cast.ToInt64(data.TenantId) != tenantId {
		return nil, status.Error(code.ConvertToGrpc(code.NoPermission), code.StatusText(code.NoPermission))
	}
	return res, nil
}

// SystemLoginLogCreate 创建数据
func SystemLoginLogCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:登录日志:system_login_log:SystemLoginLogCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemLoginLogServiceClient(grpcClient)
	request := &logger.SystemLoginLogCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemLoginLog
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
	request.Data = logger.SystemLoginLogProto(reqInfo)
	// 执行服务
	res, err := client.SystemLoginLogCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:system_login_log:SystemLoginLogCreate")
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

// SystemLoginLogUpdate 更新数据
func SystemLoginLogUpdate(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := SystemLoginLogItem(ctx, newCtx); err != nil {
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
		}).Error("Grpc:登录日志:system_login_log:SystemLoginLogUpdate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemLoginLogServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &logger.SystemLoginLogUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SystemLoginLog
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
	request.Data = logger.SystemLoginLogProto(reqInfo)
	// 执行服务
	res, err := client.SystemLoginLogUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:system_login_log:SystemLoginLogUpdate")
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

// SystemLoginLogDelete 删除数据
func SystemLoginLogDelete(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := SystemLoginLogItem(ctx, newCtx); err != nil {
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
		}).Error("Grpc:登录日志:system_login_log:SystemLoginLogDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemLoginLogServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &logger.SystemLoginLogDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemLoginLogDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:system_login_log:SystemLoginLogDelete")
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

// SystemLoginLog 查询单条数据
func SystemLoginLog(ctx context.Context, newCtx *app.RequestContext) {
	res, err := SystemLoginLogItem(ctx, newCtx)
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
		"data": logger.SystemLoginLogDao(res.GetData()),
	})
}

// SystemLoginLogRecover 恢复数据
func SystemLoginLogRecover(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := SystemLoginLogItem(ctx, newCtx); err != nil {
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
		}).Error("Grpc:登录日志:system_login_log:SystemLoginLogRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSystemLoginLogServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &logger.SystemLoginLogRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemLoginLogRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:system_login_log:SystemLoginLogRecover")
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

// SystemLoginLogList 列表数据
func SystemLoginLogList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:登录日志:system_login_log:SystemLoginLogList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	clientUser := user.NewSystemUserServiceClient(grpcClient)

	userDataScopeReq := &user.SystemUserRoleDataScopeRequest{}
	userDataScopeReq.UserId = newCtx.GetInt64("userId")
	userDataScopeReq.TenantId = newCtx.GetInt64("tenantId")
	userRes, err := clientUser.SystemUserRoleDataScope(ctx, userDataScopeReq)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": userDataScopeReq,
			"err": err,
		}).Error("GrpcCall:部门:system_dept:SystemDeptList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	userScope := user.SystemUserRoleDataScopeDao(userRes.GetData())
	if len(userScope.DataScopeDept) < 1 {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.DeptError,
			"msg":  code.StatusText(code.DeptError),
		})
		return
	}

	//链接服务
	client := logger.NewSystemLoginLogServiceClient(grpcClient)
	// 构造查询条件
	request := &logger.SystemLoginLogListRequest{}
	requestTotal := &logger.SystemLoginLogListTotalRequest{}

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

	request.UserId = proto.Int64(newCtx.GetInt64("userId"))      // 用户ID
	requestTotal.UserId = proto.Int64(newCtx.GetInt64("userId")) // 用户ID

	request.DataScope = userScope.DataScope
	requestTotal.DataScope = userScope.DataScope
	dataScopeDept, _ := json.Marshal(userScope.DataScopeDept)
	request.DataScopeDept = dataScopeDept
	requestTotal.DataScopeDept = dataScopeDept

	// 执行服务,获取数据量
	resTotal, err := client.SystemLoginLogListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:system_login_log:SystemLoginLogList")
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
	res, err := client.SystemLoginLogList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:system_login_log:SystemLoginLogList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
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
	if item.Deleted != nil {
		res.Deleted = item.Deleted
	}
	if item.Ids.IsValid() {
		res.Ids = *item.Ids.Ptr()
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
		response.JSON(newCtx, consts.StatusOK, utils.H{
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
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request := SystemLoginLogDropProto(reqInfo)
	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户
	// 执行服务
	res, err := client.SystemLoginLogDrop(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:system_login_log:SystemLoginLogDrop")
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
