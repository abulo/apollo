import { SystemUserDept } from "@/api/interface/systemUserDept";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";

// 获取列表
export const getSystemUserDeptListApi = (id: number) => {
  return http.get<SystemUserDept.ResSystemUserDeptItem>(PORT + `/system/user/${id}/dept`);
};

// 新增
export const addSystemUserDeptApi = (id: number, params: SystemUserDept.ResSystemUserDeptItem) => {
  return http.post(PORT + `/system/user/${id}/dept`, params);
};
