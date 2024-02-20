package main

import (
	"cloud/cmd/init/plugin"
	"cloud/initial"

	"github.com/abulo/ratel/v3/util"
)

func init() {
	// 全局设置
	global := initial.New()
	configPath := global.GetEnvironment(global.Path+"/config/env", "configDir")
	if util.Empty(configPath) {
		panic("configPath is empty")
	}
	//配置加载 toml 文件
	dirs := make([]string, 0)
	dirs = append(dirs, global.Path+"/config/"+configPath)
	//加载配置文件
	global.InitConfig(dirs...)
	global.InitMongoDB()
	global.InitRedis()
	global.InitSql()
	global.InitRegistry()
	global.InitTrace()
}

func main() {
	plugin.InitRegion()
}
