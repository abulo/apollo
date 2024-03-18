import { SystemUserRole } from "@/api/interface/systemUserRole";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";

// 获取列表
export const getSystemUserRoleListApi = (id: number) => {
  return http.get<SystemUserRole.ResSystemUserRoleItem>(PORT + `/system/user/${id}/role`);
};

// 新增
export const addSystemUserRoleApi = (id: number, params: SystemUserRole.ResSystemUserRoleItem) => {
  return http.post(PORT + `/system/user/${id}/role`, params);
};
