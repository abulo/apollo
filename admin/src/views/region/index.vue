<template>
  <div class="table-box">
    <VirtualTable
      ref="proTable"
      title="地区列表"
      row-key="id"
      :columns="columns"
      :request-api="getRegionListApi"
      :request-auto="true"
      :pagination="false"
      :tree-config="{ transform: true, iconOpen: 'vxe-icon-arrow-down', iconClose: 'vxe-icon-arrow-right' }"
      :scroll-y="{ enabled: true }"
      :init-param="initParam"
      height="600"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'region.RegionCreate'" type="primary" :icon="CirclePlus" @click="handleAdd()">新增</el-button>
        <el-button type="primary" :icon="Sort" @click="toggleExpandAll">展开/折叠</el-button>
      </template>
      <!-- 状态-->
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 地区操作 -->
      <template #operation="scope">
        <el-button v-auth="'region.RegionCreate'" type="primary" link :icon="CirclePlus" @click="handleAdd(scope.row)">
          新增
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['region.RegionUpdate', 'region.RegionDelete', 'region.RegionRecover']"
            type="primary"
            link
            :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'region.RegionUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                编辑
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 0"
                v-auth="'region.RegionDelete'"
                :icon="Delete"
                @click="handleDelete(scope.row)">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'region.RegionRecover'"
                :icon="Refresh"
                @click="handleRecover(scope.row)">
                恢复
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </template>
    </VirtualTable>
    <!-- 新增/编辑弹窗 -->
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
      <el-form-item label="地区编号" prop="id">
        <el-input v-model="regionItemFrom.id" placeholder="请输入地区编号" />
      </el-form-item>
      <el-form ref="refRegionItemFrom" :model="regionItemFrom" :rules="rulesRegionItemFrom" label-width="100px">
        <el-form-item label="上级地区" prop="parentId">
          <el-tree-select
            v-model="regionItemFrom.parentId"
            :data="regionOptions"
            :props="{ value: 'id', label: 'name' }"
            value-key="id"
            node-key="id"
            placeholder="请选择"
            check-strictly
            :render-after-expand="false" />
        </el-form-item>
        <el-form-item label="地区名称" prop="name">
          <el-input v-model="regionItemFrom.name" placeholder="请输入地区名称" />
        </el-form-item>
        <el-form-item label="显示排序" prop="sort">
          <el-input-number v-model="regionItemFrom.sort" controls-position="right" :min="0" />
        </el-form-item>
        <el-form-item label="地区状态" prop="status">
          <el-radio-group v-model="regionItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refRegionItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refRegionItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx" name="region">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/VirtualTable/interface";
import { EditPen, CirclePlus, Delete, Sort, Refresh, DArrowRight } from "@element-plus/icons-vue";
import VirtualTable from "@/components/VirtualTable/index.vue";
import { Region } from "@/api/interface/region";
import {
  getRegionListApi,
  getRegionItemApi,
  addRegionApi,
  updateRegionApi,
  deleteRegionApi,
  recoverRegionApi
} from "@/api/modules/region";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
const initParam = reactive({ tree: 0 });

//加载
const loading = ref(false);
//table数据
const proTable = ref<ProTableInstance>();
//是否展开，默认全部折叠
const isExpandAll = ref(false);
//弹出层标题
const title = ref();
//是否显示弹出层
const centerDialogVisible = ref(false);
//地区树选项
const regionOptions = ref<Region.ResRegionList[]>([]);
//数据接口
const regionItemFrom = ref<Region.ResRegionItem>({
  id: 0, //地区ID
  name: "", //地区名称
  sort: 0, //显示顺序
  parentId: 0, //父地区ID
  status: 0, //地区状态(0开启/1关闭)
  deleted: 0 //是否删除
});
//校验
const refRegionItemFrom = ref<FormInstance>();
//校验
const rulesRegionItemFrom = reactive<FormRules>({
  id: [{ required: true, message: "请输入地区编号", trigger: "blur" }],
  name: [{ required: true, message: "请输入地区名称", trigger: "blur" }]
});

// 重置数据
const reset = () => {
  regionItemFrom.value = {
    id: 0, //地区ID
    name: "", //地区名称
    sort: 0, //显示顺序
    parentId: 0, //父地区ID
    status: 0, //地区状态(0开启/1关闭)
    deleted: 0 //是否删除
  };
};

// resetForm
const resetForm = (formEl: FormInstance | undefined) => {
  centerDialogVisible.value = false;
  if (!formEl) return;
  formEl.resetFields();
};

//地区状态
const statusEnum = getIntDictOptions("status");
//地区删除状态
const deletedEnum = getIntDictOptions("delete");

// 设置展开合并
const toggleExpandAll = () => {
  isExpandAll.value = !isExpandAll.value;
  proTable.value?.element?.setAllTreeExpand(isExpandAll.value);
};

// 添加按钮
const handleAdd = (row?: Region.ResRegionItem) => {
  title.value = "新增地区";
  centerDialogVisible.value = true;
  reset();
  getTreeSelect();
  if (row != null && row.id) {
    regionItemFrom.value.parentId = row.id;
  } else {
    regionItemFrom.value.parentId = 0;
  }
};

// 编辑按钮
const handleUpdate = async (row: Region.ResRegionItem) => {
  title.value = "编辑地区";
  centerDialogVisible.value = true;
  reset();
  getTreeSelect();
  const { data } = await getRegionItemApi(Number(row.id));
  regionItemFrom.value = data;
};

// 删除按钮
const handleDelete = async (row: Region.ResRegionItem) => {
  await useHandleData(deleteRegionApi, Number(row.id), "删除地区");
  proTable.value?.getTableList();
};

// 恢复按钮
const handleRecover = async (row: Region.ResRegionItem) => {
  await useHandleData(recoverRegionApi, Number(row.id), "恢复地区");
  proTable.value?.getTableList();
};

// 获取地区树选项
const getTreeSelect = async () => {
  regionOptions.value = [];
  const res = await getRegionListApi();
  let obj: Region.ResRegionList = {
    id: 0,
    name: "主类目",
    children: res.data,
    sort: 0,
    parentId: 0,
    status: 0,
    deleted: 0
  };
  regionOptions.value.push(obj);
};

// 提交数据
const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loading.value = true;
    const data = regionItemFrom.value as unknown as Region.ResRegionItem;
    if (data.id !== 0) {
      await useHandleSet(updateRegionApi, data.id, data, "修改地区");
    } else {
      await useHandleData(addRegionApi, data, "添加地区");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("region.RegionDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);

const columns: ColumnProps<Region.ResRegionList>[] = [
  { field: "id", title: "地区编号", width: 100 },
  { field: "name", title: "地区名称", align: "left", treeNode: true },
  { field: "sort", title: "排序", width: 100 },
  { field: "status", title: "状态", width: 100, tag: true, enum: statusEnum, search: { el: "select", span: 2 } },
  {
    field: "deleted",
    title: "删除",
    tag: true,
    enum: deletedEnum,
    search: deleteSearch,
    width: 100
  },
  {
    field: "operation",
    title: "操作",
    width: 160,
    fixed: "right",
    showOverflow: true,
    visible: HasPermission("region.RegionUpdate", "region.RegionDelete", "region.RegionRecover", "region.RegionCreate")
  }
];
</script>
<style lang="scss">
@import "@/styles/custom.scss";
</style>
