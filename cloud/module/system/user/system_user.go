package user

import (
	"cloud/dao"
	"cloud/initial"
	"cloud/module/system/dept"
	"cloud/module/system/tenant"
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
)

// system_user 系统用户
// SystemUserCreate 创建数据
func SystemUserCreate(ctx context.Context, data dao.SystemUser) (res int64, err error) {

	db := initial.Core.Store.LoadSQL("mysql").Write()
	res = 0
	err = db.Transact(ctx, func(ctx context.Context, session sql.Session) error {
		builder := sql.NewBuilder()
		query, args, err := builder.Table("`system_user`").Insert(data)
		if err != nil {
			return err
		}
		userId, err := session.Insert(ctx, query, args...)
		if err != nil {
			return err
		}
		var userDeptList []any
		var userDeptIds []int64
		err = json.Unmarshal(data.DeptIds.JSON, &userDeptIds)
		if err != nil {
			return err
		}
		for _, deptId := range userDeptIds {
			userDeptList = append(userDeptList, dao.SystemUserDept{
				UserId:     proto.Int64(userId),
				DeptId:     proto.Int64(deptId),
				Deleted:    proto.Int32(0),
				TenantId:   data.TenantId,
				Creator:    data.Creator,
				CreateTime: data.CreateTime,
				Updater:    data.Updater,
				UpdateTime: data.UpdateTime,
			})
		}
		if len(userDeptList) > 0 {
			builder = sql.NewBuilder()
			query, args, err = builder.Table("`system_user_dept`").MultiInsert(userDeptList...)
			if err != nil {
				return err
			}
			_, err = session.MultiInsert(ctx, query, args...)
			if err != nil {
				return err
			}
		}

		var userPostList []any
		var userPostIds []int64
		err = json.Unmarshal(data.PostIds.JSON, &userPostIds)
		if err != nil {
			return err
		}
		for _, postId := range userPostIds {
			userPostList = append(userPostList, dao.SystemUserPost{
				UserId:     proto.Int64(userId),
				PostId:     proto.Int64(postId),
				Deleted:    proto.Int32(0),
				TenantId:   data.TenantId,
				Creator:    data.Creator,
				CreateTime: data.CreateTime,
				Updater:    data.Updater,
				UpdateTime: data.UpdateTime,
			})
		}
		if len(userPostList) > 0 {
			builder = sql.NewBuilder()
			query, args, err = builder.Table("`system_user_post`").MultiInsert(userPostList...)
			if err != nil {
				return err
			}
			_, err = session.MultiInsert(ctx, query, args...)
			if err != nil {
				return err
			}
		}

		var userRoleList []any
		var userRoleIds []int64
		err = json.Unmarshal(data.RoleIds.JSON, &userRoleIds)
		if err != nil {
			return err
		}
		for _, roleId := range userRoleIds {
			userRoleList = append(userRoleList, dao.SystemUserRole{
				UserId:     proto.Int64(userId),
				RoleId:     proto.Int64(roleId),
				Deleted:    proto.Int32(0),
				TenantId:   data.TenantId,
				Creator:    data.Creator,
				CreateTime: data.CreateTime,
				Updater:    data.Updater,
				UpdateTime: data.UpdateTime,
			})
		}
		if len(userRoleList) > 0 {
			builder = sql.NewBuilder()
			query, args, err = builder.Table("`system_user_role`").MultiInsert(userRoleList...)
			if err != nil {
				return err
			}
			_, err = session.MultiInsert(ctx, query, args...)
			if err != nil {
				return err
			}
		}
		res = userId
		return nil
	})
	return
}

