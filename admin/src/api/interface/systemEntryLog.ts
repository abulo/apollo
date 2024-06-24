import { ReqPage } from "./index";
export namespace SystemEntryLog {
  // 列表接口请求数据
  export interface ReqSystemEntryLogList extends ReqPage {
    startTime?: string; // 开始时间
    endTime?: string; // 结束时间
  }
  export interface ResSystemEntryLogItem {
    id: string; // id
    host: string; // 主机
    timestamp: string; // 时间
    file: string; // 文件
    func: string; // 函数
    message: string; // 消息
    data: any; // 数据
    level: string; // 级别
  }
  export interface ReqSystemEntryLogDrop {
    ids?: string[]; // 日志编号
    startTime?: string; // 开始时间
    endTime?: string; // 结束时间
  }
}
