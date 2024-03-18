package user

import (
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/tools"
	"cloud/service/captcha"
	"cloud/service/system/menu"
	"cloud/service/system/user"
	"context"
	"encoding/json"
	"time"

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

// SystemUserLogin 查询单条数据
func SystemUserLogin(ctx context.Context, newCtx *app.RequestContext) {
	//1.验证参数必传
	var req dao.SystemUserLogin
	if err := newCtx.BindAndValidate(&req); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}

	//2.判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:用户信息表:system_user:SystemUserLogin")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}

	//链接验证码服务
	clientCaptcha := captcha.NewCaptchaServiceClient(grpcClient)
	requestCaptcha := &captcha.CaptchaVerifyRequest{}
	requestCaptcha.CaptchaId = req.CaptchaId
	requestCaptcha.CaptchaCode = req.CaptchaCode
	// 执行服务
	resCaptcha, err := clientCaptcha.CaptchaVerify(ctx, requestCaptcha)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestCaptcha,
			"err": err,
		}).Error("GrpcCall:用户信息表:system_user:SystemUserLogin")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	if resCaptcha.GetCode() != code.Success {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": resCaptcha.GetCode(),
			"msg":  resCaptcha.GetMsg(),
		})
		return
	}
	if !resCaptcha.GetData().GetResult() {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.VerifyCodeError,
			"msg":  code.StatusText(code.VerifyCodeError),
		})
		return
	}
	//3.链接用户信息表服务
	clientSystemUser := user.NewSystemUserServiceClient(grpcClient)
	requestSystemUser := &user.SystemUserLoginRequest{}
	requestSystemUser.Username = proto.String(req.Username)
	// 执行服务
	res, err := clientSystemUser.SystemUserLogin(ctx, requestSystemUser)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestSystemUser,
			"err": err,
		}).Error("GrpcCall:用户信息表:system_user:SystemUserLogin")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	//4.生成 token
	userInfo := res.GetData()
	if userInfo.Id == nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.LoginFail,
			"msg":  code.StatusText(code.LoginFail),
		})
		return
	}
	// 比对密码
	if req.Password != userInfo.GetPassword() {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.LoginFail,
			"msg":  code.StatusText(code.LoginFail),
		})
		return
	}
	second := initial.Core.Config.Int64("Cache.SystemUser.PermissionExpire")
	var systemUserToken dao.SystemUserToken
	systemUserToken.SystemUserId = userInfo.GetId()
	systemUserToken.SystemUserName = userInfo.GetUsername()
	systemUserToken.SystemNickName = userInfo.GetNickname()
	systemUserToken.SystemTenantId = userInfo.GetSystemTenantId()
	token, err := tools.GenerateToken(systemUserToken, "WEB", second)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestCaptcha,
			"err": err,
		}).Error("GrpcCall:用户信息表:system_user:SystemUserLogin")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.TokenError,
			"msg":  code.StatusText(code.TokenError),
		})
		return
	}
	//这里需要将登录信息写入操作列队
	nowTime := util.Now()
	var systemLoginLog dao.SystemLoginLog
	systemLoginLog.Username = userInfo.Username
	systemLoginLog.Channel = proto.String("WEB")
	systemLoginLog.UserIp = proto.String(newCtx.ClientIP())
	systemLoginLog.UserAgent = null.StringFrom(cast.ToString(newCtx.Request.Header.UserAgent()))
	systemLoginLog.LoginTime = null.DateTimeFrom(nowTime)
	systemLoginLog.Creator = null.StringFrom(userInfo.GetUsername()) //创建者
	systemLoginLog.CreateTime = null.DateTimeFrom(nowTime)           //创建时间
	systemLoginLog.Updater = null.StringFrom(userInfo.GetUsername()) //更新者
	systemLoginLog.UpdateTime = null.DateTimeFrom(nowTime)           //更新时间
	systemLoginLog.Deleted = proto.Int32(0)                          // 删除状态
	systemLoginLog.SystemTenantId = userInfo.SystemTenantId          // 租户ID
	// 将这些数据需要全部存储在消息列队中,然后后台去执行消息列队
	redisHandler := initial.Core.Store.LoadRedis("redis")
	key := util.NewReplacer(initial.Core.Config.String("Cache.SystemLoginLog.Queue"))
	bytes, _ := json.Marshal(systemLoginLog)
	redisHandler.LPush(ctx, key, cast.ToString(bytes))
	//todo
	//链接服务
	clientMenu := menu.NewSystemMenuServiceClient(grpcClient)
	// 构造查询条件
	requestMenu := &menu.SystemMenuListRequest{}
	requestMenu.Deleted = proto.Int32(0)
	if res, err := clientMenu.SystemMenuList(ctx, requestMenu); err == nil {
		permissionList := make([]string, 0)
		if res.GetCode() == code.Success {
			rpcMenuList := res.GetData()
			for _, item := range rpcMenuList {
				if !util.Empty(item.GetPermission()) {
					permissionList = append(permissionList, item.GetPermission())
				}
			}
		}
		// 保存权限信息
		keyMenu := util.NewReplacer(initial.Core.Config.String("Cache.SystemUser.Permission"), ":UserId", userInfo.GetId())
		permission, _ := json.Marshal(permissionList)
		redisHandler.Set(ctx, keyMenu, cast.ToString(permission), time.Duration(second)*time.Second)
	}

	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"accessToken": token,
			"nickname":    userInfo.GetNickname(),
		},
	})
}
