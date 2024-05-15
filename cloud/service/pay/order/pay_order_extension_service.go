package order

import (
	"cloud/code"
	"cloud/module/pay/order"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// pay_order_extension 支付订单拓展

// SrvPayOrderExtensionServiceServer 支付订单拓展
type SrvPayOrderExtensionServiceServer struct {
	UnimplementedPayOrderExtensionServiceServer
	Server *xgrpc.Server
}

// PayOrderExtensionCreate 创建数据
func (srv SrvPayOrderExtensionServiceServer) PayOrderExtensionCreate(ctx context.Context, request *PayOrderExtensionCreateRequest) (*PayOrderExtensionCreateResponse, error) {
	req := PayOrderExtensionDao(request.GetData())
	data, err := order.PayOrderExtensionCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:支付订单拓展:pay_order_extension:PayOrderExtensionCreate")
		return &PayOrderExtensionCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayOrderExtensionCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// PayOrderExtensionUpdate 更新数据
func (srv SrvPayOrderExtensionServiceServer) PayOrderExtensionUpdate(ctx context.Context, request *PayOrderExtensionUpdateRequest) (*PayOrderExtensionUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayOrderExtensionUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := PayOrderExtensionDao(request.GetData())
	_, err := order.PayOrderExtensionUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:支付订单拓展:pay_order_extension:PayOrderExtensionUpdate")
		return &PayOrderExtensionUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayOrderExtensionUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayOrderExtensionDelete 删除数据
func (srv SrvPayOrderExtensionServiceServer) PayOrderExtensionDelete(ctx context.Context, request *PayOrderExtensionDeleteRequest) (*PayOrderExtensionDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayOrderExtensionDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := order.PayOrderExtensionDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付订单拓展:pay_order_extension:PayOrderExtensionDelete")
		return &PayOrderExtensionDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayOrderExtensionDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayOrderExtension 查询单条数据
func (srv SrvPayOrderExtensionServiceServer) PayOrderExtension(ctx context.Context, request *PayOrderExtensionRequest) (*PayOrderExtensionResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayOrderExtensionResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := order.PayOrderExtension(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付订单拓展:pay_order_extension:PayOrderExtension")
		return &PayOrderExtensionResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayOrderExtensionResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayOrderExtensionProto(res),
	}, nil
}

// PayOrderExtensionRecover 恢复数据
func (srv SrvPayOrderExtensionServiceServer) PayOrderExtensionRecover(ctx context.Context, request *PayOrderExtensionRecoverRequest) (*PayOrderExtensionRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayOrderExtensionRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := order.PayOrderExtensionRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付订单拓展:pay_order_extension:PayOrderExtensionRecover")
		return &PayOrderExtensionRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayOrderExtensionRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayOrderExtensionItem 查询单条数据
func (srv SrvPayOrderExtensionServiceServer) PayOrderExtensionItem(ctx context.Context, request *PayOrderExtensionItemRequest) (*PayOrderExtensionItemResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.OrderId != nil {
		condition["orderId"] = request.GetOrderId()
	}

	if util.Empty(condition) {
		err := errors.New("condition is empty")
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:支付订单拓展:pay_order_extension:PayOrderExtensionItem")
		return &PayOrderExtensionItemResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := order.PayOrderExtensionItem(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:支付订单拓展:pay_order_extension:PayOrderExtensionItem")
		return &PayOrderExtensionItemResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayOrderExtensionItemResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayOrderExtensionProto(res),
	}, nil
}
func (srv SrvPayOrderExtensionServiceServer) PayOrderExtensionList(ctx context.Context, request *PayOrderExtensionListRequest) (*PayOrderExtensionListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.OrderId != nil {
		condition["orderId"] = request.GetOrderId()
	}

	// 获取数据集合
	list, err := order.PayOrderExtensionList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:支付订单拓展:pay_order_extension:PayOrderExtensionList")
		return &PayOrderExtensionListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*PayOrderExtensionObject
	for _, item := range list {
		res = append(res, PayOrderExtensionProto(item))
	}
	return &PayOrderExtensionListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}
