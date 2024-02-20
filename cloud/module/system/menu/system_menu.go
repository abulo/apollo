package menu

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
)

// system_menu 系统菜单
// SystemMenuCreate 创建数据
func SystemMenuCreate(ctx context.Context, data dao.SystemMenu) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_menu`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemMenuUpdate 更新数据
func SystemMenuUpdate(ctx context.Context, systemMenuId int64, data dao.SystemMenu) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_menu`").Where("`id`", systemMenuId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemMenuDelete 删除数据
func SystemMenuDelete(ctx context.Context, systemMenuId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_menu`").Where("`id`", systemMenuId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemMenu 查询单条数据
func SystemMenu(ctx context.Context, systemMenuId int64) (res dao.SystemMenu, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_menu`").Where("`id`", systemMenuId).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemMenuRecover 恢复数据
func SystemMenuRecover(ctx context.Context, systemMenuId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_menu`").Where("`id`", systemMenuId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemMenuList 查询列表数据
func SystemMenuList(ctx context.Context, condition map[string]any) (res []dao.SystemMenu, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_menu`")
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["type"]; ok {
		builder.Where("`type`", val)
	}
	if val, ok := condition["sort"]; ok {
		builder.Where("`sort`", val)
	}
	if val, ok := condition["parentId"]; ok {
		builder.Where("`parent_id`", val)
	}

	builder.OrderBy("`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}
