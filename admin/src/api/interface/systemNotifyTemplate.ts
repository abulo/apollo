// system_notify_template 站内信模板表
import { ReqPage } from "./index";
export namespace SystemNotifyTemplate {
  export interface ReqSystemNotifyTemplateList extends ReqPage {
    deleted?: number; // 删除
    status?: number; // 状态
    type?: number; // 类型
    name?: string; // 模板名称
  }
  export interface ResSystemNotifyTemplateItem {
    id: number; // 主键
    name: string; // 模板名称
    code: string; // 模版编码
    nickname: string; // 发送人名称
    content: string; // 模版内容
    type: number; // 类型
    params: any | undefined; // 参数数组
    status: number; // 状态
    remark: string | undefined; // 备注
    deleted: number; // 删除
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
}
