// pay_wallet_recharge 会员钱包充值
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { PayWalletRecharge } from "@/api/interface/payWalletRecharge";
// 会员钱包充值删除数据
export const deletePayWalletRechargeApi = (walletId: number, id: number) => {
  return http.delete(PORT + `/pay/wallet/${walletId}/recharge/${id}/delete`);
};
// 会员钱包充值查询单条数据
export const getPayWalletRechargeItemApi = (walletId: number, id: number) => {
  return http.get<PayWalletRecharge.ResPayWalletRechargeItem>(PORT + `/pay/wallet/${walletId}/recharge/${id}/item`);
};
// 会员钱包充值恢复数据
export const recoverPayWalletRechargeApi = (walletId: number, id: number) => {
  return http.put(PORT + `/pay/wallet/${walletId}/recharge/${id}/recover`);
};
// 会员钱包充值列表数据
export const getPayWalletRechargeListApi = (walletId: number, params?: PayWalletRecharge.ReqPayWalletRechargeList) => {
  return http.get<ResPage<PayWalletRecharge.ResPayWalletRechargeItem>>(PORT + `/pay/wallet/${walletId}/recharge`, params);
};
