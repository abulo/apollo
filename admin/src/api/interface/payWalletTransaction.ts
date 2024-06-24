// pay_wallet_transaction 会员钱包流水表
import { ReqPage } from "./index";
export namespace PayWalletTransaction {
  export interface ReqPayWalletTransactionList extends ReqPage {
    deleted?: number; // 删除
    bizType?: number; // 关联类型
    bizId?: string; // 关联业务编号
    no?: string; // 流水号
    title?: string; // 流水标题
  }
  export interface ResPayWalletTransactionItem {
    id: number; // 编号
    walletId: number; // 会员钱包 id
    bizType: number; // 关联类型
    bizId: string; // 关联业务编号
    no: string; // 流水号
    title: string; // 流水标题
    price: number; // 交易金额, 单位分
    balance: number; // 余额, 单位分
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
}
