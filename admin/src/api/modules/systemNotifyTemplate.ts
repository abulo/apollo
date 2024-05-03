import { PORT } from "@/api/config/servicePort";
import { ResPage } from "@/api/interface/index";
import http from "@/api";
import { SystemNotifyTemplate } from "@/api/interface/systemNotifyTemplate";

// 获取列表
export const getSystemNotifyTemplateListApi = (params?: SystemNotifyTemplate.ReqSystemNotifyTemplateList) => {
  return http.get<ResPage<SystemNotifyTemplate.ResSystemNotifyTemplateItem>>(PORT + `/system/notify/template`, params);
};

// 获取单个信息
export const getSystemNotifyTemplateItemApi = (id: number) => {
  return http.get<SystemNotifyTemplate.ResSystemNotifyTemplateItem>(PORT + `/system/notify/template/${id}/item`);
};

// 添加
export const addSystemNotifyTemplateApi = (params: SystemNotifyTemplate.ResSystemNotifyTemplateItem) => {
  return http.post(PORT + `/system/notify/template`, params);
};

// 修改
export const updateSystemNotifyTemplateApi = (id: number, params: SystemNotifyTemplate.ResSystemNotifyTemplateItem) => {
  return http.put(PORT + `/system/notify/template/${id}/update`, params);
};

// 删除
export const deleteSystemNotifyTemplateApi = (id: number) => {
  return http.delete(PORT + `/system/notify/template/${id}/delete`);
};

// 恢复
export const recoverSystemNotifyTemplateApi = (id: number) => {
  return http.put(PORT + `/system/notify/template/${id}/recover`);
};
