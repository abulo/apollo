// system_operate_log 操作日志
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemOperateLog } from "@/api/interface/systemOperateLog";
// 操作日志删除数据
export const deleteSystemOperateLogApi = (id: number) => {
  return http.delete(PORT + `/system/logger/operate/${id}/delete`);
};
// 操作日志查询单条数据
export const getSystemOperateLogApi = (id: number) => {
  return http.get<SystemOperateLog.ResSystemOperateLogItem>(PORT + `/system/logger/operate/${id}/item`);
};
// 操作日志恢复数据
export const recoverSystemOperateLogApi = (id: number) => {
  return http.put(PORT + `/system/logger/operate/${id}/recover`);
};
// 操作日志清理数据
export const dropSystemOperateLogApi = (id: number) => {
  return http.delete(PORT + `/system/logger/operate/${id}/drop`);
};
// 操作日志列表数据
export const getSystemOperateLogListApi = (params?: SystemOperateLog.ReqSystemOperateLogList) => {
  return http.get<ResPage<SystemOperateLog.ResSystemOperateLogItem>>(PORT + `/system/logger/operate`, params);
};

// 操作日志批量删除数据
export const deleteSystemOperateLogMultipleApi = (params: SystemOperateLog.ReqSystemOperateLogMultiple) => {
  return http.put(PORT + `/system/logger/operate/delete`, params);
};
// 操作日志批量恢复数据
export const recoverSystemOperateLogMultipleApi = (params: SystemOperateLog.ReqSystemOperateLogMultiple) => {
  return http.put(PORT + `/system/logger/operate/recover`, params);
};
// 操作日志批量清理数据
export const dropSystemOperateLogMultipleApi = (params: SystemOperateLog.ReqSystemOperateLogMultiple) => {
  return http.put(PORT + `/system/logger/operate/drop`, params);
};
