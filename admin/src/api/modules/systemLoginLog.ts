import { PORT } from "@/api/config/servicePort";
import { ResPage } from "@/api/interface/index";
import http from "@/api";
import { SystemLoginLog } from "@/api/interface/systemLoginLog";

// 获取列表
export const getSystemLoginLogListApi = (params?: SystemLoginLog.ReqSystemLoginLogList) => {
  return http.get<ResPage<SystemLoginLog.ResSystemLoginLogItem>>(PORT + `/system/logger/login`, params);
};

// 恢复
export const recoverSystemLoginLogApi = (id: number) => {
  return http.put(PORT + `/system/logger/login/${id}/recover`);
};

// 删除
export const deleteSystemLoginLogApi = (id: number) => {
  return http.delete(PORT + `/system/logger/login/${id}/delete`);
};

// 单个数据
export const getSystemLoginLogItemApi = (id: number) => {
  return http.get<SystemLoginLog.ResSystemLoginLogItem>(PORT + `/system/logger/login/${id}/item`);
};

// 清空
export const dropSystemLoginLogApi = (params?: SystemLoginLog.ReqSystemLoginLogDrop) => {
  return http.post(PORT + `/system/logger/login/drop`, params);
};
