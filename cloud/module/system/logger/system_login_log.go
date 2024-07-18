package logger

import (
	"cloud/dao"
	"cloud/initial"
	"cloud/module/system/user"
	"context"
	"strconv"
	"strings"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// system_login_log 登录日志
// SystemLoginLogCreate 创建数据
func SystemLoginLogCreate(ctx context.Context, data dao.SystemLoginLog) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_login_log`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemLoginLogUpdate 更新数据
func SystemLoginLogUpdate(ctx context.Context, id int64, data dao.SystemLoginLog) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_login_log`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemLoginLogDelete 删除数据
func SystemLoginLogDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_login_log`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemLoginLog 查询单条数据
func SystemLoginLog(ctx context.Context, id int64) (res dao.SystemLoginLog, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_login_log`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemLoginLogRecover 恢复数据
func SystemLoginLogRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_login_log`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemLoginLogList 查询列表数据
func SystemLoginLogList(ctx context.Context, condition map[string]any) (res []dao.SystemLoginLog, err error) {
	if util.Empty(condition["userId"]) || util.Empty(condition["dataScope"]) || util.Empty(condition["dataScopeDept"]) || util.Empty(condition["tenantId"]) {
		return nil, errors.New("userId or dataScope or dataScopeDept or tenantId is empty")
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_login_log`")
	builder.Select("`system_login_log`.*")
	builder.LeftJoin("`system_user_dept`", "system_login_log.tenant_id = system_user_dept.tenant_id AND `system_login_log`.`user_id` = `system_user_dept`.`user_id` ")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`system_login_log`.`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`system_login_log`.`deleted`", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("`system_login_log`.`username`", val)
	}
	if val, ok := condition["beginLoginTime"]; ok {
		builder.GreaterEqual("`system_login_log`.`login_time`", val)
	}
	if val, ok := condition["finishLoginTime"]; ok {
		builder.LessEqual("`system_login_log`.`login_time`", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("`system_login_log`.`channel`", val)
	}

	builder.GroupBy("`system_login_log`.`id`")
	builder.OrderBy("`system_login_log`.`id`", sql.DESC)

	query, args, err := builder.Rows()
	if err != nil {
		return
	}

	dataScope := cast.ToInt32(condition["dataScope"])
	dataScopeDept := condition["dataScopeDept"].([]int64)
	userId := cast.ToInt64(condition["userId"])

	newSql := "SELECT newTable.* FROM (" + query + ") AS newTable "

	dataScopeDeptTotal := len(dataScopeDept) - 1
	newSql += " LEFT JOIN system_user_dept ON newTable.tenant_id = system_user_dept.tenant_id AND newTable.user_id = system_user_dept.user_id  WHERE system_user_dept.dept_id IN (?" + strings.Repeat(",?", dataScopeDeptTotal) + ") "
	for _, v := range dataScopeDept {
		args = append(args, v)
	}

	switch dataScope {
	case 1: // 全部数据权限
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				if deptList, err := user.SystemDeptListRecursive(ctx, deptId); err == nil {
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
				if deptList, err := user.SystemDeptListRecursive(ctx, deptId); err == nil {
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
				if deptList, err := user.SystemDeptListRecursive(ctx, deptId); err == nil {
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

// SystemLoginLogListTotal 查询列表数据总量
func SystemLoginLogListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	if util.Empty(condition["userId"]) || util.Empty(condition["dataScope"]) || util.Empty(condition["dataScopeDept"]) || util.Empty(condition["tenantId"]) {
		return 0, errors.New("userId or dataScope or dataScopeDept or tenantId is empty")
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()

	builder.Table("`system_login_log`")
	builder.Select("`system_login_log`.*")
	builder.LeftJoin("`system_user_dept`", "system_login_log.tenant_id = system_user_dept.tenant_id AND `system_login_log`.`user_id` = `system_user_dept`.`user_id` ")

	if val, ok := condition["tenantId"]; ok {
		builder.Where("`system_login_log`.`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`system_login_log`.`deleted`", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("`system_login_log`.`username`", val)
	}
	if val, ok := condition["beginLoginTime"]; ok {
		builder.GreaterEqual("`system_login_log`.`login_time`", val)
	}
	if val, ok := condition["finishLoginTime"]; ok {
		builder.LessEqual("`system_login_log`.`login_time`", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("`system_login_log`.`channel`", val)
	}
	builder.GroupBy("`system_login_log`.`id`")
	builder.OrderBy("`system_login_log`.`id`", sql.DESC)

	query, args, err := builder.Rows()
	if err != nil {
		return
	}

	newSql := "SELECT COUNT(1) AS C FROM ( SELECT newTable.* FROM(" + query + ") AS newTable"

	dataScope := cast.ToInt32(condition["dataScope"])
	dataScopeDept := condition["dataScopeDept"].([]int64)
	userId := cast.ToInt64(condition["userId"])

	dataScopeDeptTotal := len(dataScopeDept) - 1
	newSql += " LEFT JOIN system_user_dept ON newTable.tenant_id = system_user_dept.tenant_id AND newTable.user_id = system_user_dept.user_id  WHERE system_user_dept.dept_id IN (?" + strings.Repeat(",?", dataScopeDeptTotal) + ") "
	for _, v := range dataScopeDept {
		args = append(args, v)
	}

	switch dataScope {
	case 1: // 全部数据权限
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				if deptList, err := user.SystemDeptListRecursive(ctx, deptId); err == nil {
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
				if deptList, err := user.SystemDeptListRecursive(ctx, deptId); err == nil {
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
				if deptList, err := user.SystemDeptListRecursive(ctx, deptId); err == nil {
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

func SystemLoginLogDrop(ctx context.Context, condition map[string]any) (res int64, err error) {
	data := make(map[string]any)
	data["deleted"] = 1
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	builder.Table("`system_login_log`")
	var query string
	var args []any
	if val, ok := condition["ids"]; ok {
		ids := val.([]int64)
		id := make([]any, 0)
		for _, v := range ids {
			id = append(id, v)
		}
		builder.In("`id`", id...)
	}
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("`username`", val)
	}
	if val, ok := condition["beginLoginTime"]; ok {
		builder.GreaterEqual("`login_time`", val)
	}
	if val, ok := condition["finishLoginTime"]; ok {
		builder.LessEqual("`login_time`", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("`channel`", val)
	}
	query, args, err = builder.Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}
