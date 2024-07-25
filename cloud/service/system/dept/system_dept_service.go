package dept

import (
	"cloud/code"
	"cloud/module/system/dept"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_dept 部门

// SrvSystemDeptServiceServer 部门
type SrvSystemDeptServiceServer struct {
	UnimplementedSystemDeptServiceServer
	Server *xgrpc.Server
}

// SystemDeptCreate 创建数据
func (srv SrvSystemDeptServiceServer) SystemDeptCreate(ctx context.Context, request *SystemDeptCreateRequest) (*SystemDeptCreateResponse, error) {
	req := SystemDeptDao(request.GetData())
	data, err := dept.SystemDeptCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:部门:system_dept:SystemDeptCreate")
		return &SystemDeptCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDeptCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemDeptUpdate 更新数据
func (srv SrvSystemDeptServiceServer) SystemDeptUpdate(ctx context.Context, request *SystemDeptUpdateRequest) (*SystemDeptUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemDeptUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemDeptDao(request.GetData())
	_, err := dept.SystemDeptUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:部门:system_dept:SystemDeptUpdate")
		return &SystemDeptUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDeptUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemDeptDelete 删除数据
func (srv SrvSystemDeptServiceServer) SystemDeptDelete(ctx context.Context, request *SystemDeptDeleteRequest) (*SystemDeptDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemDeptDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dept.SystemDeptDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:部门:system_dept:SystemDeptDelete")
		return &SystemDeptDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDeptDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemDept 查询单条数据
func (srv SrvSystemDeptServiceServer) SystemDept(ctx context.Context, request *SystemDeptRequest) (*SystemDeptResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemDeptResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := dept.SystemDept(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:部门:system_dept:SystemDept")
		return &SystemDeptResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDeptResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemDeptProto(res),
	}, nil
}

// SystemDeptRecover 恢复数据
func (srv SrvSystemDeptServiceServer) SystemDeptRecover(ctx context.Context, request *SystemDeptRecoverRequest) (*SystemDeptRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemDeptRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dept.SystemDeptRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:部门:system_dept:SystemDeptRecover")
		return &SystemDeptRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDeptRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemDeptDrop 清理数据
func (srv SrvSystemDeptServiceServer) SystemDeptDrop(ctx context.Context, request *SystemDeptDropRequest) (*SystemDeptDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemDeptDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dept.SystemDeptDrop(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:部门:system_dept:SystemDeptDrop")
		return &SystemDeptDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDeptDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemDeptServiceServer) SystemDeptList(ctx context.Context, request *SystemDeptListRequest) (*SystemDeptListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.ParentId != nil {
		condition["parentId"] = request.GetParentId()
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
	list, err := dept.SystemDeptList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:部门:system_dept:SystemDeptList")
		return &SystemDeptListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemDeptObject
	for _, item := range list {
		res = append(res, SystemDeptProto(item))
	}
	return &SystemDeptListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemDeptListTotal 获取总数
func (srv SrvSystemDeptServiceServer) SystemDeptListTotal(ctx context.Context, request *SystemDeptListTotalRequest) (*SystemDeptTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.ParentId != nil {
		condition["parentId"] = request.GetParentId()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}

	// 获取数据集合
	total, err := dept.SystemDeptListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:部门:system_dept:SystemDeptListTotal")
		return &SystemDeptTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemDeptTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
