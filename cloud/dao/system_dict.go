package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemDict 字典数据表 system_dict
type SystemDict struct {
	Id         *int64        `db:"id" json:"id"`                  //bigint 字典编码,PRI
	Sort       *int32        `db:"sort" json:"sort"`              //int 字典排序
	Label      *string       `db:"label" json:"label"`            //varchar 字典标签
	Value      *string       `db:"value" json:"value"`            //varchar 字典键值
	DictType   *string       `db:"dict_type" json:"dictType"`     //varchar 字典类型
	Status     *int32        `db:"status" json:"status"`          //tinyint 状态（0正常 1停用）
	Remark     null.String   `db:"remark" json:"remark"`          //varchar 备注
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
