package region

import (
	"cloud/code"
	"cloud/module/region"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// region 地区表

// SrvRegionServiceServer 地区表
type SrvRegionServiceServer struct {
	UnimplementedRegionServiceServer
	Server *xgrpc.Server
}

// RegionCreate 创建数据
func (srv SrvRegionServiceServer) RegionCreate(ctx context.Context, request *RegionCreateRequest) (*RegionCreateResponse, error) {
	req := RegionDao(request.GetData())
	data, err := region.RegionCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:地区表:region:RegionCreate")
		return &RegionCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &RegionCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// RegionUpdate 更新数据
func (srv SrvRegionServiceServer) RegionUpdate(ctx context.Context, request *RegionUpdateRequest) (*RegionUpdateResponse, error) {
	regionId := request.GetId()
	if regionId < 1 {
		return &RegionUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := RegionDao(request.GetData())
	_, err := region.RegionUpdate(ctx, regionId, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:地区表:region:RegionUpdate")
		return &RegionUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &RegionUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// RegionDelete 删除数据
func (srv SrvRegionServiceServer) RegionDelete(ctx context.Context, request *RegionDeleteRequest) (*RegionDeleteResponse, error) {
	regionId := request.GetId()
	if regionId < 1 {
		return &RegionDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := region.RegionDelete(ctx, regionId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": regionId,
			"err": err,
		}).Error("Sql:地区表:region:RegionDelete")
		return &RegionDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &RegionDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// Region 查询单条数据
func (srv SrvRegionServiceServer) Region(ctx context.Context, request *RegionRequest) (*RegionResponse, error) {
	regionId := request.GetId()
	if regionId < 1 {
		return &RegionResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := region.Region(ctx, regionId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": regionId,
			"err": err,
		}).Error("Sql:地区表:region:Region")
		return &RegionResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &RegionResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: RegionProto(res),
	}, nil
}

// RegionRecover 恢复数据
func (srv SrvRegionServiceServer) RegionRecover(ctx context.Context, request *RegionRecoverRequest) (*RegionRecoverResponse, error) {
	regionId := request.GetId()
	if regionId < 1 {
		return &RegionRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := region.RegionRecover(ctx, regionId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": regionId,
			"err": err,
		}).Error("Sql:地区表:region:RegionRecover")
		return &RegionRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &RegionRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvRegionServiceServer) RegionList(ctx context.Context, request *RegionListRequest) (*RegionListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}
	if request.ParentId != nil {
		condition["parentId"] = request.GetParentId()
	}

	// 获取数据集合
	list, err := region.RegionList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:地区表:region:RegionList")
		return &RegionListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*RegionObject
	for _, item := range list {
		res = append(res, RegionProto(item))
	}
	return &RegionListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
