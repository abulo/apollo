package dao

import "github.com/abulo/ratel/v3/stores/null"

// PayApp 支付应用信息 pay_app
type PayApp struct {
	Id                *int64        `db:"id,-" json:"id"`                               //bigint 应用编号,PRI
	Name              *string       `db:"name" json:"name"`                             //varchar 应用名称
	Status            *int32        `db:"status" json:"status"`                         //tinyint 开启状态
	Remark            null.String   `db:"remark" json:"remark"`                         //varchar 备注
	OrderNotifyUrl    *string       `db:"order_notify_url" json:"orderNotifyUrl"`       //varchar 支付结果的回调地址
	RefundNotifyUrl   *string       `db:"refund_notify_url" json:"refundNotifyUrl"`     //varchar 退款结果的回调地址
	TransferNotifyUrl null.String   `db:"transfer_notify_url" json:"transferNotifyUrl"` //varchar 转账结果的回调地址
	Deleted           *int32        `db:"deleted" json:"deleted"`                       //tinyint 删除
	TenantId          *int64        `db:"tenant_id" json:"tenantId"`                    //bigint 租户
	Creator           null.String   `db:"creator" json:"creator"`                       //varchar 创建人
	CreateTime        null.DateTime `db:"create_time" json:"createTime"`                //datetime 创建时间
	Updater           null.String   `db:"updater" json:"updater"`                       //varchar 更新人
	UpdateTime        null.DateTime `db:"update_time" json:"updateTime"`                //datetime 更新时间
	ChannelList       null.JSON     `db:"-" json:"channelList"`                         // 支付渠道列表
}
