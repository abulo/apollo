package order

import (
	"cloud/code"
	"cloud/module/pay/order"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// pay_order 支付订单

// SrvPayOrderServiceServer 支付订单
type SrvPayOrderServiceServer struct {
	UnimplementedPayOrderServiceServer
	Server *xgrpc.Server
}

// PayOrderCreate 创建数据
func (srv SrvPayOrderServiceServer) PayOrderCreate(ctx context.Context, request *PayOrderCreateRequest) (*PayOrderCreateResponse, error) {
	req := PayOrderDao(request.GetData())
	data, err := order.PayOrderCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:支付订单:pay_order:PayOrderCreate")
		return &PayOrderCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayOrderCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// PayOrderUpdate 更新数据
func (srv SrvPayOrderServiceServer) PayOrderUpdate(ctx context.Context, request *PayOrderUpdateRequest) (*PayOrderUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayOrderUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := PayOrderDao(request.GetData())
	_, err := order.PayOrderUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:支付订单:pay_order:PayOrderUpdate")
		return &PayOrderUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayOrderUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayOrderDelete 删除数据
func (srv SrvPayOrderServiceServer) PayOrderDelete(ctx context.Context, request *PayOrderDeleteRequest) (*PayOrderDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayOrderDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := order.PayOrderDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付订单:pay_order:PayOrderDelete")
		return &PayOrderDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayOrderDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayOrder 查询单条数据
func (srv SrvPayOrderServiceServer) PayOrder(ctx context.Context, request *PayOrderRequest) (*PayOrderResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayOrderResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := order.PayOrder(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付订单:pay_order:PayOrder")
		return &PayOrderResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayOrderResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayOrderProto(res),
	}, nil
}

// PayOrderRecover 恢复数据
func (srv SrvPayOrderServiceServer) PayOrderRecover(ctx context.Context, request *PayOrderRecoverRequest) (*PayOrderRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayOrderRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := order.PayOrderRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:支付订单:pay_order:PayOrderRecover")
		return &PayOrderRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayOrderRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvPayOrderServiceServer) PayOrderList(ctx context.Context, request *PayOrderListRequest) (*PayOrderListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.AppId != nil {
		condition["appId"] = request.GetAppId()
	}
	if request.ChannelId != nil {
		condition["channelId"] = request.GetChannelId()
	}
	if request.ChannelCode != nil {
		condition["channelCode"] = request.GetChannelCode()
	}
	if request.MerchantOrderId != nil {
		condition["merchantOrderId"] = request.GetMerchantOrderId()
	}
	if request.No != nil {
		condition["no"] = request.GetNo()
	}
	if request.ChannelOrderNo != nil {
		condition["channelOrderNo"] = request.GetChannelOrderNo()
	}
	if request.BeginCreateTime != nil {
		condition["beginCreateTime"] = request.GetBeginCreateTime()
	}
	if request.FinishCreateTime != nil {
		condition["finishCreateTime"] = request.GetFinishCreateTime()
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
	list, err := order.PayOrderList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:支付订单:pay_order:PayOrderList")
		return &PayOrderListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*PayOrderObject
	for _, item := range list {
		res = append(res, PayOrderProto(item))
	}
	return &PayOrderListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// PayOrderListTotal 获取总数
func (srv SrvPayOrderServiceServer) PayOrderListTotal(ctx context.Context, request *PayOrderListTotalRequest) (*PayOrderTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.AppId != nil {
		condition["appId"] = request.GetAppId()
	}
	if request.ChannelId != nil {
		condition["channelId"] = request.GetChannelId()
	}
	if request.ChannelCode != nil {
		condition["channelCode"] = request.GetChannelCode()
	}
	if request.MerchantOrderId != nil {
		condition["merchantOrderId"] = request.GetMerchantOrderId()
	}
	if request.No != nil {
		condition["no"] = request.GetNo()
	}
	if request.ChannelOrderNo != nil {
		condition["channelOrderNo"] = request.GetChannelOrderNo()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.BeginCreateTime != nil {
		condition["beginCreateTime"] = request.GetBeginCreateTime()
	}
	if request.FinishCreateTime != nil {
		condition["finishCreateTime"] = request.GetFinishCreateTime()
	}

	// 获取数据集合
	total, err := order.PayOrderListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:支付订单:pay_order:PayOrderListTotal")
		return &PayOrderTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayOrderTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
