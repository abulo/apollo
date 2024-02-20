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
	// 解析全局列队
	// _ = cron.AddFunc("InitRegion", task.JobLocal, "*/10 * * * * *", InitRegion)
	cron.Start()
	return func() { cron.Stop() }
}
