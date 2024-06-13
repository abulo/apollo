import { ReqPage } from "./index";
export namespace SystemOperateLog {
  // 列表接口请求数据
  export interface ReqSystemOperateLogList extends ReqPage {
    deleted?: number; // 是否删除
    username?: string; // 用户账号
    module?: string; // 模块名称
    beginStartTime?: string; // 开始操作时间
    finishStartTime?: string; // 结束操作时间
    result?: number; // 结果(0 成功/1 失败)
  }
  export interface ResSystemOperateLogItem {
    id: number; // id
    username: string; // 用户账号
    module: string; // 模块名称
    requestMethod: string; // 请求方法名
    requestUrl: string; // 请求地址
    userIp: string; // 用户 ip
    userAgent: string; // UA
    goMethod: string; // 方法名
    goMethodArgs: string; // 方法的参数
    startTime: string; // 操作开始时间
    duration: number; // 执行时长
    channel: string; // 渠道
    result: number; // 结果(0 成功/1 失败)
    deleted: number; // 是否删除
    creator: string; // 创建人
    createTime: string; // 创建时间
    updater: string; // 更新人
    updateTime: string; // 更新时间
  }
  export interface ReqSystemOperateLogDrop {
    ids?: number[]; // 操作日志编号
    deleted?: number; // 是否删除
    username?: string; // 用户账号
    module?: string; // 模块名称
    beginStartTime?: string; // 开始操作时间
    finishStartTime?: string; // 结束操作时间
    result?: number; // 结果(0 成功/1 失败)
  }
}
