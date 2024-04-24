export namespace SystemNotice {
  // 接口请求数据
  export interface ReqSystemNoticeList {
    deleted?: number; //是否删除
    status?: number; //状态
    type?: number; //类型
    title?: string; //标题
  }
  // 单个数据接口
  export interface ResSystemNoticeItem {
    id: number; // 公告ID
    title: string; // 公告标题
    content: string; // 公告内容
    type: number; // 公告类型（1通知 2公告）
    status: number; // 公告状态（0正常 1关闭）
    deleted: number; // 删除
  }
}
