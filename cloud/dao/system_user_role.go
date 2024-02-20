package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemUserRole 系统用户和系统角色关联表 system_user_role
type SystemUserRole struct {
	Id           *int64        `db:"id" json:"id"`                       //bigint 自增编号,PRI
	SystemUserId *int64        `db:"system_user_id" json:"systemUserId"` //bigint 用户编号
	SystemRoleId *int64        `db:"system_role_id" json:"systemRoleId"` //bigint 角色编号
	Creator      null.String   `db:"creator" json:"creator"`             //varchar 创建者
	CreateTime   null.DateTime `db:"create_time" json:"createTime"`      //datetime 创建时间
	Updater      null.String   `db:"updater" json:"updater"`             //varchar 更新者
	UpdateTime   null.DateTime `db:"update_time" json:"updateTime"`      //datetime 更新时间
	Deleted      *int32        `db:"deleted" json:"deleted"`             //tinyint 是否删除
}
