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
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_user 系统用户

// SrvSystemUserServiceServer 系统用户
type SrvSystemUserServiceServer struct {
	UnimplementedSystemUserServiceServer
	Server *xgrpc.Server
}

func (srv SrvSystemUserServiceServer) SystemUserConvert(request *SystemUserObject) dao.SystemUser {
	var res dao.SystemUser

	if request != nil && request.Id != nil {
		res.Id = request.Id // 用户编号
	}
	if request != nil && request.Nickname != nil {
		res.Nickname = null.StringFrom(request.GetNickname()) // 昵称
	}
	if request != nil && request.Username != nil {
		res.Username = request.Username // 用户名称
	}
	if request != nil && request.Password != nil {
		res.Password = request.Password // 用户密码
	}
	if request != nil && request.Status != nil {
		res.Status = request.Status // 用户状态（0正常 1停用）
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

func (srv SrvSystemUserServiceServer) SystemUserResult(item dao.SystemUser) *SystemUserObject {
	res := &SystemUserObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Nickname.IsValid() {
		res.Nickname = item.Nickname.Ptr()
	}
	if item.Username != nil {
		res.Username = item.Username
	}
	if item.Password != nil {
		res.Password = item.Password
	}
	if item.Status != nil {
		res.Status = item.Status
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

// SystemUserCreate 创建数据
func (srv SrvSystemUserServiceServer) SystemUserCreate(ctx context.Context, request *SystemUserCreateRequest) (*SystemUserCreateResponse, error) {
	req := srv.SystemUserConvert(request.GetData())
	data, err := user.SystemUserCreate(ctx, req)
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
	systemUserId := request.GetSystemUserId()
	if systemUserId < 1 {
		return &SystemUserUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := srv.SystemUserConvert(request.GetData())
	_, err := user.SystemUserUpdate(ctx, systemUserId, req)
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
	systemUserId := request.GetSystemUserId()
	if systemUserId < 1 {
		return &SystemUserDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := user.SystemUserDelete(ctx, systemUserId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemUserId,
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
	systemUserId := request.GetSystemUserId()
	if systemUserId < 1 {
		return &SystemUserResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := user.SystemUser(ctx, systemUserId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemUserId,
			"err": err,
		}).Error("Sql:系统用户:system_user:SystemUser")
		return &SystemUserResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemUserResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: srv.SystemUserResult(res),
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
		Data: srv.SystemUserResult(res),
	}, nil
}
func (srv SrvSystemUserServiceServer) SystemUserList(ctx context.Context, request *SystemUserListRequest) (*SystemUserListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}

	// 当前页面
	pageNum := request.GetPageNum()
	// 每页多少数据
	pageSize := request.GetPageSize()
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	// 分页数据
	offset := pageSize * (pageNum - 1)
	condition["offset"] = offset
	condition["limit"] = pageSize
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
		res = append(res, srv.SystemUserResult(item))
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
	if request.Status != nil {
		condition["status"] = request.GetStatus()
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

func (srv SrvSystemUserServiceServer) SystemMenuResult(item dao.SystemMenu) *SystemMenuObject {
	res := &SystemMenuObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.Permission.IsValid() {
		res.Permission = item.Permission.Ptr()
	}
	if item.Type != nil {
		res.Type = item.Type
	}
	if item.Sort != nil {
		res.Sort = item.Sort
	}
	if item.ParentId != nil {
		res.ParentId = item.ParentId
	}
	if item.Path.IsValid() {
		res.Path = item.Path.Ptr()
	}
	if item.Icon.IsValid() {
		res.Icon = item.Icon.Ptr()
	}
	if item.Component.IsValid() {
		res.Component = item.Component.Ptr()
	}
	if item.ComponentName.IsValid() {
		res.ComponentName = item.ComponentName.Ptr()
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Hide != nil {
		res.Hide = item.Hide
	}
	if item.Link.IsValid() {
		res.Link = item.Link.Ptr()
	}
	if item.KeepAlive != nil {
		res.KeepAlive = item.KeepAlive
	}
	if item.Affix != nil {
		res.Affix = item.Affix
	}
	if item.ActivePath.IsValid() {
		res.ActivePath = item.ActivePath.Ptr()
	}
	if item.FullScreen != nil {
		res.FullScreen = item.FullScreen
	}
	if item.Redirect.IsValid() {
		res.Redirect = item.Redirect.Ptr()
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
	if item.Deleted != nil {
		res.Deleted = item.Deleted
	}

	return res
}

// SystemUserMenuList 获取用户菜单
func (srv *SrvSystemUserServiceServer) SystemUserMenuList(ctx context.Context, request *SystemUserMenuListRequest) (*SystemUserMenuListResponse, error) {
	systemUserId := request.GetSystemUserId()
	if systemUserId < 1 {
		return &SystemUserMenuListResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	list, err := user.SystemUserMenuList(ctx, systemUserId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemUserId,
			"err": err,
		}).Error("Sql:系统用户:system_user:SystemUserMenuList")
		return &SystemUserMenuListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}

	var res []*SystemMenuObject
	for _, item := range list {
		res = append(res, srv.SystemMenuResult(item))
	}
	return &SystemUserMenuListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
