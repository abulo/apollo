<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="角色列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemRoleListApi"
      :request-auto="true"
      :pagination="false"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'role.SystemRoleCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <!-- 状态-->
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <template #operation="scope">
        <el-button v-auth="'role.SystemRoleUpdate'" type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)">
          编辑
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['role.SystemRoleMenuList', 'role.SystemRoleDataScope', 'role.SystemRoleRecover', 'role.SystemRoleDelete']"
            type="primary"
            link
            :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'role.SystemRoleMenuList'" :icon="Menu" @click="handleMenu(scope.row)">
                菜单权限
              </el-dropdown-item>
              <el-dropdown-item v-auth="'role.SystemRoleDataScope'" :icon="DataBoard" @click="handleScope(scope.row)">
                数据权限
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
                @click="handleDelete(scope.row)">
                恢复
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </template>
    </ProTable>
    <!-- 添加或者修改 -->
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
          <el-input v-model="systemRoleItemFrom.name" />
        </el-form-item>
        <el-form-item label="角色编码" prop="code">
          <el-input v-model="systemRoleItemFrom.code" placeholder="请输入角色编码" />
        </el-form-item>
        <el-form-item label="角色顺序" prop="sort">
          <el-input-number v-model="systemRoleItemFrom.sort" controls-position="right" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="systemRoleItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="备注" prop="remark">
          <el-input v-model="systemRoleItemFrom.remark" placeholder="请输备注" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemRoleItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSystemRoleItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
    <!-- 菜单权限 -->
    <el-dialog
      v-model="centerMenuDialogVisible"
      :title="titleMenu"
      width="40%"
      destroy-on-close
      align-center
      center
      append-to-body
      draggable
      :lock-scroll="false"
      class="dialog-settings">
      <el-form
        ref="refSystemRoleMenuItemFrom"
        :model="systemRoleMenuItemFrom"
        :rules="rulesSystemRoleMenuItemFrom"
        label-width="100px">
        <el-form-item label="角色名称">
          <el-tag>{{ systemRoleItemFrom.name }}</el-tag>
        </el-form-item>
        <el-form-item label="角色编码">
          <el-tag>{{ systemRoleItemFrom.code }}</el-tag>
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
              :data="menuOptions"
              :props="defaultProps"
              :list="systemRoleMenuItemFrom.menuIds"
              empty-text="加载中，请稍候"
              node-key="id"
              show-checkbox />
          </el-card>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetMenuForm(refSystemRoleMenuItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitMenuForm(refSystemRoleMenuItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
    <!-- 数据权限 -->
    <el-dialog
      v-model="centerScopeDialogVisible"
      :title="titleScope"
      width="40%"
      destroy-on-close
      align-center
      center
      append-to-body
      draggable
      :lock-scroll="false"
      class="dialog-settings">
      <el-form
        ref="refSystemRoleScopeItemFrom"
        :model="systemRoleScopeItemFrom"
        :rules="rulesSystemRoleScopeItemFrom"
        label-width="100px">
        <el-form-item label="角色名称">
          <el-tag>{{ systemRoleItemFrom.name }}</el-tag>
        </el-form-item>
        <el-form-item label="角色编码">
          <el-tag>{{ systemRoleItemFrom.code }}</el-tag>
        </el-form-item>
        <el-form-item label="权限类别" prop="dataScope">
          <el-select v-model="systemRoleScopeItemFrom.dataScope">
            <el-option v-for="dict in dataScopeEnum" :key="Number(dict.value)" :label="dict.label" :value="Number(dict.value)" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="systemRoleScopeItemFrom.dataScope === 2" label="权限范围">
          <el-card class="cardHeight">
            <template #header>
              全选/全不选:
              <el-switch v-model="deptNodeAll" active-text="是" inactive-text="否" inline-prompt @change="handleDataNodeAll" />
              全部展开/折叠:
              <el-switch v-model="deptExpand" active-text="展开" inactive-text="折叠" inline-prompt @change="handleDataExpand" />
            </template>
            <el-tree
              ref="deptRef"
              :data="deptOptions"
              :props="defaultProps"
              :list="systemRoleScopeItemFrom.dataScopeDept"
              empty-text="加载中，请稍候"
              node-key="id"
              show-checkbox />
          </el-card>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetScopeForm(refSystemRoleScopeItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitScopeForm(refSystemRoleScopeItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemRole">
import { ref, reactive } from "vue";
import { useTimeoutFn } from "@vueuse/core";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, DataBoard, DArrowRight, Menu, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { DictTag } from "@/components/DictTag";
import { FormInstance, FormRules, ElTree } from "element-plus";
import Node from "element-plus/es/components/tree/src/model/node";
import { SystemRole } from "@/api/interface/systemRole";
import {
  getSystemRoleListApi,
  deleteSystemRoleApi,
  getSystemRoleItemApi,
  addSystemRoleApi,
  updateSystemRoleApi,
  getSystemRoleMenuListApi,
  addSystemRoleMenuApi,
  getSystemRoleScopeListApi,
  addSystemRoleScopeApi
} from "@/api/modules/systemRole";
import { SystemMenu } from "@/api/interface/systemMenu";
import { getSystemTenantMenuListApi } from "@/api/modules/systemTenant";
import { SystemDept } from "@/api/interface/systemDept";
import { getSystemDeptSearchApi } from "@/api/modules/systemDept";
import { getIntDictOptions } from "@/utils/dict";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
//加载
const loading = ref(false);
//table数据
const proTable = ref<ProTableInstance>();
//弹出层标题
const title = ref();
//是否显示弹出层
const centerDialogVisible = ref(false);
// 默认
const defaultProps = {
  children: "children",
  label: "name",
  value: "id"
};
//弹出层标题
const titleMenu = ref();
//是否显示弹出层
const centerMenuDialogVisible = ref(false);
//部门树选项
const menuOptions = ref<SystemMenu.ResSystemMenuList[]>([]);
const menuNodeAll = ref(false); // 全选/全不选
const menuRef = ref<InstanceType<typeof ElTree>>();
const menuExpand = ref(false); // 展开/折叠

//部门树选项
const deptOptions = ref<SystemDept.ResSystemDeptList[]>([]);
const deptNodeAll = ref(false); // 全选/全不选
const deptRef = ref<InstanceType<typeof ElTree>>();
const deptExpand = ref(false); // 展开/折叠

//数据接口
const systemRoleItemFrom = ref<SystemRole.ResSystemRoleItem>({
  id: 0, //bigint 角色ID,PRI
  name: "", //varchar 角色名称
  code: "", //varchar 角色权限字符串
  sort: 0, //int 显示顺序
  status: 0, //tinyint 角色状态（0正常 1停用）
  type: 2, //tinyint 角色类型(1内置/2定义)
  remark: "", //varchar 备注
  deleted: 0 //tinyint 删除标志
});
// 校验
const refSystemRoleItemFrom = ref<FormInstance>();
const rulesSystemRoleItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "角色标题不能为空", trigger: "blur" }],
  code: [{ required: true, message: "角色编码不能为空", trigger: "change" }],
  sort: [{ required: true, message: "角色顺序不能为空", trigger: "change" }],
  status: [{ required: true, message: "角色状态不能为空", trigger: "change" }],
  remark: [{ required: false, message: "角色内容不能为空", trigger: "blur" }]
});

