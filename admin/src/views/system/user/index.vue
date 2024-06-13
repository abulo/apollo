<template>
  <div class="main-box">
    <TreeFilter label="name" title="部门列表" :data="deptList" :default-value="initParam.deptId" @change="changeDept" />
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
        <!-- 菜单操作 -->
        <template #operation="scope">
          <el-button v-auth="'user.SystemUserUpdate'" type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)">
            编辑
          </el-button>
          <el-dropdown trigger="click">
            <el-button
              v-auth="[
                'user.SystemUserRoleList',
                'user.SystemUserDeptList',
                'user.SystemUserPostList',
                'user.SystemUserPassword',
                'user.SystemUserDelete',
                'user.SystemUserRecover'
              ]"
              type="primary"
              link
              :icon="DArrowRight">
              更多
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item v-auth="'user.SystemUserRoleList'" :icon="CircleCheck" @click="handleRole(scope.row)">
                  分配角色
                </el-dropdown-item>
                <el-dropdown-item v-auth="'user.SystemUserDeptList'" :icon="CircleCheck" @click="handleDept(scope.row)">
                  分配部门
                </el-dropdown-item>
                <el-dropdown-item v-auth="'user.SystemUserPostList'" :icon="CircleCheck" @click="handlePost(scope.row)">
                  分配职位
                </el-dropdown-item>
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
          <!-- <el-form-item label="用户角色">
            <el-select v-model="systemUserItemFrom.roleIds" multiple placeholder="请选择角色">
              <el-option v-for="item in roleList" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item> -->
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="resetForm(refSystemUserItemFrom)">取消</el-button>
            <el-button type="primary" :loading="loading" @click="submitForm(refSystemUserItemFrom)">确定</el-button>
          </span>
        </template>
      </el-dialog>
      <el-dialog
        v-model="centerRoleDialogVisible"
        :title="titleRole"
        width="40%"
        destroy-on-close
        align-center
        center
        append-to-body
        draggable
        :lock-scroll="false"
        class="dialog-settings">
        <el-form
          ref="refSystemUserRoleItemFrom"
          :model="systemUserRoleItemFrom"
          :rules="rulesSystemUserRoleItemFrom"
          label-width="100px">
          <el-form-item label="用户名">
            <el-tag>{{ systemUserItemFrom.username }}</el-tag>
          </el-form-item>
          <el-form-item label="角色" prop="roleIds">
            <el-select v-model="systemUserRoleItemFrom.roleIds" multiple placeholder="请选择角色">
              <el-option v-for="item in roleSelect" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="resetRoleForm(refSystemUserRoleItemFrom)">取消</el-button>
            <el-button type="primary" :loading="loading" @click="submitRoleForm(refSystemUserRoleItemFrom)">确定</el-button>
          </span>
        </template>
      </el-dialog>

      <el-dialog
        v-model="centerPostDialogVisible"
        :title="titlePost"
        width="40%"
        destroy-on-close
        align-center
        center
        append-to-body
        draggable
        :lock-scroll="false"
        class="dialog-settings">
        <el-form
          ref="refSystemUserPostItemFrom"
          :model="systemUserPostItemFrom"
          :rules="rulesSystemUserPostItemFrom"
          label-width="100px">
          <el-form-item label="用户名">
            <el-tag>{{ systemUserItemFrom.username }}</el-tag>
          </el-form-item>
          <el-form-item label="职位" prop="postIds">
            <el-select v-model="systemUserPostItemFrom.postIds" multiple placeholder="请选择职位">
              <el-option v-for="item in postSelect" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="resetPostForm(refSystemUserPostItemFrom)">取消</el-button>
            <el-button type="primary" :loading="loading" @click="submitPostForm(refSystemUserPostItemFrom)">确定</el-button>
          </span>
        </template>
      </el-dialog>

      <el-dialog
        v-model="centerDeptDialogVisible"
        :title="titleDept"
        width="40%"
        destroy-on-close
        align-center
        center
        append-to-body
        draggable
        :lock-scroll="false"
        class="dialog-settings">
        <el-form
          ref="refSystemUserDeptItemFrom"
          :model="systemUserDeptItemFrom"
          :rules="rulesSystemUserDeptItemFrom"
          label-width="100px">
          <el-form-item label="用户名">
            <el-tag>{{ systemUserItemFrom.username }}</el-tag>
          </el-form-item>
          <el-form-item label="部门" prop="deptIds">
            <el-tree-select
              v-model="systemUserDeptItemFrom.deptIds"
              :data="deptSelect"
              :props="{ value: 'id', label: 'name' }"
              value-key="id"
              node-key="id"
              placeholder="请选择"
              default-expand-all
              multiple
              check-strictly
              :render-after-expand="false" />
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="resetDeptForm(refSystemUserDeptItemFrom)">取消</el-button>
            <el-button type="primary" :loading="loading" @click="submitDeptForm(refSystemUserDeptItemFrom)">确定</el-button>
          </span>
        </template>
      </el-dialog>
    </div>
  </div>
