import { ReqPage } from "./index";
export namespace SystemUser {
  // ReqLogin 登录请求
  export interface ReqSystemUserLogin {
    username: string; // 用户名
    password: string; // 密码
    captchaCode: string; // 验证码
    captchaId: string; // 验证码id
  }
  // ResLogin 登录返回的数据
  export interface ResSystemUserLogin {
    accessToken: string;
    nickname: string;
  }
  // 按钮权限
  export interface ResAuthButtons {
    [key: string]: string[];
  }

  // 列表接口请求数据
  export interface ReqSystemUserList extends ReqPage {
    username?: string; //用户名
    status?: number; // 帐号状态（0正常 1停用）
    deptId?: number; // 部门
  }
  export interface ResSystemUserItem {
    id: number | undefined; // 用户编号
    nickname: string; //昵称
    mobile: string; //手机号码
    username: string; //用户名称
    password: string | undefined; //用户密码
    status: number; // 用户状态（0正常 1停用）
    deleted: number; // 是否删除(0否 1是)
    deptIds: number[] | undefined;
    postIds: number[] | undefined;
    roleIds: number[] | undefined;
  }
  export interface ReqSystemUserPassword {
    password: string;
  }
}
