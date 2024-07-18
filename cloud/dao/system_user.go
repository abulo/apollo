package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemUser 系统用户 system_user
type SystemUser struct {
	SystemUserCustom
	DeptIds null.JSON `db:"dept_ids" json:"deptIds"` // 部门ID
	RoleIds null.JSON `db:"role_ids" json:"roleIds"` // 角色ID
	PostIds null.JSON `db:"post_ids" json:"postIds"` // 岗位ID
}

type SystemUserCustom struct {
	Id         *int64        `db:"id" json:"id"`                  //bigint 用户编号,PRI
	Nickname   null.String   `db:"nickname" json:"nickname"`      //varchar 昵称
	Mobile     null.String   `db:"mobile" json:"mobile"`          //varchar 手机号码
	Username   *string       `db:"username" json:"username"`      //varchar 用户名称
	Password   *string       `db:"password" json:"password"`      //varchar 用户密码
	Status     *int32        `db:"status" json:"status"`          //tinyint 用户状态（0正常 1停用）
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 是否删除(0否 1是)
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户ID
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}

type SystemUserRoleDataScope struct {
	DataScope     *int32  `db:"data_scope" json:"dataScope"`          //tinyint 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	DataScopeDept []int64 `db:"data_scope_dept" json:"dataScopeDept"` //json 数据范围(指定部门数组)
}
