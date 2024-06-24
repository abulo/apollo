package wallet

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
)

// pay_wallet_recharge 会员钱包充值
// PayWalletRechargeCreate 创建数据
func PayWalletRechargeCreate(ctx context.Context, data dao.PayWalletRecharge) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_wallet_recharge`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// PayWalletRechargeUpdate 更新数据
func PayWalletRechargeUpdate(ctx context.Context, id int64, data dao.PayWalletRecharge) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_wallet_recharge`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayWalletRechargeDelete 删除数据
func PayWalletRechargeDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`pay_wallet_recharge`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayWalletRecharge 查询单条数据
func PayWalletRecharge(ctx context.Context, id int64) (res dao.PayWalletRecharge, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_wallet_recharge`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// PayWalletRechargeRecover 恢复数据
func PayWalletRechargeRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`pay_wallet_recharge`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayWalletRechargeList 查询列表数据
func PayWalletRechargeList(ctx context.Context, condition map[string]any) (res []dao.PayWalletRecharge, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_wallet_recharge`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["walletId"]; ok {
		builder.Where("`wallet_id`", val)
	}
	if val, ok := condition["packageId"]; ok {
		builder.Where("`package_id`", val)
	}
	if val, ok := condition["payStatus"]; ok {
		builder.Where("`pay_status`", val)
	}
	if val, ok := condition["payOrderId"]; ok {
		builder.Where("`pay_order_id`", val)
	}
	if val, ok := condition["payChannelCode"]; ok {
		builder.Where("`pay_channel_code`", val)
	}
	if val, ok := condition["beginPayTime"]; ok {
		builder.GreaterEqual("`pay_time`", val)
	}
	if val, ok := condition["finishPayTime"]; ok {
		builder.LessEqual("`pay_time`", val)
	}
	if val, ok := condition["refundStatus"]; ok {
		builder.Where("`refund_status`", val)
	}
	if val, ok := condition["payRefundId"]; ok {
		builder.Where("`pay_refund_id`", val)
	}
	if val, ok := condition["beginRefundTime"]; ok {
		builder.GreaterEqual("`refund_time`", val)
	}
	if val, ok := condition["finishRefundTime"]; ok {
		builder.LessEqual("`refund_time`", val)
	}

	if val, ok := condition["pagination"]; ok {
		pagination := val.(*sql.Pagination)
		if pagination != nil {
			builder.Offset(pagination.GetOffset())
			builder.Limit(pagination.GetLimit())
		}
	}
	builder.OrderBy("`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}

// PayWalletRechargeListTotal 查询列表数据总量
func PayWalletRechargeListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_wallet_recharge`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["walletId"]; ok {
		builder.Where("`wallet_id`", val)
	}
	if val, ok := condition["packageId"]; ok {
		builder.Where("`package_id`", val)
	}
	if val, ok := condition["payStatus"]; ok {
		builder.Where("`pay_status`", val)
	}
	if val, ok := condition["payOrderId"]; ok {
		builder.Where("`pay_order_id`", val)
	}
	if val, ok := condition["payChannelCode"]; ok {
		builder.Where("`pay_channel_code`", val)
	}
	if val, ok := condition["beginPayTime"]; ok {
		builder.GreaterEqual("`pay_time`", val)
	}
	if val, ok := condition["finishPayTime"]; ok {
		builder.LessEqual("`pay_time`", val)
	}
	if val, ok := condition["refundStatus"]; ok {
		builder.Where("`refund_status`", val)
	}
	if val, ok := condition["payRefundId"]; ok {
		builder.Where("`pay_refund_id`", val)
	}
	if val, ok := condition["beginRefundTime"]; ok {
		builder.GreaterEqual("`refund_time`", val)
	}
	if val, ok := condition["finishRefundTime"]; ok {
		builder.LessEqual("`refund_time`", val)
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}
