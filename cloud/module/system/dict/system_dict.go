package dict

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
)

// system_dict 字典数据表
// SystemDictCreate 创建数据
func SystemDictCreate(ctx context.Context, data dao.SystemDict) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dict`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemDictUpdate 更新数据
func SystemDictUpdate(ctx context.Context, id int64, data dao.SystemDict) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dict`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemDictDelete 删除数据
func SystemDictDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dict`").Where("`id`", id).Delete()
	if err != nil {
		return
	}
	res, err = db.Delete(ctx, query, args...)
	return
}

// SystemDict 查询单条数据
func SystemDict(ctx context.Context, id int64) (res dao.SystemDict, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dict`").Select("`system_dict`.*", "`system_dict_type`.`type` AS dict_type").LeftJoin("`system_dict_type`", "`system_dict`.`dict_type_id` = `system_dict_type`.`id`").Where("`system_dict`.`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemDictList 查询列表数据
func SystemDictList(ctx context.Context, condition map[string]any) (res []dao.SystemDict, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_dict`").Select("`system_dict`.*", "`system_dict_type`.`type` AS dict_type").LeftJoin("`system_dict_type`", "`system_dict`.`dict_type_id` = `system_dict_type`.`id`")
	if val, ok := condition["dictTypeId"]; ok {
		builder.Where("`system_dict`.`dict_type_id`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`system_dict`.`status`", val)
	}

	if val, ok := condition["pagination"]; ok {
		pagination := val.(*sql.Pagination)
		if pagination != nil {
			builder.Offset(pagination.GetOffset())
			builder.Limit(pagination.GetLimit())
		}
	}
	builder.OrderBy("`system_dict`.`sort`", sql.ASC)
	builder.OrderBy("`system_dict`.`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}

// SystemDictListTotal 查询列表数据总量
func SystemDictListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_dict`")
	if val, ok := condition["dictTypeId"]; ok {
		builder.Where("`dict_type_id`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}
