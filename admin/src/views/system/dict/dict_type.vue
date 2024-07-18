<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="字典列表"
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
        <el-button v-auth="'dict.SystemDictTypeUpdate'" type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)">
          编辑
        </el-button>
        <el-dropdown trigger="click">
          <el-button v-auth="['dict.SystemDictTypeDelete', 'dict.SystemDictList']" type="primary" link :icon="DArrowRight"
            >更多</el-button
          >
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'dict.SystemDictList'" :icon="DataBoard" @click="toRouteLabel(scope.row)">
                数据
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
          <el-input v-model="systemDictTypeItemFrom.name" />
        </el-form-item>
        <el-form-item label="字典类型" prop="type">
          <el-input v-model="systemDictTypeItemFrom.type" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="systemDictTypeItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="systemDictTypeItemFrom.remark" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
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
import { useRouter } from "vue-router";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, DataBoard, DArrowRight } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemDictType } from "@/api/interface/systemDictType";
import {
  getSystemDictTypeListApi,
  deleteSystemDictTypeApi,
  getSystemDictTypeItemApi,
  addSystemDictTypeApi,
  updateSystemDictTypeApi
} from "@/api/modules/systemDictType";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
const router = useRouter();
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
  id: 0, //字典主键
  name: "", //字典名称
  type: "", //字典类型
  status: 0, //状态（0正常 1停用）
  remark: "", //备注
  creator: "", //创建人
  createTime: "", //创建时间
  updater: "", //更新人
  updateTime: "" //更新时间
});
//校验
const refSystemDictTypeItemFrom = ref<FormInstance>();
//校验
const rulesSystemDictTypeItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "请输入字典名称", trigger: "blur" }],
  type: [{ required: true, message: "请输入字典类型", trigger: "blur" }],
  status: [{ required: true, message: "请选择状态", trigger: "blur" }]
});
//菜单状态
const statusEnum = getIntDictOptions("status");

const columns: ColumnProps<SystemDictType.ResSystemDictTypeItem>[] = [
  { prop: "id", label: "编号", width: 100, fixed: "left" },
  { prop: "name", label: "字典名称", search: { el: "input", span: 2 } },
  { prop: "type", label: "字典类型", search: { el: "input", span: 2 } },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 }, width: 100 },
  { prop: "remark", label: "备注" },
  {
    prop: "operation",
    label: "操作",
    width: 160,
    fixed: "right",
    isShow: HasPermission("dict.SystemDictTypeDelete", "dict.SystemDictList", "dict.SystemDictTypeUpdate")
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  systemDictTypeItemFrom.value = {
    id: 0, //字典主键
    name: "", //字典名称
    type: "", //字典类型
    status: 0, //状态（0正常 1停用）
    remark: "", //备注
    creator: "", //创建人
    createTime: "", //创建时间
    updater: "", //更新人
    updateTime: "" //更新时间
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
    const data = systemDictTypeItemFrom.value as unknown as SystemDictType.ResSystemDictTypeItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemDictTypeApi, data.id, data, "修改字典");
    } else {
      await useHandleData(addSystemDictTypeApi, data, "新增字典");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

// 删除按钮
const handleDelete = async (row: SystemDictType.ResSystemDictTypeItem) => {
  await useHandleData(deleteSystemDictTypeApi, row.id, "删除字典");
  proTable.value?.getTableList();
};

// 添加按钮
const handleAdd = () => {
  title.value = "新增字典";
  centerDialogVisible.value = true;
  reset();
};

// 编辑按钮
const handleUpdate = async (row: SystemDictType.ResSystemDictTypeItem) => {
  title.value = "编辑字典";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemDictTypeItemApi(Number(row.id));
  systemDictTypeItemFrom.value = data;
};

// 跳转链接
const toRouteLabel = (row: SystemDictType.ResSystemDictTypeItem) => {
  router.push({
    name: "systemDict",
    params: {
      type: row.type
    }
  });
};
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
