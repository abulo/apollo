export namespace Monitor {
  export interface MonitorRunItem {
    arch: string; // 架构
    buildTime: string; // 构建时间
    buildVersion: string; // 构建版本
    goArch: string; // go 架构
    goMem: string; // go 内存
    goSize: string; // go 大小
    goVersion: string; // go 版本
    goroutine: number; // go 协程数
    hostname: string; // 主机名
    intranetIP: string; // 内网IP
    launchTime: string; // 启动时间
    os: string; // 操作系统
    platform: string; // 平台
    publicIP: string; // 公网IP
    pwd: string; // 运行路径
    ratelVersion: string; // ratel 版本
    rootPath: string; // 根路径
    runTime: number; // 运行时间
  }

  export interface MonitorTrendItem {
    head: Head[];
    load: Load[];
    net: Net[];
  }

  export interface Head {
    title: string;
    data: string;
    bottomTitle: string;
    totalSum: string;
    iconClass: string;
    extra: Extra;
  }

  export interface Extra {
    data?: string;
    data1?: string;
  }

  export interface Load {
    time: string;
    avg: number;
    ratio: number;
  }

  export interface Net {
    time: string;
    bytesSent: string;
    bytesRecv: string;
    down: number;
    up: number;
  }
}
