package dao

import "github.com/abulo/ratel/v3/stores/null"

// Region 地区表 region
type Region struct {
	Id          *int64        `db:"id" json:"id"`                    //bigint 区域编号,PRI
	Name        *string       `db:"name" json:"name"`                //varchar 区域名称
	ParentId    *int64        `db:"parent_id" json:"parentId"`       //bigint 父级编号
	WeatherCode null.String   `db:"weather_code" json:"weatherCode"` //varchar 天气code
	Creator     null.String   `db:"creator" json:"creator"`          //varchar 创建人
	CreateTime  null.DateTime `db:"create_time" json:"createTime"`   //datetime 创建时间
	Updater     null.String   `db:"updater" json:"updater"`          //varchar 更新人
	UpdateTime  null.DateTime `db:"update_time" json:"updateTime"`   //datetime 更新时间
	Children    []*Region     `json:"children"`                      // 子级
}
