package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemUserTenant 系统租户用户 system_user_tenant
type SystemUserTenant struct {
	Id         *int64        `db:"id" json:"id"`                  //bigint 编号,PRI
	UserId     *int64        `db:"user_id" json:"userId"`         //bigint 系统用户 ID
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户 id
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
