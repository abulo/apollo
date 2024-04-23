export namespace Region {
  // 接口请求数据
  export interface ReqRegionList {
    name?: string; //区域名称
    deleted?: number; //是否删除
    status?: number; // 状态;
    tree?: number; // 是否树形结构
  }
  // 单个数据接口
  export interface ResRegionItem {
    id: number; // 区域编号
    name: string; // 区域名称
    parentId: number; // 父级编号
    status: number; // 状态;
    deleted: number; // 是否删除;
    sort: number; // 排序;
  }
  // 返回数据
  export interface ResRegionList extends ResRegionItem {
    children?: ResRegionList[]; // 子菜单
  }
}
