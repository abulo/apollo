package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemRoleMenuCustom 角色和菜单关联表 system_role_menu
type SystemRoleMenuCustom struct {
	RoleId     *int64        `db:"role_id" json:"roleId"`                   //bigint 角色编号
	MenuIds    null.JSON     `db:"menu_ids" json:"menuIds"`                 //bigint 菜单编号
	Deleted    *int32        `db:"deleted" json:"deleted"`                  //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`               //bigint 租户
	Creator    null.String   `db:"creator" json:"creator,omitempty"`        //varchar 创建者
	CreateTime null.DateTime `db:"create_time" json:"createTime,omitempty"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater,omitempty"`        //varchar 更新者
	UpdateTime null.DateTime `db:"update_time" json:"updateTime,omitempty"` //datetime 更新时间
}
