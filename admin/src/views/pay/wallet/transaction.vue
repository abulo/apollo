<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="会员钱包流水表列表"
      row-key="id"
      :columns="columns"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader><el-button type="primary" :icon="Remove" @click="closeCurrentTab">关闭</el-button> </template>
      <template #price="scope"> ￥{{ parseFloat(String(scope.row.price / 100)).toFixed(2) }} </template>
      <template #balance="scope"> ￥{{ parseFloat(String(scope.row.balance / 100)).toFixed(2) }} </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'wallet.PayWalletTransactionDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'wallet.PayWalletTransactionRecover'"
          type="primary"
          link
          :icon="Refresh"
          @click="handleRecover(scope.row)">
          恢复
        </el-button>
      </template>
    </ProTable>
  </div>
</template>
<script setup lang="ts" name="payWalletTransaction">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { Delete, Refresh, Remove } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { PayWalletTransaction } from "@/api/interface/payWalletTransaction";
import {
  getPayWalletTransactionListApi,
  deletePayWalletTransactionApi,
  recoverPayWalletTransactionApi,
  getPayWalletTransactionItemApi
} from "@/api/modules/payWalletTransaction";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import { useRoute } from "vue-router";
import { useTabsStore } from "@/stores/modules/tabs";
import { useKeepAliveStore } from "@/stores/modules/keepAlive";
const route = useRoute();
const tabStore = useTabsStore();
const keepAliveStore = useKeepAliveStore();
// 钱包 id
const walletId = ref(Number(route.params.walletId));
//table数据
const proTable = ref<ProTableInstance>();
//删除状态
const deletedEnum = getIntDictOptions("delete");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("wallet.PayWalletTransactionDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);

const columns: ColumnProps<PayWalletTransaction.ResPayWalletTransactionItem>[] = [
  { prop: "id", label: "编号" },
  { prop: "bizType", label: "关联类型", search: { el: "input", span: 2, props: { placeholder: "请输入关联类型" } } },
  { prop: "bizId", label: "关联业务编号", search: { el: "input", span: 2, props: { placeholder: "请输入关联业务编号" } } },
  { prop: "no", label: "流水号", search: { el: "input", span: 2, props: { placeholder: "请输入流水号" } } },
  { prop: "title", label: "流水标题", search: { el: "input", span: 2, props: { placeholder: "请输入流水标题" } } },
  { prop: "price", label: "交易金额" },
  { prop: "balance", label: "余额" },
  { prop: "deleted", label: "删除", tag: true, enum: deletedEnum, search: deleteSearch },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission("wallet.PayWalletTransactionDelete", "wallet.PayWalletTransactionRecover")
  }
];

// 删除按钮
const handleDelete = async (row: PayWalletTransaction.ResPayWalletTransactionItem) => {
  await useHandleData(deletePayWalletTransactionApiCustom, Number(row.id), "删除会员钱包流水表");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: PayWalletTransaction.ResPayWalletTransactionItem) => {
  await useHandleData(recoverPayWalletTransactionApiCustom, Number(row.id), "恢复会员钱包流水表");
  proTable.value?.getTableList();
};

// 关闭当前页
const closeCurrentTab = () => {
  if (route.meta.isAffix) return;
  tabStore.removeTabs(route.fullPath);
  keepAliveStore.removeKeepAliveName(route.name as string);
};

const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  return getPayWalletTransactionListApi(walletId.value, newParams);
};
const recoverPayWalletTransactionApiCustom = (id: number) => {
  return recoverPayWalletTransactionApi(walletId.value, id);
};
const deletePayWalletTransactionApiCustom = (id: number) => {
  return deletePayWalletTransactionApi(walletId.value, id);
};
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
