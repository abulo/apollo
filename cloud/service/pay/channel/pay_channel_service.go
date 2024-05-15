package channel

import (
	"cloud/code"
	"cloud/module/pay/channel"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// pay_channel 支付渠道

// SrvPayChannelServiceServer 支付渠道
type SrvPayChannelServiceServer struct {
	UnimplementedPayChannelServiceServer
	Server *xgrpc.Server
}

// PayChannelCreate 创建数据
func (srv SrvPayChannelServiceServer) PayChannelCreate(ctx context.Context, request *PayChannelCreateRequest) (*PayChannelCreateResponse, error) {
	req := PayChannelDao(request.GetData())
	data, err := channel.PayChannelCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:支付渠道:pay_channel:PayChannelCreate")
		return &PayChannelCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayChannelCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// PayChannelUpdate 更新数据
func (srv SrvPayChannelServiceServer) PayChannelUpdate(ctx context.Context, request *PayChannelUpdateRequest) (*PayChannelUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayChannelUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := PayChannelDao(request.GetData())
	_, err := channel.PayChannelUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:支付渠道:pay_channel:PayChannelUpdate")
		return &PayChannelUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayChannelUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayChannelDelete 删除数据
func (srv SrvPayChannelServiceServer) PayChannelDelete(ctx context.Context, request *PayChannelDeleteRequest) (*PayChannelDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayChannelDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := channel.PayChannelDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付渠道:pay_channel:PayChannelDelete")
		return &PayChannelDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayChannelDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayChannel 查询单条数据
func (srv SrvPayChannelServiceServer) PayChannel(ctx context.Context, request *PayChannelRequest) (*PayChannelResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayChannelResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := channel.PayChannel(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付渠道:pay_channel:PayChannel")
		return &PayChannelResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayChannelResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayChannelProto(res),
	}, nil
}

// PayChannelRecover 恢复数据
func (srv SrvPayChannelServiceServer) PayChannelRecover(ctx context.Context, request *PayChannelRecoverRequest) (*PayChannelRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayChannelRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := channel.PayChannelRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付渠道:pay_channel:PayChannelRecover")
		return &PayChannelRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayChannelRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayChannelCode 查询单条数据
func (srv SrvPayChannelServiceServer) PayChannelCode(ctx context.Context, request *PayChannelCodeRequest) (*PayChannelCodeResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.AppId != nil {
		condition["appId"] = request.GetAppId()
	}
	if request.Code != nil {
		condition["code"] = request.GetCode()
	}

	if util.Empty(condition) {
		err := errors.New("condition is empty")
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:支付渠道:pay_channel:PayChannelCode")
		return &PayChannelCodeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := channel.PayChannelCode(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:支付渠道:pay_channel:PayChannelCode")
		return &PayChannelCodeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayChannelCodeResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayChannelProto(res),
	}, nil
}
func (srv SrvPayChannelServiceServer) PayChannelList(ctx context.Context, request *PayChannelListRequest) (*PayChannelListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.AppId != nil {
		condition["appId"] = request.GetAppId()
	}
	if request.Code != nil {
		condition["code"] = request.GetCode()
	}

	// 获取数据集合
	list, err := channel.PayChannelList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:支付渠道:pay_channel:PayChannelList")
		return &PayChannelListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*PayChannelObject
	for _, item := range list {
		res = append(res, PayChannelProto(item))
	}
	return &PayChannelListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
