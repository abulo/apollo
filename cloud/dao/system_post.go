package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemPost 职位 system_post
type SystemPost struct {
	Id             *int64        `db:"id" json:"id"`                           //bigint 职位ID,PRI
	Name           *string       `db:"name" json:"name"`                       //varchar 职位名称
	Sort           *int32        `db:"sort" json:"sort"`                       //int 显示顺序
	Status         *int32        `db:"status" json:"status"`                   //tinyint 状态
	Deleted        *int32        `db:"deleted" json:"deleted"`                 //tinyint 是否删除
	SystemTenantId *int64        `db:"system_tenant_id" json:"systemTenantId"` //bigint 租户ID
	Creator        null.String   `db:"creator" json:"creator"`                 //varchar 创建人
	CreateTime     null.DateTime `db:"create_time" json:"createTime"`          //datetime 创建时间
	Updater        null.String   `db:"updater" json:"updater"`                 //varchar 更新人
	UpdateTime     null.DateTime `db:"update_time" json:"updateTime"`          //datetime 更新时间
}
