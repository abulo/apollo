import { PORT } from "@/api/config/servicePort";
import { ResPage } from "@/api/interface/index";
import http from "@/api";
import { SystemNotifyMessage } from "@/api/interface/systemNotifyMessage";

// 获取列表
export const getSystemNotifyMessageListApi = (params?: SystemNotifyMessage.ReqSystemNotifyMessageList) => {
  return http.get<ResPage<SystemNotifyMessage.ResSystemNotifyMessageItem>>(PORT + `/system/notify/message`, params);
};

// 获取单个信息
export const getSystemNotifyMessageItemApi = (id: number) => {
  return http.get<SystemNotifyMessage.ResSystemNotifyMessageItem>(PORT + `/system/notify/message/${id}/item`);
};

// 添加
// export const addSystemNotifyMessageApi = (params: SystemNotifyMessage.ResSystemNotifyMessageItem) => {
//   return http.post(PORT + `/system/notify/message`, params);
// };

// // 修改
// export const updateSystemNotifyMessageApi = (id: number, params: SystemNotifyMessage.ResSystemNotifyMessageItem) => {
//   return http.put(PORT + `/system/notify/message/${id}/update`, params);
// };

// 删除
export const deleteSystemNotifyMessageApi = (id: number) => {
  return http.delete(PORT + `/system/notify/message/${id}/delete`);
};

// 恢复
export const recoverSystemNotifyMessageApi = (id: number) => {
  return http.put(PORT + `/system/notify/message/${id}/recover`);
};
