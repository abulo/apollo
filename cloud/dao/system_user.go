package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemUser 系统用户 system_user
type SystemUser struct {
	Id             *int64        `db:"id" json:"id"`                           //bigint 用户编号,PRI
	Nickname       null.String   `db:"nickname" json:"nickname"`               //varchar 昵称
	Mobile         null.String   `db:"mobile" json:"mobile"`                   //varchar 手机号码
	Username       *string       `db:"username" json:"username"`               //varchar 用户名称
	Password       *string       `db:"password" json:"password"`               //varchar 用户密码
	Status         *int32        `db:"status" json:"status"`                   //tinyint 用户状态（0正常 1停用）
	Deleted        *int32        `db:"deleted" json:"deleted"`                 //tinyint 是否删除(0否 1是)
	SystemTenantId *int64        `db:"system_tenant_id" json:"systemTenantId"` //bigint 租户ID
	Creator        null.String   `db:"creator" json:"creator"`                 //varchar 创建人
	CreateTime     null.DateTime `db:"create_time" json:"createTime"`          //datetime 创建时间
	Updater        null.String   `db:"updater" json:"updater"`                 //varchar 更新人
	UpdateTime     null.DateTime `db:"update_time" json:"updateTime"`          //datetime 更新时间
}
