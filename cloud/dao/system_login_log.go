package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemLoginLog 登录日志 system_login_log
type SystemLoginLog struct {
	Id         *int64        `db:"id,-" json:"id"`                //bigint 主键,PRI
	Username   *string       `db:"username" json:"username"`      //varchar 用户账号
	UserIp     *string       `db:"user_ip" json:"userIp"`         //varchar 用户ip
	UserAgent  null.String   `db:"user_agent" json:"userAgent"`   //varchar UA
	LoginTime  null.DateTime `db:"login_time" json:"loginTime"`   //datetime 登录时间
	Channel    *string       `db:"channel" json:"channel"`        //varchar 渠道
	UserId     *int64        `db:"user_id" json:"userId"`         //bigint
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}

// SystemLoginLogMultiple  多选
type SystemLoginLogMultiple struct {
	Ids null.JSON `json:"ids,omitempty"` // 日志编号
}
