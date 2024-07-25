<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="角色列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemRoleListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'role.SystemRoleCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'role.SystemRole'" type="primary" link :icon="View" @click="handleItem(scope.row)"> 查看 </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['role.SystemRoleUpdate', 'role.SystemRoleDelete', 'role.SystemRoleRecover', 'role.SystemRoleDrop']"
            type="primary"
            link
            :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'role.SystemRoleUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                编辑
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 0"
                v-auth="'role.SystemRoleDelete'"
                :icon="Delete"
                @click="handleDelete(scope.row)">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'role.SystemRoleRecover'"
                :icon="Refresh"
                @click="handleRecover(scope.row)">
                恢复
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'role.SystemRoleDrop'"
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
      <el-form ref="refSystemRoleItemFrom" :model="systemRoleItemFrom" :rules="rulesSystemRoleItemFrom" label-width="100px">
        <el-form-item label="角色名称" prop="name">
          <el-input v-model="systemRoleItemFrom.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="角色编码" prop="code">
          <el-input v-model="systemRoleItemFrom.code" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="显示顺序" prop="sort">
          <el-input-number v-model="systemRoleItemFrom.sort" controls-position="right" :min="0" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="systemRoleItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value" :disabled="disabled">
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
              :list="systemRoleItemFrom.menuIds"
              empty-text="加载中，请稍候"
              node-key="id"
              show-checkbox />
          </el-card>
        </el-form-item>
        <el-form-item label="权限类别" prop="dataScope">
          <el-select v-model="systemRoleItemFrom.dataScope" :disabled="disabled">
            <el-option v-for="dict in dataScopeEnum" :key="Number(dict.value)" :label="dict.label" :value="Number(dict.value)" />
          </el-select>
        </el-form-item>
        <el-form-item
          v-model="systemRoleItemFrom.dataScopeDept"
          v-if="systemRoleItemFrom.dataScope === 2"
          label="权限范围"
          prop="dataScopeDept">
          <el-card class="cardHeight">
            <template #header>
              全选/不选:
              <el-switch v-model="deptNodeAll" active-text="是" inactive-text="否" inline-prompt @change="handleDeptNodeAll" />
              展开/折叠:
              <el-switch v-model="deptExpand" active-text="展开" inactive-text="折叠" inline-prompt @change="handleDeptExpand" />
              父子联动:
              <el-switch v-model="checkStrictly" active-text="是" inactive-text="否" inline-prompt />
            </template>
            <el-tree
              ref="deptRef"
              :data="deptList"
              :props="deptProps"
              :list="systemRoleItemFrom.dataScopeDept"
              empty-text="加载中，请稍候"
              :disabled="disabled"
              :check-strictly="!checkStrictly"
              node-key="id"
              show-checkbox />
          </el-card>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="systemRoleItemFrom.remark" :disabled="disabled" />
        </el-form-item>
      </el-form>
      <template #footer v-if="!disabled">
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemRoleItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSystemRoleItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemRole">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh, DeleteFilled, View, DArrowRight } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemRole } from "@/api/interface/systemRole";
import {
  getSystemRoleListApi,
  deleteSystemRoleApi,
  dropSystemRoleApi,
  recoverSystemRoleApi,
  getSystemRoleApi,
  addSystemRoleApi,
  updateSystemRoleApi
} from "@/api/modules/systemRole";
import { FormInstance, FormRules, ElTree, ElMessage } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import { SystemMenu } from "@/api/interface/systemMenu";
import { SystemDept } from "@/api/interface/systemDept";
import { getSystemTenantMenuListApi } from "@/api/modules/systemTenant";
import { getSystemDeptListSimpleApi } from "@/api/modules/systemDept";
import { useTimeoutFn } from "@vueuse/core";
import Node from "element-plus/es/components/tree/src/model/node";
const deptProps = {
  children: "children",
  label: "name",
  value: "id"
};

const deptRef = ref<InstanceType<typeof ElTree>>();
const deptList = ref<SystemDept.ResSystemDeptItem[]>([]);
const deptNodeAll = ref(false); // 全选/全不选
const deptExpand = ref(false); // 展开/折叠
const checkStrictly = ref(true); // 是否严格模式，即父子不关联

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
const systemRoleItemFrom = ref<SystemRole.ResSystemRoleItem>({
  id: 0, // 角色编号
  name: "", // 角色名称
  code: "", // 角色权限字符串
  sort: 0, // 显示顺序
  dataScope: undefined, // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
  dataScopeDept: undefined, // 数据范围(指定部门数组)
  status: 0, // 角色状态（0正常 1停用）
  type: 0, // 角色类型(1内置/2定义)
  remark: undefined, // 备注
  deleted: 0, // 删除
  tenantId: 0, // 租户
  creator: undefined, // 创建者
  createTime: undefined, // 创建时间
  updater: undefined, // 更新者
  updateTime: undefined, // 更新时间
  menuIds: []
});
//校验
const refSystemRoleItemFrom = ref<FormInstance>();
//校验
const rulesSystemRoleItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "角色标题不能为空", trigger: "blur" }],
  code: [{ required: true, message: "角色编码不能为空", trigger: "change" }],
  sort: [{ required: true, message: "角色顺序不能为空", trigger: "change" }],
  status: [{ required: true, message: "角色状态不能为空", trigger: "change" }],
  dataScope: [{ required: true, message: "必须选择", trigger: "change" }],
  menuIds: [{ required: true, message: "必须选择", trigger: "change" }]
});
// 状态
const statusEnum = getIntDictOptions("status");
// 数据权限
const dataScopeEnum = getIntDictOptions("role.scope");
// 删除状态
const deletedEnum = getIntDictOptions("delete");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("role.SystemRoleDelete")
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

