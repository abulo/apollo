package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemTenant 租户 system_tenant
type SystemTenant struct {
	Id              *int64        `db:"id,-" json:"id"`                           //bigint 租户编号,PRI
	Name            *string       `db:"name" json:"name"`                         //varchar 租户名称
	UserId          null.Int64    `db:"user_id" json:"userId"`                    //bigint 联系人ID
	ContactName     *string       `db:"contact_name" json:"contactName"`          //varchar 联系人
	ContactMobile   *string       `db:"contact_mobile" json:"contactMobile"`      //varchar 租户联系电话
	Status          *int32        `db:"status" json:"status"`                     //tinyint 状态（0正常 1停用）
	Domain          null.String   `db:"domain" json:"domain"`                     //varchar 域名
	ExpireDate      null.Date     `db:"expire_date" json:"expireDate"`            //date 过期时间
	AccountCount    *int64        `db:"account_count" json:"accountCount"`        //bigint 账号数量
	TenantPackageId *int64        `db:"tenant_package_id" json:"tenantPackageId"` //bigint 套餐编号
	Deleted         *int32        `db:"deleted" json:"deleted"`                   //tinyint 是否删除(0否 1是)
	Creator         null.String   `db:"creator" json:"creator"`                   //varchar 创建人
	CreateTime      null.DateTime `db:"create_time" json:"createTime"`            //datetime 创建时间
	Updater         null.String   `db:"updater" json:"updater"`                   //varchar 更新人
	UpdateTime      null.DateTime `db:"update_time" json:"updateTime"`            //datetime 更新时间
	Username        *string       `db:"username,-" json:"username,omitempty"`     //varchar 用户名称
	Password        *string       `db:"password,-" json:"password,omitempty"`     //varchar 用户密码
}
