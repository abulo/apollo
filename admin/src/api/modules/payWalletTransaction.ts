// pay_wallet_transaction 会员钱包流水表
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { PayWalletTransaction } from "@/api/interface/payWalletTransaction";
// 会员钱包流水表删除数据
export const deletePayWalletTransactionApi = (id: number) => {
  return http.delete(PORT + `/pay/wallet/:walletId/transaction/${id}/delete`);
};
// 会员钱包流水表查询单条数据
export const getPayWalletTransactionItemApi = (id: number) => {
  return http.get<PayWalletTransaction.ResPayWalletTransactionItem>(PORT + `/pay/wallet/:walletId/transaction/${id}/item`);
};
// 会员钱包流水表恢复数据
export const recoverPayWalletTransactionApi = (id: number) => {
  return http.put(PORT + `/pay/wallet/:walletId/transaction/${id}/recover`);
};
// 会员钱包流水表列表数据
export const getPayWalletTransactionListApi = (params?: PayWalletTransaction.ReqPayWalletTransactionList) => {
  return http.get<ResPage<PayWalletTransaction.ResPayWalletTransactionItem>>(PORT + `/pay/wallet/:walletId/transaction`, params);
};
