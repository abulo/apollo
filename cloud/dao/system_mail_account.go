package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemMailAccount 邮箱账号表 system_mail_account
type SystemMailAccount struct {
	Id         *int64        `db:"id" json:"id"`                  //bigint 主键,PRI
	Mail       *string       `db:"mail" json:"mail"`              //varchar 邮箱
	Username   *string       `db:"username" json:"username"`      //varchar 用户名
	Password   *string       `db:"password" json:"password"`      //varchar 密码
	Host       *string       `db:"host" json:"host"`              //varchar SMTP 服务器域名
	Port       *int32        `db:"port" json:"port"`              //int SMTP 服务器端口
	SslEnable  *int32        `db:"ssl_enable" json:"sslEnable"`   //tinyint 是否开启 SSL
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建者
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新者
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 是否删除
}
