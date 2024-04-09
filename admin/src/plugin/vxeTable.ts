import type { App } from "vue";

import { VXETable, Icon, Column, Table } from "vxe-table";
import VXETablePluginElement from "vxe-table-plugin-element";
import "vxe-table/lib/style.css";
import "vxe-table-plugin-element/dist/style.css";

// 全局默认参数
VXETable.config({});
VXETable.use(VXETablePluginElement);

export function useVxeTable(app: App) {
  app.use(Icon).use(Column).use(Table);
}
