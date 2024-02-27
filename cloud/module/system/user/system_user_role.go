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
	var roleIds []int64
	err = json.Unmarshal(data.SystemRoleIds.JSON, &roleIds)
	if err != nil {
		return
	}
	var list []any
	for _, roleId := range roleIds {
		list = append(list, dao.SystemUserRole{
			SystemRoleId: proto.Int64(roleId),
			SystemUserId: data.SystemUserId,
			Creator:      data.Creator,
			CreateTime:   data.CreateTime,
			Updater:      data.Updater,
			UpdateTime:   data.UpdateTime,
		})
	}
	res = 0
	err = db.Transact(ctx, func(ctx context.Context, session sql.Session) error {
		if len(list) == 0 {
			return nil
		}
		// 需要先将数据删除了在添加
		delete := sql.NewBuilder()
		queryDelete, argsDelete, err := delete.Table("`system_user_role`").Where("system_user_id", data.SystemUserId).Delete()
		if err != nil {
			return err
		}
		_, err = session.Delete(ctx, queryDelete, argsDelete...)
		if err != nil {
			return err
		}
		builder := sql.NewBuilder()
		query, args, err := builder.Table("`system_user_role`").MultiInsert(list...)
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

// SystemUserRoleList 查询列表数据
func SystemUserRoleList(ctx context.Context, condition map[string]any) (res []dao.SystemUserRole, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_user_role`")
	if val, ok := condition["systemRoleId"]; ok {
		builder.Where("`system_role_id`", val)
	}
	if val, ok := condition["systemUserId"]; ok {
		builder.Where("`system_user_id`", val)
	}

	builder.OrderBy("`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}
