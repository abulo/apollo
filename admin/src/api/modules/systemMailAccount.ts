import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemMailAccount } from "@/api/interface/systemMailAccount";

// 获取邮箱账号列表
export const getSystemMailAccountListApi = (params?: SystemMailAccount.ReqSystemMailAccountList) => {
  return http.get<SystemMailAccount.ResSystemMailAccountItem[]>(PORT + `/system/mail/account`, params);
};

// 获取单个邮箱账号
export const getSystemMailAccountItemApi = (id: number) => {
  return http.get<SystemMailAccount.ResSystemMailAccountItem>(PORT + `/system/mail/account/${id}/item`);
};

// 添加邮箱账号
export const addSystemMailAccountApi = (params: SystemMailAccount.ResSystemMailAccountItem) => {
  return http.post(PORT + `/system/mail/account`, params);
};

// 修改邮箱账号
export const updateSystemMailAccountApi = (id: number, params: SystemMailAccount.ResSystemMailAccountItem) => {
  return http.put(PORT + `/system/mail/account/${id}/update`, params);
};

// 删除邮箱账号
export const deleteSystemMailAccountApi = (id: number) => {
  return http.delete(PORT + `/system/mail/account/${id}/delete`);
};

// 恢复邮箱账号
export const recoverSystemMailAccountApi = (id: number) => {
  return http.put(PORT + `/system/mail/account/${id}/recover`);
};

// 搜索邮箱账号
export const getSystemMailAccountSearchApi = (params?: SystemMailAccount.ReqSystemMailAccountList) => {
  return http.get<SystemMailAccount.ResSystemMailAccountItem[]>(PORT + `/system/mail/account/search`, params);
};
