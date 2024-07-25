// system_mail_account 邮箱账号表
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemMailAccount } from "@/api/interface/systemMailAccount";
// 邮箱账号表创建数据
export const addSystemMailAccountApi = (params: SystemMailAccount.ResSystemMailAccountItem) => {
  return http.post(PORT + `/system/mail/account`, params);
};
// 邮箱账号表更新数据
export const updateSystemMailAccountApi = (id: number, params: SystemMailAccount.ResSystemMailAccountItem) => {
  return http.put(PORT + `/system/mail/account/${id}/update`, params);
};
// 邮箱账号表删除数据
export const deleteSystemMailAccountApi = (id: number) => {
  return http.delete(PORT + `/system/mail/account/${id}/delete`);
};
// 邮箱账号表查询单条数据
export const getSystemMailAccountApi = (id: number) => {
  return http.get<SystemMailAccount.ResSystemMailAccountItem>(PORT + `/system/mail/account/${id}/item`);
};
// 邮箱账号表恢复数据
export const recoverSystemMailAccountApi = (id: number) => {
  return http.put(PORT + `/system/mail/account/${id}/recover`);
};
// 邮箱账号表清理数据
export const dropSystemMailAccountApi = (id: number) => {
  return http.delete(PORT + `/system/mail/account/${id}/drop`);
};
// 邮箱账号表列表数据
export const getSystemMailAccountListApi = (params?: SystemMailAccount.ReqSystemMailAccountList) => {
  return http.get<ResPage<SystemMailAccount.ResSystemMailAccountItem>>(PORT + `/system/mail/account`, params);
};
// 邮箱账号表列表数据 精简
export const getSystemMailAccountListSimpleApi = (params?: SystemMailAccount.ReqSystemMailAccountList) => {
  return http.get<SystemMailAccount.ResSystemMailAccountItem[]>(PORT + `/system/mail/account/simple`, params);
};
