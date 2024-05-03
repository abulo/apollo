import { ReqPage } from "./index";
export namespace SystemNotifyTemplate {
  export interface ReqSystemNotifyTemplateList extends ReqPage {
    deleted?: number; // 删除
    status?: number; // 状态
    type?: number; // 类型
    name?: string; // 模板名称
  }
  export interface ResSystemNotifyTemplateItem {
    id: number; // id
    name: string; // 模板名称
    code: string; // 模版编码
    nickname: string; // 发送人名称
    content: string; // 模版内容
    type: number; // 类型
    params: any | undefined; // 参数数组
    status: number; // 状态
    remark: string; // 备注
    deleted: number; // 删除
  }
}
