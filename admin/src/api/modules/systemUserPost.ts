import { SystemUserPost } from "@/api/interface/systemUserPost";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";

// 获取列表
export const getSystemUserPostListApi = (id: number) => {
  return http.get<SystemUserPost.ResSystemUserPostItem>(PORT + `/system/user/${id}/post`);
};

// 新增
export const addSystemUserPostApi = (id: number, params: SystemUserPost.ResSystemUserPostItem) => {
  return http.post(PORT + `/system/user/${id}/post`, params);
};
