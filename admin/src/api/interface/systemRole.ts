export namespace SystemRole {
  export interface ReqSystemRoleList {
    type?: number; // 角色类型(1内置/2定义)
    status?: number; //状态(0:正常/1:禁用)
    name?: string; // 名称
    deleted?: number; // 删除
  }
  export interface ResSystemRoleItem {
    id: number; // 角色编号
    name: string; // 角色名称
    code: string; // 角色 code
    sort: number; // 排序
    type: number; // 类别
    status: number; // 状态
    remark: string; // 备注
    deleted: number; // 删除
  }
  export interface ResSystemRoleScopeItem {
    id: number; // 角色编号
    dataScope: number | undefined; // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
    dataScopeDept: number[] | undefined; //数据范围(指定部门数组)
  }

  export interface ReqSystemRoleMenuItem {
    deleted?: number; // 删除
    menuId?: number; // 菜单
  }
  export interface ResSystemRoleMenuItem {
    roleId: number; // 角色编号
    menuIds: number[] | undefined; //菜单编号
  }
}
