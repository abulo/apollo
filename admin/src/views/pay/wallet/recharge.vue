<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="会员钱包充值列表"
      row-key="id"
      :columns="columns"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button type="primary" :icon="Remove" @click="closeCurrentTab">关闭</el-button>
      </template>
      <template #totalPrice="scope"> ￥{{ parseFloat(String(scope.row.totalPrice / 100)).toFixed(2) }} </template>
      <template #payPrice="scope"> ￥{{ parseFloat(String(scope.row.payPrice / 100)).toFixed(2) }} </template>
      <template #bonusPrice="scope"> ￥{{ parseFloat(String(scope.row.bonusPrice / 100)).toFixed(2) }} </template>
      <template #refundTotalPrice="scope"> ￥{{ parseFloat(String(scope.row.refundTotalPrice / 100)).toFixed(2) }} </template>
      <template #refundPayPrice="scope"> ￥{{ parseFloat(String(scope.row.refundPayPrice / 100)).toFixed(2) }} </template>
      <template #refundBonusPrice="scope"> ￥{{ parseFloat(String(scope.row.refundBonusPrice / 100)).toFixed(2) }} </template>
      <template #payChannelCode="scope">
        <DictTag type="payChannel.code" :value="scope.row.payChannelCode" />
      </template>
      <template #payStatus="scope">
        <DictTag type="pay.orderStatus" :value="scope.row.payStatus" />
      </template>
      <template #refundStatus="scope">
        <DictTag type="pay.refundStatus" :value="scope.row.refundStatus" />
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'wallet.PayWalletRechargeDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'wallet.PayWalletRechargeRecover'"
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
<script setup lang="ts" name="payWalletRecharge">
import { ref, reactive, onMounted } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { Delete, Refresh, Remove } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { PayWalletRecharge } from "@/api/interface/payWalletRecharge";
import {
  getPayWalletRechargeListApi,
  deletePayWalletRechargeApi,
  recoverPayWalletRechargeApi
} from "@/api/modules/payWalletRecharge";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions, getStrDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import { useRoute } from "vue-router";
import { useTabsStore } from "@/stores/modules/tabs";
import { useKeepAliveStore } from "@/stores/modules/keepAlive";
import { getPayWalletRechargePackageListSimpleApi } from "@/api/modules/payWalletRechargePackage";
import { PayWalletRechargePackage } from "@/api/interface/payWalletRechargePackage";
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
  HasPermission("wallet.PayWalletRechargeDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);

const payWalletRechargePackageList = ref<PayWalletRechargePackage.ResPayWalletRechargePackageItem[]>([]);
// 支付渠道
const payChannel = getStrDictOptions("payChannel.code");
//支付状态
const payOrderStatusEnum = getIntDictOptions("pay.orderStatus");
//退款状态
const payRefundStatusEnum = getIntDictOptions("pay.refundStatus");

const columns: ColumnProps<PayWalletRecharge.ResPayWalletRechargeItem>[] = [
  { prop: "id", label: "编号" },
  {
    prop: "-",
    label: "支付",
    _children: [
      { prop: "payOrderId", label: "支付订单", search: { el: "input", span: 2, props: { placeholder: "请输入支付订单编号" } } },
      { prop: "totalPrice", label: "到账金额" },
      { prop: "payPrice", label: "支付金额" },
      { prop: "bonusPrice", label: "赠送金额" },
      {
        prop: "packageId",
        label: "充值套餐",
        enum: payWalletRechargePackageList,
        fieldNames: { label: "name", value: "id" },
        search: { el: "select", span: 2 }
      },
      {
        prop: "payStatus",
        label: "支付状态",
        tag: true,
        enum: payOrderStatusEnum,
        search: { el: "select", span: 2 }
      },
      {
        prop: "payChannelCode",
        label: "支付渠道",
        tag: true,
        enum: payChannel,
        search: { el: "select", span: 2 }
      },
      {
        prop: "payTime",
        label: "订单时间",
        search: {
          el: "date-picker",
          span: 4,
          props: { type: "datetimerange", valueFormat: "YYYY-MM-DD HH:mm:ss" }
        }
      }
    ]
  },
  {
    prop: "-",
    label: "退款",
    _children: [
      {
        prop: "payRefundId",
        label: "退款编号",
        search: { el: "input", span: 2, props: { placeholder: "请输入支付退款单编号" } }
      },
      { prop: "refundTotalPrice", label: "退款总金额" },
      { prop: "refundPayPrice", label: "退款支付金额" },
      { prop: "refundBonusPrice", label: "退款赠送金额" },
      {
        prop: "refundTime",
        label: "退款时间",
        search: {
          el: "date-picker",
          span: 4,
          props: { type: "datetimerange", valueFormat: "YYYY-MM-DD HH:mm:ss" }
        }
      },
      { prop: "refundStatus", label: "退款状态", tag: true, enum: payRefundStatusEnum, search: { el: "select", span: 2 } }
    ]
  },
  { prop: "deleted", label: "删除", tag: true, enum: deletedEnum, search: deleteSearch },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission("wallet.PayWalletRechargeDelete", "wallet.PayWalletRechargeRecover")
  }
];

// 删除按钮
const handleDelete = async (row: PayWalletRecharge.ResPayWalletRechargeItem) => {
  await useHandleData(deletePayWalletRechargeApiCustom, Number(row.id), "删除会员钱包充值");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: PayWalletRecharge.ResPayWalletRechargeItem) => {
  await useHandleData(recoverPayWalletRechargeApiCustom, Number(row.id), "恢复会员钱包充值");
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
  newParams.payTime && (newParams.beginPayTime = newParams.payTime[0]);
  newParams.payTime && (newParams.finishPayTime = newParams.payTime[1]);
  newParams.refundTime && (newParams.beginRefundTime = newParams.refundTime[0]);
  newParams.refundTime && (newParams.finishRefundTime = newParams.refundTime[1]);
  delete newParams.payTime;
  delete newParams.refundTime;
  return getPayWalletRechargeListApi(walletId.value, newParams);
};
const recoverPayWalletRechargeApiCustom = (id: number) => {
  return recoverPayWalletRechargeApi(walletId.value, id);
};
const deletePayWalletRechargeApiCustom = (id: number) => {
  return deletePayWalletRechargeApi(walletId.value, id);
};

onMounted(async () => {
  const { data } = await getPayWalletRechargePackageListSimpleApi();
  payWalletRechargePackageList.value = data;
});
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
