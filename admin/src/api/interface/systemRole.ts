// system_role 系统角色
import { ReqPage } from "./index";
export namespace SystemRole {
  export interface ReqSystemRoleList extends ReqPage {
    deleted?: number; // 删除
    type?: number; // 角色类型(1内置/2定义)
    status?: number; // 角色状态（0正常 1停用）
    name?: string; // 角色名称
  }
  export interface ResSystemRoleItem {
    id: number; // 角色编号
    name: string; // 角色名称
    code: string; // 角色权限字符串
    sort: number; // 显示顺序
    dataScope: number | undefined; // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
    dataScopeDept: number[] | undefined; // 数据范围(指定部门数组)
    status: number; // 角色状态（0正常 1停用）
    type: number; // 角色类型(1内置/2定义)
    remark: string | undefined; // 备注
    deleted: number; // 删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建者
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新者
    updateTime: string | undefined; // 更新时间
    menuIds: number[] | undefined; // 菜单编号数组
  }
}
