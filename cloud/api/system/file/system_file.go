package file

import (
	"context"
	"errors"
	"io"
	"strings"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/system/file"

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
// SystemFileCreate 创建数据
func SystemFileCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFileCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	fileInfo, err := newCtx.FormFile("file")
	if err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	fileReadInfo, err := fileInfo.Open()
	if err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	defer fileReadInfo.Close()
	byteContainer, err := io.ReadAll(fileReadInfo)
	if err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	kind, curFileType, err := FileType(byteContainer)
	if err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
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
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	if err := newCtx.SaveUploadedFile(fileInfo, fileFullPath); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}

	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	request := &file.SystemFileCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemFile
	reqInfo.FileName = proto.String(fileInfo.Filename)                             // 文件名称
	reqInfo.FileType = proto.Int32(curFileType)                                    // 文件类型
	reqInfo.FileMimeType = proto.String(kind.MIME.Value)                           // 文件MIME
	reqInfo.FileSize = proto.Int64(fileInfo.Size)                                  // 文件大小
	reqInfo.FilePath = proto.String(strings.Join([]string{subDir, fileName}, "/")) // 文件路径
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))                    // 租户
	reqInfo.Creator = null.StringFrom(newCtx.GetString("userName"))                // 添加人
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())                             // 添加时间
	request.Data = file.SystemFileProto(reqInfo)
	// 执行服务
	res, err := client.SystemFileCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFileCreate")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": reqInfo,
	})
}

// SystemFileUpdate 更新数据
func SystemFileUpdate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFileUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &file.SystemFileUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SystemFile
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Updater = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	request.Data = file.SystemFileProto(reqInfo)
	// 执行服务
	res, err := client.SystemFileUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFileUpdate")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemFileDelete 删除数据
func SystemFileDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFileDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	requestItem := &file.SystemFileRequest{}
	requestItem.Id = id
	// 执行服务
	item, err := client.SystemFile(ctx, requestItem)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestItem,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFile")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	fileItem := file.SystemFileDao(item.GetData())
	tenantId := newCtx.GetInt64("tenantId")
	if cast.ToInt64(fileItem.TenantId) != tenantId {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}

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
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	// 删除文件夹
	filePath := strings.Join([]string{initial.Core.Path, "upload", cast.ToString(fileItem.FilePath)}, "/")
	if err := util.Unlink(filePath); err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"filePath": filePath,
			"err":      err,
		}).Error("删除文件失败")
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SystemFile 查询单条数据
func SystemFile(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFile")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	id := cast.ToInt64(newCtx.Param("id"))
	request := &file.SystemFileRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemFile(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFile")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": file.SystemFileDao(res.GetData()),
	})
}

// SystemFileList 列表数据
func SystemFileList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:文件管理:system_file:SystemFileList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := file.NewSystemFileServiceClient(grpcClient)
	// 构造查询条件
	request := &file.SystemFileListRequest{}
	requestTotal := &file.SystemFileListTotalRequest{}

	if val, ok := newCtx.GetQuery("tenantId"); ok {
		request.TenantId = proto.Int64(cast.ToInt64(val))      // 租户
		requestTotal.TenantId = proto.Int64(cast.ToInt64(val)) // 租户
	}
	if val, ok := newCtx.GetQuery("fileType"); ok {
		request.FileType = proto.Int32(cast.ToInt32(val))      // 文件类型
		requestTotal.FileType = proto.Int32(cast.ToInt32(val)) // 文件类型
	}
	if val, ok := newCtx.GetQuery("fileName"); ok {
		request.FileName = proto.String(val)      // 文件名称
		requestTotal.FileName = proto.String(val) // 文件名称
	}

	// 执行服务,获取数据量
	resTotal, err := client.SystemFileListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:文件管理:system_file:SystemFileList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var total int64
	request.PageNum = proto.Int64(cast.ToInt64(newCtx.Query("pageNum")))
	request.PageSize = proto.Int64(cast.ToInt64(newCtx.Query("pageSize")))
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
		newCtx.JSON(consts.StatusOK, utils.H{
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
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"total":    total,
			"list":     list,
			"pageNum":  request.PageNum,
			"pageSize": request.PageSize,
		},
	})
}
