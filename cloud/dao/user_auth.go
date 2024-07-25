package dao

import "github.com/abulo/ratel/v3/stores/null"

// UserAuth 用户三方登录授权 user_auth
type UserAuth struct {
	Id           *int64        `db:"id,-" json:"id"`                    //bigint 编号,PRI
	UserId       *int64        `db:"user_id" json:"userId"`             //bigint 用户编号
	IdentityType null.String   `db:"identity_type" json:"identityType"` //varchar 登录类型(手机号/邮箱) 或第三方应用名称 (微信/微博等)
	Identifier   null.String   `db:"identifier" json:"identifier"`      //varchar 手机号/邮箱/第三方的唯一标识
	Credential   null.String   `db:"credential" json:"credential"`      //varchar 密码凭证 (自建账号的保存密码, 第三方的保存 token)
	Creator      null.String   `db:"creator" json:"creator"`            //varchar 创建人
	CreateTime   null.DateTime `db:"create_time" json:"createTime"`     //datetime 创建时间
	Updater      null.String   `db:"updater" json:"updater"`            //varchar 更新人
	UpdateTime   null.DateTime `db:"update_time" json:"updateTime"`     //datetime 更新时间
}
