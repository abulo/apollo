package dao

import "github.com/abulo/ratel/v3/stores/null"

// PayRefund 退款订单 pay_refund
type PayRefund struct {
	Id                *int64        `db:"id" json:"id"`                                 //bigint 支付退款编号,PRI
	No                *string       `db:"no" json:"no"`                                 //varchar 退款单号
	AppId             *int64        `db:"app_id" json:"appId"`                          //bigint 应用编号
	ChannelId         *int64        `db:"channel_id" json:"channelId"`                  //bigint 渠道编号
	ChannelCode       *string       `db:"channel_code" json:"channelCode"`              //varchar 渠道编码
	OrderId           *int64        `db:"order_id" json:"orderId"`                      //bigint 支付订单编号 pay_order 表id
	OrderNo           *string       `db:"order_no" json:"orderNo"`                      //varchar 支付订单 no
	MerchantOrderId   *string       `db:"merchant_order_id" json:"merchantOrderId"`     //varchar 商户订单编号（商户系统生成）
	MerchantRefundId  *string       `db:"merchant_refund_id" json:"merchantRefundId"`   //varchar 商户退款订单号（商户系统生成）
	NotifyUrl         *string       `db:"notify_url" json:"notifyUrl"`                  //varchar 异步通知商户地址
	Status            *int32        `db:"status" json:"status"`                         //tinyint 退款状态
	PayPrice          *int64        `db:"pay_price" json:"payPrice"`                    //bigint 支付金额,单位分
	RefundPrice       *int64        `db:"refund_price" json:"refundPrice"`              //bigint 退款金额,单位分
	Reason            *string       `db:"reason" json:"reason"`                         //varchar 退款原因
	UserIp            null.String   `db:"user_ip" json:"userIp"`                        //varchar ip
	ChannelOrderNo    *string       `db:"channel_order_no" json:"channelOrderNo"`       //varchar 渠道订单号，pay_order 中的 channel_order_no 对应
	ChannelRefundNo   null.String   `db:"channel_refund_no" json:"channelRefundNo"`     //varchar 渠道退款单号，渠道返
	SuccessTime       null.DateTime `db:"success_time" json:"successTime"`              //datetime 退款成功时间
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
