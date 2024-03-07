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
	systemRoleId := request.GetSystemRoleId()
	if systemRoleId < 1 {
		return &SystemRoleUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemRoleDao(request.GetData())
	_, err := role.SystemRoleUpdate(ctx, systemRoleId, *req)
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
	systemRoleId := request.GetSystemRoleId()
	if systemRoleId < 1 {
		return &SystemRoleDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := role.SystemRoleDelete(ctx, systemRoleId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemRoleId,
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
	systemRoleId := request.GetSystemRoleId()
	if systemRoleId < 1 {
		return &SystemRoleResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := role.SystemRole(ctx, systemRoleId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemRoleId,
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
func (srv SrvSystemRoleServiceServer) SystemRoleList(ctx context.Context, request *SystemRoleListRequest) (*SystemRoleListResponse, error) {
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
