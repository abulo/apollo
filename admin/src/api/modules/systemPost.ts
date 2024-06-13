import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { ResPage } from "@/api/interface/index";
import { SystemPost } from "@/api/interface/systemPost";
// 获取职位列表
export const getSystemPostListApi = (params?: SystemPost.ReqSystemPostList) => {
  return http.get<ResPage<SystemPost.ResSystemPostItem>>(PORT + `/system/post`, params);
};
// 获取单个职位
export const getSystemPostItemApi = (id: number) => {
  return http.get<SystemPost.ResSystemPostItem>(PORT + `/system/post/${id}/item`);
};
// 添加职位
export const addSystemPostApi = (params: SystemPost.ResSystemPostItem) => {
  return http.post(PORT + `/system/post`, params);
};
// 修改职位
export const updateSystemPostApi = (id: number, params: SystemPost.ResSystemPostItem) => {
  return http.put(PORT + `/system/post/${id}/update`, params);
};
// 删除职位
export const deleteSystemPostApi = (id: number) => {
  return http.delete(PORT + `/system/post/${id}/delete`);
};
// 恢复职位
export const recoverSystemPostApi = (id: number) => {
  return http.put(PORT + `/system/post/${id}/recover`);
};
// 获取职位树
export const getSystemPostListSimpleApi = (params?: SystemPost.ReqSystemPostList) => {
  return http.get<SystemPost.ResSystemPostItem[]>(PORT + `/system/post/simple`, params);
};
