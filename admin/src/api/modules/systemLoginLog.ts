// system_login_log 登录日志
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemLoginLog } from "@/api/interface/systemLoginLog";
// 登录日志删除数据
export const deleteSystemLoginLogApi = (id: number) => {
  return http.delete(PORT + `/system/logger/login/${id}/delete`);
};

// 登录日志查询单条数据
export const getSystemLoginLogApi = (id: number) => {
  return http.get<SystemLoginLog.ResSystemLoginLogItem>(PORT + `/system/logger/login/${id}/item`);
};
// 登录日志恢复数据
export const recoverSystemLoginLogApi = (id: number) => {
  return http.put(PORT + `/system/logger/login/${id}/recover`);
};

// 登录日志清理数据
export const dropSystemLoginLogApi = (id: number) => {
  return http.delete(PORT + `/system/logger/login/${id}/drop`);
};

// 登录日志列表数据
export const getSystemLoginLogListApi = (params?: SystemLoginLog.ReqSystemLoginLogList) => {
  return http.get<ResPage<SystemLoginLog.ResSystemLoginLogItem>>(PORT + `/system/logger/login`, params);
};

// 登录日志批量删除数据
export const deleteSystemLoginLogMultipleApi = (params: SystemLoginLog.ReqSystemLoginLogMultiple) => {
  return http.put(PORT + `/system/logger/login/delete`, params);
};
// 登录日志批量恢复数据
export const recoverSystemLoginLogMultipleApi = (params: SystemLoginLog.ReqSystemLoginLogMultiple) => {
  return http.put(PORT + `/system/logger/login/recover`, params);
};
// 登录日志批量清理数据
export const dropSystemLoginLogMultipleApi = (params: SystemLoginLog.ReqSystemLoginLogMultiple) => {
  return http.put(PORT + `/system/logger/login/drop`, params);
};
