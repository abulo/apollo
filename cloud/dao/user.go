package dao

import "github.com/abulo/ratel/v3/stores/null"

// User 用户 user
type User struct {
	Id         *int64        `db:"id" json:"id"`                  //bigint 用户编号,PRI
	Nickname   null.String   `db:"nickname" json:"nickname"`      //varchar 用户名称
	Avatar     null.String   `db:"avatar" json:"avatar"`          //varchar 用户头像
	Birthday   null.Date     `db:"birthday" json:"birthday"`      //date 用户生日
	Gender     null.Int32    `db:"gender" json:"gender"`          //tinyint 用户性别(0女/1男)
	Status     *int32        `db:"status" json:"status"`          //tinyint 用户状态(0正常/1锁定)
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
