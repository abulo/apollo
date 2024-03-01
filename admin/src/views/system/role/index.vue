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
        <el-button type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <!-- 状态-->
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 角色类型 -->
      <template #type="scope">
        <DictTag type="role.type" :value="scope.row.type" />
      </template>
      <template #operation="scope">
        <el-button type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)"> 编辑 </el-button>
        <el-button type="primary" link :icon="Delete" @click="handleDelete(scope.row)"> 删除 </el-button>
      </template>
    </ProTable>
    <!-- 添加或者修改 -->
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
      <el-form ref="refSystemRoleItemFrom" :model="systemRoleItemFrom" :rules="rulesSystemRoleItemFrom" label-width="100px">
        <el-form-item label="角色名称" prop="name">
          <el-input v-model="systemRoleItemFrom.name" />
        </el-form-item>
        <el-form-item label="角色编码" prop="code">
          <el-input v-model="systemRoleItemFrom.code" placeholder="请输入角色编码" />
        </el-form-item>
        <el-form-item label="角色类型" prop="type">
          <el-radio-group v-model="systemRoleItemFrom.type">
            <el-radio-button v-for="dict in typeEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
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
        <el-form-item label="菜单权限">
          <el-card class="cardHeight">
            <template #header>
              全选/全不选:
              <el-switch
                v-model="menuNodeAll"
                active-text="是"
                inactive-text="否"
                inline-prompt
                @change="handleMenuCheckedTreeNodeAll" />
              全部展开/折叠:
              <el-switch
                v-model="menuExpand"
                active-text="展开"
                inactive-text="折叠"
                inline-prompt
                @change="handleMenuCheckedTreeExpand" />
            </template>
            <el-tree
              ref="menuRef"
              :data="menuOptions"
              :props="defaultProps"
              :list="systemRoleItemFrom.menuIds"
              empty-text="加载中，请稍候"
              node-key="id"
              show-checkbox />
          </el-card>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="systemRoleItemFrom.remark" placeholder="请输备注" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemRoleItemFrom)">取消</el-button>
          <el-button type="primary" @click="submitForm(refSystemRoleItemFrom)" :loading="loading">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemRole">
import { ref, reactive } from "vue";
import { useTimeoutFn } from "@vueuse/core";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { DictTag } from "@/components/DictTag";
import { FormInstance, FormRules, ElTree } from "element-plus";
import type Node from "element-plus/es/components/tree/src/model/node";
import { SystemRole } from "@/api/interface/systemRole";
import {
  getSystemRoleListApi,
  deleteSystemRoleApi,
  getSystemRoleItemApi,
  addSystemRoleApi,
  updateSystemRoleApi
} from "@/api/modules/systemRole";
import { SystemMenu } from "@/api/interface/systemMenu";
import { getSystemMenuListApi } from "@/api/modules/systemMenu";
import { getIntDictOptions } from "@/utils/dict";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
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
//部门树选项
const menuOptions = ref<SystemMenu.ResSystemMenuList[]>([]);
const menuNodeAll = ref(false); // 全选/全不选
const menuRef = ref<InstanceType<typeof ElTree>>();
const menuExpand = ref(false); // 展开/折叠
//数据接口
const systemRoleItemFrom = ref<SystemRole.ResSystemRoleItem>({
  id: 0, //bigint 角色ID,PRI
  name: "", //varchar 角色名称
  code: "", //varchar 角色权限字符串
  sort: 0, //int 显示顺序
  status: 0, //tinyint 角色状态（0正常 1停用）
  type: 2, //tinyint 角色类型(1内置/2定义)
  remark: "", //varchar 备注
  creator: "", //varchar 创建者
  createTime: "", //datetime 创建时间
  updater: "", //varchar 更新者
  updateTime: "", //datetime 更新时间
  menuIds: [] //array 角色菜单
});
// 校验
const refSystemRoleItemFrom = ref<FormInstance>();
const rulesSystemRoleItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "角色标题不能为空", trigger: "blur" }],
  code: [{ required: true, message: "角色编码不能为空", trigger: "change" }],
  sort: [{ required: true, message: "角色顺序不能为空", trigger: "change" }],
  status: [{ required: true, message: "角色状态不能为空", trigger: "change" }],
  remark: [{ required: false, message: "角色内容不能为空", trigger: "blur" }],
  menuIds: [{ required: true, message: "必须选择", trigger: "change" }]
});

// 状态
const statusEnum = getIntDictOptions("status");
// 角色类型
const typeEnum = getIntDictOptions("role.type");

const columns: ColumnProps<SystemRole.ResSystemRoleItem>[] = [
  { prop: "id", label: "编号", width: 100 },
  { prop: "name", label: "角色名称" },
  { prop: "code", label: "角色编码" },
  { prop: "type", label: "角色类型", tag: true, enum: typeEnum, search: { el: "select", span: 2 }, width: 100 },
  { prop: "sort", label: "角色顺序" },
  { prop: "remark", label: "角色备注" },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 }, width: 100 },
  { prop: "operation", label: "操作", width: 160, fixed: "right" }
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
    creator: "", //varchar 创建者
    createTime: "", //datetime 创建时间
    updater: "", //varchar 更新者
    updateTime: "", //datetime 更新时间
    menuIds: [] //array 角色菜单
  };
  menuNodeAll.value = false;
  menuExpand.value = false;
  menuRef.value?.setCheckedNodes([]);
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
  getMenuTreeSelect();
  reset();
};

// 编辑按钮
const handleUpdate = async (row: SystemRole.ResSystemRoleItem) => {
  title.value = "编辑角色";
  centerDialogVisible.value = true;
  reset();
  getMenuTreeSelect();
  const { data } = await getSystemRoleItemApi(Number(row.id));
  systemRoleItemFrom.value = data;
  const menuIds = data.menuIds;
  useTimeoutFn(() => {
    menuIds?.forEach((menuId: number) => {
      menuRef.value?.setChecked(menuId, true, false);
    });
  }, 200);
};

// 获取菜单树选项
const getMenuTreeSelect = async () => {
  menuOptions.value = [];
  const { data } = await getSystemMenuListApi();
  menuOptions.value = data;
};

// 全选/全不选
const handleMenuCheckedTreeNodeAll = () => {
  let data = menuNodeAll.value ? menuOptions.value : [];
  menuRef.value!.setCheckedNodes(data as unknown as Node[]);
};

// 展开/折叠全部
const handleMenuCheckedTreeExpand = () => {
  const nodes = menuRef.value?.store.nodesMap;
  for (let node in nodes) {
    if (nodes[node].expanded === menuExpand.value) {
      continue;
    }
    nodes[node].expanded = menuExpand.value;
  }
};
</script>

<style lang="scss" scoped>
.cardHeight {
  width: 100%;
  max-height: 400px;
  overflow-y: auto;
}
</style>
