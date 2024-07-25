package role

import (
	"cloud/code"
	"cloud/module/system/role"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_role 系统角色

// SrvSystemRoleServiceServer 系统角色
type SrvSystemRoleServiceServer struct {
	UnimplementedSystemRoleServiceServer
	Server *xgrpc.Server
}

// SystemRoleCreate 创建数据
func (srv SrvSystemRoleServiceServer) SystemRoleCreate(ctx context.Context, request *SystemRoleCreateRequest) (*SystemRoleCreateResponse, error) {
	req := SystemRoleDao(request.GetData())
	data, err := role.SystemRoleCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:系统角色:system_role:SystemRoleCreate")
		return &SystemRoleCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemRoleCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemRoleUpdate 更新数据
func (srv SrvSystemRoleServiceServer) SystemRoleUpdate(ctx context.Context, request *SystemRoleUpdateRequest) (*SystemRoleUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemRoleUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemRoleDao(request.GetData())
	_, err := role.SystemRoleUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:系统角色:system_role:SystemRoleUpdate")
		return &SystemRoleUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemRoleUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemRoleDelete 删除数据
func (srv SrvSystemRoleServiceServer) SystemRoleDelete(ctx context.Context, request *SystemRoleDeleteRequest) (*SystemRoleDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemRoleDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := role.SystemRoleDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:系统角色:system_role:SystemRoleDelete")
		return &SystemRoleDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemRoleDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemRole 查询单条数据
func (srv SrvSystemRoleServiceServer) SystemRole(ctx context.Context, request *SystemRoleRequest) (*SystemRoleResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemRoleResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := role.SystemRole(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:系统角色:system_role:SystemRole")
		return &SystemRoleResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemRoleResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemRoleProto(res),
	}, nil
}

// SystemRoleRecover 恢复数据
func (srv SrvSystemRoleServiceServer) SystemRoleRecover(ctx context.Context, request *SystemRoleRecoverRequest) (*SystemRoleRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemRoleRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := role.SystemRoleRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:系统角色:system_role:SystemRoleRecover")
		return &SystemRoleRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemRoleRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemRoleDrop 清理数据
func (srv SrvSystemRoleServiceServer) SystemRoleDrop(ctx context.Context, request *SystemRoleDropRequest) (*SystemRoleDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemRoleDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := role.SystemRoleDrop(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:系统角色:system_role:SystemRoleDrop")
		return &SystemRoleDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemRoleDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemRoleList 列表数据
func (srv SrvSystemRoleServiceServer) SystemRoleList(ctx context.Context, request *SystemRoleListRequest) (*SystemRoleListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Type != nil {
		condition["type"] = request.GetType()
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
	list, err := role.SystemRoleList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:系统角色:system_role:SystemRoleList")
		return &SystemRoleListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemRoleObject
	for _, item := range list {
		res = append(res, SystemRoleProto(item))
	}
	return &SystemRoleListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemRoleListTotal 获取总数
func (srv SrvSystemRoleServiceServer) SystemRoleListTotal(ctx context.Context, request *SystemRoleListTotalRequest) (*SystemRoleTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}

	// 获取数据集合
	total, err := role.SystemRoleListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:系统角色:system_role:SystemRoleListTotal")
		return &SystemRoleTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemRoleTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
