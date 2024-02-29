<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="系统日志列表"
      row-key="id"
      :columns="columns"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader="scope">
        <el-button type="danger" :icon="Delete" plain :disabled="!scope.isSelected" @click="handleDelete(scope.selectedListIds)">
          删除
        </el-button>
        <el-button type="danger" :icon="Delete" @click="handleDrop"> 清空 </el-button>
      </template>
      <template #operation="scope">
        <el-button type="primary" link :icon="View" @click="handleItem(scope.row)">查看</el-button>
      </template>
    </ProTable>
    <el-dialog
      :title="title"
      v-model="centerDialogVisible"
      width="40%"
      destroy-on-close
      align-center
      center
      append-to-body
      draggable
      :lock-scroll="false"
      class="dialog-settings">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="主键" min-width="120">
          {{ systemOperateLogItemFrom.id }}
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx" name="systemLoggerOperate">
import { ref } from "vue";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import { Delete, View } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemOperateLog } from "@/api/interface/systemOperateLog";
import {
  getSystemOperateLogListApi,
  deleteSystemOperateLogApi,
  dropSystemOperateLogApi,
  getSystemOperateLogItemApi
} from "@/api/modules/systemOperateLog";
import { useHandleData } from "@/hooks/useHandleData";
//table数据
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns: ColumnProps<SystemOperateLog.ResSystemOperateLogItem>[] = [
  { type: "selection", fixed: "left", width: 70 },
  { prop: "id", label: "编号" },
  { prop: "username", label: "用户账号" },
  { prop: "module", label: "模块名称" },
  { prop: "requestMethod", label: "请求方法" },
  { prop: "requestUrl", label: "请求地址" },
  { prop: "userIp", label: "用户ip" },
  { prop: "userAgent", label: "UA" },
  { prop: "goMethod", label: "方法" },
  {
    prop: "startTime",
    label: "日志时间",
    search: {
      el: "date-picker",
      span: 4,
      props: { type: "datetimerange", valueFormat: "YYYY-MM-DD HH:mm:ss" }
    }
  },
  { prop: "duration", label: "耗时" },
  { prop: "channel", label: "渠道" },
  { prop: "result", label: "结果" },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right"
  }
];
const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.timestamp && (newParams.beginStartTime = newParams.timestamp[0]);
  newParams.timestamp && (newParams.finishStartTime = newParams.timestamp[1]);
  delete newParams.timestamp;
  return getSystemOperateLogListApi(newParams);
};

// 批量删除用户信息
const handleDelete = async (id: string[]) => {
  const data = ref<SystemOperateLog.ReqSystemOperateLogDelete>({
    systemOperateLogIds: id
  });
  await useHandleData(deleteSystemOperateLogApi, data.value, "删除所选信息");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

const handleDrop = async () => {
  const data = ref<SystemOperateLog.ReqSystemOperateLogDrop>({});
  let newParams = proTable.value?.searchParam;
  let time = newParams?.timestamp as string[];
  time && (data.value.beginStartTime = time[0]);
  time && (data.value.finishStartTime = time[1]);
  await useHandleData(dropSystemOperateLogApi, data.value, "清空日志");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

const systemOperateLogItemFrom = ref<SystemOperateLog.ResSystemOperateLogItem>({
  id: 0,
  username: "",
  module: "",
  requestMethod: "",
  requestUrl: "",
  userIp: "",
  userAgent: "",
  goMethod: "",
  goMethodArgs: "",
  startTime: "",
  duration: 0,
  channel: "",
  result: 0,
  creator: "",
  createTime: "",
  updater: "",
  updateTime: ""
});
//弹出层标题
const title = ref();
//是否显示弹出层
const centerDialogVisible = ref(false);

// 重置数据
const reset = () => {
  systemOperateLogItemFrom.value = {
    id: 0,
    username: "",
    module: "",
    requestMethod: "",
    requestUrl: "",
    userIp: "",
    userAgent: "",
    goMethod: "",
    goMethodArgs: "",
    startTime: "",
    duration: 0,
    channel: "",
    result: 0,
    creator: "",
    createTime: "",
    updater: "",
    updateTime: ""
  };
};

// 编辑按钮
const handleItem = async (row: SystemOperateLog.ResSystemOperateLogItem) => {
  title.value = "查看信息";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemOperateLogItemApi(row.id);
  systemOperateLogItemFrom.value = data;
};
</script>
