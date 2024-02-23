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

	builder.OrderBy("`parent_id`", sql.ASC)
	builder.OrderBy("`sort`", sql.ASC)
	builder.OrderBy("`id`", sql.ASC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}

// SystemMenuListRecursive 递归查询向上查询, 用于在计划任务中使用
func SystemMenuListRecursive(ctx context.Context, id int64) (res []dao.SystemMenu, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	query := "WITH RECURSIVE filter_system_menu (id,name,permission,type,sort,parent_id,path,icon,component,component_name,status,hide,link,keep_alive,affix,active_path,full_screen,redirect,creator,create_time,updater,update_time,deleted) AS ( SELECT * FROM system_menu WHERE id=? UNION ALL SELECT t.* FROM system_menu t INNER JOIN filter_system_menu ON filter_system_menu.parent_id=t.id) SELECT DISTINCT * FROM filter_system_menu ORDER BY filter_system_menu.parent_id ASC,filter_system_menu.sort ASC,filter_system_menu.id ASC"
	err = db.QueryRows(ctx, query, id).ToStruct(&res)
	return
}
