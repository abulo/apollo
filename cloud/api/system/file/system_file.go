package file

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strings"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/system/file"
	"cloud/service/system/user"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/matchers"
	"github.com/h2non/filetype/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var pdfType = filetype.NewType("pdf", "application/pdf")

func pdfMatcher(buf []byte) bool {
	return len(buf) > 4 && buf[0] == 0x25 && buf[1] == 0x50 && buf[2] == 0x44 && buf[3] == 0x46
}

func FileType(buf []byte) (kind types.Type, fileType int32, err error) {
	filetype.AddMatcher(pdfType, pdfMatcher)
	kind, err = filetype.Match(buf)
	// 内部错误
	if err != nil {
		return
	}
	// 未识别
	if kind == filetype.Unknown {
		err = errors.New("unknown file type")
		return
	}
	if filetype.IsImage(buf) {
		fileType = 1
		return
	}
	if filetype.IsDocument(buf) {
		if matchers.Doc(buf) || matchers.Docx(buf) {
			fileType = 2
			return
		}
		if matchers.Xls(buf) || matchers.Xlsx(buf) {
			fileType = 3
			return
		}
		if matchers.Ppt(buf) || matchers.Pptx(buf) {
			fileType = 4
			return
		}
		err = errors.New("unknown file type")
		return
	}
	if filetype.IsSupported("pdf") {
		fileType = 5
		return
	}
	err = errors.New("unknown file type")
	return
}

// system_file 文件管理
// SystemFileItem 查询单条数据
func SystemFileItem(ctx context.Context, newCtx *app.RequestContext, id int64) (*file.SystemFileResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFileItem")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	request := &file.SystemFileRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemFile(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFileItem")
		return nil, err
	}
	tenantId := cast.ToInt64(newCtx.GetInt64("tenantId"))
	data := file.SystemFileDao(res.GetData())
	if cast.ToInt64(data.TenantId) != tenantId {
		return nil, status.Error(code.ConvertToGrpc(code.NoPermission), code.StatusText(code.NoPermission))
	}
	return res, nil
}

// SystemFileCreate 创建数据
func SystemFileCreate(ctx context.Context, newCtx *app.RequestContext) {
	fileInfo, err := newCtx.FormFile("file")
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	fileReadInfo, err := fileInfo.Open()
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	defer fileReadInfo.Close()
	byteContainer, err := io.ReadAll(fileReadInfo)
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	kind, curFileType, err := FileType(byteContainer)
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	tenantDir := cast.ToString(newCtx.GetInt64("tenantId"))
	// 新的文件名称
	fileName := primitive.NewObjectID().Hex() + "." + kind.Extension

	subDir := tenantDir + "/" + util.Date("Y/m/d", util.Now())
	fileDir := strings.Join([]string{initial.Core.Path, "upload", subDir}, "/")
	fileFullPath := strings.Join([]string{fileDir, fileName}, "/")
	if err := util.MkdirAll(fileDir, 0777); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	if err := newCtx.SaveUploadedFile(fileInfo, fileFullPath); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}

	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFileCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	request := &file.SystemFileCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemFile

	reqInfo.Id = nil
	reqInfo.FileName = proto.String(fileInfo.Filename)                             // 文件名称
	reqInfo.FileType = proto.Int32(curFileType)                                    // 文件类型
	reqInfo.FileMimeType = proto.String(kind.MIME.Value)                           // 文件MIME
	reqInfo.FileSize = proto.Int64(fileInfo.Size)                                  // 文件大小
	reqInfo.FilePath = proto.String(strings.Join([]string{subDir, fileName}, "/")) // 文件路径
	reqInfo.UserId = proto.Int64(newCtx.GetInt64("userId"))                        // 用户 ID
	reqInfo.Deleted = proto.Int32(0)
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户
	reqInfo.Creator = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	reqInfo.Updater = null.StringFromPtr(nil)
	reqInfo.UpdateTime = null.DateTimeFromPtr(nil)
	request.Data = file.SystemFileProto(reqInfo)
	// 执行服务
	res, err := client.SystemFileCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFileCreate")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": reqInfo,
	})
}

