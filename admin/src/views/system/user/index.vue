<template>
  <div class="main-box">
    <TreeFilter label="name" :data="deptList" :default-value="initParam.deptId" @change="changeDept" />
    <div class="table-box">
      <ProTable
        ref="proTable"
        title="用户列表"
        row-key="id"
        :columns="columns"
        :request-api="getSystemUserListApi"
        :init-param="initParam"
        :request-auto="true"
        :pagination="true"
        :search-col="12">
        <!-- 表格 header 按钮 -->
        <template #tableHeader>
          <el-button v-auth="'user.SystemUserCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
        </template>
        <template #status="scope">
          <DictTag type="status" :value="scope.row.status" />
        </template>
        <template #deleted="scope">
          <DictTag type="delete" :value="scope.row.deleted" />
        </template>
        <template #postIds="scope">
          <p class="order-font" v-for="postId in scope.row.postIds" :key="postId">
            <el-tag type="info">
              {{ postSelect && postSelect.find(item => item.id === Number(postId))?.name }}
            </el-tag>
          </p>
        </template>
        <template #roleIds="scope">
          <p class="order-font" v-for="roleId in scope.row.roleIds" :key="roleId">
            <el-tag type="info">
              {{ roleSelect && roleSelect.find(item => item.id === Number(roleId))?.name }}
            </el-tag>
          </p>
        </template>
        <template #deptIds="scope">
          <p class="order-font" v-for="deptId in scope.row.deptIds" :key="deptId">
            <el-tag type="info">
              {{ deptTree && getFlatList(deptTree).find(item => item.id === Number(deptId))?.name }}
            </el-tag>
          </p>
        </template>
        <!-- 菜单操作 -->
        <template #operation="scope">
          <el-button v-auth="'user.SystemUserUpdate'" type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)">
            编辑
          </el-button>
          <el-dropdown trigger="click">
            <el-button
              v-auth="['user.SystemUserPassword', 'user.SystemUserDelete', 'user.SystemUserRecover']"
              type="primary"
              link
              :icon="DArrowRight">
              更多
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item v-auth="'user.SystemUserPassword'" :icon="Key" @click="handlePassword(scope.row)">
                  重置密码
                </el-dropdown-item>
                <el-dropdown-item
                  v-if="scope.row.deleted === 0"
                  v-auth="'user.SystemUserDelete'"
                  :icon="Delete"
                  @click="handleDelete(scope.row)">
                  删除
                </el-dropdown-item>
                <el-dropdown-item
                  v-if="scope.row.deleted === 1"
                  v-auth="'user.SystemUserRecover'"
                  :icon="Refresh"
                  @click="handleRecover(scope.row)">
                  恢复
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
        <el-form ref="refSystemUserItemFrom" :model="systemUserItemFrom" :rules="rulesSystemUserItemFrom" label-width="100px">
          <el-form-item label="用户昵称" prop="nickname">
            <el-input v-model="systemUserItemFrom.nickname" />
          </el-form-item>
          <el-form-item label="用户手机" prop="mobile">
            <el-input v-model="systemUserItemFrom.mobile" />
          </el-form-item>
          <el-form-item v-if="systemUserItemFrom.id === undefined" label="用户名" prop="username">
            <el-input v-model="systemUserItemFrom.username" />
          </el-form-item>
          <el-form-item v-if="systemUserItemFrom.id === undefined" label="用户密码" prop="password">
            <el-input v-model="systemUserItemFrom.password" show-password type="password" />
          </el-form-item>
          <el-form-item label="用户状态" prop="status">
            <el-radio-group v-model="systemUserItemFrom.status">
              <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value">
                {{ dict.label }}
              </el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="角色" prop="roleIds">
            <el-select v-model="systemUserItemFrom.roleIds" multiple placeholder="请选择角色">
              <el-option v-for="item in roleSelect" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item>
          <el-form-item label="职位" prop="postIds">
            <el-select v-model="systemUserItemFrom.postIds" multiple placeholder="请选择职位">
              <el-option v-for="item in postSelect" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item>

          <el-form-item label="部门" prop="deptIds" v-model="systemUserItemFrom.deptIds">
            <el-card class="cardHeight">
              <template #header>
                全选/不选:
                <el-switch v-model="deptNodeAll" active-text="是" inactive-text="否" inline-prompt @change="handleDataNodeAll" />
                展开/折叠:
                <el-switch
                  v-model="deptExpand"
                  active-text="展开"
                  inactive-text="折叠"
                  inline-prompt
                  @change="handleDataExpand" />
                关联/不关联:
                <el-switch
                  v-model="deptStrictly"
                  active-text="关联"
                  inactive-text="不关联"
                  inline-prompt
                  @change="handleDataStrictly" />
              </template>
              <el-tree
                ref="deptRef"
                :data="deptSelect"
                :props="defaultProps"
                :list="systemUserItemFrom.deptIds"
                empty-text="加载中，请稍候"
                node-key="id"
                :check-strictly="deptCheckStrictly"
                show-checkbox />
            </el-card>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="resetForm(refSystemUserItemFrom)">取消</el-button>
            <el-button type="primary" :loading="loading" @click="submitForm(refSystemUserItemFrom)">确定</el-button>
          </span>
        </template>
      </el-dialog>
    </div>
  </div>
