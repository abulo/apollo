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

// system_dict 字典数据表

// SrvSystemDictServiceServer 字典数据表
type SrvSystemDictServiceServer struct {
	UnimplementedSystemDictServiceServer
	Server *xgrpc.Server
}

// SystemDictCreate 创建数据
func (srv SrvSystemDictServiceServer) SystemDictCreate(ctx context.Context, request *SystemDictCreateRequest) (*SystemDictCreateResponse, error) {
	req := SystemDictDao(request.GetData())
	data, err := dict.SystemDictCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:字典数据表:system_dict:SystemDictCreate")
		return &SystemDictCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDictCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemDictUpdate 更新数据
func (srv SrvSystemDictServiceServer) SystemDictUpdate(ctx context.Context, request *SystemDictUpdateRequest) (*SystemDictUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemDictUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemDictDao(request.GetData())
	_, err := dict.SystemDictUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:字典数据表:system_dict:SystemDictUpdate")
		return &SystemDictUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDictUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemDictDelete 删除数据
func (srv SrvSystemDictServiceServer) SystemDictDelete(ctx context.Context, request *SystemDictDeleteRequest) (*SystemDictDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemDictDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dict.SystemDictDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:字典数据表:system_dict:SystemDictDelete")
		return &SystemDictDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDictDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemDict 查询单条数据
func (srv SrvSystemDictServiceServer) SystemDict(ctx context.Context, request *SystemDictRequest) (*SystemDictResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemDictResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := dict.SystemDict(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:字典数据表:system_dict:SystemDict")
		return &SystemDictResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDictResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemDictProto(res),
	}, nil
}
func (srv SrvSystemDictServiceServer) SystemDictList(ctx context.Context, request *SystemDictListRequest) (*SystemDictListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.DictType != nil {
		condition["dictType"] = request.GetDictType()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}

	// 获取数据集合
	list, err := dict.SystemDictList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典数据表:system_dict:SystemDictList")
		return &SystemDictListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemDictObject
	for _, item := range list {
		res = append(res, SystemDictProto(item))
	}
	return &SystemDictListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
