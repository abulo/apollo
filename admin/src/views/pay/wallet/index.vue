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
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'wallet.PayWalletCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
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
      <el-form ref="refPayWalletItemFrom" :model="payWalletItemFrom" :rules="rulesPayWalletItemFrom" label-width="100px">
        <el-form-item label="编号" prop="id">
          <el-input v-model="payWalletItemFrom.id" />
        </el-form-item>
        <el-form-item label="用户编号" prop="userId">
          <el-input v-model="payWalletItemFrom.userId" />
        </el-form-item>
        <el-form-item label="用户类型" prop="userType">
          <el-input v-model="payWalletItemFrom.userType" />
        </el-form-item>
        <el-form-item label="余额，单位分" prop="balance">
          <el-input v-model="payWalletItemFrom.balance" />
        </el-form-item>
        <el-form-item label="累计支出，单位分" prop="totalExpense">
          <el-input v-model="payWalletItemFrom.totalExpense" />
        </el-form-item>
        <el-form-item label="累计充值，单位分" prop="totalRecharge">
          <el-input v-model="payWalletItemFrom.totalRecharge" />
        </el-form-item>
        <el-form-item label="冻结金额，单位分" prop="freezePrice">
          <el-input v-model="payWalletItemFrom.freezePrice" />
        </el-form-item>
        <el-form-item label="删除" prop="deleted">
          <el-input v-model="payWalletItemFrom.deleted" />
        </el-form-item>
        <el-form-item label="租户" prop="tenantId">
          <el-input v-model="payWalletItemFrom.tenantId" />
        </el-form-item>
        <el-form-item label="创建人" prop="creator">
          <el-input v-model="payWalletItemFrom.creator" />
        </el-form-item>
        <el-form-item label="创建时间" prop="createTime">
          <el-input v-model="payWalletItemFrom.createTime" />
        </el-form-item>
        <el-form-item label="更新人" prop="updater">
          <el-input v-model="payWalletItemFrom.updater" />
        </el-form-item>
        <el-form-item label="更新时间" prop="updateTime">
          <el-input v-model="payWalletItemFrom.updateTime" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refPayWalletItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refPayWalletItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
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
//加载
const loading = ref(false);
//弹出层标题
const title = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
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
  updateTime: undefined // 更新时间
});
//校验
const refPayWalletItemFrom = ref<FormInstance>();
//校验
const rulesPayWalletItemFrom = reactive<FormRules>({
  id: [{ required: true, message: "编号不能为空", trigger: "blur" }],
  userId: [{ required: true, message: "用户编号不能为空", trigger: "blur" }],
  userType: [{ required: true, message: "用户类型不能为空", trigger: "blur" }],
  balance: [{ required: true, message: "余额，单位分不能为空", trigger: "blur" }],
  totalExpense: [{ required: true, message: "累计支出，单位分不能为空", trigger: "blur" }],
  totalRecharge: [{ required: true, message: "累计充值，单位分不能为空", trigger: "blur" }],
  freezePrice: [{ required: true, message: "冻结金额，单位分不能为空", trigger: "blur" }],
  deleted: [{ required: true, message: "删除不能为空", trigger: "blur" }],
  tenantId: [{ required: true, message: "租户不能为空", trigger: "blur" }]
});
//删除状态
const deletedEnum = getIntDictOptions("delete");
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
  { prop: "userId", label: "用户编号", search: { el: "input", span: 2, props: { placeholder: "请输入用户编号" } } },
  { prop: "userType", label: "用户类型", search: { el: "input", span: 2, props: { placeholder: "请输入用户类型" } } },
  { prop: "balance", label: "余额，单位分" },
  { prop: "totalExpense", label: "累计支出，单位分" },
  { prop: "totalRecharge", label: "累计充值，单位分" },
  { prop: "freezePrice", label: "冻结金额，单位分" },
  { prop: "deleted", label: "删除", search: { el: "input", span: 2, props: { placeholder: "请输入删除" } } },
  { prop: "tenantId", label: "租户" },
  { prop: "creator", label: "创建人" },
  { prop: "createTime", label: "创建时间" },
  { prop: "updater", label: "更新人" },
  { prop: "updateTime", label: "更新时间" },

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
    balance: 0, // 余额，单位分
    totalExpense: 0, // 累计支出，单位分
    totalRecharge: 0, // 累计充值，单位分
    freezePrice: 0, // 冻结金额，单位分
    deleted: 0, // 删除
    tenantId: 0, // 租户
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined // 更新时间
  };
};

// resetForm
const resetForm = (formEl: FormInstance | undefined) => {
  centerDialogVisible.value = false;
  if (!formEl) return;
  formEl.resetFields();
};

// 提交数据
const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loading.value = true;
    const data = payWalletItemFrom.value as unknown as PayWallet.ResPayWalletItem;
    if (data.id !== 0) {
      await useHandleSet(updatePayWalletApi, data.id, data, "修改会员钱包表");
    } else {
      await useHandleData(addPayWalletApi, data, "添加会员钱包表");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
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
// 添加按钮
const handleAdd = () => {
  title.value = "新增会员钱包表";
  centerDialogVisible.value = true;
  reset();
};

// 编辑按钮
const handleUpdate = async (row: PayWallet.ResPayWalletItem) => {
  title.value = "编辑会员钱包表";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getPayWalletItemApi(Number(row.id));
  payWalletItemFrom.value = data;
};
</script>
<style lang="scss">
@import "@/styles/custom.scss";
</style>
