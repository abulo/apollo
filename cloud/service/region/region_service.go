package region

import (
	"cloud/code"
	"cloud/dao"
	"cloud/module/region"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// region 地区表

// SrvRegionServiceServer 地区表
type SrvRegionServiceServer struct {
	UnimplementedRegionServiceServer
	Server *xgrpc.Server
}

func (srv SrvRegionServiceServer) RegionConvert(request *RegionObject) dao.Region {
	var res dao.Region

	if request != nil && request.Id != nil {
		res.Id = request.Id // 区域编号
	}
	if request != nil && request.Name != nil {
		res.Name = request.Name // 区域名称
	}
	if request != nil && request.ParentId != nil {
		res.ParentId = request.ParentId // 父级编号
	}
	if request != nil && request.WeatherCode != nil {
		res.WeatherCode = null.StringFrom(request.GetWeatherCode()) // 天气code
	}
	if request != nil && request.Creator != nil {
		res.Creator = null.StringFrom(request.GetCreator()) // 创建人
	}
	if request != nil && request.CreateTime != nil {
		res.CreateTime = null.DateTimeFrom(util.GrpcTime(request.CreateTime)) // 创建时间
	}
	if request != nil && request.Updater != nil {
		res.Updater = null.StringFrom(request.GetUpdater()) // 更新人
	}
	if request != nil && request.UpdateTime != nil {
		res.UpdateTime = null.DateTimeFrom(util.GrpcTime(request.UpdateTime)) // 更新时间
	}

	return res
}

func (srv SrvRegionServiceServer) RegionResult(item dao.Region) *RegionObject {
	res := &RegionObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.ParentId != nil {
		res.ParentId = item.ParentId
	}
	if item.WeatherCode.IsValid() {
		res.WeatherCode = item.WeatherCode.Ptr()
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

// RegionCreate 创建数据
func (srv SrvRegionServiceServer) RegionCreate(ctx context.Context, request *RegionCreateRequest) (*RegionCreateResponse, error) {
	req := srv.RegionConvert(request.GetData())
	data, err := region.RegionCreate(ctx, req)
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
	regionId := request.GetRegionId()
	if regionId < 1 {
		return &RegionUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := srv.RegionConvert(request.GetData())
	_, err := region.RegionUpdate(ctx, regionId, req)
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
	regionId := request.GetRegionId()
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
	regionId := request.GetRegionId()
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
		Data: srv.RegionResult(res),
	}, nil
}
func (srv SrvRegionServiceServer) RegionList(ctx context.Context, request *RegionListRequest) (*RegionListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
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
		res = append(res, srv.RegionResult(item))
	}
	return &RegionListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
