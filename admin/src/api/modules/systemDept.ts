// system_dept 部门
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemDept } from "@/api/interface/systemDept";
// 部门创建数据
export const addSystemDeptApi = (params: SystemDept.ResSystemDeptItem) => {
  return http.post(PORT + `/system/dept`, params);
};
// 部门更新数据
export const updateSystemDeptApi = (id: number, params: SystemDept.ResSystemDeptItem) => {
  return http.put(PORT + `/system/dept/${id}/update`, params);
};
// 部门删除数据
export const deleteSystemDeptApi = (id: number) => {
  return http.delete(PORT + `/system/dept/${id}/delete`);
};
// 部门查询单条数据
export const getSystemDeptApi = (id: number) => {
  return http.get<SystemDept.ResSystemDeptItem>(PORT + `/system/dept/${id}/item`);
};
// 部门恢复数据
export const recoverSystemDeptApi = (id: number) => {
  return http.put(PORT + `/system/dept/${id}/recover`);
};
// 部门清理数据
export const dropSystemDeptApi = (id: number) => {
  return http.delete(PORT + `/system/dept/${id}/drop`);
};
// 部门列表数据
export const getSystemDeptListApi = (params?: SystemDept.ReqSystemDeptList) => {
  return http.get<SystemDept.ResSystemDeptItem[]>(PORT + `/system/dept`, params);
};

// 获取精简列表
export const getSystemDeptListSimpleApi = (params?: SystemDept.ReqSystemDeptList) => {
  return http.get<SystemDept.ResSystemDeptItem[]>(PORT + `/system/dept/simple`, params);
};

// 获取全部部门
export const getSystemDeptListLabelApi = (params?: SystemDept.ReqSystemDeptList) => {
  return http.get<SystemDept.ResSystemDeptItem[]>(PORT + `/system/dept/label`, params);
};
