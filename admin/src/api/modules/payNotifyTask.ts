// pay_notify_task 商户支付-任务通知
import { ResPage } from "@/api/interface/index";
import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { PayNotifyTask } from "@/api/interface/payNotifyTask";
// 商户支付-任务通知删除数据
export const deletePayNotifyTaskApi = (id: number) => {
  return http.delete(PORT + `/pay/notify/${id}/delete`);
};
// 商户支付-任务通知查询单条数据
export const getPayNotifyTaskItemApi = (id: number) => {
  return http.get<PayNotifyTask.ResPayNotifyTaskItem>(PORT + `/pay/notify/${id}/item`);
};
// 商户支付-任务通知恢复数据
export const recoverPayNotifyTaskApi = (id: number) => {
  return http.put(PORT + `/pay/notify/${id}/recover`);
};
// 商户支付-任务通知列表数据
export const getPayNotifyTaskListApi = (params?: PayNotifyTask.ReqPayNotifyTaskList) => {
  return http.get<ResPage<PayNotifyTask.ResPayNotifyTaskItem>>(PORT + `/pay/notify`, params);
};
