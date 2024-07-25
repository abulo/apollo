// system_login_log 登录日志
import { ReqPage } from "./index";
export namespace SystemLoginLog {
  export interface ReqSystemLoginLogList extends ReqPage {
    deleted?: number; // 是否删除
    username?: string; // 用户账号
    beginLoginTime?: string; // 开始登录时间
    finishLoginTime?: string; // 结束登录时间
    channel?: string; // 渠道
    deptId?: number; // 部门
  }
  export interface ResSystemLoginLogItem {
    id: number; // 主键
    username: string; // 用户账号
    userIp: string; // 用户ip
    userAgent: string | undefined; // UA
    loginTime: string; // 登录时间
    channel: string; // 渠道
    userId: number; //
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
  export interface ReqSystemLoginLogMultiple {
    ids?: number[]; // 访问日志编号
  }
}
