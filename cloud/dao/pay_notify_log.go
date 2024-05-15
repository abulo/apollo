package dao

import "github.com/abulo/ratel/v3/stores/null"

// PayNotifyLog 支付通知日志 pay_notify_log
type PayNotifyLog struct {
	Id          *int64        `db:"id" json:"id"`                    //bigint 日志编号,PRI
	TaskId      *int64        `db:"task_id" json:"taskId"`           //bigint 通知任务编号
	NotifyTimes *int32        `db:"notify_times" json:"notifyTimes"` //tinyint 第几次被通知
	Response    null.JSON     `db:"response" json:"response"`        //json 请求参数
	Status      *int32        `db:"status" json:"status"`            //tinyint 通知状态
	Deleted     *int32        `db:"deleted" json:"deleted"`          //tinyint 删除
	TenantId    *int64        `db:"tenant_id" json:"tenantId"`       //bigint 租户
	Creator     null.String   `db:"creator" json:"creator"`          //varchar 创建人
	CreateTime  null.DateTime `db:"create_time" json:"createTime"`   //datetime 创建时间
	Updater     null.String   `db:"updater" json:"updater"`          //varchar 更新人
	UpdateTime  null.DateTime `db:"update_time" json:"updateTime"`   //datetime 更新时间
}
