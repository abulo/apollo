package user

import (
	"cloud/dao"
	"cloud/initial"
	"context"
	"encoding/json"

	"github.com/abulo/ratel/v3/stores/sql"
	"google.golang.org/protobuf/proto"
)

// system_user_role 系统用户和系统角色关联表
// SystemUserRoleCreate 创建数据
func SystemUserRoleCreate(ctx context.Context, data dao.SystemUserRoleCustom) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var list []any
	if data.RoleIds.IsValid() {
		var roleIds []int64
		err = json.Unmarshal(data.RoleIds.JSON, &roleIds)
		if err != nil {
			return
		}
		for _, roleId := range roleIds {
			list = append(list, dao.SystemUserRole{
				RoleId:     proto.Int64(roleId),
				UserId:     data.UserId,
				Deleted:    proto.Int32(0),
				TenantId:   data.TenantId,
				Creator:    data.Creator,
				CreateTime: data.CreateTime,
				Updater:    data.Updater,
				UpdateTime: data.UpdateTime,
			})
		}
	}
	res = 0
	err = db.Transact(ctx, func(ctx context.Context, session sql.Session) error {
		// 先删除数据, 在添加数据
		// 需要先将数据删除了在添加
		builder := sql.NewBuilder()
		query, args, err := builder.Table("`system_user_role`").Where("`tenant_id`", data.TenantId).Where("user_id", data.UserId).Delete()
		if err != nil {
			return err
		}
		_, err = session.Delete(ctx, query, args...)
		if err != nil {
			return err
		}

		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_user_role`").MultiInsert(list...)
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

// SystemUserRoleList 查询列表数据
func SystemUserRoleList(ctx context.Context, condition map[string]any) (res []dao.SystemUserRole, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_user_role`")

	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("`user_id`", val)
	}
	if val, ok := condition["roleId"]; ok {
		builder.Where("`role_id`", val)
	}

	builder.OrderBy("`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}
