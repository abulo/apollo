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

// system_user_post 系统用户职位

// SrvSystemUserPostServiceServer 系统用户职位
type SrvSystemUserPostServiceServer struct {
	UnimplementedSystemUserPostServiceServer
	Server *xgrpc.Server
}

// SystemUserPostCreate 创建数据
func (srv SrvSystemUserPostServiceServer) SystemUserPostCreate(ctx context.Context, request *SystemUserPostCreateRequest) (*SystemUserPostCreateResponse, error) {
	req := SystemUserPostCustomDao(request.GetData())
	data, err := user.SystemUserPostCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:系统用户职位:system_user_post:SystemUserPostCreate")
		return &SystemUserPostCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserPostCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}
func (srv SrvSystemUserPostServiceServer) SystemUserPostList(ctx context.Context, request *SystemUserPostListRequest) (*SystemUserPostListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.PostId != nil {
		condition["postId"] = request.GetPostId()
	}
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}

	// 获取数据集合
	list, err := user.SystemUserPostList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:系统用户职位:system_user_post:SystemUserPostList")
		return &SystemUserPostListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemUserPostObject
	for _, item := range list {
		res = append(res, SystemUserPostProto(item))
	}
	return &SystemUserPostListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
