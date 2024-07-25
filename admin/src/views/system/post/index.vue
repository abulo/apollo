<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="职位列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemPostListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'post.SystemPostCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <!-- 删除状态 -->
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
        <el-button v-auth="'post.SystemPost'" type="primary" link :icon="View" @click="handleItem(scope.row)"> 查看 </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['post.SystemPostUpdate', 'post.SystemPostDelete', 'post.SystemPostRecover', 'post.SystemPostDrop']"
            type="primary"
            link
            :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'post.SystemPostUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                编辑
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 0"
                v-auth="'post.SystemPostDelete'"
                :icon="Delete"
                @click="handleDelete(scope.row)">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'post.SystemPostRecover'"
                :icon="Refresh"
                @click="handleRecover(scope.row)">
                恢复
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'post.SystemPostDrop'"
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
      <el-form ref="refSystemPostItemFrom" :model="systemPostItemFrom" :rules="rulesSystemPostItemFrom" label-width="100px">
        <el-form-item label="职位名称" prop="name">
          <el-input v-model="systemPostItemFrom.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="显示顺序" prop="sort">
          <el-input-number v-model="systemPostItemFrom.sort" controls-position="right" :min="0" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="systemPostItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value" :disabled="disabled">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer v-if="!disabled">
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemPostItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSystemPostItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemPost">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh, DeleteFilled, View, DArrowRight } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemPost } from "@/api/interface/systemPost";
import {
  getSystemPostListApi,
  deleteSystemPostApi,
  dropSystemPostApi,
  recoverSystemPostApi,
  getSystemPostApi,
  addSystemPostApi,
  updateSystemPostApi
} from "@/api/modules/systemPost";
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
const systemPostItemFrom = ref<SystemPost.ResSystemPostItem>({
  id: 0, // 职位ID
  name: "", // 职位名称
  sort: 0, // 显示顺序
  status: 0, // 状态
  deleted: 0, // 是否删除
  tenantId: 0, // 租户ID
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined // 更新时间
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
const deleteSearch = reactive<SearchProps>(
  HasPermission("post.SystemPostDelete")
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

const columns: ColumnProps<SystemPost.ResSystemPostItem>[] = [
  { prop: "id", label: "编号", width: 100, fixed: "left" },
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
      "post.SystemPostUpdate",
      "post.SystemPostDelete",
      "post.SystemPostDrop",
      "post.SystemPostRecover",
      "post.SystemPost"
    )
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  systemPostItemFrom.value = {
    id: 0, // 职位ID
    name: "", // 职位名称
    sort: 0, // 显示顺序
    status: 0, // 状态
    deleted: 0, // 是否删除
    tenantId: 0, // 租户ID
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
// 清理按钮
const handleDrop = async (row: SystemPost.ResSystemPostItem) => {
  await useHandleData(dropSystemPostApi, Number(row.id), "清理职位");
  proTable.value?.getTableList();
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
  title.value = "新增职位";
  centerDialogVisible.value = true;
  reset();
  disabled.value = false;
};
// 编辑按钮
const handleUpdate = async (row: SystemPost.ResSystemPostItem) => {
  title.value = "编辑职位";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemPostApi(Number(row.id));
  systemPostItemFrom.value = data;
  disabled.value = false;
};
// 查看按钮
const handleItem = async (row: SystemPost.ResSystemPostItem) => {
  title.value = "查看职位";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemPostApi(Number(row.id));
  systemPostItemFrom.value = data;
  disabled.value = true;
};
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
