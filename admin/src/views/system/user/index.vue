<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="用户列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemUserListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)">编辑</el-button>
        <el-button type="primary" link :icon="Delete" @click="handleDelete(scope.row)">删除</el-button>
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
      <el-form ref="refSystemUserItemFrom" :model="systemUserItemFrom" :rules="rulesSystemUserItemFrom" label-width="100px">
        <el-form-item label="用户昵称" prop="nickname">
          <el-input v-model="systemUserItemFrom.nickname" />
        </el-form-item>
        <el-form-item v-if="systemUserItemFrom.id === undefined" label="用户名" prop="username">
          <el-input v-model="systemUserItemFrom.username" />
        </el-form-item>
        <el-form-item v-if="systemUserItemFrom.id === undefined" label="用户密码" prop="password">
          <el-input v-model="systemUserItemFrom.password" show-password type="password" />
        </el-form-item>
        <el-form-item label="用户状态" prop="status">
          <el-radio-group v-model="systemUserItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="用户角色">
          <el-select v-model="systemUserItemFrom.roleIds" multiple placeholder="请选择角色">
            <el-option v-for="item in roleList" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemUserItemFrom)">取消</el-button>
          <el-button type="primary" @click="submitForm(refSystemUserItemFrom)" :loading="loading">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemUser">
import { ref, reactive } from "vue";
import { ColumnProps, ProTableInstance } from "@/components/ProTable/interface";
import { DictTag } from "@/components/DictTag";
import { EditPen, CirclePlus, Delete } from "@element-plus/icons-vue";
import { FormInstance, FormRules, ElMessage, ElMessageBox } from "element-plus";
import ProTable from "@/components/ProTable/index.vue";
import { SystemUser } from "@/api/interface/systemUser";
import {
  getSystemUserListApi,
  deleteSystemUserApi,
  getSystemUserItemApi,
  addSystemUserApi,
  updateSystemUserApi
} from "@/api/modules/systemUser";
import { getIntDictOptions } from "@/utils/dict";
import { SystemRole } from "@/api/interface/systemRole";
import { getSystemRoleListApi } from "@/api/modules/systemRole";
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
// 角色列表
const roleList = ref<SystemRole.ResSystemRoleItem[]>([]);
//数据接口
const systemUserItemFrom = ref<SystemUser.ResSystemUserItem>({
  id: undefined,
  nickname: "",
  username: "",
  password: "",
  status: 0,
  roleIds: [],
  creator: "",
  createTime: "",
  updater: "",
  updateTime: ""
});
//校验
const refSystemUserItemFrom = ref<FormInstance>();
//校验
const rulesSystemUserItemFrom = reactive<FormRules>({
  username: [{ required: true, message: "用户名称不能为空", trigger: "blur" }],
  nickname: [{ required: true, message: "用户昵称不能为空", trigger: "blur" }],
  password: [{ required: true, message: "用户密码不能为空", trigger: "blur" }]
});

// 定义列表
const columns: ColumnProps<SystemUser.ResSystemUserItem>[] = [
  { prop: "id", label: "编号" },
  { prop: "username", label: "用户名", search: { el: "input", span: 2 } },
  { prop: "nickname", label: "用户昵称" },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 } },
  { prop: "operation", label: "操作", width: 160, fixed: "right" }
];

// 重置数据
const reset = () => {
  systemUserItemFrom.value = {
    id: undefined,
    nickname: "",
    username: "",
    password: "",
    status: 0,
    roleIds: [],
    creator: "",
    createTime: "",
    updater: "",
    updateTime: ""
  };
};
// 获取角色列表
const getRoleList = async () => {
  //角色信息
  const { data } = await getSystemRoleListApi();
  roleList.value = data;
};
// 添加按钮
const handleAdd = () => {
  title.value = "新增用户";
  centerDialogVisible.value = true;
  reset();
  getRoleList();
};

// 编辑按钮
const handleUpdate = async (row: SystemUser.ResSystemUserItem) => {
  title.value = "编辑用户";
  centerDialogVisible.value = true;
  reset();
  getRoleList();
  const { data } = await getSystemUserItemApi(Number(row.id));
  systemUserItemFrom.value = data;
};

// 删除按钮
const handleDelete = (row: SystemUser.ResSystemUserItem) => {
  ElMessageBox.confirm("此操作将删除该数据, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  })
    .then(() => {
      deleteSystemUserApi(Number(row.id)).then(() => {
        ElMessage.success({ message: `${row.nickname}删除成功！` });
        proTable.value?.getTableList();
      });
    })
    .catch(() => {
      ElMessage.info({ message: "已取消删除" });
    });
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
    try {
      if (data.id !== undefined) {
        delete data.password;
        await updateSystemUserApi(data.id, data);
      } else {
        await addSystemUserApi(data);
      }
      ElMessage.success({ message: `${data.nickname}操作成功！` });
    } catch (error) {
      ElMessage.error({ message: `${data.nickname}操作失败！` });
    } finally {
      resetForm(formEl);
      loading.value = false;
      await proTable.value?.getTableList();
    }
  });
};
</script>
