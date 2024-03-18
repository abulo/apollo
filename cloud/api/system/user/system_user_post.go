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

// system_user_post 系统用户职位
// SystemUserPostCreate 创建数据
func SystemUserPostCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户职位:system_user_post:SystemUserPostCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserPostServiceClient(grpcClient)
	request := &user.SystemUserPostCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemUserPostCustom
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
	request.Data = user.SystemUserPostCustomProto(reqInfo)
	// 执行服务
	res, err := client.SystemUserPostCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户职位:system_user_post:SystemUserPostCreate")
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

// SystemUserPostList 列表数据
func SystemUserPostList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户职位:system_user_post:SystemUserPostList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserPostServiceClient(grpcClient)
	// 构造查询条件
	request := &user.SystemUserPostListRequest{}

	systemUserId := proto.Int64(cast.ToInt64(newCtx.Param("systemUserId")))
	deleted := proto.Int32(0)
	systemTenantId := proto.Int64(newCtx.GetInt64("systemTenantId")) // 租户
	request.SystemUserId = systemUserId                              // 系统用户 ID
	if val, ok := newCtx.GetQuery("systemPostId"); ok {
		request.SystemPostId = proto.Int64(cast.ToInt64(val)) // 职位 id
	}
	request.SystemTenantId = systemTenantId // 租户
	request.Deleted = deleted               // 删除

	// 执行服务
	res, err := client.SystemUserPostList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户职位:system_user_post:SystemUserPostList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var listPost []int64
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			listPost = append(listPost, *item.SystemPostId)
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"systemUserId":  systemUserId,
			"systemPostIds": listPost,
		},
	})
}
