<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="邮件日志列表"
      row-key="id"
      :columns="columns"
      :request-api="getCustomSystemMailLogListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader="scope">
        <el-dropdown trigger="click" class="table-el-dropdown">
          <el-button
            type="danger"
            v-auth="['mail.SystemMailLogMultipleDelete', 'mail.SystemMailLogMultipleRecover', 'mail.SystemMailLogMultipleDrop']"
            :icon="ArrowDownBold"
            plain
            :disabled="!scope.isSelected">
            操作
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item
                v-auth="'mail.SystemMailLogMultipleDelete'"
                :icon="Delete"
                @click="handleMultipleDelete(scope.selectedListIds)">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                v-auth="'mail.SystemMailLogMultipleDrop'"
                :icon="DeleteFilled"
                @click="handleMultipleDrop(scope.selectedListIds)">
                清理
              </el-dropdown-item>
              <el-dropdown-item
                v-auth="'mail.SystemMailLogMultipleRecover'"
                :icon="Refresh"
                @click="handleMultipleRecover(scope.selectedListIds)">
                恢复
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'mail.SystemMailLog'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['mail.SystemMailLogDelete', 'mail.SystemMailLogRecover', 'mail.SystemMailLogDrop']"
            type="primary"
            link
            :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item
                v-if="scope.row.deleted === 0"
                v-auth="'mail.SystemMailLogDelete'"
                :icon="Delete"
                @click="handleDelete(scope.row)">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'mail.SystemMailLogRecover'"
                :icon="Refresh"
                @click="handleRecover(scope.row)">
                恢复
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'mail.SystemMailLogDrop'"
                :icon="DeleteFilled"
                @click="handleDrop(scope.row)">
                清理
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
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
        <el-descriptions-item label="编号" min-width="120">
          {{ systemMailLogItemFrom.id }}
        </el-descriptions-item>
        <el-descriptions-item label="用户编号" min-width="120">
          {{ systemMailLogItemFrom.userId }}
        </el-descriptions-item>
        <el-descriptions-item label="用户类型" min-width="120">
          <DictTag type="notifyMessage.userType" :value="String(systemMailLogItemFrom.userType)" />
        </el-descriptions-item>
        <el-descriptions-item label="接收邮箱地址" min-width="120">
          {{ systemMailLogItemFrom.toMail }}
        </el-descriptions-item>
        <el-descriptions-item label="邮箱账号" min-width="120">
          {{ mailAccount.find(item => item.id === systemMailLogItemFrom.accountId)?.mail }}
        </el-descriptions-item>
        <el-descriptions-item label="发送邮箱地址" min-width="120">
          {{ systemMailLogItemFrom.fromMail }}
        </el-descriptions-item>
        <el-descriptions-item label="模板编码" min-width="120">
          {{ mailTemplate.find(item => item.code === systemMailLogItemFrom.templateCode)?.name }}
        </el-descriptions-item>
        <el-descriptions-item label="邮件标题" min-width="120">
          {{ systemMailLogItemFrom.templateTitle }}
        </el-descriptions-item>
        <el-descriptions-item label="发送人名称" min-width="120">
          {{ systemMailLogItemFrom.templateNickname }}
        </el-descriptions-item>
        <el-descriptions-item label="邮件内容" min-width="120">
          {{ systemMailLogItemFrom.templateContent }}
        </el-descriptions-item>
        <el-descriptions-item label="邮件参数" min-width="120">
          {{ systemMailLogItemFrom.templateParams }}
        </el-descriptions-item>
        <el-descriptions-item label="发送状态" min-width="120">
          <DictTag type="mailLog.sendStatus" :value="systemMailLogItemFrom.sendStatus" />
        </el-descriptions-item>
        <el-descriptions-item label="发送时间" min-width="120">
          {{ systemMailLogItemFrom.sendTime }}
        </el-descriptions-item>
        <el-descriptions-item label="发送返回的消息 ID" min-width="120">
          {{ systemMailLogItemFrom.sendMessageId }}
        </el-descriptions-item>
        <el-descriptions-item label="发送异常" min-width="120">
          {{ systemMailLogItemFrom.sendException }}
        </el-descriptions-item>
        <el-descriptions-item label="是否删除" min-width="120">
          <DictTag type="delete" :value="systemMailLogItemFrom.deleted" />
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemMailLog">
import { ref, reactive, onMounted } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { Delete, Refresh, DeleteFilled, View, DArrowRight, ArrowDownBold } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemMailLog } from "@/api/interface/systemMailLog";
import {
  getSystemMailLogListApi,
  deleteSystemMailLogApi,
  deleteSystemMailLogMultipleApi,
  dropSystemMailLogApi,
  dropSystemMailLogMultipleApi,
  recoverSystemMailLogApi,
  recoverSystemMailLogMultipleApi,
  getSystemMailLogApi
} from "@/api/modules/systemMailLog";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import { SystemMailAccount } from "@/api/interface/systemMailAccount";
import { getSystemMailAccountListSimpleApi } from "@/api/modules/systemMailAccount";
import { SystemMailTemplate } from "@/api/interface/systemMailTemplate";
import { getSystemMailTemplateListSimpleApi } from "@/api/modules/systemMailTemplate";
//弹出层标题
const title = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
//数据接口
const systemMailLogItemFrom = ref<SystemMailLog.ResSystemMailLogItem>({
  id: 0, // 编号
  userId: undefined, // 用户编号
  userType: undefined, // 用户类型
  toMail: "", // 接收邮箱地址
  accountId: 0, // 邮箱账号编号
  fromMail: "", // 发送邮箱地址
  templateId: 0, // 模板编号
  templateCode: "", // 模板编码
  templateNickname: undefined, // 模版发送人名称
  templateTitle: "", // 邮件标题
  templateContent: "", // 邮件内容
  templateParams: undefined, // 邮件参数
  sendStatus: 0, // 发送状态
  sendTime: undefined, // 发送时间
  sendMessageId: undefined, // 发送返回的消息 ID
  sendException: undefined, // 发送异常
  creator: undefined, // 创建者
  createTime: "", // 创建时间
  updater: undefined, // 更新者
  updateTime: "", // 更新时间
  deleted: 0 // 是否删除
});
//删除状态
const deletedEnum = getIntDictOptions("delete");
//会员类型
const userTypeEnum = getIntDictOptions("notifyMessage.userType");
// 阅读状态
const sendStatus = getIntDictOptions("mailLog.sendStatus");
// 邮箱账号
const mailAccount = ref<SystemMailAccount.ResSystemMailAccountItem[]>([]);
// 模板
const mailTemplate = ref<SystemMailTemplate.ResSystemMailTemplateItem[]>([]);
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("mail.SystemMailLogDelete")
    ? {
        el: "switch",
        span: 2,
        props: {
          activeValue: 1,
          inactiveValue: 0
        }
      }
    : {}
);

