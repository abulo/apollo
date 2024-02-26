import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemDictType } from "@/api/interface/systemDictType";

// 获取字典类型列表
export const getSystemDictTypeListApi = (params?: SystemDictType.ReqSystemDictTypeList) => {
  return http.get<ResPage<SystemDictType.ResSystemDictTypeItem>>(PORT + `/system/dict`, params);
};

// 获取单个字典类型
export const getSystemDictTypeItemApi = (id: number) => {
  return http.get<SystemDictType.ResSystemDictTypeItem>(PORT + `/system/dict/${id}/item`);
};

// 添加字典类型
export const addSystemDictTypeApi = (params: SystemDictType.ResSystemDictTypeItem) => {
  return http.post(PORT + `/system/dict`, params);
};

// 修改字典类型
export const updateSystemDictTypeApi = (id: number, params: SystemDictType.ResSystemDictTypeItem) => {
  return http.put(PORT + `/system/dict/${id}/update`, params);
};

// 删除字典类型
export const deleteSystemDictTypeApi = (id: number) => {
  return http.delete(PORT + `/system/dict/${id}/delete`);
};

// 获取所有的字典类型
export const getAllSystemDictTypeApi = () => {
  return http.get<SystemDictType.ResSystemDictTypeItem[]>(PORT + `/system/dict/type`);
};
