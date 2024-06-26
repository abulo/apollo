import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { Monitor } from "@/api/interface/monitor";

export const getMonitorRunApi = () => {
  return http.get<Monitor.MonitorRunItem>(PORT + `/monitor/run`, {}, { loading: false });
};

export const getMonitorTrendApi = () => {
  return http.get<Monitor.MonitorTrendItem>(PORT + `/monitor/trend`, {}, { loading: false });
};
