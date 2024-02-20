package captcha

import (
	"cloud/code"
	"cloud/initial"
	"context"
	"time"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/redis"
	"github.com/abulo/ratel/v3/util"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// SrvCaptchaServiceServer 验证码服务
type SrvCaptchaServiceServer struct {
	UnimplementedCaptchaServiceServer
	Server *xgrpc.Server
}

// CaptchaDriver 驱动
type CaptchaDriver struct {
	Driver base64Captcha.Driver
	Store  *redis.Client
}

// Generate  生成验证码
func (c *CaptchaDriver) Generate(ctx context.Context) (id, b64s string, err error) {
	id, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Captcha:验证码:Generate")
		return
	}
	captchaKey := util.NewReplacer(initial.Core.Config.String("Cache.Captcha.Id"), ":Id", id)
	_, err = c.Store.Set(ctx, captchaKey, util.StrToLower(answer), time.Second*90)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Captcha:验证码:Generate")
		return
	}
	b64s = item.EncodeB64string()
	return
}

// Verify 验证
func (c *CaptchaDriver) Verify(ctx context.Context, id, answer string, clear bool) bool {
	captchaKey := util.NewReplacer(initial.Core.Config.String("Cache.Captcha.Id"), ":Id", id)
	redisAnswer, err := c.Store.Get(ctx, captchaKey)
	if err != nil {
		return false
	}
	match := util.StrToLower(redisAnswer) == util.StrToLower(answer)
	if clear {
		c.Store.Del(ctx, id)
	}
	return match
}

// NewCaptcha 创建验证码构造器
func (srv SrvCaptchaServiceServer) NewCaptcha(store *redis.Client) *CaptchaDriver {
	randCode := util.Rand(1, 4)
	var driver base64Captcha.Driver
	switch randCode {
	case 1:
		driverString := &base64Captcha.DriverString{
			Width:           360,
			Height:          120,
			Length:          5,
			NoiseCount:      0,
			ShowLineOptions: 14,
			Source:          "23456789qwertyuplkjhgfdsazxcvbnm",
			Fonts:           []string{"wqy-microhei.ttc"},
		}
		driver = driverString.ConvertFonts()
	case 2:
		driverMath := &base64Captcha.DriverMath{
			Width:           360,
			Height:          120,
			NoiseCount:      0,
			ShowLineOptions: 14,
			Fonts:           []string{"wqy-microhei.ttc"},
		}
		driver = driverMath.ConvertFonts()
	case 3:
		// driver = &base64Captcha.DriverDigit{
		// 	Height:   96,
		// 	Width:    360,
		// 	Length:   5,
		// 	DotCount: 99,
		// 	MaxSkew:  3,
		// }
		driverDigit := &base64Captcha.DriverString{
			Width:           360,
			Height:          120,
			Length:          5,
			NoiseCount:      0,
			ShowLineOptions: 14,
			Source:          "123456789",
			Fonts:           []string{"wqy-microhei.ttc"},
		}
		driver = driverDigit.ConvertFonts()
	case 4:
		driverChinese := &base64Captcha.DriverChinese{
			Width:           360,
			Height:          120,
			Length:          4,
			NoiseCount:      0,
			ShowLineOptions: 14,
			Source:          base64Captcha.TxtChineseCharaters,
			Fonts:           []string{"wqy-microhei.ttc"},
		}
		driver = driverChinese.ConvertFonts()
	}
	return &CaptchaDriver{
		Driver: driver,
		Store:  store,
	}
}

// CaptchaGenerate 生成验证码
func (srv SrvCaptchaServiceServer) CaptchaGenerate(ctx context.Context, request *CaptchaGenerateRequest) (*CaptchaGenerateResponse, error) {
	//这里需要去生成验证码
	redisHandler := initial.Core.Store.LoadRedis("redis")
	captchaDriver := srv.NewCaptcha(redisHandler)
	id, b64s, err := captchaDriver.Generate(ctx)
	if err != nil {
		return &CaptchaGenerateResponse{}, status.Error(code.ConvertToGrpc(code.RedisError), err.Error())
	}
	return &CaptchaGenerateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: &CaptchaGenerateObject{
			CaptchaId:    id,
			CaptchaImage: b64s,
		},
	}, nil
}

// CaptchaVerify 验证验证码
func (srv SrvCaptchaServiceServer) CaptchaVerify(ctx context.Context, request *CaptchaVerifyRequest) (*CaptchaVerifyResponse, error) {
	//这里需要去生成验证码
	redisHandler := initial.Core.Store.LoadRedis("redis")
	captchaDriver := srv.NewCaptcha(redisHandler)
	captchaId := request.GetCaptchaId()
	captchaCode := request.GetCaptchaCode()
	res := captchaDriver.Verify(ctx, captchaId, captchaCode, true)
	return &CaptchaVerifyResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: &CaptchaVerifyObject{
			Result: res,
		},
	}, nil
}