const columns: ColumnProps<SystemMailLog.ResSystemMailLogItem>[] = [
  { prop: "id", label: "编号", width: 100, fixed: "left" },
  { prop: "userType", label: "会员类型", tag: true, enum: userTypeEnum, search: { el: "select", span: 2 } },
  { prop: "userId", label: "用户编号", search: { el: "input", span: 2, props: { placeholder: "请输入用户编号" } } },
  {
    prop: "accountId",
    label: "邮箱账号",
    tag: true,
    enum: mailAccount,
    fieldNames: { label: "mail", value: "id" },
    search: { el: "select", span: 2 }
  },
  {
    prop: "templateCode",
    label: "模板编码",
    tag: true,
    enum: mailTemplate,
    fieldNames: { label: "name", value: "code" },
    search: { el: "select", span: 2 }
  },
  { prop: "templateTitle", label: "邮件标题", search: { el: "input", span: 2, props: { placeholder: "请输入标题" } } },
  { prop: "templateNickname", label: "发送人名称" },
  { prop: "templateContent", label: "模板内容" },
  { prop: "templateParams", label: "模板参数" },
  { prop: "sendStatus", label: "发送状态", tag: true, enum: sendStatus, search: { el: "select", span: 2 } },
  {
    prop: "sendTime",
    label: "发送时间",
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
    search: deleteSearch
  },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission("mail.SystemMailLogDelete", "mail.SystemMailLogDrop", "mail.SystemMailLogRecover", "mail.SystemMailLog")
  }
];

