<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="消息列表"
      row-key="id"
      :columns="columns"
      :request-api="getCustomSystemNotifyMessageListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <template #tableHeader="scope">
        <el-dropdown trigger="click" class="table-el-dropdown">
          <el-button
            type="danger"
            v-auth="[
              'notify.SystemNotifyMessageMultipleDelete',
              'notify.SystemNotifyMessageMultipleRecover',
              'notify.SystemNotifyMessageMultipleDrop'
            ]"
            :icon="ArrowDownBold"
            plain
            :disabled="!scope.isSelected">
            操作
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item
                v-auth="'notify.SystemNotifyMessageMultipleDelete'"
                :icon="Delete"
                @click="handleMultipleDelete(scope.selectedListIds)">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                v-auth="'notify.SystemNotifyMessageMultipleDrop'"
                :icon="DeleteFilled"
                @click="handleMultipleDrop(scope.selectedListIds)">
                清理
              </el-dropdown-item>
              <el-dropdown-item
                v-auth="'notify.SystemNotifyMessageMultipleRecover'"
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
        <el-button v-auth="'notify.SystemNotifyMessage'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['notify.SystemNotifyMessageDelete', 'notify.SystemNotifyMessageRecover', 'notify.SystemNotifyMessageDrop']"
            type="primary"
            link
            :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item
                v-if="scope.row.deleted === 0"
                v-auth="'notify.SystemNotifyMessageDelete'"
                :icon="Delete"
                @click="handleDelete(scope.row)">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'notify.SystemNotifyMessageRecover'"
                :icon="Refresh"
                @click="handleRecover(scope.row)">
                恢复
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'notify.SystemNotifyMessageDrop'"
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
          {{ systemNotifyMessageItemFrom.id }}
        </el-descriptions-item>
        <el-descriptions-item label="会员类型">
          <DictTag type="notifyMessage.userType" :value="systemNotifyMessageItemFrom.userType" />
        </el-descriptions-item>
        <el-descriptions-item label="用户编号">
          {{ systemNotifyMessageItemFrom.userId }}
        </el-descriptions-item>
        <el-descriptions-item label="模板编码">
          {{ systemNotifyMessageItemFrom.templateCode }}
        </el-descriptions-item>
        <el-descriptions-item label="发送人名称">
          {{ systemNotifyMessageItemFrom.templateNickname }}
        </el-descriptions-item>
        <el-descriptions-item label="模板内容">
          {{ systemNotifyMessageItemFrom.templateContent }}
        </el-descriptions-item>
        <el-descriptions-item label="模板参数">
          {{ systemNotifyMessageItemFrom.templateParams }}
        </el-descriptions-item>
        <el-descriptions-item label="模板类型">
          <DictTag type="notifyTemplate.type" :value="systemNotifyMessageItemFrom.templateType" />
        </el-descriptions-item>
        <el-descriptions-item label="阅读状态">
          <DictTag type="notifyMessage.readStatus" :value="systemNotifyMessageItemFrom.readStatus" />
        </el-descriptions-item>
        <el-descriptions-item label="阅读时间">
          {{ systemNotifyMessageItemFrom.readTime }}
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemNotifyMessage">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { Delete, Refresh, DeleteFilled, View, DArrowRight, ArrowDownBold } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemNotifyMessage } from "@/api/interface/systemNotifyMessage";
import {
  getSystemNotifyMessageListApi,
  deleteSystemNotifyMessageApi,
  dropSystemNotifyMessageApi,
  recoverSystemNotifyMessageApi,
  getSystemNotifyMessageApi,
  deleteSystemNotifyMessageMultipleApi,
  dropSystemNotifyMessageMultipleApi,
  recoverSystemNotifyMessageMultipleApi
} from "@/api/modules/systemNotifyMessage";
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
const systemNotifyMessageItemFrom = ref<SystemNotifyMessage.ResSystemNotifyMessageItem>({
  id: 0, // 消息
  userId: 0, // 用户id
  userType: 0, // 用户类型
  templateId: 0, // 模版编号
  templateCode: "", // 模板编码
  templateNickname: "", // 模版发送人名称
  templateContent: "", // 模版内容
  templateType: 0, // 模版类型
  templateParams: {}, // 模版参数
  readStatus: 0, // 是否已读
  readTime: undefined, // 阅读时间
  deleted: 0, // 删除
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined // 更新时间
});
//删除状态
const deletedEnum = getIntDictOptions("delete");
//消息类型
const typeEnum = getIntDictOptions("notifyTemplate.type");
//会员类型
const userTypeEnum = getIntDictOptions("notifyMessage.userType");
// 阅读状态
const readStatusEnum = getIntDictOptions("notifyMessage.readStatus");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("notify.SystemNotifyMessageDelete")
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