</template>
<script setup lang="ts" name="systemUser">
import { onMounted, ref, reactive } from "vue";
import { useTimeoutFn } from "@vueuse/core";
import { ColumnProps, ProTableInstance, SearchProps } from "@/components/ProTable/interface";
import { DictTag } from "@/components/DictTag";
import { EditPen, CirclePlus, Delete, DArrowRight, Refresh, CircleCheck, Key } from "@element-plus/icons-vue";
import { FormInstance, FormRules, ElMessage, ElMessageBox, ElTree } from "element-plus";
import ProTable from "@/components/ProTable/index.vue";
import { SystemUser } from "@/api/interface/systemUser";
import {
  getSystemUserListApi,
  deleteSystemUserApi,
  getSystemUserItemApi,
  addSystemUserApi,
  updateSystemUserApi,
  recoverSystemUserApi,
  putSystemUserPasswordApi
} from "@/api/modules/systemUser";
import { getIntDictOptions } from "@/utils/dict";
import TreeFilter from "@/components/TreeFilter/index.vue";
import { SystemRole } from "@/api/interface/systemRole";
import { getSystemRoleListSimpleApi } from "@/api/modules/systemRole";
import { SystemDept } from "@/api/interface/systemDept";
import { getSystemDeptListSimpleApi, getSystemDeptListLabelApi } from "@/api/modules/systemDept";
import { SystemPost } from "@/api/interface/systemPost";
import { getSystemPostListSimpleApi } from "@/api/modules/systemPost";
import { HasPermission } from "@/utils/permission";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import Node from "element-plus/es/components/tree/src/model/node";

const initParam = reactive({ deptId: "" });
//加载
const loading = ref(false);
//弹出层标题
const title = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
// 状态
const statusEnum = getIntDictOptions("status");
const deletedEnum = getIntDictOptions("delete");
const userDeptId = ref<Number[]>([]);

//部门树选项
const deptOptions = ref<SystemDept.ResSystemDeptList[]>([]);
const deptNodeAll = ref(false); // 全选/全不选
const deptRef = ref<InstanceType<typeof ElTree>>();
const deptExpand = ref(false); // 展开/折叠
const deptStrictly = ref(true); // 关联/不关联
const deptCheckStrictly = ref(true);
const defaultProps = {
  children: "children",
  label: "name",
  value: "id"
};

//数据接口
const systemUserItemFrom = ref<SystemUser.ResSystemUserItem>({
  id: undefined,
  nickname: "",
  mobile: "",
  username: "",
  password: "",
  status: 0,
  deleted: 0,
  roleIds: [],
  deptIds: [],
  postIds: []
});
//校验
const refSystemUserItemFrom = ref<FormInstance>();
//校验
const rulesSystemUserItemFrom = reactive<FormRules>({
  username: [{ required: true, message: "用户名称不能为空", trigger: "blur" }],
  nickname: [{ required: true, message: "用户昵称不能为空", trigger: "blur" }],
  password: [{ required: true, message: "用户密码不能为空", trigger: "blur" }],
  mobile: [{ required: true, message: "用户手机不能为空", trigger: "blur" }],
  status: [{ required: true, message: "用户状态不能为空", trigger: "blur" }],
  deptIds: [{ required: true, message: "用户部门不能为空", trigger: "blur" }],
  roleIds: [{ required: true, message: "用户角色不能为空", trigger: "blur" }],
  postIds: [{ required: true, message: "用户职位不能为空", trigger: "blur" }]
});

const deptTree = ref<SystemDept.ResSystemDeptList[]>([]);
//部门树选项
const deptSelect = ref<SystemDept.ResSystemDeptList[]>([]);
// 职位列表
const postSelect = ref<SystemPost.ResSystemPostItem[]>([]);
// 角色列表
const roleSelect = ref<SystemRole.ResSystemRoleItem[]>([]);

// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("user.SystemUserDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);
// 定义列表
const columns: ColumnProps<SystemUser.ResSystemUserItem>[] = [
  { prop: "id", label: "编号", fixed: "left" },
  { prop: "username", label: "用户名", search: { el: "input", span: 2 } },
  { prop: "nickname", label: "用户昵称" },
  { prop: "mobile", label: "用户手机" },
  { prop: "deptIds", label: "部门" },
  { prop: "postIds", label: "职位" },
  { prop: "roleIds", label: "角色" },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 } },
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
    isShow: HasPermission("user.SystemUserPassword", "user.SystemUserDelete", "user.SystemUserUpdate", "user.SystemUserRecover")
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  systemUserItemFrom.value = {
    id: undefined,
    nickname: "",
    mobile: "",
    username: "",
    password: "",
    status: 0,
    deleted: 0,
    roleIds: [],
    deptIds: [],
    postIds: []
  };
  deptNodeAll.value = false;
  deptExpand.value = false;
  deptRef.value?.setCheckedNodes([]);
  deptStrictly.value = true;
  deptCheckStrictly.value = true;
  userDeptId.value = [];
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
// 关联/不关联
const handleDataStrictly = () => {
  // let data = [];
  deptCheckStrictly.value = !deptStrictly.value;
  // deptRef.value!.setCheckedNodes(data as unknown as Node[]);
};

