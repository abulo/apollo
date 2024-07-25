// system_post 职位
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemPost } from "@/api/interface/systemPost";
// 职位创建数据
export const addSystemPostApi = (params: SystemPost.ResSystemPostItem) => {
  return http.post(PORT + `/system/post`, params);
};
// 职位更新数据
export const updateSystemPostApi = (id: number, params: SystemPost.ResSystemPostItem) => {
  return http.put(PORT + `/system/post/${id}/update`, params);
};
// 职位删除数据
export const deleteSystemPostApi = (id: number) => {
  return http.delete(PORT + `/system/post/${id}/delete`);
};
// 职位查询单条数据
export const getSystemPostApi = (id: number) => {
  return http.get<SystemPost.ResSystemPostItem>(PORT + `/system/post/${id}/item`);
};
// 职位恢复数据
export const recoverSystemPostApi = (id: number) => {
  return http.put(PORT + `/system/post/${id}/recover`);
};
// 职位清理数据
export const dropSystemPostApi = (id: number) => {
  return http.delete(PORT + `/system/post/${id}/drop`);
};
// 职位列表数据
export const getSystemPostListApi = (params?: SystemPost.ReqSystemPostList) => {
  return http.get<ResPage<SystemPost.ResSystemPostItem>>(PORT + `/system/post`, params);
};
// 职位列表数据精简
export const getSystemPostListSimpleApi = (params?: SystemPost.ReqSystemPostList) => {
  return http.get<SystemPost.ResSystemPostItem[]>(PORT + `/system/post/simple`, params);
};
