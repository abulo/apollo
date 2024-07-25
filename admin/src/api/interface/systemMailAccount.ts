// system_mail_account 邮箱账号表
import { ReqPage } from "./index";
export namespace SystemMailAccount {
  export interface ReqSystemMailAccountList extends ReqPage {
    deleted?: number; // 是否删除
    mail?: string; // 邮箱
    username?: string; // 用户名
  }
  export interface ResSystemMailAccountItem {
    id: number; // 主键
    mail: string; // 邮箱
    username: string; // 用户名
    password: string; // 密码
    host: string; // SMTP 服务器域名
    port: number; // SMTP 服务器端口
    sslEnable: number; // 是否开启 SSL
    creator: string | undefined; // 创建者
    createTime: string; // 创建时间
    updater: string | undefined; // 更新者
    updateTime: string; // 更新时间
    deleted: number; // 是否删除
  }
}
