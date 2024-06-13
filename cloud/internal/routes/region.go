package routes

import (
	"cloud/api/region"
	"cloud/internal/middleware"

	"github.com/abulo/ratel/v3/server/xhertz"
)

func RegionInitRoute(handle *xhertz.Server) {
	auth := handle.Group("/admin").Use(middleware.AuthMiddleware())
	{
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
		// region->地区表->精简列表
		auth.GET("/region/simple", region.RegionListSimple)
	}
}
