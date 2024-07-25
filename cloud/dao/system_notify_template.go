package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemNotifyTemplate 站内信模板表 system_notify_template
type SystemNotifyTemplate struct {
	Id         *int64        `db:"id,-" json:"id"`                //bigint 主键,PRI
	Name       *string       `db:"name" json:"name"`              //varchar 模板名称
	Code       *string       `db:"code" json:"code"`              //varchar 模版编码
	Nickname   *string       `db:"nickname" json:"nickname"`      //varchar 发送人名称
	Content    *string       `db:"content" json:"content"`        //varchar 模版内容
	Type       *int32        `db:"type" json:"type"`              //tinyint 类型
	Params     null.JSON     `db:"params" json:"params"`          //json 参数数组
	Status     *int32        `db:"status" json:"status"`          //tinyint 状态
	Remark     null.String   `db:"remark" json:"remark"`          //varchar 备注
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
