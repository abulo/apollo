<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="商户支付-任务通知列表"
      row-key="id"
      :columns="columns"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button type="primary" link :icon="View" @click="handleItem(scope.row)"> 查看 </el-button>
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'notify.PayNotifyTaskDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'notify.PayNotifyTaskRecover'"
          type="primary"
          link
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
      <el-descriptions :column="2" border>
        <el-descriptions-item label="商户订单编号">
          <el-tag>{{ payNotifyTaskItemFrom.merchantOrderId }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="通知状态">
          <DictTag type="pay.notifyStatus" :value="payNotifyTaskItemFrom.status" />
        </el-descriptions-item>

        <el-descriptions-item label="应用编号">{{ payNotifyTaskItemFrom.appId }}</el-descriptions-item>
        <el-descriptions-item label="应用名称">
          {{ payAppList.find(item => item.id === payNotifyTaskItemFrom.appId)?.name }}
        </el-descriptions-item>

        <el-descriptions-item label="关联编号">{{ payNotifyTaskItemFrom.dataId }}</el-descriptions-item>
        <el-descriptions-item label="通知类型">
          <DictTag type="pay.notifyType" :value="payNotifyTaskItemFrom.type" />
        </el-descriptions-item>

        <el-descriptions-item label="通知次数">{{ payNotifyTaskItemFrom.notifyTimes }}</el-descriptions-item>
        <el-descriptions-item label="最大通知次数">
          {{ payNotifyTaskItemFrom.maxNotifyTimes }}
        </el-descriptions-item>

        <el-descriptions-item label="最后通知时间">
          {{ payNotifyTaskItemFrom.lastExecuteTime }}
        </el-descriptions-item>
        <el-descriptions-item label="下次通知时间">
          {{ payNotifyTaskItemFrom.nextNotifyTime }}
        </el-descriptions-item>

        <el-descriptions-item label="创建时间">
          {{ payNotifyTaskItemFrom.createTime }}
        </el-descriptions-item>
        <el-descriptions-item label="更新时间">
          {{ payNotifyTaskItemFrom.updateTime }}
        </el-descriptions-item>
      </el-descriptions>
      <el-divider />
      <el-descriptions :column="1" direction="vertical" border>
        <el-descriptions-item label="回调日志">
          <el-table :data="payNotifyTaskItemFrom.logs">
            <el-table-column label="日志编号" align="center" prop="id" />
            <el-table-column label="通知状态" align="center" prop="status">
              <template #default="scope">
                <DictTag type="pay.notifyStatus" :value="scope.row.status" />
              </template>
            </el-table-column>
            <el-table-column label="通知次数" align="center" prop="notifyTimes" />
            <el-table-column label="通知时间" align="center" prop="lastExecuteTime" width="180">
              <template #default="scope">
                <span>{{ scope.row.createTime }}</span>
              </template>
            </el-table-column>
            <el-table-column label="响应结果" align="center" prop="response" />
          </el-table>
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="payNotifyTask">
import { ref, reactive, onMounted } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { View, CirclePlus, Delete, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { PayNotifyTask } from "@/api/interface/payNotifyTask";
import {
  getPayNotifyTaskListApi,
  deletePayNotifyTaskApi,
  recoverPayNotifyTaskApi,
  getPayNotifyTaskItemApi
} from "@/api/modules/payNotifyTask";
import { PayApp } from "@/api/interface/payApp";
import { getPayAppListApi } from "@/api/modules/payApp";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
//加载
const loading = ref(false);
//弹出层标题
const title = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
//数据接口
const payNotifyTaskItemFrom = ref<PayNotifyTask.ResPayNotifyTaskItem>({
  id: 0, // 任务编号
  appId: 0, // 应用编号
  type: 0, // 通知类型
  dataId: 0, // 数据编号
  status: 0, // 通知状态
  merchantOrderId: "", // 商户订单编号
  merchantTransferId: "", // 商户转账单编号
  nextNotifyTime: "", // 下一次通知时间
  lastExecuteTime: "", // 最后一次执行时间
  notifyTimes: 0, // 当前通知次数
  maxNotifyTimes: undefined, // 最大可通知次数
  notifyUrl: "", // 异步通知地址
  deleted: 0, // 删除
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
  logs: undefined // 日志
});

//应用列表
const payAppList = ref<PayApp.ResPayAppItem[]>([]);
//删除状态
const deletedEnum = getIntDictOptions("delete");
//支付通知类型
const payNotifyTypeEnum = getIntDictOptions("pay.notifyType");
//支付通知状态
const payNotifyStatusEnum = getIntDictOptions("pay.notifyStatus");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("notify.PayNotifyTaskDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);

const columns: ColumnProps<PayNotifyTask.ResPayNotifyTaskItem>[] = [
  { prop: "id", label: "任务编号" },
  {
    prop: "appId",
    label: "应用编号",
    enum: payAppList,
    fieldNames: { label: "name", value: "id" },
    search: { el: "select", span: 2 }
  },
  { prop: "type", label: "通知类型", enum: payNotifyTypeEnum, search: { el: "select", span: 2 } },
  { prop: "dataId", label: "数据编号", search: { el: "input", span: 2, props: { placeholder: "请输入数据编号" } } },
  { prop: "status", label: "通知状态", enum: payNotifyStatusEnum, search: { el: "select", span: 2 } },
  {
    prop: "merchantOrderId",
    label: "商户订单编号",
    search: { el: "input", span: 2, props: { placeholder: "请输入商户订单编号" } }
  },
  { prop: "deleted", label: "删除", tag: true, enum: deletedEnum, search: deleteSearch },
  {
    prop: "createTime",
    label: "创建时间",
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
    isShow: HasPermission("notify.PayNotifyTaskDelete", "notify.PayNotifyTaskRecover")
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  payNotifyTaskItemFrom.value = {
    id: 0, // 任务编号
    appId: 0, // 应用编号
    type: 0, // 通知类型
    dataId: 0, // 数据编号
    status: 0, // 通知状态
    merchantOrderId: "", // 商户订单编号
    merchantTransferId: "", // 商户转账单编号
    nextNotifyTime: "", // 下一次通知时间
    lastExecuteTime: "", // 最后一次执行时间
    notifyTimes: 0, // 当前通知次数
    maxNotifyTimes: undefined, // 最大可通知次数
    notifyUrl: "", // 异步通知地址
    deleted: 0, // 删除
    tenantId: 0, // 租户
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined, // 更新时间
    logs: undefined // 日志
  };
};
// 删除按钮
const handleDelete = async (row: PayNotifyTask.ResPayNotifyTaskItem) => {
  await useHandleData(deletePayNotifyTaskApi, Number(row.id), "删除任务通知");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: PayNotifyTask.ResPayNotifyTaskItem) => {
  await useHandleData(recoverPayNotifyTaskApi, Number(row.id), "恢复任务通知");
  proTable.value?.getTableList();
};

// 编辑按钮
const handleItem = async (row: PayNotifyTask.ResPayNotifyTaskItem) => {
  title.value = "任务通知";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getPayNotifyTaskItemApi(Number(row.id));
  payNotifyTaskItemFrom.value = data;
};

const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.createTime && (newParams.beginCreateTimeTime = newParams.createTime[0]);
  newParams.createTime && (newParams.finishCreateTimeTime = newParams.createTime[1]);
  delete newParams.createTime;
  return getPayNotifyTaskListApi(newParams);
};

onMounted(async () => {
  const { data } = await getPayAppListApi();
  payAppList.value = data;
});
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
