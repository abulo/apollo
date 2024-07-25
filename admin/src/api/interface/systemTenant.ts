// system_tenant 租户
import { ReqPage } from "./index";
export namespace SystemTenant {
  export interface ReqSystemTenantList extends ReqPage {
    deleted?: number; // 是否删除(0否 1是)
    status?: number; // 状态（0正常 1停用）
    name?: string; // 租户名称
    beginExpireDate?: string; // 过期时间
    finishExpireDate?: string; // 过期时间
    tenantPackageId?: number; // 套餐编号
  }
  export interface ResSystemTenantItem {
    id: number; // 租户编号
    name: string; // 租户名称
    userId: number | undefined; // 联系人ID
    contactName: string; // 联系人
    contactMobile: string; // 租户联系电话
    status: number; // 状态（0正常 1停用）
    domain: string | undefined; // 域名
    expireDate: string; // 过期时间
    accountCount: number; // 账号数量
    tenantPackageId: number; // 套餐编号
    deleted: number; // 是否删除(0否 1是)
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
    username: string | undefined; // 用户名
    password: string | undefined; // 密码
  }
}
