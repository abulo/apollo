package routes

import (
	"cloud/internal/middleware"

	"github.com/abulo/ratel/v3/server/xhertz"
)

func InitRoute(handle *xhertz.Server) {
	handle.Use(middleware.Limiter(), middleware.Request())
	MissInitRoute(handle)
	SystemInitRoute(handle)
	RegionInitRoute(handle)
	PayInitRoute(handle)
}
