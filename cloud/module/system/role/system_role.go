package role

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
)

// system_role 系统角色
// SystemRoleCreate 创建数据
func SystemRoleCreate(ctx context.Context, data dao.SystemRole) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_role`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemRoleUpdate 更新数据
func SystemRoleUpdate(ctx context.Context, systemRoleId int64, data dao.SystemRole) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_role`").Where("`id`", systemRoleId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemRoleDelete 删除数据
func SystemRoleDelete(ctx context.Context, systemRoleId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_role`").Where("`id`", systemRoleId).Delete()
	if err != nil {
		return
	}
	res, err = db.Delete(ctx, query, args...)
	return
}

// SystemRole 查询单条数据
func SystemRole(ctx context.Context, systemRoleId int64) (res dao.SystemRole, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_role`").Where("`id`", systemRoleId).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemRoleList 查询列表数据
func SystemRoleList(ctx context.Context, condition map[string]any) (res []dao.SystemRole, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_role`")
	if val, ok := condition["type"]; ok {
		builder.Where("`type`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}

	builder.OrderBy("`sort`", sql.ASC)
	builder.OrderBy("`id`", sql.ASC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}
