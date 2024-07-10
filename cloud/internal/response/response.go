package response

import (
	"cloud/code"
	"cloud/initial"
	"cloud/internal/encrypt"
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/spf13/cast"
)

func JSON(ctx *app.RequestContext, status int, data utils.H) {
	if !initial.Core.Config.Bool("encrypt.Disable", true) {
		if val, ok := data["data"]; ok {
			stringByte, err := json.Marshal(val)
			if err != nil {
				data["code"] = code.ParamInvalid
				data["msg"] = code.StatusText(code.ParamInvalid)
				data["data"] = nil
			} else {
				data["data"] = encrypt.AesCBCEncrypt(cast.ToString(stringByte), initial.Core.Config.String("encrypt.SecretKey"), initial.Core.Config.String("encrypt.Iv"), encrypt.PKCS7)
			}
		}
	}
	ctx.JSON(status, data)
}
