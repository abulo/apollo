<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="充值套餐表列表"
      row-key="id"
      :columns="columns"
      :request-api="getPayWalletRechargePackageListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'wallet.PayWalletRechargePackageCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">
          新增
        </el-button>
      </template>
      <template #payPrice="scope">
        <span>￥{{ parseFloat(String(scope.row.payPrice / 100)).toFixed(2) }}</span>
      </template>
      <template #bonusPrice="scope">
        <span>￥{{ parseFloat(String(scope.row.bonusPrice / 100)).toFixed(2) }}</span>
      </template>
      <!-- 状态-->
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button
          v-auth="'wallet.PayWalletRechargePackageUpdate'"
          type="primary"
          link
          :icon="EditPen"
          @click="handleUpdate(scope.row)">
          编辑
        </el-button>
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'wallet.PayWalletRechargePackageDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'wallet.PayWalletRechargePackageRecover'"
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
      <el-form
        ref="refPayWalletRechargePackageItemFrom"
        :model="payWalletRechargePackageItemFrom"
        :rules="rulesPayWalletRechargePackageItemFrom"
        label-width="100px">
        <el-form-item label="套餐名称" prop="name">
          <el-input v-model="payWalletRechargePackageItemFrom.name" />
        </el-form-item>
        <el-form-item label="支付金额" prop="payPrice">
          <el-input-number v-model="payWalletRechargePackageItemFrom.payPrice" :min="0" :precision="2" :step="0.01" />
        </el-form-item>
        <el-form-item label="赠送金额" prop="bonusPrice">
          <el-input-number v-model="payWalletRechargePackageItemFrom.bonusPrice" :min="0" :precision="2" :step="0.01" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="payWalletRechargePackageItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refPayWalletRechargePackageItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refPayWalletRechargePackageItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="payWalletRechargePackage">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { PayWalletRechargePackage } from "@/api/interface/payWalletRechargePackage";
import {
  getPayWalletRechargePackageListApi,
  deletePayWalletRechargePackageApi,
  recoverPayWalletRechargePackageApi,
  getPayWalletRechargePackageItemApi,
  addPayWalletRechargePackageApi,
  updatePayWalletRechargePackageApi
} from "@/api/modules/payWalletRechargePackage";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import { fenToYuan, yuanToFen } from "@/utils/price";
//加载
const loading = ref(false);
//弹出层标题
const title = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
//数据接口
const payWalletRechargePackageItemFrom = ref<PayWalletRechargePackage.ResPayWalletRechargePackageItem>({
  id: 0, // 编号
  name: "", // 套餐名称
  payPrice: 0, // 支付金额
  bonusPrice: 0, // 赠送金额
  status: 0, // 状态
  deleted: 0, // 删除
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined // 更新时间
});
//校验
const refPayWalletRechargePackageItemFrom = ref<FormInstance>();
//校验
const rulesPayWalletRechargePackageItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "套餐名称不能为空", trigger: "blur" }],
  payPrice: [{ required: true, message: "支付金额不能为空", trigger: "blur" }],
  bonusPrice: [{ required: true, message: "赠送金额不能为空", trigger: "blur" }],
  status: [{ required: true, message: "状态不能为空", trigger: "blur" }]
});
//删除状态
const deletedEnum = getIntDictOptions("delete");
const statusEnum = getIntDictOptions("status");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("wallet.PayWalletRechargePackageDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);

const columns: ColumnProps<PayWalletRechargePackage.ResPayWalletRechargePackageItem>[] = [
  { prop: "id", label: "编号" },
  { prop: "name", label: "套餐名称", search: { el: "input", span: 2, props: { placeholder: "请输入套餐名称" } } },
  { prop: "payPrice", label: "支付金额" },
  { prop: "bonusPrice", label: "赠送金额" },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 } },
  { prop: "deleted", label: "删除", tag: true, enum: deletedEnum, search: deleteSearch },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission(
      "wallet.PayWalletRechargePackageUpdate",
      "wallet.PayWalletRechargePackageDelete",
      "wallet.PayWalletRechargePackageRecover"
    )
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  payWalletRechargePackageItemFrom.value = {
    id: 0, // 编号
    name: "", // 套餐名称
    payPrice: 0, // 支付金额
    bonusPrice: 0, // 赠送金额
    status: 0, // 状态
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
    const data = { ...payWalletRechargePackageItemFrom.value };
    data.payPrice = yuanToFen(data.payPrice);
    data.bonusPrice = yuanToFen(data.bonusPrice);
    if (data.id !== 0) {
      await useHandleSet(updatePayWalletRechargePackageApi, data.id, data, "修改充值套餐表");
    } else {
      await useHandleData(addPayWalletRechargePackageApi, data, "添加充值套餐表");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};
// 删除按钮
const handleDelete = async (row: PayWalletRechargePackage.ResPayWalletRechargePackageItem) => {
  await useHandleData(deletePayWalletRechargePackageApi, Number(row.id), "删除充值套餐表");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: PayWalletRechargePackage.ResPayWalletRechargePackageItem) => {
  await useHandleData(recoverPayWalletRechargePackageApi, Number(row.id), "恢复充值套餐表");
  proTable.value?.getTableList();
};
// 添加按钮
const handleAdd = () => {
  title.value = "新增充值套餐表";
  centerDialogVisible.value = true;
  reset();
};

// 编辑按钮
const handleUpdate = async (row: PayWalletRechargePackage.ResPayWalletRechargePackageItem) => {
  title.value = "编辑充值套餐表";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getPayWalletRechargePackageItemApi(Number(row.id));
  payWalletRechargePackageItemFrom.value = data;
  payWalletRechargePackageItemFrom.value.payPrice = fenToYuan(payWalletRechargePackageItemFrom.value.payPrice);
  payWalletRechargePackageItemFrom.value.bonusPrice = fenToYuan(payWalletRechargePackageItemFrom.value.bonusPrice);
};
</script>
<style lang="scss">
@import "@/styles/custom.scss";
</style>
