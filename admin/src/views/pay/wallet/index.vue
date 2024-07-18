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
      <template #username="scope">
        <p class="order-font" v-if="scope.row.username"><el-tag size="small"> 用户名</el-tag> {{ scope.row.username }}</p>
        <p class="order-font" v-if="scope.row.nickname">
          <el-tag size="small" type="warning">昵称</el-tag> {{ scope.row.nickname }}
        </p>
        <p class="order-font" v-if="scope.row.mobile"><el-tag size="small" type="success">手机</el-tag> {{ scope.row.mobile }}</p>
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
        <el-dropdown trigger="click">
          <el-button
            v-auth="['wallet.PayWalletRechargeList', 'wallet.PayWalletTransactionList']"
            type="primary"
            link
            :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'wallet.PayWalletRechargeList'" :icon="DataBoard" @click="toRouteRecharge(scope.row)">
                钱包充值
              </el-dropdown-item>
              <el-dropdown-item
                v-auth="'wallet.PayWalletTransactionList'"
                :icon="DataBoard"
                @click="toRouteTransaction(scope.row)">
                钱包流水
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </template>
    </ProTable>
  </div>
</template>
<script setup lang="ts" name="payWallet">
import { ref, reactive } from "vue";
import { useRouter } from "vue-router";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { Delete, Refresh, DataBoard, DArrowRight } from "@element-plus/icons-vue";
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
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
const router = useRouter();
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
  username: "", // 用户
  nickname: "", // 用户名称
  mobile: "" // 手机号
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
  { prop: "id", label: "编号", fixed: "left" },
  { prop: "username", label: "用户名称", search: { el: "input", span: 2, props: { placeholder: "昵称/用户名/手机号" } } },
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
    isShow: HasPermission(
      "wallet.PayWalletRechargeList",
      "wallet.PayWalletTransactionList",
      "wallet.PayWalletDelete",
      "wallet.PayWalletRecover"
    )
  }
];

// 重置数据
const reset = () => {
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
    username: "", // 用户
    nickname: "", // 用户名称
    mobile: "" // 手机号
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

const toRouteRecharge = (row: PayWallet.ResPayWalletItem) => {
  router.push({
    name: "payWalletRecharge",
    params: {
      walletId: String(row.id)
    }
  });
};
const toRouteTransaction = (row: PayWallet.ResPayWalletItem) => {
  router.push({
    name: "payWalletTransaction",
    params: {
      walletId: String(row.id)
    }
  });
};
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
