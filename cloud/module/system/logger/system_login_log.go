package logger

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
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

// SystemLoginLogDelete 删除数据
func SystemLoginLogDelete(ctx context.Context, systemLoginLogIds []int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	id := make([]any, 0)
	for _, v := range systemLoginLogIds {
		id = append(id, v)
	}
	query, args, err := builder.Table("`system_login_log`").In("`id`", id...).Delete()
	if err != nil {
		return
	}
	res, err = db.Delete(ctx, query, args...)
	return
}

// SystemLoginLog 查询单条数据
func SystemLoginLog(ctx context.Context, systemLoginLogId int64) (res dao.SystemLoginLog, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_login_log`").Where("`id`", systemLoginLogId).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemLoginLogList 查询列表数据
func SystemLoginLogList(ctx context.Context, condition map[string]any) (res []dao.SystemLoginLog, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_login_log`")
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

// SystemLoginLogListTotal 查询列表数据总量
func SystemLoginLogListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_login_log`")
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

// SystemLoginLogDrop 查询列表数据总量
func SystemLoginLogDrop(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_login_log`")
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
	query, args, err := builder.Delete()
	if err != nil {
		return
	}
	res, err = db.Delete(ctx, query, args...)
	return
}
