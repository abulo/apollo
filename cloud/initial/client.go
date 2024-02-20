package initial

import (
	"time"

	"github.com/abulo/ratel/v3/client"
	"github.com/abulo/ratel/v3/client/grpc"
	"github.com/abulo/ratel/v3/client/grpc/balancer"
	"github.com/abulo/ratel/v3/client/grpc/balancer/p2c"
	"github.com/abulo/ratel/v3/registry/etcdv3"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
	"google.golang.org/grpc/keepalive"
)

func (initial *Initial) InitRegistry() *Initial {
	configs := initial.Config.Get("etcd")
	list := configs.(map[string]interface{})
	links := make(map[string]*etcdv3.Config)
	for node, nodeConfig := range list {
		config := etcdv3.New()
		res := nodeConfig.(map[string]interface{})
		if Name := cast.ToString(res["Name"]); Name != "" {
			config.SetName(Name)
			config.SetNode(Name)
		}
		if Endpoints := cast.ToStringSlice(res["Endpoints"]); len(Endpoints) > 0 {
			config.SetEndpoints(Endpoints)
		}
		if CertFile := cast.ToString(res["CertFile"]); CertFile != "" {
			config.SetCertFile(CertFile)
		}
		if KeyFile := cast.ToString(res["KeyFile"]); KeyFile != "" {
			config.SetKeyFile(KeyFile)
		}
		if CaCert := cast.ToString(res["CaCert"]); CaCert != "" {
			config.SetCaCert(CaCert)
		}
		config.SetBasicAuth(cast.ToBool(res["BasicAuth"]))
		if UserName := cast.ToString(res["UserName"]); UserName != "" {
			config.SetUserName(UserName)
		}
		if Password := cast.ToString(res["Password"]); Password != "" {
			config.SetPassword(Password)
		}
		if ConnectTimeout := cast.ToString(res["ConnectTimeout"]); ConnectTimeout != "" {
			config.SetConnectTimeout(util.Duration(ConnectTimeout))
		}
		config.SetSecure(cast.ToBool(res["Secure"]))
		if AutoSyncInterval := cast.ToString(res["AutoSyncInterval"]); AutoSyncInterval != "" {
			config.SetAutoSyncInterval(util.Duration(AutoSyncInterval))
		}
		if Prefix := cast.ToString(res["Prefix"]); Prefix != "" {
			config.SetPrefix(Prefix)
		}
		links["etcd."+node] = config
	}
	proxyConfigs := initial.Config.Get("proxyetcd")
	proxyRes := proxyConfigs.([]map[string]interface{})
	for _, val := range proxyRes {
		proxyPool := client.NewEtcdConfig()
		if node := cast.ToString(val["Node"]); node != "" {
			proxyPool.Store(links[node])
		}
		if Name := cast.ToString(val["Name"]); Name != "" {
			initial.Client.StoreEtcd(Name, proxyPool)
		}
	}
	return initial
}

func (initial *Initial) InitGrpc() *Initial {
	configs := initial.Config.Get("grpc")
	list := configs.(map[string]interface{})
	links := make(map[string]*grpc.Config)
	for node, nodeConfig := range list {
		config := grpc.New()
		res := nodeConfig.(map[string]interface{})
		if Name := cast.ToString(res["Name"]); Name != "" {
			config.SetName(Name)
		}
		if BalancerName := cast.ToString(res["BalancerName"]); BalancerName != "" {
			config.SetBalancerName(BalancerName)
		}
		if Address := cast.ToString(res["Address"]); Address != "" {
			config.SetAddress(Address)
		}
		config.SetBlock(cast.ToBool(res["Block"]))
		if DialTimeout := cast.ToString(res["DialTimeout"]); DialTimeout != "" {
			config.SetDialTimeout(util.Duration(DialTimeout))
		}
		if ReadTimeout := cast.ToString(res["ReadTimeout"]); ReadTimeout != "" {
			config.SetReadTimeout(util.Duration(ReadTimeout))
		}
		config.SetDirect(cast.ToBool(res["Direct"]))
		if SlowThreshold := cast.ToString(res["SlowThreshold"]); SlowThreshold != "" {
			config.SetSlowThreshold(util.Duration(SlowThreshold))
		}
		config.SetDebug(cast.ToBool(res["Debug"]))
		config.SetDisableTraceInterceptor(cast.ToBool(res["DisableTraceInterceptor"]))
		config.SetDisableAidInterceptor(cast.ToBool(res["DisableAidInterceptor"]))
		config.SetDisableTimeoutInterceptor(cast.ToBool(res["DisableTimeoutInterceptor"]))
		config.SetDisableMetricInterceptor(cast.ToBool(res["DisableMetricInterceptor"]))
		config.SetDisableAccessInterceptor(cast.ToBool(res["DisableAccessInterceptor"]))
		if AccessInterceptorLevel := cast.ToString(res["AccessInterceptorLevel"]); AccessInterceptorLevel != "" {
			config.SetAccessInterceptorLevel(AccessInterceptorLevel)
		}
		config.KeepAlive = &keepalive.ClientParameters{
			Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
			Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
			PermitWithoutStream: true,             // send pings even without active streams
		}
		if Etcd := cast.ToString(res["Etcd"]); Etcd != "" {
			config.SetEtcd(initial.Client.LoadEtcd(Etcd))
		}
		config.BalancerName = balancer.NameSmoothWeightRoundRobin
		config.BalancerName = p2c.Name
		links["grpc."+node] = config
	}

	proxyConfigs := initial.Config.Get("proxygrpc")
	proxyRes := proxyConfigs.([]map[string]interface{})
	for _, val := range proxyRes {
		proxyPool := client.NewGrpcConfig()
		if node := cast.ToString(val["Node"]); node != "" {
			proxyPool.Store(links[node])
		}
		if Name := cast.ToString(val["Name"]); Name != "" {
			initial.Client.StoreGrpc(Name, proxyPool)
		}
	}

	return initial
}
