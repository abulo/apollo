package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemUserRoleCustom 系统用户和系统角色关联表 system_user_role
type SystemUserRoleCustom struct {
	SystemUserId  *int64        `db:"system_user_id" json:"systemUserId"`      //bigint 用户编号
	SystemRoleIds null.JSON     `db:"system_menu_ids" json:"systemRoleIds"`    //bigint 角色编号
	Creator       null.String   `db:"creator" json:"creator,omitempty"`        //varchar 创建者
	CreateTime    null.DateTime `db:"create_time" json:"createTime,omitempty"` //datetime 创建时间
	Updater       null.String   `db:"updater" json:"updater,omitempty"`        //varchar 更新者
	UpdateTime    null.DateTime `db:"update_time" json:"updateTime,omitempty"` //datetime 更新时间
}
