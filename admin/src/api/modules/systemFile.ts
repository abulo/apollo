import { SystemFile } from "@/api/interface/systemFile";
import { PORT } from "@/api/config/servicePort";
import { ResPage } from "@/api/interface/index";
import http from "@/api";

// 新增

export const getSystemFileListApi = (params?: SystemFile.ReqSystemFileList) => {
  return http.get<ResPage<SystemFile.ResSystemFileItem>>(PORT + `/system/file`, params);
};

export const getSystemFileItemApi = (id: number) => {
  return http.get<SystemFile.ResSystemFileItem>(PORT + `/system/file/${id}/item`);
};

export const addSystemFileApi = (params: FormData) => {
  return http.post<SystemFile.ResSystemFileItem>(PORT + `/system/file`, params);
};

export const updateSystemFileApi = (id: number, params: SystemFile.ResSystemFileItem) => {
  return http.put(PORT + `/system/file/${id}/update`, params);
};

export const deleteSystemFileApi = (id: number) => {
  return http.delete(PORT + `/system/file/${id}/delete`);
};
