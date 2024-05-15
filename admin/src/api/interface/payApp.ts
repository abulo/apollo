export namespace PayApp {
  // 接口请求数据
  export interface ReqPayAppList {
    name?: string; // 名称
    deleted?: number; //是否删除
    status?: number; //状态
  }
  // 单个数据接口
  export interface ResPayAppItem {
    id: number; // 应用编号
    name: string; // 应用名称
    status: number; // 开启状态
    remark: string; // 备注
    orderNotifyUrl: string; // 支付结果的回调地址
    refundNotifyUrl: string; // 退款结果的回调地址
    transferNotifyUrl: string | undefined; // 转账结果的回调地址
    deleted: number; // 删除
    channelList?: any | undefined; // 支付渠道列表
  }
}
