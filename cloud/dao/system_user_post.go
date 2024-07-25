package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemUserPost 系统用户职位 system_user_post
type SystemUserPost struct {
	Id         *int64        `db:"id,-" json:"id"`                //bigint 编号,PRI
	UserId     *int64        `db:"user_id" json:"userId"`         //bigint 系统用户 ID
	PostId     *int64        `db:"post_id" json:"postId"`         //bigint 职位 id
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}

type SystemUserPostCustom struct {
	UserId     *int64        `db:"user_id" json:"userId"`         //bigint 系统用户 id
	PostIds    null.JSON     `db:"post_ids" json:"postIds"`       //bigint 职位 id
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
