package routes

import (
	"cloud/api/system/dict"
	"cloud/api/system/logger"
	"cloud/api/system/menu"
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
		auth.PUT("/system/user/:systemUserId/update", user.SystemUserUpdate)
		// system_user->系统用户->删除
		auth.DELETE("/system/user/:systemUserId/delete", user.SystemUserDelete)
		// system_user->系统用户->单条数据信息查看
		auth.GET("/system/user/:systemUserId/item", user.SystemUser)
		// system_user->系统用户->列表
		auth.GET("/system/user", user.SystemUserList)
		// system_user->用户信息表->用户的菜单
		auth.GET("/system/user/menu", user.SystemUserMenu)
		// system_user->用户信息表->菜单权限
		auth.GET("/system/user/btn", user.SystemUserBtn)
		// system_menu->系统菜单->列表
		auth.GET("/system/menu", menu.SystemMenuList)
		// system_menu->系统菜单->创建
		auth.POST("/system/menu", menu.SystemMenuCreate)
		// system_menu->系统菜单->更新
		auth.PUT("/system/menu/:systemMenuId/update", menu.SystemMenuUpdate)
		// system_menu->系统菜单->删除
		auth.DELETE("/system/menu/:systemMenuId/delete", menu.SystemMenuDelete)
		// system_menu->系统菜单->单条数据信息查看
		auth.GET("/system/menu/:systemMenuId/item", menu.SystemMenu)
		// system_menu->系统菜单->恢复
		auth.PUT("/system/menu/:systemMenuId/recover", menu.SystemMenuRecover)
		// system_dict_type->字典类型->创建
		auth.POST("/system/dict", dict.SystemDictTypeCreate)
		// system_dict_type->字典类型->更新
		auth.PUT("/system/dict/:systemDictTypeId/update", dict.SystemDictTypeUpdate)
		// system_dict_type->字典类型->删除
		auth.DELETE("/system/dict/:systemDictTypeId/delete", dict.SystemDictTypeDelete)
		// system_dict_type->字典类型->单条数据信息查看
		auth.GET("/system/dict/:systemDictTypeId/item", dict.SystemDictType)
		// system_dict_type->字典类型->所有数据
		auth.GET("/system/dict/type", dict.SystemDictTypeAll)
		// system_dict_type->字典类型->列表
		auth.GET("/system/dict", dict.SystemDictTypeList)
		// system_dict->字典数据表->创建
		auth.POST("/system/dict/data", dict.SystemDictCreate)
		// system_dict->字典数据表->更新
		auth.PUT("/system/dict/data/:systemDictId/update", dict.SystemDictUpdate)
		// system_dict->字典数据表->删除
		auth.DELETE("/system/dict/data/:systemDictId/delete", dict.SystemDictDelete)
		// system_dict->字典数据表->单条数据信息查看
		auth.GET("/system/dict/data/:systemDictId/item", dict.SystemDict)
		// system_dict->字典数据表->列表
		auth.GET("/system/dict/data", dict.SystemDictList)
		// system_dict->字典数据表->所有数据
		auth.GET("/system/dict/all", dict.SystemDictAll)
		// system_role->系统角色->创建
		auth.POST("/system/role", role.SystemRoleCreate)
		// system_role->系统角色->更新
		auth.PUT("/system/role/:systemRoleId/update", role.SystemRoleUpdate)
		// system_role->系统角色->删除
		auth.DELETE("/system/role/:systemRoleId/delete", role.SystemRoleDelete)
		// system_role->系统角色->单条数据信息查看
		auth.GET("/system/role/:systemRoleId/item", role.SystemRole)
		// system_role->系统角色->列表
		auth.GET("/system/role", role.SystemRoleList)
		// system_login_log->登录日志->删除
		auth.POST("/system/logger/login/delete", logger.SystemLoginLogDelete)
		// system_login_log->登录日志->清空
		auth.POST("/system/logger/login/drop", logger.SystemLoginLogDrop)
		// system_login_log->登录日志->单条数据信息查看
		auth.GET("/system/logger/login/:systemLoginLogId/item", logger.SystemLoginLog)
		// system_login_log->登录日志->列表
		auth.GET("/system/logger/login", logger.SystemLoginLogList)
		// system_operate_log->操作日志->删除
		auth.POST("/system/logger/operate/delete", logger.SystemOperateLogDelete)
		// system_operate_log->操作日志->清空
		auth.POST("/system/logger/operate/drop", logger.SystemOperateLogDrop)
		// system_operate_log->操作日志->单条数据信息查看
		auth.GET("/system/logger/operate/:systemOperateLogId/item", logger.SystemOperateLog)
		// system_operate_log->操作日志->列表
		auth.GET("/system/logger/operate", logger.SystemOperateLogList)
		// system_entry_log -> 访问日志 -> 删除
		auth.POST("/system/logger/entry/delete", logger.SystemEntryLogDelete)
		// system_entry_log -> 访问日志 -> 清空
		auth.POST("/system/logger/entry/drop", logger.SystemEntryLogDrop)
		// system_entry_log -> 访问日志 -> 单条数据信息查看
		auth.GET("/system/logger/entry/:systemEntryLogId/item", logger.SystemEntryLog)
		// system_entry_log -> 访问日志 -> 列表
		auth.GET("/system/logger/entry", logger.SystemEntryLogList)
		// system_tenant->租户->创建
		auth.POST("/system/tenant", tenant.SystemTenantCreate)
		// system_tenant->租户->更新
		auth.PUT("/system/tenant/:systemTenantId/update", tenant.SystemTenantUpdate)
		// system_tenant->租户->删除
		auth.DELETE("/system/tenant/:systemTenantId/delete", tenant.SystemTenantDelete)
		// system_tenant->租户->单条数据信息查看
		auth.GET("/system/tenant/:systemTenantId/item", tenant.SystemTenant)
		// system_tenant->租户->恢复
		auth.PUT("/system/tenant/:systemTenantId/recover", tenant.SystemTenantRecover)
		// system_tenant->租户->列表
		auth.GET("/system/tenant", tenant.SystemTenantList)
	}
}
