<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="字典数据"
      row-key="id"
      :columns="columns"
      :request-api="getCustomSystemDictListApi"
      :request-auto="true"
      :pagination="false"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
        <el-button type="primary" :icon="Remove" @click="closeCurrentTab">关闭</el-button>
      </template>
      <!-- 状态-->
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)"> 编辑 </el-button>
        <el-button type="primary" link :icon="Delete" @click="handleDelete(scope.row)"> 删除 </el-button>
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
      <el-form ref="refSystemDictItemFrom" :model="systemDictItemFrom" :rules="rulesSystemDictItemFrom" label-width="100px">
        <el-form-item label="数据标签" prop="label">
          <el-input v-model="systemDictItemFrom.label" />
        </el-form-item>
        <el-form-item label="数据键值" prop="value">
          <el-input v-model="systemDictItemFrom.value" />
        </el-form-item>
        <el-form-item label="编码排序" prop="sort">
          <el-input-number v-model="systemDictItemFrom.sort" controls-position="right" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="systemDictItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="颜色类型" prop="colorType">
          <el-select v-model="systemDictItemFrom.colorType" clearable>
            <el-option
              v-for="item in colorTypeOptions"
              :key="item.value"
              :label="item.label + '(' + item.value + ')'"
              :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="CSS Class" prop="cssClass">
          <el-input v-model="systemDictItemFrom.cssClass" placeholder="请输入 CSS Class" />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input type="textarea" v-model="systemDictItemFrom.remark" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemDictItemFrom)">取消</el-button>
          <el-button type="primary" @click="submitForm(refSystemDictItemFrom)" :loading="loading">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemDict">
import { useRoute } from "vue-router";
import { onMounted, reactive, ref } from "vue";
import { CirclePlus, EditPen, Delete, Remove } from "@element-plus/icons-vue";
import { FormInstance, FormRules } from "element-plus";
import { useTabsStore } from "@/stores/modules/tabs";
import { useKeepAliveStore } from "@/stores/modules/keepAlive";
import ProTable from "@/components/ProTable/index.vue";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import { SystemDict } from "@/api/interface/systemDict";
import {
  getSystemDictListApi,
  addSystemDictApi,
  updateSystemDictApi,
  getSystemDictItemApi,
  deleteSystemDictApi
} from "@/api/modules/systemDict";
import { getAllSystemDictTypeApi } from "@/api/modules/systemDictType";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
const route = useRoute();
const tabStore = useTabsStore();
const keepAliveStore = useKeepAliveStore();
// 字典 systemDictId
const systemDictType = ref(String(route.params.type));
//加载
const loading = ref(false);
//弹出层标题
const title = ref();
//是否显示弹出层
const centerDialogVisible = ref(false);
//字典数据
const systemDictItemFrom = ref<SystemDict.ResSystemDictItem>({
  id: 0, //字典主键
  sort: 0, //字典排序
  label: "", //字典标签
  value: "", //字典键值
  dictType: systemDictType.value, //字典类型
  status: 0, //状态（0正常 1停用）
  colorType: "", //颜色类型
  cssClass: "", //css 样式
  remark: "", //备注
  creator: "", //创建人
  createTime: "", //创建时间
  updater: "", //更新人
  updateTime: "" //更新时间
});
//table数据
const proTable = ref<ProTableInstance>();
//校验
const refSystemDictItemFrom = ref<FormInstance>();
//校验
const rulesSystemDictItemFrom = reactive<FormRules>({
  sort: [{ required: true, message: "请输入字典排序", trigger: "blur" }],
  label: [{ required: true, message: "请输入字典标签", trigger: "blur" }],
  value: [{ required: true, message: "请输入字典键值", trigger: "blur" }],
  status: [{ required: true, message: "请选择状态", trigger: "blur" }],
  dictType: [{ required: true, message: "请选择字典类型", trigger: "blur" }]
});
// 关闭当前页
const closeCurrentTab = () => {
  if (route.meta.isAffix) return;
  tabStore.removeTabs(route.fullPath);
  keepAliveStore.removeKeepAliveName(route.name as string);
};

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
const dictTypeEnum = ref([]);
// 表头信息
const columns: ColumnProps<SystemDict.ResSystemDictItem>[] = [
  { prop: "id", label: "编号", width: 100 },
  { prop: "label", label: "数据标签" },
  { prop: "value", label: "数据键值" },
  {
    prop: "dictType",
    label: "字典类型",
    enum: dictTypeEnum,
    search: { el: "tree-select", span: 2, defaultValue: systemDictType.value }
  },
  { prop: "colorType", label: "颜色类型", tag: true, enum: colorTypeOptions, width: 100 },
  { prop: "cssClass", label: "CSS Class" },
  { prop: "sort", label: "字典排序" },
  { prop: "remark", label: "备注" },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 }, width: 100 },
  { prop: "operation", label: "操作", width: 160, fixed: "right" }
];

// 重置数据
const reset = () => {
  systemDictItemFrom.value = {
    id: 0, //字典主键
    sort: 0, //字典排序
    label: "", //字典标签
    value: "", //字典键值
    dictType: systemDictType.value, //字典类型
    status: 0, //状态（0正常 1停用）
    colorType: "", //颜色类型
    cssClass: "", //css 样式
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

// 添加按钮
const handleAdd = () => {
  title.value = "新增字典数据";
  centerDialogVisible.value = true;
  reset();
};

// 编辑按钮
const handleUpdate = async (row: SystemDict.ResSystemDictItem) => {
  title.value = "编辑字典数据";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemDictItemApi(Number(row.id));
  systemDictItemFrom.value = data;
};

// 删除按钮
const handleDelete = async (row: SystemDict.ResSystemDictItem) => {
  await useHandleData(deleteSystemDictApi, row.id, "删除字典");
  proTable.value?.getTableList();
};

// 提交数据
const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loading.value = true;
    const data = systemDictItemFrom.value as unknown as SystemDict.ResSystemDictItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemDictApi, data.id, data, "修改字典");
    } else {
      await useHandleData(addSystemDictApi, data, "新增字典");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};
// 数据初始化时候需要去获取字典类型数据
onMounted(async () => {
  const { data } = await getAllSystemDictTypeApi();
  for (const item of data) {
    (dictTypeEnum.value as Array<{ label: string; value: string }>).push({ label: item.name, value: item.type });
  }
});

const getCustomSystemDictListApi = (params: any) => {
  const newParams = JSON.parse(JSON.stringify(params));
  //判断 newParams 中是否存在 dictType, 如果存在, 就将值赋值给 systemDictType
  if (newParams.dictType) {
    systemDictType.value = String(newParams.dictType);
    systemDictItemFrom.value.dictType = systemDictType.value;
  }
  return getSystemDictListApi(params);
};
</script>
<style lang="scss">
@import "@/styles/custom.scss";
</style>
