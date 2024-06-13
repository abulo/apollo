package logger

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
)

// system_login_log 登录日志
// SystemLoginLogCreate 创建数据
func SystemLoginLogCreate(ctx context.Context, data dao.SystemLoginLog) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_login_log`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemLoginLogUpdate 更新数据
func SystemLoginLogUpdate(ctx context.Context, id int64, data dao.SystemLoginLog) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_login_log`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemLoginLogDelete 删除数据
func SystemLoginLogDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_login_log`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemLoginLog 查询单条数据
func SystemLoginLog(ctx context.Context, id int64) (res dao.SystemLoginLog, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_login_log`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemLoginLogRecover 恢复数据
func SystemLoginLogRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_login_log`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemLoginLogList 查询列表数据
func SystemLoginLogList(ctx context.Context, condition map[string]any) (res []dao.SystemLoginLog, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_login_log`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("`username`", val)
	}
	if val, ok := condition["beginLoginTime"]; ok {
		builder.GreaterEqual("`login_time`", val)
	}
	if val, ok := condition["finishLoginTime"]; ok {
		builder.LessEqual("`login_time`", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("`channel`", val)
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

// SystemLoginLogListTotal 查询列表数据总量
func SystemLoginLogListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_login_log`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("`username`", val)
	}
	if val, ok := condition["beginLoginTime"]; ok {
		builder.GreaterEqual("`login_time`", val)
	}
	if val, ok := condition["finishLoginTime"]; ok {
		builder.LessEqual("`login_time`", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("`channel`", val)
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}

func SystemLoginLogDrop(ctx context.Context, condition map[string]any) (res int64, err error) {
	data := make(map[string]any)
	data["deleted"] = 1
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	builder.Table("`system_login_log`")
	var query string
	var args []any
	if val, ok := condition["ids"]; ok {
		ids := val.([]int64)
		id := make([]any, 0)
		for _, v := range ids {
			id = append(id, v)
		}
		builder.In("`id`", id...)
	}
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("`username`", val)
	}
	if val, ok := condition["beginLoginTime"]; ok {
		builder.GreaterEqual("`login_time`", val)
	}
	if val, ok := condition["finishLoginTime"]; ok {
		builder.LessEqual("`login_time`", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("`channel`", val)
	}
	query, args, err = builder.Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}
