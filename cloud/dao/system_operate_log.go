package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemOperateLog 操作日志 system_operate_log
type SystemOperateLog struct {
	Id            *int64        `db:"id" json:"id"`                        //bigint 主键,PRI
	Username      *string       `db:"username" json:"username"`            //varchar 用户账号
	UserId        *int64        `db:"user_id" json:"userId"`               //bigint 用户id
	Module        *string       `db:"module" json:"module"`                //varchar 模块名称
	RequestMethod *string       `db:"request_method" json:"requestMethod"` //varchar 请求方法名
	RequestUrl    *string       `db:"request_url" json:"requestUrl"`       //varchar 请求地址
	UserIp        *string       `db:"user_ip" json:"userIp"`               //varchar 用户 ip
	UserAgent     null.String   `db:"user_agent" json:"userAgent"`         //varchar UA
	GoMethod      *string       `db:"go_method" json:"goMethod"`           //varchar 方法名
	GoMethodArgs  null.JSON     `db:"go_method_args" json:"goMethodArgs"`  //json 方法的参数
	StartTime     null.DateTime `db:"start_time" json:"startTime"`         //datetime 操作开始时间
	Duration      *int32        `db:"duration" json:"duration"`            //int 执行时长
	Channel       *string       `db:"channel" json:"channel"`              //varchar 渠道
	Result        *int32        `db:"result" json:"result"`                //tinyint 结果(0 成功/1 失败)
	Deleted       *int32        `db:"deleted" json:"deleted"`              //tinyint 删除
	TenantId      *int64        `db:"tenant_id" json:"tenantId"`           //bigint 租户
	Creator       null.String   `db:"creator" json:"creator"`              //varchar 创建人
	CreateTime    null.DateTime `db:"create_time" json:"createTime"`       //datetime 创建时间
	Updater       null.String   `db:"updater" json:"updater"`              //varchar 更新人
	UpdateTime    null.DateTime `db:"update_time" json:"updateTime"`       //datetime 更新时间
}
