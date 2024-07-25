// system_tenant_package 租户套餐包
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemTenantPackage } from "@/api/interface/systemTenantPackage";
// 租户套餐包创建数据
export const addSystemTenantPackageApi = (params: SystemTenantPackage.ResSystemTenantPackageItem) => {
  return http.post(PORT + `/system/tenant/package`, params);
};
// 租户套餐包更新数据
export const updateSystemTenantPackageApi = (id: number, params: SystemTenantPackage.ResSystemTenantPackageItem) => {
  return http.put(PORT + `/system/tenant/package/${id}/update`, params);
};
// 租户套餐包删除数据
export const deleteSystemTenantPackageApi = (id: number) => {
  return http.delete(PORT + `/system/tenant/package/${id}/delete`);
};
// 租户套餐包查询单条数据
export const getSystemTenantPackageApi = (id: number) => {
  return http.get<SystemTenantPackage.ResSystemTenantPackageItem>(PORT + `/system/tenant/package/${id}/item`);
};
// 租户套餐包恢复数据
export const recoverSystemTenantPackageApi = (id: number) => {
  return http.put(PORT + `/system/tenant/package/${id}/recover`);
};
// 租户套餐包清理数据
export const dropSystemTenantPackageApi = (id: number) => {
  return http.delete(PORT + `/system/tenant/package/${id}/drop`);
};
// 租户套餐包列表数据
export const getSystemTenantPackageListApi = (params?: SystemTenantPackage.ReqSystemTenantPackageList) => {
  return http.get<ResPage<SystemTenantPackage.ResSystemTenantPackageItem>>(PORT + `/system/tenant/package`, params);
};
// 租户套餐包列表数据精简
export const getSystemTenantPackageListSimpleApi = (params?: SystemTenantPackage.ReqSystemTenantPackageList) => {
  return http.get<SystemTenantPackage.ResSystemTenantPackageItem[]>(PORT + `/system/tenant/package/simple`, params);
};
