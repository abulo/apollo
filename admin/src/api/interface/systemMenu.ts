// system_menu 系统菜单
export namespace SystemMenu {
  export interface ReqSystemMenuList {
    deleted?: number; // 是否删除
    status?: number; // 菜单状态(0开启/1关闭)
    type?: number; // 菜单类型(1:目录/2: 菜单/3: 按钮)
  }
  export interface ResSystemMenuItem {
    id: number; // 菜单编号
    name: string; // 菜单名称
    permission: string | undefined; // 权限标识
    type: number; // 菜单类型(1:目录/2: 菜单/3: 按钮)
    sort: number; // 显示顺序
    parentId: number; // 父菜单ID
    path: string | undefined; // 路由地址
    icon: string | undefined; // 菜单图标
    component: string | undefined; // 组件路径
    componentName: string | undefined; // 组件名
    status: number; // 菜单状态(0开启/1关闭)
    hide: number; // 是否隐藏(0:否/1是)
    link: string | undefined; // 路由外链时填写的访问地址
    keepAlive: number; // 是否缓存(0不/ 1是)
    affix: number; // 是否总是显示(0 不显示/1 显示)
    activePath: string | undefined; // 激活链接
    fullScreen: number; // 是否全屏
    redirect: string | undefined; // 路由重定向地址
    creator: string | undefined; // 创建者
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新者
    updateTime: string | undefined; // 更新时间
    deleted: number; // 是否删除
    children?: ResSystemMenuItem[]; // 子菜单
  }
  // 按钮权限
  export interface ResSystemMenuButtons {
    [key: string]: string[];
  }
}
