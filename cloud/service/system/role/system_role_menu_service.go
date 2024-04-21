package role

import (
	"cloud/code"
	"cloud/module/system/role"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_role_menu 系统角色和系统菜单关联表

// SrvSystemRoleMenuServiceServer 系统角色和系统菜单关联表
type SrvSystemRoleMenuServiceServer struct {
	UnimplementedSystemRoleMenuServiceServer
	Server *xgrpc.Server
}

// SystemRoleMenuCreate 创建数据
func (srv SrvSystemRoleMenuServiceServer) SystemRoleMenuCreate(ctx context.Context, request *SystemRoleMenuCreateRequest) (*SystemRoleMenuCreateResponse, error) {
	req := SystemRoleMenuCustomDao(request.GetData())
	data, err := role.SystemRoleMenuCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:系统角色和系统菜单关联表:system_role_menu:SystemRoleMenuCreate")
		return &SystemRoleMenuCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemRoleMenuCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}
func (srv SrvSystemRoleMenuServiceServer) SystemRoleMenuList(ctx context.Context, request *SystemRoleMenuListRequest) (*SystemRoleMenuListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.RoleId != nil {
		condition["roleId"] = request.GetRoleId()
	}
	if request.MenuId != nil {
		condition["menuId"] = request.GetMenuId()
	}
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}

	// 获取数据集合
	list, err := role.SystemRoleMenuList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:系统角色和系统菜单关联表:system_role_menu:SystemRoleMenuList")
		return &SystemRoleMenuListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemRoleMenuObject
	for _, item := range list {
		res = append(res, SystemRoleMenuProto(item))
	}
	return &SystemRoleMenuListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
