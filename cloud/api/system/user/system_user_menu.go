package user

import (
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/system/menu"
	"cloud/service/system/role"
	"cloud/service/system/tenant"
	"cloud/service/system/user"
	"context"
	"encoding/json"
	"strings"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// SystemUserMenuDao 数据转换
func SystemUserMenuDao(item *menu.SystemMenuObject) *dao.SystemMenuTree {
	daoItem := &dao.SystemMenuTree{}
	if !util.Empty(item.Id) {
		daoItem.Id = item.GetId() // 菜单ID
	}
	if !util.Empty(item.ParentId) {
		daoItem.ParentId = item.GetParentId() // 父菜单ID
	}
	if !util.Empty(item.Path) {
		daoItem.Path = item.GetPath() // 路由地址
	}
	if !util.Empty(item.Name) {
		daoItem.Name = item.GetComponentName() // 路由名称
	}
	if !util.Empty(item.Component) {
		daoItem.Component = item.GetComponent() // 组件路径
	}
	if !util.Empty(item.Redirect) {
		daoItem.Redirect = item.GetRedirect() // 路由重定向地址
	}
	if !util.Empty(item.Type) {
		daoItem.Type = item.GetType()
	}
	daoMetaItem := &dao.SystemMenuTreeMeta{}
	if !util.Empty(item.Icon) {
		daoMetaItem.Icon = item.GetIcon() // 菜单图标
	}
	if !util.Empty(item.Name) {
		daoMetaItem.Title = item.GetName() // 菜单标题
	}
	if !util.Empty(item.Link) {
		daoMetaItem.IsLink = item.GetLink() // 是否外链
	}
	daoMetaItem.IsHide = cast.ToBool(cast.ToInt(item.GetHide()))           // 是否隐藏(0:否/1是)
	daoMetaItem.IsFull = cast.ToBool(cast.ToInt(item.GetFullScreen()))     // 是否全屏
	daoMetaItem.IsAffix = cast.ToBool(cast.ToInt(item.GetAffix()))         // 是否固定
	daoMetaItem.IsKeepAlive = cast.ToBool(cast.ToInt(item.GetKeepAlive())) // 是否缓存
	if !util.Empty(item.ActivePath) {
		daoMetaItem.ActiveMenu = item.GetActivePath() // 激活菜单
	}
	daoItem.Meta = daoMetaItem
	return daoItem
}

// SystemUserMenu 用户菜单
func SystemUserMenu(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:菜单权限表:system_menu:SystemUserMenu")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}

	systemTenantId := newCtx.GetInt64("systemTenantId") // 租户 Id
	systemUserId := newCtx.GetInt64("systemUserId")     // 用户 Id
	//链接服务
	tenantClient := tenant.NewSystemTenantServiceClient(grpcClient)
	tenantRequest := &tenant.SystemTenantRequest{}
	tenantRequest.SystemTenantId = systemTenantId
	// 执行服务
	tenantRes, tenantErr := tenantClient.SystemTenant(ctx, tenantRequest)
	if tenantErr != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": tenantRequest,
			"err": tenantErr,
		}).Error("GrpcCall:租户:system_tenant:SystemTenant")
		fromError := status.Convert(tenantErr)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	// 判断这个有没有值
	if tenantRes.GetCode() != code.Success {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	// 这是租户的信息
	tenantItem := tenant.SystemTenantDao(tenantRes.GetData())
	// 获取套餐服务
	tenantPackageId := tenantItem.SystemTenantPackageId
	//链接服务
	menuClient := menu.NewSystemMenuServiceClient(grpcClient)
	// 构造查询条件
	menuRequest := &menu.SystemMenuListRequest{}
	menuRequest.Deleted = nil
	menuRequest.Status = nil
	var list []*dao.SystemMenuTree
	var menuIds []int64
	var currentMenuIds []int64 // 当前租户所有的菜单 id
	if cast.ToInt64(tenantPackageId) != 0 {
		menuRequest.Deleted = proto.Int32(0)
		menuRequest.Status = proto.Int32(0)
		//链接服务
		tenantPackageClient := tenant.NewSystemTenantPackageServiceClient(grpcClient)
		systemTenantPackageId := cast.ToInt64(tenantPackageId)
		tenantPackageRequest := &tenant.SystemTenantPackageRequest{}
		tenantPackageRequest.SystemTenantPackageId = systemTenantPackageId
		// 执行服务
		if res, err := tenantPackageClient.SystemTenantPackage(ctx, tenantPackageRequest); err == nil {
			tenantPackageItem := tenant.SystemTenantPackageDao(res.GetData())
			if tenantPackageItem.SystemMenuIds.IsValid() {
				json.Unmarshal(tenantPackageItem.SystemMenuIds.JSON, &menuIds)
			}
		}
	}
	// 执行服务
	menuRes, menuErr := menuClient.SystemMenuList(ctx, menuRequest)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": menuRequest,
			"err": menuErr,
		}).Error("GrpcCall:系统菜单:system_menu:SystemMenuList")
		fromError := status.Convert(menuErr)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}

	if menuRes.GetCode() == code.Success {
		rpcList := menuRes.GetData()
		for _, item := range rpcList {
			currentMenuIds = append(currentMenuIds, *item.Id)
			if cast.ToInt64(tenantPackageId) != 0 {
				if !util.InArray(*item.Id, menuIds) {
					continue
				}
			}
			list = append(list, SystemUserMenuDao(item))
		}
	}
	if systemUserId != *tenantItem.SystemUserId.Ptr() {
		//链接服务
		userRoleClient := user.NewSystemUserRoleServiceClient(grpcClient)
		// 构造查询条件
		userRoleRequest := &user.SystemUserRoleListRequest{}
		userRoleRequest.SystemUserId = proto.Int64(systemUserId)     // 系统用户 ID
		userRoleRequest.SystemTenantId = proto.Int64(systemTenantId) // 租户
		userRoleRequest.Deleted = proto.Int32(0)                     // 删除

		roleMenuClient := role.NewSystemRoleMenuServiceClient(grpcClient)
		if userRoleRes, err := userRoleClient.SystemUserRoleList(ctx, userRoleRequest); err == nil {
			if userRoleRes.GetCode() == code.Success {
				rpcUserRoleList := userRoleRes.GetData()
				for _, item := range rpcUserRoleList {
					systemRoleId := item.GetSystemRoleId()
					//链接服务
					roleMenuRequest := &role.SystemRoleMenuListRequest{}
					roleMenuRequest.SystemTenantId = proto.Int64(systemTenantId) // 租户ID
					roleMenuRequest.Deleted = proto.Int32(0)                     // 删除状态
					roleMenuRequest.SystemRoleId = proto.Int64(systemRoleId)     // 角色ID
					if roleMenuRes, err := roleMenuClient.SystemRoleMenuList(ctx, roleMenuRequest); err == nil {
						if roleMenuRes.GetCode() == code.Success {
							rpcRoleMenuList := roleMenuRes.GetData()
							var newCurrentMenuIds []int64
							for _, item := range rpcRoleMenuList {
								newCurrentMenuIds = append(newCurrentMenuIds, *item.SystemMenuId)
							}
							currentMenuIds = newCurrentMenuIds
						}
					}
				}
			}
		}
	}
	var curList []*dao.SystemMenuTree
	for _, item := range list {
		if item.Type == 3 {
			continue
		}
		if util.InArray(item.Id, currentMenuIds) {
			curList = append(curList, item)
		}
	}
	newList := SystemUserMenuTree(curList, 0)
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": menuRes.GetCode(),
		"msg":  menuRes.GetMsg(),
		"data": newList,
	})
}

