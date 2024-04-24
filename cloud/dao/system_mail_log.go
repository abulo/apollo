package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemMailLog 邮件日志表 system_mail_log
type SystemMailLog struct {
	Id               *int64        `db:"id" json:"id"`                              //bigint 编号,PRI
	UserId           null.Int64    `db:"user_id" json:"userId"`                     //bigint 用户编号
	UserType         null.Int32    `db:"user_type" json:"userType"`                 //tinyint 用户类型
	ToMail           *string       `db:"to_mail" json:"toMail"`                     //varchar 接收邮箱地址
	AccountId        *int64        `db:"account_id" json:"accountId"`               //bigint 邮箱账号编号
	FromMail         *string       `db:"from_mail" json:"fromMail"`                 //varchar 发送邮箱地址
	TemplateId       *int64        `db:"template_id" json:"templateId"`             //bigint 模板编号
	TemplateCode     *string       `db:"template_code" json:"templateCode"`         //varchar 模板编码
	TemplateNickname null.String   `db:"template_nickname" json:"templateNickname"` //varchar 模版发送人名称
	TemplateTitle    *string       `db:"template_title" json:"templateTitle"`       //varchar 邮件标题
	TemplateContent  *string       `db:"template_content" json:"templateContent"`   //varchar 邮件内容
	TemplateParams   *string       `db:"template_params" json:"templateParams"`     //varchar 邮件参数
	SendStatus       *int32        `db:"send_status" json:"sendStatus"`             //tinyint 发送状态
	SendTime         null.DateTime `db:"send_time" json:"sendTime"`                 //datetime 发送时间
	SendMessageId    null.String   `db:"send_message_id" json:"sendMessageId"`      //varchar 发送返回的消息 ID
	SendException    null.String   `db:"send_exception" json:"sendException"`       //varchar 发送异常
	Creator          null.String   `db:"creator" json:"creator"`                    //varchar 创建者
	CreateTime       null.DateTime `db:"create_time" json:"createTime"`             //datetime 创建时间
	Updater          null.String   `db:"updater" json:"updater"`                    //varchar 更新者
	UpdateTime       null.DateTime `db:"update_time" json:"updateTime"`             //datetime 更新时间
	Deleted          *int32        `db:"deleted" json:"deleted"`                    //tinyint 是否删除
}