const systemRoleMenuItemFrom = ref<SystemRole.ResSystemRoleMenuItem>({
  menuIds: [], //json 关联的菜单编号
  roleId: 0 //bigint 租户编号
});
const refSystemRoleMenuItemFrom = ref<FormInstance>();
const rulesSystemRoleMenuItemFrom = reactive<FormRules>({
  menuIds: [{ required: true, message: "必须选择", trigger: "change" }]
});

//弹出层标题
const titleScope = ref();
//是否显示弹出层
const centerScopeDialogVisible = ref(false);

const systemRoleScopeItemFrom = ref<SystemRole.ResSystemRoleScopeItem>({
  id: 0, // 角色编号
  dataScope: undefined, // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
  dataScopeDept: undefined //数据范围(指定部门数组)
});
const refSystemRoleScopeItemFrom = ref<FormInstance>();
const rulesSystemRoleScopeItemFrom = reactive<FormRules>({
  dataScope: [{ required: true, message: "必须选择", trigger: "change" }]
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
        span: 2
      }
    : {}
);
const columns: ColumnProps<SystemRole.ResSystemRoleItem>[] = [
  { prop: "id", label: "编号", width: 100 },
  { prop: "name", label: "角色名称" },
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
  {
    prop: "operation",
    label: "操作",
    width: 160,
    fixed: "right",
    isShow: HasPermission(
      "role.SystemRoleMenuList",
      "role.SystemRoleDataScope",
      "role.SystemRoleRecover",
      "role.SystemRoleDelete",
      "role.SystemRoleUpdate"
    )
  }
];
// 重置数据
const reset = () => {
  systemRoleItemFrom.value = {
    id: 0, //bigint 角色ID,PRI
    name: "", //varchar 角色名称
    code: "", //varchar 角色权限字符串
    sort: 0, //int 显示顺序
    status: 0, //tinyint 角色状态（0正常 1停用）
    type: 2, //tinyint 角色类型(1内置/2定义)
    remark: "", //varchar 备注
    deleted: 0 //tinyint 删除标志
  };
  menuNodeAll.value = false;
  menuExpand.value = false;
  menuRef.value?.setCheckedNodes([]);
  deptNodeAll.value = false;
  deptExpand.value = false;
  deptRef.value?.setCheckedNodes([]);
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

// 删除按钮
const handleDelete = async (row: SystemRole.ResSystemRoleItem) => {
  await useHandleData(deleteSystemRoleApi, Number(row.id), "删除角色");
  proTable.value?.getTableList();
};

// 添加按钮
const handleAdd = () => {
  title.value = "新增角色";
  centerDialogVisible.value = true;
  reset();
};

// 编辑按钮
const handleUpdate = async (row: SystemRole.ResSystemRoleItem) => {
  title.value = "编辑角色";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemRoleItemApi(Number(row.id));
  systemRoleItemFrom.value = data;
};

// 获取菜单树选项
const getMenuTreeSelect = async () => {
  menuOptions.value = [];
  const { data } = await getSystemTenantMenuListApi();
  menuOptions.value = data;
};

// 全选/全不选
const handleMenuNodeAll = () => {
  let data = menuNodeAll.value ? menuOptions.value : [];
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

const handleMenu = async (row: SystemRole.ResSystemRoleItem) => {
  titleMenu.value = "菜单权限";
  centerMenuDialogVisible.value = true;
  reset();
  getMenuTreeSelect();
  const { data } = await getSystemRoleMenuListApi(Number(row.id));
  systemRoleItemFrom.value = row;
  systemRoleMenuItemFrom.value = data;
  const menuIds = data.menuIds;
  useTimeoutFn(() => {
    menuIds?.forEach((menuId: number) => {
      menuRef.value?.setChecked(menuId, true, false);
    });
  }, 200);
};

// 重置
const resetMenuForm = (formEl: FormInstance | undefined) => {
  centerMenuDialogVisible.value = false;
  if (!formEl) return;
  formEl.resetFields();
};

// 提交数据
const submitMenuForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loading.value = true;
    const data = systemRoleMenuItemFrom.value as unknown as SystemRole.ResSystemRoleMenuItem;
    data.menuIds = [
      ...(menuRef.value!.getCheckedKeys(false) as unknown as Array<number>), // 获得当前选中节点
      ...(menuRef.value!.getHalfCheckedKeys() as unknown as Array<number>) // 获得半选中的父节点
    ];
    await useHandleSet(addSystemRoleMenuApi, data.roleId, data, "绑定菜单权限");
    resetMenuForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

// 获取部门树选项
const getDeptTreeSelect = async () => {
  deptOptions.value = [];
  const { data } = await getSystemDeptSearchApi({ tree: 1 });
  deptOptions.value = data;
};

const handleScope = async (row: SystemRole.ResSystemRoleItem) => {
  titleScope.value = "数据权限";
  centerScopeDialogVisible.value = true;
  reset();
  getDeptTreeSelect();
  const { data } = await getSystemRoleScopeListApi(Number(row.id));
  systemRoleItemFrom.value = row;
  systemRoleScopeItemFrom.value = data;
  const dataScopeDept = data.dataScopeDept;
  useTimeoutFn(() => {
    dataScopeDept?.forEach((deptId: number) => {
      deptRef.value?.setChecked(deptId, true, false);
    });
  }, 200);
};

/** 全选/全不选 */
const handleDataNodeAll = () => {
  let data = deptNodeAll.value ? deptOptions.value : [];
  deptRef.value!.setCheckedNodes(data as unknown as Node[]);
};

/** 展开/折叠全部 */
const handleDataExpand = () => {
  const nodes = deptRef.value?.store.nodesMap;
  for (let node in nodes) {
    if (nodes[node].expanded === deptExpand.value) {
      continue;
    }
    nodes[node].expanded = deptExpand.value;
  }
};

// resetForm
const resetScopeForm = (formEl: FormInstance | undefined) => {
  centerScopeDialogVisible.value = false;
  if (!formEl) return;
  formEl.resetFields();
};

// 提交数据
const submitScopeForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loading.value = true;
    const data = systemRoleScopeItemFrom.value as unknown as SystemRole.ResSystemRoleScopeItem;
    data.dataScopeDept = [];
    if (data.dataScope === 2) {
      data.dataScopeDept = [
        ...(deptRef.value!.getCheckedKeys(false) as unknown as Array<number>), // 获得当前选中节点
        ...(deptRef.value!.getHalfCheckedKeys() as unknown as Array<number>) // 获得半选中的父节点
      ];
    }
    await useHandleSet(addSystemRoleScopeApi, data.id, data, "绑定数据权限");
    resetScopeForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};
</script>

<style lang="scss">
@import "@/styles/custom.scss";
</style>
