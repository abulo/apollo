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
        <el-button v-auth="'logger.SystemEntryLogDrop'" type="danger" :icon="Delete" @click="handleDrop(scope.selectedListIds)">
          清空
        </el-button>
      </template>
      <template #operation="scope">
        <el-button v-auth="'logger.SystemEntryLog'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-button type="primary" link v-auth="'logger.SystemEntryLogDelete'" :icon="Delete" @click="handleDelete(scope.row)">
          删除
        </el-button>
      </template>
    </ProTable>
    <el-dialog
      v-model="centerDialogVisible"
      :title="title"
      width="40%"
      destroy-on-close
      align-center
      center
      append-to-body
      draggable
      :lock-scroll="false"
      class="dialog-settings">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="日志主键" min-width="120">
          {{ systemEntryLogItemFrom.id }}
        </el-descriptions-item>
        <el-descriptions-item label="服务器">
          {{ systemEntryLogItemFrom.host }}
        </el-descriptions-item>
        <el-descriptions-item label="日志信息">
          {{ systemEntryLogItemFrom.message }}
        </el-descriptions-item>
        <el-descriptions-item label="日志等级">
          {{ systemEntryLogItemFrom.level }}
        </el-descriptions-item>
        <el-descriptions-item label="文件">
          {{ systemEntryLogItemFrom.file }}
        </el-descriptions-item>
        <el-descriptions-item label="方法">
          {{ systemEntryLogItemFrom.func }}
        </el-descriptions-item>
        <el-descriptions-item label="数据">
          {{ systemEntryLogItemFrom.data }}
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx" name="systemLoggerEntry">
import { ref } from "vue";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import { Delete, View } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemEntryLog } from "@/api/interface/systemEntryLog";
import {
  getSystemEntryLogListApi,
  deleteSystemEntryLogApi,
  dropSystemEntryLogApi,
  getSystemEntryLogItemApi
} from "@/api/modules/systemEntryLog";
import { useHandleData } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
//table数据
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns: ColumnProps<SystemEntryLog.ResSystemEntryLogItem>[] = [
  { type: "selection", fixed: "left", width: 70 },
  { prop: "id", label: "编号" },
  { prop: "host", label: "服务器" },
  { prop: "message", label: "日志信息" },
  { prop: "level", label: "日志等级" },
  {
    prop: "timestamp",
    label: "日志时间",
    search: {
      el: "date-picker",
      span: 4,
      props: { type: "datetimerange", valueFormat: "YYYY-MM-DD HH:mm:ss" }
    }
  },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission("logger.SystemEntryLog")
  }
];
const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.timestamp && (newParams.startTime = newParams.timestamp[0]);
  newParams.timestamp && (newParams.endTime = newParams.timestamp[1]);
  delete newParams.timestamp;
  return getSystemEntryLogListApi(newParams);
};

// 删除按钮
const handleDelete = async (row: SystemEntryLog.ResSystemEntryLogItem) => {
  await useHandleData(deleteSystemEntryLogApi, String(row.id), "删除日志");
  proTable.value?.getTableList();
};

const handleDrop = async (id: string[]) => {
  const data = ref<SystemEntryLog.ReqSystemEntryLogDrop>({});
  let newParams = proTable.value?.searchParam;
  let time = newParams?.timestamp as string[];
  time && (data.value.startTime = time[0]);
  time && (data.value.endTime = time[1]);
  // 判断 ids 是不是空
  if (id.length > 0) {
    data.value.ids = id.map(String);
  }
  await useHandleData(dropSystemEntryLogApi, data.value, "清空日志");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

const systemEntryLogItemFrom = ref<SystemEntryLog.ResSystemEntryLogItem>({
  id: "",
  host: "",
  timestamp: "",
  file: "",
  func: "",
  message: "",
  data: undefined,
  level: ""
});
//弹出层标题
const title = ref();
//是否显示弹出层
const centerDialogVisible = ref(false);

// 重置数据
const reset = () => {
  systemEntryLogItemFrom.value = {
    id: "",
    host: "",
    timestamp: "",
    file: "",
    func: "",
    message: "",
    data: undefined,
    level: ""
  };
};

// 编辑按钮
const handleItem = async (row: SystemEntryLog.ResSystemEntryLogItem) => {
  title.value = "查看信息";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemEntryLogItemApi(row.id);
  systemEntryLogItemFrom.value = data;
};
</script>
<style lang="scss">
@import "@/styles/custom.scss";
</style>
