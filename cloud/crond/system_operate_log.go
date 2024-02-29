package crond

import (
	"cloud/dao"
	"cloud/initial"
	"cloud/module/system/logger"
	"context"
	"encoding/json"

	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
)

func SystemOperateLogQueue() {
	ctx := context.Background()

	redisHandler := initial.Core.Store.LoadRedis("redis")
	key := util.NewReplacer(initial.Core.Config.String("Cache.SystemOperateLog.Queue"))
	systemMenuModuleKey := util.NewReplacer(initial.Core.Config.String("Cache.SystemMenu.Module"))
	//获取列队长度
	currentLen, err := redisHandler.LLen(ctx, key)
	if err != nil {
		return
	}
	length := cast.ToInt64(300)
	if currentLen > length {
		currentLen = length
	}
	for i := 0; i < cast.ToInt(currentLen); i++ {
		data, err := redisHandler.RPop(ctx, key)
		if err != nil {
			continue
		}
		var res dao.SystemOperateLog
		err = json.Unmarshal([]byte(data), &res)
		if err != nil {
			continue
		}
		method := util.Explode("/", *res.GoMethod)
		handlerName := method[len(method)-1]
		// 获取菜单模块名称
		moduleName, err := redisHandler.HGet(ctx, systemMenuModuleKey, handlerName)
		if err != nil {
			continue
		}
		if moduleName == "" {
			continue
		}
		res.Module = proto.String(moduleName)
		logger.SystemOperateLogCreate(ctx, res)
	}
}
