package user

import (
	"cloud/dao"
	"cloud/initial"
	"context"
	"encoding/json"

	"github.com/abulo/ratel/v3/stores/sql"
	"google.golang.org/protobuf/proto"
)

// system_user_dept 系统用户部门
// SystemUserDeptCreate 创建数据
func SystemUserDeptCreate(ctx context.Context, data dao.SystemUserDeptCustom) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var list []any
	if data.DeptIds.IsValid() {
		var deptIds []int64
		err = json.Unmarshal(data.DeptIds.JSON, &deptIds)
		if err != nil {
			return
		}
		for _, deptId := range deptIds {
			list = append(list, dao.SystemUserDept{
				DeptId:     proto.Int64(deptId),
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
		query, args, err := builder.Table("`system_user_dept`").Where("`tenant_id`", data.TenantId).Where("user_id", data.UserId).Delete()
		if err != nil {
			return err
		}
		_, err = session.Delete(ctx, query, args...)
		if err != nil {
			return err
		}

		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_user_dept`").MultiInsert(list...)
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

// SystemUserDeptList 查询列表数据
func SystemUserDeptList(ctx context.Context, condition map[string]any) (res []dao.SystemUserDept, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_user_dept`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["deptId"]; ok {
		builder.Where("`dept_id`", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("`user_id`", val)
	}

	builder.OrderBy("`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}

// SystemDeptListRecursive 递归查询
func SystemDeptListRecursive(ctx context.Context, id int64) (res []dao.SystemDept, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	query := "WITH RECURSIVE filter_system_dept AS (SELECT `id`,`name`,`parent_id`,`sort`,`system_user_id`,`phone`,`email`,`status`,`deleted`,`tenant_id`,`creator`,`create_time`,`updater`,`update_time` FROM `system_dept` WHERE `id`=? UNION ALL SELECT d.* FROM system_dept AS d INNER JOIN filter_system_dept f ON d.parent_id=f.id) SELECT filter_system_dept.* FROM filter_system_dept"
	err = db.QueryRows(ctx, query, id).ToStruct(&res)
	return
}
