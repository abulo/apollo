package file

import (
	"cloud/code"
	"cloud/module/system/file"
	"context"
	"encoding/json"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_file 文件管理

// SrvSystemFileServiceServer 文件管理
type SrvSystemFileServiceServer struct {
	UnimplementedSystemFileServiceServer
	Server *xgrpc.Server
}

// SystemFileCreate 创建数据
func (srv SrvSystemFileServiceServer) SystemFileCreate(ctx context.Context, request *SystemFileCreateRequest) (*SystemFileCreateResponse, error) {
	req := SystemFileDao(request.GetData())
	data, err := file.SystemFileCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:文件管理:system_file:SystemFileCreate")
		return &SystemFileCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemFileCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemFileUpdate 更新数据
func (srv SrvSystemFileServiceServer) SystemFileUpdate(ctx context.Context, request *SystemFileUpdateRequest) (*SystemFileUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemFileUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemFileDao(request.GetData())
	_, err := file.SystemFileUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:文件管理:system_file:SystemFileUpdate")
		return &SystemFileUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemFileUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemFileDelete 删除数据
func (srv SrvSystemFileServiceServer) SystemFileDelete(ctx context.Context, request *SystemFileDeleteRequest) (*SystemFileDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemFileDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := file.SystemFileDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:文件管理:system_file:SystemFileDelete")
		return &SystemFileDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemFileDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemFile 查询单条数据
func (srv SrvSystemFileServiceServer) SystemFile(ctx context.Context, request *SystemFileRequest) (*SystemFileResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemFileResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := file.SystemFile(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:文件管理:system_file:SystemFile")
		return &SystemFileResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemFileResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemFileProto(res),
	}, nil
}

// SystemFileRecover 恢复数据
func (srv SrvSystemFileServiceServer) SystemFileRecover(ctx context.Context, request *SystemFileRecoverRequest) (*SystemFileRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemFileRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := file.SystemFileRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:文件管理:system_file:SystemFileRecover")
		return &SystemFileRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemFileRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemFileDrop 清理数据
func (srv SrvSystemFileServiceServer) SystemFileDrop(ctx context.Context, request *SystemFileDropRequest) (*SystemFileDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SystemFileDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := file.SystemFileDrop(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:文件管理:system_file:SystemFileDrop")
		return &SystemFileDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemFileDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemFileServiceServer) SystemFileList(ctx context.Context, request *SystemFileListRequest) (*SystemFileListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.FileType != nil {
		condition["fileType"] = request.GetFileType()
	}
	if request.FileName != nil {
		condition["fileName"] = request.GetFileName()
	}
	if request.DeptId != nil {
		condition["deptId"] = request.GetDeptId()
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
	list, err := file.SystemFileList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:文件管理:system_file:SystemFileList")
		return &SystemFileListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemFileObject
	for _, item := range list {
		res = append(res, SystemFileProto(item))
	}
	return &SystemFileListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemFileListTotal 获取总数
func (srv SrvSystemFileServiceServer) SystemFileListTotal(ctx context.Context, request *SystemFileListTotalRequest) (*SystemFileTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.FileType != nil {
		condition["fileType"] = request.GetFileType()
	}
	if request.FileName != nil {
		condition["fileName"] = request.GetFileName()
	}
	if request.DeptId != nil {
		condition["deptId"] = request.GetDeptId()
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
	total, err := file.SystemFileListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:文件管理:system_file:SystemFileListTotal")
		return &SystemFileTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemFileTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