</template>
<script setup lang="ts" name="systemUser">
import { onMounted, ref, reactive } from "vue";
import { ColumnProps, ProTableInstance, SearchProps } from "@/components/ProTable/interface";
import { DictTag } from "@/components/DictTag";
import { EditPen, CirclePlus, Delete, DArrowRight, Refresh, CircleCheck, Key } from "@element-plus/icons-vue";
import { FormInstance, FormRules, ElMessage, ElMessageBox } from "element-plus";
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
import { SystemRole } from "@/api/interface/systemRole";
import { SystemUserRole } from "@/api/interface/systemUserRole";
import { getSystemRoleListSimpleApi } from "@/api/modules/systemRole";
import { getSystemUserRoleListApi, addSystemUserRoleApi } from "@/api/modules/systemUserRole";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import TreeFilter from "@/components/TreeFilter/index.vue";
import { SystemDept } from "@/api/interface/systemDept";
import { getSystemDeptListSimpleApi } from "@/api/modules/systemDept";
import { SystemPost } from "@/api/interface/systemPost";
import { getSystemPostListSimpleApi } from "@/api/modules/systemPost";
import { SystemUserPost } from "@/api/interface/systemUserPost";
import { getSystemUserPostListApi, addSystemUserPostApi } from "@/api/modules/systemUserPost";
import { SystemUserDept } from "@/api/interface/systemUserDept";
import { getSystemUserDeptListApi, addSystemUserDeptApi } from "@/api/modules/systemUserDept";
import { HasPermission } from "@/utils/permission";

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

//数据接口
const systemUserItemFrom = ref<SystemUser.ResSystemUserItem>({
  id: undefined,
  nickname: "",
  mobile: "",
  username: "",
  password: "",
  status: 0,
  deleted: 0,
  roleIds: undefined,
  deptIds: undefined,
  postIds: undefined
});
//校验
const refSystemUserItemFrom = ref<FormInstance>();
//校验
const rulesSystemUserItemFrom = reactive<FormRules>({
  username: [{ required: true, message: "用户名称不能为空", trigger: "blur" }],
  nickname: [{ required: true, message: "用户昵称不能为空", trigger: "blur" }],
  password: [{ required: true, message: "用户密码不能为空", trigger: "blur" }]
});

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
  { prop: "id", label: "编号" },
  { prop: "username", label: "用户名", search: { el: "input", span: 2 } },
  { prop: "nickname", label: "用户昵称" },
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
    isShow: HasPermission(
      "user.SystemUserRoleList",
      "user.SystemUserDeptList",
      "user.SystemUserPostList",
      "user.SystemUserPassword",
      "user.SystemUserDelete",
      "user.SystemUserUpdate",
      "user.SystemUserRecover"
    )
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
    roleIds: undefined,
    deptIds: undefined,
    postIds: undefined
  };
  systemUserRoleItemFrom.value = {
    userId: 0,
    roleIds: []
  };
  systemUserPostItemFrom.value = {
    userId: 0,
    postIds: []
  };
  systemUserDeptItemFrom.value = {
    userId: 0,
    deptIds: []
  };
};
// 弹出层标题
const titleDept = ref();
// 弹出层标题
const centerDeptDialogVisible = ref(false);
//部门树选项
const deptSelect = ref<SystemDept.ResSystemDeptList[]>([]);
// 当前用户的职位
const systemUserDeptItemFrom = ref<SystemUserDept.ResSystemUserDeptItem>({
  userId: 0,
  deptIds: []
});
//校验
const refSystemUserDeptItemFrom = ref<FormInstance>();
//校验
const rulesSystemUserDeptItemFrom = reactive<FormRules>({
  deptIds: [{ required: true, message: "用户部门不能为空", trigger: "blur" }]
});

// 部门处理
const handleDept = async (row: SystemUser.ResSystemUserItem) => {
  reset();
  titleDept.value = "分配部门";
  centerDeptDialogVisible.value = true;
  // 获取用户信息
  systemUserItemFrom.value = row;
  const deptList = await getSystemDeptListSimpleApi();
  deptSelect.value = deptList.data;
  // 获取用户当前的部门
  const userDeptItem = await getSystemUserDeptListApi(Number(row.id));
  systemUserDeptItemFrom.value = userDeptItem.data;
};

