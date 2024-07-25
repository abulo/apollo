package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemDept 部门 system_dept
type SystemDept struct {
	Id         *int64        `db:"id,-" json:"id"`                //bigint 部门ID,PRI
	Name       *string       `db:"name" json:"name"`              //varchar 部门名称
	ParentId   *int64        `db:"parent_id" json:"parentId"`     //bigint 父部门ID
	Sort       *int32        `db:"sort" json:"sort"`              //int 显示顺序
	UserId     null.Int64    `db:"user_id" json:"userId"`         //bigint 负责人
	Phone      null.String   `db:"phone" json:"phone"`            //varchar 联系电话
	Email      null.String   `db:"email" json:"email"`            //varchar 邮件
	Status     *int32        `db:"status" json:"status"`          //tinyint 部门状态
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 是否删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户ID
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
	Children   []*SystemDept `db:"-" json:"children"`
}
