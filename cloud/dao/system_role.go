package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemRole 系统角色 system_role
type SystemRole struct {
	Id            *int64        `db:"id" json:"id"`                         //bigint 角色编号,PRI
	Name          *string       `db:"name" json:"name"`                     //varchar 角色名称
	Code          *string       `db:"code" json:"code"`                     //varchar 角色权限字符串
	Sort          *int32        `db:"sort" json:"sort"`                     //int 显示顺序
	DataScope     *int32        `db:"data_scope" json:"dataScope"`          //tinyint 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	DataScopeDept null.JSON     `db:"data_scope_dept" json:"dataScopeDept"` //json 数据范围(指定部门数组)
	Status        *int32        `db:"status" json:"status"`                 //tinyint 角色状态（0正常 1停用）
	Type          *int32        `db:"type" json:"type"`                     //tinyint 角色类型(1内置/2定义)
	Remark        null.String   `db:"remark" json:"remark"`                 //varchar 备注
	Deleted       *int32        `db:"deleted" json:"deleted"`               //tinyint 删除
	TenantId      *int64        `db:"tenant_id" json:"tenantId"`            //bigint 租户
	Creator       null.String   `db:"creator" json:"creator"`               //varchar 创建者
	CreateTime    null.DateTime `db:"create_time" json:"createTime"`        //datetime 创建时间
	Updater       null.String   `db:"updater" json:"updater"`               //varchar 更新者
	UpdateTime    null.DateTime `db:"update_time" json:"updateTime"`        //datetime 更新时间
}