const columns: ColumnProps<SystemRole.ResSystemRoleItem>[] = [
  { prop: "id", label: "编号", width: 100, fixed: "left" },
  { prop: "name", label: "角色名称", search: { el: "input", span: 2, props: { placeholder: "请输入名称" } } },
  { prop: "code", label: "角色编码" },
  { prop: "sort", label: "角色顺序" },
  { prop: "remark", label: "角色备注" },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 }, width: 100 },
  {
    prop: "deleted",
    label: "删除",
    tag: true,
    enum: deletedEnum,
    search: deleteSearch,
    width: 100
  },
  { prop: "creator", label: "创建者" },
  { prop: "createTime", label: "创建时间" },
  { prop: "updater", label: "更新者" },
  { prop: "updateTime", label: "更新时间" },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission(
      "role.SystemRoleUpdate",
      "role.SystemRoleDelete",
      "role.SystemRoleDrop",
      "role.SystemRoleRecover",
      "role.SystemRole"
    )
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  systemRoleItemFrom.value = {
    id: 0, // 角色编号
    name: "", // 角色名称
    code: "", // 角色权限字符串
    sort: 0, // 显示顺序
    dataScope: undefined, // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
    dataScopeDept: undefined, // 数据范围(指定部门数组)
    status: 0, // 角色状态（0正常 1停用）
    type: 0, // 角色类型(1内置/2定义)
    remark: undefined, // 备注
    deleted: 0, // 删除
    tenantId: 0, // 租户
    creator: undefined, // 创建者
    createTime: undefined, // 创建时间
    updater: undefined, // 更新者
    updateTime: undefined, // 更新时间
    menuIds: []
  };
  disabled.value = true;
  menuNodeAll.value = false;
  menuExpand.value = false;
  deptNodeAll.value = false;
  deptExpand.value = false;
  checkStrictly.value = true;
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
    const data = systemRoleItemFrom.value as unknown as SystemRole.ResSystemRoleItem;
    data.menuIds = [
      ...(menuRef.value!.getCheckedKeys(false) as unknown as Array<number>), // 获得当前选中节点
      ...(menuRef.value!.getHalfCheckedKeys() as unknown as Array<number>) // 获得半选中的父节点
    ];
    if (data.dataScope === 2) {
      data.dataScopeDept = [
        ...(deptRef.value!.getCheckedKeys(false) as unknown as Array<number>), // 获得当前选中节点
        ...(deptRef.value!.getHalfCheckedKeys() as unknown as Array<number>) // 获得半选中的父节点
      ];
      // 判断这个数据的长度
      if (data.dataScopeDept.length === 0) {
        loading.value = false;
        ElMessage({
          type: "warning",
          message: "请选择权限范围"
        });
        return;
      }
    }
    if (data.id !== 0) {
      await useHandleSet(updateSystemRoleApi, data.id, data, "修改角色");
    } else {
      await useHandleData(addSystemRoleApi, data, "添加角色");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};
// 清理按钮
const handleDrop = async (row: SystemRole.ResSystemRoleItem) => {
  await useHandleData(dropSystemRoleApi, Number(row.id), "清理角色");
  proTable.value?.getTableList();
};
// 删除按钮
const handleDelete = async (row: SystemRole.ResSystemRoleItem) => {
  await useHandleData(deleteSystemRoleApi, Number(row.id), "删除角色");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: SystemRole.ResSystemRoleItem) => {
  await useHandleData(recoverSystemRoleApi, Number(row.id), "恢复角色");
  proTable.value?.getTableList();
};
// 添加按钮
const handleAdd = () => {
  title.value = "新增角色";
  centerDialogVisible.value = true;
  reset();
  getMenuList();
  getDeptList();
  disabled.value = false;
};
// 编辑按钮
const handleUpdate = async (row: SystemRole.ResSystemRoleItem) => {
  title.value = "编辑角色";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemRoleApi(Number(row.id));
  systemRoleItemFrom.value = data;
  getMenuList();
  getDeptList();
  useTimeoutFn(() => {
    data?.dataScopeDept?.forEach((deptId: number) => {
      deptRef.value?.setChecked(deptId, true, false);
    });
    data?.menuIds?.forEach((menuId: number) => {
      menuRef.value?.setChecked(menuId, true, false);
    });
  }, 200);
  disabled.value = false;
};
// 查看按钮
const handleItem = async (row: SystemRole.ResSystemRoleItem) => {
  title.value = "查看角色";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemRoleApi(Number(row.id));
  systemRoleItemFrom.value = data;
  getMenuList();
  getDeptList();
  useTimeoutFn(() => {
    data?.dataScopeDept?.forEach((deptId: number) => {
      deptRef.value?.setChecked(deptId, true, false);
    });
    data?.menuIds?.forEach((menuId: number) => {
      menuRef.value?.setChecked(menuId, true, false);
    });
  }, 200);
  disabled.value = true;
};

const getMenuList = async () => {
  menuList.value = [];
  const { data } = await getSystemTenantMenuListApi();
  menuList.value = data;
};

const getDeptList = async () => {
  deptList.value = [];
  const { data } = await getSystemDeptListSimpleApi();
  deptList.value = data;
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

// 全选/全不选
const handleDeptNodeAll = () => {
  let data = deptNodeAll.value ? deptList.value : [];
  deptRef.value!.setCheckedNodes(data as unknown as Node[]);
};

// 展开/折叠全部
const handleDeptExpand = () => {
  const nodes = deptRef.value?.store.nodesMap;
  for (let node in nodes) {
    if (nodes[node].expanded === deptExpand.value) {
      continue;
    }
    nodes[node].expanded = deptExpand.value;
  }
};
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
