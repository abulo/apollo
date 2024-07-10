package captcha

import (
	"cloud/code"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/captcha"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// CaptchaGenerate 生成验证码
func CaptchaGenerate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:验证码:CaptchaGenerate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := captcha.NewCaptchaServiceClient(grpcClient)
	request := &captcha.CaptchaGenerateRequest{}
	// 执行服务
	res, err := client.CaptchaGenerate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:验证码:CaptchaGenerate")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	if res.GetCode() != code.Success {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": res.GetCode(),
			"msg":  code.StatusText(res.GetCode()),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": code.Success,
		"msg":  code.StatusText(code.Success),
		"data": utils.H{
			"captchaId":    res.GetData().GetCaptchaId(),
			"captchaImage": res.GetData().GetCaptchaImage(),
		},
	})
}
