// pay_order 支付订单
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { PayOrder } from "@/api/interface/payOrder";
// 支付订单创建数据
export const addPayOrderApi = (params: PayOrder.ResPayOrderItem) => {
  return http.post(PORT + `/pay/order`, params);
};
// 支付订单更新数据
export const updatePayOrderApi = (id: number, params: PayOrder.ResPayOrderItem) => {
  return http.put(PORT + `/pay/order/${id}/update`, params);
};
// 支付订单删除数据
export const deletePayOrderApi = (id: number) => {
  return http.delete(PORT + `/pay/order/${id}/delete`);
};
// 支付订单查询单条数据
export const getPayOrderItemApi = (id: number) => {
  return http.get<PayOrder.ResPayOrderItem>(PORT + `/pay/order/${id}/item`);
};
// 支付订单恢复数据
export const recoverPayOrderApi = (id: number) => {
  return http.put(PORT + `/pay/order/${id}/recover`);
};
// 支付订单列表数据
export const getPayOrderListApi = (params?: PayOrder.ReqPayOrderList) => {
  return http.get<ResPage<PayOrder.ResPayOrderItem>>(PORT + `/pay/order`, params);
};
