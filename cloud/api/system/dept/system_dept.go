package dept

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/system/dept"

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
)

// system_dept 部门
// SystemDeptItem 查询单条数据
func SystemDeptItem(ctx context.Context, newCtx *app.RequestContext) (*dept.SystemDeptResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:system_dept:SystemDept")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := dept.NewSystemDeptServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &dept.SystemDeptRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemDept(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:system_dept:SystemDept")
		return nil, err
	}
	tenantId := cast.ToInt64(newCtx.GetInt64("tenantId"))
	data := dept.SystemDeptDao(res.GetData())
	if cast.ToInt64(data.TenantId) != tenantId {
		return nil, status.Error(code.ConvertToGrpc(code.NoPermission), code.StatusText(code.NoPermission))
	}
	return res, nil
}

// SystemDeptCreate 创建数据
func SystemDeptCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:system_dept:SystemDeptCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dept.NewSystemDeptServiceClient(grpcClient)
	request := &dept.SystemDeptCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemDept
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Id = nil
	reqInfo.Deleted = proto.Int32(0)
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户
	reqInfo.Creator = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	request.Data = dept.SystemDeptProto(reqInfo)
	// 执行服务
	res, err := client.SystemDeptCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:system_dept:SystemDeptCreate")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemDeptUpdate 更新数据
func SystemDeptUpdate(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := SystemDeptItem(ctx, newCtx); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:system_dept:SystemDeptUpdate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dept.NewSystemDeptServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &dept.SystemDeptUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SystemDept
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户
	reqInfo.Updater = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	reqInfo.Creator = null.StringFromPtr(nil)
	reqInfo.CreateTime = null.DateTimeFromPtr(nil)
	request.Data = dept.SystemDeptProto(reqInfo)
	// 执行服务
	res, err := client.SystemDeptUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:system_dept:SystemDeptUpdate")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemDeptDelete 删除数据
func SystemDeptDelete(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := SystemDeptItem(ctx, newCtx); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:system_dept:SystemDeptDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dept.NewSystemDeptServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &dept.SystemDeptDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemDeptDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:system_dept:SystemDeptDelete")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemDept 查询单条数据
func SystemDept(ctx context.Context, newCtx *app.RequestContext) {
	res, err := SystemDeptItem(ctx, newCtx)
	if err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": dept.SystemDeptDao(res.GetData()),
	})
}

// SystemDeptRecover 恢复数据
func SystemDeptRecover(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := SystemDeptItem(ctx, newCtx); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:system_dept:SystemDeptRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dept.NewSystemDeptServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &dept.SystemDeptRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemDeptRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:system_dept:SystemDeptRecover")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

func SystemDeptListSimple(ctx context.Context, newCtx *app.RequestContext) {
	SystemDeptList(ctx, newCtx)
}

// SystemDeptList 列表数据
func SystemDeptList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:system_dept:SystemDeptList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dept.NewSystemDeptServiceClient(grpcClient)
	// 构造查询条件
	request := &dept.SystemDeptListRequest{}
	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                            // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val)) // 部门状态
	}
	if val, ok := newCtx.GetQuery("parentId"); ok {
		request.ParentId = proto.Int64(cast.ToInt64(val)) // 父部门ID
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val) // 部门名称
	}
	// 执行服务
	res, err := client.SystemDeptList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:system_dept:SystemDeptList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemDept
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, dept.SystemDeptDao(item))
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": SystemDeptTree(list, 0),
	})
}

// SystemDeptTree 树形菜单
func SystemDeptTree(list []*dao.SystemDept, pid int64) []*dao.SystemDept {
	var tree []*dao.SystemDept
	for _, item := range list {
		if *item.ParentId == pid {
			item.Children = SystemDeptTree(list, *item.Id)
			tree = append(tree, item)
		}
	}
	return tree
}
