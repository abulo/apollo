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
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_dict 字典数据表

// SystemDictDao 数据转换
func SystemDictDao(item *dict.SystemDictObject) *dao.SystemDict {
	daoItem := &dao.SystemDict{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 字典编码
	}
	if item != nil && item.Sort != nil {
		daoItem.Sort = item.Sort // 字典排序
	}
	if item != nil && item.Label != nil {
		daoItem.Label = item.Label // 字典标签
	}
	if item != nil && item.Value != nil {
		daoItem.Value = item.Value // 字典键值
	}
	if item != nil && item.DictType != nil {
		daoItem.DictType = item.DictType // 字典类型
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态（0正常 1停用）
	}
	if item != nil && item.ColorType != nil {
		daoItem.ColorType = null.StringFrom(item.GetColorType()) // 颜色类型
	}
	if item != nil && item.CssClass != nil {
		daoItem.CssClass = null.StringFrom(item.GetCssClass()) // css 样式
	}
	if item != nil && item.Remark != nil {
		daoItem.Remark = null.StringFrom(item.GetRemark()) // 备注
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

// SystemDictProto 数据绑定
func SystemDictProto(item dao.SystemDict) *dict.SystemDictObject {
	res := &dict.SystemDictObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Sort != nil {
		res.Sort = item.Sort
	}
	if item.Label != nil {
		res.Label = item.Label
	}
	if item.Value != nil {
		res.Value = item.Value
	}
	if item.DictType != nil {
		res.DictType = item.DictType
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.ColorType.IsValid() {
		res.ColorType = item.ColorType.Ptr()
	}
	if item.CssClass.IsValid() {
		res.CssClass = item.CssClass.Ptr()
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
	reqInfo.Creator = null.StringFrom(newCtx.GetString("systemUserName"))
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	request.Data = SystemDictProto(reqInfo)
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
	systemDictId := cast.ToInt64(newCtx.Param("systemDictId"))
	request := &dict.SystemDictUpdateRequest{}
	request.SystemDictId = systemDictId
	// 数据绑定
	var reqInfo dao.SystemDict
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Updater = null.StringFrom(newCtx.GetString("systemUserName"))
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	request.Data = SystemDictProto(reqInfo)
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
	systemDictId := cast.ToInt64(newCtx.Param("systemDictId"))
	request := &dict.SystemDictDeleteRequest{}
	request.SystemDictId = systemDictId
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
	systemDictId := cast.ToInt64(newCtx.Param("systemDictId"))
	request := &dict.SystemDictRequest{}
	request.SystemDictId = systemDictId
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
		"data": SystemDictDao(res.GetData()),
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
			list = append(list, SystemDictDao(item))
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
					dictList = append(dictList, SystemDictDao(item))
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
