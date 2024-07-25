<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="邮箱账号列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemMailAccountListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'mail.SystemMailAccountCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 是否SSL -->
      <template #sslEnable="scope">
        <DictTag type="mailAccount.ssl" :value="scope.row.sslEnable" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'mail.SystemMailAccount'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="[
              'mail.SystemMailAccountUpdate',
              'mail.SystemMailAccountDelete',
              'mail.SystemMailAccountRecover',
              'mail.SystemMailAccountDrop'
            ]"
            type="primary"
            link
            :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'mail.SystemMailAccountUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                编辑
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 0"
                v-auth="'mail.SystemMailAccountDelete'"
                :icon="Delete"
                @click="handleDelete(scope.row)">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'mail.SystemMailAccountRecover'"
                :icon="Refresh"
                @click="handleRecover(scope.row)">
                恢复
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'mail.SystemMailAccountDrop'"
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
      <el-form
        ref="refSystemMailAccountItemFrom"
        :model="systemMailAccountItemFrom"
        :rules="rulesSystemMailAccountItemFrom"
        label-width="100px">
        <el-form-item label="邮箱" prop="mail">
          <el-input v-model="systemMailAccountItemFrom.mail" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="用户名" prop="username">
          <el-input v-model="systemMailAccountItemFrom.username" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="systemMailAccountItemFrom.password" :disabled="disabled" type="password" />
        </el-form-item>
        <el-form-item label="服务器域名" prop="host">
          <el-input v-model="systemMailAccountItemFrom.host" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="服务器端口" prop="port">
          <el-input v-model="systemMailAccountItemFrom.port" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="是否SSL" prop="sslEnable">
          <el-radio-group v-model="systemMailAccountItemFrom.sslEnable">
            <el-radio-button v-for="dict in sslEnum" :key="Number(dict.value)" :value="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer v-if="!disabled">
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemMailAccountItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSystemMailAccountItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemMailAccount">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh, DeleteFilled, View, DArrowRight } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemMailAccount } from "@/api/interface/systemMailAccount";
import {
  getSystemMailAccountListApi,
  deleteSystemMailAccountApi,
  dropSystemMailAccountApi,
  recoverSystemMailAccountApi,
  getSystemMailAccountApi,
  addSystemMailAccountApi,
  updateSystemMailAccountApi
} from "@/api/modules/systemMailAccount";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
const disabled = ref(true);
//加载
const loading = ref(false);
//弹出层标题
const title = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
//数据接口
const systemMailAccountItemFrom = ref<SystemMailAccount.ResSystemMailAccountItem>({
  id: 0, // 主键
  mail: "", // 邮箱
  username: "", // 用户名
  password: "", // 密码
  host: "", // SMTP 服务器域名
  port: 0, // SMTP 服务器端口
  sslEnable: 0, // 是否开启 SSL
  creator: undefined, // 创建者
  createTime: "", // 创建时间
  updater: undefined, // 更新者
  updateTime: "", // 更新时间
  deleted: 0 // 是否删除
});
//校验
const refSystemMailAccountItemFrom = ref<FormInstance>();
//校验
const rulesSystemMailAccountItemFrom = reactive<FormRules>({
  mail: [{ required: true, message: "邮箱不能为空", trigger: "blur" }],
  username: [{ required: true, message: "用户名不能为空", trigger: "blur" }],
  password: [{ required: true, message: "密码不能为空", trigger: "blur" }],
  host: [{ required: true, message: "服务器域名不能为空", trigger: "blur" }],
  port: [{ required: true, message: "服务器端口不能为空", trigger: "blur" }],
  sslEnable: [{ required: true, message: "是否SSL不能为空", trigger: "change" }]
});
//删除状态
const deletedEnum = getIntDictOptions("delete");
//状态
const sslEnum = getIntDictOptions("mailAccount.ssl");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("mail.SystemMailAccountDelete")
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

const columns: ColumnProps<SystemMailAccount.ResSystemMailAccountItem>[] = [
  { prop: "id", label: "编号", width: 100, fixed: "left" },
  { prop: "mail", label: "邮箱", search: { el: "input", span: 2, props: { placeholder: "请输入" } } },
  { prop: "username", label: "用户名", search: { el: "input", span: 2, props: { placeholder: "请输入" } } },
  { prop: "password", label: "密码" },
  { prop: "host", label: "服务器域名" },
  { prop: "port", label: "服务器端口" },
  { prop: "sslEnable", label: "是否SSL" },
  { prop: "creator", label: "创建者" },
  { prop: "createTime", label: "创建时间" },
  { prop: "updater", label: "更新者" },
  { prop: "updateTime", label: "更新时间" },
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
    isShow: HasPermission(
      "mail.SystemMailAccountUpdate",
      "mail.SystemMailAccountDelete",
      "mail.SystemMailAccountDrop",
      "mail.SystemMailAccountRecover",
      "mail.SystemMailAccount"
    )
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  systemMailAccountItemFrom.value = {
    id: 0, // 主键
    mail: "", // 邮箱
    username: "", // 用户名
    password: "", // 密码
    host: "", // SMTP 服务器域名
    port: 0, // SMTP 服务器端口
    sslEnable: 0, // 是否开启 SSL
    creator: undefined, // 创建者
    createTime: "", // 创建时间
    updater: undefined, // 更新者
    updateTime: "", // 更新时间
    deleted: 0 // 是否删除
  };
  disabled.value = true;
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
    const data = systemMailAccountItemFrom.value as unknown as SystemMailAccount.ResSystemMailAccountItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemMailAccountApi, data.id, data, "修改邮箱账号");
    } else {
      await useHandleData(addSystemMailAccountApi, data, "添加邮箱账号");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};
// 清理按钮
const handleDrop = async (row: SystemMailAccount.ResSystemMailAccountItem) => {
  await useHandleData(dropSystemMailAccountApi, Number(row.id), "清理邮箱账号");
  proTable.value?.getTableList();
};
// 删除按钮
const handleDelete = async (row: SystemMailAccount.ResSystemMailAccountItem) => {
  await useHandleData(deleteSystemMailAccountApi, Number(row.id), "删除邮箱账号");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: SystemMailAccount.ResSystemMailAccountItem) => {
  await useHandleData(recoverSystemMailAccountApi, Number(row.id), "恢复邮箱账号");
  proTable.value?.getTableList();
};
// 添加按钮
const handleAdd = () => {
  title.value = "新增邮箱账号";
  centerDialogVisible.value = true;
  reset();
  disabled.value = false;
};
// 编辑按钮
const handleUpdate = async (row: SystemMailAccount.ResSystemMailAccountItem) => {
  title.value = "编辑邮箱账号";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemMailAccountApi(Number(row.id));
  systemMailAccountItemFrom.value = data;
  disabled.value = false;
};
// 查看按钮
const handleItem = async (row: SystemMailAccount.ResSystemMailAccountItem) => {
  title.value = "查看邮箱账号";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemMailAccountApi(Number(row.id));
  systemMailAccountItemFrom.value = data;
  disabled.value = true;
};
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
