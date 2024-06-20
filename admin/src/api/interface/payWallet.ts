// pay_wallet 会员钱包表
import { ReqPage } from "./index";
export namespace PayWallet {
  export interface ReqPayWalletList extends ReqPage {
    userType?: number; // 用户类型
    userId?: number; // 用户编号
    deleted?: number; // 删除
    username?: string; // 用户名
  }
  export interface ResPayWalletItem {
    id: number; // 编号
    userId: number; // 用户编号
    userType: number; // 用户类型
    balance: number; // 余额，单位分
    totalExpense: number; // 累计支出，单位分
    totalRecharge: number; // 累计充值，单位分
    freezePrice: number; // 冻结金额，单位分
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
    username?: string; // 用户名
  }
}
