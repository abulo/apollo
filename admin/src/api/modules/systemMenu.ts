import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemMenu } from "@/api/interface/systemMenu";

// 获取系统菜单列表
export const getSystemMenuListApi = (params?: SystemMenu.ReqSystemMenuList) => {
  return http.get<SystemMenu.ResSystemMenuList[]>(PORT + `/system/menu`, params);
};

// 获取单个系统菜单
export const getSystemMenuItemApi = (id: number) => {
  return http.get<SystemMenu.ResSystemMenuItem>(PORT + `/system/menu/${id}/item`);
};

// 添加系统菜单
export const addSystemMenuApi = (params: SystemMenu.ResSystemMenuItem) => {
  return http.post(PORT + `/system/menu`, params);
};

// 修改系统菜单
export const updateSystemMenuApi = (id: number, params: SystemMenu.ResSystemMenuItem) => {
  return http.put(PORT + `/system/menu/${id}/update`, params);
};

// 删除系统菜单
export const deleteSystemMenuApi = (id: number) => {
  return http.delete(PORT + `/system/menu/${id}/delete`);
};

// 恢复系统菜单
export const recoverSystemMenuApi = (id: number) => {
  return http.put(PORT + `/system/menu/${id}/recover`);
};
