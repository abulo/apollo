package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemDictType 字典类型 system_dict_type
type SystemDictType struct {
	Id         *int64        `db:"id,-" json:"id"`                //bigint 字典主键,PRI
	Name       *string       `db:"name" json:"name"`              //varchar 字典名称
	Type       *string       `db:"type" json:"type"`              //varchar 字典类型
	Status     *int32        `db:"status" json:"status"`          //tinyint 状态（0正常 1停用）
	Remark     null.String   `db:"remark" json:"remark"`          //varchar 备注
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
