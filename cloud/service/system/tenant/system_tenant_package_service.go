package tenant

import (
	"cloud/code"
	"cloud/module/system/tenant"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_tenant_package 租户套餐包

// SrvSystemTenantPackageServiceServer 租户套餐包
type SrvSystemTenantPackageServiceServer struct {
	UnimplementedSystemTenantPackageServiceServer
	Server *xgrpc.Server
}

// SystemTenantPackageCreate 创建数据
func (srv SrvSystemTenantPackageServiceServer) SystemTenantPackageCreate(ctx context.Context, request *SystemTenantPackageCreateRequest) (*SystemTenantPackageCreateResponse, error) {
	req := SystemTenantPackageDao(request.GetData())
	data, err := tenant.SystemTenantPackageCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户套餐包:system_tenant_package:SystemTenantPackageCreate")
		return &SystemTenantPackageCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantPackageCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemTenantPackageUpdate 更新数据
func (srv SrvSystemTenantPackageServiceServer) SystemTenantPackageUpdate(ctx context.Context, request *SystemTenantPackageUpdateRequest) (*SystemTenantPackageUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemTenantPackageUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemTenantPackageDao(request.GetData())
	_, err := tenant.SystemTenantPackageUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户套餐包:system_tenant_package:SystemTenantPackageUpdate")
		return &SystemTenantPackageUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantPackageUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemTenantPackageDelete 删除数据
func (srv SrvSystemTenantPackageServiceServer) SystemTenantPackageDelete(ctx context.Context, request *SystemTenantPackageDeleteRequest) (*SystemTenantPackageDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemTenantPackageDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.SystemTenantPackageDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户套餐包:system_tenant_package:SystemTenantPackageDelete")
		return &SystemTenantPackageDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantPackageDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemTenantPackage 查询单条数据
func (srv SrvSystemTenantPackageServiceServer) SystemTenantPackage(ctx context.Context, request *SystemTenantPackageRequest) (*SystemTenantPackageResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemTenantPackageResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := tenant.SystemTenantPackage(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户套餐包:system_tenant_package:SystemTenantPackage")
		return &SystemTenantPackageResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantPackageResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemTenantPackageProto(res),
	}, nil
}

// SystemTenantPackageRecover 恢复数据
func (srv SrvSystemTenantPackageServiceServer) SystemTenantPackageRecover(ctx context.Context, request *SystemTenantPackageRecoverRequest) (*SystemTenantPackageRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemTenantPackageRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.SystemTenantPackageRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户套餐包:system_tenant_package:SystemTenantPackageRecover")
		return &SystemTenantPackageRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantPackageRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemTenantPackageDrop 清理数据
func (srv SrvSystemTenantPackageServiceServer) SystemTenantPackageDrop(ctx context.Context, request *SystemTenantPackageDropRequest) (*SystemTenantPackageDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemTenantPackageDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.SystemTenantPackageDrop(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户套餐包:system_tenant_package:SystemTenantPackageDrop")
		return &SystemTenantPackageDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantPackageDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemTenantPackageList 列表数据
func (srv SrvSystemTenantPackageServiceServer) SystemTenantPackageList(ctx context.Context, request *SystemTenantPackageListRequest) (*SystemTenantPackageListResponse, error) {
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

	paginationRequest := request.GetPagination()
	if paginationRequest != nil {
		// 当前页面
		pageNum := paginationRequest.GetPageNum()
		// 每页多少数据
		pageSize := paginationRequest.GetPageSize()
		if pageNum < 1 {
			pageNum = 1
		}
		if pageSize < 1 {
			pageSize = 10
		}
		// 分页数据
		offset := pageSize * (pageNum - 1)
		pagination := &sql.Pagination{
			Offset: &offset,
			Limit:  &pageSize,
		}
		condition["pagination"] = pagination
	}
	// 获取数据集合
	list, err := tenant.SystemTenantPackageList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户套餐包:system_tenant_package:SystemTenantPackageList")
		return &SystemTenantPackageListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemTenantPackageObject
	for _, item := range list {
		res = append(res, SystemTenantPackageProto(item))
	}
	return &SystemTenantPackageListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemTenantPackageListTotal 获取总数
func (srv SrvSystemTenantPackageServiceServer) SystemTenantPackageListTotal(ctx context.Context, request *SystemTenantPackageListTotalRequest) (*SystemTenantPackageTotalResponse, error) {
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

	// 获取数据集合
	total, err := tenant.SystemTenantPackageListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户套餐包:system_tenant_package:SystemTenantPackageListTotal")
		return &SystemTenantPackageTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantPackageTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