const columns: ColumnProps<SystemNotifyMessage.ResSystemNotifyMessageItem>[] = [
  { prop: "id", label: "编号", width: 100, fixed: "left" },
  { prop: "userType", label: "会员类型", tag: true, enum: userTypeEnum, search: { el: "select", span: 2 } },
  { prop: "userId", label: "用户编号" },
  { prop: "templateCode", label: "模板编码" },
  { prop: "templateNickname", label: "发送人名称" },
  { prop: "templateContent", label: "模板内容" },
  { prop: "templateParams", label: "模板参数" },
  { prop: "templateType", label: "模板类型", tag: true, enum: typeEnum, search: { el: "select", span: 2 } },
  { prop: "readStatus", label: "阅读状态", tag: true, enum: readStatusEnum, search: { el: "select", span: 2 } },
  {
    prop: "readTime",
    label: "阅读时间",
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
  { prop: "creator", label: "创建人" },
  { prop: "createTime", label: "创建时间" },
  { prop: "updater", label: "更新人" },
  { prop: "updateTime", label: "更新时间" },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission(
      "notify.SystemNotifyMessageDelete",
      "notify.SystemNotifyMessageDrop",
      "notify.SystemNotifyMessageRecover",
      "notify.SystemNotifyMessage"
    )
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  systemNotifyMessageItemFrom.value = {
    id: 0, // 消息
    userId: 0, // 用户id
    userType: 0, // 用户类型
    templateId: 0, // 模版编号
    templateCode: "", // 模板编码
    templateNickname: "", // 模版发送人名称
    templateContent: "", // 模版内容
    templateType: 0, // 模版类型
    templateParams: {}, // 模版参数
    readStatus: 0, // 是否已读
    readTime: undefined, // 阅读时间
    deleted: 0, // 删除
    tenantId: 0, // 租户
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined // 更新时间
  };
};

// 清理按钮
const handleDrop = async (row: SystemNotifyMessage.ResSystemNotifyMessageItem) => {
  await useHandleData(dropSystemNotifyMessageApi, Number(row.id), "清理消息");
  proTable.value?.getTableList();
};
// 删除按钮
const handleDelete = async (row: SystemNotifyMessage.ResSystemNotifyMessageItem) => {
  await useHandleData(deleteSystemNotifyMessageApi, Number(row.id), "删除消息");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: SystemNotifyMessage.ResSystemNotifyMessageItem) => {
  await useHandleData(recoverSystemNotifyMessageApi, Number(row.id), "恢复消息");
  proTable.value?.getTableList();
};
// 查看按钮
const handleItem = async (row: SystemNotifyMessage.ResSystemNotifyMessageItem) => {
  title.value = "查看消息";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemNotifyMessageApi(Number(row.id));
  systemNotifyMessageItemFrom.value = data;
};

const getCustomSystemNotifyMessageListApi = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.readTime && (newParams.beginReadTime = newParams.readTime[0]);
  newParams.readTime && (newParams.finishReadTime = newParams.readTime[1]);
  delete newParams.readTime;
  return getSystemNotifyMessageListApi(newParams);
};

// 批量删除
const handleMultipleDelete = async (ids: string[]) => {
  const data = ref<SystemNotifyMessage.ReqSystemNotifyMessageMultiple>({});
  if (ids.length > 0) {
    data.value.ids = ids.map(Number);
  }
  await useHandleData(deleteSystemNotifyMessageMultipleApi, data.value, "删除消息");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

// 批量恢复
const handleMultipleRecover = async (ids: string[]) => {
  const data = ref<SystemNotifyMessage.ReqSystemNotifyMessageMultiple>({});
  if (ids.length > 0) {
    data.value.ids = ids.map(Number);
  }
  await useHandleData(recoverSystemNotifyMessageMultipleApi, data.value, "恢复消息");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

// 批量清理
const handleMultipleDrop = async (ids: string[]) => {
  const data = ref<SystemNotifyMessage.ReqSystemNotifyMessageMultiple>({});
  if (ids.length > 0) {
    data.value.ids = ids.map(Number);
  }
  await useHandleData(dropSystemNotifyMessageMultipleApi, data.value, "清理消息");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
