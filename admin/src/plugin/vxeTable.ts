import { App } from "vue";

import VxeTable from "vxe-table";
import "vxe-table/lib/style.css";

import VxeUI from "vxe-pc-ui";
import "vxe-pc-ui/lib/style.css";
// import { VXETable, Icon, Column, Table } from "vxe-table";
// import VXETablePluginElement from "vxe-table-plugin-element";
// import "vxe-table/lib/style.css";
// import "vxe-table-plugin-element/dist/style.css";

// // 全局默认参数
// VXETable.config({});
// VXETable.use(VXETablePluginElement);

export function useVxeTable(app: App) {
  app.use(VxeUI).use(VxeTable);
}
