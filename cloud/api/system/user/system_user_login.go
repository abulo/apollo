package user

import (
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/tools"
	"cloud/service/captcha"
	"cloud/service/system/menu"
	"cloud/service/system/role"
	"cloud/service/system/tenant"
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
	systemUserToken.UserId = userInfo.GetId()
	systemUserToken.UserName = userInfo.GetUsername()
	systemUserToken.NickName = userInfo.GetNickname()
	systemUserToken.TenantId = userInfo.GetTenantId()
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
	systemLoginLog.TenantId = userInfo.TenantId                      // 租户ID
	// 将这些数据需要全部存储在消息列队中,然后后台去执行消息列队
	redisHandler := initial.Core.Store.LoadRedis("redis")
	key := util.NewReplacer(initial.Core.Config.String("Cache.SystemLoginLog.Queue"))
	bytes, _ := json.Marshal(systemLoginLog)
	redisHandler.LPush(ctx, key, cast.ToString(bytes))

	tenantId := userInfo.GetTenantId() // 租户 Id
	systemUserId := userInfo.GetId()   // 用户 Id

	//链接服务
	tenantClient := tenant.NewSystemTenantServiceClient(grpcClient)
	tenantRequest := &tenant.SystemTenantRequest{}
	tenantRequest.Id = tenantId
	// 执行服务
	tenantRes, tenantErr := tenantClient.SystemTenant(ctx, tenantRequest)
	if tenantErr != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": tenantRequest,
			"err": tenantErr,
		}).Error("GrpcCall:租户:system_tenant:SystemTenant")
		fromError := status.Convert(tenantErr)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	// 判断这个有没有值
	if tenantRes.GetCode() != code.Success {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	// 这是租户的信息
	tenantItem := tenant.SystemTenantDao(tenantRes.GetData())
	// 获取套餐服务
	tenantPackageId := tenantItem.TenantPackageId
	//链接服务
	menuClient := menu.NewSystemMenuServiceClient(grpcClient)
	// 构造查询条件
	menuRequest := &menu.SystemMenuListRequest{}
	menuRequest.Deleted = nil
	menuRequest.Status = nil
	var listMenu []*menu.SystemMenuObject
	var menuIds []int64
	var currentMenuIds []int64 // 当前租户所有的菜单 id
	if cast.ToInt64(tenantPackageId) != 0 {
		menuRequest.Deleted = proto.Int32(0)
		menuRequest.Status = proto.Int32(0)
		//链接服务
		tenantPackageClient := tenant.NewSystemTenantPackageServiceClient(grpcClient)
		systemTenantPackageId := cast.ToInt64(tenantPackageId)
		tenantPackageRequest := &tenant.SystemTenantPackageRequest{}
		tenantPackageRequest.Id = systemTenantPackageId
		// 执行服务
		if res, err := tenantPackageClient.SystemTenantPackage(ctx, tenantPackageRequest); err == nil {
			tenantPackageItem := tenant.SystemTenantPackageDao(res.GetData())
			if tenantPackageItem.MenuIds.IsValid() {
				json.Unmarshal(tenantPackageItem.MenuIds.JSON, &menuIds)
			}
		}
	}
	// 执行服务
	menuRes, menuErr := menuClient.SystemMenuList(ctx, menuRequest)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": menuRequest,
			"err": menuErr,
		}).Error("GrpcCall:系统菜单:system_menu:SystemMenuList")
		fromError := status.Convert(menuErr)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}

	if menuRes.GetCode() == code.Success {
		rpcList := menuRes.GetData()
		for _, item := range rpcList {
			currentMenuIds = append(currentMenuIds, *item.Id)
			if cast.ToInt64(tenantPackageId) != 0 {
				if !util.InArray(*item.Id, menuIds) {
					continue
				}
			}
			// listMenu = append(listMenu, menu.SystemMenuDao(item))
			listMenu = append(listMenu, item)
		}
	}
	if systemUserId != *tenantItem.UserId.Ptr() {
		//链接服务
		userRoleClient := user.NewSystemUserRoleServiceClient(grpcClient)
		// 构造查询条件
		userRoleRequest := &user.SystemUserRoleListRequest{}
		userRoleRequest.UserId = proto.Int64(systemUserId) // 系统用户 ID
		userRoleRequest.TenantId = proto.Int64(tenantId)   // 租户
		userRoleRequest.Deleted = proto.Int32(0)           // 删除

		roleMenuClient := role.NewSystemRoleMenuServiceClient(grpcClient)
		if userRoleRes, err := userRoleClient.SystemUserRoleList(ctx, userRoleRequest); err == nil {
			if userRoleRes.GetCode() == code.Success {
				rpcUserRoleList := userRoleRes.GetData()
				for _, item := range rpcUserRoleList {
					systemRoleId := item.GetRoleId()
					//链接服务
					roleMenuRequest := &role.SystemRoleMenuListRequest{}
					roleMenuRequest.TenantId = proto.Int64(tenantId)   // 租户ID
					roleMenuRequest.Deleted = proto.Int32(0)           // 删除状态
					roleMenuRequest.RoleId = proto.Int64(systemRoleId) // 角色ID
					if roleMenuRes, err := roleMenuClient.SystemRoleMenuList(ctx, roleMenuRequest); err == nil {
						if roleMenuRes.GetCode() == code.Success {
							rpcRoleMenuList := roleMenuRes.GetData()
							var newCurrentMenuIds []int64
							for _, item := range rpcRoleMenuList {
								newCurrentMenuIds = append(newCurrentMenuIds, *item.MenuId)
							}
							currentMenuIds = newCurrentMenuIds
						}
					}
				}
			}
		}
	}
	var permissionList []string
	for _, item := range listMenu {
		if util.InArray(*item.Id, currentMenuIds) {
			if !util.Empty(item.GetPermission()) {
				permissionList = append(permissionList, item.GetPermission())
			}
		}
	}
	// 保存权限信息
	keyMenu := util.NewReplacer(initial.Core.Config.String("Cache.SystemUser.Permission"), ":UserId", userInfo.GetId())
	permission, _ := json.Marshal(permissionList)
	redisHandler.Set(ctx, keyMenu, cast.ToString(permission), time.Duration(second)*time.Second)
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"accessToken": token,
			"nickname":    userInfo.GetNickname(),
		},
	})
}