// SystemUserUpdate 更新数据
func SystemUserUpdate(ctx context.Context, userId int64, data dao.SystemUser) (res int64, err error) {

	db := initial.Core.Store.LoadSQL("mysql").Write()
	res = 0
	err = db.Transact(ctx, func(ctx context.Context, session sql.Session) error {
		builder := sql.NewBuilder()
		query, args, err := builder.Table("`system_user`").Where("`id`", userId).Update(data)
		if err != nil {
			return err
		}
		res, err = session.Update(ctx, query, args...)
		if err != nil {
			return err
		}

		var userDeptList []any
		var userDeptIds []int64
		err = json.Unmarshal(data.DeptIds.JSON, &userDeptIds)
		if err != nil {
			return err
		}
		for _, deptId := range userDeptIds {
			userDeptList = append(userDeptList, dao.SystemUserDept{
				UserId:     proto.Int64(userId),
				DeptId:     proto.Int64(deptId),
				Deleted:    proto.Int32(0),
				TenantId:   data.TenantId,
				Creator:    data.Creator,
				CreateTime: data.CreateTime,
				Updater:    data.Updater,
				UpdateTime: data.UpdateTime,
			})
		}
		if len(userDeptList) > 0 {

			builder = sql.NewBuilder()
			query, args, err := builder.Table("`system_user_dept`").Where("`tenant_id`", data.TenantId).Where("`user_id`", userId).Delete()
			if err != nil {
				return err
			}
			_, err = session.Delete(ctx, query, args...)
			if err != nil {
				return err
			}

			builder = sql.NewBuilder()
			query, args, err = builder.Table("`system_user_dept`").MultiInsert(userDeptList...)
			if err != nil {
				return err
			}
			_, err = session.MultiInsert(ctx, query, args...)
			if err != nil {
				return err
			}
		}

		var userPostList []any
		var userPostIds []int64
		err = json.Unmarshal(data.PostIds.JSON, &userPostIds)
		if err != nil {
			return err
		}
		for _, postId := range userPostIds {
			userPostList = append(userPostList, dao.SystemUserPost{
				UserId:     proto.Int64(userId),
				PostId:     proto.Int64(postId),
				Deleted:    proto.Int32(0),
				TenantId:   data.TenantId,
				Creator:    data.Creator,
				CreateTime: data.CreateTime,
				Updater:    data.Updater,
				UpdateTime: data.UpdateTime,
			})
		}
		if len(userPostList) > 0 {
			builder = sql.NewBuilder()
			query, args, err := builder.Table("`system_user_post`").Where("`tenant_id`", data.TenantId).Where("`user_id`", userId).Delete()
			if err != nil {
				return err
			}
			_, err = session.Delete(ctx, query, args...)
			if err != nil {
				return err
			}

			builder = sql.NewBuilder()
			query, args, err = builder.Table("`system_user_post`").MultiInsert(userPostList...)
			if err != nil {
				return err
			}
			_, err = session.MultiInsert(ctx, query, args...)
			if err != nil {
				return err
			}
		}

		var userRoleList []any
		var userRoleIds []int64
		err = json.Unmarshal(data.RoleIds.JSON, &userRoleIds)
		if err != nil {
			return err
		}
		for _, roleId := range userRoleIds {
			userRoleList = append(userRoleList, dao.SystemUserRole{
				UserId:     proto.Int64(userId),
				RoleId:     proto.Int64(roleId),
				Deleted:    proto.Int32(0),
				TenantId:   data.TenantId,
				Creator:    data.Creator,
				CreateTime: data.CreateTime,
				Updater:    data.Updater,
				UpdateTime: data.UpdateTime,
			})
		}
		if len(userRoleList) > 0 {
			builder = sql.NewBuilder()
			query, args, err := builder.Table("`system_user_role`").Where("`tenant_id`", data.TenantId).Where("`user_id`", userId).Delete()
			if err != nil {
				return err
			}
			_, err = session.Delete(ctx, query, args...)
			if err != nil {
				return err
			}
			builder = sql.NewBuilder()
			query, args, err = builder.Table("`system_user_role`").MultiInsert(userRoleList...)
			if err != nil {
				return err
			}
			_, err = session.MultiInsert(ctx, query, args...)
			if err != nil {
				return err
			}
		}
		return nil
	})
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

	builder.Table("`system_user`")
	builder.Select("`system_user`.*", "JSON_ARRAYAGG(system_user_dept.dept_id) AS dept_ids",
		"JSON_ARRAYAGG(system_user_post.post_id) AS post_ids",
		"JSON_ARRAYAGG(system_user_role.role_id) AS role_ids ")
	builder.LeftJoin("`system_user_dept`", "`system_user`.`id` = `system_user_dept`.`user_id`   AND system_user.tenant_id = system_user_dept.tenant_id")
	builder.LeftJoin("`system_user_post`", "`system_user`.`id` = `system_user_post`.`user_id`  AND system_user.tenant_id = system_user_post.tenant_id")
	builder.LeftJoin("`system_user_role`", "`system_user`.`id` = `system_user_role`.`user_id`  AND system_user.tenant_id = system_user_role.tenant_id")
	builder.Where("`system_user`.`id`", id)
	builder.GroupBy("`system_user`.`id`")

	query, args, err := builder.Row()
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

// SystemUserDrop 恢复数据
func SystemUserDrop(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	userInfo, err := SystemUser(ctx, id)
	if err != nil {
		return 0, err
	}
	err = db.Transact(ctx, func(ctx context.Context, session sql.Session) error {
		builder := sql.NewBuilder()
		data := make(map[string]any)
		data["tenant_id"] = 0
		query, args, err := builder.Table("`system_user`").Where("`id`", id).Update(data)
		if err != nil {
			return err
		}
		res, err = session.Update(ctx, query, args...)
		if err != nil {
			return err
		}
		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_user_role`").Where("`tenant_id`", userInfo.TenantId).Where("`user_id`", userInfo.Id).Delete()
		if err != nil {
			return err
		}
		_, err = session.Delete(ctx, query, args...)
		if err != nil {
			return err
		}

		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_user_post`").Where("`tenant_id`", userInfo.TenantId).Where("`user_id`", userInfo.Id).Delete()
		if err != nil {
			return err
		}
		_, err = session.Delete(ctx, query, args...)
		if err != nil {
			return err
		}

		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_user_dept`").Where("`tenant_id`", userInfo.TenantId).Where("`user_id`", userInfo.Id).Delete()
		if err != nil {
			return err
		}
		_, err = session.Delete(ctx, query, args...)
		if err != nil {
			return err
		}

		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_user_tenant`").Where("`tenant_id`", userInfo.TenantId).Where("`user_id`", userInfo.Id).Delete()
		if err != nil {
			return err
		}
		_, err = session.Delete(ctx, query, args...)
		if err != nil {
			return err
		}

		return nil
	})
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
	builder.Select("`system_user`.*", "JSON_ARRAYAGG(system_user_dept.dept_id) AS dept_ids",
		"JSON_ARRAYAGG(system_user_post.post_id) AS post_ids",
		"JSON_ARRAYAGG(system_user_role.role_id) AS role_ids ")
	builder.LeftJoin("`system_user_dept`", "`system_user`.`id` = `system_user_dept`.`user_id`   AND system_user.tenant_id = system_user_dept.tenant_id")
	builder.LeftJoin("`system_user_post`", "`system_user`.`id` = `system_user_post`.`user_id`  AND system_user.tenant_id = system_user_post.tenant_id")
	builder.LeftJoin("`system_user_role`", "`system_user`.`id` = `system_user_role`.`user_id`  AND system_user.tenant_id = system_user_role.tenant_id")
	if val, ok := condition["username"]; ok {
		builder.Where("system_user.`username`", val)
	}
	builder.GroupBy("`system_user`.`id`")
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
	builder.Select("`system_user`.*", "JSON_ARRAYAGG(system_user_dept.dept_id) AS dept_ids",
		"JSON_ARRAYAGG(system_user_post.post_id) AS post_ids",
		"JSON_ARRAYAGG(system_user_role.role_id) AS role_ids ")
	builder.LeftJoin("`system_user_dept`", "`system_user`.`id` = `system_user_dept`.`user_id`   AND system_user.tenant_id = system_user_dept.tenant_id")
	builder.LeftJoin("`system_dept`", "`system_dept`.tenant_id = system_user_dept.tenant_id  AND system_dept.deleted = 0 ")
	builder.LeftJoin("`system_user_post`", "`system_user`.`id` = `system_user_post`.`user_id`  AND system_user.tenant_id = system_user_post.tenant_id")
	builder.LeftJoin("`system_user_role`", "`system_user`.`id` = `system_user_role`.`user_id`  AND system_user.tenant_id = system_user_role.tenant_id")

	if val, ok := condition["tenantId"]; ok {
		builder.Where("`system_user`.`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`system_user`.`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`system_user`.`status`", val)
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
	builder.OrderBy("`system_user`.`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}

	if util.Empty(condition["userId"]) || util.Empty(condition["dataScope"]) || util.Empty(condition["dataScopeDept"]) || util.Empty(condition["tenantId"]) {
		return nil, errors.New("userId or dataScope or dataScopeDept or tenantId is empty")
	}
	newSql := "SELECT newTable.* FROM (" + query + ") AS newTable"

	dataScope := cast.ToInt32(condition["dataScope"])
	dataScopeDept := condition["dataScopeDept"].([]int64)
	userId := cast.ToInt64(condition["userId"])

	dataScopeDeptTotal := len(dataScopeDept) - 1
	newSql += " LEFT JOIN system_user_dept ON newTable.id = system_user_dept.user_id AND newTable.tenant_id = system_user_dept.tenant_id WHERE system_user_dept.dept_id IN (?" + strings.Repeat(",?", dataScopeDeptTotal) + ") "
	for _, v := range dataScopeDept {
		args = append(args, v)
	}

	switch dataScope {
	case 1: // 全部数据权限
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				if deptList, err := SystemDeptListRecursive(ctx, deptId); err == nil {
					for _, v := range deptList {
						args = append(args, cast.ToInt64(v.Id))
					}
					total := cast.ToInt(len(deptList)) - 1
					newSql += " AND system_user_dept.dept_id IN (?" + strings.Repeat(",?", total) + ") "
				}
			}
		}
	case 2: // 指定部门数据权限
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				if deptList, err := SystemDeptListRecursive(ctx, deptId); err == nil {
					for _, v := range deptList {
						args = append(args, cast.ToInt64(v.Id))
					}
					total := cast.ToInt(len(deptList)) - 1
					newSql += " AND system_user_dept.dept_id IN (?" + strings.Repeat(",?", total) + ") "
				}
			}
		}
	case 3: // 本部门数据权限
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				newSql += " AND system_user_dept.dept_id = ? "
				args = append(args, deptId)
			}
		}
	case 4: // 本部门及以下数据权限
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				if deptList, err := SystemDeptListRecursive(ctx, deptId); err == nil {
					for _, v := range deptList {
						args = append(args, cast.ToInt64(v.Id))
					}
					total := cast.ToInt(len(deptList)) - 1
					newSql += " AND system_user_dept.dept_id IN (?" + strings.Repeat(",?", total) + ") "
				}
			}
		}
	case 5: // 仅本人数据权限
		newSql += " AND system_user_dept.user_id = ? "
		args = append(args, userId)
	}

	newSql += " GROUP BY newTable.id "
	if val, ok := condition["pagination"]; ok {
		pagination := val.(*sql.Pagination)
		if pagination != nil {
			newSql += " LIMIT ?,?"
			args = append(args, pagination.GetOffset(), pagination.GetLimit())
		}
	}
	err = db.QueryRows(ctx, newSql, args...).ToStruct(&res)
	return
}

