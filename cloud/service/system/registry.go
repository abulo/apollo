package system

import (
	"cloud/service/system/dict"
	"cloud/service/system/logger"
	"cloud/service/system/menu"
	"cloud/service/system/role"
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
	// 字典类型->system_dict_type
	dict.RegisterSystemDictTypeServiceServer(server.Server, &dict.SrvSystemDictTypeServiceServer{
		Server: server,
	})
	// 字典数据表->system_dict
	dict.RegisterSystemDictServiceServer(server.Server, &dict.SrvSystemDictServiceServer{
		Server: server,
	})
	// 系统角色->system_role
	role.RegisterSystemRoleServiceServer(server.Server, &role.SrvSystemRoleServiceServer{
		Server: server,
	})
	// 系统角色和系统菜单关联表->system_role_menu
	role.RegisterSystemRoleMenuServiceServer(server.Server, &role.SrvSystemRoleMenuServiceServer{
		Server: server,
	})
	// 系统用户和系统角色关联表->system_user_role
	user.RegisterSystemUserRoleServiceServer(server.Server, &user.SrvSystemUserRoleServiceServer{
		Server: server,
	})
	// 登录日志->system_login_log
	logger.RegisterSystemLoginLogServiceServer(server.Server, &logger.SrvSystemLoginLogServiceServer{
		Server: server,
	})
	// 操作日志->system_operate_log
	logger.RegisterSystemOperateLogServiceServer(server.Server, &logger.SrvSystemOperateLogServiceServer{
		Server: server,
	})
}
