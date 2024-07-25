// system_file 文件管理
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemFile } from "@/api/interface/systemFile";
// 文件管理创建数据
export const addSystemFileApi = (params: FormData) => {
  return http.post<SystemFile.ResSystemFileItem>(PORT + `/system/file`, params);
};
// 文件管理重新上传数据
export const reUploadSystemFileApi = (id: number, params: FormData) => {
  return http.post<SystemFile.ResSystemFileItem>(PORT + `/system/file/${id}/reupload`, params);
};
// 文件管理更新数据
export const updateSystemFileApi = (id: number, params: SystemFile.ResSystemFileItem) => {
  return http.put(PORT + `/system/file/${id}/update`, params);
};
// 文件管理删除数据
export const deleteSystemFileApi = (id: number) => {
  return http.delete(PORT + `/system/file/${id}/delete`);
};
// 文件管理查询单条数据
export const getSystemFileApi = (id: number) => {
  return http.get<SystemFile.ResSystemFileItem>(PORT + `/system/file/${id}/item`);
};
// 文件管理恢复数据
export const recoverSystemFileApi = (id: number) => {
  return http.put(PORT + `/system/file/${id}/recover`);
};
// 文件管理清理数据
export const dropSystemFileApi = (id: number) => {
  return http.delete(PORT + `/system/file/${id}/drop`);
};
// 文件管理列表数据
export const getSystemFileListApi = (params?: SystemFile.ReqSystemFileList) => {
  return http.get<ResPage<SystemFile.ResSystemFileItem>>(PORT + `/system/file`, params);
};
