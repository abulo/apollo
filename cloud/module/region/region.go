package region

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
)

// region 地区表
// RegionCreate 创建数据
func RegionCreate(ctx context.Context, data dao.Region) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`region`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// RegionUpdate 更新数据
func RegionUpdate(ctx context.Context, regionId int64, data dao.Region) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`region`").Where("`id`", regionId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// RegionDelete 删除数据
func RegionDelete(ctx context.Context, regionId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`region`").Where("`id`", regionId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// Region 查询单条数据
func Region(ctx context.Context, regionId int64) (res dao.Region, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`region`").Where("`id`", regionId).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// RegionRecover 恢复数据
func RegionRecover(ctx context.Context, regionId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`region`").Where("`id`", regionId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// RegionList 查询列表数据
func RegionList(ctx context.Context, condition map[string]any) (res []dao.Region, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`region`")
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("`name`", val)
	}
	if val, ok := condition["parentId"]; ok {
		builder.Where("`parent_id`", val)
	}

	builder.OrderBy("`parent_id`", sql.ASC)
	builder.OrderBy("`sort`", sql.ASC)
	builder.OrderBy("`id`", sql.ASC)
	// builder.Limit(20)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}
