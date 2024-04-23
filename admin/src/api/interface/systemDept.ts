export namespace SystemDept {
  // 接口请求数据
  export interface ReqSystemDeptList {
    name?: string; // 名称
    deleted?: number; //是否删除
    status?: number; //菜单状态(0开启/1关闭)
    tree?: number; // 是否树形结构
  }
  // 单个数据接口
  export interface ResSystemDeptItem {
    id: number; // 部门ID
    name: string; // 部门名称
    parentId: number; // 父部门ID
    sort: number; // 显示顺序
    userId: number; // 负责人
    phone: string; // 联系电话
    email: string; // 邮件
    status: number; // 部门状态
    deleted: number; // 是否删除
  }
  // 返回数据
  export interface ResSystemDeptList extends ResSystemDeptItem {
    children?: ResSystemDeptList[]; // 子菜单
  }
}
