export namespace SystemDict {
  export interface ReqSystemDictList {
    dictType?: number; //字典类型
    status?: number; //状态（0正常 1停用）
  }
  export interface ResSystemDictItem {
    id: number; //字典主键
    sort: number; //字典排序
    label: string; //字典标签
    value: string; //字典键值
    dictType: string; //字典类型
    status: number; //状态（0正常 1停用）
    colorType: string; //颜色类型
    cssClass: string; //css 样式
    remark: string; //备注
    creator: string; //创建人
    createTime: string; //创建时间
    updater: string; //更新人
    updateTime: string; //更新时间
  }
  export interface Dict {
    [key: string]: ResSystemDictItem[];
  }
}
