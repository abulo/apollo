export namespace PayOrderExtension {
  export interface ResPayOrderExtensionItem {
    id: number; // 编号
    no: string; // 支付订单号
    orderId: number; // 支付订单编号
    channelId: number; // 渠道编号
    channelCode: string; // 渠道编码
    userIp: string | undefined; // 用户 IP
    status: number; // 支付状态
    channelExtras: string | undefined; // 支付渠道的额外参数
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
