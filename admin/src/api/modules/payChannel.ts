import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { PayChannel } from "@/api/interface/payChannel";

// 渠道信息设置
export const setPayChannelApi = (payAppId: number, channelCode: string, params: PayChannel.ResPayChannelItem) => {
  return http.post(PORT + `/pay/app/${payAppId}/channel/${channelCode}`, params);
};

// 渠道信息获取
export const getPayChannelItemApi = (payAppId: number, channelCode: string) => {
  return http.get<PayChannel.ResPayChannelItem>(PORT + `/pay/app/${payAppId}/channel/${channelCode}/item`);
};
