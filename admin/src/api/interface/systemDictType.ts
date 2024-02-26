import { ReqPage } from "./index";
export namespace SystemDictType {
  export interface ReqSystemDictTypeList extends ReqPage {
    type?: number; //字典类型
    status?: number; //状态（0正常 1停用）
  }
  export interface ResSystemDictTypeItem {
    id: number; //字典主键
    name: string; //字典名称
    type: string; //字典类型
    status: number; //状态（0正常 1停用）
    remark: string; //备注
    creator: string; //创建人
    createTime: string; //创建时间
    updater: string; //更新人
    updateTime: string; //更新时间
  }
}
