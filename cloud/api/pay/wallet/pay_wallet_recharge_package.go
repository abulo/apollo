package wallet

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/pay/wallet"

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

// pay_wallet_recharge_package 充值套餐表
// PayWalletRechargePackage 查询单条数据
func PayWalletRechargePackageItem(ctx context.Context, newCtx *app.RequestContext) (*wallet.PayWalletRechargePackageResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackage")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := wallet.NewPayWalletRechargePackageServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &wallet.PayWalletRechargePackageRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayWalletRechargePackage(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackage")
		return nil, err
	}
	tenantId := cast.ToInt64(newCtx.GetInt64("tenantId"))
	data := wallet.PayWalletRechargePackageDao(res.GetData())
	if cast.ToInt64(data.TenantId) != tenantId {
		return nil, status.Error(code.ConvertToGrpc(code.NoPermission), code.StatusText(code.NoPermission))
	}
	return res, nil
}

// PayWalletRechargePackageCreate 创建数据
func PayWalletRechargePackageCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := wallet.NewPayWalletRechargePackageServiceClient(grpcClient)
	request := &wallet.PayWalletRechargePackageCreateRequest{}
	// 数据绑定
	var reqInfo dao.PayWalletRechargePackage
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Deleted = proto.Int32(0)
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户
	reqInfo.Creator = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	request.Data = wallet.PayWalletRechargePackageProto(reqInfo)
	// 执行服务
	res, err := client.PayWalletRechargePackageCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageCreate")
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

// PayWalletRechargePackageUpdate 更新数据
func PayWalletRechargePackageUpdate(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := PayWalletRechargePackageItem(ctx, newCtx); err != nil {
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
		}).Error("Grpc:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageUpdate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := wallet.NewPayWalletRechargePackageServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &wallet.PayWalletRechargePackageUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.PayWalletRechargePackage
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
	request.Data = wallet.PayWalletRechargePackageProto(reqInfo)
	// 执行服务
	res, err := client.PayWalletRechargePackageUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageUpdate")
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

// PayWalletRechargePackageDelete 删除数据
func PayWalletRechargePackageDelete(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := PayWalletRechargePackageItem(ctx, newCtx); err != nil {
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
		}).Error("Grpc:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := wallet.NewPayWalletRechargePackageServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &wallet.PayWalletRechargePackageDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayWalletRechargePackageDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageDelete")
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

// PayWalletRechargePackage 查询单条数据
func PayWalletRechargePackage(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	res, err := PayWalletRechargePackageItem(ctx, newCtx)
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
		"data": wallet.PayWalletRechargePackageDao(res.GetData()),
	})
}

// PayWalletRechargePackageRecover 恢复数据
func PayWalletRechargePackageRecover(ctx context.Context, newCtx *app.RequestContext) {
	if _, err := PayWalletRechargePackageItem(ctx, newCtx); err != nil {
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
		}).Error("Grpc:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := wallet.NewPayWalletRechargePackageServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &wallet.PayWalletRechargePackageRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.PayWalletRechargePackageRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageRecover")
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

// PayWalletRechargePackageList 列表数据
func PayWalletRechargePackageList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := wallet.NewPayWalletRechargePackageServiceClient(grpcClient)
	// 构造查询条件
	request := &wallet.PayWalletRechargePackageListRequest{}
	requestTotal := &wallet.PayWalletRechargePackageListTotalRequest{}

	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))      // 租户ID
	requestTotal.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                                 // 删除状态
	requestTotal.Deleted = proto.Int32(0)                            // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
			requestTotal.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 状态
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 状态
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val)      // 套餐名称
		requestTotal.Name = proto.String(val) // 套餐名称
	}

	// 执行服务,获取数据量
	resTotal, err := client.PayWalletRechargePackageListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var total int64
	paginationRequest := &pagination.PaginationRequest{}
	paginationRequest.PageNum = proto.Int64(cast.ToInt64(newCtx.Query("pageNum")))
	paginationRequest.PageSize = proto.Int64(cast.ToInt64(newCtx.Query("pageSize")))
	request.Pagination = paginationRequest
	if resTotal.GetCode() == code.Success {
		total = resTotal.GetData()
	}
	// 执行服务
	res, err := client.PayWalletRechargePackageList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.PayWalletRechargePackage
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, wallet.PayWalletRechargePackageDao(item))
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"total":    total,
			"list":     list,
			"pageNum":  paginationRequest.PageNum,
			"pageSize": paginationRequest.PageSize,
		},
	})
}

// PayWalletRechargePackageListSimple 精简列表数据
func PayWalletRechargePackageListSimple(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := wallet.NewPayWalletRechargePackageServiceClient(grpcClient)
	// 构造查询条件
	request := &wallet.PayWalletRechargePackageListRequest{}

	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                            // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val)) // 状态
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val) // 套餐名称
	}
	// 执行服务
	res, err := client.PayWalletRechargePackageList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:充值套餐表:pay_wallet_recharge_package:PayWalletRechargePackageList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.PayWalletRechargePackage
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, wallet.PayWalletRechargePackageDao(item))
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": list,
	})
}
