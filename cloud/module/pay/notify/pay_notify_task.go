package notify

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
)

// pay_notify_task 商户支付-退款等的通知
// PayNotifyTaskCreate 创建数据
func PayNotifyTaskCreate(ctx context.Context, data dao.PayNotifyTask) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_notify_task`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// PayNotifyTaskUpdate 更新数据
func PayNotifyTaskUpdate(ctx context.Context, id int64, data dao.PayNotifyTask) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_notify_task`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayNotifyTaskDelete 删除数据
func PayNotifyTaskDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`pay_notify_task`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayNotifyTask 查询单条数据
func PayNotifyTask(ctx context.Context, id int64) (res dao.PayNotifyTask, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`pay_notify_task`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// PayNotifyTaskRecover 恢复数据
func PayNotifyTaskRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`pay_notify_task`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// PayNotifyTaskList 查询列表数据
func PayNotifyTaskList(ctx context.Context, condition map[string]any) (res []dao.PayNotifyTask, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_notify_task`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["appId"]; ok {
		builder.Where("`app_id`", val)
	}
	if val, ok := condition["type"]; ok {
		builder.Where("`type`", val)
	}
	if val, ok := condition["dataId"]; ok {
		builder.Where("`data_id`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["merchantOrderId"]; ok {
		builder.Where("`merchant_order_id`", val)
	}
	if val, ok := condition["beginCreateTime"]; ok {
		builder.GreaterEqual("`create_time`", val)
	}
	if val, ok := condition["finishCreateTime"]; ok {
		builder.LessEqual("`create_time`", val)
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

// PayNotifyTaskListTotal 查询列表数据总量
func PayNotifyTaskListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`pay_notify_task`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["appId"]; ok {
		builder.Where("`app_id`", val)
	}
	if val, ok := condition["type"]; ok {
		builder.Where("`type`", val)
	}
	if val, ok := condition["dataId"]; ok {
		builder.Where("`data_id`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["merchantOrderId"]; ok {
		builder.Where("`merchant_order_id`", val)
	}
	if val, ok := condition["beginCreateTime"]; ok {
		builder.GreaterEqual("`create_time`", val)
	}
	if val, ok := condition["finishCreateTime"]; ok {
		builder.LessEqual("`create_time`", val)
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}
