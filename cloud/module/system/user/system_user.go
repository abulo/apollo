package user

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
)

// system_user 系统用户
// SystemUserCreate 创建数据
func SystemUserCreate(ctx context.Context, data dao.SystemUser) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_user`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemUserUpdate 更新数据
func SystemUserUpdate(ctx context.Context, systemUserId int64, data dao.SystemUser) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_user`").Where("`id`", systemUserId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemUserDelete 删除数据
func SystemUserDelete(ctx context.Context, systemUserId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_user`").Where("`id`", systemUserId).Delete()
	if err != nil {
		return
	}
	res, err = db.Delete(ctx, query, args...)
	return
}

// SystemUser 查询单条数据
func SystemUser(ctx context.Context, systemUserId int64) (res dao.SystemUser, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_user`").Where("`id`", systemUserId).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemUserList 查询列表数据
func SystemUserList(ctx context.Context, condition map[string]any) (res []dao.SystemUser, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_user`")
	if val, ok := condition["username"]; ok {
		builder.Where("`username`", val)
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

// SystemUserListTotal 查询列表数据总量
func SystemUserListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_user`")
	if val, ok := condition["username"]; ok {
		builder.Where("`username`", val)
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
