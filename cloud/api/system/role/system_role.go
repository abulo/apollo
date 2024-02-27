package role

import (
	"context"
	"encoding/json"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/system/role"

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

// system_role 系统角色

// SystemRoleDao 数据转换
func SystemRoleDao(item *role.SystemRoleObject) *dao.SystemRole {
	daoItem := &dao.SystemRole{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 角色编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 角色名称
	}
	if item != nil && item.Code != nil {
		daoItem.Code = item.Code // 角色权限字符串
	}
	if item != nil && item.Sort != nil {
		daoItem.Sort = item.Sort // 显示顺序
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 角色状态（0正常 1停用）
	}
	if item != nil && item.Type != nil {
		daoItem.Type = item.Type // 角色类型(1内置/2定义)
	}
	if item != nil && item.Remark != nil {
		daoItem.Remark = null.StringFrom(item.GetRemark()) // 备注
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

	return daoItem
}

// SystemRoleProto 数据绑定
func SystemRoleProto(item dao.SystemRole) *role.SystemRoleObject {
	res := &role.SystemRoleObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.Code != nil {
		res.Code = item.Code
	}
	if item.Sort != nil {
		res.Sort = item.Sort
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Type != nil {
		res.Type = item.Type
	}
	if item.Remark.IsValid() {
		res.Remark = item.Remark.Ptr()
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

	return res
}

// SystemRoleMenuProto 数据绑定
func SystemRoleMenuProto(item dao.SystemRole) *role.SystemRoleMenuCreateRequest {
	res := &role.SystemRoleMenuCreateRequest{}
	if item.MenuIds.IsValid() {
		res.SystemMenuIds = *item.MenuIds.Ptr()
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

	return res
}

// SystemRoleCreate 创建数据
func SystemRoleCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统角色:system_role:SystemRoleCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSystemRoleServiceClient(grpcClient)
	request := &role.SystemRoleCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemRole
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request.Data = SystemRoleProto(reqInfo)
	// 执行服务
	res, err := client.SystemRoleCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统角色:system_role:SystemRoleCreate")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	if res.GetCode() == code.Success {
		// 获取角色ID
		roleId := res.GetData()
		clientRoleMenu := role.NewSystemRoleMenuServiceClient(grpcClient)
		requestRoleMenu := SystemRoleMenuProto(reqInfo)
		requestRoleMenu.SystemRoleId = proto.Int64(roleId)
		clientRoleMenu.SystemRoleMenuCreate(ctx, requestRoleMenu)
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemRoleUpdate 更新数据
func SystemRoleUpdate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统角色:system_role:SystemRoleUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSystemRoleServiceClient(grpcClient)
	systemRoleId := cast.ToInt64(newCtx.Param("systemRoleId"))
	request := &role.SystemRoleUpdateRequest{}
	request.SystemRoleId = systemRoleId
	// 数据绑定
	var reqInfo dao.SystemRole
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request.Data = SystemRoleProto(reqInfo)
	// 执行服务
	res, err := client.SystemRoleUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统角色:system_role:SystemRoleUpdate")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	if res.GetCode() == code.Success {
		// 获取角色ID
		clientRoleMenu := role.NewSystemRoleMenuServiceClient(grpcClient)
		requestRoleMenu := SystemRoleMenuProto(reqInfo)
		requestRoleMenu.SystemRoleId = proto.Int64(systemRoleId)
		clientRoleMenu.SystemRoleMenuCreate(ctx, requestRoleMenu)
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemRoleDelete 删除数据
func SystemRoleDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统角色:system_role:SystemRoleDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSystemRoleServiceClient(grpcClient)
	systemRoleId := cast.ToInt64(newCtx.Param("systemRoleId"))
	request := &role.SystemRoleDeleteRequest{}
	request.SystemRoleId = systemRoleId
	// 执行服务
	res, err := client.SystemRoleDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统角色:system_role:SystemRoleDelete")
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

// SystemRole 查询单条数据
func SystemRole(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统角色:system_role:SystemRole")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSystemRoleServiceClient(grpcClient)
	systemRoleId := cast.ToInt64(newCtx.Param("systemRoleId"))
	request := &role.SystemRoleRequest{}
	request.SystemRoleId = systemRoleId
	// 执行服务
	res, err := client.SystemRole(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统角色:system_role:SystemRole")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}

	clientRoleMenu := role.NewSystemRoleMenuServiceClient(grpcClient)
	requestRoleMenu := &role.SystemRoleMenuListRequest{}
	requestRoleMenu.SystemRoleId = proto.Int64(systemRoleId)
	var listMenuId []int64
	if resRoleMenu, err := clientRoleMenu.SystemRoleMenuList(ctx, requestRoleMenu); err == nil {
		if resRoleMenu.GetCode() == code.Success {
			rpcList := resRoleMenu.GetData()
			for _, item := range rpcList {
				listMenuId = append(listMenuId, *item.SystemMenuId)
			}
		}
	}
	byteMenuIds, _ := json.Marshal(listMenuId)
	roleInfo := SystemRoleDao(res.GetData())
	roleInfo.MenuIds = null.JSONFrom(byteMenuIds)
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": roleInfo,
	})
}

// SystemRoleList 列表数据
func SystemRoleList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统角色:system_role:SystemRoleList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSystemRoleServiceClient(grpcClient)
	// 构造查询条件
	request := &role.SystemRoleListRequest{}

	if val, ok := newCtx.GetQuery("type"); ok {
		request.Type = proto.Int32(cast.ToInt32(val)) // 角色类型(1内置/2定义)
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val)) // 角色状态（0正常 1停用）
	}

	// 执行服务
	res, err := client.SystemRoleList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统角色:system_role:SystemRoleList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemRole
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, SystemRoleDao(item))
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": list,
	})
}
