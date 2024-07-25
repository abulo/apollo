// system_operate_log 操作日志
import { ReqPage } from "./index";
export namespace SystemOperateLog {
  export interface ReqSystemOperateLogList extends ReqPage {
    deleted?: number; // 删除
    username?: string; // 用户账号
    module?: string; // 模块名称
    beginStartTime?: string; // 开始操作时间
    finishStartTime?: string; // 结束操作时间
    result?: number; // 结果(0 成功/1 失败)
    deptId?: number; // 部门
  }
  export interface ResSystemOperateLogItem {
    id: number; // 主键
    username: string; // 用户账号
    module: string; // 模块名称
    requestMethod: string; // 请求方法名
    requestUrl: string; // 请求地址
    userIp: string; // 用户 ip
    userAgent: string | undefined; // UA
    goMethod: string; // 方法名
    goMethodArgs: any | undefined; // 方法的参数
    startTime: string; // 操作开始时间
    duration: number; // 执行时长
    channel: string; // 渠道
    result: number; // 结果(0 成功/1 失败)
    userId: number; // 用户 ID
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
  export interface ReqSystemOperateLogMultiple {
    ids?: number[]; // 访问日志编号
  }
}
