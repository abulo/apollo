import { ReqPage } from "./index";
export namespace SystemLoginLog {
  // 列表接口请求数据
  export interface ReqSystemLoginLogList extends ReqPage {
    username?: string; // 用户账号
    beginLoginTime?: string; // 开始登录时间
    finishLoginTime?: string; // 结束登录时间
    channel?: string; // 渠道
  }
  export interface ResSystemLoginLogItem {
    id: number; // id
    username: string; // 用户账号
    userIp: string; // 用户ip
    userAgent: string; // UA
    loginTime: string; // 登录时间
    channel: string; // 渠道
    creator: string; // 创建人
    createTime: string; // 创建时间
    updater: string; // 更新人
    updateTime: string; // 更新时间
  }
  export interface ReqSystemLoginLogDelete {
    systemLoginLogIds: number[]; // 访问日志编号
  }
  export interface ReqSystemLoginLogDrop {
    username?: string; // 用户账号
    beginLoginTime?: string; // 开始登录时间
    finishLoginTime?: string; // 结束登录时间
    channel?: string; // 渠道
  }
}
