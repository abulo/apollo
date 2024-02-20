import { UPLOAD } from "@/api/config/servicePort";
export const getFileUrl = (url: string | undefined) => {
  const domain = import.meta.env.VITE_UPLOAD_URL as string;
  return domain + "/" + UPLOAD + "/" + url;
};
