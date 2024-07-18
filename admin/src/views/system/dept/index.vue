<template>
  <div class="table-box">
    <VirtualTable
      ref="proTable"
      title="部门列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemDeptListApi"
      :data-callback="getFlatList"
      :request-auto="true"
      :pagination="false"
      :column-config="{ resizable: true }"
      :row-config="{ height: 45, isHover: true, keyField: 'id', useKey: true }"
      :tree-config="{ transform: true, iconOpen: 'vxe-icon-arrow-down', iconClose: 'vxe-icon-arrow-right', reserve: true }"
      :scroll-y="{ enabled: true }"
      height="600"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'dept.SystemDeptCreate'" type="primary" :icon="CirclePlus" @click="handleAdd()">新增</el-button>
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
      <template #operation="scope">
        <el-button v-auth="'dept.SystemDeptCreate'" type="primary" link :icon="CirclePlus" @click="handleAdd(scope.row)">
          新增
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['dept.SystemDeptUpdate', 'dept.SystemDeptDelete', 'dept.SystemDeptRecover']"
            type="primary"
            link
            :icon="DArrowRight"
            >更多</el-button
          >
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'dept.SystemDeptUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                编辑
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 0"
                v-auth="'dept.SystemDeptDelete'"
                :icon="Delete"
                @click="handleDelete(scope.row)">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 1"
                v-auth="'dept.SystemDeptRecover'"
                :icon="Refresh"
                @click="handleRecover(scope.row)">
                恢复
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </template>
    </VirtualTable>
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
      class="dialog-settings"
      @click="handleDialogClick">
      <el-form ref="refSystemDeptItemFrom" :model="systemDeptItemFrom" :rules="rulesSystemDeptItemFrom" label-width="100px">
        <el-form-item label="上级部门" prop="parentId">
          <el-tree-select
            v-model="systemDeptItemFrom.parentId"
            :data="deptSelect"
            :props="{ value: 'id', label: 'name' }"
            value-key="id"
            node-key="id"
            placeholder="请选择"
            default-expand-all
            check-strictly
            :render-after-expand="false" />
        </el-form-item>
        <el-form-item label="部门名称" prop="name">
          <el-input v-model="systemDeptItemFrom.name" placeholder="请输入部门名称" />
        </el-form-item>
        <el-form-item label="负责人" prop="userId">
          <el-popover placement="bottom-start" :width="600" :show-arrow="false" trigger="click" :visible="isUserOpenPopover">
            <template #reference>
              <el-button style="margin-right: 16px" @click.stop="userOpenPopover">{{ userItem }}</el-button>
            </template>
            <div class="table-box">
              <ProTable
                title="用户列表"
                row-key="id"
                :columns="userColumns"
                :request-api="getSystemUserSearch"
                :request-auto="true"
                :tool-button="false"
                :pagination-layout="'prev, pager, next'"
                :padding="'20px 0 20px 0'"
                :init-param="initParam"
                :pagination="true">
                <template #operation="scope">
                  <el-button type="primary" link :icon="CirclePlus" @click.stop="handleUser(scope.row)">确定</el-button>
                </template>
              </ProTable>
            </div>
          </el-popover>
        </el-form-item>
        <el-form-item label="显示排序" prop="sort">
          <el-input-number v-model="systemDeptItemFrom.sort" controls-position="right" :min="0" />
        </el-form-item>
        <el-form-item label="联系电话" prop="phone">
          <el-input v-model="systemDeptItemFrom.phone" maxlength="11" placeholder="请输入联系电话" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="systemDeptItemFrom.email" maxlength="50" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="systemDeptItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click.stop="resetForm(refSystemDeptItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click.stop="submitForm(refSystemDeptItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="tsx" name="systemDept">
import { ref, reactive } from "vue";
import {
  ProTableInstance as VirtualProTableInstance,
  ColumnProps as VirtualColumnProps,
  SearchProps as VirtualSearchProps
} from "@/components/VirtualTable/interface";
import { EditPen, CirclePlus, Sort, Delete, Refresh, DArrowRight } from "@element-plus/icons-vue";
import VirtualTable from "@/components/VirtualTable/index.vue";
import ProTable from "@/components/ProTable/index.vue";
import { ColumnProps } from "@/components/ProTable/interface";
import { SystemUser } from "@/api/interface/systemUser";
import { getSystemUserItemApi, getSystemUserListSimpleApi } from "@/api/modules/systemUser";
import { SystemDept } from "@/api/interface/systemDept";
import {
  getSystemDeptListApi,
  getSystemDeptItemApi,
  addSystemDeptApi,
  updateSystemDeptApi,
  deleteSystemDeptApi,
  recoverSystemDeptApi
} from "@/api/modules/systemDept";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";

//菜单状态
const statusEnum = getIntDictOptions("status");
//菜单删除状态
const deletedEnum = getIntDictOptions("delete");
//加载
const loading = ref(false);
//table数据
const proTable = ref<VirtualProTableInstance>();
//是否展开，默认全部折叠
const isExpandAll = ref(true);
//弹出层标题
const title = ref();
//是否显示弹出层
const centerDialogVisible = ref(false);
//部门树选项
const deptSelect = ref<SystemDept.ResSystemDeptList[]>([]);
//数据接口
const systemDeptItemFrom = ref<SystemDept.ResSystemDeptItem>({
  id: 0, //部门ID
  name: "", //部门名称
  parentId: 0, //父部门ID
  sort: 0, //显示顺序
  userId: 0, //负责人id
  phone: "", //联系电话
  email: "", //邮箱
  status: 0, //部门状态（0正常 1停用）
  deleted: 0 //是否删除
});
//校验
const refSystemDeptItemFrom = ref<FormInstance>();
//校验
const rulesSystemDeptItemFrom = reactive<FormRules>({
  parentId: [{ required: true, message: "上级部门不能为空", trigger: "blur" }],
  name: [{ required: true, message: "部门名称不能为空", trigger: "blur" }],
  sort: [{ required: true, message: "显示排序不能为空", trigger: "blur" }],
  email: [{ type: "email", message: "请输入正确的邮箱地址", trigger: ["blur", "change"] }],
  phone: [{ pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/, message: "请输入正确的手机号码", trigger: "blur" }],
  status: [{ required: true, message: "状态不能为空", trigger: "blur" }]
});

// 重置数据
const reset = () => {
  loading.value = false;
  systemDeptItemFrom.value = {
    id: 0, //部门ID
    name: "", //部门名称
    parentId: 0, //父部门ID
    sort: 0, //显示顺序
    userId: 0, //负责人id
    phone: "", //联系电话
    email: "", //邮箱
    status: 0, //部门状态（0正常 1停用）
    deleted: 0 //是否删除
  };
  userItem.value = "点击选择";
  isUserOpenPopover.value = false;
};

const getFlatList = (list: SystemDept.ResSystemDeptList[]) => {
  let newList: SystemDept.ResSystemDeptList[] = JSON.parse(JSON.stringify(list));
  return newList.flatMap(item => [item, ...(item.children ? getFlatList(item.children) : [])]);
};

// 设置展开合并
const toggleExpandAll = () => {
  isExpandAll.value = !isExpandAll.value;
  proTable.value?.element?.setAllTreeExpand(isExpandAll.value);
};
// resetForm
const resetForm = (formEl: FormInstance | undefined) => {
  centerDialogVisible.value = false;
  if (!formEl) return;
  formEl.resetFields();
};
// 添加按钮
const handleAdd = (row?: SystemDept.ResSystemDeptItem) => {
  title.value = "新增部门";
  centerDialogVisible.value = true;
  reset();
  getTreeSelect();
  if (row != null && row.id) {
    systemDeptItemFrom.value.parentId = row.id;
  } else {
    systemDeptItemFrom.value.parentId = 0;
  }
};

// 获取部门树选项
const getTreeSelect = async () => {
  deptSelect.value = [];
  const { data } = await getSystemDeptListApi();
  let obj: SystemDept.ResSystemDeptList = {
    id: 0,
    name: "顶级部门",
    children: data,
    parentId: 0, //父部门ID
    sort: 0, //显示顺序
    userId: 0, //负责人id
    phone: "", //联系电话
    email: "", //邮箱
    status: 0, //部门状态（0正常 1停用）
    deleted: 0 //是否删除
  };
  deptSelect.value.push(obj);
};

// 编辑按钮
const handleUpdate = async (row: SystemDept.ResSystemDeptItem) => {
  title.value = "编辑部门";
  centerDialogVisible.value = true;
  reset();
  getTreeSelect();
  const { data } = await getSystemDeptItemApi(Number(row.id));
  systemDeptItemFrom.value = data;
  if (Number(data.userId) !== 0) {
    const user = await getSystemUserItemApi(Number(data.userId));
    userItem.value = user.data.nickname;
  }
};

// 删除按钮
const handleDelete = async (row: SystemDept.ResSystemDeptItem) => {
  await useHandleData(deleteSystemDeptApi, Number(row.id), "删除部门");
  reloadData();
};

// 恢复按钮
const handleRecover = async (row: SystemDept.ResSystemDeptItem) => {
  await useHandleData(recoverSystemDeptApi, Number(row.id), "恢复部门");
  reloadData();
};

// 提交数据
const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loading.value = true;
    const data = systemDeptItemFrom.value as unknown as SystemDept.ResSystemDeptItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemDeptApi, data.id, data, "修改部门");
    } else {
      await useHandleData(addSystemDeptApi, data, "添加部门");
    }
    resetForm(formEl);
    loading.value = false;
    reloadData();
  });
};
// 表格配置项
const deleteSearch = reactive<VirtualSearchProps>(
  HasPermission("dept.SystemDeptDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);
// 表格配置项
const columns: VirtualColumnProps<SystemDept.ResSystemDeptList>[] = [
  { field: "id", title: "编号", width: 100, fixed: "left" },
  { field: "name", title: "部门名称", align: "left", treeNode: true },
  { field: "sort", title: "排序" },
  { field: "status", title: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 } },
  {
    field: "deleted",
    title: "删除",
    tag: true,
    enum: deletedEnum,
    search: deleteSearch
  },
  {
    field: "operation",
    title: "操作",
    width: 160,
    fixed: "right",
    isShow: HasPermission("dept.SystemDeptUpdate", "dept.SystemDeptDelete", "dept.SystemDeptRecover", "dept.SystemDeptCreate")
  }
];

// 定义负责人
const userItem = ref("点击选择");
const isUserOpenPopover = ref(false);
const initParam = reactive({ pageSize: 5 });
const userOpenPopover = () => {
  isUserOpenPopover.value = true;
};
// 定义用户数据
const userColumns: ColumnProps<SystemUser.ResSystemUserItem>[] = [
  { prop: "id", label: "编码", fixed: "left" },
  { prop: "username", label: "用户名", search: { el: "input", span: 2, props: { placeholder: "请输入用户名/昵称/手机号" } } },
  { prop: "nickname", label: "用户名称" },
  {
    prop: "operation",
    label: "操作",
    fixed: "right"
  }
];

// 在 el-dialog 上添加点击事件监听器
const handleDialogClick = () => {
  isUserOpenPopover.value = false;
};
// 列表数据接口
const getSystemUserSearch = (params: any) => {
  const newParams = Object.assign(params, JSON.parse(JSON.stringify(initParam)));
  return getSystemUserListSimpleApi(newParams);
};
// 当用户被选择
const handleUser = (row: SystemUser.ResSystemUserItem) => {
  userItem.value = row.nickname;
  systemDeptItemFrom.value.userId = Number(row.id);
  isUserOpenPopover.value = false;
};

const reloadData = async () => {
  // const recordList = proTable.value?.element?.getTreeExpandRecords();
  const { data } = await getSystemDeptListApi();
  proTable.value?.element?.loadData(getFlatList(data));
};
</script>

<style lang="scss">
@import "@/styles/custom";
@import "@/styles/popover";
</style>
