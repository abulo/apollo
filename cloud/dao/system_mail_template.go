package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemMailTemplate 邮件模版表 system_mail_template
type SystemMailTemplate struct {
	Id         *int64        `db:"id,-" json:"id"`                //bigint 编号,PRI
	Name       *string       `db:"name" json:"name"`              //varchar 模板名称
	Code       *string       `db:"code" json:"code"`              //varchar 模板编码
	AccountId  *int64        `db:"account_id" json:"accountId"`   //bigint 发送的邮箱账号编号
	Nickname   null.String   `db:"nickname" json:"nickname"`      //varchar 发送人名称
	Title      *string       `db:"title" json:"title"`            //varchar 模板标题
	Content    *string       `db:"content" json:"content"`        //varchar 模板内容
	Params     null.JSON     `db:"params" json:"params"`          //json 参数数组
	Status     *int32        `db:"status" json:"status"`          //tinyint 开启状态
	Remark     null.String   `db:"remark" json:"remark"`          //varchar 备注
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建者
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新者
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 是否删除
}
