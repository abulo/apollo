import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemDict } from "@/api/interface/systemDict";

// 获取字典列表
export const getSystemDictListApi = (params?: SystemDict.ReqSystemDictList) => {
  return http.get<SystemDict.ResSystemDictItem[]>(PORT + `/system/dict/data`, params);
};

// 获取单个字典
export const getSystemDictItemApi = (id: number) => {
  return http.get<SystemDict.ResSystemDictItem>(PORT + `/system/dict/data/${id}/item`);
};

// 添加字典
export const addSystemDictApi = (params: SystemDict.ResSystemDictItem) => {
  return http.post(PORT + `/system/dict/data`, params);
};

// 修改字典
export const updateSystemDictApi = (id: number, params: SystemDict.ResSystemDictItem) => {
  return http.put(PORT + `/system/dict/data/${id}/update`, params);
};

// 删除字典
export const deleteSystemDictApi = (id: number) => {
  return http.delete(PORT + `/system/dict/data/${id}/delete`);
};

// 获取所有字典
export const getAllSystemDictApi = () => {
  return http.get<SystemDict.Dict>(PORT + `/system/dict/all`);
};
