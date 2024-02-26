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

// system_dict 字典数据表

// SrvSystemDictServiceServer 字典数据表
type SrvSystemDictServiceServer struct {
	UnimplementedSystemDictServiceServer
	Server *xgrpc.Server
}

func (srv SrvSystemDictServiceServer) SystemDictConvert(request *SystemDictObject) dao.SystemDict {
	var res dao.SystemDict

	if request != nil && request.Id != nil {
		res.Id = request.Id // 字典编码
	}
	if request != nil && request.Sort != nil {
		res.Sort = request.Sort // 字典排序
	}
	if request != nil && request.Label != nil {
		res.Label = request.Label // 字典标签
	}
	if request != nil && request.Value != nil {
		res.Value = request.Value // 字典键值
	}
	if request != nil && request.DictType != nil {
		res.DictType = request.DictType // 字典类型
	}
	if request != nil && request.Status != nil {
		res.Status = request.Status // 状态（0正常 1停用）
	}
	if request != nil && request.ColorType != nil {
		res.ColorType = null.StringFrom(request.GetColorType()) // 颜色类型
	}
	if request != nil && request.CssClass != nil {
		res.CssClass = null.StringFrom(request.GetCssClass()) // css 样式
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

func (srv SrvSystemDictServiceServer) SystemDictResult(item dao.SystemDict) *SystemDictObject {
	res := &SystemDictObject{}
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
func (srv SrvSystemDictServiceServer) SystemDictCreate(ctx context.Context, request *SystemDictCreateRequest) (*SystemDictCreateResponse, error) {
	req := srv.SystemDictConvert(request.GetData())
	data, err := dict.SystemDictCreate(ctx, req)
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
	systemDictId := request.GetSystemDictId()
	if systemDictId < 1 {
		return &SystemDictUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := srv.SystemDictConvert(request.GetData())
	_, err := dict.SystemDictUpdate(ctx, systemDictId, req)
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
	systemDictId := request.GetSystemDictId()
	if systemDictId < 1 {
		return &SystemDictDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dict.SystemDictDelete(ctx, systemDictId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemDictId,
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
	systemDictId := request.GetSystemDictId()
	if systemDictId < 1 {
		return &SystemDictResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := dict.SystemDict(ctx, systemDictId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemDictId,
			"err": err,
		}).Error("Sql:字典数据表:system_dict:SystemDict")
		return &SystemDictResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDictResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: srv.SystemDictResult(res),
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
		res = append(res, srv.SystemDictResult(item))
	}
	return &SystemDictListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
