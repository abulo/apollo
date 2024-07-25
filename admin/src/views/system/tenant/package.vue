<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="套餐列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemTenantPackageListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'tenant.SystemTenantPackageCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">
          新增
        </el-button>
      </template>
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'tenant.SystemTenantPackage'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="[
              'tenant.SystemTenantPackageUpdate',
              'tenant.SystemTenantPackageDelete',
              'tenant.SystemTenantPackageRecover',
              'tenant.SystemTenantPackageDrop'
            ]"
            type="primary"
            link
            :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'tenant.SystemTenantPackageUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                编辑
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 0"
                v-auth="'tenant.SystemTenantPackageDelete'"
                :icon="Delete"
                @click="handleDelete(scope.row)">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'tenant.SystemTenantPackageRecover'"
                :icon="Refresh"
                @click="handleRecover(scope.row)">
                恢复
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'tenant.SystemTenantPackageDrop'"
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
        ref="refSystemTenantPackageItemFrom"
        :model="systemTenantPackageItemFrom"
        :rules="rulesSystemTenantPackageItemFrom"
        label-width="100px">
        <el-form-item label="套餐名称" prop="name">
          <el-input v-model="systemTenantPackageItemFrom.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="systemTenantPackageItemFrom.status">
            <el-radio-button
              v-for="dict in statusEnum"
              :key="Number(dict.value)"
              :label="Number(dict.value)"
              :disabled="disabled">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="菜单权限" prop="menuIds">
          <el-card class="cardHeight">
            <template #header>
              全选/全不选:
              <el-switch v-model="menuNodeAll" active-text="是" inactive-text="否" inline-prompt @change="handleMenuNodeAll" />
              全部展开/折叠:
              <el-switch v-model="menuExpand" active-text="展开" inactive-text="折叠" inline-prompt @change="handleMenuExpand" />
            </template>
            <el-tree
              ref="menuRef"
              :data="menuList"
              :props="menuProps"
              :list="systemTenantPackageItemFrom.menuIds"
              empty-text="加载中，请稍候"
              node-key="id"
              show-checkbox />
          </el-card>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="systemTenantPackageItemFrom.remark" :disabled="disabled" />
        </el-form-item>
      </el-form>
      <template #footer v-if="!disabled">
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
import { EditPen, CirclePlus, Delete, Refresh, DeleteFilled, View, DArrowRight } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemTenantPackage } from "@/api/interface/systemTenantPackage";
import {
  getSystemTenantPackageListApi,
  deleteSystemTenantPackageApi,
  dropSystemTenantPackageApi,
  recoverSystemTenantPackageApi,
  getSystemTenantPackageApi,
  addSystemTenantPackageApi,
  updateSystemTenantPackageApi
} from "@/api/modules/systemTenantPackage";
import { FormInstance, FormRules, ElTree } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import { SystemMenu } from "@/api/interface/systemMenu";
import { getSystemTenantMenuListApi } from "@/api/modules/systemTenant";
import { useTimeoutFn } from "@vueuse/core";
import Node from "element-plus/es/components/tree/src/model/node";
import { getSystemMenuListSimpleApi } from "@/api/modules/systemMenu";

const menuProps = {
  children: "children",
  label: "name",
  value: "id"
};

const menuRef = ref<InstanceType<typeof ElTree>>();
const menuList = ref<SystemMenu.ResSystemMenuItem[]>([]);
const menuNodeAll = ref(false); // 全选/全不选
const menuExpand = ref(false); // 展开/折叠

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
const systemTenantPackageItemFrom = ref<SystemTenantPackage.ResSystemTenantPackageItem>({
  id: 0, // 套餐编号
  name: "", // 套餐名称
  status: 0, // 状态（0正常 1停用）
  menuIds: [], // 目录编号
  remark: undefined, // 备注
  deleted: 0, // 是否删除(0否 1是)
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined // 更新时间
});
//校验
const refSystemTenantPackageItemFrom = ref<FormInstance>();
//校验
const rulesSystemTenantPackageItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "套餐名称不能为空", trigger: "blur" }],
  status: [{ required: true, message: "状态不能为空", trigger: "blur" }]
});
const statusEnum = getIntDictOptions("status");
const deletedEnum = getIntDictOptions("delete");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("tenant.SystemTenantPackageDelete")
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
      "tenant.SystemTenantPackageUpdate",
      "tenant.SystemTenantPackageDelete",
      "tenant.SystemTenantPackageDrop",
      "tenant.SystemTenantPackageRecover",
      "tenant.SystemTenantPackage"
    )
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  systemTenantPackageItemFrom.value = {
    id: 0, // 套餐编号
    name: "", // 套餐名称
    status: 0, // 状态（0正常 1停用）
    menuIds: [], // 目录编号
    remark: undefined, // 备注
    deleted: 0, // 是否删除(0否 1是)
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined // 更新时间
  };
  menuNodeAll.value = false;
  menuExpand.value = false;
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
// 清理按钮
const handleDrop = async (row: SystemTenantPackage.ResSystemTenantPackageItem) => {
  await useHandleData(dropSystemTenantPackageApi, Number(row.id), "清理套餐");
  proTable.value?.getTableList();
};
// 删除按钮
const handleDelete = async (row: SystemTenantPackage.ResSystemTenantPackageItem) => {
  await useHandleData(deleteSystemTenantPackageApi, Number(row.id), "删除套餐");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: SystemTenantPackage.ResSystemTenantPackageItem) => {
  await useHandleData(recoverSystemTenantPackageApi, Number(row.id), "恢复套餐");
  proTable.value?.getTableList();
};
// 添加按钮
const handleAdd = () => {
  title.value = "新增套餐";
  centerDialogVisible.value = true;
  reset();
  getMenuList();
  disabled.value = false;
};
// 编辑按钮
const handleUpdate = async (row: SystemTenantPackage.ResSystemTenantPackageItem) => {
  title.value = "编辑套餐";
  centerDialogVisible.value = true;
  reset();
  getMenuList();
  const { data } = await getSystemTenantPackageApi(Number(row.id));
  systemTenantPackageItemFrom.value = data;
  useTimeoutFn(() => {
    data?.menuIds?.forEach((menuId: number) => {
      menuRef.value?.setChecked(menuId, true, false);
    });
  }, 200);
  disabled.value = false;
};
// 查看按钮
const handleItem = async (row: SystemTenantPackage.ResSystemTenantPackageItem) => {
  title.value = "查看套餐";
  centerDialogVisible.value = true;
  reset();
  getMenuList();
  const { data } = await getSystemTenantPackageApi(Number(row.id));
  systemTenantPackageItemFrom.value = data;
  useTimeoutFn(() => {
    data?.menuIds?.forEach((menuId: number) => {
      menuRef.value?.setChecked(menuId, true, false);
    });
  }, 200);
  disabled.value = true;
};
const getMenuList = async () => {
  menuList.value = [];
  const { data } = await getSystemMenuListSimpleApi({ deleted: 0, status: 0 });
  menuList.value = data;
};

// 全选/全不选
const handleMenuNodeAll = () => {
  let data = menuNodeAll.value ? menuList.value : [];
  menuRef.value!.setCheckedNodes(data as unknown as Node[]);
};

// 展开/折叠全部
const handleMenuExpand = () => {
  const nodes = menuRef.value?.store.nodesMap;
  for (let node in nodes) {
    if (nodes[node].expanded === menuExpand.value) {
      continue;
    }
    nodes[node].expanded = menuExpand.value;
  }
};
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
