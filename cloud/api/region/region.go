package region

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/region"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// region 地区表
// RegionCreate 创建数据
func RegionCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:地区表:region:RegionCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := region.NewRegionServiceClient(grpcClient)
	request := &region.RegionCreateRequest{}
	// 数据绑定
	var reqInfo dao.Region
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request.Data = region.RegionProto(reqInfo)
	// 执行服务
	res, err := client.RegionCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:地区表:region:RegionCreate")
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

// RegionUpdate 更新数据
func RegionUpdate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:地区表:region:RegionUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := region.NewRegionServiceClient(grpcClient)
	regionId := cast.ToInt64(newCtx.Param("id"))
	request := &region.RegionUpdateRequest{}
	request.Id = regionId
	// 数据绑定
	var reqInfo dao.Region
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request.Data = region.RegionProto(reqInfo)
	// 执行服务
	res, err := client.RegionUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:地区表:region:RegionUpdate")
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

// RegionDelete 删除数据
func RegionDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:地区表:region:RegionDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := region.NewRegionServiceClient(grpcClient)
	regionId := cast.ToInt64(newCtx.Param("id"))
	request := &region.RegionDeleteRequest{}
	request.Id = regionId
	// 执行服务
	res, err := client.RegionDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:地区表:region:RegionDelete")
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

// Region 查询单条数据
func Region(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:地区表:region:Region")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := region.NewRegionServiceClient(grpcClient)
	regionId := cast.ToInt64(newCtx.Param("id"))
	request := &region.RegionRequest{}
	request.Id = regionId
	// 执行服务
	res, err := client.Region(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:地区表:region:Region")
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
		"data": region.RegionDao(res.GetData()),
	})
}

// RegionRecover 恢复数据
func RegionRecover(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:地区表:region:RegionRecover")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := region.NewRegionServiceClient(grpcClient)
	regionId := cast.ToInt64(newCtx.Param("id"))
	request := &region.RegionRecoverRequest{}
	request.Id = regionId
	// 执行服务
	res, err := client.RegionRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:地区表:region:RegionRecover")
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
func RegionListSimple(ctx context.Context, newCtx *app.RequestContext) {
	RegionList(ctx, newCtx)
}

// RegionList 列表数据
func RegionList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:地区表:region:RegionList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := region.NewRegionServiceClient(grpcClient)
	// 构造查询条件
	request := &region.RegionListRequest{}

	request.Deleted = proto.Int32(0)
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val)) // 状态;
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val) // 区域名称
	}

	// 执行服务
	res, err := client.RegionList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:地区表:region:RegionList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.Region
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, region.RegionDao(item))
		}
	}
	var newList []*dao.Region
	newList = list
	tree := false
	if val, ok := newCtx.GetQuery("tree"); ok {
		tree = cast.ToBool(val)
	}
	if tree {
		newList = RegionTree(list, 0)
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": newList,
	})
}

// RegionTree 树形菜单
func RegionTree(list []*dao.Region, pid int64) []*dao.Region {
	var tree []*dao.Region
	for _, item := range list {
		if *item.ParentId == pid {
			item.Children = RegionTree(list, *item.Id)
			tree = append(tree, item)
		}
	}
	return tree
}
