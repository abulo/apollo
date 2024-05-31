// pay_refund 退款订单
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { PayRefund } from "@/api/interface/payRefund";
// 退款订单创建数据
export const addPayRefundApi = (params: PayRefund.ResPayRefundItem) => {
  return http.post(PORT + `/pay/refund`, params);
};
// 退款订单更新数据
export const updatePayRefundApi = (id: number, params: PayRefund.ResPayRefundItem) => {
  return http.put(PORT + `/pay/refund/${id}/update`, params);
};
// 退款订单删除数据
export const deletePayRefundApi = (id: number) => {
  return http.delete(PORT + `/pay/refund/${id}/delete`);
};
// 退款订单查询单条数据
export const getPayRefundItemApi = (id: number) => {
  return http.get<PayRefund.ResPayRefundItem>(PORT + `/pay/refund/${id}/item`);
};
// 退款订单恢复数据
export const recoverPayRefundApi = (id: number) => {
  return http.put(PORT + `/pay/refund/${id}/recover`);
};
// 退款订单列表数据
export const getPayRefundListApi = (params?: PayRefund.ReqPayRefundList) => {
  return http.get<ResPage<PayRefund.ResPayRefundItem>>(PORT + `/pay/refund`, params);
};
