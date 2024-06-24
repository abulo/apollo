import { PORT } from "@/api/config/servicePort";
import { ResPage } from "@/api/interface/index";
import http from "@/api";
import { SystemEntryLog } from "@/api/interface/systemEntryLog";

// 获取列表
export const getSystemEntryLogListApi = (params?: SystemEntryLog.ReqSystemEntryLogList) => {
  return http.get<ResPage<SystemEntryLog.ResSystemEntryLogItem>>(PORT + `/system/logger/entry`, params);
};

// 删除
export const deleteSystemEntryLogApi = (id: string) => {
  return http.delete(PORT + `/system/logger/entry/${id}/delete`);
};

// 单个数据
export const getSystemEntryLogItemApi = (id: string) => {
  return http.get<SystemEntryLog.ResSystemEntryLogItem>(PORT + `/system/logger/entry/${id}/item`);
};

// 清空
export const dropSystemEntryLogApi = (params?: SystemEntryLog.ReqSystemEntryLogDrop) => {
  return http.post(PORT + `/system/logger/entry/drop`, params);
};
