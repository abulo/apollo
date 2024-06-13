package notify

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
)

// system_notify_message 站内信消息表
// SystemNotifyMessageCreate 创建数据
func SystemNotifyMessageCreate(ctx context.Context, data dao.SystemNotifyMessage) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_notify_message`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemNotifyMessageUpdate 更新数据
func SystemNotifyMessageUpdate(ctx context.Context, id int64, data dao.SystemNotifyMessage) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_notify_message`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemNotifyMessageDelete 删除数据
func SystemNotifyMessageDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_notify_message`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemNotifyMessage 查询单条数据
func SystemNotifyMessage(ctx context.Context, id int64) (res dao.SystemNotifyMessage, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_notify_message`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemNotifyMessageRecover 恢复数据
func SystemNotifyMessageRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_notify_message`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemNotifyMessageList 查询列表数据
func SystemNotifyMessageList(ctx context.Context, condition map[string]any) (res []dao.SystemNotifyMessage, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_notify_message`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["templateType"]; ok {
		builder.Where("`template_type`", val)
	}
	if val, ok := condition["readStatus"]; ok {
		builder.Where("`read_status`", val)
	}
	if val, ok := condition["beginReadTime"]; ok {
		builder.GreaterEqual("`read_time`", val)
	}
	if val, ok := condition["finishReadTime"]; ok {
		builder.LessEqual("`read_time`", val)
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

// SystemNotifyMessageListTotal 查询列表数据总量
func SystemNotifyMessageListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_notify_message`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["templateType"]; ok {
		builder.Where("`template_type`", val)
	}
	if val, ok := condition["readStatus"]; ok {
		builder.Where("`read_status`", val)
	}
	if val, ok := condition["beginReadTime"]; ok {
		builder.GreaterEqual("`read_time`", val)
	}
	if val, ok := condition["finishReadTime"]; ok {
		builder.LessEqual("`read_time`", val)
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}
