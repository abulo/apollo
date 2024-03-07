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
  }
  export interface ResSystemUserItem {
    id: number | undefined; // 用户编号
    nickname: string; // 昵称
    username: string; // 用户名称
    password: string | undefined; // 用户密码
    status: number; // 用户状态（0正常 1停用）
    creator: string; // 创建人
    createTime: string; // 创建时间
    updater: string; // 更新人
    updateTime: string; // 更新时间
    systemRoleIds: number[]; // 角色ID
  }
}
