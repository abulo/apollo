import { partial } from "filesize";

// 文件格式化
export const FileSize = (value: number) => {
  const size = partial({ base: 2, standard: "jedec" });
  return size(value);
};

// 文件地址
export const FileUrl = (url: string | undefined) => {
  if (!url) return "";
  const domain = import.meta.env.VITE_STATIC_URL as string;
  return domain + "/" + url;
};
