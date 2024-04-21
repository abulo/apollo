package user

import (
	"cloud/code"
	"cloud/module/system/user"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_user_dept 系统用户部门

// SrvSystemUserDeptServiceServer 系统用户部门
type SrvSystemUserDeptServiceServer struct {
	UnimplementedSystemUserDeptServiceServer
	Server *xgrpc.Server
}

// SystemUserDeptCreate 创建数据
func (srv SrvSystemUserDeptServiceServer) SystemUserDeptCreate(ctx context.Context, request *SystemUserDeptCreateRequest) (*SystemUserDeptCreateResponse, error) {
	req := SystemUserDeptCustomDao(request.GetData())
	data, err := user.SystemUserDeptCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:系统用户部门:system_user_dept:SystemUserDeptCreate")
		return &SystemUserDeptCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserDeptCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}
func (srv SrvSystemUserDeptServiceServer) SystemUserDeptList(ctx context.Context, request *SystemUserDeptListRequest) (*SystemUserDeptListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.DeptId != nil {
		condition["deptId"] = request.GetDeptId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}

	// 获取数据集合
	list, err := user.SystemUserDeptList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:系统用户部门:system_user_dept:SystemUserDeptList")
		return &SystemUserDeptListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemUserDeptObject
	for _, item := range list {
		res = append(res, SystemUserDeptProto(item))
	}
	return &SystemUserDeptListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
