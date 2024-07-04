export namespace SystemMenu {
  // 接口请求数据
  export interface ReqSystemMenuList {
    type?: number; //菜单类型(1:目录/2: 菜单/3: 按钮)
    deleted?: number; //是否删除
    status?: number; //菜单状态(0开启/1关闭)
  }
  // 单个数据接口
  export interface ResSystemMenuItem {
    id: number; //菜单ID
    name: string; //菜单名称
    permission: string; //权限标识
    type: number; //菜单类型(1:目录/2: 菜单/3: 按钮)
    sort: number; //显示顺序
    parentId: number; //父菜单ID
    path: string; //路由地址
    icon: string; //菜单图标
    component: string; //组件路径
    componentName: string; //组件名
    status: number; //菜单状态(0开启/1关闭)
    hide: number; //是否隐藏(0:否/1是)
    link: string; //路由外链时填写的访问地址
    keepAlive: number; //是否缓存(0不/ 是)
    affix: number; //是否总是显示(0 不显示/1 显示)
    activePath: string; //激活链接
    fullScreen: number; //是否全屏
    redirect: string; //路由重定向地址
    creator: string; //创建者
    createTime: string; //创建时间
    updater: string; //更新者
    updateTime: string; //更新时间
    deleted: number; //是否删除
  }
  // 返回数据
  export interface ResSystemMenuList extends ResSystemMenuItem {
    children?: ResSystemMenuList[]; // 子菜单
  }
  // 按钮权限
  export interface ResSystemMenuButtons {
    [key: string]: string[];
  }
}
