package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemDict 字典数据表 system_dict
type SystemDict struct {
	Id         *int64        `db:"id,-" json:"id"`                 //bigint 字典编码,PRI
	Sort       *int32        `db:"sort" json:"sort"`               //int 字典排序
	Label      *string       `db:"label" json:"label"`             //varchar 字典标签
	Value      *string       `db:"value" json:"value"`             //varchar 字典键值
	DictTypeId *int64        `db:"dict_type_id" json:"dictTypeId"` //bigint 字段类型 ID;
	Status     *int32        `db:"status" json:"status"`           //tinyint 状态（0正常 1停用）
	ColorType  null.String   `db:"color_type" json:"colorType"`    //varchar 颜色类型
	CssClass   null.String   `db:"css_class" json:"cssClass"`      //varchar css 样式
	Remark     null.String   `db:"remark" json:"remark"`           //varchar 备注
	Creator    null.String   `db:"creator" json:"creator"`         //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"`  //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`         //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"`  //datetime 更新时间
	DictType   null.String   `db:"dict_type,-" json:"dictType"`
}
