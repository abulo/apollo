package dao

import "github.com/abulo/ratel/v3/stores/null"

// PayOrder 支付订单 pay_order
type PayOrder struct {
	Id              *int64             `db:"id" json:"id"`                             //bigint 支付订单编号,PRI
	AppId           *int64             `db:"app_id" json:"appId"`                      //bigint 应用编号
	ChannelId       *int64             `db:"channel_id" json:"channelId"`              //bigint 渠道编号
	ChannelCode     *string            `db:"channel_code" json:"channelCode"`          //varchar 渠道编码
	MerchantOrderId *string            `db:"merchant_order_id" json:"merchantOrderId"` //varchar 商户订单编号
	Subject         *string            `db:"subject" json:"subject"`                   //varchar 商品标题
	Body            *string            `db:"body" json:"body"`                         //varchar 商品描述
	NotifyUrl       *string            `db:"notify_url" json:"notifyUrl"`              //varchar 异步通知地址
	Price           *int64             `db:"price" json:"price"`                       //bigint 支付金额，单位：分
	ChannelFeeRate  *float64           `db:"channel_fee_rate" json:"channelFeeRate"`   //double 渠道手续费，单位：百分比
	ChannelFeePrice *int64             `db:"channel_fee_price" json:"channelFeePrice"` //bigint 渠道手续金额，单位：分
	Status          *int32             `db:"status" json:"status"`                     //tinyint 支付状态
	UserIp          null.String        `db:"user_ip" json:"userIp"`                    //varchar 用户 IP
	ExpireTime      null.DateTime      `db:"expire_time" json:"expireTime"`            //datetime 订单失效时间
	SuccessTime     null.DateTime      `db:"success_time" json:"successTime"`          //datetime 订单支付成功时间
	ExtensionId     null.Int64         `db:"extension_id" json:"extensionId"`          //bigint 支付成功的订单拓展单编号
	No              null.String        `db:"no" json:"no"`                             //varchar 支付订单号
	RefundPrice     *int64             `db:"refund_price" json:"refundPrice"`          //bigint 退款总金额，单位：分
	ChannelUserId   null.String        `db:"channel_user_id" json:"channelUserId"`     //varchar 渠道用户编号
	ChannelOrderNo  null.String        `db:"channel_order_no" json:"channelOrderNo"`   //varchar 渠道订单号
	Deleted         *int32             `db:"deleted" json:"deleted"`                   //tinyint 删除
	TenantId        *int64             `db:"tenant_id" json:"tenantId"`                //bigint 租户
	Creator         null.String        `db:"creator" json:"creator"`                   //varchar 创建人
	CreateTime      null.DateTime      `db:"create_time" json:"createTime"`            //datetime 创建时间
	Updater         null.String        `db:"updater" json:"updater"`                   //varchar 更新人
	UpdateTime      null.DateTime      `db:"update_time" json:"updateTime"`            //datetime 更新时间
	Extension       *PayOrderExtension `db:"-" json:"extension,omitempty"`             // 支付订单拓展
}
