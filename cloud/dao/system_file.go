package dao

import "github.com/abulo/ratel/v3/stores/null"

// SystemFile 文件管理 system_file
type SystemFile struct {
	Id           *int64        `db:"id" json:"id"`                       //bigint 编号,PRI
	FileName     *string       `db:"file_name" json:"fileName"`          //varchar 文件名称
	FileType     *int32        `db:"file_type" json:"fileType"`          //tinyint 文件类型
	FileMimeType *string       `db:"file_mime_type" json:"fileMimeType"` //varchar Mime类型
	FileSize     *int64        `db:"file_size" json:"fileSize"`          //bigint 文件大小
	FilePath     *string       `db:"file_path" json:"filePath"`          //varchar 文件路径
	TenantId     *int64        `db:"tenant_id" json:"tenantId"`          //bigint 租户
	Creator      null.String   `db:"creator" json:"creator"`             //varchar 创建人
	CreateTime   null.DateTime `db:"create_time" json:"createTime"`      //datetime 创建时间
	Updater      null.String   `db:"updater" json:"updater"`             //varchar 更新人
	UpdateTime   null.DateTime `db:"update_time" json:"updateTime"`      //datetime 更新时间
}
