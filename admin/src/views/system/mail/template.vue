<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="模板列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemMailTemplateListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'mail.SystemMailTemplateCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">
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
        <el-button v-auth="'mail.SystemMailTemplateUpdate'" type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)">
          编辑
        </el-button>
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'mail.SystemMailTemplateDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'mail.SystemMailTemplateRecover'"
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
        ref="refSystemMailTemplateItemFrom"
        :model="systemMailTemplateItemFrom"
        :rules="rulesSystemMailTemplateItemFrom"
        label-width="100px">
        <el-form-item label="模板名称" prop="name">
          <el-input v-model="systemMailTemplateItemFrom.name" />
        </el-form-item>
        <el-form-item label="模板编码" prop="code">
          <el-input v-model="systemMailTemplateItemFrom.code" />
        </el-form-item>
        <el-form-item label="发送人" prop="nickname">
          <el-input v-model="systemMailTemplateItemFrom.nickname" />
        </el-form-item>
        <el-form-item label="模板内容" prop="content">
          <el-input v-model="systemMailTemplateItemFrom.content" type="textarea" />
        </el-form-item>
        <el-form-item label="邮箱账号" prop="accountId">
          <el-select v-model="systemMailTemplateItemFrom.accountId" placeholder="请选择邮箱账号">
            <el-option v-for="item in accountEnum" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="模板状态" prop="status">
          <el-radio-group v-model="systemMailTemplateItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="systemMailTemplateItemFrom.remark" />
        </el-form-item>
      </el-form>
      <template #footer>
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
import { EditPen, CirclePlus, Delete, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemMailTemplate } from "@/api/interface/systemMailTemplate";
import {
  getSystemMailTemplateListApi,
  deleteSystemMailTemplateApi,
  recoverSystemMailTemplateApi,
  getSystemMailTemplateItemApi,
  addSystemMailTemplateApi,
  updateSystemMailTemplateApi
} from "@/api/modules/systemMailTemplate";
import { FormInstance, FormRules } from "element-plus";
import { getSystemMailAccountSearchApi } from "@/api/modules/systemMailAccount";
import { getIntDictOptions, DictDataType } from "@/utils/dict";
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
const systemMailTemplateItemFrom = ref<SystemMailTemplate.ResSystemMailTemplateItem>({
  id: 0, // 模板 id
  name: "", // 模板名称
  code: "", // 模版编码
  accountId: undefined, // 发送的邮箱账号编号
  nickname: "", // 发送人名称
  title: "", // 模板标题
  content: "", // 模版内容
  params: {}, // 参数数组
  status: 0, // 状态
  remark: "", // 备注
  deleted: 0 // 删除
});
//校验
const refSystemMailTemplateItemFrom = ref<FormInstance>();
//校验
const rulesSystemMailTemplateItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "模板名称不能为空", trigger: "blur" }],
  code: [{ required: true, message: "模板编码不能为空", trigger: "blur" }],
  accountId: [{ required: true, message: "发送的邮箱账号编号不能为空", trigger: "blur" }],
  nickname: [{ required: true, message: "发送人名称不能为空", trigger: "blur" }],
  type: [{ required: true, message: "模板类别不能为空", trigger: "blur" }],
  content: [{ required: true, message: "模板内容不能为空", trigger: "blur" }],
  status: [{ required: true, message: "模板状态不能为空", trigger: "change" }]
});

//状态
const statusEnum = getIntDictOptions("status");
//删除状态
const deletedEnum = getIntDictOptions("delete");
//账号
const accountEnum = reactive<DictDataType[]>([]);
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("mail.SystemMailTemplateDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);
const columns: ColumnProps<SystemMailTemplate.ResSystemMailTemplateItem>[] = [
  { prop: "id", label: "编号", width: 100 },
  { prop: "name", label: "模板名称", search: { el: "input", span: 2, props: { placeholder: "请输入名称" } } },
  { prop: "code", label: "模板编码", search: { el: "input", span: 2, props: { placeholder: "请输入编码" } } },
  { prop: "accountId", label: "邮箱账号", tag: true, enum: accountEnum, search: { el: "select", span: 2 } },
  { prop: "nickname", label: "发送人" },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 } },
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
    isShow: HasPermission("mail.SystemMailTemplateUpdate", "mail.SystemMailTemplateDelete", "mail.SystemMailTemplateRecover")
  }
];
// 重置数据
const reset = () => {
  loading.value = false;
  systemMailTemplateItemFrom.value = {
    id: 0, // 模板 id
    name: "", // 模板名称
    code: "", // 模版编码
    accountId: undefined, // 发送的邮箱账号编号
    nickname: "", // 发送人名称
    title: "", // 模板标题
    content: "", // 模版内容
    params: {}, // 参数数组
    status: 0, // 状态
    remark: "", // 备注
    deleted: 0 // 删除
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
    const data = systemMailTemplateItemFrom.value as unknown as SystemMailTemplate.ResSystemMailTemplateItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemMailTemplateApi, data.id, data, "修改模板");
    } else {
      await useHandleData(addSystemMailTemplateApi, data, "添加模板");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

// 删除按钮
const handleDelete = async (row: SystemMailTemplate.ResSystemMailTemplateItem) => {
  await useHandleData(deleteSystemMailTemplateApi, Number(row.id), "删除模板");
  proTable.value?.getTableList();
};

// 恢复按钮
const handleRecover = async (row: SystemMailTemplate.ResSystemMailTemplateItem) => {
  await useHandleData(recoverSystemMailTemplateApi, Number(row.id), "恢复模板");
  proTable.value?.getTableList();
};

// 添加按钮
const handleAdd = () => {
  title.value = "新增模板";
  centerDialogVisible.value = true;
  reset();
};

// 编辑按钮
const handleUpdate = async (row: SystemMailTemplate.ResSystemMailTemplateItem) => {
  title.value = "编辑模板";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemMailTemplateItemApi(Number(row.id));
  systemMailTemplateItemFrom.value = data;
};

const getAccount = async () => {
  const { data } = await getSystemMailAccountSearchApi();
  //扁平化输出数据
  data.forEach(item => {
    let obj: DictDataType = {
      label: item.mail,
      value: Number(item.id),
      colorType: "",
      cssClass: "",
      dictType: "account"
    };
    accountEnum.push(obj);
  });
};

onMounted(() => {
  getAccount();
});
</script>

<style lang="scss">
@import "@/styles/custom.scss";
</style>
