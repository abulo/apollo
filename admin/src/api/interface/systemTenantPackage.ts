export namespace SystemTenantPackage {
  export interface ReqSystemTenantPackageList {
    delete?: number; //是否删除(0否 1是)
    status?: number; //状态（0正常 1停用）
    name?: string; // 套餐名称
  }
  export interface ResSystemTenantPackageItem {
    id: number; //bigint 套餐编号,PRI
    name: string; //varchar 套餐名称
    status: number; //tinyint 状态（0正常 1停用）
    menuIds: number[]; //json 目录编号
    remark: string; //varchar 备注
    deleted: number; //tinyint 是否删除(0否 1是)
  }
}
