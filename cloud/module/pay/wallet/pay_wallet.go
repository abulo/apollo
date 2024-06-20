package wallet

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// pay_wallet 会员钱包表
// PayWalletCreate 创建数据
func PayWalletCreate(ctx context.Context, data dao.PayWallet) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_wallet`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// PayWalletUpdate 更新数据
func PayWalletUpdate(ctx context.Context, id int64, data dao.PayWallet) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_wallet`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayWalletDelete 删除数据
func PayWalletDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`pay_wallet`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayWallet 查询单条数据
func PayWallet(ctx context.Context, id int64) (res dao.PayWallet, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_wallet`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// PayWalletRecover 恢复数据
func PayWalletRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`pay_wallet`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayWalletUser 查询单条数据
func PayWalletUser(ctx context.Context, condition map[string]any) (res dao.PayWallet, err error) {
	if util.Empty(condition) {
		err = errors.New("condition is empty")
		return
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_wallet`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["userType"]; ok {
		builder.Where("`user_type`", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("`user_id`", val)
	}

	query, args, err := builder.Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// PayWalletList 查询列表数据
func PayWalletList(ctx context.Context, condition map[string]any) (res []dao.PayWalletCustom, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_wallet`").Select("`pay_wallet`.*", "CASE WHEN pay_wallet.user_type = 1 THEN member.nickname WHEN pay_wallet.user_type = 0 THEN system_user.nickname  END AS username")
	builder.LeftJoin("system_user", "pay_wallet.user_id = system_user.id AND pay_wallet.user_type = 0")
	builder.LeftJoin("member", "pay_wallet.user_id = member.id AND pay_wallet.user_type = 1")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`pay_wallet`.`tenant_id`", val)
	}
	if val, ok := condition["userType"]; ok {
		builder.Where("`pay_wallet`.`user_type`", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("`pay_wallet`.`user_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`pay_wallet`.`deleted`", val)
	}
	if val, ok := condition["username"]; ok {
		builder.And()
		builder.LeftBracket()
		builder.FirstLike("`system_user`.`username`", "%"+cast.ToString(val)+"%")
		builder.OrLike("`system_user`.`nickname`", "%"+cast.ToString(val)+"%")
		builder.OrLike("`system_user`.`mobile`", "%"+cast.ToString(val)+"%")
		builder.OrLike("`member`.`nickname`", "%"+cast.ToString(val)+"%")
		builder.RightBracket()
	}

	if val, ok := condition["pagination"]; ok {
		pagination := val.(*sql.Pagination)
		if pagination != nil {
			builder.Offset(pagination.GetOffset())
			builder.Limit(pagination.GetLimit())
		}
	}
	builder.OrderBy("`pay_wallet`.`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}

// PayWalletListTotal 查询列表数据总量
func PayWalletListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_wallet`")
	builder.LeftJoin("system_user", "pay_wallet.user_id = system_user.id AND pay_wallet.user_type = 0")
	builder.LeftJoin("member", "pay_wallet.user_id = member.id AND pay_wallet.user_type = 1")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`pay_wallet`.`tenant_id`", val)
	}
	if val, ok := condition["userType"]; ok {
		builder.Where("`pay_wallet`.`user_type`", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("`pay_wallet`.`user_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`pay_wallet`.`deleted`", val)
	}

	if val, ok := condition["username"]; ok {
		builder.And()
		builder.LeftBracket()
		builder.FirstLike("`system_user`.`username`", "%"+cast.ToString(val)+"%")
		builder.OrLike("`system_user`.`nickname`", "%"+cast.ToString(val)+"%")
		builder.OrLike("`system_user`.`mobile`", "%"+cast.ToString(val)+"%")
		builder.OrLike("`member`.`nickname`", "%"+cast.ToString(val)+"%")
		builder.RightBracket()
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}
