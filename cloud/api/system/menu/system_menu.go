package menu

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/system/menu"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_menu 系统菜单

// SystemMenuDao 数据转换
func SystemMenuDao(item *menu.SystemMenuObject) *dao.SystemMenu {
	daoItem := &dao.SystemMenu{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 菜单编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 菜单名称
	}
	if item != nil && item.Permission != nil {
		daoItem.Permission = null.StringFrom(item.GetPermission()) // 权限标识
	}
	if item != nil && item.Type != nil {
		daoItem.Type = item.Type // 菜单类型(1:目录/2: 菜单/3: 按钮)
	}
	if item != nil && item.Sort != nil {
		daoItem.Sort = item.Sort // 显示顺序
	}
	if item != nil && item.ParentId != nil {
		daoItem.ParentId = item.ParentId // 父菜单ID
	}
	if item != nil && item.Path != nil {
		daoItem.Path = null.StringFrom(item.GetPath()) // 路由地址
	}
	if item != nil && item.Icon != nil {
		daoItem.Icon = null.StringFrom(item.GetIcon()) // 菜单图标
	}
	if item != nil && item.Component != nil {
		daoItem.Component = null.StringFrom(item.GetComponent()) // 组件路径
	}
	if item != nil && item.ComponentName != nil {
		daoItem.ComponentName = null.StringFrom(item.GetComponentName()) // 组件名
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 菜单状态(0开启/1关闭)
	}
	if item != nil && item.Hide != nil {
		daoItem.Hide = item.Hide // 是否隐藏(0:否/1是)
	}
	if item != nil && item.Link != nil {
		daoItem.Link = null.StringFrom(item.GetLink()) // 路由外链时填写的访问地址
	}
	if item != nil && item.KeepAlive != nil {
		daoItem.KeepAlive = item.KeepAlive // 是否缓存(0不/ 是)
	}
	if item != nil && item.Affix != nil {
		daoItem.Affix = item.Affix // 是否总是显示(0 不显示/1 显示)
	}
	if item != nil && item.ActivePath != nil {
		daoItem.ActivePath = null.StringFrom(item.GetActivePath()) // 激活链接
	}
	if item != nil && item.FullScreen != nil {
		daoItem.FullScreen = item.FullScreen // 是否全屏
	}
	if item != nil && item.Redirect != nil {
		daoItem.Redirect = null.StringFrom(item.GetRedirect()) // 路由重定向地址
	}
	if item != nil && item.Creator != nil {
		daoItem.Creator = null.StringFrom(item.GetCreator()) // 创建者
	}
	if item != nil && item.CreateTime != nil {
		daoItem.CreateTime = null.DateTimeFrom(util.GrpcTime(item.CreateTime)) // 创建时间
	}
	if item != nil && item.Updater != nil {
		daoItem.Updater = null.StringFrom(item.GetUpdater()) // 更新者
	}
	if item != nil && item.UpdateTime != nil {
		daoItem.UpdateTime = null.DateTimeFrom(util.GrpcTime(item.UpdateTime)) // 更新时间
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 是否删除
	}

	return daoItem
}

// SystemMenuProto 数据绑定
func SystemMenuProto(item dao.SystemMenu) *menu.SystemMenuObject {
	res := &menu.SystemMenuObject{}
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
func SystemMenuCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统菜单:system_menu:SystemMenuCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := menu.NewSystemMenuServiceClient(grpcClient)
	request := &menu.SystemMenuCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemMenu
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request.Data = SystemMenuProto(reqInfo)
	// 执行服务
	res, err := client.SystemMenuCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统菜单:system_menu:SystemMenuCreate")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemMenuUpdate 更新数据
func SystemMenuUpdate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统菜单:system_menu:SystemMenuUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := menu.NewSystemMenuServiceClient(grpcClient)
	systemMenuId := cast.ToInt64(newCtx.Param("systemMenuId"))
	request := &menu.SystemMenuUpdateRequest{}
	request.SystemMenuId = systemMenuId
	// 数据绑定
	var reqInfo dao.SystemMenu
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request.Data = SystemMenuProto(reqInfo)
	// 执行服务
	res, err := client.SystemMenuUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统菜单:system_menu:SystemMenuUpdate")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemMenuDelete 删除数据
func SystemMenuDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统菜单:system_menu:SystemMenuDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := menu.NewSystemMenuServiceClient(grpcClient)
	systemMenuId := cast.ToInt64(newCtx.Param("systemMenuId"))
	request := &menu.SystemMenuDeleteRequest{}
	request.SystemMenuId = systemMenuId
	// 执行服务
	res, err := client.SystemMenuDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统菜单:system_menu:SystemMenuDelete")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemMenu 查询单条数据
func SystemMenu(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统菜单:system_menu:SystemMenu")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := menu.NewSystemMenuServiceClient(grpcClient)
	systemMenuId := cast.ToInt64(newCtx.Param("systemMenuId"))
	request := &menu.SystemMenuRequest{}
	request.SystemMenuId = systemMenuId
	// 执行服务
	res, err := client.SystemMenu(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统菜单:system_menu:SystemMenu")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": SystemMenuDao(res.GetData()),
	})
}

// SystemMenuRecover 恢复数据
func SystemMenuRecover(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统菜单:system_menu:SystemMenuRecover")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := menu.NewSystemMenuServiceClient(grpcClient)
	systemMenuId := cast.ToInt64(newCtx.Param("systemMenuId"))
	request := &menu.SystemMenuRecoverRequest{}
	request.SystemMenuId = systemMenuId
	// 执行服务
	res, err := client.SystemMenuRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统菜单:system_menu:SystemMenuRecover")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemMenuList 列表数据
func SystemMenuList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统菜单:system_menu:SystemMenuList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := menu.NewSystemMenuServiceClient(grpcClient)
	// 构造查询条件
	request := &menu.SystemMenuListRequest{}

	if val, ok := newCtx.GetQuery("deleted"); ok {
		request.Deleted = proto.Int32(cast.ToInt32(val)) // 是否删除
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val)) // 菜单状态(0开启/1关闭)
	}
	if val, ok := newCtx.GetQuery("type"); ok {
		request.Type = proto.Int32(cast.ToInt32(val)) // 菜单类型(1:目录/2: 菜单/3: 按钮)
	}
	if val, ok := newCtx.GetQuery("sort"); ok {
		request.Sort = proto.Int32(cast.ToInt32(val)) // 显示顺序
	}
	if val, ok := newCtx.GetQuery("parentId"); ok {
		request.ParentId = proto.Int64(cast.ToInt64(val)) // 父菜单ID
	}

	// 执行服务
	res, err := client.SystemMenuList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统菜单:system_menu:SystemMenuList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemMenu
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, SystemMenuDao(item))
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": list,
	})
}
