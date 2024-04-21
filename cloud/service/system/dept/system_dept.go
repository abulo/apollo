package dept

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_dept 部门

// SystemDeptDao 数据转换
func SystemDeptDao(item *SystemDeptObject) *dao.SystemDept {
	daoItem := &dao.SystemDept{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 部门ID
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 部门名称
	}
	if item != nil && item.ParentId != nil {
		daoItem.ParentId = item.ParentId // 父部门ID
	}
	if item != nil && item.Sort != nil {
		daoItem.Sort = item.Sort // 显示顺序
	}
	if item != nil && item.UserId != nil {
		daoItem.UserId = null.Int64From(item.GetUserId()) // 负责人
	}
	if item != nil && item.Phone != nil {
		daoItem.Phone = null.StringFrom(item.GetPhone()) // 联系电话
	}
	if item != nil && item.Email != nil {
		daoItem.Email = null.StringFrom(item.GetEmail()) // 邮件
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 部门状态
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 是否删除
	}
	if item != nil && item.TenantId != nil {
		daoItem.TenantId = item.TenantId // 租户ID
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

// SystemDeptProto 数据绑定
func SystemDeptProto(item dao.SystemDept) *SystemDeptObject {
	res := &SystemDeptObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.ParentId != nil {
		res.ParentId = item.ParentId
	}
	if item.Sort != nil {
		res.Sort = item.Sort
	}
	if item.UserId.IsValid() {
		res.UserId = item.UserId.Ptr()
	}
	if item.Phone.IsValid() {
		res.Phone = item.Phone.Ptr()
	}
	if item.Email.IsValid() {
		res.Email = item.Email.Ptr()
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Deleted != nil {
		res.Deleted = item.Deleted
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
