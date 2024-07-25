package notify

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
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

// SystemNotifyMessageDrop 清理数据
func SystemNotifyMessageDrop(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_notify_message`").Where("`id`", id).Delete()
	if err != nil {
		return
	}
	res, err = db.Delete(ctx, query, args...)
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

// SystemNotifyMessageMultipleDelete 批量删除数据
func SystemNotifyMessageMultipleDelete(ctx context.Context, condition map[string]any) (res int64, err error) {
	if util.Empty(condition["ids"]) || util.Empty(condition["tenantId"]) {
		return 0, errors.New("ids or tenantId is empty")
	}
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	builder.Table("`system_notify_message`")
	if val, ok := condition["ids"]; ok {
		ids := val.([]int64)
		id := make([]any, 0)
		for _, v := range ids {
			id = append(id, v)
		}
		builder.In("`id`", id...)
	}
	tenantId := cast.ToInt64(condition["tenantId"])
	builder.Where("tenant_id", tenantId)
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemNotifyMessageMultipleRecover 批量恢复数据
func SystemNotifyMessageMultipleRecover(ctx context.Context, condition map[string]any) (res int64, err error) {
	if util.Empty(condition["ids"]) || util.Empty(condition["tenantId"]) {
		return 0, errors.New("ids or tenantId is empty")
	}
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	builder.Table("`system_notify_message`")
	if val, ok := condition["ids"]; ok {
		ids := val.([]int64)
		id := make([]any, 0)
		for _, v := range ids {
			id = append(id, v)
		}
		builder.In("`id`", id...)
	}
	tenantId := cast.ToInt64(condition["tenantId"])
	builder.Where("tenant_id", tenantId)
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemNotifyMessageMultipleDrop 批量清理数据
func SystemNotifyMessageMultipleDrop(ctx context.Context, condition map[string]any) (res int64, err error) {
	if util.Empty(condition["ids"]) || util.Empty(condition["tenantId"]) {
		return 0, errors.New("ids or tenantId is empty")
	}
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	builder.Table("`system_notify_message`")
	if val, ok := condition["ids"]; ok {
		ids := val.([]int64)
		id := make([]any, 0)
		for _, v := range ids {
			id = append(id, v)
		}
		builder.In("`id`", id...)
	}
	tenantId := cast.ToInt64(condition["tenantId"])
	builder.Where("tenant_id", tenantId)
	query, args, err := builder.Delete()
	if err != nil {
		return
	}
	res, err = db.Delete(ctx, query, args...)
	return
}
