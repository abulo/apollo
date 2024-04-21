package dict

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/system/dict"

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

// SystemDictCreate 创建数据
func SystemDictCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:字典数据表:system_dict:SystemDictCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dict.NewSystemDictServiceClient(grpcClient)
	request := &dict.SystemDictCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemDict
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Creator = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	request.Data = dict.SystemDictProto(reqInfo)
	// 执行服务
	res, err := client.SystemDictCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:字典数据表:system_dict:SystemDictCreate")
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

// SystemDictUpdate 更新数据
func SystemDictUpdate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:字典数据表:system_dict:SystemDictUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dict.NewSystemDictServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &dict.SystemDictUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SystemDict
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Updater = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	request.Data = dict.SystemDictProto(reqInfo)
	// 执行服务
	res, err := client.SystemDictUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:字典数据表:system_dict:SystemDictUpdate")
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

// SystemDictDelete 删除数据
func SystemDictDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:字典数据表:system_dict:SystemDictDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dict.NewSystemDictServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &dict.SystemDictDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemDictDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:字典数据表:system_dict:SystemDictDelete")
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

// SystemDict 查询单条数据
func SystemDict(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:字典数据表:system_dict:SystemDict")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dict.NewSystemDictServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &dict.SystemDictRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemDict(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:字典数据表:system_dict:SystemDict")
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
		"data": dict.SystemDictDao(res.GetData()),
	})
}

// SystemDictList 列表数据
func SystemDictList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:字典数据表:system_dict:SystemDictList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dict.NewSystemDictServiceClient(grpcClient)
	// 构造查询条件
	request := &dict.SystemDictListRequest{}

	if val, ok := newCtx.GetQuery("dictType"); ok {
		request.DictType = proto.String(val) // 字典类型
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val)) // 状态（0正常 1停用）
	}

	// 执行服务
	res, err := client.SystemDictList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:字典数据表:system_dict:SystemDictList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemDict
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, dict.SystemDictDao(item))
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": list,
	})
}

func SystemDictAll(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:字典类型:system_dict_type:SystemDictAll")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dict.NewSystemDictTypeServiceClient(grpcClient)
	// 构造查询条件
	request := &dict.SystemDictTypeListRequest{}
	requestTotal := &dict.SystemDictTypeListTotalRequest{}
	request.Status = proto.Int32(0)      // 状态（0正常 1停用）
	requestTotal.Status = proto.Int32(0) // 状态（0正常 1停用）
	// 执行服务,获取数据量
	resTotal, err := client.SystemDictTypeListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:字典类型:system_dict_type:SystemDictAll")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var total int64
	request.PageNum = proto.Int64(1)
	if resTotal.GetCode() == code.Success {
		total = resTotal.GetData()
	}
	request.PageSize = proto.Int64(total)
	// 执行服务
	res, err := client.SystemDictTypeList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:字典类型:system_dict_type:SystemDictAll")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	dictRes := make(map[string][]*dao.SystemDict)
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		//链接服务
		clientDict := dict.NewSystemDictServiceClient(grpcClient)
		// 构造查询条件
		requestDict := &dict.SystemDictListRequest{}
		requestDict.Status = proto.Int32(0) // 状态（0正常 1停用）
		for _, item := range rpcList {
			requestDict.DictType = item.Type
			// 执行服务
			resDict, errDict := clientDict.SystemDictList(ctx, requestDict)
			if errDict != nil {
				continue
			}
			var dictList []*dao.SystemDict
			if resDict.GetCode() == code.Success {
				rpcDictList := resDict.GetData()
				for _, item := range rpcDictList {
					dictList = append(dictList, dict.SystemDictDao(item))
				}
			}
			if len(dictList) > 0 {
				dictRes[cast.ToString(item.Type)] = dictList
			}
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": dictRes,
	})
}
