package routes

import (
	"github.com/abulo/ratel/v3/server/xhertz"
)

func InitRoute(handle *xhertz.Server) {
	MissInitRoute(handle)
	ManagerInitRoute(handle)
}
