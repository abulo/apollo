import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemDept } from "@/api/interface/systemDept";
// 获取部门列表
export const getSystemDeptListApi = (params?: SystemDept.ReqSystemDeptList) => {
  return http.get<SystemDept.ResSystemDeptList[]>(PORT + `/system/dept`, params);
};
// 获取单个部门
export const getSystemDeptItemApi = (id: number) => {
  return http.get<SystemDept.ResSystemDeptItem>(PORT + `/system/dept/${id}/item`);
};
// 添加部门
export const addSystemDeptApi = (params: SystemDept.ResSystemDeptItem) => {
  return http.post(PORT + `/system/dept`, params);
};
// 修改部门
export const updateSystemDeptApi = (id: number, params: SystemDept.ResSystemDeptItem) => {
  return http.put(PORT + `/system/dept/${id}/update`, params);
};
// 删除部门
export const deleteSystemDeptApi = (id: number) => {
  return http.delete(PORT + `/system/dept/${id}/delete`);
};
// 恢复部门
export const recoverSystemDeptApi = (id: number) => {
  return http.put(PORT + `/system/dept/${id}/recover`);
};
// 获取部门树
export const getSystemDeptListSimpleApi = (params?: SystemDept.ReqSystemDeptList) => {
  return http.get<SystemDept.ResSystemDeptList[]>(PORT + `/system/dept/simple`, params);
};
