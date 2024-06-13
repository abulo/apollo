package dao

import "github.com/abulo/ratel/v3/stores/null"

// PayWalletTransaction 会员钱包流水表 pay_wallet_transaction
type PayWalletTransaction struct {
	Id         *int64        `db:"id" json:"id"`                  //bigint 编号,PRI
	WalletId   *int64        `db:"wallet_id" json:"walletId"`     //bigint 会员钱包 id
	BizType    *int32        `db:"biz_type" json:"bizType"`       //tinyint 关联类型
	BizId      *string       `db:"biz_id" json:"bizId"`           //varchar 关联业务编号
	No         *string       `db:"no" json:"no"`                  //varchar 流水号
	Title      *string       `db:"title" json:"title"`            //varchar 流水标题
	Price      *int64        `db:"price" json:"price"`            //bigint 交易金额, 单位分
	Balance    *int64        `db:"balance" json:"balance"`        //bigint 余额, 单位分
	Deleted    *int32        `db:"deleted" json:"deleted"`        //tinyint 删除
	TenantId   *int64        `db:"tenant_id" json:"tenantId"`     //bigint 租户
	Creator    null.String   `db:"creator" json:"creator"`        //varchar 创建人
	CreateTime null.DateTime `db:"create_time" json:"createTime"` //datetime 创建时间
	Updater    null.String   `db:"updater" json:"updater"`        //varchar 更新人
	UpdateTime null.DateTime `db:"update_time" json:"updateTime"` //datetime 更新时间
}
