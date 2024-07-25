package tenant

import (
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/system/tenant"
	"cloud/service/system/user"
	"context"
	"encoding/json"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// SystemTenantUserList 列表数据
func SystemTenantUserList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:system_user:SystemTenantUserList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	id := cast.ToInt64(newCtx.Param("id"))

	clientTenant := tenant.NewSystemTenantServiceClient(grpcClient)
	requestTenant := &tenant.SystemTenantRequest{}
	requestTenant.Id = id
	// 执行服务
	resTenant, err := clientTenant.SystemTenant(ctx, requestTenant)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestTenant,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenant")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	tenantItem := tenant.SystemTenantDao(resTenant.GetData())

	clientUser := user.NewSystemUserServiceClient(grpcClient)
	userDataScopeReq := &user.SystemUserRoleDataScopeRequest{}
	userDataScopeReq.UserId = *tenantItem.UserId.Ptr()
	userDataScopeReq.TenantId = id
	userRes, err := clientUser.SystemUserRoleDataScope(ctx, userDataScopeReq)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": userDataScopeReq,
			"err": err,
		}).Error("GrpcCall:部门:system_dept:SystemDeptList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	userScope := user.SystemUserRoleDataScopeDao(userRes.GetData())
	if len(userScope.DataScopeDept) < 1 {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.DeptError,
			"msg":  code.StatusText(code.DeptError),
		})
		return
	}
	//链接服务
	client := user.NewSystemUserServiceClient(grpcClient)
	// 构造查询条件
	request := &user.SystemUserListRequest{}
	requestTotal := &user.SystemUserListTotalRequest{}

	if val, ok := newCtx.GetQuery("username"); ok {
		request.Username = proto.String(val)      // 用户名称
		requestTotal.Username = proto.String(val) // 用户名称
	}

	request.TenantId = proto.Int64(id)      // 租户ID
	requestTotal.TenantId = proto.Int64(id) // 租户ID
	request.Deleted = proto.Int32(0)        // 删除状态
	requestTotal.Deleted = proto.Int32(0)   // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
			requestTotal.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 用户状态（0正常 1停用）
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 用户状态（0正常 1停用）
	}

	if val, ok := newCtx.GetQuery("deptId"); ok {
		request.DeptId = proto.Int64(cast.ToInt64(val))      // 用户状态（0正常 1停用）
		requestTotal.DeptId = proto.Int64(cast.ToInt64(val)) // 用户状态（0正常 1停用）
	}

	request.UserId = tenantItem.UserId.Ptr()      // 用户ID
	requestTotal.UserId = tenantItem.UserId.Ptr() // 用户ID

	request.DataScope = userScope.DataScope
	requestTotal.DataScope = userScope.DataScope
	dataScopeDept, _ := json.Marshal(userScope.DataScopeDept)
	request.DataScopeDept = dataScopeDept
	requestTotal.DataScopeDept = dataScopeDept

	// 执行服务,获取数据量
	resTotal, err := client.SystemUserListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestTotal,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemTenantUserList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var total int64
	paginationRequest := &pagination.PaginationRequest{}
	paginationRequest.PageNum = proto.Int64(cast.ToInt64(newCtx.Query("pageNum")))
	paginationRequest.PageSize = proto.Int64(cast.ToInt64(newCtx.Query("pageSize")))
	if resTotal.GetCode() == code.Success {
		total = resTotal.GetData()
	}
	// 执行服务
	res, err := client.SystemUserList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:system_user:SystemTenantUserList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemUser
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			userInfo := user.SystemUserDao(item)
			userInfo.Password = nil
			list = append(list, userInfo)
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"total":    total,
			"list":     list,
			"pageNum":  paginationRequest.PageNum,
			"pageSize": paginationRequest.PageSize,
		},
	})
}
