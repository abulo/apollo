import { SystemUser } from "@/api/interface/systemUser";
import { SystemMenu } from "@/api/interface/systemMenu";
import { PORT } from "@/api/config/servicePort";
import { ResPage } from "@/api/interface/index";
import http from "@/api";

/**
 * @name 用户模块
 */
// 用户登录
export const postSystemUserLoginApi = (params: SystemUser.ReqSystemUserLogin) => {
  return http.post<SystemUser.ResSystemUserLogin>(PORT + `/login`, params, { loading: false }); // 正常 post json 请求  ==>  application/json
  // return http.post<Login.ResLogin>(PORT1 + `/login`, params, { loading: false }); // 控制当前请求不显示 loading
  // return http.post<Login.ResLogin>(PORT1 + `/login`, {}, { params }); // post 请求携带 query 参数  ==>  ?username=admin&password=123456
  // return http.post<Login.ResLogin>(PORT1 + `/login`, qs.stringify(params)); // post 请求携带表单参数  ==>  application/x-www-form-urlencoded
  // return http.get<Login.ResLogin>(PORT1 + `/login?${qs.stringify(params, { arrayFormat: "repeat" })}`); // get 请求可以携带数组等复杂参数
};

// 获取列表
export const getSystemUserListApi = (params?: SystemUser.ReqSystemUserList) => {
  return http.get<ResPage<SystemUser.ResSystemUserItem>>(PORT + `/system/user`, params);
};

// 获取单个数据接口
export const getSystemUserItemApi = (id: number) => {
  return http.get<SystemUser.ResSystemUserItem>(PORT + `/system/user/${id}/item`);
};

// 新增
export const addSystemUserApi = (params: SystemUser.ResSystemUserItem) => {
  return http.post(PORT + `/system/user`, params);
};

// 修改
export const updateSystemUserApi = (id: number, params: SystemUser.ResSystemUserItem) => {
  return http.put(PORT + `/system/user/${id}/update`, params);
};

// 删除
export const deleteSystemUserApi = (id: number) => {
  return http.delete(PORT + `/system/user/${id}/delete`);
};

// // 密码秀
// export const setSystemUserPasswordApi = (id: number, params: SystemUser.ReqSystemUserPassword) => {
//   return http.put(PORT + `/system/user/${id}/password`, params);
// };

// 获取用户菜单
export const getSystemUserMenuApi = () => {
  return http.get<Menu.MenuOptions[]>(PORT + `/system/user/menu`, {}, { loading: false });
};

// 获取用户按钮
export const getSystemUserBtnApi = () => {
  return http.get<SystemMenu.ResSystemMenuButtons>(PORT + `/system/user/btn`, {}, { loading: false });
};
