package mail

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// system_mail_log 邮件日志表
// SystemMailLogCreate 创建数据
func SystemMailLogCreate(ctx context.Context, data dao.SystemMailLog) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_mail_log`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemMailLogUpdate 更新数据
func SystemMailLogUpdate(ctx context.Context, id int64, data dao.SystemMailLog) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_mail_log`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemMailLogDelete 删除数据
func SystemMailLogDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_mail_log`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemMailLog 查询单条数据
func SystemMailLog(ctx context.Context, id int64) (res dao.SystemMailLog, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_mail_log`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemMailLogRecover 恢复数据
func SystemMailLogRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_mail_log`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemMailLogDrop 清理数据
func SystemMailLogDrop(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_mail_log`").Where("`id`", id).Delete()
	if err != nil {
		return
	}
	res, err = db.Delete(ctx, query, args...)
	return
}

// SystemMailLogList 查询列表数据
func SystemMailLogList(ctx context.Context, condition map[string]any) (res []dao.SystemMailLog, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_mail_log`")
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["sendStatus"]; ok {
		builder.Where("`send_status`", val)
	}
	if val, ok := condition["beginSendTime"]; ok {
		builder.GreaterEqual("`send_time`", val)
	}
	if val, ok := condition["finishSendTime"]; ok {
		builder.LessEqual("`send_time`", val)
	}
	if val, ok := condition["templateTitle"]; ok {
		builder.Like("`template_title`", "%"+cast.ToString(val)+"%")
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("`user_id`", val)
	}
	if val, ok := condition["userType"]; ok {
		builder.Where("`user_type`", val)
	}
	if val, ok := condition["accountId"]; ok {
		builder.Where("`account_id`", val)
	}
	if val, ok := condition["templateCode"]; ok {
		builder.Where("`template_code`", val)
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

// SystemMailLogListTotal 查询列表数据总量
func SystemMailLogListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_mail_log`")
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["sendStatus"]; ok {
		builder.Where("`send_status`", val)
	}
	if val, ok := condition["beginSendTime"]; ok {
		builder.GreaterEqual("`send_time`", val)
	}
	if val, ok := condition["finishSendTime"]; ok {
		builder.LessEqual("`send_time`", val)
	}
	if val, ok := condition["templateTitle"]; ok {
		builder.Like("`template_title`", "%"+cast.ToString(val)+"%")
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("`user_id`", val)
	}
	if val, ok := condition["userType"]; ok {
		builder.Where("`user_type`", val)
	}
	if val, ok := condition["accountId"]; ok {
		builder.Where("`account_id`", val)
	}
	if val, ok := condition["templateCode"]; ok {
		builder.Where("`template_code`", val)
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}

// SystemMailLogMultipleDelete 批量删除数据
func SystemMailLogMultipleDelete(ctx context.Context, condition map[string]any) (res int64, err error) {
	if util.Empty(condition["ids"]) {
		return 0, errors.New("ids is empty")
	}
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	builder.Table("`system_mail_log`")
	if val, ok := condition["ids"]; ok {
		ids := val.([]int64)
		id := make([]any, 0)
		for _, v := range ids {
			id = append(id, v)
		}
		builder.In("`id`", id...)
	}
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemMailLogMultipleRecover 批量恢复数据
func SystemMailLogMultipleRecover(ctx context.Context, condition map[string]any) (res int64, err error) {
	if util.Empty(condition["ids"]) {
		return 0, errors.New("ids is empty")
	}
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	builder.Table("`system_mail_log`")
	if val, ok := condition["ids"]; ok {
		ids := val.([]int64)
		id := make([]any, 0)
		for _, v := range ids {
			id = append(id, v)
		}
		builder.In("`id`", id...)
	}
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemMailLogMultipleDrop 批量清理数据
func SystemMailLogMultipleDrop(ctx context.Context, condition map[string]any) (res int64, err error) {
	if util.Empty(condition["ids"]) {
		return 0, errors.New("ids is empty")
	}
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	builder.Table("`system_mail_log`")
	if val, ok := condition["ids"]; ok {
		ids := val.([]int64)
		id := make([]any, 0)
		for _, v := range ids {
			id = append(id, v)
		}
		builder.In("`id`", id...)
	}
	query, args, err := builder.Delete()
	if err != nil {
		return
	}
	res, err = db.Delete(ctx, query, args...)
	return
}
