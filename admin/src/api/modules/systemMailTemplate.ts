// system_mail_template 邮件模版表
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemMailTemplate } from "@/api/interface/systemMailTemplate";
// 邮件模版表创建数据
export const addSystemMailTemplateApi = (params: SystemMailTemplate.ResSystemMailTemplateItem) => {
  return http.post(PORT + `/system/mail/template`, params);
};
// 邮件模版表更新数据
export const updateSystemMailTemplateApi = (id: number, params: SystemMailTemplate.ResSystemMailTemplateItem) => {
  return http.put(PORT + `/system/mail/template/${id}/update`, params);
};
// 邮件模版表删除数据
export const deleteSystemMailTemplateApi = (id: number) => {
  return http.delete(PORT + `/system/mail/template/${id}/delete`);
};
// 邮件模版表查询单条数据
export const getSystemMailTemplateApi = (id: number) => {
  return http.get<SystemMailTemplate.ResSystemMailTemplateItem>(PORT + `/system/mail/template/${id}/item`);
};
// 邮件模版表恢复数据
export const recoverSystemMailTemplateApi = (id: number) => {
  return http.put(PORT + `/system/mail/template/${id}/recover`);
};
// 邮件模版表清理数据
export const dropSystemMailTemplateApi = (id: number) => {
  return http.delete(PORT + `/system/mail/template/${id}/drop`);
};
// 邮件模版表列表数据
export const getSystemMailTemplateListApi = (params?: SystemMailTemplate.ReqSystemMailTemplateList) => {
  return http.get<ResPage<SystemMailTemplate.ResSystemMailTemplateItem>>(PORT + `/system/mail/template`, params);
};
// 邮件模版表列表数据精简
export const getSystemMailTemplateListSimpleApi = (params?: SystemMailTemplate.ReqSystemMailTemplateList) => {
  return http.get<SystemMailTemplate.ResSystemMailTemplateItem[]>(PORT + `/system/mail/template/simple`, params);
};
