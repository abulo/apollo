package flash

import (
	"context"

	"cloud/initial"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
)

func PutUrl(ctx context.Context, url string) {
	val := cast.ToString(ctx.Value("CookiesID"))
	if util.Empty(val) {
		return
	}
	session := initial.Core.Session(cast.ToString(val))
	session.Put(ctx, util.NewReplacer(initial.Core.Config.String("cache.manager.url")), url)
}

func GetUrl(ctx context.Context) string {
	var res string
	val := cast.ToString(ctx.Value("CookiesID"))
	if util.Empty(val) {
		return res
	}
	session := initial.Core.Session(cast.ToString(val))
	backUrl := session.Get(ctx, util.NewReplacer(initial.Core.Config.String("cache.manager.url")))
	return cast.ToString(backUrl)
}
