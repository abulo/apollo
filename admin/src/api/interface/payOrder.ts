// pay_order 支付订单
import { ReqPage } from "./index";
import { PayOrderExtension } from "./payOrderExtension";
export namespace PayOrder {
  export interface ReqPayOrderList extends ReqPage {
    deleted?: number; // 删除
    appId?: number; // 应用编号
    channelId?: number; // 渠道编号
    channelCode?: string; // 渠道编码
    merchantOrderId?: string; // 商户订单编号
    no?: string; // 支付订单号
    channelOrderNo?: string; // 渠道订单号
    status?: number; // 支付状态
    beginCreateTime?: string; // 开始创建时间
    finishCreateTime?: string; // 结束创建时间
  }
  export interface ResPayOrderItem {
    id: number; // 支付订单编号
    appId: number; // 应用编号
    channelId: number; // 渠道编号
    channelCode: string; // 渠道编码
    merchantOrderId: string; // 商户订单编号
    subject: string; // 商品标题
    body: string; // 商品描述
    notifyUrl: string; // 异步通知地址
    price: number; // 支付金额，单位：分
    channelFeeRate: number; // 渠道手续费，单位：百分比
    channelFeePrice: number; // 渠道手续金额，单位：分
    status: number; // 支付状态
    userIp: string | undefined; // 用户 IP
    expireTime: string; // 订单失效时间
    successTime: string | undefined; // 订单支付成功时间
    extensionId: number | undefined; // 支付成功的订单拓展单编号
    no: string | undefined; // 支付订单号
    refundPrice: number; // 退款总金额，单位：分
    channelUserId: string | undefined; // 渠道用户编号
    channelOrderNo: string | undefined; // 渠道订单号
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
    extension?: PayOrderExtension.ResPayOrderExtensionItem | undefined; // 支付订单拓展单
  }
}
