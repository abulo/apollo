package menu

import (
	"cloud/code"
	"cloud/module/system/menu"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_menu 系统菜单

// SrvSystemMenuServiceServer 系统菜单
type SrvSystemMenuServiceServer struct {
	UnimplementedSystemMenuServiceServer
	Server *xgrpc.Server
}

// SystemMenuCreate 创建数据
func (srv SrvSystemMenuServiceServer) SystemMenuCreate(ctx context.Context, request *SystemMenuCreateRequest) (*SystemMenuCreateResponse, error) {
	req := SystemMenuDao(request.GetData())
	data, err := menu.SystemMenuCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:系统菜单:system_menu:SystemMenuCreate")
		return &SystemMenuCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMenuCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemMenuUpdate 更新数据
func (srv SrvSystemMenuServiceServer) SystemMenuUpdate(ctx context.Context, request *SystemMenuUpdateRequest) (*SystemMenuUpdateResponse, error) {
	systemMenuId := request.GetSystemMenuId()
	if systemMenuId < 1 {
		return &SystemMenuUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemMenuDao(request.GetData())
	_, err := menu.SystemMenuUpdate(ctx, systemMenuId, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:系统菜单:system_menu:SystemMenuUpdate")
		return &SystemMenuUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMenuUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemMenuDelete 删除数据
func (srv SrvSystemMenuServiceServer) SystemMenuDelete(ctx context.Context, request *SystemMenuDeleteRequest) (*SystemMenuDeleteResponse, error) {
	systemMenuId := request.GetSystemMenuId()
	if systemMenuId < 1 {
		return &SystemMenuDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := menu.SystemMenuDelete(ctx, systemMenuId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemMenuId,
			"err": err,
		}).Error("Sql:系统菜单:system_menu:SystemMenuDelete")
		return &SystemMenuDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMenuDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemMenu 查询单条数据
func (srv SrvSystemMenuServiceServer) SystemMenu(ctx context.Context, request *SystemMenuRequest) (*SystemMenuResponse, error) {
	systemMenuId := request.GetSystemMenuId()
	if systemMenuId < 1 {
		return &SystemMenuResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := menu.SystemMenu(ctx, systemMenuId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemMenuId,
			"err": err,
		}).Error("Sql:系统菜单:system_menu:SystemMenu")
		return &SystemMenuResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMenuResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemMenuProto(res),
	}, nil
}

// SystemMenuRecover 恢复数据
func (srv SrvSystemMenuServiceServer) SystemMenuRecover(ctx context.Context, request *SystemMenuRecoverRequest) (*SystemMenuRecoverResponse, error) {
	systemMenuId := request.GetSystemMenuId()
	if systemMenuId < 1 {
		return &SystemMenuRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := menu.SystemMenuRecover(ctx, systemMenuId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemMenuId,
			"err": err,
		}).Error("Sql:系统菜单:system_menu:SystemMenuRecover")
		return &SystemMenuRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemMenuRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemMenuServiceServer) SystemMenuList(ctx context.Context, request *SystemMenuListRequest) (*SystemMenuListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Type != nil {
		condition["type"] = request.GetType()
	}

	// 获取数据集合
	list, err := menu.SystemMenuList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:系统菜单:system_menu:SystemMenuList")
		return &SystemMenuListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemMenuObject
	for _, item := range list {
		res = append(res, SystemMenuProto(item))
	}
	return &SystemMenuListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
