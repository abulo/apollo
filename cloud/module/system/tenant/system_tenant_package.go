package tenant

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
)

// system_tenant_package 租户套餐包
// SystemTenantPackageCreate 创建数据
func SystemTenantPackageCreate(ctx context.Context, data dao.SystemTenantPackage) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_tenant_package`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemTenantPackageUpdate 更新数据
func SystemTenantPackageUpdate(ctx context.Context, systemTenantPackageId int64, data dao.SystemTenantPackage) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_tenant_package`").Where("`id`", systemTenantPackageId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemTenantPackageDelete 删除数据
func SystemTenantPackageDelete(ctx context.Context, systemTenantPackageId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_tenant_package`").Where("`id`", systemTenantPackageId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemTenantPackage 查询单条数据
func SystemTenantPackage(ctx context.Context, systemTenantPackageId int64) (res dao.SystemTenantPackage, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_tenant_package`").Where("`id`", systemTenantPackageId).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemTenantPackageRecover 恢复数据
func SystemTenantPackageRecover(ctx context.Context, systemTenantPackageId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_tenant_package`").Where("`id`", systemTenantPackageId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemTenantPackageList 查询列表数据
func SystemTenantPackageList(ctx context.Context, condition map[string]any) (res []dao.SystemTenantPackage, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_tenant_package`")
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Like("`name`", "%"+cast.ToString(val)+"%")
	}

	builder.OrderBy("`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}