// SystemMenuTree 树形菜单
func SystemUserMenuTree(list []*dao.SystemMenuTree, pid int64) []*dao.SystemMenuTree {
	var tree []*dao.SystemMenuTree
	for _, item := range list {
		if item.ParentId == pid {
			item.Children = SystemUserMenuTree(list, item.Id)
			tree = append(tree, item)
		}
	}
	return tree
}

// SystemUserBtn 用户权限按钮
func SystemUserBtn(ctx context.Context, newCtx *app.RequestContext) {

	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:菜单权限表:system_menu:SystemUserMenu")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}

	systemTenantId := newCtx.GetInt64("systemTenantId") // 租户 Id
	systemUserId := newCtx.GetInt64("systemUserId")     // 用户 Id
	//链接服务
	tenantClient := tenant.NewSystemTenantServiceClient(grpcClient)
	tenantRequest := &tenant.SystemTenantRequest{}
	tenantRequest.SystemTenantId = systemTenantId
	// 执行服务
	tenantRes, tenantErr := tenantClient.SystemTenant(ctx, tenantRequest)
	if tenantErr != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": tenantRequest,
			"err": tenantErr,
		}).Error("GrpcCall:租户:system_tenant:SystemTenant")
		fromError := status.Convert(tenantErr)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	// 判断这个有没有值
	if tenantRes.GetCode() != code.Success {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	// 这是租户的信息
	tenantItem := tenant.SystemTenantDao(tenantRes.GetData())
	// 获取套餐服务
	tenantPackageId := tenantItem.SystemTenantPackageId
	//链接服务
	menuClient := menu.NewSystemMenuServiceClient(grpcClient)
	// 构造查询条件
	menuRequest := &menu.SystemMenuListRequest{}
	menuRequest.Deleted = nil
	menuRequest.Status = nil
	var listMenu []*menu.SystemMenuObject
	var menuIds []int64
	var currentMenuIds []int64 // 当前租户所有的菜单 id
	if cast.ToInt64(tenantPackageId) != 0 {
		menuRequest.Deleted = proto.Int32(0)
		menuRequest.Status = proto.Int32(0)
		//链接服务
		tenantPackageClient := tenant.NewSystemTenantPackageServiceClient(grpcClient)
		systemTenantPackageId := cast.ToInt64(tenantPackageId)
		tenantPackageRequest := &tenant.SystemTenantPackageRequest{}
		tenantPackageRequest.SystemTenantPackageId = systemTenantPackageId
		// 执行服务
		if res, err := tenantPackageClient.SystemTenantPackage(ctx, tenantPackageRequest); err == nil {
			tenantPackageItem := tenant.SystemTenantPackageDao(res.GetData())
			if tenantPackageItem.SystemMenuIds.IsValid() {
				json.Unmarshal(tenantPackageItem.SystemMenuIds.JSON, &menuIds)
			}
		}
	}
	// 执行服务
	menuRes, menuErr := menuClient.SystemMenuList(ctx, menuRequest)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": menuRequest,
			"err": menuErr,
		}).Error("GrpcCall:系统菜单:system_menu:SystemMenuList")
		fromError := status.Convert(menuErr)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}

	if menuRes.GetCode() == code.Success {
		rpcList := menuRes.GetData()
		for _, item := range rpcList {
			currentMenuIds = append(currentMenuIds, *item.Id)
			if cast.ToInt64(tenantPackageId) != 0 {
				if !util.InArray(*item.Id, menuIds) {
					continue
				}
			}
			// listMenu = append(listMenu, menu.SystemMenuDao(item))
			listMenu = append(listMenu, item)
		}
	}
	if systemUserId != *tenantItem.SystemUserId.Ptr() {
		//链接服务
		userRoleClient := user.NewSystemUserRoleServiceClient(grpcClient)
		// 构造查询条件
		userRoleRequest := &user.SystemUserRoleListRequest{}
		userRoleRequest.SystemUserId = proto.Int64(systemUserId)     // 系统用户 ID
		userRoleRequest.SystemTenantId = proto.Int64(systemTenantId) // 租户
		userRoleRequest.Deleted = proto.Int32(0)                     // 删除

		roleMenuClient := role.NewSystemRoleMenuServiceClient(grpcClient)
		if userRoleRes, err := userRoleClient.SystemUserRoleList(ctx, userRoleRequest); err == nil {
			if userRoleRes.GetCode() == code.Success {
				rpcUserRoleList := userRoleRes.GetData()
				for _, item := range rpcUserRoleList {
					systemRoleId := item.GetSystemRoleId()
					//链接服务
					roleMenuRequest := &role.SystemRoleMenuListRequest{}
					roleMenuRequest.SystemTenantId = proto.Int64(systemTenantId) // 租户ID
					roleMenuRequest.Deleted = proto.Int32(0)                     // 删除状态
					roleMenuRequest.SystemRoleId = proto.Int64(systemRoleId)     // 角色ID
					if roleMenuRes, err := roleMenuClient.SystemRoleMenuList(ctx, roleMenuRequest); err == nil {
						if roleMenuRes.GetCode() == code.Success {
							rpcRoleMenuList := roleMenuRes.GetData()
							var newCurrentMenuIds []int64
							for _, item := range rpcRoleMenuList {
								newCurrentMenuIds = append(newCurrentMenuIds, *item.SystemMenuId)
							}
							currentMenuIds = newCurrentMenuIds
						}
					}
				}
			}
		}
	}
	var curList []*menu.SystemMenuObject
	var list []string
	for _, item := range listMenu {
		if util.InArray(*item.Id, currentMenuIds) {
			curList = append(curList, item)
			if !util.Empty(item.GetPermission()) {
				list = append(list, item.GetPermission())
			}
		}
	}
	resList := make(map[string][]string)
	for _, v := range list {
		if !strings.Contains(v, ".") {
			resList[v] = append(resList[v], v)
		} else {
			tmp := strings.Split(v, ".")
			resList[tmp[0]] = append(resList[tmp[0]], v)
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": menuRes.GetCode(),
		"msg":  menuRes.GetMsg(),
		"data": resList,
	})
}
