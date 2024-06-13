package notice

import (
	"cloud/code"
	"cloud/module/system/notice"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_notice 通知公告表

// SrvSystemNoticeServiceServer 通知公告表
type SrvSystemNoticeServiceServer struct {
	UnimplementedSystemNoticeServiceServer
	Server *xgrpc.Server
}

// SystemNoticeCreate 创建数据
func (srv SrvSystemNoticeServiceServer) SystemNoticeCreate(ctx context.Context, request *SystemNoticeCreateRequest) (*SystemNoticeCreateResponse, error) {
	req := SystemNoticeDao(request.GetData())
	data, err := notice.SystemNoticeCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:通知公告表:system_notice:SystemNoticeCreate")
		return &SystemNoticeCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNoticeCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemNoticeUpdate 更新数据
func (srv SrvSystemNoticeServiceServer) SystemNoticeUpdate(ctx context.Context, request *SystemNoticeUpdateRequest) (*SystemNoticeUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNoticeUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemNoticeDao(request.GetData())
	_, err := notice.SystemNoticeUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:通知公告表:system_notice:SystemNoticeUpdate")
		return &SystemNoticeUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNoticeUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemNoticeDelete 删除数据
func (srv SrvSystemNoticeServiceServer) SystemNoticeDelete(ctx context.Context, request *SystemNoticeDeleteRequest) (*SystemNoticeDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNoticeDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := notice.SystemNoticeDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:通知公告表:system_notice:SystemNoticeDelete")
		return &SystemNoticeDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNoticeDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemNotice 查询单条数据
func (srv SrvSystemNoticeServiceServer) SystemNotice(ctx context.Context, request *SystemNoticeRequest) (*SystemNoticeResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNoticeResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := notice.SystemNotice(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:通知公告表:system_notice:SystemNotice")
		return &SystemNoticeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNoticeResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemNoticeProto(res),
	}, nil
}

// SystemNoticeRecover 恢复数据
func (srv SrvSystemNoticeServiceServer) SystemNoticeRecover(ctx context.Context, request *SystemNoticeRecoverRequest) (*SystemNoticeRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemNoticeRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := notice.SystemNoticeRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:通知公告表:system_notice:SystemNoticeRecover")
		return &SystemNoticeRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNoticeRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemNoticeServiceServer) SystemNoticeList(ctx context.Context, request *SystemNoticeListRequest) (*SystemNoticeListResponse, error) {
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
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.Title != nil {
		condition["title"] = request.GetTitle()
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
	list, err := notice.SystemNoticeList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:通知公告表:system_notice:SystemNoticeList")
		return &SystemNoticeListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemNoticeObject
	for _, item := range list {
		res = append(res, SystemNoticeProto(item))
	}
	return &SystemNoticeListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemNoticeListTotal 获取总数
func (srv SrvSystemNoticeServiceServer) SystemNoticeListTotal(ctx context.Context, request *SystemNoticeListTotalRequest) (*SystemNoticeTotalResponse, error) {
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
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.Title != nil {
		condition["title"] = request.GetTitle()
	}

	// 获取数据集合
	total, err := notice.SystemNoticeListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:通知公告表:system_notice:SystemNoticeListTotal")
		return &SystemNoticeTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemNoticeTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
