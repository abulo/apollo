package tenant

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
)

// system_tenant 租户
// SystemTenantCreate 创建数据
func SystemTenantCreate(ctx context.Context, data dao.SystemTenant) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	res = 0
	err = db.Transact(ctx, func(ctx context.Context, session sql.Session) error {
		builder := sql.NewBuilder()
		data.UserId = null.Int64From(0)
		query, args, err := builder.Table("`system_tenant`").Insert(data)
		if err != nil {
			return err
		}
		// 插入租户信息
		id, err := session.Insert(ctx, query, args...)
		if err != nil {
			return err
		}
		// 插入租户管理员信息
		var user dao.SystemUserCustom
		user.Username = data.Username
		user.Password = data.Password
		user.Nickname = null.StringFrom(cast.ToString(data.Name))
		user.Mobile = null.StringFrom(cast.ToString(data.ContactMobile))
		user.Creator = data.Creator
		user.CreateTime = data.CreateTime
		user.Updater = data.Updater
		user.UpdateTime = data.UpdateTime
		user.TenantId = proto.Int64(id)
		user.Status = proto.Int32(0)
		user.Deleted = proto.Int32(0)
		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_user`").Insert(user)
		if err != nil {
			return err
		}
		// 插入用户和租户的绑定信息
		userId, err := session.Insert(ctx, query, args...)
		if err != nil {
			return err
		}

		// 绑定用户和租户的关系

		var userTenant dao.SystemUserTenant
		userTenant.UserId = proto.Int64(userId)
		userTenant.TenantId = proto.Int64(id)
		userTenant.Deleted = proto.Int32(0)
		userTenant.Creator = data.Creator
		userTenant.CreateTime = data.CreateTime
		userTenant.Updater = data.Updater
		userTenant.UpdateTime = data.UpdateTime
		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_user_tenant`").Insert(user)
		if err != nil {
			return err
		}
		// 插入用户信息
		_, err = session.Insert(ctx, query, args...)
		if err != nil {
			return err
		}

		// 更新租户管理员信息
		var updateTenant dao.SystemTenant
		updateTenant.UserId = null.Int64From(userId)
		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_tenant`").Where("`id`", id).Update(updateTenant)
		if err != nil {
			return err
		}
		_, err = session.Update(ctx, query, args...)
		if err != nil {
			return err
		}
		res = id
		return nil
	})

	return res, err
}

// SystemTenantUpdate 更新数据
func SystemTenantUpdate(ctx context.Context, id int64, data dao.SystemTenant) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_tenant`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemTenantDelete 删除数据
func SystemTenantDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_tenant`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemTenant 查询单条数据
func SystemTenant(ctx context.Context, id int64) (res dao.SystemTenant, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_tenant`").Where("`id`", id).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemTenantRecover 恢复数据
func SystemTenantRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_tenant`").Where("`id`", id).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemTenantList 查询列表数据
func SystemTenantList(ctx context.Context, condition map[string]any) (res []dao.SystemTenant, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_tenant`")
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Like("`name`", "%"+cast.ToString(val)+"%")
	}
	if val, ok := condition["beginExpireDate"]; ok {
		builder.GreaterEqual("`expire_date`", val)
	}
	if val, ok := condition["finishExpireDate"]; ok {
		builder.LessEqual("`expire_date`", val)
	}
	if val, ok := condition["systemTenantPackageId"]; ok {
		builder.Where("`system_tenant_package_id`", val)
	}

	if val, ok := condition["pagination"]; ok {
		pagination := val.(*sql.Pagination)
		if pagination != nil {
			builder.Offset(pagination.GetOffset())
			builder.Limit(pagination.GetLimit())
		}
	}
	builder.OrderBy("`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}

// SystemTenantListTotal 查询列表数据总量
func SystemTenantListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_tenant`")
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Like("`name`", "%"+cast.ToString(val)+"%")
	}
	if val, ok := condition["beginExpireDate"]; ok {
		builder.GreaterEqual("`expire_date`", val)
	}
	if val, ok := condition["finishExpireDate"]; ok {
		builder.LessEqual("`expire_date`", val)
	}
	if val, ok := condition["systemTenantPackageId"]; ok {
		builder.Where("`system_tenant_package_id`", val)
	}

	query, args, err := builder.Count()
	if err != nil {
		return
	}
	res, err = db.Count(ctx, query, args...)
	return
}
