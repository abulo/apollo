export namespace PayChannel {
  // 单个数据接口
  export interface ResPayChannelItem {
    id: number; // 渠道编号
    code: string; // 渠道编码
    status: number; // 开启状态
    remark: string; // 备注
    feeRate: number; // 渠道费率，单位：百分比
    appId: number; // 应用编号
    config: any; // 支付渠道配置
    deleted: number; // 删除
    tenantId: number; // 租户
  }

  export interface ResPayChannelConfig {
    appId?: string | undefined; // 微信和支付宝走需要这个
    serverUrl?: string | undefined; // varchar 编号,,
    signType?: string | undefined;
    mode?: number | undefined; // varchar 编号,,
    privateKey?: string | undefined;
    alipayPublicKey?: string | undefined;
    appCertContent?: string | undefined;
    alipayPublicCertContent?: string | undefined;
    rootCertContent?: string | undefined;
    // 下面是微信
    mchId?: string | undefined;
    apiVersion?: string | undefined;
    mchKey?: string | undefined;
    keyContent?: string | undefined;
    privateKeyContent?: string | undefined;
    privateCertContent?: string | undefined;
    apiV3Key?: string | undefined;
  }
}
