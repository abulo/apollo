<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="消息列表"
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
        <el-button v-auth="'notify.SystemNotifyTemplate'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          详情
        </el-button>
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'notify.SystemNotifyMessageDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'notify.SystemNotifyMessageRecover'"
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
import { View, CirclePlus, Delete, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemNotifyMessage } from "@/api/interface/systemNotifyMessage";
import {
  getSystemNotifyMessageListApi,
  deleteSystemNotifyMessageApi,
  recoverSystemNotifyMessageApi,
  getSystemNotifyMessageItemApi
} from "@/api/modules/systemNotifyMessage";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
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
const systemNotifyMessageItemFrom = ref<SystemNotifyMessage.ResSystemNotifyMessageItem>({
  id: 0, // id
  userId: 0, // 用户 Id
  userType: 0, // 用户类型
  templateId: 0, // 消息编号
  templateCode: "", // 消息编码
  templateNickname: "", // 发送人名称
  templateContent: "", // 模版内容
  templateType: 0, // 模版类型
  templateParams: {}, // 模版参数
  readStatus: 0, // 是否已读
  readTime: "", // 阅读时间
  deleted: 0 // 删除
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
        span: 2
      }
    : {}
);
const columns: ColumnProps<SystemNotifyMessage.ResSystemNotifyMessageItem>[] = [
  { prop: "id", label: "编号", width: 100 },
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
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission("notify.SystemNotifyMessageDelete", "notify.SystemNotifyMessageRecover", "notify.SystemNotifyTemplate")
  }
];
// 重置数据
const reset = () => {
  systemNotifyMessageItemFrom.value = {
    id: 0, // id
    userId: 0, // 用户 Id
    userType: 0, // 用户类型
    templateId: 0, // 消息编号
    templateCode: "", // 消息编码
    templateNickname: "", // 发送人名称
    templateContent: "", // 模版内容
    templateType: 0, // 模版类型
    templateParams: {}, // 模版参数
    readStatus: 0, // 是否已读
    readTime: "", // 阅读时间
    deleted: 0 // 删除
  };
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

const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.readTime && (newParams.beginReadTime = newParams.readTime[0]);
  newParams.readTime && (newParams.finishReadTime = newParams.readTime[1]);
  delete newParams.readTime;
  return getSystemNotifyMessageListApi(newParams);
};

// 编辑按钮
const handleItem = async (row: SystemNotifyMessage.ResSystemNotifyMessageItem) => {
  title.value = "查看信息";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemNotifyMessageItemApi(row.id);
  systemNotifyMessageItemFrom.value = data;
};
</script>

<style lang="scss">
@import "@/styles/custom";
</style>
