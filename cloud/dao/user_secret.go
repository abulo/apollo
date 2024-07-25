package dao

import "github.com/abulo/ratel/v3/stores/null"

// UserSecret 用户密码 user_secret
type UserSecret struct {
	Id         *int64        `db:"id,-" json:"id"`                //bigint 编号,PRI
	UserId     *int64        `db:"user_id" json:"userId"`         //bigint 用户编号
	Password   null.String   `db:"password" json:"password"`      //varchar 用户密码
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
