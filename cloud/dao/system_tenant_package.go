package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemTenantPackage 租户套餐包 system_tenant_package
type SystemTenantPackage struct {
	Id            *int64        `db:"id" json:"id"`                         //bigint 套餐编号,PRI
	Name          *string       `db:"name" json:"name"`                     //varchar 套餐名称
	Status        *int32        `db:"status" json:"status"`                 //tinyint 状态（0正常 1停用）
	SystemMenuIds null.JSON     `db:"system_menu_ids" json:"systemMenuIds"` //json 目录编号
	Remark        null.String   `db:"remark" json:"remark"`                 //varchar 备注
	Deleted       *int32        `db:"deleted" json:"deleted"`               //tinyint 是否删除(0否 1是)
	Creator       null.String   `db:"creator" json:"creator"`               //varchar 创建人
	CreateTime    null.DateTime `db:"create_time" json:"createTime"`        //datetime 创建时间
	Updater       null.String   `db:"updater" json:"updater"`               //varchar 更新人
	UpdateTime    null.DateTime `db:"update_time" json:"updateTime"`        //datetime 更新时间
}
