package dict

import (
	"cloud/code"
	"cloud/dao"
	"cloud/module/system/dict"
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

// system_dict_type 字典类型

// SrvSystemDictTypeServiceServer 字典类型
type SrvSystemDictTypeServiceServer struct {
	UnimplementedSystemDictTypeServiceServer
	Server *xgrpc.Server
}

func (srv SrvSystemDictTypeServiceServer) SystemDictTypeConvert(request *SystemDictTypeObject) dao.SystemDictType {
	var res dao.SystemDictType

	if request != nil && request.Id != nil {
		res.Id = request.Id // 字典主键
	}
	if request != nil && request.Name != nil {
		res.Name = request.Name // 字典名称
	}
	if request != nil && request.Type != nil {
		res.Type = request.Type // 字典类型
	}
	if request != nil && request.Status != nil {
		res.Status = request.Status // 状态（0正常 1停用）
	}
	if request != nil && request.Remark != nil {
		res.Remark = null.StringFrom(request.GetRemark()) // 备注
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

func (srv SrvSystemDictTypeServiceServer) SystemDictTypeResult(item dao.SystemDictType) *SystemDictTypeObject {
	res := &SystemDictTypeObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.Type != nil {
		res.Type = item.Type
	}
	if item.Status != nil {
		res.Status = item.Status
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

// SystemDictTypeCreate 创建数据
func (srv SrvSystemDictTypeServiceServer) SystemDictTypeCreate(ctx context.Context, request *SystemDictTypeCreateRequest) (*SystemDictTypeCreateResponse, error) {
	req := srv.SystemDictTypeConvert(request.GetData())
	data, err := dict.SystemDictTypeCreate(ctx, req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:字典类型:system_dict_type:SystemDictTypeCreate")
		return &SystemDictTypeCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDictTypeCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemDictTypeUpdate 更新数据
func (srv SrvSystemDictTypeServiceServer) SystemDictTypeUpdate(ctx context.Context, request *SystemDictTypeUpdateRequest) (*SystemDictTypeUpdateResponse, error) {
	systemDictTypeId := request.GetSystemDictTypeId()
	if systemDictTypeId < 1 {
		return &SystemDictTypeUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := srv.SystemDictTypeConvert(request.GetData())
	_, err := dict.SystemDictTypeUpdate(ctx, systemDictTypeId, req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:字典类型:system_dict_type:SystemDictTypeUpdate")
		return &SystemDictTypeUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDictTypeUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemDictTypeDelete 删除数据
func (srv SrvSystemDictTypeServiceServer) SystemDictTypeDelete(ctx context.Context, request *SystemDictTypeDeleteRequest) (*SystemDictTypeDeleteResponse, error) {
	systemDictTypeId := request.GetSystemDictTypeId()
	if systemDictTypeId < 1 {
		return &SystemDictTypeDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dict.SystemDictTypeDelete(ctx, systemDictTypeId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemDictTypeId,
			"err": err,
		}).Error("Sql:字典类型:system_dict_type:SystemDictTypeDelete")
		return &SystemDictTypeDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDictTypeDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemDictType 查询单条数据
func (srv SrvSystemDictTypeServiceServer) SystemDictType(ctx context.Context, request *SystemDictTypeRequest) (*SystemDictTypeResponse, error) {
	systemDictTypeId := request.GetSystemDictTypeId()
	if systemDictTypeId < 1 {
		return &SystemDictTypeResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := dict.SystemDictType(ctx, systemDictTypeId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemDictTypeId,
			"err": err,
		}).Error("Sql:字典类型:system_dict_type:SystemDictType")
		return &SystemDictTypeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDictTypeResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: srv.SystemDictTypeResult(res),
	}, nil
}
func (srv SrvSystemDictTypeServiceServer) SystemDictTypeList(ctx context.Context, request *SystemDictTypeListRequest) (*SystemDictTypeListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}

	// 当前页面
	pageNum := request.GetPageNum()
	// 每页多少数据
	pageSize := request.GetPageSize()
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	// 分页数据
	offset := pageSize * (pageNum - 1)
	condition["offset"] = offset
	condition["limit"] = pageSize
	// 获取数据集合
	list, err := dict.SystemDictTypeList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典类型:system_dict_type:SystemDictTypeList")
		return &SystemDictTypeListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemDictTypeObject
	for _, item := range list {
		res = append(res, srv.SystemDictTypeResult(item))
	}
	return &SystemDictTypeListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemDictTypeListTotal 获取总数
func (srv SrvSystemDictTypeServiceServer) SystemDictTypeListTotal(ctx context.Context, request *SystemDictTypeListTotalRequest) (*SystemDictTypeTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}

	// 获取数据集合
	total, err := dict.SystemDictTypeListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典类型:system_dict_type:SystemDictTypeListTotal")
		return &SystemDictTypeTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDictTypeTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
