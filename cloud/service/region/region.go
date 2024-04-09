package region

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// region 地区表

// RegionDao 数据转换
func RegionDao(item *RegionObject) *dao.Region {
	daoItem := &dao.Region{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 区域编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 区域名称
	}
	if item != nil && item.ParentId != nil {
		daoItem.ParentId = item.ParentId // 父级编号
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态;
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 是否删除;
	}
	if item != nil && item.Sort != nil {
		daoItem.Sort = item.Sort // 排序;
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

// RegionProto 数据绑定
func RegionProto(item dao.Region) *RegionObject {
	res := &RegionObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.ParentId != nil {
		res.ParentId = item.ParentId
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Deleted != nil {
		res.Deleted = item.Deleted
	}
	if item.Sort != nil {
		res.Sort = item.Sort
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
