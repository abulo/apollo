package user

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/tools"
	"cloud/service/captcha"
	"cloud/service/system/user"

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

// system_user 系统用户

// SystemUserDao 数据转换
func SystemUserDao(item *user.SystemUserObject) *dao.SystemUser {
	daoItem := &dao.SystemUser{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 用户编号
	}
	if item != nil && item.Nickname != nil {
		daoItem.Nickname = null.StringFrom(item.GetNickname()) // 昵称
	}
	if item != nil && item.Username != nil {
		daoItem.Username = item.Username // 用户名称
	}
	if item != nil && item.Password != nil {
		daoItem.Password = item.Password // 用户密码
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 用户状态（0正常 1停用）
	}
	if item != nil && item.Creator != nil {
		daoItem.Creator = null.StringFrom(item.GetCreator()) // 创建人
	}
	if item != nil && item.CreateTime != nil {
		daoItem.CreateTime = null.DateTimeFrom(util.GrpcTime(item.CreateTime)) // 创建时间
	}
	if item != nil && item.Updater != nil {
		daoItem.Updater = null.StringFrom(item.GetUpdater()) // 更新人
	}
	if item != nil && item.UpdateTime != nil {
		daoItem.UpdateTime = null.DateTimeFrom(util.GrpcTime(item.UpdateTime)) // 更新时间
	}

	return daoItem
}

// SystemUserProto 数据绑定
func SystemUserProto(item dao.SystemUser) *user.SystemUserObject {
	res := &user.SystemUserObject{}
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
func SystemUserCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:system_user:SystemUserCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	request := &user.SystemUserCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemUser
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request.Data = SystemUserProto(reqInfo)
	// 执行服务
	res, err := client.SystemUserCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUserCreate")
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

// SystemUserUpdate 更新数据
func SystemUserUpdate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:system_user:SystemUserUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	systemUserId := cast.ToInt64(newCtx.Param("systemUserId"))
	request := &user.SystemUserUpdateRequest{}
	request.SystemUserId = systemUserId
	// 数据绑定
	var reqInfo dao.SystemUser
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	request.Data = SystemUserProto(reqInfo)
	// 执行服务
	res, err := client.SystemUserUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUserUpdate")
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

// SystemUserDelete 删除数据
func SystemUserDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:system_user:SystemUserDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	systemUserId := cast.ToInt64(newCtx.Param("systemUserId"))
	request := &user.SystemUserDeleteRequest{}
	request.SystemUserId = systemUserId
	// 执行服务
	res, err := client.SystemUserDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUserDelete")
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

// SystemUser 查询单条数据
func SystemUser(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:system_user:SystemUser")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	systemUserId := cast.ToInt64(newCtx.Param("systemUserId"))
	request := &user.SystemUserRequest{}
	request.SystemUserId = systemUserId
	// 执行服务
	res, err := client.SystemUser(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUser")
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
		"data": SystemUserDao(res.GetData()),
	})
}

// SystemUserList 列表数据
func SystemUserList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:system_user:SystemUserList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	// 构造查询条件
	request := &user.SystemUserListRequest{}
	requestTotal := &user.SystemUserListTotalRequest{}

	if val, ok := newCtx.GetQuery("username"); ok {
		request.Username = proto.String(val)      // 用户名称
		requestTotal.Username = proto.String(val) // 用户名称
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 用户状态（0正常 1停用）
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 用户状态（0正常 1停用）
	}

	// 执行服务,获取数据量
	resTotal, err := client.SystemUserListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUserList")
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
	res, err := client.SystemUserList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemUserList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemUser
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, SystemUserDao(item))
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

// SystemUserLogin 查询单条数据
func SystemUserLogin(ctx context.Context, newCtx *app.RequestContext) {
	//1.验证参数必传
	var req dao.SystemUserLogin
	if err := newCtx.BindAndValidate(&req); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}

	//2.判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:用户信息表:system_user:SystemUserLogin")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}

	//链接验证码服务
	clientCaptcha := captcha.NewCaptchaServiceClient(grpcClient)
	requestCaptcha := &captcha.CaptchaVerifyRequest{}
	requestCaptcha.CaptchaId = req.CaptchaId
	requestCaptcha.CaptchaCode = req.CaptchaCode
	// 执行服务
	resCaptcha, err := clientCaptcha.CaptchaVerify(ctx, requestCaptcha)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestCaptcha,
			"err": err,
		}).Error("GrpcCall:用户信息表:system_user:SystemUserLogin")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	if resCaptcha.GetCode() != code.Success {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": resCaptcha.GetCode(),
			"msg":  resCaptcha.GetMsg(),
		})
		return
	}
	if !resCaptcha.GetData().GetResult() {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.VerifyCodeError,
			"msg":  code.StatusText(code.VerifyCodeError),
		})
		return
	}
	//3.链接用户信息表服务
	clientSystemUser := user.NewSystemUserServiceClient(grpcClient)
	requestSystemUser := &user.SystemUserLoginRequest{}
	requestSystemUser.Username = proto.String(req.Username)
	// 执行服务
	res, err := clientSystemUser.SystemUserLogin(ctx, requestSystemUser)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestSystemUser,
			"err": err,
		}).Error("GrpcCall:用户信息表:system_user:SystemUserLogin")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	//4.生成 token
	userInfo := res.GetData()
	if userInfo.Id == nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.LoginFail,
			"msg":  code.StatusText(code.LoginFail),
		})
		return
	}
	// 比对密码
	if req.Password != userInfo.GetPassword() {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.LoginFail,
			"msg":  code.StatusText(code.LoginFail),
		})
		return
	}
	token, err := tools.GenerateToken(*userInfo.Id, *userInfo.Username, *userInfo.Nickname, "WEB")
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestCaptcha,
			"err": err,
		}).Error("GrpcCall:用户信息表:system_user:SystemUserLogin")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.TokenError,
			"msg":  code.StatusText(code.TokenError),
		})
		return
	}
	// TODO 需要添加后台用户登录日志
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"accessToken": token,
			"nickname":    userInfo.GetNickname(),
		},
	})
}
