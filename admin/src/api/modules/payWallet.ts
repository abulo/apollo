// pay_wallet 会员钱包表
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { PayWallet } from "@/api/interface/payWallet";
// 会员钱包表创建数据
export const addPayWalletApi = (params: PayWallet.ResPayWalletItem) => {
  return http.post(PORT + `/pay/wallet`, params);
};
// 会员钱包表更新数据
export const updatePayWalletApi = (id: number, params: PayWallet.ResPayWalletItem) => {
  return http.put(PORT + `/pay/wallet/${id}/update`, params);
};
// 会员钱包表删除数据
export const deletePayWalletApi = (id: number) => {
  return http.delete(PORT + `/pay/wallet/${id}/delete`);
};
// 会员钱包表查询单条数据
export const getPayWalletItemApi = (id: number) => {
  return http.get<PayWallet.ResPayWalletItem>(PORT + `/pay/wallet/${id}/item`);
};
// 会员钱包表恢复数据
export const recoverPayWalletApi = (id: number) => {
  return http.put(PORT + `/pay/wallet/${id}/recover`);
};
// 会员钱包表列表数据
export const getPayWalletListApi = (params?: PayWallet.ReqPayWalletList) => {
  return http.get<ResPage<PayWallet.ResPayWalletItem>>(PORT + `/pay/wallet`, params);
};
