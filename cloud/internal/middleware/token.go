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
		newCtx.Set("systemUserId", rsp.SystemUserId) // 用户ID
		newCtx.Set("systemNickName", rsp.SystemNickName)
		newCtx.Set("systemUserName", rsp.SystemUserName)
		newCtx.Next(ctx)

		//添加日志收集流程
		var response Response
		json.Unmarshal(newCtx.Response.Body(), &response)
		channelList, _ := rsp.RegisteredClaims.Audience.MarshalJSON()
		var channel []string
		json.Unmarshal(channelList, &channel)
		if len(channel) < 1 {
			channel = append(channel, "unknown")
		}
		var resSystemOperateLog dao.SystemOperateLog
		resSystemOperateLog.Username = proto.String(rsp.SystemUserName)                                   //用户名称
		resSystemOperateLog.Module = nil                                                                  //模块标题
		resSystemOperateLog.RequestMethod = proto.String(cast.ToString(newCtx.Request.Method()))          //请求方法名
		resSystemOperateLog.RequestUrl = proto.String(cast.ToString(newCtx.Request.URI().RequestURI()))   //请求地址
		resSystemOperateLog.UserIp = proto.String(newCtx.ClientIP())                                      //用户IP
		resSystemOperateLog.UserAgent = null.StringFrom(cast.ToString(newCtx.Request.Header.UserAgent())) //浏览器UA
		resSystemOperateLog.GoMethod = proto.String(newCtx.HandlerName())                                 //方法名
		resSystemOperateLog.GoMethodArgs = null.JSONFrom(newCtx.Request.Body())                           //方法参数
		resSystemOperateLog.StartTime = null.DateTimeFrom(startTime)                                      //开始时间
		resSystemOperateLog.Channel = proto.String(channel[0])                                            //渠道
		resSystemOperateLog.Duration = proto.Int32(cast.ToInt32(time.Since(startTime).Milliseconds()))    //执行时长
		if response.Code == 200 {
			resSystemOperateLog.Result = proto.Int32(0) //结果(0 成功/1 失败)
		} else {
			resSystemOperateLog.Result = proto.Int32(1) //结果(0 成功/1 失败)
		}
		resSystemOperateLog.Creator = null.StringFrom(rsp.SystemUserName) //创建者
		resSystemOperateLog.CreateTime = null.DateTimeFrom(startTime)     //创建时间
		resSystemOperateLog.Updater = null.StringFrom(rsp.SystemUserName) //更新者
		resSystemOperateLog.UpdateTime = null.DateTimeFrom(startTime)     //更新时间
		// 将这些数据需要全部存储在消息列队中,然后后台去执行消息列队
		redisHandler := initial.Core.Store.LoadRedis("redis")
		key := util.NewReplacer(initial.Core.Config.String("Cache.SystemOperateLog.Queue"))
		bytes, _ := json.Marshal(resSystemOperateLog)
		redisHandler.LPush(ctx, key, cast.ToString(bytes))
	}
}
