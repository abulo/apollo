package initial

import (
	"github.com/abulo/ratel/v3/stores/proxy"
	"github.com/abulo/ratel/v3/stores/redis"
	"github.com/spf13/cast"
)

// InitRedis load redis && returns an redis instance.
func (initial *Initial) InitRedis() *Initial {
	configs := initial.Config.Get("redis")
	list := configs.(map[string]interface{})
	links := make(map[string]*redis.Client)
	for node, nodeConfig := range list {
		opts := make([]redis.Option, 0)
		res := nodeConfig.(map[string]interface{})
		if ClientType := cast.ToString(res["ClientType"]); ClientType != "" {
			opts = append(opts, redis.WithClientType(ClientType))
		}
		if Hosts := cast.ToStringSlice(res["Hosts"]); len(Hosts) > 0 {
			opts = append(opts, redis.WithHosts(Hosts))
		}
		if Password := cast.ToString(res["Password"]); Password != "" {
			opts = append(opts, redis.WithPassword(Password))
		}
		if Database := cast.ToInt(res["Database"]); Database > 0 {
			opts = append(opts, redis.WithDatabase(Database))
		}
		if PoolSize := cast.ToInt(res["PoolSize"]); PoolSize > 0 {
			opts = append(opts, redis.WithPoolSize(PoolSize))
		}
		if KeyPrefix := cast.ToString(res["KeyPrefix"]); KeyPrefix != "" {
			opts = append(opts, redis.WithKeyPrefix(KeyPrefix))
		}
		opts = append(opts, redis.WithDisableMetric(cast.ToBool(res["DisableMetric"])))
		opts = append(opts, redis.WithDisableTrace(cast.ToBool(res["DisableTrace"])))
		if Addr := cast.ToString(res["Addr"]); Addr != "" {
			opts = append(opts, redis.WithAddr(Addr))
		}
		if Addrs := cast.ToStringMapString(res["Addrs"]); len(Addrs) > 0 {
			opts = append(opts, redis.WithAddrs(Addrs))
		}
		if MasterName := cast.ToString(res["MasterName"]); MasterName != "" {
			opts = append(opts, redis.WithMasterName(MasterName))
		}
		conn, err := redis.NewRedisClient(opts...)
		if err != nil {
			panic(err)
		}
		links["redis."+node] = conn
	}
	proxyConfigs := initial.Config.Get("proxyredis")
	proxyRes := proxyConfigs.([]map[string]interface{})
	for _, val := range proxyRes {
		proxyPool := proxy.NewRedis()
		if node := cast.ToStringSlice(val["Node"]); len(node) > 0 {
			for _, v := range node {
				proxyPool.Store(links[v])
			}
		}
		if Name := cast.ToString(val["Name"]); Name != "" {
			initial.Store.StoreRedis(Name, proxyPool)
		}
	}
	return initial
}
