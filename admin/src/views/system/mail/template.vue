<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="邮件模版列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemMailTemplateListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'mail.SystemMailTemplateCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'mail.SystemMailTemplate'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="[
              'mail.SystemMailTemplateUpdate',
              'mail.SystemMailTemplateDelete',
              'mail.SystemMailTemplateRecover',
              'mail.SystemMailTemplateDrop'
            ]"
            type="primary"
            link
            :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'mail.SystemMailTemplateUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                编辑
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 0"
                v-auth="'mail.SystemMailTemplateDelete'"
                :icon="Delete"
                @click="handleDelete(scope.row)">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'mail.SystemMailTemplateRecover'"
                :icon="Refresh"
                @click="handleRecover(scope.row)">
                恢复
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'mail.SystemMailTemplateDrop'"
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
        ref="refSystemMailTemplateItemFrom"
        :model="systemMailTemplateItemFrom"
        :rules="rulesSystemMailTemplateItemFrom"
        label-width="100px">
        <el-form-item label="模板名称" prop="name">
          <el-input v-model="systemMailTemplateItemFrom.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="模板编码" prop="code">
          <el-input v-model="systemMailTemplateItemFrom.code" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="邮箱账号" prop="accountId">
          <el-select v-model="systemMailTemplateItemFrom.accountId" placeholder="请选择邮箱账号" :disabled="disabled">
            <el-option v-for="item in mailAccount" :key="item.id" :label="item.mail" :value="item.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="发送人名称" prop="nickname">
          <el-input v-model="systemMailTemplateItemFrom.nickname" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="模板标题" prop="title">
          <el-input v-model="systemMailTemplateItemFrom.title" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="模板内容" prop="content">
          <el-input v-model="systemMailTemplateItemFrom.content" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="参数数组" prop="params">
          <el-input v-model="systemMailTemplateItemFrom.params" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="开启状态" prop="status">
          <el-radio-group v-model="systemMailTemplateItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value" :disabled="disabled">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="systemMailTemplateItemFrom.remark" :disabled="disabled" />
        </el-form-item>
      </el-form>
      <template #footer v-if="!disabled">
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemMailTemplateItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSystemMailTemplateItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemMailTemplate">
import { ref, reactive, onMounted } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh, DeleteFilled, View, DArrowRight } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemMailTemplate } from "@/api/interface/systemMailTemplate";
import {
  getSystemMailTemplateListApi,
  deleteSystemMailTemplateApi,
  dropSystemMailTemplateApi,
  recoverSystemMailTemplateApi,
  getSystemMailTemplateApi,
  addSystemMailTemplateApi,
  updateSystemMailTemplateApi
} from "@/api/modules/systemMailTemplate";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import { SystemMailAccount } from "@/api/interface/systemMailAccount";
import { getSystemMailAccountListSimpleApi } from "@/api/modules/systemMailAccount";
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
const systemMailTemplateItemFrom = ref<SystemMailTemplate.ResSystemMailTemplateItem>({
  id: 0, // 编号
  name: "", // 模板名称
  code: "", // 模板编码
  accountId: 0, // 发送的邮箱账号编号
  nickname: undefined, // 发送人名称
  title: "", // 模板标题
  content: "", // 模板内容
  params: undefined, // 参数数组
  status: 0, // 开启状态
  remark: undefined, // 备注
  creator: undefined, // 创建者
  createTime: "", // 创建时间
  updater: undefined, // 更新者
  updateTime: "", // 更新时间
  deleted: 0 // 是否删除
});
//校验
const refSystemMailTemplateItemFrom = ref<FormInstance>();
//校验
const rulesSystemMailTemplateItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "模板名称不能为空", trigger: "blur" }],
  code: [{ required: true, message: "模板编码不能为空", trigger: "blur" }],
  accountId: [{ required: true, message: "发送的邮箱账号编号不能为空", trigger: "blur" }],
  title: [{ required: true, message: "模板标题不能为空", trigger: "blur" }],
  content: [{ required: true, message: "模板内容不能为空", trigger: "blur" }],
  status: [{ required: true, message: "开启状态不能为空", trigger: "blur" }]
});
//状态
const statusEnum = getIntDictOptions("status");
//删除状态
const deletedEnum = getIntDictOptions("delete");
// 邮箱账号
const mailAccount = ref<SystemMailAccount.ResSystemMailAccountItem[]>([]);
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("mail.SystemMailTemplateDelete")
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

