<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="岗位列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemPostListApi"
      :request-auto="true"
      :pagination="false"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
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
        <el-button type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)"> 编辑 </el-button>
        <el-button v-if="scope.row.deleted === 0" type="primary" link :icon="Delete" @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button v-if="scope.row.deleted === 1" type="primary" link :icon="Refresh" @click="handleRecover(scope.row)">
          恢复
        </el-button>
      </template>
    </ProTable>
    <el-dialog
      :title="title"
      v-model="centerDialogVisible"
      width="40%"
      destroy-on-close
      align-center
      center
      append-to-body
      draggable
      :lock-scroll="false"
      class="dialog-settings">
      <el-form ref="refSystemPostItemFrom" :model="systemPostItemFrom" :rules="rulesSystemPostItemFrom" label-width="100px">
        <el-form-item label="岗位名称" prop="name">
          <el-input v-model="systemPostItemFrom.name" />
        </el-form-item>
        <el-form-item label="岗位顺序" prop="sort">
          <el-input-number v-model="systemPostItemFrom.sort" controls-position="right" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="systemPostItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemPostItemFrom)">取消</el-button>
          <el-button type="primary" @click="submitForm(refSystemPostItemFrom)" :loading="loading">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemPost">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemPost } from "@/api/interface/systemPost";
import {
  getSystemPostListApi,
  deleteSystemPostApi,
  recoverSystemPostApi,
  getSystemPostItemApi,
  addSystemPostApi,
  updateSystemPostApi
} from "@/api/modules/systemPost";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
//加载
const loading = ref(false);
//弹出层标题
const title = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
//数据接口
const systemPostItemFrom = ref<SystemPost.ResSystemPostItem>({
  id: 0, //bigint 岗位ID,PRI
  name: "", //varchar 岗位名称
  sort: 0, //int 显示顺序
  status: 0, //tinyint 状态（0正常 1停用）
  deleted: 0 //tinyint 是否删除
});
//校验
const refSystemPostItemFrom = ref<FormInstance>();
//校验
const rulesSystemPostItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "岗位标题不能为空", trigger: "blur" }],
  status: [{ required: true, message: "岗位状态不能为空", trigger: "change" }]
});

//状态
const statusEnum = getIntDictOptions("status");
//删除状态
const deletedEnum = getIntDictOptions("delete");
// 表格配置项
const deleteSearch = reactive<SearchProps>({
  el: "switch",
  span: 2
});
const columns: ColumnProps<SystemPost.ResSystemPostItem>[] = [
  { prop: "id", label: "编号", width: 100 },
  { prop: "name", label: "岗位名称", search: { el: "input", span: 2, props: { placeholder: "请输入名称" } } },
  { prop: "sort", label: "岗位顺序" },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "tree-select", span: 2 } },
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
    fixed: "right"
  }
];
// 重置数据
const reset = () => {
  systemPostItemFrom.value = {
    id: 0, //bigint 岗位ID,PRI
    name: "", //varchar 岗位名称
    sort: 0, //int 显示顺序
    status: 0, //tinyint 状态（0正常 1停用）
    deleted: 0 //tinyint 是否删除
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
    const data = systemPostItemFrom.value as unknown as SystemPost.ResSystemPostItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemPostApi, data.id, data, "修改职位");
    } else {
      await useHandleData(addSystemPostApi, data, "添加职位");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

// 删除按钮
const handleDelete = async (row: SystemPost.ResSystemPostItem) => {
  await useHandleData(deleteSystemPostApi, Number(row.id), "删除职位");
  proTable.value?.getTableList();
};

// 恢复按钮
const handleRecover = async (row: SystemPost.ResSystemPostItem) => {
  await useHandleData(recoverSystemPostApi, Number(row.id), "恢复职位");
  proTable.value?.getTableList();
};

// 添加按钮
const handleAdd = () => {
  title.value = "新增岗位";
  centerDialogVisible.value = true;
  reset();
};

// 编辑按钮
const handleUpdate = async (row: SystemPost.ResSystemPostItem) => {
  title.value = "编辑岗位";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemPostItemApi(Number(row.id));
  systemPostItemFrom.value = data;
};
</script>

<style lang="scss">
@import "@/styles/custom.scss";
</style>
