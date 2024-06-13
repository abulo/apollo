package dao

import "github.com/abulo/ratel/v3/stores/null"

// PayWallet 会员钱包表 pay_wallet
type PayWallet struct {
	Id            *int64        `db:"id" json:"id"`                        //bigint 编号,PRI
	UserId        *int64        `db:"user_id" json:"userId"`               //bigint 用户编号
	UserType      *int32        `db:"user_type" json:"userType"`           //tinyint 用户类型
	Balance       *int64        `db:"balance" json:"balance"`              //bigint 余额，单位分
	TotalExpense  *int64        `db:"total_expense" json:"totalExpense"`   //bigint 累计支出，单位分
	TotalRecharge *int64        `db:"total_recharge" json:"totalRecharge"` //bigint 累计充值，单位分
	FreezePrice   *int64        `db:"freeze_price" json:"freezePrice"`     //bigint 冻结金额，单位分
	Deleted       *int32        `db:"deleted" json:"deleted"`              //tinyint 删除
	TenantId      *int64        `db:"tenant_id" json:"tenantId"`           //bigint 租户
	Creator       null.String   `db:"creator" json:"creator"`              //varchar 创建人
	CreateTime    null.DateTime `db:"create_time" json:"createTime"`       //datetime 创建时间
	Updater       null.String   `db:"updater" json:"updater"`              //varchar 更新人
	UpdateTime    null.DateTime `db:"update_time" json:"updateTime"`       //datetime 更新时间
}
