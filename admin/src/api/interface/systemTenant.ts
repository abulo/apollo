import { ReqPage } from "./index";
export namespace SystemTenant {
  export interface ReqSystemTenantList extends ReqPage {
    delete?: number; //是否删除(0否 1是)
    status?: number; //状态（0正常 1停用）
    name?: string; // 租户名称
    tenantPackageId?: number; // 套餐编号
    beginExpireDate?: string; // 过期时间
    finishExpireDate?: string; // 过期时间
  }
  export interface ResSystemTenantItem {
    id: number | undefined; // 租户编号
    name: string; // 租户名称
    userId?: number | undefined; // 联系人 id
    contactName: string; // 联系人
    contactMobile: string; // 联系人电话
    status: number; // 状态
    domain: string; // 域名
    expireDate: string; // 过期时间
    accountCount: number; // 账号数量
    tenantPackageId: number | undefined; // 套餐编号
    deleted: number; // 是否删除
    username?: string | undefined; // 用户名称
    password?: string | undefined; // 用户密码
  }
}
