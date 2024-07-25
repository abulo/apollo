// system_dict 字典数据表
import { ReqPage } from "./index";
export namespace SystemDict {
  export interface ReqSystemDictList extends ReqPage {
    status?: number; // 状态（0正常 1停用）
  }
  export interface ResSystemDictItem {
    id: number; // 字典编码
    sort: number; // 字典排序
    label: string; // 字典标签
    value: string; // 字典键值
    dictTypeId: number; // 字段类型 ID;
    status: number; // 状态（0正常 1停用）
    colorType: string | undefined; // 颜色类型
    cssClass: string | undefined; // css 样式
    remark: string | undefined; // 备注
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
    dictType?: string | undefined;
  }
  export interface Dict {
    [key: string]: ResSystemDictItem[];
  }
}
