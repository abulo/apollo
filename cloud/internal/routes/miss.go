package routes

import (
	"cloud/api/captcha"
	"cloud/api/system/user"

	"github.com/abulo/ratel/v3/server/xhertz"
)

func MissInitRoute(handle *xhertz.Server) {
	miss := handle.Group("/admin")
	{
		// 验证码生成
		miss.GET("/captcha", captcha.CaptchaGenerate)
		// 用户登录
		miss.POST("/login", user.SystemUserLogin)

	}
}
