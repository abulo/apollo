import { ReqPage } from "./index";
export namespace SystemMailTemplate {
  export interface ReqSystemMailTemplateList extends ReqPage {
    deleted?: number; // 是否删除
    status?: number; // 开启状态
    title?: string; // 模板标题
    name?: string; // 模板名称
    code?: string; // 模板编码
    accountId?: number; // 发送的邮箱账号编号
  }
  export interface ResSystemMailTemplateItem {
    id: number; // 主键
    name: string; // 模板名称
    code: string; // 模板编码
    accountId: number | undefined; // 发送的邮箱账号编号
    nickname: string; // 发送人名称
    title: string; // 模板标题
    content: string; // 模板内容
    params: any | undefined; // 参数数组
    status: number; // 开启状态
    remark: string; // 备注
    deleted: number; // 是否删除
  }
}
