package dao

import "github.com/abulo/ratel/v3/stores/null"

// PayOrderExtension 支付订单拓展 pay_order_extension
type PayOrderExtension struct {
	Id                *int64        `db:"id" json:"id"`                                 //bigint 编号,PRI
	No                *string       `db:"no" json:"no"`                                 //varchar 支付订单号
	OrderId           *int64        `db:"order_id" json:"orderId"`                      //bigint 支付订单编号
	ChannelId         *int64        `db:"channel_id" json:"channelId"`                  //bigint 渠道编号
	ChannelCode       *string       `db:"channel_code" json:"channelCode"`              //varchar 渠道编码
	UserIp            null.String   `db:"user_ip" json:"userIp"`                        //varchar ip
	Status            *int32        `db:"status" json:"status"`                         //tinyint 支付状态
	ChannelExtras     null.String   `db:"channel_extras" json:"channelExtras"`          //varchar 支付渠道的额外参数
	ChannelErrorCode  null.String   `db:"channel_error_code" json:"channelErrorCode"`   //varchar 渠道调用报错时，错误码
	ChannelErrorMsg   null.String   `db:"channel_error_msg" json:"channelErrorMsg"`     //varchar 渠道调用报错时，错误信息
	ChannelNotifyData null.String   `db:"channel_notify_data" json:"channelNotifyData"` //varchar 支付渠道异步通知的内容
	Deleted           *int32        `db:"deleted" json:"deleted"`                       //tinyint 删除
	TenantId          *int64        `db:"tenant_id" json:"tenantId"`                    //bigint 租户
	Creator           null.String   `db:"creator" json:"creator"`                       //varchar 创建人
	CreateTime        null.DateTime `db:"create_time" json:"createTime"`                //datetime 创建时间
	Updater           null.String   `db:"updater" json:"updater"`                       //varchar 更新人
	UpdateTime        null.DateTime `db:"update_time" json:"updateTime"`                //datetime 更新时间
}
