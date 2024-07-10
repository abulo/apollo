package middleware

import (
	"cloud/initial"
	"cloud/internal/response"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func Limiter() app.HandlerFunc {
	return func(ctx context.Context, newCtx *app.RequestContext) {
		if !initial.Core.Limiter.Allow() {
			response.JSON(newCtx, 429, utils.H{
				"code": "429",
				"msg":  "too many requests",
			})
			newCtx.Abort()
			return
		}
		newCtx.Next(ctx)
	}
}
