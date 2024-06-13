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
        <el-button v-auth="'logger.SystemLoginLogDrop'" type="danger" :icon="Delete" @click="handleDrop(scope.selectedListIds)">
          清空
        </el-button>
      </template>
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <template #operation="scope">
        <el-button v-auth="'logger.SystemLoginLog'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-button
          type="primary"
          link
          v-if="scope.row.deleted === 0"
          v-auth="'logger.SystemLoginLogDelete'"
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          type="primary"
          link
          v-if="scope.row.deleted === 1"
          v-auth="'logger.SystemLoginLogRecover'"
          :icon="Refresh"
          @click="handleRecover(scope.row)">
          恢复
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
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { Delete, View, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemLoginLog } from "@/api/interface/systemLoginLog";
import {
  getSystemLoginLogListApi,
  deleteSystemLoginLogApi,
  dropSystemLoginLogApi,
  recoverSystemLoginLogApi,
  getSystemLoginLogItemApi
} from "@/api/modules/systemLoginLog";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";

//删除状态
const deletedEnum = getIntDictOptions("delete");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("logger.SystemLoginLogDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);
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
    prop: "deleted",
    label: "删除",
    tag: true,
    enum: deletedEnum,
    search: deleteSearch,
    width: 100
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

// 删除按钮
const handleDelete = async (row: SystemLoginLog.ResSystemLoginLogItem) => {
  await useHandleData(deleteSystemLoginLogApi, Number(row.id), "删除日志");
  proTable.value?.getTableList();
};

// 恢复按钮
const handleRecover = async (row: SystemLoginLog.ResSystemLoginLogItem) => {
  await useHandleData(recoverSystemLoginLogApi, Number(row.id), "恢复日志");
  proTable.value?.getTableList();
};

const handleDrop = async (id: string[]) => {
  const data = ref<SystemLoginLog.ReqSystemLoginLogDrop>({});
  let newParams = proTable.value?.searchParam;
  let time = newParams?.loginTime as string[];
  time && (data.value.beginLoginTime = time[0]);
  time && (data.value.finishLoginTime = time[1]);
  // 判断 ids 是不是空
  if (id.length > 0) {
    data.value.ids = id.map(Number);
  }
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
  deleted: 0,
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
    deleted: 0,
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
