import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { PayApp } from "@/api/interface/payApp";

// 获取支付应用列表
export const getPayAppListApi = (params?: PayApp.ReqPayAppList) => {
  return http.get<PayApp.ResPayAppItem[]>(PORT + `/pay/app`, params);
};

// 获取单个支付应用
export const getPayAppItemApi = (id: number) => {
  return http.get<PayApp.ResPayAppItem>(PORT + `/pay/app/${id}/item`);
};

// 添加支付应用
export const addPayAppApi = (params: PayApp.ResPayAppItem) => {
  return http.post(PORT + `/pay/app`, params);
};

// 修改支付应用
export const updatePayAppApi = (id: number, params: PayApp.ResPayAppItem) => {
  return http.put(PORT + `/pay/app/${id}/update`, params);
};

// 删除支付应用
export const deletePayAppApi = (id: number) => {
  return http.delete(PORT + `/pay/app/${id}/delete`);
};

// 恢复支付应用
export const recoverPayAppApi = (id: number) => {
  return http.put(PORT + `/pay/app/${id}/recover`);
};
