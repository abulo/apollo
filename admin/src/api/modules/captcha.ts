import { Captcha } from "@/api/interface/captcha";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";

// captchaGenerateApi 调用服务端验证码生成
export const captchaGenerateApi = () => {
  return http.get<Captcha.ResCaptcha>(PORT + `/captcha`, {}, { loading: false });
};
