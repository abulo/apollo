package file

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_file 文件管理

// SystemFileDao 数据转换
func SystemFileDao(item *SystemFileObject) *dao.SystemFile {
	daoItem := &dao.SystemFile{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.FileName != nil {
		daoItem.FileName = item.FileName // 文件名称
	}
	if item != nil && item.FileType != nil {
		daoItem.FileType = item.FileType // 文件类型
	}
	if item != nil && item.FileMimeType != nil {
		daoItem.FileMimeType = item.FileMimeType // Mime类型
	}
	if item != nil && item.FileSize != nil {
		daoItem.FileSize = item.FileSize // 文件大小
	}
	if item != nil && item.FilePath != nil {
		daoItem.FilePath = item.FilePath // 文件路径
	}
	if item != nil && item.TenantId != nil {
		daoItem.TenantId = item.TenantId // 租户
	}
	if item != nil && item.Creator != nil {
		daoItem.Creator = null.StringFrom(item.GetCreator()) // 创建人
	}
	if item != nil && item.CreateTime != nil {
		daoItem.CreateTime = null.DateTimeFrom(util.GrpcTime(item.CreateTime)) // 创建时间
	}
	if item != nil && item.Updater != nil {
		daoItem.Updater = null.StringFrom(item.GetUpdater()) // 更新人
	}
	if item != nil && item.UpdateTime != nil {
		daoItem.UpdateTime = null.DateTimeFrom(util.GrpcTime(item.UpdateTime)) // 更新时间
	}

	return daoItem
}

// SystemFileProto 数据绑定
func SystemFileProto(item dao.SystemFile) *SystemFileObject {
	res := &SystemFileObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.FileName != nil {
		res.FileName = item.FileName
	}
	if item.FileType != nil {
		res.FileType = item.FileType
	}
	if item.FileMimeType != nil {
		res.FileMimeType = item.FileMimeType
	}
	if item.FileSize != nil {
		res.FileSize = item.FileSize
	}
	if item.FilePath != nil {
		res.FilePath = item.FilePath
	}
	if item.TenantId != nil {
		res.TenantId = item.TenantId
	}
	if item.Creator.IsValid() {
		res.Creator = item.Creator.Ptr()
	}
	if item.CreateTime.IsValid() {
		res.CreateTime = timestamppb.New(*item.CreateTime.Ptr())
	}
	if item.Updater.IsValid() {
		res.Updater = item.Updater.Ptr()
	}
	if item.UpdateTime.IsValid() {
		res.UpdateTime = timestamppb.New(*item.UpdateTime.Ptr())
	}

	return res
}
