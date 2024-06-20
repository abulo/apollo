package wallet

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
)

// pay_wallet_recharge_package 充值套餐表
// PayWalletRechargePackageCreate 创建数据
func PayWalletRechargePackageCreate(ctx context.Context, data dao.PayWalletRechargePackage) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_wallet_recharge_package`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// PayWalletRechargePackageUpdate 更新数据
func PayWalletRechargePackageUpdate(ctx context.Context, id int64, data dao.PayWalletRechargePackage) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_wallet_recharge_package`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayWalletRechargePackageDelete 删除数据
func PayWalletRechargePackageDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`pay_wallet_recharge_package`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayWalletRechargePackage 查询单条数据
func PayWalletRechargePackage(ctx context.Context, id int64) (res dao.PayWalletRechargePackage, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_wallet_recharge_package`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// PayWalletRechargePackageRecover 恢复数据
func PayWalletRechargePackageRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`pay_wallet_recharge_package`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayWalletRechargePackageList 查询列表数据
func PayWalletRechargePackageList(ctx context.Context, condition map[string]any) (res []dao.PayWalletRechargePackage, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_wallet_recharge_package`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("`name`", val)
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

// PayWalletRechargePackageListTotal 查询列表数据总量
func PayWalletRechargePackageListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_wallet_recharge_package`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("`name`", val)
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}
