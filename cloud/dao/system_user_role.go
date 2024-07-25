package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemUserRole 系统用户和系统角色关联表 system_user_role
type SystemUserRole struct {
	Id         *int64        `db:"id,-" json:"id"`                //bigint 自增编号,PRI
	UserId     *int64        `db:"user_id" json:"userId"`         //bigint 用户编号
	RoleId     *int64        `db:"role_id" json:"roleId"`         //bigint 角色编号
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建者
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新者
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}

type SystemUserRoleCustom struct {
	UserId     *int64        `db:"user_id" json:"userId"`         //bigint 用户编号
	RoleIds    null.JSON     `db:"role_ids" json:"roleIds"`       //bigint 角色编号
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建者
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新者
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
