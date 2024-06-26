import { IFetchOptions } from "@/api/sse/interface";
import { parseServerSentEvent } from "@/api/sse/sse";
import { checkOk } from "@/api/sse/utils";
import { PORT } from "@/api/config/servicePort";
import { useUserStore } from "@/stores/modules/user";
import { generateUUID } from "@/utils/index";

export async function fetchEventData(url: string, options: IFetchOptions = {}): Promise<void> {
  const { method, data, headers = {}, signal, onMessage, onError, onOpen, onClose } = options;

  // 默认地址请求地址，可在 .env.** 文件中修改
  const baseURL = import.meta.env.VITE_API_URL as string;

  const defaultHeaders = {
    Accept: "text/event-stream",
    "Content-Type": "application/json"
  };
  const mergedHeaders = {
    ...defaultHeaders,
    ...headers
  };
  const userStore = useUserStore();
  // mergedHeaders需要添加 X-Access-Token 头部
  mergedHeaders["X-Access-Token"] = userStore.token;
  // mergedHeaders需要添加 Last-Event-ID 头部
  mergedHeaders["Last-Event-ID"] = generateUUID();
  let fetchUrl = baseURL + PORT + url;
  // 这里需要替换一下, 除了 http://和https://的字符,  其余字符串中包含了两个或两个以上的/的, 都替换成一个/
  const reg = /([^:]\/)\/+/g;
  fetchUrl = fetchUrl.replace(reg, "$1");
  try {
    const res = await fetch(fetchUrl, {
      method,
      headers: mergedHeaders,
      body: JSON.stringify(data),
      signal: signal
    });
    await checkOk(res);
    onOpen?.(res);
    if (typeof onMessage === "function" && res.body) {
      await parseServerSentEvent(res.body, event => {
        onMessage(event);
      });
      onClose?.();
    }
  } catch (err) {
    onError?.(err);
  }
}
