package menu

import (
	"cloud/code"
	"cloud/dao"
	"cloud/module/system/menu"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_menu 系统菜单

// SrvSystemMenuServiceServer 系统菜单
type SrvSystemMenuServiceServer struct {
	UnimplementedSystemMenuServiceServer
	Server *xgrpc.Server
}

func (srv SrvSystemMenuServiceServer) SystemMenuConvert(request *SystemMenuObject) dao.SystemMenu {
	var res dao.SystemMenu

	if request != nil && request.Id != nil {
		res.Id = request.Id // 菜单编号
	}
	if request != nil && request.Name != nil {
		res.Name = request.Name // 菜单名称
	}
	if request != nil && request.Permission != nil {
		res.Permission = null.StringFrom(request.GetPermission()) // 权限标识
	}
	if request != nil && request.Type != nil {
		res.Type = request.Type // 菜单类型(1:目录/2: 菜单/3: 按钮)
	}
	if request != nil && request.Sort != nil {
		res.Sort = request.Sort // 显示顺序
	}
	if request != nil && request.ParentId != nil {
		res.ParentId = request.ParentId // 父菜单ID
	}
	if request != nil && request.Path != nil {
		res.Path = null.StringFrom(request.GetPath()) // 路由地址
	}
	if request != nil && request.Icon != nil {
		res.Icon = null.StringFrom(request.GetIcon()) // 菜单图标
	}
	if request != nil && request.Component != nil {
		res.Component = null.StringFrom(request.GetComponent()) // 组件路径
	}
	if request != nil && request.ComponentName != nil {
		res.ComponentName = null.StringFrom(request.GetComponentName()) // 组件名
	}
	if request != nil && request.Status != nil {
		res.Status = request.Status // 菜单状态(0开启/1关闭)
	}
	if request != nil && request.Hide != nil {
		res.Hide = request.Hide // 是否隐藏(0:否/1是)
	}
	if request != nil && request.Link != nil {
		res.Link = null.StringFrom(request.GetLink()) // 路由外链时填写的访问地址
	}
	if request != nil && request.KeepAlive != nil {
		res.KeepAlive = request.KeepAlive // 是否缓存(0不/ 是)
	}
	if request != nil && request.Affix != nil {
		res.Affix = request.Affix // 是否总是显示(0 不显示/1 显示)
	}
	if request != nil && request.ActivePath != nil {
		res.ActivePath = null.StringFrom(request.GetActivePath()) // 激活链接
	}
	if request != nil && request.FullScreen != nil {
		res.FullScreen = request.FullScreen // 是否全屏
	}
	if request != nil && request.Redirect != nil {
		res.Redirect = null.StringFrom(request.GetRedirect()) // 路由重定向地址
	}
	if request != nil && request.Creator != nil {
		res.Creator = null.StringFrom(request.GetCreator()) // 创建者
	}
	if request != nil && request.CreateTime != nil {
		res.CreateTime = null.DateTimeFrom(util.GrpcTime(request.CreateTime)) // 创建时间
	}
	if request != nil && request.Updater != nil {
		res.Updater = null.StringFrom(request.GetUpdater()) // 更新者
	}
	if request != nil && request.UpdateTime != nil {
		res.UpdateTime = null.DateTimeFrom(util.GrpcTime(request.UpdateTime)) // 更新时间
	}
	if request != nil && request.Deleted != nil {
		res.Deleted = request.Deleted // 是否删除
	}

	return res
}

func (srv SrvSystemMenuServiceServer) SystemMenuResult(item dao.SystemMenu) *SystemMenuObject {
	res := &SystemMenuObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.Permission.IsValid() {
		res.Permission = item.Permission.Ptr()
	}
	if item.Type != nil {
		res.Type = item.Type
	}
	if item.Sort != nil {
		res.Sort = item.Sort
	}
	if item.ParentId != nil {
		res.ParentId = item.ParentId
	}
	if item.Path.IsValid() {
		res.Path = item.Path.Ptr()
	}
	if item.Icon.IsValid() {
		res.Icon = item.Icon.Ptr()
	}
	if item.Component.IsValid() {
		res.Component = item.Component.Ptr()
	}
	if item.ComponentName.IsValid() {
		res.ComponentName = item.ComponentName.Ptr()
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Hide != nil {
		res.Hide = item.Hide
	}
	if item.Link.IsValid() {
		res.Link = item.Link.Ptr()
	}
	if item.KeepAlive != nil {
		res.KeepAlive = item.KeepAlive
	}
	if item.Affix != nil {
		res.Affix = item.Affix
	}
	if item.ActivePath.IsValid() {
		res.ActivePath = item.ActivePath.Ptr()
	}
	if item.FullScreen != nil {
		res.FullScreen = item.FullScreen
	}
	if item.Redirect.IsValid() {
		res.Redirect = item.Redirect.Ptr()
	}
	if item.Creator.IsValid() {
		res.Creator = item.Creator.Ptr()
	}
	if item.CreateTime.IsValid() {
		res.CreateTime = timestamppb.New(*item.CreateTime.Ptr())
	}
	if item.Updater.IsValid() {
		res.Updater = item.Updater.Ptr()
	}
	if item.UpdateTime.IsValid() {
		res.UpdateTime = timestamppb.New(*item.UpdateTime.Ptr())
	}
	if item.Deleted != nil {
		res.Deleted = item.Deleted
	}

	return res
}

// SystemMenuCreate 创建数据
func (srv SrvSystemMenuServiceServer) SystemMenuCreate(ctx context.Context, request *SystemMenuCreateRequest) (*SystemMenuCreateResponse, error) {
	req := srv.SystemMenuConvert(request.GetData())
	data, err := menu.SystemMenuCreate(ctx, req)
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
	req := srv.SystemMenuConvert(request.GetData())
	_, err := menu.SystemMenuUpdate(ctx, systemMenuId, req)
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
		Data: srv.SystemMenuResult(res),
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
		res = append(res, srv.SystemMenuResult(item))
	}
	return &SystemMenuListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
