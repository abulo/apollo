import { ReqPage } from "./index";
export namespace SystemMailLog {
  export interface ReqSystemMailLogList extends ReqPage {
    deleted?: number; // 是否删除
    sendStatus?: number; // 发送状态
    beginSendTime?: string; // 发送时间
    finishSendTime?: string; // 发送时间
    templateTitle?: string; // 邮件标题
    userId?: number; // 用户编号
    userType?: number; // 用户类型
    accountId?: number; // 邮箱账号编号
    templateCode?: string; // 模板编码
  }
  export interface ResSystemMailLogItem {
    id: number; // 主键
    userId: number; // 用户编号
    userType: number; // 用户类型
    toMail: string; // 接收邮箱地址
    accountId: number; // 邮箱账号编号
    fromMail: string; // 发送邮箱地址
    templateId: number; // 模板编号
    templateCode: string; // 模板编码
    templateNickname: string; // 模版发送人名称
    templateTitle: string; // 邮件标题
    templateContent: string; // 邮件内容
    templateParams: any | undefined; // 邮件参数
    sendStatus: number; // 发送状态
    sendTime: string; // 发送时间
    sendMessageId: string; // 发送返回的消息 ID
    sendException: string; // 发送异常
    deleted: number; // 是否删除
  }
}
