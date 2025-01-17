package mail

import (
	"context"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/system/mail"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_mail_log 邮件日志表
// SystemMailLogItem 查询单条数据
func SystemMailLogItem(ctx context.Context, newCtx *app.RequestContext, id int64) (*mail.SystemMailLogResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:邮件日志表:system_mail_log:SystemMailLogItem")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := mail.NewSystemMailLogServiceClient(grpcClient)
	request := &mail.SystemMailLogRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemMailLog(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:邮件日志表:system_mail_log:SystemMailLogItem")
		return nil, err
	}
	return res, nil
}

// SystemMailLogCreate 创建数据
func SystemMailLogCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:邮件日志表:system_mail_log:SystemMailLogCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := mail.NewSystemMailLogServiceClient(grpcClient)
	request := &mail.SystemMailLogCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemMailLog
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Id = nil
	reqInfo.Deleted = proto.Int32(0)
	reqInfo.Creator = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	request.Data = mail.SystemMailLogProto(reqInfo)
	// 执行服务
	res, err := client.SystemMailLogCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:邮件日志表:system_mail_log:SystemMailLogCreate")
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

// SystemMailLogUpdate 更新数据
func SystemMailLogUpdate(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SystemMailLogItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:邮件日志表:system_mail_log:SystemMailLogUpdate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := mail.NewSystemMailLogServiceClient(grpcClient)
	request := &mail.SystemMailLogUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SystemMailLog
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Id = nil
	reqInfo.Updater = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	reqInfo.Creator = null.StringFromPtr(nil)
	reqInfo.CreateTime = null.DateTimeFromPtr(nil)
	request.Data = mail.SystemMailLogProto(reqInfo)
	// 执行服务
	res, err := client.SystemMailLogUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:邮件日志表:system_mail_log:SystemMailLogUpdate")
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

// SystemMailLogDelete 删除数据
func SystemMailLogDelete(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SystemMailLogItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:邮件日志表:system_mail_log:SystemMailLogDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := mail.NewSystemMailLogServiceClient(grpcClient)
	request := &mail.SystemMailLogDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemMailLogDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:邮件日志表:system_mail_log:SystemMailLogDelete")
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

// SystemMailLog 查询单条数据
func SystemMailLog(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	// 执行服务
	res, err := SystemMailLogItem(ctx, newCtx, id)
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
		"data": mail.SystemMailLogDao(res.GetData()),
	})
}

// SystemMailLogRecover 恢复数据
func SystemMailLogRecover(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SystemMailLogItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:邮件日志表:system_mail_log:SystemMailLogRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := mail.NewSystemMailLogServiceClient(grpcClient)
	request := &mail.SystemMailLogRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemMailLogRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:邮件日志表:system_mail_log:SystemMailLogRecover")
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

// SystemMailLogDrop 清理数据
func SystemMailLogDrop(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SystemMailLogItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:邮件日志表:system_mail_log:SystemMailLogDrop")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := mail.NewSystemMailLogServiceClient(grpcClient)
	request := &mail.SystemMailLogDropRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SystemMailLogDrop(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:邮件日志表:system_mail_log:SystemMailLogDrop")
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

// SystemMailLogList 列表数据
func SystemMailLogList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:邮件日志表:system_mail_log:SystemMailLogList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := mail.NewSystemMailLogServiceClient(grpcClient)
	// 构造查询条件
	request := &mail.SystemMailLogListRequest{}
	requestTotal := &mail.SystemMailLogListTotalRequest{}

	request.Deleted = proto.Int32(0)      // 删除状态
	requestTotal.Deleted = proto.Int32(0) // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
			requestTotal.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("sendStatus"); ok {
		request.SendStatus = proto.Int32(cast.ToInt32(val))      // 发送状态
		requestTotal.SendStatus = proto.Int32(cast.ToInt32(val)) // 发送状态
	}
	if val, ok := newCtx.GetQuery("beginSendTime"); ok {
		request.BeginSendTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 登录时间
		requestTotal.BeginSendTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 登录时间
	}
	if val, ok := newCtx.GetQuery("finishSendTime"); ok {
		request.FinishSendTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 登录时间
		requestTotal.FinishSendTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 登录时间
	}
	if val, ok := newCtx.GetQuery("templateTitle"); ok {
		request.TemplateTitle = proto.String(val)      // 邮件标题
		requestTotal.TemplateTitle = proto.String(val) // 邮件标题
	}
	if val, ok := newCtx.GetQuery("userId"); ok {
		request.UserId = proto.Int64(cast.ToInt64(val))      // 用户编号
		requestTotal.UserId = proto.Int64(cast.ToInt64(val)) // 用户编号
	}
	if val, ok := newCtx.GetQuery("userType"); ok {
		request.UserType = proto.Int32(cast.ToInt32(val))      // 用户类型
		requestTotal.UserType = proto.Int32(cast.ToInt32(val)) // 用户类型
	}
	if val, ok := newCtx.GetQuery("accountId"); ok {
		request.AccountId = proto.Int64(cast.ToInt64(val))      // 邮箱账号编号
		requestTotal.AccountId = proto.Int64(cast.ToInt64(val)) // 邮箱账号编号
	}
	if val, ok := newCtx.GetQuery("templateCode"); ok {
		request.TemplateCode = proto.String(val)      // 模板编码
		requestTotal.TemplateCode = proto.String(val) // 模板编码
	}

	// 执行服务,获取数据量
	resTotal, err := client.SystemMailLogListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:邮件日志表:system_mail_log:SystemMailLogList")
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
	res, err := client.SystemMailLogList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:邮件日志表:system_mail_log:SystemMailLogList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemMailLog
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, mail.SystemMailLogDao(item))
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

// SystemMailLogMultipleDelete 批量删除数据
func SystemMailLogMultipleDelete(ctx context.Context, newCtx *app.RequestContext) {
	var reqInfo dao.SystemMailLogMultiple
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:邮件日志表:system_mail_log:SystemMailLogDrop")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := mail.NewSystemMailLogServiceClient(grpcClient)
	request := &mail.SystemMailLogMultipleRequest{}
	if reqInfo.Ids.IsValid() {
		request.Ids = *reqInfo.Ids.Ptr()
	}
	// 执行服务
	res, err := client.SystemMailLogMultipleDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:邮件日志表:system_mail_log:SystemMailLogMultipleDelete")
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

// SystemMailLogMultipleRecover 批量恢复数据
func SystemMailLogMultipleRecover(ctx context.Context, newCtx *app.RequestContext) {
	var reqInfo dao.SystemMailLogMultiple
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:邮件日志表:system_mail_log:SystemMailLogDrop")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := mail.NewSystemMailLogServiceClient(grpcClient)
	request := &mail.SystemMailLogMultipleRequest{}
	if reqInfo.Ids.IsValid() {
		request.Ids = *reqInfo.Ids.Ptr()
	}
	// 执行服务
	res, err := client.SystemMailLogMultipleRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:邮件日志表:system_mail_log:SystemMailLogMultipleRecover")
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

// SystemMailLogMultipleDrop 批量清理数据
func SystemMailLogMultipleDrop(ctx context.Context, newCtx *app.RequestContext) {
	var reqInfo dao.SystemMailLogMultiple
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:邮件日志表:system_mail_log:SystemMailLogDrop")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := mail.NewSystemMailLogServiceClient(grpcClient)
	request := &mail.SystemMailLogMultipleRequest{}
	if reqInfo.Ids.IsValid() {
		request.Ids = *reqInfo.Ids.Ptr()
	}
	// 执行服务
	res, err := client.SystemMailLogMultipleDrop(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:邮件日志表:system_mail_log:SystemMailLogMultipleDrop")
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
