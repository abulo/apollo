// system_dept 部门
export namespace SystemDept {
  export interface ReqSystemDeptList {
    deleted?: number; // 是否删除
    status?: number; // 部门状态
    name?: string; // 部门名称
  }
  export interface ResSystemDeptItem {
    id: number; // 部门ID
    name: string; // 部门名称
    parentId: number; // 父部门ID
    sort: number; // 显示顺序
    userId: number | undefined; // 负责人
    phone: string | undefined; // 联系电话
    email: string | undefined; // 邮件
    status: number; // 部门状态
    deleted: number; // 是否删除
    tenantId: number; // 租户ID
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
    children?: ResSystemDeptItem[]; // 子部门
  }
}
