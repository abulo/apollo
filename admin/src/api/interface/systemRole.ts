export namespace SystemRole {
  export interface ReqSystemRoleList {
    type?: number; // 角色类型(1内置/2定义)
    status?: number; //状态(0:正常/1:禁用)
  }
  export interface ResSystemRoleItem {
    id: number; //角色编号
    name: string; //角色名称
    code: string; //角色权限字符串
    sort: number; //显示顺序
    status: number; //角色状态（0正常 1停用）
    type: number; //角色类型(1内置/2定义)
    remark: string; //备注
    creator: string; //创建者
    createTime: string; //创建时间
    updater: string; //更新者
    updateTime: string; //更新时间
    systemMenuIds: number[]; //json 数据范围(指定菜单数组)
  }
}
