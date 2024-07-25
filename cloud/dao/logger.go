package dao

import "github.com/abulo/ratel/v3/stores/null"

// type SystemLoginLogDelete struct {
// 	Ids null.JSON `json:"ids,omitempty"` // 登录日志编号
// }

// type SystemOperateLogDelete struct {
// 	Ids null.JSON `json:"ids,omitempty"` // 操作日志编号
// }

// type SystemEntryLogDelete struct {
// 	Ids null.JSON `json:"ids,omitempty"` // 访问日志编号
// }

type SystemEntryLogDrop struct {
	Ids       null.JSON     `json:"ids,omitempty"`       // 日志编号
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
