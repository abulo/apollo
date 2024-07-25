package dao

import "github.com/abulo/ratel/v3/stores/null"

// PayChannel 支付渠道 pay_channel
type PayChannel struct {
	Id         *int64        `db:"id,-" json:"id"`                //bigint 商户编号,PRI
	Code       *string       `db:"code" json:"code"`              //varchar 渠道编码
	Status     *int32        `db:"status" json:"status"`          //tinyint 开启状态
	Remark     null.String   `db:"remark" json:"remark"`          //varchar 备注
	FeeRate    *float64      `db:"fee_rate" json:"feeRate"`       //double 渠道费率，单位：百分比
	AppId      *int64        `db:"app_id" json:"appId"`           //bigint 应用编号
	Config     null.JSON     `db:"config" json:"config"`          //json 支付渠道配置
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
