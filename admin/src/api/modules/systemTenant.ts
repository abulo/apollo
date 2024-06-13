import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemTenant } from "@/api/interface/systemTenant";
import { SystemUser } from "@/api/interface/systemUser";
import { SystemMenu } from "@/api/interface/systemMenu";

// 获取租户列表
export const getSystemTenantListApi = (params?: SystemTenant.ReqSystemTenantList) => {
  return http.get<ResPage<SystemTenant.ResSystemTenantItem>>(PORT + `/system/tenant`, params);
};

// 获取单个租户
export const getSystemTenantItemApi = (id: number) => {
  return http.get<SystemTenant.ResSystemTenantItem>(PORT + `/system/tenant/${id}/item`);
};
// 添加租户
export const addSystemTenantApi = (params: SystemTenant.ResSystemTenantItem) => {
  return http.post(PORT + `/system/tenant`, params);
};
// 修改租户
export const updateSystemTenantApi = (id: number, params: SystemTenant.ResSystemTenantItem) => {
  return http.put(PORT + `/system/tenant/${id}/update`, params);
};
// 删除租户
export const deleteSystemTenantApi = (id: number) => {
  return http.delete(PORT + `/system/tenant/${id}/delete`);
};
// 恢复租户
export const recoverSystemTenantApi = (id: number) => {
  return http.put(PORT + `/system/tenant/${id}/recover`);
};
// 搜索租户
export const getSystemTenantListSimpleApi = (params?: SystemTenant.ReqSystemTenantList) => {
  return http.get<ResPage<SystemTenant.ResSystemTenantItem>>(PORT + `/system/tenant/simple`, params);
};
// 获取租户用户列表
export const getSystemTenantUserSearchApi = (id: number, params?: SystemUser.ReqSystemUserList) => {
  return http.get<ResPage<SystemUser.ResSystemUserItem>>(PORT + `/system/tenant/${id}/user`, params);
};
// 获取租户菜单树
export const getSystemTenantMenuListApi = (params?: SystemMenu.ReqSystemMenuList) => {
  return http.get<SystemMenu.ResSystemMenuList[]>(PORT + `/system/tenant/menu`, params);
};
// 快捷登录
export const getSystemTenantLogin = (id: number) => {
  return http.get<SystemUser.ResSystemUserLogin>(PORT + `/system/tenant/${id}/login`);
};
