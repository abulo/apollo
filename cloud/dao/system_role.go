package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemRole 系统角色 system_role
type SystemRole struct {
	Id         *int64        `db:"id" json:"id"`                  //bigint 角色编号,PRI
	Name       *string       `db:"name" json:"name"`              //varchar 角色名称
	Code       *string       `db:"code" json:"code"`              //varchar 角色权限字符串
	Sort       *int32        `db:"sort" json:"sort"`              //int 显示顺序
	Status     *int32        `db:"status" json:"status"`          //tinyint 角色状态（0正常 1停用）
	Type       *int32        `db:"type" json:"type"`              //tinyint 角色类型(1内置/2定义)
	Remark     null.String   `db:"remark" json:"remark"`          //varchar 备注
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建者
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新者
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
	MenuIds    null.JSON     `json:"menuIds"`                     //json 关联的菜单编号
}
