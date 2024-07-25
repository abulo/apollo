package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemNotifyMessage 站内信消息表 system_notify_message
type SystemNotifyMessage struct {
	Id               *int64        `db:"id,-" json:"id"`                            //bigint 消息,PRI
	UserId           *int64        `db:"user_id" json:"userId"`                     //bigint 用户id
	UserType         *int32        `db:"user_type" json:"userType"`                 //tinyint 用户类型
	TemplateId       *int64        `db:"template_id" json:"templateId"`             //bigint 模版编号
	TemplateCode     *string       `db:"template_code" json:"templateCode"`         //varchar 模板编码
	TemplateNickname *string       `db:"template_nickname" json:"templateNickname"` //varchar 模版发送人名称
	TemplateContent  *string       `db:"template_content" json:"templateContent"`   //varchar 模版内容
	TemplateType     *int32        `db:"template_type" json:"templateType"`         //int 模版类型
	TemplateParams   null.JSON     `db:"template_params" json:"templateParams"`     //json 模版参数
	ReadStatus       *int32        `db:"read_status" json:"readStatus"`             //tinyint 是否已读
	ReadTime         null.DateTime `db:"read_time" json:"readTime"`                 //datetime 阅读时间
	Deleted          *int32        `db:"deleted" json:"deleted"`                    //tinyint 删除
	TenantId         *int64        `db:"tenant_id" json:"tenantId"`                 //bigint 租户
	Creator          null.String   `db:"creator" json:"creator"`                    //varchar 创建人
	CreateTime       null.DateTime `db:"create_time" json:"createTime"`             //datetime 创建时间
	Updater          null.String   `db:"updater" json:"updater"`                    //varchar 更新人
	UpdateTime       null.DateTime `db:"update_time" json:"updateTime"`             //datetime 更新时间
}

// SystemNotifyMessageMultiple  多选
type SystemNotifyMessageMultiple struct {
	Ids null.JSON `json:"ids,omitempty"` // 日志编号
}
