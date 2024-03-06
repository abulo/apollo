package tenant

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/system/tenant"

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

// system_tenant 租户

// SystemTenantDao 数据转换
func SystemTenantDao(item *tenant.SystemTenantObject) *dao.SystemTenant {
	daoItem := &dao.SystemTenant{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 租户编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 租户名称
	}
	if item != nil && item.SystemUserId != nil {
		daoItem.SystemUserId = null.Int64From(item.GetSystemUserId()) // 联系人ID
	}
	if item != nil && item.ContactName != nil {
		daoItem.ContactName = item.ContactName // 联系人
	}
	if item != nil && item.ContactMobile != nil {
		daoItem.ContactMobile = item.ContactMobile // 租户联系电话
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态（0正常 1停用）
	}
	if item != nil && item.Domain != nil {
		daoItem.Domain = null.StringFrom(item.GetDomain()) // 域名
	}
	if item != nil && item.ExpireDate != nil {
		daoItem.ExpireDate = null.DateFrom(util.GrpcTime(item.ExpireDate)) // 过期时间
	}
	if item != nil && item.AccountCont != nil {
		daoItem.AccountCont = item.AccountCont // 账号数量
	}
	if item != nil && item.SystemTenantPackageId != nil {
		daoItem.SystemTenantPackageId = item.SystemTenantPackageId // 套餐编号
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 是否删除(0否 1是)
	}
	if item != nil && item.Creator != nil {
		daoItem.Creator = null.StringFrom(item.GetCreator()) // 创建人
	}
	if item != nil && item.CreateTime != nil {
		daoItem.CreateTime = null.DateTimeFrom(util.GrpcTime(item.CreateTime)) // 创建时间
	}
	if item != nil && item.Updater != nil {
		daoItem.Updater = null.StringFrom(item.GetUpdater()) // 更新人
	}
	if item != nil && item.UpdateTime != nil {
		daoItem.UpdateTime = null.DateTimeFrom(util.GrpcTime(item.UpdateTime)) // 更新时间
	}

	return daoItem
}

// SystemTenantProto 数据绑定
func SystemTenantProto(item dao.SystemTenant) *tenant.SystemTenantObject {
	res := &tenant.SystemTenantObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.SystemUserId.IsValid() {
		res.SystemUserId = item.SystemUserId.Ptr()
	}
	if item.ContactName != nil {
		res.ContactName = item.ContactName
	}
	if item.ContactMobile != nil {
		res.ContactMobile = item.ContactMobile
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Domain.IsValid() {
		res.Domain = item.Domain.Ptr()
	}
	if item.ExpireDate.IsValid() {
		res.ExpireDate = timestamppb.New(*item.ExpireDate.Ptr())
	}
	if item.AccountCont != nil {
		res.AccountCont = item.AccountCont
	}
	if item.SystemTenantPackageId != nil {
		res.SystemTenantPackageId = item.SystemTenantPackageId
	}
	if item.Deleted != nil {
		res.Deleted = item.Deleted
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
	if item.Username != nil {
		res.Username = item.Username
	}
	if item.Password != nil {
		res.Password = item.Password
	}

	return res
}

// SystemTenantCreate 创建数据
func SystemTenantCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:system_tenant:SystemTenantCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSystemTenantServiceClient(grpcClient)
	request := &tenant.SystemTenantCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemTenant
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request.Data = SystemTenantProto(reqInfo)
	// 执行服务
	res, err := client.SystemTenantCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenantCreate")
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

// SystemTenantUpdate 更新数据
func SystemTenantUpdate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:system_tenant:SystemTenantUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSystemTenantServiceClient(grpcClient)
	systemTenantId := cast.ToInt64(newCtx.Param("systemTenantId"))
	request := &tenant.SystemTenantUpdateRequest{}
	request.SystemTenantId = systemTenantId
	// 数据绑定
	var reqInfo dao.SystemTenant
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request.Data = SystemTenantProto(reqInfo)
	// 执行服务
	res, err := client.SystemTenantUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenantUpdate")
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

// SystemTenantDelete 删除数据
func SystemTenantDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:system_tenant:SystemTenantDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSystemTenantServiceClient(grpcClient)
	systemTenantId := cast.ToInt64(newCtx.Param("systemTenantId"))
	request := &tenant.SystemTenantDeleteRequest{}
	request.SystemTenantId = systemTenantId
	// 执行服务
	res, err := client.SystemTenantDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenantDelete")
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

// SystemTenant 查询单条数据
func SystemTenant(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:system_tenant:SystemTenant")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSystemTenantServiceClient(grpcClient)
	systemTenantId := cast.ToInt64(newCtx.Param("systemTenantId"))
	request := &tenant.SystemTenantRequest{}
	request.SystemTenantId = systemTenantId
	// 执行服务
	res, err := client.SystemTenant(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenant")
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
		"data": SystemTenantDao(res.GetData()),
	})
}

// SystemTenantRecover 恢复数据
func SystemTenantRecover(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:system_tenant:SystemTenantRecover")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSystemTenantServiceClient(grpcClient)
	systemTenantId := cast.ToInt64(newCtx.Param("systemTenantId"))
	request := &tenant.SystemTenantRecoverRequest{}
	request.SystemTenantId = systemTenantId
	// 执行服务
	res, err := client.SystemTenantRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenantRecover")
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

// SystemTenantList 列表数据
func SystemTenantList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:system_tenant:SystemTenantList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSystemTenantServiceClient(grpcClient)
	// 构造查询条件
	request := &tenant.SystemTenantListRequest{}
	requestTotal := &tenant.SystemTenantListTotalRequest{}
	// 是否删除(0否 1是)
	request.Deleted = proto.Int32(0)
	requestTotal.Deleted = proto.Int32(0)
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
			requestTotal.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 状态（0正常 1停用）
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 状态（0正常 1停用）
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val)      // 租户名称
		requestTotal.Name = proto.String(val) // 租户名称
	}
	if val, ok := newCtx.GetQuery("beginExpireDate"); ok {
		request.BeginExpireDate = timestamppb.New(cast.ToTime(val))      // 过期时间
		requestTotal.BeginExpireDate = timestamppb.New(cast.ToTime(val)) // 过期时间
	}
	if val, ok := newCtx.GetQuery("finishExpireDate"); ok {
		request.FinishExpireDate = timestamppb.New(cast.ToTime(val))      // 过期时间
		requestTotal.FinishExpireDate = timestamppb.New(cast.ToTime(val)) // 过期时间
	}
	if val, ok := newCtx.GetQuery("systemTenantPackageId"); ok {
		request.SystemTenantPackageId = proto.Int64(cast.ToInt64(val))      // 套餐编号
		requestTotal.SystemTenantPackageId = proto.Int64(cast.ToInt64(val)) // 套餐编号
	}
	// 执行服务,获取数据量
	resTotal, err := client.SystemTenantListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenantList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var total int64
	request.PageNum = proto.Int64(cast.ToInt64(newCtx.Query("pageNum")))
	request.PageSize = proto.Int64(cast.ToInt64(newCtx.Query("pageSize")))
	if resTotal.GetCode() == code.Success {
		total = resTotal.GetData()
	}
	// 执行服务
	res, err := client.SystemTenantList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenantList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemTenant
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, SystemTenantDao(item))
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"total":    total,
			"list":     list,
			"pageNum":  request.PageNum,
			"pageSize": request.PageSize,
		},
	})
}
