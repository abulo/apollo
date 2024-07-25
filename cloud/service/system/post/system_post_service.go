package post

import (
	"cloud/code"
	"cloud/module/system/post"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_post 职位

// SrvSystemPostServiceServer 职位
type SrvSystemPostServiceServer struct {
	UnimplementedSystemPostServiceServer
	Server *xgrpc.Server
}

// SystemPostCreate 创建数据
func (srv SrvSystemPostServiceServer) SystemPostCreate(ctx context.Context, request *SystemPostCreateRequest) (*SystemPostCreateResponse, error) {
	req := SystemPostDao(request.GetData())
	data, err := post.SystemPostCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:职位:system_post:SystemPostCreate")
		return &SystemPostCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemPostCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemPostUpdate 更新数据
func (srv SrvSystemPostServiceServer) SystemPostUpdate(ctx context.Context, request *SystemPostUpdateRequest) (*SystemPostUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemPostUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemPostDao(request.GetData())
	_, err := post.SystemPostUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:职位:system_post:SystemPostUpdate")
		return &SystemPostUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemPostUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemPostDelete 删除数据
func (srv SrvSystemPostServiceServer) SystemPostDelete(ctx context.Context, request *SystemPostDeleteRequest) (*SystemPostDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemPostDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := post.SystemPostDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:职位:system_post:SystemPostDelete")
		return &SystemPostDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemPostDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemPost 查询单条数据
func (srv SrvSystemPostServiceServer) SystemPost(ctx context.Context, request *SystemPostRequest) (*SystemPostResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemPostResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := post.SystemPost(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:职位:system_post:SystemPost")
		return &SystemPostResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemPostResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemPostProto(res),
	}, nil
}

// SystemPostRecover 恢复数据
func (srv SrvSystemPostServiceServer) SystemPostRecover(ctx context.Context, request *SystemPostRecoverRequest) (*SystemPostRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemPostRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := post.SystemPostRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:职位:system_post:SystemPostRecover")
		return &SystemPostRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemPostRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemPostDrop 清理数据
func (srv SrvSystemPostServiceServer) SystemPostDrop(ctx context.Context, request *SystemPostDropRequest) (*SystemPostDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemPostDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := post.SystemPostDrop(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:职位:system_post:SystemPostDrop")
		return &SystemPostDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemPostDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemPostServiceServer) SystemPostList(ctx context.Context, request *SystemPostListRequest) (*SystemPostListResponse, error) {
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
	if request.Name != nil {
		condition["name"] = request.GetName()
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
	list, err := post.SystemPostList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:职位:system_post:SystemPostList")
		return &SystemPostListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemPostObject
	for _, item := range list {
		res = append(res, SystemPostProto(item))
	}
	return &SystemPostListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemPostListTotal 获取总数
func (srv SrvSystemPostServiceServer) SystemPostListTotal(ctx context.Context, request *SystemPostListTotalRequest) (*SystemPostTotalResponse, error) {
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
	if request.Name != nil {
		condition["name"] = request.GetName()
	}

	// 获取数据集合
	total, err := post.SystemPostListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:职位:system_post:SystemPostListTotal")
		return &SystemPostTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemPostTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
