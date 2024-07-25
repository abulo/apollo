// system_notify_template 站内信模板表
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemNotifyTemplate } from "@/api/interface/systemNotifyTemplate";
// 站内信模板表创建数据
export const addSystemNotifyTemplateApi = (params: SystemNotifyTemplate.ResSystemNotifyTemplateItem) => {
  return http.post(PORT + `/system/notify/template`, params);
};
// 站内信模板表更新数据
export const updateSystemNotifyTemplateApi = (id: number, params: SystemNotifyTemplate.ResSystemNotifyTemplateItem) => {
  return http.put(PORT + `/system/notify/template/${id}/update`, params);
};
// 站内信模板表删除数据
export const deleteSystemNotifyTemplateApi = (id: number) => {
  return http.delete(PORT + `/system/notify/template/${id}/delete`);
};
// 站内信模板表查询单条数据
export const getSystemNotifyTemplateApi = (id: number) => {
  return http.get<SystemNotifyTemplate.ResSystemNotifyTemplateItem>(PORT + `/system/notify/template/${id}/item`);
};
// 站内信模板表恢复数据
export const recoverSystemNotifyTemplateApi = (id: number) => {
  return http.put(PORT + `/system/notify/template/${id}/recover`);
};
// 站内信模板表清理数据
export const dropSystemNotifyTemplateApi = (id: number) => {
  return http.delete(PORT + `/system/notify/template/${id}/drop`);
};
// 站内信模板表列表数据
export const getSystemNotifyTemplateListApi = (params?: SystemNotifyTemplate.ReqSystemNotifyTemplateList) => {
  return http.get<ResPage<SystemNotifyTemplate.ResSystemNotifyTemplateItem>>(PORT + `/system/notify/template`, params);
};
// 站内信模板表列表数据精简
export const getSystemNotifyTemplateListSimpleApi = (params?: SystemNotifyTemplate.ReqSystemNotifyTemplateList) => {
  return http.get<SystemNotifyTemplate.ResSystemNotifyTemplateItem[]>(PORT + `/system/notify/template/simple`, params);
};
