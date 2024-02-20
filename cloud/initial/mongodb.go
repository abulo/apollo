package initial

import (
	"time"

	"github.com/abulo/ratel/v3/stores/mongodb"
	"github.com/abulo/ratel/v3/stores/proxy"
	"github.com/spf13/cast"
)

// InitMongoDB load mongodb && returns an mongodb instance.
func (initial *Initial) InitMongoDB() *Initial {
	configs := initial.Config.Get("mongodb")
	list := configs.(map[string]interface{})
	links := make(map[string]*mongodb.MongoDB)
	for node, nodeConfig := range list {
		opts := make([]mongodb.Option, 0)
		res := nodeConfig.(map[string]interface{})
		uri := cast.ToString(res["URI"])
		if uri == "" {
			panic("mongodb uri is empty")
		}
		if MaxConnIdleTime := cast.ToInt64(res["MaxConnIdleTime"]); MaxConnIdleTime > 0 {
			opts = append(opts, mongodb.WithMaxConnIdleTime(cast.ToDuration(MaxConnIdleTime)*time.Minute))
		}
		if MaxPoolSize := cast.ToInt64(res["MaxPoolSize"]); MaxPoolSize > 0 {
			opts = append(opts, mongodb.WithMaxPoolSize(cast.ToUint64(MaxPoolSize)))
		}
		if MinPoolSize := cast.ToInt64(res["MinPoolSize"]); MinPoolSize > 0 {
			opts = append(opts, mongodb.WithMinPoolSize(cast.ToUint64(MinPoolSize)))
		}
		client, err := mongodb.NewMongoDBClient(uri, opts...)
		if err != nil {
			panic(err)
		}
		client.SetDisableMetric(cast.ToBool(res["DisableMetric"]))
		client.SetDisableTrace(cast.ToBool(res["DisableTrace"]))
		links["mongodb."+node] = client
	}
	proxyConfigs := initial.Config.Get("proxymongodb")
	proxyRes := proxyConfigs.([]map[string]interface{})
	for _, val := range proxyRes {
		proxyPool := proxy.NewMongoDB()
		if node := cast.ToStringSlice(val["Node"]); len(node) > 0 {
			for _, v := range node {
				proxyPool.Store(links[v])
			}
		}
		if Name := cast.ToString(val["Name"]); Name != "" {
			initial.Store.StoreMongoDB(Name, proxyPool)
		}
	}
	return initial
}
