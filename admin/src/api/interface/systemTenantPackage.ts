// system_tenant_package 租户套餐包
import { ReqPage } from "./index";
export namespace SystemTenantPackage {
  export interface ReqSystemTenantPackageList extends ReqPage {
    deleted?: number; // 是否删除(0否 1是)
    status?: number; // 状态（0正常 1停用）
    name?: string; // 套餐名称
  }
  export interface ResSystemTenantPackageItem {
    id: number; // 套餐编号
    name: string; // 套餐名称
    status: number; // 状态（0正常 1停用）
    menuIds: number[] | undefined; // 目录编号
    remark: string | undefined; // 备注
    deleted: number; // 是否删除(0否 1是)
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
}
