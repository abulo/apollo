export namespace SystemPost {
  // 接口请求数据
  export interface ReqSystemPostList {
    name?: string; // 名称
    deleted?: number; //是否删除
    status?: number; //状态
  }
  // 单个数据接口
  export interface ResSystemPostItem {
    id: number; // 职位ID
    name: string; // 职位名称
    sort: number; // 显示顺序
    status: number; // 职位状态
    deleted: number; // 是否删除
  }
}
