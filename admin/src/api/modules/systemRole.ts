// system_role 系统角色
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemRole } from "@/api/interface/systemRole";
// 系统角色创建数据
export const addSystemRoleApi = (params: SystemRole.ResSystemRoleItem) => {
  return http.post(PORT + `/system/role`, params);
};
// 系统角色更新数据
export const updateSystemRoleApi = (id: number, params: SystemRole.ResSystemRoleItem) => {
  return http.put(PORT + `/system/role/${id}/update`, params);
};
// 系统角色删除数据
export const deleteSystemRoleApi = (id: number) => {
  return http.delete(PORT + `/system/role/${id}/delete`);
};
// 系统角色查询单条数据
export const getSystemRoleApi = (id: number) => {
  return http.get<SystemRole.ResSystemRoleItem>(PORT + `/system/role/${id}/item`);
};
// 系统角色恢复数据
export const recoverSystemRoleApi = (id: number) => {
  return http.put(PORT + `/system/role/${id}/recover`);
};
// 系统角色清理数据
export const dropSystemRoleApi = (id: number) => {
  return http.delete(PORT + `/system/role/${id}/drop`);
};
// 系统角色列表数据
export const getSystemRoleListApi = (params?: SystemRole.ReqSystemRoleList) => {
  return http.get<ResPage<SystemRole.ResSystemRoleItem>>(PORT + `/system/role`, params);
};
// 系统角色列表精简数据
export const getSystemRoleListSimpleApi = (params?: SystemRole.ReqSystemRoleList) => {
  return http.get<SystemRole.ResSystemRoleItem[]>(PORT + `/system/role/simple`, params);
};
