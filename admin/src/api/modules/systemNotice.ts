// system_notice 通知公告表
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemNotice } from "@/api/interface/systemNotice";
// 通知公告表创建数据
export const addSystemNoticeApi = (params: SystemNotice.ResSystemNoticeItem) => {
  return http.post(PORT + `/system/notice`, params);
};
// 通知公告表更新数据
export const updateSystemNoticeApi = (id: number, params: SystemNotice.ResSystemNoticeItem) => {
  return http.put(PORT + `/system/notice/${id}/update`, params);
};
// 通知公告表删除数据
export const deleteSystemNoticeApi = (id: number) => {
  return http.delete(PORT + `/system/notice/${id}/delete`);
};
// 通知公告表查询单条数据
export const getSystemNoticeApi = (id: number) => {
  return http.get<SystemNotice.ResSystemNoticeItem>(PORT + `/system/notice/${id}/item`);
};
// 通知公告表恢复数据
export const recoverSystemNoticeApi = (id: number) => {
  return http.put(PORT + `/system/notice/${id}/recover`);
};
// 通知公告表清理数据
export const dropSystemNoticeApi = (id: number) => {
  return http.delete(PORT + `/system/notice/${id}/drop`);
};
// 通知公告表列表数据
export const getSystemNoticeListApi = (params?: SystemNotice.ReqSystemNoticeList) => {
  return http.get<ResPage<SystemNotice.ResSystemNoticeItem>>(PORT + `/system/notice`, params);
};
// 通知公告表列表数据精简
export const getSystemNoticeListSimpleApi = (params?: SystemNotice.ReqSystemNoticeList) => {
  return http.get<SystemNotice.ResSystemNoticeItem[]>(PORT + `/system/notice/simple`, params);
};
