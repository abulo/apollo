export namespace SystemMailAccount {
  export interface ReqSystemMailAccountList {
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
    deleted: number; // 是否删除
  }
}
