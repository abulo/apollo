package main

import (
	"io"
	"os"

	"cloud/crond"
	"cloud/initial"
	"cloud/server"

	"github.com/abulo/ratel/v3/core/env"
	"github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/core/logger/mongo"
	"github.com/abulo/ratel/v3/util"
	"github.com/sirupsen/logrus"
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
	global.InitRate()
}

var BuildVersion string // 编译版本
var BuildTime string    // 编译时间
// 程序主入口
func main() {
	env.SetName("ApolloGrpc")
	env.SetAppID("1")
	env.SetAppRegion("sichuan")
	env.SetAppZone("chengdu")
	env.SetAppMode("product")
	env.SetAppHost("golang")
	env.SetBuildTime(BuildTime)
	env.SetBuildVersion(BuildVersion)
	mgClient := initial.Core.Store.LoadMongoDB("mongodb")
	loggerHook := mongo.DefaultWithURL(mgClient)
	defer loggerHook.Flush()
	logger.Logger.AddHook(loggerHook)
	logger.Logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Logger.SetReportCaller(true)
	if initial.Core.Config.Bool("DisableDebug", true) {
		logger.Logger.SetOutput(io.Discard)

	} else {
		logger.Logger.SetOutput(os.Stdout)
	}
	registryHandle := initial.Core.Client.LoadEtcd("etcd").MustBuild()
	eng := server.NewGrpcEngine()
	//注册 etcd
	eng.SetRegistry(registryHandle)
	//计划任务
	stop := crond.CronJob()
	defer stop()
	if err := eng.Run(); err != nil {
		logger.Logger.Panic(err.Error())
	}
}
