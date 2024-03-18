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

// system_user_role 系统用户和系统角色关联表
// SystemUserRoleCreate 创建数据
func SystemUserRoleCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户和系统角色关联表:system_user_role:SystemUserRoleCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserRoleServiceClient(grpcClient)
	request := &user.SystemUserRoleCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemUserRoleCustom
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	systemUserId := proto.Int64(cast.ToInt64(newCtx.Param("systemUserId")))
	deleted := proto.Int32(0)
	systemTenantId := proto.Int64(newCtx.GetInt64("systemTenantId")) // 租户
	reqInfo.SystemUserId = systemUserId
	reqInfo.Deleted = deleted
	reqInfo.SystemTenantId = systemTenantId
	reqInfo.Updater = null.StringFrom(newCtx.GetString("systemUserName"))
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	request.Data = user.SystemUserRoleCustomProto(reqInfo)
	// 执行服务
	res, err := client.SystemUserRoleCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户和系统角色关联表:system_user_role:SystemUserRoleCreate")
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

// SystemUserRoleList 列表数据
func SystemUserRoleList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户和系统角色关联表:system_user_role:SystemUserRoleList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserRoleServiceClient(grpcClient)
	// 构造查询条件
	request := &user.SystemUserRoleListRequest{}

	systemUserId := proto.Int64(cast.ToInt64(newCtx.Param("systemUserId")))
	deleted := proto.Int32(0)
	systemTenantId := proto.Int64(newCtx.GetInt64("systemTenantId")) // 租户
	request.SystemUserId = systemUserId                              // 系统用户 ID
	request.SystemTenantId = systemTenantId                          // 租户
	request.Deleted = deleted                                        // 删除
	if val, ok := newCtx.GetQuery("systemRoleId"); ok {
		request.SystemRoleId = proto.Int64(cast.ToInt64(val)) // 角色编号
	}

	// 执行服务
	res, err := client.SystemUserRoleList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户和系统角色关联表:system_user_role:SystemUserRoleList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var listRole []int64
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			listRole = append(listRole, *item.SystemRoleId)
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"systemUserId":  systemUserId,
			"systemRoleIds": listRole,
		},
	})
}
