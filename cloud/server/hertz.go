package server

import (
	"context"

	"cloud/initial"
	"cloud/internal/routes"
	"github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xhertz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/spf13/cast"
)

func hertzPanicRecoveryHandler(ctx context.Context, newCtx *app.RequestContext, err interface{}, stack []byte) {
	logger.Logger.Error("startup")
	hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
	hlog.SystemLogger().Infof("Client: %s", newCtx.Request.Header.UserAgent())
	newCtx.AbortWithStatus(consts.StatusInternalServerError)
}

func (eng *Engine) NewHertzServer() error {
	configApi := initial.Core.Config.Get("server.api")
	cfg := configApi.(map[string]interface{})
	//先获取这个服务是否是需要开启
	if disable := cast.ToBool(cfg["Disable"]); disable {
		logger.Logger.Error("server.api is disabled")
		return nil
	}
	client := xhertz.New()
	client.Host = cast.ToString(cfg["Host"])
	client.Port = cast.ToInt(cfg["Port"])
	client.Deployment = cast.ToString(cfg["Deployment"])
	client.DisableMetric = cast.ToBool(cfg["DisableMetric"])
	client.DisableTrace = cast.ToBool(cfg["DisableTrace"])
	client.DisableSlowQuery = cast.ToBool(cfg["DisableSlowQuery"])
	client.ServiceAddress = cast.ToString(cfg["ServiceAddress"])
	client.SlowQueryThresholdInMilli = cast.ToInt64(cfg["SlowQueryThresholdInMilli"])

	res := client.Build()
	res.Use(recovery.Recovery(recovery.WithRecoveryHandler(hertzPanicRecoveryHandler)))
	routes.InitRoute(res)
	return eng.Serve(res)
}
