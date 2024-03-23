<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="登录日志列表"
      row-key="id"
      :columns="columns"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader="scope">
        <el-button
          type="danger"
          :icon="Delete"
          plain
          :disabled="!scope.isSelected"
          @click="handleDelete(scope.selectedListIds)"
          v-auth="'logger.SystemLoginLogDelete'">
          删除
        </el-button>
        <el-button type="danger" :icon="Delete" @click="handleDrop" v-auth="'logger.SystemLoginLogDrop'"> 清空 </el-button>
      </template>
      <template #operation="scope">
        <el-button type="primary" link :icon="View" @click="handleItem(scope.row)" v-auth="'logger.SystemLoginLog'">
          查看
        </el-button>
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
        <el-descriptions-item label="日志主键" min-width="120">
          {{ systemLoginLogItemFrom.id }}
        </el-descriptions-item>
        <el-descriptions-item label="用户账号">
          {{ systemLoginLogItemFrom.username }}
        </el-descriptions-item>
        <el-descriptions-item label="用户ip">
          {{ systemLoginLogItemFrom.userIp }}
        </el-descriptions-item>
        <el-descriptions-item label="UA">
          {{ systemLoginLogItemFrom.userAgent }}
        </el-descriptions-item>
        <el-descriptions-item label="渠道">
          {{ systemLoginLogItemFrom.channel }}
        </el-descriptions-item>
        <el-descriptions-item label="登录时间">
          {{ systemLoginLogItemFrom.loginTime }}
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx" name="systemLoggerLogin">
import { ref } from "vue";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import { Delete, View } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemLoginLog } from "@/api/interface/systemLoginLog";
import {
  getSystemLoginLogListApi,
  deleteSystemLoginLogApi,
  dropSystemLoginLogApi,
  getSystemLoginLogItemApi
} from "@/api/modules/systemLoginLog";
import { useHandleData } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
// 表格配置项
const columns: ColumnProps<SystemLoginLog.ResSystemLoginLogItem>[] = [
  { type: "selection", fixed: "left", width: 70 },
  { prop: "id", label: "编号" },
  { prop: "username", label: "用户账号", search: { el: "input", span: 2 } },
  { prop: "userIp", label: "用户ip" },
  { prop: "userAgent", label: "UA" },
  { prop: "channel", label: "渠道", search: { el: "input", span: 2 } },
  {
    prop: "loginTime",
    label: "登录时间",
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
    isShow: HasPermission("logger.SystemLoginLog")
  }
];

const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.loginTime && (newParams.beginLoginTime = newParams.loginTime[0]);
  newParams.loginTime && (newParams.finishLoginTime = newParams.loginTime[1]);
  delete newParams.loginTime;
  return getSystemLoginLogListApi(newParams);
};
//table数据
const proTable = ref<ProTableInstance>();

// 批量删除用户信息
const handleDelete = async (id: string[]) => {
  const data = ref<SystemLoginLog.ReqSystemLoginLogDelete>({
    systemLoginLogIds: id.map(Number)
  });
  await useHandleData(deleteSystemLoginLogApi, data.value, "删除所选信息");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

const handleDrop = async () => {
  const data = ref<SystemLoginLog.ReqSystemLoginLogDrop>({});
  let newParams = proTable.value?.searchParam;
  let time = newParams?.loginTime as string[];
  time && (data.value.beginLoginTime = time[0]);
  time && (data.value.finishLoginTime = time[1]);
  await useHandleData(dropSystemLoginLogApi, data.value, "清空日志");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

const systemLoginLogItemFrom = ref<SystemLoginLog.ResSystemLoginLogItem>({
  id: 0,
  username: "",
  userIp: "",
  userAgent: "",
  loginTime: "",
  channel: "",
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
  systemLoginLogItemFrom.value = {
    id: 0,
    username: "",
    userIp: "",
    userAgent: "",
    loginTime: "",
    channel: "",
    creator: "",
    createTime: "",
    updater: "",
    updateTime: ""
  };
};

// 编辑按钮
const handleItem = async (row: SystemLoginLog.ResSystemLoginLogItem) => {
  title.value = "查看信息";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemLoginLogItemApi(row.id);
  systemLoginLogItemFrom.value = data;
};
</script>
<style lang="scss">
@import "@/styles/custom.scss";
</style>
