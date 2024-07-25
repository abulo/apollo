<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="字典数据列表"
      row-key="id"
      :columns="columns"
      :request-api="getCustomSystemDictListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'dict.SystemDictCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
        <el-button type="primary" :icon="Remove" @click="closeCurrentTab">关闭</el-button>
      </template>
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'dict.SystemDict'" type="primary" link :icon="View" @click="handleItem(scope.row)"> 查看 </el-button>
        <el-dropdown trigger="click">
          <el-button v-auth="['dict.SystemDictUpdate', 'dict.SystemDictDelete']" type="primary" link :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'dict.SystemDictUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                编辑
              </el-dropdown-item>
              <el-dropdown-item v-auth="'dict.SystemDictDelete'" :icon="Delete" @click="handleDelete(scope.row)">
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
      <el-form ref="refSystemDictItemFrom" :model="systemDictItemFrom" :rules="rulesSystemDictItemFrom" label-width="100px">
        <el-form-item label="数据标签" prop="label">
          <el-input v-model="systemDictItemFrom.label" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="数据键值" prop="value">
          <el-input v-model="systemDictItemFrom.value" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="编码排序" prop="sort">
          <el-input-number v-model="systemDictItemFrom.sort" controls-position="right" :min="0" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="systemDictItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value" :disabled="disabled">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="颜色类型" prop="colorType">
          <el-select v-model="systemDictItemFrom.colorType" clearable :disabled="disabled">
            <el-option
              v-for="item in colorTypeOptions"
              :key="item.value"
              :label="item.label + '(' + item.value + ')'"
              :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="CSS Class" prop="cssClass">
          <el-input v-model="systemDictItemFrom.cssClass" placeholder="请输入 CSS Class" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="systemDictItemFrom.remark" type="textarea" :disabled="disabled" />
        </el-form-item>
      </el-form>
      <template #footer v-if="!disabled">
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemDictItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSystemDictItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemDict">
import { ref, reactive, onMounted } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, View, DArrowRight, Remove } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemDict } from "@/api/interface/systemDict";
import {
  getSystemDictListApi,
  deleteSystemDictApi,
  getSystemDictApi,
  addSystemDictApi,
  updateSystemDictApi
} from "@/api/modules/systemDict";
import { FormInstance, FormRules } from "element-plus";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import { useRoute } from "vue-router";
import { useTabsStore } from "@/stores/modules/tabs";
import { useKeepAliveStore } from "@/stores/modules/keepAlive";
import { getSystemDictTypeListSimpleApi } from "@/api/modules/systemDictType";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { SystemDictType } from "@/api/interface/systemDictType";
const route = useRoute();
const tabStore = useTabsStore();
const keepAliveStore = useKeepAliveStore();
const dictTypeId = ref(Number(route.params.dictTypeId));
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
const systemDictItemFrom = ref<SystemDict.ResSystemDictItem>({
  id: 0, // 字典编码
  sort: 0, // 字典排序
  label: "", // 字典标签
  value: "", // 字典键值
  dictTypeId: dictTypeId.value, // 字段类型 ID;
  status: 0, // 状态（0正常 1停用）
  colorType: undefined, // 颜色类型
  cssClass: undefined, // css 样式
  remark: undefined, // 备注
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined // 更新时间
});
//校验
const refSystemDictItemFrom = ref<FormInstance>();
//校验
const rulesSystemDictItemFrom = reactive<FormRules>({
  sort: [{ required: true, message: "请输入字典排序", trigger: "blur" }],
  label: [{ required: true, message: "请输入字典标签", trigger: "blur" }],
  value: [{ required: true, message: "请输入字典键值", trigger: "blur" }],
  status: [{ required: true, message: "请选择状态", trigger: "blur" }],
  dictTypeId: [{ required: true, message: "请选择字典类型", trigger: "blur" }]
});
//菜单状态
const statusEnum = getIntDictOptions("status");

// 数据标签回显样式
const colorTypeOptions = ref([
  {
    value: "default",
    label: "默认"
  },
  {
    value: "primary",
    label: "主要"
  },
  {
    value: "success",
    label: "成功"
  },
  {
    value: "info",
    label: "信息"
  },
  {
    value: "warning",
    label: "警告"
  },
  {
    value: "danger",
    label: "危险"
  }
]);

