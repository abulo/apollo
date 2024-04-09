import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { Region } from "@/api/interface/region";

// 获取列表
export const getRegionListApi = (params?: Region.ReqRegionList) => {
  return http.get<Region.ResRegionList[]>(PORT + `/region`, params);
};

// 获取单个
export const getRegionItemApi = (id: number) => {
  return http.get<Region.ResRegionItem>(PORT + `/region/${id}/item`);
};

// 添加
export const addRegionApi = (params: Region.ResRegionItem) => {
  return http.post(PORT + `/region`, params);
};

// 修改
export const updateRegionApi = (id: number, params: Region.ResRegionItem) => {
  return http.put(PORT + `/region/${id}/update`, params);
};

// 删除
export const deleteRegionApi = (id: number) => {
  return http.delete(PORT + `/region/${id}/delete`);
};

// 恢复
export const recoverRegionApi = (id: number) => {
  return http.put(PORT + `/region/${id}/recover`);
};

// 搜索
export const getRegionSearchApi = (params?: Region.ReqRegionList) => {
  return http.get<Region.ResRegionList[]>(PORT + `/region/search`, params);
};
