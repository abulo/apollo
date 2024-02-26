package dict

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
)

// system_dict_type 字典类型
// SystemDictTypeCreate 创建数据
func SystemDictTypeCreate(ctx context.Context, data dao.SystemDictType) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dict_type`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemDictTypeUpdate 更新数据
func SystemDictTypeUpdate(ctx context.Context, systemDictTypeId int64, data dao.SystemDictType) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dict_type`").Where("`id`", systemDictTypeId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemDictTypeDelete 删除数据
func SystemDictTypeDelete(ctx context.Context, systemDictTypeId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dict_type`").Where("`id`", systemDictTypeId).Delete()
	if err != nil {
		return
	}
	res, err = db.Delete(ctx, query, args...)
	return
}

// SystemDictType 查询单条数据
func SystemDictType(ctx context.Context, systemDictTypeId int64) (res dao.SystemDictType, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dict_type`").Where("`id`", systemDictTypeId).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemDictTypeList 查询列表数据
func SystemDictTypeList(ctx context.Context, condition map[string]any) (res []dao.SystemDictType, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_dict_type`")
	if val, ok := condition["type"]; ok {
		builder.Where("`type`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
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

// SystemDictTypeListTotal 查询列表数据总量
func SystemDictTypeListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_dict_type`")
	if val, ok := condition["type"]; ok {
		builder.Where("`type`", val)
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
