package middleware

import (
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/tools"
	"context"
	"encoding/json"
	"time"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
)

type Response struct {
	Code int32     `json:"code"`
	Msg  string    `json:"msg"`
	Data null.JSON `json:"data,omitempty"`
}

// AuthMiddleware 验证token
func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, newCtx *app.RequestContext) {
		//加入开始时间
		startTime := util.Now()
		// startTime = cast.ToTime(util.Date("Y-m-d H:i:s", startTime))
		authHeader := newCtx.Request.Header.Get("X-Access-Token")
		if util.Empty(authHeader) {
			newCtx.JSON(consts.StatusUnauthorized, utils.H{
				"code": code.TokenEmptyError,
				"msg":  code.StatusText(code.TokenEmptyError),
			})
			newCtx.Abort()
			return
		}
		// 按空格分割
		parts := util.Explode(".", authHeader)
		if len(parts) != 3 {
			newCtx.JSON(consts.StatusUnauthorized, utils.H{
				"code": code.TokenInvalidError,
				"msg":  code.StatusText(code.TokenInvalidError),
			})
			newCtx.Abort()
			return
		}
		rsp, err := tools.ParseToken(authHeader)
		if err != nil {
			newCtx.JSON(consts.StatusUnauthorized, utils.H{
				"code": code.TokenInvalidError,
				"msg":  code.StatusText(code.TokenInvalidError),
			})
			newCtx.Abort()
			return
		}
		newCtx.Set("userId", rsp.UserId)     // 用户ID
		newCtx.Set("nickName", rsp.NickName) // 昵称
		newCtx.Set("userName", rsp.UserName) // 用户名
		newCtx.Set("tenantId", rsp.TenantId) // 租户

		handlerName := newCtx.HandlerName()
		method := util.Explode("/", handlerName)
		methodName := method[len(method)-1]
		// 检查用户有没有该权限
		redisHandler := initial.Core.Store.LoadRedis("redis")
		menuKey := util.NewReplacer(initial.Core.Config.String("Cache.SystemMenu.Permission"))
		if exist, err := redisHandler.SIsMember(ctx, menuKey, methodName); err == nil {
			if !exist {
				newCtx.Next(ctx)
			} else {
				// 判断一下这个用户的的权限
				key := util.NewReplacer(initial.Core.Config.String("Cache.SystemUser.Permission"), ":UserId", rsp.UserId)
				// 获取用户的权限
				if permission, err := redisHandler.Get(ctx, key); err == nil {
					var permissionList []string
					json.Unmarshal([]byte(permission), &permissionList)
					if util.InArray(methodName, permissionList) {
						newCtx.Next(ctx)
					} else {
						newCtx.JSON(consts.StatusForbidden, utils.H{
							"code": code.TokenInvalidError,
							"msg":  code.StatusText(code.TokenInvalidError),
						})
						newCtx.Abort()
						return
					}
				}
			}
		}
		//添加日志收集流程
		var response Response
		json.Unmarshal(newCtx.Response.Body(), &response)
		channelList, _ := rsp.RegisteredClaims.Audience.MarshalJSON()
		var channel []string
		json.Unmarshal(channelList, &channel)
		if len(channel) < 1 {
			channel = append(channel, "unknown")
		}
		var systemOperateLog dao.SystemOperateLog
		systemOperateLog.Username = proto.String(rsp.UserName)                                         //用户名称
		systemOperateLog.Module = nil                                                                  //模块标题
		systemOperateLog.RequestMethod = proto.String(cast.ToString(newCtx.Request.Method()))          //请求方法名
		systemOperateLog.RequestUrl = proto.String(cast.ToString(newCtx.Request.URI().RequestURI()))   //请求地址
		systemOperateLog.UserIp = proto.String(newCtx.ClientIP())                                      //用户IP
		systemOperateLog.UserAgent = null.StringFrom(cast.ToString(newCtx.Request.Header.UserAgent())) //浏览器UA
		systemOperateLog.GoMethod = proto.String(newCtx.HandlerName())                                 //方法名
		systemOperateLog.GoMethodArgs = null.JSONFrom(newCtx.Request.Body())                           //方法参数
		systemOperateLog.StartTime = null.DateTimeFrom(startTime)
		systemOperateLog.Channel = proto.String(channel[0])                                         //渠道
		systemOperateLog.Duration = proto.Int32(cast.ToInt32(time.Since(startTime).Milliseconds())) //执行时长
		if response.Code == 200 {
			systemOperateLog.Result = proto.Int32(0) //结果(0 成功/1 失败)
		} else {
			systemOperateLog.Result = proto.Int32(1) //结果(0 成功/1 失败)
		}
		systemOperateLog.Creator = null.StringFrom(rsp.UserName)   //创建者
		systemOperateLog.CreateTime = null.DateTimeFrom(startTime) //创建时间
		systemOperateLog.Updater = null.StringFrom(rsp.UserName)   //更新者
		systemOperateLog.UpdateTime = null.DateTimeFrom(startTime) //更新时间
		systemOperateLog.Deleted = proto.Int32(0)                  //删除状态
		systemOperateLog.TenantId = proto.Int64(rsp.TenantId)      //租户

		// 将这些数据需要全部存储在消息列队中,然后后台去执行消息列队

		key := util.NewReplacer(initial.Core.Config.String("Cache.SystemOperateLog.Queue"))
		bytes, _ := json.Marshal(systemOperateLog)
		redisHandler.LPush(ctx, key, cast.ToString(bytes))
	}
}
