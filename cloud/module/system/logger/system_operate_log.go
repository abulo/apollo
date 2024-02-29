package logger

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
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

// SystemOperateLogDelete 删除数据
func SystemOperateLogDelete(ctx context.Context, systemOperateLogIds []int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	id := make([]any, 0)
	for _, v := range systemOperateLogIds {
		id = append(id, v)
	}
	query, args, err := builder.Table("`system_operate_log`").In("`id`", id...).Delete()
	if err != nil {
		return
	}
	res, err = db.Delete(ctx, query, args...)
	return
}

// SystemOperateLog 查询单条数据
func SystemOperateLog(ctx context.Context, systemOperateLogId int64) (res dao.SystemOperateLog, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_operate_log`").Where("`id`", systemOperateLogId).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemOperateLogList 查询列表数据
func SystemOperateLogList(ctx context.Context, condition map[string]any) (res []dao.SystemOperateLog, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_operate_log`")
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

// SystemOperateLogListTotal 查询列表数据总量
func SystemOperateLogListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_operate_log`")
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

// SystemOperateLogDrop 清理数据
func SystemOperateLogDrop(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_operate_log`")
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

	query, args, err := builder.Delete()
	if err != nil {
		return
	}
	res, err = db.Delete(ctx, query, args...)
	return
}
