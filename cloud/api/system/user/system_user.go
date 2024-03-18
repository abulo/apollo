package user

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
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
)

// system_user 系统用户
// SystemUserCreate 创建数据
func SystemUserCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:system_user:SystemUserCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	request := &user.SystemUserCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemUser
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Deleted = proto.Int32(0)
	reqInfo.SystemTenantId = proto.Int64(newCtx.GetInt64("systemTenantId")) // 租户
	reqInfo.Creator = null.StringFrom(newCtx.GetString("systemUserName"))
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	if reqInfo.Password != nil {
		reqInfo.Password = proto.String(util.Md5(cast.ToString(reqInfo.Password)))
	}
	request.Data = user.SystemUserProto(reqInfo)
	// 执行服务
	res, err := client.SystemUserCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUserCreate")
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

// SystemUserUpdate 更新数据
func SystemUserUpdate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:system_user:SystemUserUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	systemUserId := cast.ToInt64(newCtx.Param("systemUserId"))
	request := &user.SystemUserUpdateRequest{}
	request.SystemUserId = systemUserId
	// 数据绑定
	var reqInfo dao.SystemUser
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.SystemTenantId = proto.Int64(newCtx.GetInt64("systemTenantId")) // 租户
	reqInfo.Updater = null.StringFrom(newCtx.GetString("systemUserName"))
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	if reqInfo.Password != nil {
		reqInfo.Password = proto.String(util.Md5(cast.ToString(reqInfo.Password)))
	}
	request.Data = user.SystemUserProto(reqInfo)
	// 执行服务
	res, err := client.SystemUserUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUserUpdate")
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

// SystemUserDelete 删除数据
func SystemUserDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:system_user:SystemUserDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	systemUserId := cast.ToInt64(newCtx.Param("systemUserId"))
	request := &user.SystemUserDeleteRequest{}
	request.SystemUserId = systemUserId
	// 执行服务
	res, err := client.SystemUserDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUserDelete")
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

// SystemUser 查询单条数据
func SystemUser(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:system_user:SystemUser")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	systemUserId := cast.ToInt64(newCtx.Param("systemUserId"))
	request := &user.SystemUserRequest{}
	request.SystemUserId = systemUserId
	// 执行服务
	res, err := client.SystemUser(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUser")
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
		"data": userInfo(res.GetData()),
	})
}

func userInfo(item *user.SystemUserObject) *dao.SystemUser {
	res := user.SystemUserDao(item)
	res.Password = nil
	return res
}

// SystemUserRecover 恢复数据
func SystemUserRecover(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:system_user:SystemUserRecover")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	systemUserId := cast.ToInt64(newCtx.Param("systemUserId"))
	request := &user.SystemUserRecoverRequest{}
	request.SystemUserId = systemUserId
	// 执行服务
	res, err := client.SystemUserRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUserRecover")
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

// SystemUserSearch 列表数据
func SystemUserSearch(ctx context.Context, newCtx *app.RequestContext) {
	SystemUserList(ctx, newCtx)
}

// SystemUserList 列表数据
func SystemUserList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:system_user:SystemUserList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	// 构造查询条件
	request := &user.SystemUserListRequest{}
	requestTotal := &user.SystemUserListTotalRequest{}

	if val, ok := newCtx.GetQuery("username"); ok {
		request.Username = proto.String(val)      // 用户名称
		requestTotal.Username = proto.String(val) // 用户名称
	}
	request.SystemTenantId = proto.Int64(newCtx.GetInt64("systemTenantId"))      // 租户ID
	requestTotal.SystemTenantId = proto.Int64(newCtx.GetInt64("systemTenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                                             // 删除状态
	requestTotal.Deleted = proto.Int32(0)                                        // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
			requestTotal.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 用户状态（0正常 1停用）
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 用户状态（0正常 1停用）
	}

	if val, ok := newCtx.GetQuery("systemDeptId"); ok {
		request.SystemDeptId = proto.Int64(cast.ToInt64(val))      // 用户状态（0正常 1停用）
		requestTotal.SystemDeptId = proto.Int64(cast.ToInt64(val)) // 用户状态（0正常 1停用）
	}

	// 执行服务,获取数据量
	resTotal, err := client.SystemUserListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUserList")
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
	res, err := client.SystemUserList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUserList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemUser
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			userInfo := user.SystemUserDao(item)
			userInfo.Password = nil
			list = append(list, userInfo)
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

// SystemUserPassword 更新密码
func SystemUserPassword(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:用户信息表:system_user:SystemUserPassword")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	systemUserId := cast.ToInt64(newCtx.Param("systemUserId"))
	request := &user.SystemUserPasswordRequest{}
	request.SystemUserId = systemUserId
	// 数据绑定
	var reqInfo dao.SystemUserPassword
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	if request.Password != nil {
		request.Password = proto.String(util.Md5(cast.ToString(reqInfo.Password)))
	}
	// 执行服务
	res, err := client.SystemUserPassword(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:用户信息表:system_user:SystemUserPassword")
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
