package dao

import "github.com/abulo/ratel/v3/stores/null"

// PayNotifyTask 商户支付-退款等的通知 pay_notify_task
type PayNotifyTask struct {
	Id                 *int64        `db:"id" json:"id"`                                   //bigint 任务编号,PRI
	AppId              *int64        `db:"app_id" json:"appId"`                            //bigint 应用编号
	Type               *int32        `db:"type" json:"type"`                               //tinyint 通知类型
	DataId             *int64        `db:"data_id" json:"dataId"`                          //bigint 数据编号
	Status             *int32        `db:"status" json:"status"`                           //tinyint 通知状态
	MerchantOrderId    *string       `db:"merchant_order_id" json:"merchantOrderId"`       //varchar 商户订单编号
	MerchantTransferId *string       `db:"merchant_transfer_id" json:"merchantTransferId"` //varchar 商户转账单编号
	NextNotifyTime     null.DateTime `db:"next_notify_time" json:"nextNotifyTime"`         //datetime 下一次通知时间
	LastExecuteTime    null.DateTime `db:"last_execute_time" json:"lastExecuteTime"`       //datetime 最后一次执行时间
	NotifyTimes        *int32        `db:"notify_times" json:"notifyTimes"`                //tinyint 当前通知次数
	MaxNotifyTimes     null.Int32    `db:"max_notify_times" json:"maxNotifyTimes"`         //tinyint 最大可通知次数
	NotifyUrl          *string       `db:"notify_url" json:"notifyUrl"`                    //varchar 异步通知地址
	Deleted            *int32        `db:"deleted" json:"deleted"`                         //tinyint 删除
	TenantId           *int64        `db:"tenant_id" json:"tenantId"`                      //bigint 租户
	Creator            null.String   `db:"creator" json:"creator"`                         //varchar 创建人
	CreateTime         null.DateTime `db:"create_time" json:"createTime"`                  //datetime 创建时间
	Updater            null.String   `db:"updater" json:"updater"`                         //varchar 更新人
	UpdateTime         null.DateTime `db:"update_time" json:"updateTime"`                  //datetime 更新时间
}
