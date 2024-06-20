// pay_wallet_recharge_package 充值套餐表
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { PayWalletRechargePackage } from "@/api/interface/payWalletRechargePackage";
// 充值套餐表创建数据
export const addPayWalletRechargePackageApi = (params: PayWalletRechargePackage.ResPayWalletRechargePackageItem) => {
  return http.post(PORT + `/pay/wallet/package`, params);
};
// 充值套餐表更新数据
export const updatePayWalletRechargePackageApi = (
  id: number,
  params: PayWalletRechargePackage.ResPayWalletRechargePackageItem
) => {
  return http.put(PORT + `/pay/wallet/package/${id}/update`, params);
};
// 充值套餐表删除数据
export const deletePayWalletRechargePackageApi = (id: number) => {
  return http.delete(PORT + `/pay/wallet/package/${id}/delete`);
};
// 充值套餐表查询单条数据
export const getPayWalletRechargePackageItemApi = (id: number) => {
  return http.get<PayWalletRechargePackage.ResPayWalletRechargePackageItem>(PORT + `/pay/wallet/package/${id}/item`);
};
// 充值套餐表恢复数据
export const recoverPayWalletRechargePackageApi = (id: number) => {
  return http.put(PORT + `/pay/wallet/package/${id}/recover`);
};
// 充值套餐表列表数据
export const getPayWalletRechargePackageListApi = (params?: PayWalletRechargePackage.ReqPayWalletRechargePackageList) => {
  return http.get<ResPage<PayWalletRechargePackage.ResPayWalletRechargePackageItem>>(PORT + `/pay/wallet/package`, params);
};

// 充值套餐表精简列表数据
export const getPayWalletRechargePackageListSimpleApi = (params?: PayWalletRechargePackage.ReqPayWalletRechargePackageList) => {
  return http.get<PayWalletRechargePackage.ResPayWalletRechargePackageItem[]>(PORT + `/pay/wallet/package/simple`, params);
};
