package routes

import (
	"cloud/api/system/dept"
	"cloud/api/system/dict"
	"cloud/api/system/file"
	"cloud/api/system/logger"
	"cloud/api/system/mail"
	"cloud/api/system/menu"
	"cloud/api/system/monitor"
	"cloud/api/system/notice"
	"cloud/api/system/notify"
	"cloud/api/system/post"
	"cloud/api/system/role"
	"cloud/api/system/tenant"
	"cloud/api/system/user"
	"cloud/internal/middleware"

	"github.com/abulo/ratel/v3/server/xhertz"
)

func SystemInitRoute(handle *xhertz.Server) {
	auth := handle.Group("/admin").Use(middleware.AuthMiddleware())
	{

		// system_user->系统用户->创建
		auth.POST("/system/user", user.SystemUserCreate)
		// system_user->系统用户->更新
		auth.PUT("/system/user/:id/update", user.SystemUserUpdate)
		// system_user->系统用户->删除
		auth.DELETE("/system/user/:id/delete", user.SystemUserDelete)
		// system_user->系统用户->单条数据信息查看
		auth.GET("/system/user/:id/item", user.SystemUser)
		// system_user->系统用户->恢复
		auth.PUT("/system/user/:id/recover", user.SystemUserRecover)
		// system_user->系统用户->列表
		auth.GET("/system/user", user.SystemUserList)
		// system_user->系统用户->精简列表
		auth.GET("/system/user/simple", user.SystemUserListSimple)
		// system_user->用户信息表->用户的菜单
		auth.GET("/system/user/menu", user.SystemUserMenu)
		// system_user->用户信息表->菜单权限
		auth.GET("/system/user/btn", user.SystemUserBtn)
		// system_user->用户信息表->密码修改
		auth.PUT("/system/user/:id/password", user.SystemUserPassword)
		// system_menu->系统菜单->列表
		auth.GET("/system/menu", menu.SystemMenuList)
		// system_menu->系统菜单->创建
		auth.POST("/system/menu", menu.SystemMenuCreate)
		// system_menu->系统菜单->更新
		auth.PUT("/system/menu/:id/update", menu.SystemMenuUpdate)
		// system_menu->系统菜单->删除
		auth.DELETE("/system/menu/:id/delete", menu.SystemMenuDelete)
		// system_menu->系统菜单->单条数据信息查看
		auth.GET("/system/menu/:id/item", menu.SystemMenu)
		// system_menu->系统菜单->恢复
		auth.PUT("/system/menu/:id/recover", menu.SystemMenuRecover)
		// system_menu->系统菜单->精简列表
		auth.GET("/system/menu/simple", menu.SystemMenuListSimple)

		// system_dict_type->字典类型->创建
		auth.POST("/system/dict", dict.SystemDictTypeCreate)
		// system_dict_type->字典类型->更新
		auth.PUT("/system/dict/:id/update", dict.SystemDictTypeUpdate)
		// system_dict_type->字典类型->删除
		auth.DELETE("/system/dict/:id/delete", dict.SystemDictTypeDelete)
		// system_dict_type->字典类型->单条数据信息查看
		auth.GET("/system/dict/:id/item", dict.SystemDictType)
		// system_dict_type->字典类型->所有数据
		auth.GET("/system/dict/type", dict.SystemDictTypeAll)
		// system_dict_type->字典类型->列表
		auth.GET("/system/dict", dict.SystemDictTypeList)
		// system_dict->字典数据表->创建
		auth.POST("/system/dict/data", dict.SystemDictCreate)
		// system_dict->字典数据表->更新
		auth.PUT("/system/dict/data/:id/update", dict.SystemDictUpdate)
		// system_dict->字典数据表->删除
		auth.DELETE("/system/dict/data/:id/delete", dict.SystemDictDelete)
		// system_dict->字典数据表->单条数据信息查看
		auth.GET("/system/dict/data/:id/item", dict.SystemDict)
		// system_dict->字典数据表->列表
		auth.GET("/system/dict/data", dict.SystemDictList)
		// system_dict->字典数据表->所有数据
		auth.GET("/system/dict/all", dict.SystemDictAll)

		// system_login_log->登录日志->删除
		auth.DELETE("/system/logger/login/:id/delete", logger.SystemLoginLogDelete)
		// system_login_log->登录日志->单条数据信息查看
		auth.GET("/system/logger/login/:id/item", logger.SystemLoginLog)
		// system_login_log->登录日志->恢复
		auth.PUT("/system/logger/login/:id/recover", logger.SystemLoginLogRecover)
		// system_login_log->登录日志->列表
		auth.GET("/system/logger/login", logger.SystemLoginLogList)
		// system_login_log->登录日志->清除
		auth.POST("/system/logger/login/drop", logger.SystemLoginLogDrop)

		// system_operate_log->操作日志->删除
		auth.DELETE("/system/logger/operate/:id/delete", logger.SystemOperateLogDelete)
		// system_operate_log->操作日志->单条数据信息查看
		auth.GET("/system/logger/operate/:id/item", logger.SystemOperateLog)
		// system_operate_log->操作日志->恢复
		auth.PUT("/system/logger/operate/:id/recover", logger.SystemOperateLogRecover)
		// system_operate_log->操作日志->列表
		auth.GET("/system/logger/operate", logger.SystemOperateLogList)
		// system_operate_log->操作日志->清除
		auth.POST("/system/logger/operate/drop", logger.SystemOperateLogDrop)

		// system_entry_log -> 访问日志 -> 删除
		auth.DELETE("/system/logger/entry/:id/delete", logger.SystemEntryLogDelete)
		// system_entry_log -> 访问日志 -> 清空
		auth.POST("/system/logger/entry/drop", logger.SystemEntryLogDrop)
		// system_entry_log -> 访问日志 -> 单条数据信息查看
		auth.GET("/system/logger/entry/:id/item", logger.SystemEntryLog)
		// system_entry_log -> 访问日志 -> 列表
		auth.GET("/system/logger/entry", logger.SystemEntryLogList)

		// system_tenant->租户->创建
		auth.POST("/system/tenant", tenant.SystemTenantCreate)
		// system_tenant->租户->更新
		auth.PUT("/system/tenant/:id/update", tenant.SystemTenantUpdate)
		// system_tenant->租户->删除
		auth.DELETE("/system/tenant/:id/delete", tenant.SystemTenantDelete)
		// system_tenant->租户->单条数据信息查看
		auth.GET("/system/tenant/:id/item", tenant.SystemTenant)
		// system_tenant->租户->恢复
		auth.PUT("/system/tenant/:id/recover", tenant.SystemTenantRecover)
		// system_tenant->租户->列表
		auth.GET("/system/tenant", tenant.SystemTenantList)
		// system_tenant->租户->精简列表
		auth.GET("/system/tenant/simple", tenant.SystemTenantListSimple)
		// system_tenant->租户->租户用户
		auth.GET("/system/tenant/:id/user", tenant.SystemUserList)
		// system_tenant->租户->租户菜单
		auth.GET("/system/tenant/menu", tenant.SystemMenuList)
		// system_tenant->租户->登录
		auth.GET("/system/tenant/:id/login", tenant.SystemTenantLogin)

		// system_tenant_package->租户套餐包->创建
		auth.POST("/system/tenant/package", tenant.SystemTenantPackageCreate)
		// system_tenant_package->租户套餐包->更新
		auth.PUT("/system/tenant/package/:id/update", tenant.SystemTenantPackageUpdate)
		// system_tenant_package->租户套餐包->删除
		auth.DELETE("/system/tenant/package/:id/delete", tenant.SystemTenantPackageDelete)
		// system_tenant_package->租户套餐包->单条数据信息查看
		auth.GET("/system/tenant/package/:id/item", tenant.SystemTenantPackage)
		// system_tenant_package->租户套餐包->恢复
		auth.PUT("/system/tenant/package/:id/recover", tenant.SystemTenantPackageRecover)
		// system_tenant_package->租户套餐包->列表
		auth.GET("/system/tenant/package", tenant.SystemTenantPackageList)
		// system_tenant_package->租户套餐包->精简列表
		auth.GET("/system/tenant/package/simple", tenant.SystemTenantPackageListSimple)
		//

		// system_dept->部门->创建
		auth.POST("/system/dept", dept.SystemDeptCreate)
		// system_dept->部门->更新
		auth.PUT("/system/dept/:id/update", dept.SystemDeptUpdate)
		// system_dept->部门->删除
		auth.DELETE("/system/dept/:id/delete", dept.SystemDeptDelete)
		// system_dept->部门->单条数据信息查看
		auth.GET("/system/dept/:id/item", dept.SystemDept)
		// system_dept->部门->恢复
		auth.PUT("/system/dept/:id/recover", dept.SystemDeptRecover)
		// system_dept->部门->列表
		auth.GET("/system/dept", dept.SystemDeptList)
		// system_dept->部门->精简列表
		auth.GET("/system/dept/simple", dept.SystemDeptListSimple)
		// system_dept->部门->部门列表
		auth.GET("/system/dept/label", dept.SystemDeptListLabel)

		// system_post->职位->创建
		auth.POST("/system/post", post.SystemPostCreate)
		// system_post->职位->更新
		auth.PUT("/system/post/:id/update", post.SystemPostUpdate)
		// system_post->职位->删除
		auth.DELETE("/system/post/:id/delete", post.SystemPostDelete)
		// system_post->职位->单条数据信息查看
		auth.GET("/system/post/:id/item", post.SystemPost)
		// system_post->职位->恢复
		auth.PUT("/system/post/:id/recover", post.SystemPostRecover)
		// system_post->职位->列表
		auth.GET("/system/post", post.SystemPostList)
		// system_post->职位->精简列表
		auth.GET("/system/post/simple", post.SystemPostListSimple)

		// system_role->系统角色->创建
		auth.POST("/system/role", role.SystemRoleCreate)
		// system_role->系统角色->更新
		auth.PUT("/system/role/:id/update", role.SystemRoleUpdate)
		// system_role->系统角色->删除
		auth.DELETE("/system/role/:id/delete", role.SystemRoleDelete)
		// system_role->系统角色->单条数据信息查看
		auth.GET("/system/role/:id/item", role.SystemRole)
		// system_role->系统角色->恢复
		auth.PUT("/system/role/:id/recover", role.SystemRoleRecover)
		// system_role->系统角色->列表
		auth.GET("/system/role", role.SystemRoleList)
		// system_role->系统角色->精简列表
		auth.GET("/system/role/simple", role.SystemRoleListSimple)

		// system_role->系统角色->角色菜单
		auth.GET("/system/role/:id/menu", role.SystemRoleMenuList)
		// system_role->系统角色->角色菜单
		auth.POST("/system/role/:id/menu", role.SystemRoleMenuCreate)

		// system_notice->通知公告表->创建
		auth.POST("/system/notice", notice.SystemNoticeCreate)
		// system_notice->通知公告表->更新
		auth.PUT("/system/notice/:id/update", notice.SystemNoticeUpdate)
		// system_notice->通知公告表->删除
		auth.DELETE("/system/notice/:id/delete", notice.SystemNoticeDelete)
		// system_notice->通知公告表->单条数据信息查看
		auth.GET("/system/notice/:id/item", notice.SystemNotice)
		// system_notice->通知公告表->恢复
		auth.PUT("/system/notice/:id/recover", notice.SystemNoticeRecover)
		// system_notice->通知公告表->列表
		auth.GET("/system/notice", notice.SystemNoticeList)

		// system_notify_message->站内信消息表->创建
		// auth.POST("/system/notify/message", notify.SystemNotifyMessageCreate)
		// system_notify_message->站内信消息表->更新
		// auth.PUT("/system/notify/message/:id/update", notify.SystemNotifyMessageUpdate)
		// system_notify_message->站内信消息表->删除
		auth.DELETE("/system/notify/message/:id/delete", notify.SystemNotifyMessageDelete)
		// system_notify_message->站内信消息表->单条数据信息查看
		auth.GET("/system/notify/message/:id/item", notify.SystemNotifyMessage)
		// system_notify_message->站内信消息表->恢复
		auth.PUT("/system/notify/message/:id/recover", notify.SystemNotifyMessageRecover)
		// system_notify_message->站内信消息表->列表
		auth.GET("/system/notify/message", notify.SystemNotifyMessageList)

		// system_notify_template->站内信模板表->创建
		auth.POST("/system/notify/template", notify.SystemNotifyTemplateCreate)
		// system_notify_template->站内信模板表->更新
		auth.PUT("/system/notify/template/:id/update", notify.SystemNotifyTemplateUpdate)
		// system_notify_template->站内信模板表->删除
		auth.DELETE("/system/notify/template/:id/delete", notify.SystemNotifyTemplateDelete)
		// system_notify_template->站内信模板表->单条数据信息查看
		auth.GET("/system/notify/template/:id/item", notify.SystemNotifyTemplate)
		// system_notify_template->站内信模板表->恢复
		auth.PUT("/system/notify/template/:id/recover", notify.SystemNotifyTemplateRecover)
		// system_notify_template->站内信模板表->列表
		auth.GET("/system/notify/template", notify.SystemNotifyTemplateList)

		// system_mail_account->邮箱账号表->创建
		auth.POST("/system/mail/account", mail.SystemMailAccountCreate)
		// system_mail_account->邮箱账号表->更新
		auth.PUT("/system/mail/account/:id/update", mail.SystemMailAccountUpdate)
		// system_mail_account->邮箱账号表->删除
		auth.DELETE("/system/mail/account/:id/delete", mail.SystemMailAccountDelete)
		// system_mail_account->邮箱账号表->单条数据信息查看
		auth.GET("/system/mail/account/:id/item", mail.SystemMailAccount)
		// system_mail_account->邮箱账号表->恢复
		auth.PUT("/system/mail/account/:id/recover", mail.SystemMailAccountRecover)
		// system_mail_account->邮箱账号表->列表
		auth.GET("/system/mail/account", mail.SystemMailAccountList)
		// system_mail_account->邮箱账号表->精简列表
		auth.GET("/system/mail/account/simple", mail.SystemMailAccountListSimple)

		// system_mail_template->邮件模版表->创建
		auth.POST("/system/mail/template", mail.SystemMailTemplateCreate)
		// system_mail_template->邮件模版表->更新
		auth.PUT("/system/mail/template/:id/update", mail.SystemMailTemplateUpdate)
		// system_mail_template->邮件模版表->删除
		auth.DELETE("/system/mail/template/:id/delete", mail.SystemMailTemplateDelete)
		// system_mail_template->邮件模版表->单条数据信息查看
		auth.GET("/system/mail/template/:id/item", mail.SystemMailTemplate)
		// system_mail_template->邮件模版表->恢复
		auth.PUT("/system/mail/template/:id/recover", mail.SystemMailTemplateRecover)
		// system_mail_template->邮件模版表->列表
		auth.GET("/system/mail/template", mail.SystemMailTemplateList)

		// system_mail_log->邮件日志表->创建
		// auth.POST("/system/mail/log", mail.SystemMailLogCreate)
		// system_mail_log->邮件日志表->更新
		// auth.PUT("/system/mail/log/:id/update", mail.SystemMailLogUpdate)
		// system_mail_log->邮件日志表->删除
		auth.DELETE("/system/mail/log/:id/delete", mail.SystemMailLogDelete)
		// system_mail_log->邮件日志表->单条数据信息查看
		auth.GET("/system/mail/log/:id/item", mail.SystemMailLog)
		// system_mail_log->邮件日志表->恢复
		auth.PUT("/system/mail/log/:id/recover", mail.SystemMailLogRecover)
		// system_mail_log->邮件日志表->列表
		auth.GET("/system/mail/log", mail.SystemMailLogList)

		// system_file->文件管理->创建
		auth.POST("/system/file", file.SystemFileCreate)
		// system_file->文件管理->更新
		auth.PUT("/system/file/:id/update", file.SystemFileUpdate)
		// system_file->文件管理->删除
		auth.DELETE("/system/file/:id/delete", file.SystemFileDelete)
		// system_file->文件管理->单条数据信息查看
		auth.GET("/system/file/:id/item", file.SystemFile)
		// system_file->文件管理->列表
		auth.GET("/system/file", file.SystemFileList)

		// 监控数据-动态数据
		auth.GET("/monitor/trend", monitor.Trend)
		// 监控数据-运行信息
		auth.GET("/monitor/run", monitor.RunInfo)

	}
}
