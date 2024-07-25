// system_tenant 租户
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemTenant } from "@/api/interface/systemTenant";
import { SystemUser } from "@/api/interface/systemUser";
import { SystemMenu } from "@/api/interface/systemMenu";
// 租户创建数据
export const addSystemTenantApi = (params: SystemTenant.ResSystemTenantItem) => {
  return http.post(PORT + `/system/tenant`, params);
};
// 租户更新数据
export const updateSystemTenantApi = (id: number, params: SystemTenant.ResSystemTenantItem) => {
  return http.put(PORT + `/system/tenant/${id}/update`, params);
};
// 租户删除数据
export const deleteSystemTenantApi = (id: number) => {
  return http.delete(PORT + `/system/tenant/${id}/delete`);
};
// 租户查询单条数据
export const getSystemTenantApi = (id: number) => {
  return http.get<SystemTenant.ResSystemTenantItem>(PORT + `/system/tenant/${id}/item`);
};
// 租户恢复数据
export const recoverSystemTenantApi = (id: number) => {
  return http.put(PORT + `/system/tenant/${id}/recover`);
};
// 租户清理数据
export const dropSystemTenantApi = (id: number) => {
  return http.delete(PORT + `/system/tenant/${id}/drop`);
};
// 租户列表数据
export const getSystemTenantListApi = (params?: SystemTenant.ReqSystemTenantList) => {
  return http.get<ResPage<SystemTenant.ResSystemTenantItem>>(PORT + `/system/tenant`, params);
};
// 租户列表精简数据
export const getSystemTenantListSimpleApi = (params?: SystemTenant.ReqSystemTenantList) => {
  return http.get<SystemTenant.ResSystemTenantItem[]>(PORT + `/system/tenant/simple`, params);
};
// 获取租户用户列表
export const getSystemTenantUserListApi = (id: number, params?: SystemUser.ReqSystemUserList) => {
  return http.get<ResPage<SystemUser.ResSystemUserItem>>(PORT + `/system/tenant/${id}/user`, params);
};
// 获取租户菜单树
export const getSystemTenantMenuListApi = (params?: SystemMenu.ReqSystemMenuList) => {
  return http.get<SystemMenu.ResSystemMenuItem[]>(PORT + `/system/tenant/menu`, params);
};
// 快捷登录
export const getSystemTenantLogin = (id: number) => {
  return http.get<SystemUser.ResSystemUserLogin>(PORT + `/system/tenant/${id}/login`);
};
