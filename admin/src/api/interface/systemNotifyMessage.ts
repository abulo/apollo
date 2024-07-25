// system_notify_message 站内信消息表
import { ReqPage } from "./index";
export namespace SystemNotifyMessage {
  export interface ReqSystemNotifyMessageList extends ReqPage {
    deleted?: number; // 删除
    templateType?: number; // 模板类型
    readStatus?: number; // 阅读状态
    beginReadTime?: string; // 开始阅读时间
    finishReadTime?: string; // 结束阅读时间
  }
  export interface ResSystemNotifyMessageItem {
    id: number; // 消息
    userId: number; // 用户id
    userType: number; // 用户类型
    templateId: number; // 模版编号
    templateCode: string; // 模板编码
    templateNickname: string; // 模版发送人名称
    templateContent: string; // 模版内容
    templateType: number; // 模版类型
    templateParams: any; // 模版参数
    readStatus: number; // 是否已读
    readTime: string | undefined; // 阅读时间
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
  export interface ReqSystemNotifyMessageMultiple {
    ids?: number[]; // 访问日志编号
  }
}
