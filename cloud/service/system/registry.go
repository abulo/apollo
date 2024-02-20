package system

import (
	"cloud/service/system/menu"
	"cloud/service/system/user"

	"github.com/abulo/ratel/v3/server/xgrpc"
)

// Registry 注册服务
func Registry(server *xgrpc.Server) {
	// 用户信息表->system_user
	user.RegisterSystemUserServiceServer(server.Server, &user.SrvSystemUserServiceServer{
		Server: server,
	})
	// 系统菜单->system_menu
	menu.RegisterSystemMenuServiceServer(server.Server, &menu.SrvSystemMenuServiceServer{
		Server: server,
	})
}
