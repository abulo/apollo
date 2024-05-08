<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="记录列表"
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
        <el-button v-auth="'mail.SystemNotifyTemplate'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          详情
        </el-button>
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'mail.SystemMailLogDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'mail.SystemMailLogRecover'"
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
      <el-descriptions :column="1" border>
        <el-descriptions-item label="编号" min-width="120">
          {{ systemMailLogItemFrom.id }}
        </el-descriptions-item>
        <el-descriptions-item label="用户编号" min-width="120">
          {{ systemMailLogItemFrom.userId }}
        </el-descriptions-item>
        <el-descriptions-item label="用户类型" min-width="120">
          <DictTag type="notifyMessage.userType" :value="systemMailLogItemFrom.userType" />
        </el-descriptions-item>
        <el-descriptions-item label="接收邮箱地址" min-width="120">
          {{ systemMailLogItemFrom.toMail }}
        </el-descriptions-item>
        <el-descriptions-item label="邮箱账号" min-width="120">
          {{ systemMailLogItemFrom.accountId }}
        </el-descriptions-item>
        <el-descriptions-item label="发送邮箱地址" min-width="120">
          {{ systemMailLogItemFrom.fromMail }}
        </el-descriptions-item>
        <el-descriptions-item label="模板编码" min-width="120">
          {{ systemMailLogItemFrom.templateCode }}
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
import { View, CirclePlus, Delete, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemMailLog } from "@/api/interface/systemMailLog";
import {
  getSystemMailLogListApi,
  deleteSystemMailLogApi,
  recoverSystemMailLogApi,
  getSystemMailLogItemApi
} from "@/api/modules/systemMailLog";
import { FormInstance, FormRules } from "element-plus";
import { getSystemMailAccountSearchApi } from "@/api/modules/systemMailAccount";
import { getIntDictOptions, DictDataType } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
//table数据
const proTable = ref<ProTableInstance>();
//弹出层标题
const title = ref();
//是否显示弹出层
const centerDialogVisible = ref(false);
//数据接口
const systemMailLogItemFrom = ref<SystemMailLog.ResSystemMailLogItem>({
  id: 0, // 主键
  userId: 0, // 用户编号
  userType: 0, // 用户类型
  toMail: "", // 接收邮箱地址
  accountId: 0, // 邮箱账号编号
  fromMail: "", // 发送邮箱地址
  templateId: 0, // 模板编号
  templateCode: "", // 模板编码
  templateNickname: "", // 发送人名称
  templateTitle: "", // 邮件标题
  templateContent: "", // 邮件内容
  templateParams: {}, // 邮件参数
  sendStatus: 0, // 发送状态
  sendTime: "", // 发送时间
  sendMessageId: "", // 发送返回的消息 ID
  sendException: "", // 发送异常
  deleted: 0 // 是否删除
});

//删除状态
const deletedEnum = getIntDictOptions("delete");
//会员类型
const userTypeEnum = getIntDictOptions("notifyMessage.userType");
// 阅读状态
const sendStatus = getIntDictOptions("mailLog.sendStatus");
const accountEnum = reactive<DictDataType[]>([]);
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("mail.SystemMailLogDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);

const columns: ColumnProps<SystemMailLog.ResSystemMailLogItem>[] = [
  { prop: "id", label: "编号", width: 100 },
  { prop: "userType", label: "会员类型", tag: true, enum: userTypeEnum, search: { el: "select", span: 2 } },
  { prop: "userId", label: "用户编号", search: { el: "input", span: 2, props: { placeholder: "请输入用户编号" } } },
  { prop: "accountId", label: "邮箱账号", tag: true, enum: accountEnum, search: { el: "select", span: 2 } },
  { prop: "templateCode", label: "模板编码", search: { el: "input", span: 2, props: { placeholder: "请输入编码" } } },
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
    isShow: HasPermission("mail.SystemMailLogDelete", "mail.SystemMailLogRecover", "mail.SystemNotifyTemplate")
  }
];
// 重置数据
const reset = () => {
  systemMailLogItemFrom.value = {
    id: 0, // 主键
    userId: 0, // 用户编号
    userType: 0, // 用户类型
    toMail: "", // 接收邮箱地址
    accountId: 0, // 邮箱账号编号
    fromMail: "", // 发送邮箱地址
    templateId: 0, // 模板编号
    templateCode: "", // 模板编码
    templateNickname: "", // 模版发送人名称
    templateTitle: "", // 邮件标题
    templateContent: "", // 邮件内容
    templateParams: {}, // 邮件参数
    sendStatus: 0, // 发送状态
    sendTime: "", // 发送时间
    sendMessageId: "", // 发送返回的消息 ID
    sendException: "", // 发送异常
    deleted: 0 // 是否删除
  };
};

// 删除按钮
const handleDelete = async (row: SystemMailLog.ResSystemMailLogItem) => {
  await useHandleData(deleteSystemMailLogApi, Number(row.id), "删除记录");
  proTable.value?.getTableList();
};

// 恢复按钮
const handleRecover = async (row: SystemMailLog.ResSystemMailLogItem) => {
  await useHandleData(recoverSystemMailLogApi, Number(row.id), "恢复记录");
  proTable.value?.getTableList();
};

const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.sendTime && (newParams.beginSendTime = newParams.sendTime[0]);
  newParams.sendTime && (newParams.finishSendTime = newParams.sendTime[1]);
  delete newParams.sendTime;
  return getSystemMailLogListApi(newParams);
};

// 编辑按钮
const handleItem = async (row: SystemMailLog.ResSystemMailLogItem) => {
  title.value = "查看信息";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemMailLogItemApi(row.id);
  systemMailLogItemFrom.value = data;
};

const getAccount = async () => {
  const { data } = await getSystemMailAccountSearchApi();
  //扁平化输出数据
  data.forEach(item => {
    let obj: DictDataType = {
      label: item.mail,
      value: Number(item.id),
      colorType: "",
      cssClass: "",
      dictType: "account"
    };
    accountEnum.push(obj);
  });
};

onMounted(() => {
  getAccount();
});
</script>

<style lang="scss">
@import "@/styles/custom.scss";
</style>
