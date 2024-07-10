package initial

import (
	"time"

	"github.com/abulo/ratel/v3/client"
	"github.com/abulo/ratel/v3/config"
	"github.com/abulo/ratel/v3/registry"
	"github.com/abulo/ratel/v3/stores/proxy"
	"github.com/abulo/ratel/v3/util"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

// Initial ...
type Initial struct {
	Path           string         // 应用程序执行路径
	Config         *config.Config // 配置文件
	Store          *proxy.Proxy   // 数据库链接
	LaunchTime     time.Time      // 时间设置
	Client         *client.Proxy
	GrpcClient     *grpc.ClientConn
	RegistryClient registry.Registry
	Local          *time.Location
	Limiter        *rate.Limiter
}

// Core 系统
var Core *Initial

// New Default returns an Initial instance.
func New() *Initial {
	//时区设置
	time.Local = time.FixedZone("CST", 8*3600)
	engine := NewInitial()
	engine.Local = time.Local
	return engine
}

// NewInitial ...
func NewInitial() *Initial {
	Core = &Initial{
		Store:  proxy.NewProxy(),        // 数据库代理
		Client: client.NewClientProxy(), // 客户端代理
	}
	Core.InitPath(util.GetAppRootPath())
	Core.InitLaunchTime(util.Now())
	return Core
}
