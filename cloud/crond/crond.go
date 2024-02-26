package crond

import (
	"cloud/initial"

	"github.com/abulo/ratel/v3/core/task"
	"github.com/abulo/ratel/v3/core/task/driver/redis"
)

func CronJob() func() {
	redisHandler := initial.Core.Store.LoadRedis("redis")
	driverHandler := redis.NewDriver(redisHandler)
	cron := task.NewTask("WorkerService", driverHandler, task.WithLazyPick(true), task.WithSeconds())
	// 刷新菜单模块名称缓存数据
	_ = cron.AddFunc("SystemMenuModule", task.JobLocal, "0 */1 * * * *", SystemMenuModule)
	// 后台操作日志写入
	_ = cron.AddFunc("SystemOperateLogQueue", task.JobLocal, "*/2 * * * * *", SystemOperateLogQueue)
	// 后的登录日志写入
	_ = cron.AddFunc("SystemLoginLogQueue", task.JobLocal, "*/2 * * * * *", SystemLoginLogQueue)
	cron.Start()
	return func() { cron.Stop() }
}
