package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemRoleMenu 系统角色和系统菜单关联表 system_role_menu
type SystemRoleMenu struct {
	Id         *int64        `db:"id" json:"id"`                  //bigint 自增编号,PRI
	RoleId     *int64        `db:"role_id" json:"roleId"`         //bigint 角色编号
	MenuId     *int64        `db:"menu_id" json:"menuId"`         //bigint 菜单编号
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建者
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新者
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
