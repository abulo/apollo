export namespace SystemFile {
  export interface ReqSystemFileList {
    fileType?: number; //文件类型
    fileName?: string; //文件名称
  }
  export interface ResSystemFileItem {
    id: number; //编号
    fileName: string; //文件名称
    fileType: number; //文件类型
    fileMimeType: string; //Mime类型
    fileSize: number; //文件大小
    filePath: string; //文件路径
  }
}
