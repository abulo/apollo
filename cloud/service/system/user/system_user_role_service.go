package user

import (
	"cloud/code"
	"cloud/module/system/user"
	"context"
	"encoding/json"
	"fmt"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_user_role 系统用户和系统角色关联表

// SrvSystemUserRoleServiceServer 系统用户和系统角色关联表
type SrvSystemUserRoleServiceServer struct {
	UnimplementedSystemUserRoleServiceServer
	Server *xgrpc.Server
}

// SystemUserRoleCreate 创建数据
func (srv SrvSystemUserRoleServiceServer) SystemUserRoleCreate(ctx context.Context, request *SystemUserRoleCreateRequest) (*SystemUserRoleCreateResponse, error) {
	req := SystemUserRoleCustomDao(request.GetData())
	js, _ := json.Marshal(req)
	fmt.Println(string(js))
	data, err := user.SystemUserRoleCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:系统用户和系统角色关联表:system_user_role:SystemUserRoleCreate")
		return &SystemUserRoleCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserRoleCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}
func (srv SrvSystemUserRoleServiceServer) SystemUserRoleList(ctx context.Context, request *SystemUserRoleListRequest) (*SystemUserRoleListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.SystemUserId != nil {
		condition["systemUserId"] = request.GetSystemUserId()
	}
	if request.SystemRoleId != nil {
		condition["systemRoleId"] = request.GetSystemRoleId()
	}
	if request.SystemTenantId != nil {
		condition["systemTenantId"] = request.GetSystemTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}

	// 获取数据集合
	list, err := user.SystemUserRoleList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:系统用户和系统角色关联表:system_user_role:SystemUserRoleList")
		return &SystemUserRoleListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemUserRoleObject
	for _, item := range list {
		res = append(res, SystemUserRoleProto(item))
	}
	return &SystemUserRoleListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
