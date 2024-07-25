// system_notify_message 站内信消息表
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemNotifyMessage } from "@/api/interface/systemNotifyMessage";
// 站内信消息表删除数据
export const deleteSystemNotifyMessageApi = (id: number) => {
  return http.delete(PORT + `/system/notify/message/${id}/delete`);
};
// 站内信消息表查询单条数据
export const getSystemNotifyMessageApi = (id: number) => {
  return http.get<SystemNotifyMessage.ResSystemNotifyMessageItem>(PORT + `/system/notify/message/${id}/item`);
};
// 站内信消息表恢复数据
export const recoverSystemNotifyMessageApi = (id: number) => {
  return http.put(PORT + `/system/notify/message/${id}/recover`);
};
// 站内信消息表清理数据
export const dropSystemNotifyMessageApi = (id: number) => {
  return http.delete(PORT + `/system/notify/message/${id}/drop`);
};
// 站内信消息表列表数据
export const getSystemNotifyMessageListApi = (params?: SystemNotifyMessage.ReqSystemNotifyMessageList) => {
  return http.get<ResPage<SystemNotifyMessage.ResSystemNotifyMessageItem>>(PORT + `/system/notify/message`, params);
};
// 站内信消息表批量删除数据
export const deleteSystemNotifyMessageMultipleApi = (params: SystemNotifyMessage.ReqSystemNotifyMessageMultiple) => {
  return http.put(PORT + `/system/notify/message/delete`, params);
};
// 站内信消息表批量恢复数据
export const recoverSystemNotifyMessageMultipleApi = (params: SystemNotifyMessage.ReqSystemNotifyMessageMultiple) => {
  return http.put(PORT + `/system/notify/message/recover`, params);
};
// 站内信消息表批量清理数据
export const dropSystemNotifyMessageMultipleApi = (params: SystemNotifyMessage.ReqSystemNotifyMessageMultiple) => {
  return http.put(PORT + `/system/notify/message/drop`, params);
};
