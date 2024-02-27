package role

import (
	"cloud/code"
	"cloud/dao"
	"cloud/module/system/role"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_role 系统角色

// SrvSystemRoleServiceServer 系统角色
type SrvSystemRoleServiceServer struct {
	UnimplementedSystemRoleServiceServer
	Server *xgrpc.Server
}

func (srv SrvSystemRoleServiceServer) SystemRoleConvert(request *SystemRoleObject) dao.SystemRole {
	var res dao.SystemRole

	if request != nil && request.Id != nil {
		res.Id = request.Id // 角色编号
	}
	if request != nil && request.Name != nil {
		res.Name = request.Name // 角色名称
	}
	if request != nil && request.Code != nil {
		res.Code = request.Code // 角色权限字符串
	}
	if request != nil && request.Sort != nil {
		res.Sort = request.Sort // 显示顺序
	}
	if request != nil && request.Status != nil {
		res.Status = request.Status // 角色状态（0正常 1停用）
	}
	if request != nil && request.Type != nil {
		res.Type = request.Type // 角色类型(1内置/2定义)
	}
	if request != nil && request.Remark != nil {
		res.Remark = null.StringFrom(request.GetRemark()) // 备注
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

func (srv SrvSystemRoleServiceServer) SystemRoleResult(item dao.SystemRole) *SystemRoleObject {
	res := &SystemRoleObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.Code != nil {
		res.Code = item.Code
	}
	if item.Sort != nil {
		res.Sort = item.Sort
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Type != nil {
		res.Type = item.Type
	}
	if item.Remark.IsValid() {
		res.Remark = item.Remark.Ptr()
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

// SystemRoleCreate 创建数据
func (srv SrvSystemRoleServiceServer) SystemRoleCreate(ctx context.Context, request *SystemRoleCreateRequest) (*SystemRoleCreateResponse, error) {
	req := srv.SystemRoleConvert(request.GetData())
	data, err := role.SystemRoleCreate(ctx, req)
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
	req := srv.SystemRoleConvert(request.GetData())
	_, err := role.SystemRoleUpdate(ctx, systemRoleId, req)
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
		Data: srv.SystemRoleResult(res),
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
		res = append(res, srv.SystemRoleResult(item))
	}
	return &SystemRoleListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
