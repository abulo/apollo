package tenant

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
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
		data.SystemUserId = null.Int64From(0)
		query, args, err := builder.Table("`system_tenant`").Insert(data)
		if err != nil {
			return err
		}
		// 插入租户信息
		tenantId, err := session.Insert(ctx, query, args...)
		if err != nil {
			return err
		}
		// 插入租户管理员信息
		var user dao.SystemUser
		user.Username = data.Username
		user.Password = data.Password
		user.Nickname = null.StringFrom(cast.ToString(data.Name))
		user.Mobile = null.StringFrom(cast.ToString(data.ContactMobile))
		user.Creator = data.Creator
		user.CreateTime = data.CreateTime
		user.Updater = data.Updater
		user.UpdateTime = data.UpdateTime
		user.SystemTenantId = proto.Int64(tenantId)
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
		userTenant.SystemUserId = proto.Int64(userId)
		userTenant.SystemTenantId = proto.Int64(tenantId)
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
		updateTenant.SystemUserId = null.Int64From(userId)
		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_tenant`").Where("`id`", tenantId).Update(updateTenant)
		if err != nil {
			return err
		}
		_, err = session.Update(ctx, query, args...)
		if err != nil {
			return err
		}
		res = tenantId
		return nil
	})

	return res, err
}

// SystemTenantUpdate 更新数据
func SystemTenantUpdate(ctx context.Context, systemTenantId int64, data dao.SystemTenant) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_tenant`").Where("`id`", systemTenantId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemTenantDelete 删除数据
func SystemTenantDelete(ctx context.Context, systemTenantId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_tenant`").Where("`id`", systemTenantId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemTenant 查询单条数据
func SystemTenant(ctx context.Context, systemTenantId int64) (res dao.SystemTenant, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_tenant`").Where("`id`", systemTenantId).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemTenantRecover 恢复数据
func SystemTenantRecover(ctx context.Context, systemTenantId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_tenant`").Where("`id`", systemTenantId).Update(data)
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

	if !util.Empty(condition["offset"]) {
		builder.Offset(cast.ToInt64(condition["offset"]))
	}
	if !util.Empty(condition["limit"]) {
		builder.Limit(cast.ToInt64(condition["limit"]))
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
