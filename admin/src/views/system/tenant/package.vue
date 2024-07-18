<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="套餐列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemTenantPackageListApi"
      :request-auto="true"
      :pagination="false"
      :search-col="12">
      <template #tableHeader>
        <el-button v-auth="'tenant.SystemTenantPackageCreate'" type="primary" :icon="CirclePlus" @click="handleAdd"
          >新增</el-button
        >
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
        <el-button
          v-auth="'tenant.SystemTenantPackageUpdate'"
          type="primary"
          link
          :icon="EditPen"
          @click="handleUpdate(scope.row)">
          编辑
        </el-button>
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'tenant.SystemTenantPackageDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'tenant.SystemTenantPackageRecover'"
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
        ref="refSystemTenantPackageItemFrom"
        :model="systemTenantPackageItemFrom"
        :rules="rulesSystemTenantPackageItemFrom"
        :loading="loading"
        label-width="100px">
        <el-form-item label="套餐名称" prop="name">
          <el-input v-model="systemTenantPackageItemFrom.name" />
        </el-form-item>
        <el-form-item label="菜单权限">
          <el-card class="cardHeight">
            <template #header>
              全选/全不选:
              <el-switch v-model="menuNodeAll" active-text="是" inactive-text="否" inline-prompt @change="handleMenuNodeAll" />
              全部展开/折叠:
              <el-switch v-model="menuExpand" active-text="展开" inactive-text="折叠" inline-prompt @change="handleMenuExpand" />
            </template>
            <el-tree
              ref="menuRef"
              :data="menuSelect"
              :props="defaultProps"
              :list="systemTenantPackageItemFrom.menuIds"
              empty-text="加载中，请稍候"
              node-key="id"
              show-checkbox />
          </el-card>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="systemTenantPackageItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :label="Number(dict.value)">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="systemTenantPackageItemFrom.remark" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemTenantPackageItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSystemTenantPackageItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts" name="systemTenantPackage">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemTenantPackage } from "@/api/interface/systemTenantPackage";
import {
  getSystemTenantPackageListApi,
  deleteSystemTenantPackageApi,
  recoverSystemTenantPackageApi,
  getSystemTenantPackageItemApi,
  addSystemTenantPackageApi,
  updateSystemTenantPackageApi
} from "@/api/modules/systemTenantPackage";
import { SystemMenu } from "@/api/interface/systemMenu";
import { getSystemMenuListSimpleApi } from "@/api/modules/systemMenu";
import { FormInstance, FormRules, ElTree } from "element-plus";
import Node from "element-plus/es/components/tree/src/model/node";
import { useTimeoutFn } from "@vueuse/core";
import { getIntDictOptions } from "@/utils/dict";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { DictTag } from "@/components/DictTag";
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
const systemTenantPackageItemFrom = ref<SystemTenantPackage.ResSystemTenantPackageItem>({
  id: 0, //套餐编号,PRI
  name: "", //套餐名
  status: 0, //租户状态（0正常 1停用）
  remark: "", //备注
  menuIds: [], //关联的菜单编号
  deleted: 0 //是否删除
});
//校验
const refSystemTenantPackageItemFrom = ref<FormInstance>();
//校验
const rulesSystemTenantPackageItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "请输入套餐名称", trigger: "blur" }],
  status: [{ required: true, message: "请选择状态", trigger: "blur" }]
});

// 默认参数
const defaultProps = {
  children: "children",
  label: "name",
  value: "id"
};

const menuNodeAll = ref(false); // 全选/全不选
const menuRef = ref<InstanceType<typeof ElTree>>();
const menuSelect = ref<SystemMenu.ResSystemMenuList[]>([]);
const menuExpand = ref(false); // 展开/折叠

// 状态
const statusEnum = getIntDictOptions("status");
const deletedEnum = getIntDictOptions("delete");

// 表格配置项
// const deleteSearch = reactive<SearchProps>(
//   HasPermission("tenant.SystemPackageDelete")
//     ? {
//         el: "switch",
//         span: 2
//       }
//     : {}
// );

// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("tenant.SystemTenantPackageDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);
// 配置数据列表
const columns: ColumnProps<SystemTenantPackage.ResSystemTenantPackageItem>[] = [
  { prop: "id", label: "编号", width: 100, fixed: "left" },
  { prop: "name", label: "套餐名称", search: { el: "input", span: 2 } },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 }, width: 100 },
  { prop: "remark", label: "备注" },
  {
    prop: "deleted",
    label: "删除",
    tag: true,
    enum: deletedEnum,
    search: deleteSearch,
    width: 100
  },
  {
    prop: "operation",
    label: "操作",
    width: 160,
    fixed: "right",
    isShow: HasPermission(
      "tenant.SystemTenantPackageRecover",
      "tenant.SystemTenantPackageDelete",
      "tenant.SystemTenantPackageUpdate"
    )
  }
];

// 获取菜单树选项
const getTreeSelect = async () => {
  menuSelect.value = [];
  const { data } = await getSystemMenuListSimpleApi();
  menuSelect.value = data;
};

/** 全选/全不选 */
const handleMenuNodeAll = () => {
  let data = menuNodeAll.value ? menuSelect.value : [];
  menuRef.value!.setCheckedNodes(data as unknown as Node[]);
};

/** 展开/折叠全部 */
const handleMenuExpand = () => {
  const nodes = menuRef.value?.store.nodesMap;
  for (let node in nodes) {
    if (nodes[node].expanded === menuExpand.value) {
      continue;
    }
    nodes[node].expanded = menuExpand.value;
  }
};

// 重置数据
const reset = () => {
  loading.value = false;
  // 重置选项
  menuNodeAll.value = false;
  menuExpand.value = false;
  systemTenantPackageItemFrom.value = {
    id: 0, //套餐编号,PRI
    name: "", //套餐名
    status: 0, //租户状态（0正常 1停用）
    remark: "", //备注
    menuIds: [], //关联的菜单编号
    deleted: 0 //是否删除
  };
  menuRef.value?.setCheckedNodes([]);
  refSystemTenantPackageItemFrom.value?.resetFields();
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
    const data = systemTenantPackageItemFrom.value as unknown as SystemTenantPackage.ResSystemTenantPackageItem;
    data.menuIds = [
      ...(menuRef.value!.getCheckedKeys(false) as unknown as Array<number>), // 获得当前选中节点
      ...(menuRef.value!.getHalfCheckedKeys() as unknown as Array<number>) // 获得半选中的父节点
    ];
    if (data.id !== 0) {
      await useHandleSet(updateSystemTenantPackageApi, data.id, data, "修改套餐");
    } else {
      await useHandleData(addSystemTenantPackageApi, data, "添加套餐");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

// 添加按钮
const handleAdd = () => {
  title.value = "新增套餐";
  centerDialogVisible.value = true;
  reset();
  getTreeSelect();
};

// 编辑按钮
const handleUpdate = async (row: SystemTenantPackage.ResSystemTenantPackageItem) => {
  reset();
  getTreeSelect();
  title.value = "编辑套餐";
  centerDialogVisible.value = true;
  const { data } = await getSystemTenantPackageItemApi(Number(row.id));
  systemTenantPackageItemFrom.value = data;
  useTimeoutFn(() => {
    data.menuIds?.forEach((menuId: number) => {
      menuRef.value?.setChecked(menuId, true, false);
    });
  }, 200);
};

// 删除按钮
const handleDelete = async (row: SystemTenantPackage.ResSystemTenantPackageItem) => {
  await useHandleData(deleteSystemTenantPackageApi, Number(row.id), "删除套餐");
  proTable.value?.getTableList();
};

const handleRecover = async (row: SystemTenantPackage.ResSystemTenantPackageItem) => {
  await useHandleData(recoverSystemTenantPackageApi, Number(row.id), "恢复套餐");
  proTable.value?.getTableList();
};
</script>

<style lang="scss">
@import "@/styles/custom";
</style>
