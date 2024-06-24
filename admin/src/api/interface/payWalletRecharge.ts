// pay_wallet_recharge 会员钱包充值
import { ReqPage } from "./index";
export namespace PayWalletRecharge {
  export interface ReqPayWalletRechargeList extends ReqPage {
    deleted?: number; // 删除
    packageId?: number; // 充值套餐编号
    payStatus?: number; // 是否已支付：[0:未支付 1:已经支付过]
    payOrderId?: number; // 支付订单编号
    payChannelCode?: string; // 支付成功的支付渠道
    beginPayTime?: string; // 订单支付时间
    finishPayTime?: string; // 订单支付时间
    refundStatus?: number; // 退款状态
    payRefundId?: number; // 支付退款单编号
    beginRefundTime?: string; // 退款时间
    finishRefundTime?: string; // 退款时间
  }
  export interface ResPayWalletRechargeItem {
    id: number; // 编号
    walletId: number; // 会员钱包id
    totalPrice: number; // 用户实际到账余额，例如充 100 送 20，则该值是 120
    payPrice: number; // 实际支付金额
    bonusPrice: number; // 钱包赠送金额
    packageId: number | undefined; // 充值套餐编号
    payStatus: number; // 是否已支付：[0:未支付 1:已经支付过]
    payOrderId: number; // 支付订单编号
    payChannelCode: string; // 支付成功的支付渠道
    payTime: string | undefined; // 订单支付时间
    payRefundId: number | undefined; // 支付退款单编号
    refundTotalPrice: number | undefined; // 退款金额，包含赠送金额
    refundPayPrice: number | undefined; // 退款支付金额
    refundBonusPrice: number | undefined; // 退款钱包赠送金额
    refundTime: string | undefined; // 退款时间
    refundStatus: number | undefined; // 退款状态
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
}
