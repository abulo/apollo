package user

import (
	"cloud/code"
	"cloud/module/system/user"
	"context"
	"encoding/json"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_user 系统用户

// SrvSystemUserServiceServer 系统用户
type SrvSystemUserServiceServer struct {
	UnimplementedSystemUserServiceServer
	Server *xgrpc.Server
}

// SystemUserCreate 创建数据
func (srv SrvSystemUserServiceServer) SystemUserCreate(ctx context.Context, request *SystemUserCreateRequest) (*SystemUserCreateResponse, error) {
	req := SystemUserDao(request.GetData())
	data, err := user.SystemUserCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:系统用户:system_user:SystemUserCreate")
		return &SystemUserCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemUserUpdate 更新数据
func (srv SrvSystemUserServiceServer) SystemUserUpdate(ctx context.Context, request *SystemUserUpdateRequest) (*SystemUserUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemUserUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemUserDao(request.GetData())
	_, err := user.SystemUserUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:系统用户:system_user:SystemUserUpdate")
		return &SystemUserUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemUserDelete 删除数据
func (srv SrvSystemUserServiceServer) SystemUserDelete(ctx context.Context, request *SystemUserDeleteRequest) (*SystemUserDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemUserDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := user.SystemUserDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:系统用户:system_user:SystemUserDelete")
		return &SystemUserDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemUser 查询单条数据
func (srv SrvSystemUserServiceServer) SystemUser(ctx context.Context, request *SystemUserRequest) (*SystemUserResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemUserResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := user.SystemUser(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:系统用户:system_user:SystemUser")
		return &SystemUserResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemUserProto(res),
	}, nil
}

// SystemUserRecover 恢复数据
func (srv SrvSystemUserServiceServer) SystemUserRecover(ctx context.Context, request *SystemUserRecoverRequest) (*SystemUserRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemUserRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := user.SystemUserRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:系统用户:system_user:SystemUserRecover")
		return &SystemUserRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemUserDrop 清理数据
func (srv SrvSystemUserServiceServer) SystemUserDrop(ctx context.Context, request *SystemUserDropRequest) (*SystemUserDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemUserDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := user.SystemUserDrop(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户:system_user:SystemUserDrop")
		return &SystemUserDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemUserLogin 查询单条数据
func (srv SrvSystemUserServiceServer) SystemUserLogin(ctx context.Context, request *SystemUserLoginRequest) (*SystemUserLoginResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}

	if util.Empty(condition) {
		err := errors.New("condition is empty")
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:系统用户:system_user:SystemUserLogin")
		return &SystemUserLoginResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := user.SystemUserLogin(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:系统用户:system_user:SystemUserLogin")
		return &SystemUserLoginResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserLoginResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemUserProto(res),
	}, nil
}
func (srv SrvSystemUserServiceServer) SystemUserList(ctx context.Context, request *SystemUserListRequest) (*SystemUserListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.DeptId != nil {
		condition["deptId"] = request.GetDeptId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.DataScope != nil {
		condition["dataScope"] = request.GetDataScope()
	}
	if request.DataScopeDept != nil {
		var deptIds []int64
		json.Unmarshal(request.GetDataScopeDept(), &deptIds)
		condition["dataScopeDept"] = deptIds
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
	list, err := user.SystemUserList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:系统用户:system_user:SystemUserList")
		return &SystemUserListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemUserObject
	for _, item := range list {
		res = append(res, SystemUserProto(item))
	}
	return &SystemUserListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemUserListTotal 获取总数
func (srv SrvSystemUserServiceServer) SystemUserListTotal(ctx context.Context, request *SystemUserListTotalRequest) (*SystemUserTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.DeptId != nil {
		condition["deptId"] = request.GetDeptId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.DataScope != nil {
		condition["dataScope"] = request.GetDataScope()
	}
	if request.DataScopeDept != nil {
		var deptIds []int64
		json.Unmarshal(request.GetDataScopeDept(), &deptIds)
		condition["dataScopeDept"] = deptIds
	}

	// 获取数据集合
	total, err := user.SystemUserListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:系统用户:system_user:SystemUserListTotal")
		return &SystemUserTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}

// SystemUserUpdate 更新数据
func (srv SrvSystemUserServiceServer) SystemUserPassword(ctx context.Context, request *SystemUserPasswordRequest) (*SystemUserPasswordResponse, error) {
	id := request.GetId()

	if id < 1 {
		return &SystemUserPasswordResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	password := request.GetPassword()
	if util.Empty(password) {
		return &SystemUserPasswordResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := user.SystemUserPassword(ctx, id, password)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("Sql:用户信息表:system_user:SystemUserPassword")
		return &SystemUserPasswordResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserPasswordResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

func (srv SrvSystemUserServiceServer) SystemUserRoleDataScope(ctx context.Context, request *SystemUserRoleDataScopeRequest) (*SystemUserRoleDataScopeResponse, error) {
	userId := request.GetUserId()
	if userId < 1 {
		return &SystemUserRoleDataScopeResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	tenantId := request.GetTenantId()
	if tenantId < 1 {
		return &SystemUserRoleDataScopeResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := user.SystemUserDataScope(ctx, tenantId, userId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("Sql:系统用户:system_user:SystemUserRoleDataScope")
		return &SystemUserRoleDataScopeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserRoleDataScopeResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemUserRoleDataScopeProto(res),
	}, nil
}
