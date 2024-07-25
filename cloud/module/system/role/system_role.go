package role

import (
	"cloud/dao"
	"cloud/initial"
	"context"
	"encoding/json"

	"github.com/abulo/ratel/v3/stores/sql"
	"google.golang.org/protobuf/proto"
)

// system_role 系统角色
// SystemRoleCreate 创建数据
func SystemRoleCreate(ctx context.Context, data dao.SystemRole) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	res = 0
	err = db.Transact(ctx, func(ctx context.Context, session sql.Session) error {
		builder := sql.NewBuilder()
		query, args, err := builder.Table("`system_role`").Insert(data)
		if err != nil {
			return err
		}
		res, err = session.Insert(ctx, query, args...)
		if err != nil {
			return err
		}
		// 该去插入角色和菜单的关联关系
		var menuIds []int64
		err = json.Unmarshal(data.MenuIds.JSON, &menuIds)
		if err != nil {
			return err
		}
		var list []any
		for _, menuId := range menuIds {
			list = append(list, dao.SystemRoleMenu{
				RoleId:     proto.Int64(res),
				MenuId:     proto.Int64(menuId),
				Deleted:    proto.Int32(0),
				TenantId:   data.TenantId,
				Creator:    data.Creator,
				CreateTime: data.CreateTime,
				Updater:    data.Updater,
				UpdateTime: data.UpdateTime,
			})
		}
		if len(list) == 0 {
			return nil
		}
		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_role_menu`").MultiInsert(list...)
		if err != nil {
			return err
		}
		_, err = session.MultiInsert(ctx, query, args...)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// SystemRoleUpdate 更新数据
func SystemRoleUpdate(ctx context.Context, id int64, data dao.SystemRole) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	res = 0
	err = db.Transact(ctx, func(ctx context.Context, session sql.Session) error {
		builder := sql.NewBuilder()
		query, args, err := builder.Table("`system_role`").Where("`id`", id).Update(data)
		if err != nil {
			return err
		}
		res, err = session.Update(ctx, query, args...)
		if err != nil {
			return err
		}

		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_role_menu`").Where("tenant_id", data.TenantId).Where("role_id", data.Id).Delete()
		if err != nil {
			return err
		}
		_, err = session.Delete(ctx, query, args...)
		if err != nil {
			return err
		}
		// 该去插入角色和菜单的关联关系
		var menuIds []int64
		err = json.Unmarshal(data.MenuIds.JSON, &menuIds)
		if err != nil {
			return err
		}
		var list []any
		for _, menuId := range menuIds {
			list = append(list, dao.SystemRoleMenu{
				RoleId:     proto.Int64(res),
				MenuId:     proto.Int64(menuId),
				Deleted:    proto.Int32(0),
				TenantId:   data.TenantId,
				Creator:    data.Creator,
				CreateTime: data.CreateTime,
				Updater:    data.Updater,
				UpdateTime: data.UpdateTime,
			})
		}
		if len(list) == 0 {
			return nil
		}
		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_role_menu`").MultiInsert(list...)
		if err != nil {
			return err
		}
		_, err = session.MultiInsert(ctx, query, args...)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// SystemRoleDelete 删除数据
func SystemRoleDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_role`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemRole 查询单条数据
func SystemRole(ctx context.Context, id int64) (res dao.SystemRole, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_role`").Select("system_role.*", "JSON_ARRAYAGG(system_role_menu.menu_id) AS menu_ids")
	builder.LeftJoin("system_role_menu", "system_role_menu.tenant_id =system_role.tenant_id  AND  system_role_menu.role_id = system_role.id")
	builder.LeftJoin("system_menu", "system_menu.id = system_role_menu.menu_id  AND  system_menu.deleted =0 AND system_menu.status =0")
	query, args, err := builder.Where("`system_role`.`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemRoleRecover 恢复数据
func SystemRoleRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_role`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemRoleDrop 清理数据
func SystemRoleDrop(ctx context.Context, id int64) (res int64, err error) {
	roleItem, err := SystemRole(ctx, id)
	if err != nil {
		return 0, err
	}
	db := initial.Core.Store.LoadSQL("mysql").Write()
	res = 0
	err = db.Transact(ctx, func(ctx context.Context, session sql.Session) error {

		builder := sql.NewBuilder()
		query, args, err := builder.Table("`system_role_menu`").Where("`tenant_id`", roleItem.TenantId).Where("role_id", id).Delete()
		if err != nil {
			return err
		}
		_, err = session.Delete(ctx, query, args...)
		if err != nil {
			return err
		}

		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_user_role`").Where("`tenant_id`", roleItem.TenantId).Where("role_id", id).Delete()
		if err != nil {
			return err
		}
		_, err = session.Delete(ctx, query, args...)
		if err != nil {
			return err
		}

		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_role`").Where("`id`", id).Delete()
		if err != nil {
			return err
		}
		res, err = session.Delete(ctx, query, args...)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// SystemRoleList 查询列表数据
func SystemRoleList(ctx context.Context, condition map[string]any) (res []dao.SystemRole, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_role`").Select("system_role.*", "JSON_ARRAYAGG(system_role_menu.menu_id) AS menu_ids")
	builder.LeftJoin("system_role_menu", "system_role_menu.tenant_id =system_role.tenant_id  AND  system_role_menu.role_id = system_role.id")
	builder.LeftJoin("system_menu", "system_menu.id = system_role_menu.menu_id  AND  system_menu.deleted =0 AND system_menu.status =0")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`system_role`.`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`system_role`.`deleted`", val)
	}
	if val, ok := condition["type"]; ok {
		builder.Where("`system_role`.`type`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`system_role`.`status`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("`system_role`.`name`", val)
	}

	if val, ok := condition["pagination"]; ok {
		pagination := val.(*sql.Pagination)
		if pagination != nil {
			builder.Offset(pagination.GetOffset())
			builder.Limit(pagination.GetLimit())
		}
	}
	builder.GroupBy("`system_role`.`id`")
	builder.OrderBy("`system_role`.`sort`", sql.ASC)
	builder.OrderBy("`system_role`.`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}

// SystemRoleListTotal 查询列表数据总量
func SystemRoleListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_role`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["type"]; ok {
		builder.Where("`type`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("`name`", val)
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}
