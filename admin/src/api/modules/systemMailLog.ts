// system_mail_log 邮件日志表
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemMailLog } from "@/api/interface/systemMailLog";
// 邮件日志表删除数据
export const deleteSystemMailLogApi = (id: number) => {
  return http.delete(PORT + `/system/mail/log/${id}/delete`);
};
// 邮件日志表查询单条数据
export const getSystemMailLogApi = (id: number) => {
  return http.get<SystemMailLog.ResSystemMailLogItem>(PORT + `/system/mail/log/${id}/item`);
};
// 邮件日志表恢复数据
export const recoverSystemMailLogApi = (id: number) => {
  return http.put(PORT + `/system/mail/log/${id}/recover`);
};
// 邮件日志表清理数据
export const dropSystemMailLogApi = (id: number) => {
  return http.delete(PORT + `/system/mail/log/${id}/drop`);
};
// 邮件日志表列表数据
export const getSystemMailLogListApi = (params?: SystemMailLog.ReqSystemMailLogList) => {
  return http.get<ResPage<SystemMailLog.ResSystemMailLogItem>>(PORT + `/system/mail/log`, params);
};

// 邮件日志表批量删除数据
export const deleteSystemMailLogMultipleApi = (params: SystemMailLog.ReqSystemMailLogMultiple) => {
  return http.put(PORT + `/system/mail/log/delete`, params);
};
// 邮件日志表批量恢复数据
export const recoverSystemMailLogMultipleApi = (params: SystemMailLog.ReqSystemMailLogMultiple) => {
  return http.put(PORT + `/system/mail/log/recover`, params);
};
// 邮件日志表批量清理数据
export const dropSystemMailLogMultipleApi = (params: SystemMailLog.ReqSystemMailLogMultiple) => {
  return http.put(PORT + `/system/mail/log/drop`, params);
};
