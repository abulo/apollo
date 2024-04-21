package role

import (
	"cloud/dao"
	"cloud/initial"
	"context"
	"encoding/json"

	"github.com/abulo/ratel/v3/stores/sql"
	"google.golang.org/protobuf/proto"
)

// system_role_menu 系统角色和系统菜单关联表
// SystemRoleMenuCreate 创建数据
func SystemRoleMenuCreate(ctx context.Context, data dao.SystemRoleMenuCustom) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var menuIds []int64
	err = json.Unmarshal(data.MenuIds.JSON, &menuIds)
	if err != nil {
		return
	}
	var list []any
	for _, menuId := range menuIds {
		list = append(list, dao.SystemRoleMenu{
			RoleId:     data.RoleId,
			MenuId:     proto.Int64(menuId),
			Deleted:    proto.Int32(0),
			TenantId:   data.TenantId,
			Creator:    data.Creator,
			CreateTime: data.CreateTime,
			Updater:    data.Updater,
			UpdateTime: data.UpdateTime,
		})
	}

	res = 0
	err = db.Transact(ctx, func(ctx context.Context, session sql.Session) error {
		if len(list) == 0 {
			return nil
		}
		// 需要先将数据删除了在添加
		delete := sql.NewBuilder()
		queryDelete, argsDelete, err := delete.Table("`system_role_menu`").Where("tenant_id", data.TenantId).Where("role_id", data.RoleId).Delete()
		if err != nil {
			return err
		}
		_, err = session.Delete(ctx, queryDelete, argsDelete...)
		if err != nil {
			return err
		}
		builder := sql.NewBuilder()
		query, args, err := builder.Table("`system_role_menu`").MultiInsert(list...)
		if err != nil {
			return err
		}
		_, err = session.MultiInsert(ctx, query, args...)
		if err != nil {
			return err
		}

		return nil
	})
	return res, err
}

// SystemRoleMenuList 查询列表数据
func SystemRoleMenuList(ctx context.Context, condition map[string]any) (res []dao.SystemRoleMenu, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_role_menu`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["roleId"]; ok {
		builder.Where("`role_id`", val)
	}
	if val, ok := condition["menuId"]; ok {
		builder.Where("`menu_id`", val)
	}
	builder.OrderBy("`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}
