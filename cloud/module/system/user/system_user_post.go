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
	if data.PostIds.IsValid() {
		var postIds []int64
		err = json.Unmarshal(data.PostIds.JSON, &postIds)
		if err != nil {
			return
		}
		for _, postId := range postIds {
			list = append(list, dao.SystemUserPost{
				PostId:     proto.Int64(postId),
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
		query, args, err := builder.Table("`system_user_post`").Where("`tenant_id`", data.TenantId).Where("user_id", data.UserId).Delete()
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
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("`user_id`", val)
	}
	if val, ok := condition["postId"]; ok {
		builder.Where("`post_id`", val)
	}
	builder.OrderBy("`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}
