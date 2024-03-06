package tenant

import (
	"cloud/code"
	"cloud/dao"
	"cloud/module/system/tenant"
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

// system_tenant 租户

// SrvSystemTenantServiceServer 租户
type SrvSystemTenantServiceServer struct {
	UnimplementedSystemTenantServiceServer
	Server *xgrpc.Server
}

func (srv SrvSystemTenantServiceServer) SystemTenantConvert(request *SystemTenantObject) dao.SystemTenant {
	var res dao.SystemTenant

	if request != nil && request.Id != nil {
		res.Id = request.Id // 租户编号
	}
	if request != nil && request.Name != nil {
		res.Name = request.Name // 租户名称
	}
	if request != nil && request.SystemUserId != nil {
		res.SystemUserId = null.Int64From(request.GetSystemUserId()) // 联系人ID
	}
	if request != nil && request.ContactName != nil {
		res.ContactName = request.ContactName // 联系人
	}
	if request != nil && request.ContactMobile != nil {
		res.ContactMobile = request.ContactMobile // 租户联系电话
	}
	if request != nil && request.Status != nil {
		res.Status = request.Status // 状态（0正常 1停用）
	}
	if request != nil && request.Domain != nil {
		res.Domain = null.StringFrom(request.GetDomain()) // 域名
	}
	if request != nil && request.ExpireDate != nil {
		res.ExpireDate = null.DateFrom(util.GrpcTime(request.ExpireDate)) // 过期时间
	}
	if request != nil && request.AccountCont != nil {
		res.AccountCont = request.AccountCont // 账号数量
	}
	if request != nil && request.SystemTenantPackageId != nil {
		res.SystemTenantPackageId = request.SystemTenantPackageId // 套餐编号
	}
	if request != nil && request.Deleted != nil {
		res.Deleted = request.Deleted // 是否删除(0否 1是)
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
	if request != nil && request.Username != nil {
		res.Username = request.Username // 用户名称
	}
	if request != nil && request.Password != nil {
		res.Password = request.Password // 用户密码
	}

	return res
}

func (srv SrvSystemTenantServiceServer) SystemTenantResult(item dao.SystemTenant) *SystemTenantObject {
	res := &SystemTenantObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.SystemUserId.IsValid() {
		res.SystemUserId = item.SystemUserId.Ptr()
	}
	if item.ContactName != nil {
		res.ContactName = item.ContactName
	}
	if item.ContactMobile != nil {
		res.ContactMobile = item.ContactMobile
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Domain.IsValid() {
		res.Domain = item.Domain.Ptr()
	}
	if item.ExpireDate.IsValid() {
		res.ExpireDate = timestamppb.New(*item.ExpireDate.Ptr())
	}
	if item.AccountCont != nil {
		res.AccountCont = item.AccountCont
	}
	if item.SystemTenantPackageId != nil {
		res.SystemTenantPackageId = item.SystemTenantPackageId
	}
	if item.Deleted != nil {
		res.Deleted = item.Deleted
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

// SystemTenantCreate 创建数据
func (srv SrvSystemTenantServiceServer) SystemTenantCreate(ctx context.Context, request *SystemTenantCreateRequest) (*SystemTenantCreateResponse, error) {
	req := srv.SystemTenantConvert(request.GetData())
	data, err := tenant.SystemTenantCreate(ctx, req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenantCreate")
		return &SystemTenantCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemTenantUpdate 更新数据
func (srv SrvSystemTenantServiceServer) SystemTenantUpdate(ctx context.Context, request *SystemTenantUpdateRequest) (*SystemTenantUpdateResponse, error) {
	systemTenantId := request.GetSystemTenantId()
	if systemTenantId < 1 {
		return &SystemTenantUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := srv.SystemTenantConvert(request.GetData())
	_, err := tenant.SystemTenantUpdate(ctx, systemTenantId, req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenantUpdate")
		return &SystemTenantUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemTenantDelete 删除数据
func (srv SrvSystemTenantServiceServer) SystemTenantDelete(ctx context.Context, request *SystemTenantDeleteRequest) (*SystemTenantDeleteResponse, error) {
	systemTenantId := request.GetSystemTenantId()
	if systemTenantId < 1 {
		return &SystemTenantDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.SystemTenantDelete(ctx, systemTenantId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemTenantId,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenantDelete")
		return &SystemTenantDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemTenant 查询单条数据
func (srv SrvSystemTenantServiceServer) SystemTenant(ctx context.Context, request *SystemTenantRequest) (*SystemTenantResponse, error) {
	systemTenantId := request.GetSystemTenantId()
	if systemTenantId < 1 {
		return &SystemTenantResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := tenant.SystemTenant(ctx, systemTenantId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemTenantId,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenant")
		return &SystemTenantResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: srv.SystemTenantResult(res),
	}, nil
}

// SystemTenantRecover 恢复数据
func (srv SrvSystemTenantServiceServer) SystemTenantRecover(ctx context.Context, request *SystemTenantRecoverRequest) (*SystemTenantRecoverResponse, error) {
	systemTenantId := request.GetSystemTenantId()
	if systemTenantId < 1 {
		return &SystemTenantRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.SystemTenantRecover(ctx, systemTenantId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemTenantId,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenantRecover")
		return &SystemTenantRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemTenantServiceServer) SystemTenantList(ctx context.Context, request *SystemTenantListRequest) (*SystemTenantListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}
	if request.BeginExpireDate != nil {
		condition["beginExpireDate"] = request.GetBeginExpireDate()
	}
	if request.FinishExpireDate != nil {
		condition["finishExpireDate"] = request.GetFinishExpireDate()
	}
	if request.SystemTenantPackageId != nil {
		condition["systemTenantPackageId"] = request.GetSystemTenantPackageId()
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
	list, err := tenant.SystemTenantList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenantList")
		return &SystemTenantListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemTenantObject
	for _, item := range list {
		res = append(res, srv.SystemTenantResult(item))
	}
	return &SystemTenantListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemTenantListTotal 获取总数
func (srv SrvSystemTenantServiceServer) SystemTenantListTotal(ctx context.Context, request *SystemTenantListTotalRequest) (*SystemTenantTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}
	if request.BeginExpireDate != nil {
		condition["beginExpireDate"] = request.GetBeginExpireDate()
	}
	if request.FinishExpireDate != nil {
		condition["finishExpireDate"] = request.GetFinishExpireDate()
	}
	if request.SystemTenantPackageId != nil {
		condition["systemTenantPackageId"] = request.GetSystemTenantPackageId()
	}

	// 获取数据集合
	total, err := tenant.SystemTenantListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenantListTotal")
		return &SystemTenantTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
