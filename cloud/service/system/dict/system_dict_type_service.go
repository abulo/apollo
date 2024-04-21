package dict

import (
	"cloud/code"
	"cloud/module/system/dict"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_dict_type 字典类型

// SrvSystemDictTypeServiceServer 字典类型
type SrvSystemDictTypeServiceServer struct {
	UnimplementedSystemDictTypeServiceServer
	Server *xgrpc.Server
}

// SystemDictTypeCreate 创建数据
func (srv SrvSystemDictTypeServiceServer) SystemDictTypeCreate(ctx context.Context, request *SystemDictTypeCreateRequest) (*SystemDictTypeCreateResponse, error) {
	req := SystemDictTypeDao(request.GetData())
	data, err := dict.SystemDictTypeCreate(ctx, *req)
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
	id := request.GetId()
	if id < 1 {
		return &SystemDictTypeUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemDictTypeDao(request.GetData())
	_, err := dict.SystemDictTypeUpdate(ctx, id, *req)
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
	id := request.GetId()
	if id < 1 {
		return &SystemDictTypeDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dict.SystemDictTypeDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
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
	id := request.GetId()
	if id < 1 {
		return &SystemDictTypeResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := dict.SystemDictType(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:字典类型:system_dict_type:SystemDictType")
		return &SystemDictTypeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDictTypeResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemDictTypeProto(res),
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
		res = append(res, SystemDictTypeProto(item))
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
