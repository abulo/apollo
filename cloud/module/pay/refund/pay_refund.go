package refund

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
)

// pay_refund 退款订单
// PayRefundCreate 创建数据
func PayRefundCreate(ctx context.Context, data dao.PayRefund) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_refund`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// PayRefundUpdate 更新数据
func PayRefundUpdate(ctx context.Context, id int64, data dao.PayRefund) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_refund`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayRefundDelete 删除数据
func PayRefundDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`pay_refund`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayRefund 查询单条数据
func PayRefund(ctx context.Context, id int64) (res dao.PayRefund, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_refund`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// PayRefundRecover 恢复数据
func PayRefundRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`pay_refund`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayRefundList 查询列表数据
func PayRefundList(ctx context.Context, condition map[string]any) (res []dao.PayRefund, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_refund`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["appId"]; ok {
		builder.Where("`app_id`", val)
	}
	if val, ok := condition["channelId"]; ok {
		builder.Where("`channel_id`", val)
	}
	if val, ok := condition["channelCode"]; ok {
		builder.Where("`channel_code`", val)
	}
	if val, ok := condition["merchantOrderId"]; ok {
		builder.Where("`merchant_order_id`", val)
	}
	if val, ok := condition["orderId"]; ok {
		builder.Where("`order_id`", val)
	}
	if val, ok := condition["orderNo"]; ok {
		builder.Where("`order_no`", val)
	}
	if val, ok := condition["channelOrderNo"]; ok {
		builder.Where("`channel_order_no`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["beginCreateTime"]; ok {
		builder.GreaterEqual("`create_time`", val)
	}
	if val, ok := condition["finishCreateTime"]; ok {
		builder.LessEqual("`create_time`", val)
	}

	if !util.Empty(condition["offset"]) {
		builder.Offset(cast.ToInt64(condition["offset"]))
	}
	if !util.Empty(condition["limit"]) {
		builder.Limit(cast.ToInt64(condition["limit"]))
	}
	builder.OrderBy("`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}

// PayRefundListTotal 查询列表数据总量
func PayRefundListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_refund`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["appId"]; ok {
		builder.Where("`app_id`", val)
	}
	if val, ok := condition["channelId"]; ok {
		builder.Where("`channel_id`", val)
	}
	if val, ok := condition["channelCode"]; ok {
		builder.Where("`channel_code`", val)
	}
	if val, ok := condition["merchantOrderId"]; ok {
		builder.Where("`merchant_order_id`", val)
	}
	if val, ok := condition["orderId"]; ok {
		builder.Where("`order_id`", val)
	}
	if val, ok := condition["orderNo"]; ok {
		builder.Where("`order_no`", val)
	}
	if val, ok := condition["channelOrderNo"]; ok {
		builder.Where("`channel_order_no`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["beginCreateTime"]; ok {
		builder.GreaterEqual("`create_time`", val)
	}
	if val, ok := condition["finishCreateTime"]; ok {
		builder.LessEqual("`create_time`", val)
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}
