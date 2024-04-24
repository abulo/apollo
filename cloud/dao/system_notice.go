package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemNotice 通知公告表 system_notice
type SystemNotice struct {
	Id         *int64        `db:"id" json:"id"`                  //bigint 公告ID,PRI
	Title      *string       `db:"title" json:"title"`            //varchar 公告标题
	Content    *string       `db:"content" json:"content"`        //text 公告内容
	Type       *int32        `db:"type" json:"type"`              //tinyint 公告类型（1通知 2公告）
	Status     *int32        `db:"status" json:"status"`          //tinyint 公告状态（0正常 1关闭）
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
