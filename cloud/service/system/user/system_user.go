package user

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_user 系统用户

// SystemUserDao 数据转换
func SystemUserDao(item *SystemUserObject) *dao.SystemUser {
	daoItem := &dao.SystemUser{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 用户编号
	}
	if item != nil && item.Nickname != nil {
		daoItem.Nickname = null.StringFrom(item.GetNickname()) // 昵称
	}
	if item != nil && item.Mobile != nil {
		daoItem.Mobile = null.StringFrom(item.GetMobile()) // 手机号码
	}
	if item != nil && item.Username != nil {
		daoItem.Username = item.Username // 用户名称
	}
	if item != nil && item.Password != nil {
		daoItem.Password = item.Password // 用户密码
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 用户状态（0正常 1停用）
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 是否删除(0否 1是)
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
	// if item != nil && item.DeptIds != nil {
	// 	daoItem.DeptIds = null.JSONFrom(item.GetDeptIds())
	// }
	// if item != nil && item.PostIds != nil {
	// 	daoItem.PostIds = null.JSONFrom(item.GetPostIds())
	// }
	// if item != nil && item.RoleIds != nil {
	// 	daoItem.RoleIds = null.JSONFrom(item.GetRoleIds())
	// }

	return daoItem
}

// SystemUserProto 数据绑定
func SystemUserProto(item dao.SystemUser) *SystemUserObject {
	res := &SystemUserObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Nickname.IsValid() {
		res.Nickname = item.Nickname.Ptr()
	}
	if item.Mobile.IsValid() {
		res.Mobile = item.Mobile.Ptr()
	}
	if item.Username != nil {
		res.Username = item.Username
	}
	if item.Password != nil {
		res.Password = item.Password
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
	// if item.DeptIds.IsValid() {
	// 	res.DeptIds = *item.DeptIds.Ptr()
	// }
	// if item.PostIds.IsValid() {
	// 	res.PostIds = *item.PostIds.Ptr()
	// }
	// if item.RoleIds.IsValid() {
	// 	res.RoleIds = *item.RoleIds.Ptr()
	// }

	return res
}