// 字典类型枚举
const dictTypeList = ref<SystemDictType.ResSystemDictTypeItem[]>([]);
const columns: ColumnProps<SystemDict.ResSystemDictItem>[] = [
  { prop: "id", label: "编号", width: 100, fixed: "left" },
  { prop: "label", label: "数据标签" },
  { prop: "value", label: "数据键值" },
  {
    prop: "dictTypeId",
    label: "字典类型",
    enum: dictTypeList,
    fieldNames: { label: "name", value: "id" },
    search: { el: "select", span: 2, defaultValue: dictTypeId.value }
  },
  { prop: "colorType", label: "颜色类型", tag: true, enum: colorTypeOptions, width: 100 },
  { prop: "cssClass", label: "CSS Class" },
  { prop: "sort", label: "字典排序" },
  { prop: "remark", label: "备注" },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 }, width: 100 },
  { prop: "creator", label: "创建者" },
  { prop: "createTime", label: "创建时间" },
  { prop: "updater", label: "更新者" },
  { prop: "updateTime", label: "更新时间" },
  {
    prop: "operation",
    label: "操作",
    width: 160,
    fixed: "right",
    isShow: HasPermission("dict.SystemDictDelete", "dict.SystemDictUpdate", "dict.SystemDict")
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  systemDictItemFrom.value = {
    id: 0, // 字典编码
    sort: 0, // 字典排序
    label: "", // 字典标签
    value: "", // 字典键值
    dictTypeId: dictTypeId.value, // 字段类型 ID;
    status: 0, // 状态（0正常 1停用）
    colorType: undefined, // 颜色类型
    cssClass: undefined, // css 样式
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
    const data = systemDictItemFrom.value as unknown as SystemDict.ResSystemDictItem;
    if (data.id !== 0) {
      await useHandleSet(updateCustomSystemDictApi, data.id, data, "修改字典数据");
    } else {
      await useHandleData(addCustomSystemDictApi, data, "添加字典数据");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};
// 删除按钮
const handleDelete = async (row: SystemDict.ResSystemDictItem) => {
  await useHandleData(deleteCustomSystemDictApi, Number(row.id), "删除字典数据");
  proTable.value?.getTableList();
};
// 添加按钮
const handleAdd = () => {
  title.value = "新增字典数据";
  centerDialogVisible.value = true;
  reset();
  disabled.value = false;
};
// 编辑按钮
const handleUpdate = async (row: SystemDict.ResSystemDictItem) => {
  title.value = "编辑字典数据";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getCustomSystemDictApi(Number(row.id));
  systemDictItemFrom.value = data;
  disabled.value = false;
};
// 查看按钮
const handleItem = async (row: SystemDict.ResSystemDictItem) => {
  title.value = "查看字典数据";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getCustomSystemDictApi(Number(row.id));
  systemDictItemFrom.value = data;
  disabled.value = true;
};

// getCustomSystemDictListApi 自定义列表方法
const getCustomSystemDictListApi = (params: any) => {
  const newParams = JSON.parse(JSON.stringify(params));
  if (newParams.dictTypeId) {
    dictTypeId.value = Number(newParams.dictTypeId);
    systemDictItemFrom.value.dictTypeId = dictTypeId.value;
    delete newParams.dictTypeId;
  }
  return getSystemDictListApi(Number(dictTypeId.value), params);
};

// getCustomSystemDictApi 自定义详情方法
const getCustomSystemDictApi = (id: Number) => {
  return getSystemDictApi(Number(dictTypeId.value), Number(id));
};
// deleteCustomSystemDictApi 自定义删除方法
const deleteCustomSystemDictApi = (id: Number) => {
  return deleteSystemDictApi(Number(dictTypeId.value), Number(id));
};
// updateCustomSystemDictApi 自定义更新方法
const updateCustomSystemDictApi = (id: Number, data: SystemDict.ResSystemDictItem) => {
  return updateSystemDictApi(Number(dictTypeId.value), Number(id), data);
};

// addCustomSystemDictApi 自定义添加方法
const addCustomSystemDictApi = (data: SystemDict.ResSystemDictItem) => {
  return addSystemDictApi(Number(dictTypeId.value), data);
};
// 关闭当前页
const closeCurrentTab = () => {
  if (route.meta.isAffix) return;
  tabStore.removeTabs(route.fullPath);
  keepAliveStore.removeKeepAliveName(route.name as string);
};

onMounted(async () => {
  const { data } = await getSystemDictTypeListSimpleApi();
  dictTypeList.value = data;
});
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
