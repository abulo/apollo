import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemNotice } from "@/api/interface/systemNotice";

// 获取公告列表
export const getSystemNoticeListApi = (params?: SystemNotice.ReqSystemNoticeList) => {
  return http.get<ResPage<SystemNotice.ResSystemNoticeItem>>(PORT + `/system/notice`, params);
};

// 获取单个公告
export const getSystemNoticeItemApi = (id: number) => {
  return http.get<SystemNotice.ResSystemNoticeItem>(PORT + `/system/notice/${id}/item`);
};

// 添加公告
export const addSystemNoticeApi = (params: SystemNotice.ResSystemNoticeItem) => {
  return http.post(PORT + `/system/notice`, params);
};

// 修改公告
export const updateSystemNoticeApi = (id: number, params: SystemNotice.ResSystemNoticeItem) => {
  return http.put(PORT + `/system/notice/${id}/update`, params);
};

// 删除公告
export const deleteSystemNoticeApi = (id: number) => {
  return http.delete(PORT + `/system/notice/${id}/delete`);
};

// 恢复公告
export const recoverSystemNoticeApi = (id: number) => {
  return http.put(PORT + `/system/notice/${id}/recover`);
};
