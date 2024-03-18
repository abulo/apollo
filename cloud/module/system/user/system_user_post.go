package user

import (
	"cloud/dao"
	"cloud/initial"
	"context"
	"encoding/json"

	"github.com/abulo/ratel/v3/stores/sql"
	"google.golang.org/protobuf/proto"
)

// system_user_post 系统用户职位
// SystemUserPostCreate 创建数据
func SystemUserPostCreate(ctx context.Context, data dao.SystemUserPostCustom) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var list []any
	if data.SystemPostIds.IsValid() {
		var postIds []int64
		err = json.Unmarshal(data.SystemPostIds.JSON, &postIds)
		if err != nil {
			return
		}
		for _, postId := range postIds {
			list = append(list, dao.SystemUserPost{
				SystemPostId:   proto.Int64(postId),
				SystemUserId:   data.SystemUserId,
				Deleted:        proto.Int32(0),
				SystemTenantId: data.SystemTenantId,
				Creator:        data.Creator,
				CreateTime:     data.CreateTime,
				Updater:        data.Updater,
				UpdateTime:     data.UpdateTime,
			})
		}
	}
	res = 0
	err = db.Transact(ctx, func(ctx context.Context, session sql.Session) error {
		// 先删除数据, 在添加数据
		// 需要先将数据删除了在添加
		builder := sql.NewBuilder()
		query, args, err := builder.Table("`system_user_post`").Where("`system_tenant_id`", data.SystemTenantId).Where("system_user_id", data.SystemUserId).Delete()
		if err != nil {
			return err
		}
		_, err = session.Delete(ctx, query, args...)
		if err != nil {
			return err
		}

		builder = sql.NewBuilder()
		query, args, err = builder.Table("`system_user_post`").MultiInsert(list...)
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

// SystemUserPostList 查询列表数据
func SystemUserPostList(ctx context.Context, condition map[string]any) (res []dao.SystemUserPost, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_user_post`")
	if val, ok := condition["systemTenantId"]; ok {
		builder.Where("`system_tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["systemUserId"]; ok {
		builder.Where("`system_user_id`", val)
	}
	if val, ok := condition["systemPostId"]; ok {
		builder.Where("`system_post_id`", val)
	}
	builder.OrderBy("`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}
