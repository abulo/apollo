package wallet

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
)

// pay_wallet_transaction 会员钱包流水表
// PayWalletTransactionCreate 创建数据
func PayWalletTransactionCreate(ctx context.Context, data dao.PayWalletTransaction) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_wallet_transaction`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// PayWalletTransactionUpdate 更新数据
func PayWalletTransactionUpdate(ctx context.Context, id int64, data dao.PayWalletTransaction) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_wallet_transaction`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayWalletTransactionDelete 删除数据
func PayWalletTransactionDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`pay_wallet_transaction`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayWalletTransaction 查询单条数据
func PayWalletTransaction(ctx context.Context, id int64) (res dao.PayWalletTransaction, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_wallet_transaction`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// PayWalletTransactionRecover 恢复数据
func PayWalletTransactionRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`pay_wallet_transaction`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayWalletTransactionList 查询列表数据
func PayWalletTransactionList(ctx context.Context, condition map[string]any) (res []dao.PayWalletTransaction, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_wallet_transaction`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["walletId"]; ok {
		builder.Where("`wallet_id`", val)
	}
	if val, ok := condition["bizType"]; ok {
		builder.Where("`biz_type`", val)
	}
	if val, ok := condition["bizId"]; ok {
		builder.Where("`biz_id`", val)
	}
	if val, ok := condition["no"]; ok {
		builder.Where("`no`", val)
	}
	if val, ok := condition["title"]; ok {
		builder.Like("`title`", "%"+cast.ToString(val)+"%")
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

// PayWalletTransactionListTotal 查询列表数据总量
func PayWalletTransactionListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_wallet_transaction`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["walletId"]; ok {
		builder.Where("`wallet_id`", val)
	}
	if val, ok := condition["bizType"]; ok {
		builder.Where("`biz_type`", val)
	}
	if val, ok := condition["bizId"]; ok {
		builder.Where("`biz_id`", val)
	}
	if val, ok := condition["no"]; ok {
		builder.Where("`no`", val)
	}
	if val, ok := condition["title"]; ok {
		builder.Like("`title`", "%"+cast.ToString(val)+"%")
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}
