// system_post 职位
import { ReqPage } from "./index";
export namespace SystemPost {
  export interface ReqSystemPostList extends ReqPage {
    deleted?: number; // 是否删除
    status?: number; // 状态
    name?: string; // 职位名称
  }
  export interface ResSystemPostItem {
    id: number; // 职位ID
    name: string; // 职位名称
    sort: number; // 显示顺序
    status: number; // 状态
    deleted: number; // 是否删除
    tenantId: number; // 租户ID
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
}
