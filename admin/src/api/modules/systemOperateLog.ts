import { PORT } from "@/api/config/servicePort";
import { ResPage } from "@/api/interface/index";
import http from "@/api";
import { SystemOperateLog } from "@/api/interface/systemOperateLog";

// 获取列表
export const getSystemOperateLogListApi = (params?: SystemOperateLog.ReqSystemOperateLogList) => {
  return http.get<ResPage<SystemOperateLog.ResSystemOperateLogItem>>(PORT + `/system/logger/operate`, params);
};

// 恢复
export const recoverSystemOperateLogApi = (id: number) => {
  return http.put(PORT + `/system/logger/operate/${id}/recover`);
};

// 删除
export const deleteSystemOperateLogApi = (id: number) => {
  return http.delete(PORT + `/system/logger/operate/${id}/delete`);
};

// 单个数据
export const getSystemOperateLogItemApi = (id: number) => {
  return http.get<SystemOperateLog.ResSystemOperateLogItem>(PORT + `/system/logger/operate/${id}/item`);
};

// 清空
export const dropSystemOperateLogApi = (params?: SystemOperateLog.ReqSystemOperateLogDrop) => {
  return http.post(PORT + `/system/logger/operate/drop`, params);
};
