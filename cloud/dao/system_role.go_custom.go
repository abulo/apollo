package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemRoleDataScope 系统角色 system_role
type SystemRoleDataScope struct {
	Id            *int64        `db:"id" json:"id"`                         //bigint 角色编号,PRI
	DataScope     null.Int32    `db:"data_scope" json:"dataScope"`          //tinyint 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	DataScopeDept null.JSON     `db:"data_scope_dept" json:"dataScopeDept"` //json 数据范围(指定部门数组)
	Updater       null.String   `db:"updater" json:"updater"`               //varchar 更新者
	UpdateTime    null.DateTime `db:"update_time" json:"updateTime"`        //datetime 更新时间
}
