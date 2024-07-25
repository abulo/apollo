// system_dict_type 字典类型
import { ReqPage } from "./index";
export namespace SystemDictType {
  export interface ReqSystemDictTypeList extends ReqPage {
    type?: string; // 字典类型
    status?: number; // 状态（0正常 1停用）
    name?: string; // 字典名称
  }
  export interface ResSystemDictTypeItem {
    id: number; // 字典主键
    name: string; // 字典名称
    type: string; // 字典类型
    status: number; // 状态（0正常 1停用）
    remark: string | undefined; // 备注
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
}