// SystemFileReUpload 更新数据
func SystemFileReUpload(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	fileItem, err := SystemFileItem(ctx, newCtx, id)
	if err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}

	fileInfo, err := newCtx.FormFile("file")
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	fileReadInfo, err := fileInfo.Open()
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	defer fileReadInfo.Close()
	byteContainer, err := io.ReadAll(fileReadInfo)
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	kind, curFileType, err := FileType(byteContainer)
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	tenantDir := cast.ToString(newCtx.GetInt64("tenantId"))
	// 新的文件名称
	fileName := primitive.NewObjectID().Hex() + "." + kind.Extension

	subDir := tenantDir + "/" + util.Date("Y/m/d", util.Now())
	fileDir := strings.Join([]string{initial.Core.Path, "upload", subDir}, "/")
	fileFullPath := strings.Join([]string{fileDir, fileName}, "/")
	if err := util.MkdirAll(fileDir, 0777); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	if err := newCtx.SaveUploadedFile(fileInfo, fileFullPath); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}

	// 查询原来数据, 需要将原来的文件删除才能更新
	fileDao := file.SystemFileDao(fileItem.GetData())
	filePath := strings.Join([]string{initial.Core.Path, "upload", cast.ToString(fileDao.FilePath)}, "/")
	if err := util.Unlink(filePath); err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"filePath": filePath,
			"err":      err,
		}).Error("删除文件失败")
	}
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFileReUpload")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	request := &file.SystemFileUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SystemFile
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Id = nil
	reqInfo.FileName = proto.String(fileInfo.Filename)                             // 文件名称
	reqInfo.FileType = proto.Int32(curFileType)                                    // 文件类型
	reqInfo.FileMimeType = proto.String(kind.MIME.Value)                           // 文件MIME
	reqInfo.FileSize = proto.Int64(fileInfo.Size)                                  // 文件大小
	reqInfo.FilePath = proto.String(strings.Join([]string{subDir, fileName}, "/")) // 文件路径
	reqInfo.UserId = fileDao.UserId                                                // 用户 ID
	reqInfo.Deleted = fileDao.Deleted
	reqInfo.TenantId = fileDao.TenantId // 租户
	reqInfo.Creator = fileDao.Creator
	reqInfo.CreateTime = fileDao.CreateTime
	reqInfo.Updater = null.StringFromPtr(nil)
	reqInfo.UpdateTime = null.DateTimeFromPtr(nil)
	request.Data = file.SystemFileProto(reqInfo)
	// 执行服务
	res, err := client.SystemFileUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFileReUpload")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": reqInfo,
	})
}

// SystemFileUpdate 更新数据
func SystemFileUpdate(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	fileItem, err := SystemFileItem(ctx, newCtx, id)
	if err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFileUpdate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	request := &file.SystemFileUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SystemFile
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	fileDao := file.SystemFileDao(fileItem.GetData())
	reqInfo.Id = nil
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户
	reqInfo.Updater = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	reqInfo.Creator = null.StringFromPtr(nil)
	reqInfo.CreateTime = null.DateTimeFromPtr(nil)
	reqInfo.FileType = fileDao.FileType         // 文件类型
	reqInfo.FileMimeType = fileDao.FileMimeType // 文件MIME
	reqInfo.FileSize = fileDao.FileSize         // 文件大小
	reqInfo.FilePath = fileDao.FilePath         // 文件路径
	reqInfo.UserId = fileDao.UserId             // 用户 ID
	request.Data = file.SystemFileProto(reqInfo)
	// 执行服务
	res, err := client.SystemFileUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFileUpdate")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemFileDelete 删除数据
func SystemFileDelete(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SystemFileItem(ctx, newCtx, id); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFileDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	request := &file.SystemFileDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemFileDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFileDelete")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemFile 查询单条数据
func SystemFile(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	// 执行服务
	res, err := SystemFileItem(ctx, newCtx, id)
	if err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": file.SystemFileDao(res.GetData()),
	})
}

