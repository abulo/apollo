package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemUserPost 系统用户职位 system_user_post
type SystemUserPost struct {
	Id             *int64        `db:"id" json:"id"`                           //bigint 编号,PRI
	SystemUserId   *int64        `db:"system_user_id" json:"systemUserId"`     //bigint 系统用户 ID
	SystemPostId   *int64        `db:"system_post_id" json:"systemPostId"`     //bigint 职位 id
	Deleted        *int32        `db:"deleted" json:"deleted"`                 //tinyint 删除
	SystemTenantId *int64        `db:"system_tenant_id" json:"systemTenantId"` //bigint 租户
	Creator        null.String   `db:"creator" json:"creator"`                 //varchar 创建人
	CreateTime     null.DateTime `db:"create_time" json:"createTime"`          //datetime 创建时间
	Updater        null.String   `db:"updater" json:"updater"`                 //varchar 更新人
	UpdateTime     null.DateTime `db:"update_time" json:"updateTime"`          //datetime 更新时间
}

type SystemUserPostCustom struct {
	SystemUserId   *int64        `db:"system_user_id" json:"systemUserId"`     //bigint 系统用户 id
	SystemPostIds  null.JSON     `db:"system_post_ids" json:"systemPostIds"`   //bigint 职位 id
	Deleted        *int32        `db:"deleted" json:"deleted"`                 //tinyint 删除
	SystemTenantId *int64        `db:"system_tenant_id" json:"systemTenantId"` //bigint 租户
	Creator        null.String   `db:"creator" json:"creator"`                 //varchar 创建人
	CreateTime     null.DateTime `db:"create_time" json:"createTime"`          //datetime 创建时间
	Updater        null.String   `db:"updater" json:"updater"`                 //varchar 更新人
	UpdateTime     null.DateTime `db:"update_time" json:"updateTime"`          //datetime 更新时间
}
