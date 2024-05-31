// pay_refund 退款订单
import { ReqPage } from "./index";
export namespace PayRefund {
  export interface ReqPayRefundList extends ReqPage {
    deleted?: number; // 删除
    appId?: number; // 应用编号
    channelId?: number; // 渠道编号
    channelCode?: string; // 渠道编码
    merchantOrderId?: string; // 商户订单编号（商户系统生成）
    orderId?: number; // 支付订单编号 pay_order 表id
    orderNo?: string; // 支付订单 no
    channelOrderNo?: string; // 渠道订单号，pay_order 中的 channel_order_no 对应
    status?: number; // 退款状态
    beginCreateTime?: string; // 开始创建时间
    finishCreateTime?: string; // 结束创建时间
  }
  export interface ResPayRefundItem {
    id: number; // 支付退款编号
    no: string; // 退款单号
    appId: number; // 应用编号
    channelId: number; // 渠道编号
    channelCode: string; // 渠道编码
    orderId: number; // 支付订单编号 pay_order 表id
    orderNo: string; // 支付订单 no
    merchantOrderId: string; // 商户订单编号（商户系统生成）
    merchantRefundId: string; // 商户退款订单号（商户系统生成）
    notifyUrl: string; // 异步通知商户地址
    status: number; // 退款状态
    payPrice: number; // 支付金额,单位分
    refundPrice: number; // 退款金额,单位分
    reason: string; // 退款原因
    userIp: string | undefined; // ip
    channelOrderNo: string; // 渠道订单号，pay_order 中的 channel_order_no 对应
    channelRefundNo: string | undefined; // 渠道退款单号，渠道返
    successTime: string | undefined; // 退款成功时间
    channelErrorCode: string | undefined; // 渠道调用报错时，错误码
    channelErrorMsg: string | undefined; // 渠道调用报错时，错误信息
    channelNotifyData: string | undefined; // 支付渠道异步通知的内容
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
}
