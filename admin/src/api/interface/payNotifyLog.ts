// pay_notify_log 支付通知日志
export namespace PayNotifyLog {
  export interface ReqPayNotifyLogList {
    tenantId?: number; // 租户
    deleted?: number; // 删除
    status?: number; // 通知状态
    taskId?: number; // 通知任务编号
  }
  export interface ResPayNotifyLogItem {
    id: number; // 日志编号
    taskId: number; // 通知任务编号
    notifyTimes: number; // 第几次被通知
    response: any; // 请求参数
    status: number; // 通知状态
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
}