// SystemFileRecover 恢复数据
func SystemFileRecover(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SystemFileItem(ctx, newCtx, id); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFileRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	request := &file.SystemFileRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemFileRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFileRecover")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemFileDrop 清理数据
func SystemFileDrop(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	fileItem, err := SystemFileItem(ctx, newCtx, id)
	if err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFileDrop")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	request := &file.SystemFileDropRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemFileDrop(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFileDrop")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	fileDao := file.SystemFileDao(fileItem.GetData())
	filePath := strings.Join([]string{initial.Core.Path, "upload", cast.ToString(fileDao.FilePath)}, "/")
	if err := util.Unlink(filePath); err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"filePath": filePath,
			"err":      err,
		}).Error("删除文件失败")
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemFileList 列表数据
func SystemFileList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFileList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}

	//链接服务
	clientUser := user.NewSystemUserServiceClient(grpcClient)
	userDataScopeReq := &user.SystemUserRoleDataScopeRequest{}
	userDataScopeReq.UserId = newCtx.GetInt64("userId")
	userDataScopeReq.TenantId = newCtx.GetInt64("tenantId")
	userRes, err := clientUser.SystemUserRoleDataScope(ctx, userDataScopeReq)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": userDataScopeReq,
			"err": err,
		}).Error("GrpcCall:部门:system_dept:SystemDeptList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	userScope := user.SystemUserRoleDataScopeDao(userRes.GetData())
	if len(userScope.DataScopeDept) < 1 {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.DeptError,
			"msg":  code.StatusText(code.DeptError),
		})
		return
	}

	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	// 构造查询条件
	request := &file.SystemFileListRequest{}
	requestTotal := &file.SystemFileListTotalRequest{}
	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))      // 租户ID
	requestTotal.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                                 // 删除状态
	requestTotal.Deleted = proto.Int32(0)                            // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
			requestTotal.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("fileType"); ok {
		request.FileType = proto.Int32(cast.ToInt32(val))      // 文件类型
		requestTotal.FileType = proto.Int32(cast.ToInt32(val)) // 文件类型
	}
	if val, ok := newCtx.GetQuery("fileName"); ok {
		request.FileName = proto.String(val)      // 文件名称
		requestTotal.FileName = proto.String(val) // 文件名称
	}
	request.UserId = proto.Int64(newCtx.GetInt64("userId"))      // 用户ID
	requestTotal.UserId = proto.Int64(newCtx.GetInt64("userId")) // 用户ID
	if val, ok := newCtx.GetQuery("deptId"); ok {
		request.DeptId = proto.Int64(cast.ToInt64(val))      // 部门
		requestTotal.DeptId = proto.Int64(cast.ToInt64(val)) // 部门
	}
	request.DataScope = userScope.DataScope
	requestTotal.DataScope = userScope.DataScope
	dataScopeDept, _ := json.Marshal(userScope.DataScopeDept)
	request.DataScopeDept = dataScopeDept
	requestTotal.DataScopeDept = dataScopeDept

	// 执行服务,获取数据量
	resTotal, err := client.SystemFileListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFileList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var total int64
	paginationRequest := &pagination.PaginationRequest{}
	paginationRequest.PageNum = proto.Int64(cast.ToInt64(newCtx.Query("pageNum")))
	paginationRequest.PageSize = proto.Int64(cast.ToInt64(newCtx.Query("pageSize")))
	request.Pagination = paginationRequest
	if resTotal.GetCode() == code.Success {
		total = resTotal.GetData()
	}
	// 执行服务
	res, err := client.SystemFileList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFileList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemFile
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, file.SystemFileDao(item))
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"total":    total,
			"list":     list,
			"pageNum":  paginationRequest.PageNum,
			"pageSize": paginationRequest.PageSize,
		},
	})
}
