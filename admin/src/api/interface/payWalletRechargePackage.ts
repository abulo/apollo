// pay_wallet_recharge_package 充值套餐表
import { ReqPage } from "./index";
export namespace PayWalletRechargePackage {
  export interface ReqPayWalletRechargePackageList extends ReqPage {
    deleted?: number; // 删除
    status?: number; // 状态
    name?: string; // 套餐名称
  }
  export interface ResPayWalletRechargePackageItem {
    id: number; // 编号
    name: string; // 套餐名称
    payPrice: number; // 支付金额
    bonusPrice: number; // 赠送金额
    status: number; // 状态
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
}