// resetForm
const resetDeptForm = (formEl: FormInstance | undefined) => {
  centerDeptDialogVisible.value = false;
  if (!formEl) return;
  formEl.resetFields();
};

// 提交数据
const submitDeptForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loading.value = true;
    const data = systemUserDeptItemFrom.value as unknown as SystemUserDept.ResSystemUserDeptItem;
    await useHandleSet(addSystemUserDeptApi, data.userId, data, "分配部门");
    resetDeptForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

// 弹出层辩题
const titlePost = ref();
// 是否显示弹出层
const centerPostDialogVisible = ref(false);
// 职位列表
const postSelect = ref<SystemPost.ResSystemPostItem[]>([]);
// 当前用户的职位
const systemUserPostItemFrom = ref<SystemUserPost.ResSystemUserPostItem>({
  userId: 0,
  postIds: []
});
//校验
const refSystemUserPostItemFrom = ref<FormInstance>();
//校验
const rulesSystemUserPostItemFrom = reactive<FormRules>({
  postIds: [{ required: true, message: "用户职位不能为空", trigger: "blur" }]
});
// 职位处理
const handlePost = async (row: SystemUser.ResSystemUserItem) => {
  reset();
  titlePost.value = "分配职位";
  centerPostDialogVisible.value = true;
  // 获取用户信息
  systemUserItemFrom.value = row;
  const postList = await getSystemPostListSimpleApi();
  postSelect.value = postList.data;
  // 获取用户当前的职位
  const userPostItem = await getSystemUserPostListApi(Number(row.id));
  systemUserPostItemFrom.value = userPostItem.data;
};
// resetForm
const resetPostForm = (formEl: FormInstance | undefined) => {
  centerPostDialogVisible.value = false;
  if (!formEl) return;
  formEl.resetFields();
};
// 提交数据
const submitPostForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loading.value = true;
    const data = systemUserPostItemFrom.value as unknown as SystemUserPost.ResSystemUserPostItem;
    await useHandleSet(addSystemUserPostApi, data.userId, data, "分配职位");
    resetPostForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

//弹出层标题
const titleRole = ref();
//是否显示弹出层
const centerRoleDialogVisible = ref(false);
// 角色列表
const roleSelect = ref<SystemRole.ResSystemRoleItem[]>([]);
// 当前用户的角色
const systemUserRoleItemFrom = ref<SystemUserRole.ResSystemUserRoleItem>({
  userId: 0,
  roleIds: []
});
//校验
const refSystemUserRoleItemFrom = ref<FormInstance>();
//校验
const rulesSystemUserRoleItemFrom = reactive<FormRules>({
  roleIds: [{ required: true, message: "用户角色不能为空", trigger: "blur" }]
});
// 角色处理
const handleRole = async (row: SystemUser.ResSystemUserItem) => {
  reset();
  titleRole.value = "分配角色";
  centerRoleDialogVisible.value = true;
  // 获取用户信息
  systemUserItemFrom.value = row;
  const roleList = await getSystemRoleListSimpleApi();
  roleSelect.value = roleList.data;
  // 获取用户当前的角色
  const userRoleItem = await getSystemUserRoleListApi(Number(row.id));
  systemUserRoleItemFrom.value = userRoleItem.data;
};

// resetForm
const resetRoleForm = (formEl: FormInstance | undefined) => {
  centerRoleDialogVisible.value = false;
  if (!formEl) return;
  formEl.resetFields();
};

// 提交数据
const submitRoleForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loading.value = true;
    const data = systemUserRoleItemFrom.value as unknown as SystemUserRole.ResSystemUserRoleItem;
    await useHandleSet(addSystemUserRoleApi, data.userId, data, "分配角色");
    resetRoleForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

// 添加按钮
const handleAdd = () => {
  reset();
  title.value = "新增用户";
  centerDialogVisible.value = true;
};

// 编辑按钮
const handleUpdate = async (row: SystemUser.ResSystemUserItem) => {
  reset();
  title.value = "编辑用户";
  centerDialogVisible.value = true;
  const { data } = await getSystemUserItemApi(Number(row.id));
  systemUserItemFrom.value = data;
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
  const { data } = await getSystemDeptListSimpleApi({ tree: 1 });
  deptList.value = data;
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
@import "@/styles/custom.scss";
</style>
