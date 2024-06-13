import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemTenantPackage } from "@/api/interface/systemTenantPackage";

// 获取套餐列表
export const getSystemTenantPackageListApi = (params?: SystemTenantPackage.ReqSystemTenantPackageList) => {
  return http.get<SystemTenantPackage.ResSystemTenantPackageItem[]>(PORT + `/system/tenant/package`, params);
};
// 获取单个套餐
export const getSystemTenantPackageItemApi = (id: number) => {
  return http.get<SystemTenantPackage.ResSystemTenantPackageItem>(PORT + `/system/tenant/package/${id}/item`);
};
// 添加套餐
export const addSystemTenantPackageApi = (params: SystemTenantPackage.ResSystemTenantPackageItem) => {
  return http.post(PORT + `/system/tenant/package`, params);
};
// 修改套餐
export const updateSystemTenantPackageApi = (id: number, params: SystemTenantPackage.ResSystemTenantPackageItem) => {
  return http.put(PORT + `/system/tenant/package/${id}/update`, params);
};
// 删除套餐
export const deleteSystemTenantPackageApi = (id: number) => {
  return http.delete(PORT + `/system/tenant/package/${id}/delete`);
};
// 恢复套餐
export const recoverSystemTenantPackageApi = (id: number) => {
  return http.put(PORT + `/system/tenant/package/${id}/recover`);
};
// 搜索套餐
export const getSystemTenantPackageListSimpleApi = (params?: SystemTenantPackage.ReqSystemTenantPackageList) => {
  return http.get<SystemTenantPackage.ResSystemTenantPackageItem[]>(PORT + `/system/tenant/package/simple`, params);
};
