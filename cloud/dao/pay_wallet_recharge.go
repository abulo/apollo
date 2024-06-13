package dao

import "github.com/abulo/ratel/v3/stores/null"

// PayWalletRecharge 会员钱包充值 pay_wallet_recharge
type PayWalletRecharge struct {
	Id               *int64        `db:"id" json:"id"`                               //bigint 编号,PRI
	WalletId         *int64        `db:"wallet_id" json:"walletId"`                  //bigint 会员钱包id
	TotalPrice       *int64        `db:"total_price" json:"totalPrice"`              //bigint 用户实际到账余额，例如充 100 送 20，则该值是 120
	PayPrice         *int64        `db:"pay_price" json:"payPrice"`                  //bigint 实际支付金额
	BonusPrice       *int64        `db:"bonus_price" json:"bonusPrice"`              //bigint 钱包赠送金额
	PackageId        null.Int64    `db:"package_id" json:"packageId"`                //bigint 充值套餐编号
	PayStatus        *int32        `db:"pay_status" json:"payStatus"`                //tinyint 是否已支付：[0:未支付 1:已经支付过]
	PayOrderId       *int64        `db:"pay_order_id" json:"payOrderId"`             //bigint 支付订单编号
	PayChannelCode   *string       `db:"pay_channel_code" json:"payChannelCode"`     //varchar 支付成功的支付渠道
	PayTime          null.DateTime `db:"pay_time" json:"payTime"`                    //datetime 订单支付时间
	PayRefundId      null.Int64    `db:"pay_refund_id" json:"payRefundId"`           //bigint 支付退款单编号
	RefundTotalPrice null.Int64    `db:"refund_total_price" json:"refundTotalPrice"` //bigint 退款金额，包含赠送金额
	RefundPayPrice   null.Int64    `db:"refund_pay_price" json:"refundPayPrice"`     //bigint 退款支付金额
	RefundBonusPrice null.Int64    `db:"refund_bonus_price" json:"refundBonusPrice"` //bigint 退款钱包赠送金额
	RefundTime       null.DateTime `db:"refund_time" json:"refundTime"`              //datetime 退款时间
	RefundStatus     null.Int32    `db:"refund_status" json:"refundStatus"`          //tinyint 退款状态
	Deleted          *int32        `db:"deleted" json:"deleted"`                     //tinyint 删除
	TenantId         *int64        `db:"tenant_id" json:"tenantId"`                  //bigint 租户
	Creator          null.String   `db:"creator" json:"creator"`                     //varchar 创建人
	CreateTime       null.DateTime `db:"create_time" json:"createTime"`              //datetime 创建时间
	Updater          null.String   `db:"updater" json:"updater"`                     //varchar 更新人
	UpdateTime       null.DateTime `db:"update_time" json:"updateTime"`              //datetime 更新时间
}
