package crond

import (
	"cloud/initial"
	"cloud/module/system/menu"
	"context"

	"github.com/abulo/ratel/v3/util"
)

// SystemMenuModule 系统菜单模去遍历数据
func SystemMenuModule() {
	ctx := context.Background()
	condition := make(map[string]any)
	// 获取数据集合
	list, err := menu.SystemMenuList(ctx, condition)
	if err != nil {
		return
	}
	redisHandler := initial.Core.Store.LoadRedis("redis")
	key := util.NewReplacer(initial.Core.Config.String("Cache.SystemMenu.Module"))
	for _, v := range list {
		if v.Permission.String != "" {
			itemList, err := menu.SystemMenuListRecursive(ctx, *v.Id)
			if err != nil {
				continue
			}
			data := make([]string, 0)
			for _, item := range itemList {
				data = append(data, *item.Name)
			}
			redisHandler.HSet(ctx, key, v.Permission.String, util.Implode("-", data))
		}
	}
}
