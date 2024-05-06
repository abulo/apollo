import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemMailLog } from "@/api/interface/systemMailLog";

// 获取邮件日志列表
export const getSystemMailLogListApi = (params?: SystemMailLog.ReqSystemMailLogList) => {
  return http.get<ResPage<SystemMailLog.ResSystemMailLogItem>>(PORT + `/system/mail/log`, params);
};

// 获取单个邮件日志
export const getSystemMailLogItemApi = (id: number) => {
  return http.get<SystemMailLog.ResSystemMailLogItem>(PORT + `/system/mail/log/${id}/item`);
};

// 删除邮件日志
export const deleteSystemMailLogApi = (id: number) => {
  return http.delete(PORT + `/system/mail/log/${id}/delete`);
};

// 恢复邮件日志
export const recoverSystemMailLogApi = (id: number) => {
  return http.put(PORT + `/system/mail/log/${id}/recover`);
};

