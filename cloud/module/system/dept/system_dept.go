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
func SystemDeptUpdate(ctx context.Context, systemDeptId int64, data dao.SystemDept) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dept`").Where("`id`", systemDeptId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemDeptDelete 删除数据
func SystemDeptDelete(ctx context.Context, systemDeptId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_dept`").Where("`id`", systemDeptId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemDept 查询单条数据
func SystemDept(ctx context.Context, systemDeptId int64) (res dao.SystemDept, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_dept`").Where("`id`", systemDeptId).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemDeptRecover 恢复数据
func SystemDeptRecover(ctx context.Context, systemDeptId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_dept`").Where("`id`", systemDeptId).Update(data)
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
	if val, ok := condition["systemTenantId"]; ok {
		builder.Where("`system_tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("`name`", "%"+cast.ToString(val)+"%")
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
