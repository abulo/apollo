package dao

import "github.com/abulo/ratel/v3/stores/null"

// PayWalletRechargePackage 充值套餐表 pay_wallet_recharge_package
type PayWalletRechargePackage struct {
	Id         *int64        `db:"id" json:"id"`                  //bigint 编号,PRI
	Name       *string       `db:"name" json:"name"`              //varchar 套餐名称
	PayPrice   *int64        `db:"pay_price" json:"payPrice"`     //bigint 支付金额
	BonusPrice *int64        `db:"bonus_price" json:"bonusPrice"` //bigint 赠送金额
	Status     *int32        `db:"status" json:"status"`          //tinyint 状态
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
