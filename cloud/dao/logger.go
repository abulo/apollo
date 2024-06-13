package dao

import "github.com/abulo/ratel/v3/stores/null"

type SystemLoginLogDelete struct {
	Ids null.JSON `json:"ids,omitempty"` // 登录日志编号
}

type SystemLoginLogDrop struct {
	Ids             null.JSON     `json:"ids,omitempty"`             // 日志编号
	Deleted         *int32        `json:"deleted,omitempty"`         // 删除
	Username        *string       `json:"username,omitempty"`        // 用户账号
	BeginLoginTime  null.DateTime `json:"beginLoginTime,omitempty"`  // 开始登录时间
	FinishLoginTime null.DateTime `json:"finishLoginTime,omitempty"` // 结束登录时间
	Channel         *string       `json:"channel,omitempty"`         // 渠道
}

type SystemOperateLogDelete struct {
	Ids null.JSON `json:"ids,omitempty"` // 操作日志编号
}

type SystemOperateLogDrop struct {
	Ids             null.JSON     `json:"ids,omitempty"`             // 日志编号
	Deleted         *int32        `json:"deleted,omitempty"`         // 删除
	Username        *string       `json:"username,omitempty"`        // 用户账号
	Module          *string       `json:"module,omitempty"`          // 模块名称
	BeginStartTime  null.DateTime `json:"beginStartTime,omitempty"`  // 开始操作时间
	FinishStartTime null.DateTime `json:"finishStartTime,omitempty"` // 结束操作时间
	Result          *int32        `json:"result,omitempty"`          // 结果(0 成功/1 失败)
}

type SystemEntryLogDelete struct {
	Ids null.JSON `json:"ids,omitempty"` // 访问日志编号
}

type SystemEntryLogDrop struct {
	StartTime null.DateTime `json:"startTime,omitempty"` // 开始时间
	EndTime   null.DateTime `json:"endTime,omitempty"`   // 结束时间
}

type SystemEntryLog struct {
	Id        string `json:"id,omitempty"`
	Host      string `json:"host,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	File      string `json:"file,omitempty"`
	Func      string `json:"func,omitempty"`
	Message   string `json:"message,omitempty"`
	Data      any    `json:"data,omitempty"`
	Level     string `json:"level,omitempty"`
}
