package routes

import (
	"cloud/api/region"
	"cloud/api/system/dept"
	"cloud/api/system/dict"
	"cloud/api/system/logger"
	"cloud/api/system/menu"
	"cloud/api/system/notice"
	"cloud/api/system/notify"
	"cloud/api/system/post"
	"cloud/api/system/role"
	"cloud/api/system/tenant"
	"cloud/api/system/user"
	"cloud/internal/middleware"

	"github.com/abulo/ratel/v3/server/xhertz"
)

func ManagerInitRoute(handle *xhertz.Server) {
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
		// system_user->系统用户->搜索
		auth.GET("/system/user/search", user.SystemUserSearch)
		// system_user->用户信息表->用户的菜单
		auth.GET("/system/user/menu", user.SystemUserMenu)
		// system_user->用户信息表->菜单权限
		auth.GET("/system/user/btn", user.SystemUserBtn)

		// system_user_post->系统用户职位->创建
		auth.POST("/system/user/:id/post", user.SystemUserPostCreate)
		// system_user_post->系统用户职位->列表
		auth.GET("/system/user/:id/post", user.SystemUserPostList)

		// system_user_role->系统用户和系统角色关联表->创建
		auth.POST("/system/user/:id/role", user.SystemUserRoleCreate)
		// system_user_role->系统用户和系统角色关联表->列表
		auth.GET("/system/user/:id/role", user.SystemUserRoleList)

		// system_user_dept->系统用户部门->创建
		auth.POST("/system/user/:id/dept", user.SystemUserDeptCreate)
		// system_user_dept->系统用户部门->列表
		auth.GET("/system/user/:id/dept", user.SystemUserDeptList)

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
		// system_menu->系统菜单->搜索
		auth.GET("/system/menu/search", menu.SystemMenuSearch)

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
		auth.POST("/system/logger/login/delete", logger.SystemLoginLogDelete)
		// system_login_log->登录日志->清空
		auth.POST("/system/logger/login/drop", logger.SystemLoginLogDrop)
		// system_login_log->登录日志->单条数据信息查看
		auth.GET("/system/logger/login/:id/item", logger.SystemLoginLog)
		// system_login_log->登录日志->列表
		auth.GET("/system/logger/login", logger.SystemLoginLogList)

		// system_operate_log->操作日志->删除
		auth.POST("/system/logger/operate/delete", logger.SystemOperateLogDelete)
		// system_operate_log->操作日志->清空
		auth.POST("/system/logger/operate/drop", logger.SystemOperateLogDrop)
		// system_operate_log->操作日志->单条数据信息查看
		auth.GET("/system/logger/operate/:id/item", logger.SystemOperateLog)
		// system_operate_log->操作日志->列表
		auth.GET("/system/logger/operate", logger.SystemOperateLogList)

		// system_entry_log -> 访问日志 -> 删除
		auth.POST("/system/logger/entry/delete", logger.SystemEntryLogDelete)
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
		// system_tenant->租户->搜索
		auth.GET("/system/tenant/search", tenant.SystemTenantSearch)
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
		// system_tenant_package->租户套餐包->搜索
		auth.GET("/system/tenant/package/search", tenant.SystemTenantPackageSearch)
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
		// system_dept->部门->搜索
		auth.GET("/system/dept/search", dept.SystemDeptSearch)

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
		// system_post->职位->搜索
		auth.GET("/system/post/search", post.SystemPostSearch)

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
		// system_role->系统角色->搜索
		auth.GET("/system/role/search", role.SystemRoleSearch)

		// system_role->系统角色->角色菜单
		auth.GET("/system/role/:id/menu", role.SystemRoleMenuList)
		// system_role->系统角色->角色菜单
		auth.POST("/system/role/:id/menu", role.SystemRoleMenuCreate)
		// system_role->系统角色->数据范围
		auth.GET("/system/role/:id/scope", role.SystemRoleDataScope)
		// system_role->系统角色->数据范围
		auth.POST("/system/role/:id/scope", role.SystemRoleDataScopeCreate)

		// region->地区表->创建
		auth.POST("/region", region.RegionCreate)
		// region->地区表->更新
		auth.PUT("/region/:id/update", region.RegionUpdate)
		// region->地区表->删除
		auth.DELETE("/region/:id/delete", region.RegionDelete)
		// region->地区表->单条数据信息查看
		auth.GET("/region/:id/item", region.Region)
		// region->地区表->恢复
		auth.PUT("/region/:id/recover", region.RegionRecover)
		// region->地区表->列表
		auth.GET("/region", region.RegionList)
		// region->地区表->搜索
		auth.GET("/region/search", region.RegionSearch)

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
		auth.POST("/system/notify/message", notify.SystemNotifyMessageCreate)
		// system_notify_message->站内信消息表->更新
		auth.PUT("/system/notify/message/:id/update", notify.SystemNotifyMessageUpdate)
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
	}
}
