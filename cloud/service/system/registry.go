package system

import (
	"cloud/service/system/dept"
	"cloud/service/system/dict"
	"cloud/service/system/logger"
	"cloud/service/system/mail"
	"cloud/service/system/menu"
	"cloud/service/system/notice"
	"cloud/service/system/notify"
	"cloud/service/system/post"
	"cloud/service/system/role"
	"cloud/service/system/tenant"
	"cloud/service/system/user"

	"github.com/abulo/ratel/v3/server/xgrpc"
)

// Registry 注册服务
func Registry(server *xgrpc.Server) {
	// 用户信息表->system_user
	user.RegisterSystemUserServiceServer(server.Server, &user.SrvSystemUserServiceServer{
		Server: server,
	})
	// 系统用户职位->system_user_post
	user.RegisterSystemUserPostServiceServer(server.Server, &user.SrvSystemUserPostServiceServer{
		Server: server,
	})
	// 系统用户和系统角色关联表->system_user_role
	user.RegisterSystemUserRoleServiceServer(server.Server, &user.SrvSystemUserRoleServiceServer{
		Server: server,
	})
	// 系统用户部门->system_user_dept
	user.RegisterSystemUserDeptServiceServer(server.Server, &user.SrvSystemUserDeptServiceServer{
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
	// 登录日志->system_login_log
	logger.RegisterSystemLoginLogServiceServer(server.Server, &logger.SrvSystemLoginLogServiceServer{
		Server: server,
	})
	// 操作日志->system_operate_log
	logger.RegisterSystemOperateLogServiceServer(server.Server, &logger.SrvSystemOperateLogServiceServer{
		Server: server,
	})
	// 租户->system_tenant
	tenant.RegisterSystemTenantServiceServer(server.Server, &tenant.SrvSystemTenantServiceServer{
		Server: server,
	})
	// 租户套餐->system_tenant_package
	tenant.RegisterSystemTenantPackageServiceServer(server.Server, &tenant.SrvSystemTenantPackageServiceServer{
		Server: server,
	})
	// 部门->system_dept
	dept.RegisterSystemDeptServiceServer(server.Server, &dept.SrvSystemDeptServiceServer{
		Server: server,
	})
	// 职位->system_post
	post.RegisterSystemPostServiceServer(server.Server, &post.SrvSystemPostServiceServer{
		Server: server,
	})
	// 通知公告表->system_notice
	notice.RegisterSystemNoticeServiceServer(server.Server, &notice.SrvSystemNoticeServiceServer{
		Server: server,
	})
	// 站内信消息表->system_notify_message
	notify.RegisterSystemNotifyMessageServiceServer(server.Server, &notify.SrvSystemNotifyMessageServiceServer{
		Server: server,
	})
	// 站内信模板表->system_notify_template
	notify.RegisterSystemNotifyTemplateServiceServer(server.Server, &notify.SrvSystemNotifyTemplateServiceServer{
		Server: server,
	})
	// 邮箱账号表->system_mail_account
	mail.RegisterSystemMailAccountServiceServer(server.Server, &mail.SrvSystemMailAccountServiceServer{
		Server: server,
	})
	// 邮件模版表->system_mail_template
	mail.RegisterSystemMailTemplateServiceServer(server.Server, &mail.SrvSystemMailTemplateServiceServer{
		Server: server,
	})
	// 邮件日志表->system_mail_log
	mail.RegisterSystemMailLogServiceServer(server.Server, &mail.SrvSystemMailLogServiceServer{
		Server: server,
	})

}
