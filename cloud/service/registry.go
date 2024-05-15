package service

import (
	"cloud/service/captcha"
	"cloud/service/pay"
	"cloud/service/region"
	"cloud/service/system"

	"github.com/abulo/ratel/v3/server/xgrpc"
)

// Registry 注册服务
func Registry(server *xgrpc.Server) {
	// 验证码服务->captcha
	captcha.RegisterCaptchaServiceServer(server.Server, &captcha.SrvCaptchaServiceServer{
		Server: server,
	})
	// 地区表->region
	region.RegisterRegionServiceServer(server.Server, &region.SrvRegionServiceServer{
		Server: server,
	})
	system.Registry(server)
	pay.Registry(server)
}
