// system_menu 系统菜单
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemMenu } from "@/api/interface/systemMenu";
// 系统菜单创建数据
export const addSystemMenuApi = (params: SystemMenu.ResSystemMenuItem) => {
  return http.post(PORT + `/system/menu`, params);
};
// 系统菜单更新数据
export const updateSystemMenuApi = (id: number, params: SystemMenu.ResSystemMenuItem) => {
  return http.put(PORT + `/system/menu/${id}/update`, params);
};
// 系统菜单删除数据
export const deleteSystemMenuApi = (id: number) => {
  return http.delete(PORT + `/system/menu/${id}/delete`);
};
// 系统菜单查询单条数据
export const getSystemMenuApi = (id: number) => {
  return http.get<SystemMenu.ResSystemMenuItem>(PORT + `/system/menu/${id}/item`);
};
// 系统菜单恢复数据
export const recoverSystemMenuApi = (id: number) => {
  return http.put(PORT + `/system/menu/${id}/recover`);
};
// 系统菜单清理数据
export const dropSystemMenuApi = (id: number) => {
  return http.delete(PORT + `/system/menu/${id}/drop`);
};
// 系统菜单列表数据
export const getSystemMenuListApi = (params?: SystemMenu.ReqSystemMenuList) => {
  return http.get<SystemMenu.ResSystemMenuItem[]>(PORT + `/system/menu`, params);
};
// 系统菜单列表数据精简
export const getSystemMenuListSimpleApi = (params?: SystemMenu.ReqSystemMenuList) => {
  return http.get<SystemMenu.ResSystemMenuItem[]>(PORT + `/system/menu/simple`, params);
};
