package system

import (
	"cloud/service/system/dict"
	"cloud/service/system/login"
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
	// 登录日志->system_login_log
	login.RegisterSystemLoginLogServiceServer(server.Server, &login.SrvSystemLoginLogServiceServer{
		Server: server,
	})
	// 字典类型->system_dict_type
	dict.RegisterSystemDictTypeServiceServer(server.Server, &dict.SrvSystemDictTypeServiceServer{
		Server: server,
	})
	// 字典数据表->system_dict
	dict.RegisterSystemDictServiceServer(server.Server, &dict.SrvSystemDictServiceServer{
		Server: server,
	})
}
