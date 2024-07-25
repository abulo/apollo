// system_dict 字典数据表
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemDict } from "@/api/interface/systemDict";
// 字典数据表创建数据
export const addSystemDictApi = (dictTypeId: number, params: SystemDict.ResSystemDictItem) => {
  return http.post(PORT + `/system/dict/${dictTypeId}/data`, params);
};
// 字典数据表更新数据
export const updateSystemDictApi = (dictTypeId: number, id: number, params: SystemDict.ResSystemDictItem) => {
  return http.put(PORT + `/system/dict/${dictTypeId}/data/${id}/update`, params);
};
// 字典数据表删除数据
export const deleteSystemDictApi = (dictTypeId: number, id: number) => {
  return http.delete(PORT + `/system/dict/${dictTypeId}/data/${id}/delete`);
};
// 字典数据表查询单条数据
export const getSystemDictApi = (dictTypeId: number, id: number) => {
  return http.get<SystemDict.ResSystemDictItem>(PORT + `/system/dict/${dictTypeId}/data/${id}/item`);
};
// 字典数据表列表数据
export const getSystemDictListApi = (dictTypeId: number, params?: SystemDict.ReqSystemDictList) => {
  return http.get<ResPage<SystemDict.ResSystemDictItem>>(PORT + `/system/dict/${dictTypeId}/data`, params);
};
// 获取所有字典
export const getSystemDictAllApi = () => {
  return http.get<SystemDict.Dict>(PORT + `/system/dict/all`);
};
