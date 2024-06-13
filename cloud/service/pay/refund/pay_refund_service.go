package refund

import (
	"cloud/code"
	"cloud/module/pay/refund"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// pay_refund 退款订单

// SrvPayRefundServiceServer 退款订单
type SrvPayRefundServiceServer struct {
	UnimplementedPayRefundServiceServer
	Server *xgrpc.Server
}

// PayRefundCreate 创建数据
func (srv SrvPayRefundServiceServer) PayRefundCreate(ctx context.Context, request *PayRefundCreateRequest) (*PayRefundCreateResponse, error) {
	req := PayRefundDao(request.GetData())
	data, err := refund.PayRefundCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:退款订单:pay_refund:PayRefundCreate")
		return &PayRefundCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayRefundCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// PayRefundUpdate 更新数据
func (srv SrvPayRefundServiceServer) PayRefundUpdate(ctx context.Context, request *PayRefundUpdateRequest) (*PayRefundUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayRefundUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := PayRefundDao(request.GetData())
	_, err := refund.PayRefundUpdate(ctx, id, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:退款订单:pay_refund:PayRefundUpdate")
		return &PayRefundUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayRefundUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayRefundDelete 删除数据
func (srv SrvPayRefundServiceServer) PayRefundDelete(ctx context.Context, request *PayRefundDeleteRequest) (*PayRefundDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayRefundDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := refund.PayRefundDelete(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:退款订单:pay_refund:PayRefundDelete")
		return &PayRefundDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayRefundDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// PayRefund 查询单条数据
func (srv SrvPayRefundServiceServer) PayRefund(ctx context.Context, request *PayRefundRequest) (*PayRefundResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayRefundResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := refund.PayRefund(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:退款订单:pay_refund:PayRefund")
		return &PayRefundResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayRefundResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: PayRefundProto(res),
	}, nil
}

// PayRefundRecover 恢复数据
func (srv SrvPayRefundServiceServer) PayRefundRecover(ctx context.Context, request *PayRefundRecoverRequest) (*PayRefundRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &PayRefundRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := refund.PayRefundRecover(ctx, id)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:退款订单:pay_refund:PayRefundRecover")
		return &PayRefundRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayRefundRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvPayRefundServiceServer) PayRefundList(ctx context.Context, request *PayRefundListRequest) (*PayRefundListResponse, error) {
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
	if request.OrderId != nil {
		condition["orderId"] = request.GetOrderId()
	}
	if request.OrderNo != nil {
		condition["orderNo"] = request.GetOrderNo()
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
	list, err := refund.PayRefundList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:退款订单:pay_refund:PayRefundList")
		return &PayRefundListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*PayRefundObject
	for _, item := range list {
		res = append(res, PayRefundProto(item))
	}
	return &PayRefundListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// PayRefundListTotal 获取总数
func (srv SrvPayRefundServiceServer) PayRefundListTotal(ctx context.Context, request *PayRefundListTotalRequest) (*PayRefundTotalResponse, error) {
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
	if request.OrderId != nil {
		condition["orderId"] = request.GetOrderId()
	}
	if request.OrderNo != nil {
		condition["orderNo"] = request.GetOrderNo()
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
	total, err := refund.PayRefundListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:退款订单:pay_refund:PayRefundListTotal")
		return &PayRefundTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &PayRefundTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
