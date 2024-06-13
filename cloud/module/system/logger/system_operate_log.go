package logger

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
)

// system_operate_log 操作日志
// SystemOperateLogCreate 创建数据
func SystemOperateLogCreate(ctx context.Context, data dao.SystemOperateLog) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_operate_log`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemOperateLogUpdate 更新数据
func SystemOperateLogUpdate(ctx context.Context, id int64, data dao.SystemOperateLog) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_operate_log`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemOperateLogDelete 删除数据
func SystemOperateLogDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_operate_log`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemOperateLog 查询单条数据
func SystemOperateLog(ctx context.Context, id int64) (res dao.SystemOperateLog, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_operate_log`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemOperateLogRecover 恢复数据
func SystemOperateLogRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_operate_log`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemOperateLogList 查询列表数据
func SystemOperateLogList(ctx context.Context, condition map[string]any) (res []dao.SystemOperateLog, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_operate_log`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("`username`", val)
	}
	if val, ok := condition["module"]; ok {
		builder.Where("`module`", val)
	}
	if val, ok := condition["beginStartTime"]; ok {
		builder.GreaterEqual("`start_time`", val)
	}
	if val, ok := condition["finishStartTime"]; ok {
		builder.LessEqual("`start_time`", val)
	}
	if val, ok := condition["result"]; ok {
		builder.Where("`result`", val)
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

// SystemOperateLogListTotal 查询列表数据总量
func SystemOperateLogListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_operate_log`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("`username`", val)
	}
	if val, ok := condition["module"]; ok {
		builder.Where("`module`", val)
	}
	if val, ok := condition["beginStartTime"]; ok {
		builder.GreaterEqual("`start_time`", val)
	}
	if val, ok := condition["finishStartTime"]; ok {
		builder.LessEqual("`start_time`", val)
	}
	if val, ok := condition["result"]; ok {
		builder.Where("`result`", val)
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}

func SystemOperateLogDrop(ctx context.Context, condition map[string]any) (res int64, err error) {
	data := make(map[string]any)
	data["deleted"] = 1
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	builder.Table("`system_operate_log`")
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
	if val, ok := condition["module"]; ok {
		builder.Where("`module`", val)
	}
	if val, ok := condition["beginStartTime"]; ok {
		builder.GreaterEqual("`start_time`", val)
	}
	if val, ok := condition["finishStartTime"]; ok {
		builder.LessEqual("`start_time`", val)
	}
	if val, ok := condition["result"]; ok {
		builder.Where("`result`", val)
	}
	query, args, err = builder.Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}
