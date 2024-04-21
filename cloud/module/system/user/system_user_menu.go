package user

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
)

// SystemUserMenuList 获取用户菜单
func SystemUserMenuList(ctx context.Context, systemUserId int64) (res []dao.SystemMenu, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Select("`system_menu`.*").Table("`system_user_role`").LeftJoin("`system_role_menu`", "`system_user_role`.role_id = `system_role_menu`.role_id").LeftJoin("`system_menu`", "`system_role_menu`.menu_id = `system_menu`.id ").Where("`system_user_role`.user_id", systemUserId).Where("`system_menu`.deleted", 0).Where("`system_menu`.status", 0).GroupBy("`system_menu`.id")
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
