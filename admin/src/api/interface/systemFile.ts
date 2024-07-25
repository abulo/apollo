// system_file 文件管理
import { ReqPage } from "./index";
export namespace SystemFile {
  export interface ReqSystemFileList extends ReqPage {
    deleted?: number; // 是否删除
    fileType?: number; // 文件类型
    fileName?: string; // 文件名称
    deptId?: number; // 部门
  }
  export interface ResSystemFileItem {
    id: number; // 编号
    fileName: string; // 文件名称
    fileType: number; // 文件类型
    fileMimeType: string; // Mime类型
    fileSize: number; // 文件大小
    filePath: string; // 文件路径
    userId: number; // 用户 ID
    deleted: number; // 是否删除
    tenantId: number; // 租户
    creator: string | undefined; // 创建人
    createTime: string | undefined; // 创建时间
    updater: string | undefined; // 更新人
    updateTime: string | undefined; // 更新时间
  }
}
