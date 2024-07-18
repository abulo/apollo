package dept

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
)

// system_dept 部门
// SystemDeptCreate 创建数据
func SystemDeptCreate(ctx context.Context, data dao.SystemDept) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dept`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemDeptUpdate 更新数据
func SystemDeptUpdate(ctx context.Context, id int64, data dao.SystemDept) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dept`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemDeptDelete 删除数据
func SystemDeptDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_dept`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemDept 查询单条数据
func SystemDept(ctx context.Context, id int64) (res dao.SystemDept, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dept`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemDeptRecover 恢复数据
func SystemDeptRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_dept`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemDeptList 查询列表数据
func SystemDeptList(ctx context.Context, condition map[string]any) (res []dao.SystemDept, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_dept`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["parentId"]; ok {
		builder.Where("`parent_id`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Like("`name`", "%"+cast.ToString(val)+"%")
	}

	if val, ok := condition["pagination"]; ok {
		pagination := val.(*sql.Pagination)
		if pagination != nil {
			builder.Offset(pagination.GetOffset())
			builder.Limit(pagination.GetLimit())
		}
	}
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

// SystemDeptList 查询列表数据
func SystemDeptItem(ctx context.Context, condition map[string]any) (res dao.SystemDept, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_dept`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["parentId"]; ok {
		builder.Where("`parent_id`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Like("`name`", "%"+cast.ToString(val)+"%")
	}

	builder.OrderBy("`parent_id`", sql.ASC)
	builder.OrderBy("`sort`", sql.ASC)
	builder.OrderBy("`id`", sql.ASC)
	query, args, err := builder.Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemDeptListTotal 查询列表数据总量
func SystemDeptListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_dept`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["parentId"]; ok {
		builder.Where("`parent_id`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Like("`name`", "%"+cast.ToString(val)+"%")
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}

func SystemUserDeptTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_user_dept`")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("`user_id`", val)
	}
	if val, ok := condition["deptId"]; ok {
		builder.Where("`dept_id`", val)
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}
