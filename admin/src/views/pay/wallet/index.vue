<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="会员钱包表列表"
      row-key="id"
      :columns="columns"
      :request-api="getPayWalletListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 用户类型 -->
      <template #userType="scope">
        <DictTag type="pay.walletUserType" :value="scope.row.userType" />
      </template>
      <!-- 余额 -->
      <template #balance="scope">
        <span>￥{{ parseFloat(String(scope.row.balance / 100)).toFixed(2) }}</span>
      </template>
      <!-- 累计支出 -->
      <template #totalExpense="scope">
        <span>￥{{ parseFloat(String(scope.row.totalExpense / 100)).toFixed(2) }}</span>
      </template>
      <!-- 累计充值 -->
      <template #totalRecharge="scope">
        <span>￥{{ parseFloat(String(scope.row.totalRecharge / 100)).toFixed(2) }}</span>
      </template>
      <!-- 冻结金额 -->
      <template #freezePrice="scope">
        <span>￥{{ parseFloat(String(scope.row.freezePrice / 100)).toFixed(2) }}</span>
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'wallet.PayWalletUpdate'" type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)">
          编辑
        </el-button>
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'wallet.PayWalletDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'wallet.PayWalletRecover'"
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
<script setup lang="ts" name="payWallet">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { PayWallet } from "@/api/interface/payWallet";
import {
  getPayWalletListApi,
  deletePayWalletApi,
  recoverPayWalletApi,
  getPayWalletItemApi,
  addPayWalletApi,
  updatePayWalletApi
} from "@/api/modules/payWallet";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
//table数据
const proTable = ref<ProTableInstance>();
//数据接口
const payWalletItemFrom = ref<PayWallet.ResPayWalletItem>({
  id: 0, // 编号
  userId: 0, // 用户编号
  userType: 0, // 用户类型
  balance: 0, // 余额，单位分
  totalExpense: 0, // 累计支出，单位分
  totalRecharge: 0, // 累计充值，单位分
  freezePrice: 0, // 冻结金额，单位分
  deleted: 0, // 删除
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
  username: "" // 用户昵称
});
//删除状态
const deletedEnum = getIntDictOptions("delete");
const userType = getIntDictOptions("pay.walletUserType");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("wallet.PayWalletDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);

const columns: ColumnProps<PayWallet.ResPayWalletItem>[] = [
  { prop: "id", label: "编号" },
  { prop: "username", label: "用户昵称", search: { el: "input", span: 3, props: { placeholder: "昵称/用户名/手机号" } } },
  { prop: "userType", label: "用户类型", tag: true, enum: userType, search: { el: "select", span: 2 } },
  { prop: "balance", label: "余额" },
  { prop: "totalExpense", label: "累计支出" },
  { prop: "totalRecharge", label: "累计充值" },
  { prop: "freezePrice", label: "冻结金额" },
  { prop: "deleted", label: "删除", tag: true, enum: deletedEnum, search: deleteSearch },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission("wallet.PayWalletUpdate", "wallet.PayWalletDelete", "wallet.PayWalletRecover")
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  payWalletItemFrom.value = {
    id: 0, // 编号
    userId: 0, // 用户编号
    userType: 0, // 用户类型
    balance: 0, // 余额
    totalExpense: 0, // 累计支出
    totalRecharge: 0, // 累计充值
    freezePrice: 0, // 冻结金额
    deleted: 0, // 删除
    tenantId: 0, // 租户
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined, // 更新时间
    username: "" // 用户昵称
  };
};
// 删除按钮
const handleDelete = async (row: PayWallet.ResPayWalletItem) => {
  await useHandleData(deletePayWalletApi, Number(row.id), "删除会员钱包表");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: PayWallet.ResPayWalletItem) => {
  await useHandleData(recoverPayWalletApi, Number(row.id), "恢复会员钱包表");
  proTable.value?.getTableList();
};
</script>
<style lang="scss">
@import "@/styles/custom.scss";
</style>