// 重置数据
const reset = () => {
  systemMailLogItemFrom.value = {
    id: 0, // 编号
    userId: undefined, // 用户编号
    userType: undefined, // 用户类型
    toMail: "", // 接收邮箱地址
    accountId: 0, // 邮箱账号编号
    fromMail: "", // 发送邮箱地址
    templateId: 0, // 模板编号
    templateCode: "", // 模板编码
    templateNickname: undefined, // 模版发送人名称
    templateTitle: "", // 邮件标题
    templateContent: "", // 邮件内容
    templateParams: undefined, // 邮件参数
    sendStatus: 0, // 发送状态
    sendTime: undefined, // 发送时间
    sendMessageId: undefined, // 发送返回的消息 ID
    sendException: undefined, // 发送异常
    creator: undefined, // 创建者
    createTime: "", // 创建时间
    updater: undefined, // 更新者
    updateTime: "", // 更新时间
    deleted: 0 // 是否删除
  };
};
// 清理按钮
const handleDrop = async (row: SystemMailLog.ResSystemMailLogItem) => {
  await useHandleData(dropSystemMailLogApi, Number(row.id), "清理邮件日志");
  proTable.value?.getTableList();
};
// 删除按钮
const handleDelete = async (row: SystemMailLog.ResSystemMailLogItem) => {
  await useHandleData(deleteSystemMailLogApi, Number(row.id), "删除邮件日志");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: SystemMailLog.ResSystemMailLogItem) => {
  await useHandleData(recoverSystemMailLogApi, Number(row.id), "恢复邮件日志");
  proTable.value?.getTableList();
};
// 查看按钮
const handleItem = async (row: SystemMailLog.ResSystemMailLogItem) => {
  title.value = "查看邮件日志";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemMailLogApi(Number(row.id));
  systemMailLogItemFrom.value = data;
};

const getCustomSystemMailLogListApi = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.sendTime && (newParams.beginSendTime = newParams.sendTime[0]);
  newParams.sendTime && (newParams.finishSendTime = newParams.sendTime[1]);
  delete newParams.sendTime;
  return getSystemMailLogListApi(newParams);
};

const getMailAccount = async () => {
  const { data } = await getSystemMailAccountListSimpleApi({ deleted: 1 });
  mailAccount.value = data;
};

const getMailTemplate = async () => {
  const { data } = await getSystemMailTemplateListSimpleApi({ deleted: 1 });
  mailTemplate.value = data;
};

// 批量删除
const handleMultipleDelete = async (ids: string[]) => {
  const data = ref<SystemMailLog.ReqSystemMailLogMultiple>({});
  if (ids.length > 0) {
    data.value.ids = ids.map(Number);
  }
  await useHandleData(deleteSystemMailLogMultipleApi, data.value, "删除邮件日志");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

// 批量恢复
const handleMultipleRecover = async (ids: string[]) => {
  const data = ref<SystemMailLog.ReqSystemMailLogMultiple>({});
  if (ids.length > 0) {
    data.value.ids = ids.map(Number);
  }
  await useHandleData(recoverSystemMailLogMultipleApi, data.value, "恢复邮件日志");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

// 批量清理
const handleMultipleDrop = async (ids: string[]) => {
  const data = ref<SystemMailLog.ReqSystemMailLogMultiple>({});
  if (ids.length > 0) {
    data.value.ids = ids.map(Number);
  }
  await useHandleData(dropSystemMailLogMultipleApi, data.value, "清理邮件日志");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

onMounted(() => {
  getMailAccount();
  getMailTemplate();
});
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