// SystemUserListTotal 查询列表数据总量
func SystemUserListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_user`")
	builder.Select("`system_user`.*")
	builder.LeftJoin("`system_user_dept`", "`system_user`.`id` = `system_user_dept`.`user_id`   AND system_user.tenant_id = system_user_dept.tenant_id")
	builder.LeftJoin("`system_dept`", "`system_dept`.tenant_id = system_user_dept.tenant_id  AND system_dept.deleted = 0 ")
	builder.LeftJoin("`system_user_post`", "`system_user`.`id` = `system_user_post`.`user_id`  AND system_user.tenant_id = system_user_post.tenant_id")
	builder.LeftJoin("`system_user_role`", "`system_user`.`id` = `system_user_role`.`user_id`  AND system_user.tenant_id = system_user_role.tenant_id")
	builder.Select("`system_user`.`id`", "`system_user`.`tenant_id`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`system_user`.`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`system_user`.`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`system_user`.`status`", val)
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
	newSql := "SELECT COUNT(1) AS C FROM ( SELECT newTable.* FROM(" + query + ") AS newTable"

	if util.Empty(condition["userId"]) || util.Empty(condition["dataScope"]) || util.Empty(condition["dataScopeDept"]) || util.Empty(condition["tenantId"]) {
		return 0, errors.New("userId or dataScope or dataScopeDept or tenantId is empty")
	}

	dataScope := cast.ToInt32(condition["dataScope"])
	dataScopeDept := condition["dataScopeDept"].([]int64)
	userId := cast.ToInt64(condition["userId"])

	dataScopeDeptTotal := len(dataScopeDept) - 1
	newSql += " LEFT JOIN system_user_dept ON newTable.id = system_user_dept.user_id AND newTable.tenant_id = system_user_dept.tenant_id WHERE system_user_dept.dept_id IN (?" + strings.Repeat(",?", dataScopeDeptTotal) + ") "
	for _, v := range dataScopeDept {
		args = append(args, v)
	}

	switch dataScope {
	case 1: // 全部数据权限
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				if deptList, err := SystemDeptListRecursive(ctx, deptId); err == nil {
					for _, v := range deptList {
						args = append(args, cast.ToInt64(v.Id))
					}
					total := cast.ToInt(len(deptList)) - 1
					newSql += " AND system_user_dept.dept_id IN (?" + strings.Repeat(",?", total) + ") "
				}
			}
		}
	case 2: // 指定部门数据权限
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				if deptList, err := SystemDeptListRecursive(ctx, deptId); err == nil {
					for _, v := range deptList {
						args = append(args, cast.ToInt64(v.Id))
					}
					total := cast.ToInt(len(deptList)) - 1
					newSql += " AND system_user_dept.dept_id IN (?" + strings.Repeat(",?", total) + ") "
				}
			}
		}
	case 3: // 本部门数据权限
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				newSql += " AND system_user_dept.dept_id = ? "
				args = append(args, deptId)
			}
		}
	case 4: // 本部门及以下数据权限
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				if deptList, err := SystemDeptListRecursive(ctx, deptId); err == nil {
					for _, v := range deptList {
						args = append(args, cast.ToInt64(v.Id))
					}
					total := cast.ToInt(len(deptList)) - 1
					newSql += " AND system_user_dept.dept_id IN (?" + strings.Repeat(",?", total) + ") "
				}
			}
		}
	case 5: // 仅本人数据权限
		newSql += " AND system_user_dept.user_id = ? "
		args = append(args, userId)
	}
	newSql += " GROUP BY newTable.id) AS nTable"
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

// SystemDeptListRecursive 递归查询
func SystemDeptListRecursive(ctx context.Context, id int64) (res []dao.SystemDept, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	query := "WITH RECURSIVE filter_system_dept AS (SELECT `id`,`name`,`parent_id`,`sort`,`user_id`,`phone`,`email`,`status`,`deleted`,`tenant_id`,`creator`,`create_time`,`updater`,`update_time` FROM `system_dept` WHERE `id`=? UNION ALL SELECT d.* FROM system_dept AS d INNER JOIN filter_system_dept f ON d.parent_id=f.id) SELECT filter_system_dept.* FROM filter_system_dept"
	var deptList []dao.SystemDept
	err = db.QueryRows(ctx, query, id).ToStruct(&deptList)
	if err == nil {
		for _, v := range deptList {
			if cast.ToInt32(v.Deleted) == 0 {
				res = append(res, v)
			}
		}
	}
	return
}

// SystemUserDataScope 用户数据权限
func SystemUserDataScope(ctx context.Context, tenantId, userId int64) (res dao.SystemUserRoleDataScope, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	tenantItem, err := tenant.SystemTenant(ctx, tenantId)
	if err != nil {
		return res, err
	}
	deptList, err := dept.SystemDeptList(ctx, map[string]any{"tenantId": tenantId, "deleted": 0})
	if err != nil {
		return res, err
	}
	superDept := make([]int64, 0)
	for _, v := range deptList {
		superDept = append(superDept, *v.Id)
	}
	if len(superDept) < 1 {
		superDept = append(superDept, 0)
	}
	if tenantItem.UserId.IsValid() && cast.ToInt64(tenantItem.UserId.Ptr()) == userId {
		// 如果是租户管理员, 则数据权限为全部数据权限
		res.DataScope = proto.Int32(1)
		res.DataScopeDept = superDept
		return res, nil
	}
	builder := sql.NewBuilder()
	builder.Table("`system_role`")
	builder.Select("`system_role`.*")
	builder.LeftJoin("`system_user_role`", "`system_role`.`id` = `system_user_role`.`role_id` AND `system_role`.`tenant_id` = `system_user_role`.`tenant_id`")
	builder.Where("`system_user_role`.`tenant_id`", tenantId)
	builder.Where("`system_user_role`.`user_id`", userId)
	builder.Where("`system_role`.`tenant_id`", tenantId)
	builder.Where("`system_role`.`deleted`", 0)
	builder.GroupBy("`system_role`.`id`")
	builder.OrderBy("`system_role`.`data_scope`", sql.ASC)
	query, args, err := builder.Rows()
	if err != nil {
		return res, err
	}
	var systemRoleList []dao.SystemRole
	err = db.QueryRows(ctx, query, args...).ToStruct(&systemRoleList)
	if err != nil {
		return res, err
	}
	// systemRoleList 是一个数组, 这里需要遍历数据, dataScope就取数据的第一个对象里面的 data_scope 值, dataScopeDept 需要获取整个数组的值
	var dataScope int32
	var dataScopeDept []int64
	if len(systemRoleList) > 0 {
		if systemRoleList[0].DataScope != nil {
			dataScope = cast.ToInt32(systemRoleList[0].DataScope)
		}
		for _, v := range systemRoleList {
			if v.DataScopeDept.IsValid() {
				tmp := make([]int64, 0)
				if err := json.Unmarshal(*v.DataScopeDept.Ptr(), &tmp); err == nil {
					dataScopeDept = append(dataScopeDept, tmp...)
				}
			}
		}
	}
	builder = sql.NewBuilder()
	builder.Table("`system_dept`")
	builder.Select("`system_dept`.*")
	builder.LeftJoin("`system_user_dept`", "`system_dept`.`id` = `system_user_dept`.`dept_id` AND `system_dept`.`tenant_id` = `system_user_dept`.`tenant_id`")
	builder.Where("`system_user_dept`.`tenant_id`", tenantId)
	builder.Where("`system_user_dept`.`user_id`", userId)
	builder.Where("`system_dept`.`tenant_id`", tenantId)
	builder.Where("`system_dept`.`deleted`", 0)
	query, args, err = builder.Rows()
	if err != nil {
		return res, err
	}
	var systemDeptList []dao.SystemDept
	err = db.QueryRows(ctx, query, args...).ToStruct(&systemDeptList)
	if err != nil {
		return res, err
	}
	var dept []int64
	var deptTree []int64
	for _, v := range systemDeptList {
		dept = append(dept, *v.Id)
		deptTree = append(deptTree, *v.Id)
		if deptList, err := SystemDeptListRecursive(ctx, *v.Id); err == nil {
			for _, v := range deptList {
				if !util.InArray(*v.Id, deptTree) {
					deptTree = append(deptTree, *v.Id)
				}
			}
		}
	}

	// 查询本部门数据权限
	if dataScope == 1 { // 全部数据权限
		res.DataScopeDept = superDept
	} else if dataScope == 2 { // 指定部门数据权限
		res.DataScopeDept = dataScopeDept
	} else if dataScope == 3 { // 本部门数据权限
		res.DataScopeDept = dept
	} else if dataScope == 4 { // 本部门及以下数据权限
		res.DataScopeDept = deptTree
	} else { // 仅仅自己
		res.DataScopeDept = dept
	}
	if len(res.DataScopeDept) < 1 {
		res.DataScopeDept = append(res.DataScopeDept, 0)
	}
	res.DataScope = proto.Int32(dataScope)
	// res.DataScopeDept = dataScopeDept
	return res, nil
}