const columns: ColumnProps<SystemMailTemplate.ResSystemMailTemplateItem>[] = [
  { prop: "id", label: "编号", width: 100, fixed: "left" },
  { prop: "name", label: "模板名称", search: { el: "input", span: 2, props: { placeholder: "请输入模板名称" } } },
  { prop: "code", label: "模板编码", search: { el: "input", span: 2, props: { placeholder: "请输入模板编码" } } },
  {
    prop: "accountId",
    label: "邮箱账号",
    tag: true,
    enum: mailAccount,
    fieldNames: { label: "mail", value: "id" },
    search: { el: "select", span: 2 }
  },
  { prop: "nickname", label: "发送人名称" },
  { prop: "title", label: "模板标题", search: { el: "input", span: 2, props: { placeholder: "请输入模板标题" } } },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 } },
  { prop: "remark", label: "备注" },
  { prop: "creator", label: "创建者" },
  { prop: "createTime", label: "创建时间" },
  { prop: "updater", label: "更新者" },
  { prop: "updateTime", label: "更新时间" },
  { prop: "deleted", label: "删除", tag: true, enum: deletedEnum, search: deleteSearch },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission(
      "mail.SystemMailTemplateUpdate",
      "mail.SystemMailTemplateDelete",
      "mail.SystemMailTemplateDrop",
      "mail.SystemMailTemplateRecover",
      "mail.SystemMailTemplate"
    )
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  systemMailTemplateItemFrom.value = {
    id: 0, // 编号
    name: "", // 模板名称
    code: "", // 模板编码
    accountId: undefined, // 发送的邮箱账号编号
    nickname: undefined, // 发送人名称
    title: "", // 模板标题
    content: "", // 模板内容
    params: undefined, // 参数数组
    status: 0, // 开启状态
    remark: undefined, // 备注
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
    const data = systemMailTemplateItemFrom.value as unknown as SystemMailTemplate.ResSystemMailTemplateItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemMailTemplateApi, data.id, data, "修改邮件模版");
    } else {
      await useHandleData(addSystemMailTemplateApi, data, "添加邮件模版");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};
// 清理按钮
const handleDrop = async (row: SystemMailTemplate.ResSystemMailTemplateItem) => {
  await useHandleData(dropSystemMailTemplateApi, Number(row.id), "清理邮件模版");
  proTable.value?.getTableList();
};
// 删除按钮
const handleDelete = async (row: SystemMailTemplate.ResSystemMailTemplateItem) => {
  await useHandleData(deleteSystemMailTemplateApi, Number(row.id), "删除邮件模版");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: SystemMailTemplate.ResSystemMailTemplateItem) => {
  await useHandleData(recoverSystemMailTemplateApi, Number(row.id), "恢复邮件模版");
  proTable.value?.getTableList();
};
// 添加按钮
const handleAdd = () => {
  title.value = "新增邮件模版";
  centerDialogVisible.value = true;
  reset();
  disabled.value = false;
};
// 编辑按钮
const handleUpdate = async (row: SystemMailTemplate.ResSystemMailTemplateItem) => {
  title.value = "编辑邮件模版";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemMailTemplateApi(Number(row.id));
  systemMailTemplateItemFrom.value = data;
  disabled.value = false;
};
// 查看按钮
const handleItem = async (row: SystemMailTemplate.ResSystemMailTemplateItem) => {
  title.value = "查看邮件模版";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemMailTemplateApi(Number(row.id));
  systemMailTemplateItemFrom.value = data;
  disabled.value = true;
};

const getMailAccount = async () => {
  const { data } = await getSystemMailAccountListSimpleApi({ deleted: 1 });
  mailAccount.value = data;
};

onMounted(() => {
  getMailAccount();
});
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
