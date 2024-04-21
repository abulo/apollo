package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemUserDept 系统用户部门 system_user_dept
type SystemUserDept struct {
	Id         *int64        `db:"id" json:"id"`                  //bigint 编号,PRI
	UserId     *int64        `db:"user_id" json:"userId"`         //bigint 系统用户 id
	DeptId     *int64        `db:"dept_id" json:"deptId"`         //bigint 部门 id
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}

type SystemUserDeptCustom struct {
	UserId     *int64        `db:"user_id" json:"userId"`         //bigint 系统用户 id
	DeptIds    null.JSON     `db:"dept_ids" json:"deptIds"`       //bigint 部门 id
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
