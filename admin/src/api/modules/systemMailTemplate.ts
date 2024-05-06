import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemMailTemplate } from "@/api/interface/systemMailTemplate";

// 获取邮件模板列表
export const getSystemMailTemplateListApi = (params?: SystemMailTemplate.ReqSystemMailTemplateList) => {
  return http.get<ResPage<SystemMailTemplate.ResSystemMailTemplateItem>>(PORT + `/system/mail/template`, params);
};

// 获取单个邮件模板
export const getSystemMailTemplateItemApi = (id: number) => {
  return http.get<SystemMailTemplate.ResSystemMailTemplateItem>(PORT + `/system/mail/template/${id}/item`);
};

// 添加邮件模板
export const addSystemMailTemplateApi = (params: SystemMailTemplate.ResSystemMailTemplateItem) => {
  return http.post(PORT + `/system/mail/template`, params);
};

// 修改邮件模板
export const updateSystemMailTemplateApi = (id: number, params: SystemMailTemplate.ResSystemMailTemplateItem) => {
  return http.put(PORT + `/system/mail/template/${id}/update`, params);
};

// 删除邮件模板
export const deleteSystemMailTemplateApi = (id: number) => {
  return http.delete(PORT + `/system/mail/template/${id}/delete`);
};

// 恢复邮件模板
export const recoverSystemMailTemplateApi = (id: number) => {
  return http.put(PORT + `/system/mail/template/${id}/recover`);
};
