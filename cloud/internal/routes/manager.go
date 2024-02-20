package routes

import (
	"cloud/api/system/menu"
	"cloud/api/system/user"

	"github.com/abulo/ratel/v3/server/xhertz"
)

func ManagerInitRoute(handle *xhertz.Server) {
	auth := handle.Group("/admin")
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
		// system_menu->系统菜单->列表
		auth.GET("/system/menu", menu.SystemMenuList)
	}
}
