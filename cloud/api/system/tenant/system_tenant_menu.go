package tenant

import (
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/system/menu"
	"cloud/service/system/tenant"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

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
	tenantClient := tenant.NewSystemTenantServiceClient(grpcClient)
	systemTenantId := newCtx.GetInt64("systemTenantId")
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
	// 这里有一个数据需要来判断
	tenantItem := tenant.SystemTenantDao(tenantRes.GetData())
	// 获取套餐服务
	tenantPackageId := tenantItem.SystemTenantPackageId
	//链接服务
	menuClient := menu.NewSystemMenuServiceClient(grpcClient)
	// 构造查询条件
	menuRequest := &menu.SystemMenuListRequest{}
	menuRequest.Deleted = nil
	menuRequest.Status = nil
	var list []*dao.SystemMenu
	var menuIds []int64
	if cast.ToInt64(tenantPackageId) != 0 {
		menuRequest.Deleted = proto.Int32(0)
		menuRequest.Status = proto.Int32(0)
		//这里还需要去
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
			if cast.ToInt64(tenantPackageId) != 0 {
				if !util.InArray(*item.Id, menuIds) {
					continue
				}
			}
			list = append(list, menu.SystemMenuDao(item))
		}
	}
	newList := SystemMenuTree(list, 0)
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": menuRes.GetCode(),
		"msg":  menuRes.GetMsg(),
		"data": newList,
	})
}

// SystemMenuTree 树形菜单
func SystemMenuTree(list []*dao.SystemMenu, pid int64) []*dao.SystemMenu {
	var tree []*dao.SystemMenu
	for _, item := range list {
		if *item.ParentId == pid {
			item.Children = SystemMenuTree(list, *item.Id)
			tree = append(tree, item)
		}
	}
	return tree
}