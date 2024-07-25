<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="字典类型列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemDictTypeListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'dict.SystemDictTypeCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <!-- 状态-->
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'dict.SystemDictType'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['dict.SystemDictTypeUpdate', 'dict.SystemDictTypeDelete', 'dict.SystemDictList']"
            type="primary"
            link
            :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'dict.SystemDictList'" :icon="DataBoard" @click="routeSystemDict(scope.row)">
                数据
              </el-dropdown-item>
              <el-dropdown-item v-auth="'dict.SystemDictTypeUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                编辑
              </el-dropdown-item>
              <el-dropdown-item v-auth="'dict.SystemDictTypeDelete'" :icon="Delete" @click="handleDelete(scope.row)">
                删除
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
        ref="refSystemDictTypeItemFrom"
        :model="systemDictTypeItemFrom"
        :rules="rulesSystemDictTypeItemFrom"
        label-width="100px">
        <el-form-item label="字典名称" prop="name">
          <el-input v-model="systemDictTypeItemFrom.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="字典类型" prop="type">
          <el-input v-model="systemDictTypeItemFrom.type" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="systemDictTypeItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value" :disabled="disabled">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="systemDictTypeItemFrom.remark" :disabled="disabled" />
        </el-form-item>
      </el-form>
      <template #footer v-if="!disabled">
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemDictTypeItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSystemDictTypeItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemDictType">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, View, DArrowRight, DataBoard } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemDictType } from "@/api/interface/systemDictType";
import {
  getSystemDictTypeListApi,
  deleteSystemDictTypeApi,
  getSystemDictTypeApi,
  addSystemDictTypeApi,
  updateSystemDictTypeApi
} from "@/api/modules/systemDictType";
import { FormInstance, FormRules } from "element-plus";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import { useRouter } from "vue-router";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
const router = useRouter();
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
const systemDictTypeItemFrom = ref<SystemDictType.ResSystemDictTypeItem>({
  id: 0, // 字典主键
  name: "", // 字典名称
  type: "", // 字典类型
  status: 0, // 状态（0正常 1停用）
  remark: undefined, // 备注
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined // 更新时间
});
//校验
const refSystemDictTypeItemFrom = ref<FormInstance>();
//校验
const rulesSystemDictTypeItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "字典名称不能为空", trigger: "blur" }],
  type: [{ required: true, message: "字典类型不能为空", trigger: "blur" }],
  status: [{ required: true, message: "状态不能为空", trigger: "blur" }]
});

//菜单状态
const statusEnum = getIntDictOptions("status");
const columns: ColumnProps<SystemDictType.ResSystemDictTypeItem>[] = [
  { prop: "id", label: "编号", width: 100, fixed: "left" },
  { prop: "name", label: "字典名称", search: { el: "input", span: 2 } },
  { prop: "type", label: "字典类型", search: { el: "input", span: 2 } },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 }, width: 100 },
  { prop: "remark", label: "备注" },
  { prop: "creator", label: "创建者" },
  { prop: "createTime", label: "创建时间" },
  { prop: "updater", label: "更新者" },
  { prop: "updateTime", label: "更新时间" },
  {
    prop: "operation",
    label: "操作",
    width: 160,
    fixed: "right",
    isShow: HasPermission("dict.SystemDictTypeDelete", "dict.SystemDictList", "dict.SystemDictTypeUpdate", "dict.SystemDictType")
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  systemDictTypeItemFrom.value = {
    id: 0, // 字典主键
    name: "", // 字典名称
    type: "", // 字典类型
    status: 0, // 状态（0正常 1停用）
    remark: undefined, // 备注
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
    const data = systemDictTypeItemFrom.value as unknown as SystemDictType.ResSystemDictTypeItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemDictTypeApi, data.id, data, "修改字典类型");
    } else {
      await useHandleData(addSystemDictTypeApi, data, "添加字典类型");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};
// 删除按钮
const handleDelete = async (row: SystemDictType.ResSystemDictTypeItem) => {
  await useHandleData(deleteSystemDictTypeApi, Number(row.id), "删除字典类型");
  proTable.value?.getTableList();
};
// 添加按钮
const handleAdd = () => {
  title.value = "新增字典类型";
  centerDialogVisible.value = true;
  reset();
  disabled.value = false;
};
// 编辑按钮
const handleUpdate = async (row: SystemDictType.ResSystemDictTypeItem) => {
  title.value = "编辑字典类型";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemDictTypeApi(Number(row.id));
  systemDictTypeItemFrom.value = data;
  disabled.value = false;
};
// 查看按钮
const handleItem = async (row: SystemDictType.ResSystemDictTypeItem) => {
  title.value = "查看字典类型";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemDictTypeApi(Number(row.id));
  systemDictTypeItemFrom.value = data;
  disabled.value = true;
};

// 跳转链接
const routeSystemDict = (row: SystemDictType.ResSystemDictTypeItem) => {
  router.push({
    name: "systemDict",
    params: {
      dictTypeId: row.id
    }
  });
};
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
