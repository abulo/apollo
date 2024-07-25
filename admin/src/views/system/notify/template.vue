<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="模板列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemNotifyTemplateListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'notify.SystemNotifyTemplateCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">
          新增
        </el-button>
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
        <el-button v-auth="'notify.SystemNotifyTemplate'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="[
              'notify.SystemNotifyTemplateUpdate',
              'notify.SystemNotifyTemplateDelete',
              'notify.SystemNotifyTemplateRecover',
              'notify.SystemNotifyTemplateDrop'
            ]"
            type="primary"
            link
            :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'notify.SystemNotifyTemplateUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                编辑
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 0"
                v-auth="'notify.SystemNotifyTemplateDelete'"
                :icon="Delete"
                @click="handleDelete(scope.row)">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'notify.SystemNotifyTemplateRecover'"
                :icon="Refresh"
                @click="handleRecover(scope.row)">
                恢复
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'notify.SystemNotifyTemplateDrop'"
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
        ref="refSystemNotifyTemplateItemFrom"
        :model="systemNotifyTemplateItemFrom"
        :rules="rulesSystemNotifyTemplateItemFrom"
        label-width="100px">
        <el-form-item label="模板名称" prop="name">
          <el-input v-model="systemNotifyTemplateItemFrom.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="模版编码" prop="code">
          <el-input v-model="systemNotifyTemplateItemFrom.code" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="发送人名称" prop="nickname">
          <el-input v-model="systemNotifyTemplateItemFrom.nickname" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="模版内容" prop="content">
          <el-input v-model="systemNotifyTemplateItemFrom.content" type="textarea" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="模版类型" prop="type">
          <el-radio-group v-model="systemNotifyTemplateItemFrom.type">
            <el-radio-button v-for="dict in typeEnum" :key="Number(dict.value)" :value="dict.value" :disabled="disabled">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="模版参数" prop="params">
          <el-input v-model="systemNotifyTemplateItemFrom.params" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="模版状态" prop="status">
          <el-radio-group v-model="systemNotifyTemplateItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value" :disabled="disabled">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="systemNotifyTemplateItemFrom.remark" :disabled="disabled" />
        </el-form-item>
      </el-form>
      <template #footer v-if="!disabled">
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemNotifyTemplateItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSystemNotifyTemplateItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemNotifyTemplate">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh, DeleteFilled, View, DArrowRight } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemNotifyTemplate } from "@/api/interface/systemNotifyTemplate";
import {
  getSystemNotifyTemplateListApi,
  deleteSystemNotifyTemplateApi,
  dropSystemNotifyTemplateApi,
  recoverSystemNotifyTemplateApi,
  getSystemNotifyTemplateApi,
  addSystemNotifyTemplateApi,
  updateSystemNotifyTemplateApi
} from "@/api/modules/systemNotifyTemplate";
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
const systemNotifyTemplateItemFrom = ref<SystemNotifyTemplate.ResSystemNotifyTemplateItem>({
  id: 0, // 主键
  name: "", // 模板名称
  code: "", // 模版编码
  nickname: "", // 发送人名称
  content: "", // 模版内容
  type: 0, // 类型
  params: undefined, // 参数数组
  status: 0, // 状态
  remark: undefined, // 备注
  deleted: 0, // 删除
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined // 更新时间
});
//校验
const refSystemNotifyTemplateItemFrom = ref<FormInstance>();
//校验
const rulesSystemNotifyTemplateItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "模板名称不能为空", trigger: "blur" }],
  code: [{ required: true, message: "模板编码不能为空", trigger: "blur" }],
  nickname: [{ required: true, message: "发送人名称不能为空", trigger: "blur" }],
  type: [{ required: true, message: "模板类别不能为空", trigger: "blur" }],
  content: [{ required: true, message: "模板内容不能为空", trigger: "blur" }],
  status: [{ required: true, message: "模板状态不能为空", trigger: "change" }]
});
//状态
const statusEnum = getIntDictOptions("status");
//删除状态
const deletedEnum = getIntDictOptions("delete");
//模板类型
const typeEnum = getIntDictOptions("notifyTemplate.type");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("notify.SystemNotifyTemplateDelete")
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

const columns: ColumnProps<SystemNotifyTemplate.ResSystemNotifyTemplateItem>[] = [
  { prop: "id", label: "编号", width: 100, fixed: "left" },
  { prop: "name", label: "模板名称", search: { el: "input", span: 2, props: { placeholder: "请输入名称" } } },
  { prop: "code", label: "模板编码" },
  { prop: "nickname", label: "发送人" },
  { prop: "type", label: "模板类型", tag: true, enum: typeEnum, search: { el: "select", span: 2 } },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 } },
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
      "notify.SystemNotifyTemplateUpdate",
      "notify.SystemNotifyTemplateDelete",
      "notify.SystemNotifyTemplateDrop",
      "notify.SystemNotifyTemplateRecover",
      "notify.SystemNotifyTemplate"
    )
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  systemNotifyTemplateItemFrom.value = {
    id: 0, // 主键
    name: "", // 模板名称
    code: "", // 模版编码
    nickname: "", // 发送人名称
    content: "", // 模版内容
    type: 0, // 类型
    params: undefined, // 参数数组
    status: 0, // 状态
    remark: undefined, // 备注
    deleted: 0, // 删除
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined // 更新时间
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
    const data = systemNotifyTemplateItemFrom.value as unknown as SystemNotifyTemplate.ResSystemNotifyTemplateItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemNotifyTemplateApi, data.id, data, "修改模板");
    } else {
      await useHandleData(addSystemNotifyTemplateApi, data, "添加模板");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};
// 清理按钮
const handleDrop = async (row: SystemNotifyTemplate.ResSystemNotifyTemplateItem) => {
  await useHandleData(dropSystemNotifyTemplateApi, Number(row.id), "清理模板");
  proTable.value?.getTableList();
};
// 删除按钮
const handleDelete = async (row: SystemNotifyTemplate.ResSystemNotifyTemplateItem) => {
  await useHandleData(deleteSystemNotifyTemplateApi, Number(row.id), "删除模板");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: SystemNotifyTemplate.ResSystemNotifyTemplateItem) => {
  await useHandleData(recoverSystemNotifyTemplateApi, Number(row.id), "恢复模板");
  proTable.value?.getTableList();
};
// 添加按钮
const handleAdd = () => {
  title.value = "新增模板";
  centerDialogVisible.value = true;
  reset();
  disabled.value = false;
};
// 编辑按钮
const handleUpdate = async (row: SystemNotifyTemplate.ResSystemNotifyTemplateItem) => {
  title.value = "编辑模板";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemNotifyTemplateApi(Number(row.id));
  systemNotifyTemplateItemFrom.value = data;
  disabled.value = false;
};
// 查看按钮
const handleItem = async (row: SystemNotifyTemplate.ResSystemNotifyTemplateItem) => {
  title.value = "查看模板";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemNotifyTemplateApi(Number(row.id));
  systemNotifyTemplateItemFrom.value = data;
  disabled.value = true;
};
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
