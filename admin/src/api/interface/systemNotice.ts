// system_notice 通知公告表
import { ReqPage } from "./index";
export namespace SystemNotice {
  export interface ReqSystemNoticeList extends ReqPage {
    deleted?: number; // 删除
    status?: number; // 公告状态（0正常 1关闭）
    type?: number; // 公告类型（1通知 2公告）
    title?: string; // 公告标题
  }
  export interface ResSystemNoticeItem {
    id: number; // 公告ID
    title: string; // 公告标题
    content: string; // 公告内容
    type: number; // 公告类型（1通知 2公告）
    status: number; // 公告状态（0正常 1关闭）
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
}
