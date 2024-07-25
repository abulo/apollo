// system_dict_type 字典类型
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemDictType } from "@/api/interface/systemDictType";
// 字典类型创建数据
export const addSystemDictTypeApi = (params: SystemDictType.ResSystemDictTypeItem) => {
  return http.post(PORT + `/system/dict`, params);
};
// 字典类型更新数据
export const updateSystemDictTypeApi = (id: number, params: SystemDictType.ResSystemDictTypeItem) => {
  return http.put(PORT + `/system/dict/${id}/update`, params);
};
// 字典类型删除数据
export const deleteSystemDictTypeApi = (id: number) => {
  return http.delete(PORT + `/system/dict/${id}/delete`);
};
// 字典类型查询单条数据
export const getSystemDictTypeApi = (id: number) => {
  return http.get<SystemDictType.ResSystemDictTypeItem>(PORT + `/system/dict/${id}/item`);
};
// 字典类型列表数据
export const getSystemDictTypeListApi = (params?: SystemDictType.ReqSystemDictTypeList) => {
  return http.get<ResPage<SystemDictType.ResSystemDictTypeItem>>(PORT + `/system/dict`, params);
};
// 字典类型列表数据（简单）
export const getSystemDictTypeListSimpleApi = (params?: SystemDictType.ReqSystemDictTypeList) => {
  return http.get<SystemDictType.ResSystemDictTypeItem[]>(PORT + `/system/dict/simple`, params);
};
