// pay_notify_task 商户支付-任务通知
import { ReqPage } from "./index";
import { PayNotifyLog } from "./payNotifyLog";
export namespace PayNotifyTask {
  export interface ReqPayNotifyTaskList extends ReqPage {
    deleted?: number; // 删除
    appId?: number; // 应用编号
    type?: number; // 通知类型
    dataId?: number; // 数据编号
    status?: number; // 通知状态
    merchantOrderId?: string; // 商户订单编号
    beginCreateTime?: string; // 开始创建时间
    finishCreateTime?: string; // 结束创建时间
  }
  export interface ResPayNotifyTaskItem {
    id: number; // 任务编号
    appId: number; // 应用编号
    type: number; // 通知类型
    dataId: number; // 数据编号
    status: number; // 通知状态
    merchantOrderId: string; // 商户订单编号
    merchantTransferId: string; // 商户转账单编号
    nextNotifyTime: string; // 下一次通知时间
    lastExecuteTime: string; // 最后一次执行时间
    notifyTimes: number; // 当前通知次数
    maxNotifyTimes: number | undefined; // 最大可通知次数
    notifyUrl: string; // 异步通知地址
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
    logs?: PayNotifyLog.ResPayNotifyLogItem[]; // 日志
  }
}