// 添加按钮
const handleAdd = () => {
  reset();
  title.value = "新增用户";
  centerDialogVisible.value = true;
};

const getFlatList = (list: SystemDept.ResSystemDeptList[]) => {
  let newList: SystemDept.ResSystemDeptList[] = JSON.parse(JSON.stringify(list));
  return newList.flatMap(item => [item, ...(item.children ? getFlatList(item.children) : [])]);
};

// 编辑按钮
const handleUpdate = async (row: SystemUser.ResSystemUserItem) => {
  reset();
  title.value = "编辑用户";
  centerDialogVisible.value = true;
  const { data } = await getSystemUserItemApi(Number(row.id));
  systemUserItemFrom.value = data;
  const deptListData = await getSystemDeptListSimpleApi();
  deptSelect.value = deptListData.data;
  const postList = await getSystemPostListSimpleApi();
  postSelect.value = postList.data;
  const roleList = await getSystemRoleListSimpleApi();
  roleSelect.value = roleList.data;

  const dataScopeDept = data.deptIds;
  useTimeoutFn(() => {
    dataScopeDept?.forEach((deptId: number) => {
      deptRef.value?.setChecked(deptId, true, false);
    });
  }, 200);

  const currUserDept = data.deptIds as unknown as Number[];
  const curDept = getFlatList(deptListData.data) as SystemDept.ResSystemDeptList[];
  // 遍历 curDept
  curDept.forEach(item => {
    if (!currUserDept.includes(item.id)) {
      userDeptId.value.push(item.id);
    }
  });
};

// 删除按钮
const handleDelete = async (row: SystemUser.ResSystemUserItem) => {
  await useHandleData(deleteSystemUserApi, Number(row.id), "删除用户");
  proTable.value?.getTableList();
};

const handleRecover = async (row: SystemUser.ResSystemUserItem) => {
  await useHandleData(recoverSystemUserApi, Number(row.id), "恢复用户");
  proTable.value?.getTableList();
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
    const data = systemUserItemFrom.value as unknown as SystemUser.ResSystemUserItem;
    data.deptIds = [];
    data.deptIds = [
      ...(deptRef.value!.getCheckedKeys(false) as unknown as Array<number>), // 获得当前选中节点
      ...(deptRef.value!.getHalfCheckedKeys() as unknown as Array<number>) // 获得半选中的父节点
    ];
    // 判断userDeptId有没有值
    if (userDeptId.value.length > 0) {
      // 有值的话 需要将 数组添加到 data.deptIds
      data.deptIds = [...(data.deptIds as unknown as Array<number>), ...(userDeptId.value as unknown as Array<number>)];
    }
    if (data.id !== undefined) {
      delete data.password;
      await useHandleSet(updateSystemUserApi, data.id, data, "修改用户");
    } else {
      await useHandleData(addSystemUserApi, data, "添加用户");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};
// 部门树选择
const deptList = ref<SystemDept.ResSystemDeptList[]>([]);
const getTreeFilter = async () => {
  const { data } = await getSystemDeptListSimpleApi();
  deptList.value = data;
  const deptListData = await getSystemDeptListSimpleApi();
  deptSelect.value = deptListData.data;
  const postList = await getSystemPostListSimpleApi();
  postSelect.value = postList.data;
  const roleList = await getSystemRoleListSimpleApi();
  roleSelect.value = roleList.data;
  const deptTreeData = await getSystemDeptListLabelApi();
  deptTree.value = deptTreeData.data;
};

// 树形筛选切换
const changeDept = (val: string) => {
  proTable.value!.pageable.pageNum = 1;
  initParam.deptId = val;
  proTable.value?.getTableList();
};

//设置密码
const systemUserItemPassword = ref<SystemUser.ReqSystemUserPassword>({
  password: ""
});

// handlePassword 修改密码
const handlePassword = (row: SystemUser.ResSystemUserItem) => {
  ElMessageBox.prompt("请输入 " + row.username + " 的新密码", "温馨提示", {
    confirmButtonText: "确认",
    cancelButtonText: "取消"
  })
    .then(({ value }) => {
      if (value === "") {
        ElMessage.error({ message: "密码不能为空" });
        return;
      }
      const data = systemUserItemPassword.value as unknown as SystemUser.ReqSystemUserPassword;
      putSystemUserPasswordApi(Number(row.id), data).then(() => {
        ElMessage.success({ message: `${row.nickname}密码修改成功！` });
      });
    })
    .catch(() => {
      ElMessage.info({ message: "已取消" });
    });
};

onMounted(() => {
  getTreeFilter();
});
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
