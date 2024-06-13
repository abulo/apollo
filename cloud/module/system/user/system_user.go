package user

import (
	"cloud/dao"
	"cloud/initial"
	"context"
	"strconv"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
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
func SystemUserUpdate(ctx context.Context, id int64, data dao.SystemUser) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_user`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemUserDelete 删除数据
func SystemUserDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_user`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemUser 查询单条数据
func SystemUser(ctx context.Context, id int64) (res dao.SystemUser, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_user`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemUserRecover 恢复数据
func SystemUserRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_user`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemUserLogin 查询单条数据
func SystemUserLogin(ctx context.Context, condition map[string]any) (res dao.SystemUser, err error) {
	if util.Empty(condition) {
		err = errors.New("condition is empty")
		return
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_user`")
	if val, ok := condition["username"]; ok {
		builder.Where("`username`", val)
	}

	query, args, err := builder.Row()
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
	builder.Select("`system_user`.*")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`system_user`.`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`system_user`.`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`system_user`.`status`", val)
	}
	if val, ok := condition["systemDeptId"]; ok {
		systemDeptId := cast.ToInt64(val)
		if systemDeptId > 0 {
			if deptList, err := SystemDeptListRecursive(ctx, systemDeptId); err == nil {
				newVal := make([]any, 0)
				for _, v := range deptList {
					newVal = append(newVal, cast.ToInt64(v.Id))
				}
				builder.LeftJoin("`system_user_dept`", "`system_user`.`id` = `system_user_dept`.`user_id`")
				builder.In("`system_user_dept`.`dept_id`", newVal...)
			}
		}
	}
	if val, ok := condition["username"]; ok {
		builder.And()
		builder.LeftBracket()
		builder.FirstLike("`system_user`.`username`", "%"+cast.ToString(val)+"%")
		builder.OrLike("`system_user`.`nickname`", "%"+cast.ToString(val)+"%")
		builder.OrLike("`system_user`.`mobile`", "%"+cast.ToString(val)+"%")
		builder.RightBracket()
	}
	if val, ok := condition["pagination"]; ok {
		pagination := val.(*sql.Pagination)
		if pagination != nil {
			builder.Offset(pagination.GetOffset())
			builder.Limit(pagination.GetLimit())
		}
	}
	builder.GroupBy("`system_user`.`id`")
	builder.OrderBy("`system_user`.`id`", sql.DESC)
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
	builder.Select("`system_user`.`id`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`system_user`.`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`system_user`.`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`system_user`.`status`", val)
	}
	if val, ok := condition["systemDeptId"]; ok {
		systemDeptId := cast.ToInt64(val)
		if systemDeptId > 0 {
			if deptList, err := SystemDeptListRecursive(ctx, systemDeptId); err == nil {
				newVal := make([]any, 0)
				for _, v := range deptList {
					newVal = append(newVal, cast.ToInt64(v.Id))
				}
				builder.LeftJoin("`system_user_dept`", "`system_user`.`id` = `system_user_dept`.`user_id`")
				builder.In("`system_user_dept`.`dept_id`", newVal...)
			}
		}
	}
	if val, ok := condition["username"]; ok {
		builder.And()
		builder.LeftBracket()
		builder.FirstLike("`system_user`.`username`", "%"+cast.ToString(val)+"%")
		builder.OrLike("`system_user`.`nickname`", "%"+cast.ToString(val)+"%")
		builder.OrLike("`system_user`.`mobile`", "%"+cast.ToString(val)+"%")
		builder.RightBracket()
	}
	builder.GroupBy("`system_user`.`id`")
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	newSql := "SELECT COUNT(1) AS C FROM(" + query + ") AS newTable"
	resMap, err := db.QueryRow(ctx, newSql, args...).ToMap()
	if err != nil {
		return 0, err
	}
	if len(resMap) < 1 {
		return 0, nil
	}
	v := resMap["C"]
	res, err = strconv.ParseInt(v, 10, 0)
	return
}

// SystemUserPassword 密码修改
func SystemUserPassword(ctx context.Context, id int64, password string) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["password"] = password
	query, args, err := builder.Table("`system_user`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}
