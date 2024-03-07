package dict

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// system_dict 字典数据表

// SystemDictDao 数据转换
func SystemDictDao(item *SystemDictObject) *dao.SystemDict {
	daoItem := &dao.SystemDict{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 字典编码
	}
	if item != nil && item.Sort != nil {
		daoItem.Sort = item.Sort // 字典排序
	}
	if item != nil && item.Label != nil {
		daoItem.Label = item.Label // 字典标签
	}
	if item != nil && item.Value != nil {
		daoItem.Value = item.Value // 字典键值
	}
	if item != nil && item.DictType != nil {
		daoItem.DictType = item.DictType // 字典类型
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态（0正常 1停用）
	}
	if item != nil && item.ColorType != nil {
		daoItem.ColorType = null.StringFrom(item.GetColorType()) // 颜色类型
	}
	if item != nil && item.CssClass != nil {
		daoItem.CssClass = null.StringFrom(item.GetCssClass()) // css 样式
	}
	if item != nil && item.Remark != nil {
		daoItem.Remark = null.StringFrom(item.GetRemark()) // 备注
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

// SystemDictProto 数据绑定
func SystemDictProto(item dao.SystemDict) *SystemDictObject {
	res := &SystemDictObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Sort != nil {
		res.Sort = item.Sort
	}
	if item.Label != nil {
		res.Label = item.Label
	}
	if item.Value != nil {
		res.Value = item.Value
	}
	if item.DictType != nil {
		res.DictType = item.DictType
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.ColorType.IsValid() {
		res.ColorType = item.ColorType.Ptr()
	}
	if item.CssClass.IsValid() {
		res.CssClass = item.CssClass.Ptr()
	}
	if item.Remark.IsValid() {
		res.Remark = item.Remark.Ptr()
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
