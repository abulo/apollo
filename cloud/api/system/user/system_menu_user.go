package user

import (
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/system/menu"
	"context"
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
func SystemUserMenuDao(item *menu.SystemMenuObject) dao.SystemMenuTree {
	daoItem := dao.SystemMenuTree{}
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
	daoMetaItem := dao.SystemMenuTreeMeta{}
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

	client := menu.NewSystemMenuServiceClient(grpcClient)
	// 构造查询条件
	request := &menu.SystemMenuListRequest{}
	// 构造查询条件
	request.Deleted = proto.Int32(0)

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

	var list []dao.SystemMenuTree
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			if item.GetType() == 3 {
				continue
			}
			list = append(list, SystemUserMenuDao(item))
		}
	}
	newList := SystemUserMenuTree(list, 0)
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": newList,
	})
}

// SystemMenuTree 树形菜单
func SystemUserMenuTree(list []dao.SystemMenuTree, pid int64) []dao.SystemMenuTree {
	var tree []dao.SystemMenuTree
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
		}).Error("Grpc:菜单权限表:system_menu:SystemUserBtn")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}

	client := menu.NewSystemMenuServiceClient(grpcClient)
	// 构造查询条件
	request := &menu.SystemMenuListRequest{}
	// 构造查询条件
	request.Deleted = proto.Int32(0)

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
	var list []string
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
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
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": resList,
	})
}
