// system_mail_log 邮件日志表
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
    id: number; // 编号
    userId: number | undefined; // 用户编号
    userType: number | undefined; // 用户类型
    toMail: string; // 接收邮箱地址
    accountId: number; // 邮箱账号编号
    fromMail: string; // 发送邮箱地址
    templateId: number; // 模板编号
    templateCode: string; // 模板编码
    templateNickname: string | undefined; // 模版发送人名称
    templateTitle: string; // 邮件标题
    templateContent: string; // 邮件内容
    templateParams: any | undefined; // 邮件参数
    sendStatus: number; // 发送状态
    sendTime: string | undefined; // 发送时间
    sendMessageId: string | undefined; // 发送返回的消息 ID
    sendException: string | undefined; // 发送异常
    creator: string | undefined; // 创建者
    createTime: string; // 创建时间
    updater: string | undefined; // 更新者
    updateTime: string; // 更新时间
    deleted: number; // 是否删除
  }
  export interface ReqSystemMailLogMultiple {
    ids?: number[]; // 访问日志编号
  }
}
