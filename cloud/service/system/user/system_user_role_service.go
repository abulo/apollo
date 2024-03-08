package user

import (
	"cloud/code"
	"cloud/dao"
	"cloud/module/system/user"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_user_role 系统用户和系统角色关联表

// SrvSystemUserRoleServiceServer 系统用户和系统角色关联表
type SrvSystemUserRoleServiceServer struct {
	UnimplementedSystemUserRoleServiceServer
	Server *xgrpc.Server
}

func (srv SrvSystemUserRoleServiceServer) SystemUserRoleCustomConvert(request *SystemUserRoleCreateRequest) dao.SystemUserRoleCustom {
	var res dao.SystemUserRoleCustom

	if request != nil && request.SystemUserId != nil {
		res.SystemUserId = request.SystemUserId // 用户编号
	}
	if request != nil && request.SystemRoleIds != nil {
		res.SystemRoleIds = null.JSONFrom(request.GetSystemRoleIds()) // 角色编号
	}
	if request != nil && request.Creator != nil {
		res.Creator = null.StringFrom(request.GetCreator()) // 创建者
	}
	if request != nil && request.CreateTime != nil {
		res.CreateTime = null.DateTimeFrom(util.GrpcTime(request.CreateTime)) // 创建时间
	}
	if request != nil && request.Updater != nil {
		res.Updater = null.StringFrom(request.GetUpdater()) // 更新者
	}
	if request != nil && request.UpdateTime != nil {
		res.UpdateTime = null.DateTimeFrom(util.GrpcTime(request.UpdateTime)) // 更新时间
	}

	return res
}

// SystemUserRoleCreate 创建数据
func (srv SrvSystemUserRoleServiceServer) SystemUserRoleCreate(ctx context.Context, request *SystemUserRoleCreateRequest) (*SystemUserRoleCreateResponse, error) {
	req := srv.SystemUserRoleCustomConvert(request)
	data, err := user.SystemUserRoleCreate(ctx, req)
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
	if request.SystemRoleId != nil {
		condition["systemRoleId"] = request.GetSystemRoleId()
	}
	if request.SystemUserId != nil {
		condition["systemUserId"] = request.GetSystemUserId()
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
