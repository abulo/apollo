<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="租户列表"
      row-key="id"
      :columns="columns"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button type="primary" :icon="CirclePlus" @click="handleAdd" v-auth="'tenant.SystemTenantCreate'">新增</el-button>
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
        <el-button type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)" v-auth="'tenant.SystemTenantUpdate'">
          编辑
        </el-button>
        <el-button type="primary" link :icon="Connection" @click="handleLogin(scope.row)" v-auth="'tenant.SystemTenantLogin'">
          登录
        </el-button>
        <el-button
          v-if="scope.row.deleted === 0"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)"
          v-auth="'tenant.SystemTenantDelete'">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          type="primary"
          link
          :icon="Refresh"
          @click="handleRecover(scope.row)"
          v-auth="'tenant.SystemTenantRecover'">
          恢复
        </el-button>
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
      @click="handleDialogClick"
      class="dialog-settings">
      <el-form ref="refSystemTenantItemFrom" :model="systemTenantItemFrom" :rules="rulesSystemTenantItemFrom" label-width="100px">
        <el-form-item v-if="systemTenantItemFrom.id === 0" label="用户名" prop="username">
          <el-input v-model="systemTenantItemFrom.username" />
        </el-form-item>
        <el-form-item v-if="systemTenantItemFrom.id === 0" label="用户密码" prop="password">
          <el-input v-model="systemTenantItemFrom.password" show-password type="password" />
        </el-form-item>
        <el-form-item label="负责人" prop="systemUserId" v-if="systemTenantItemFrom.id !== 0">
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
        <el-form-item label="租户名" prop="name">
          <el-input v-model="systemTenantItemFrom.name" />
        </el-form-item>
        <el-form-item label="租户套餐" prop="systemTenantPackageId">
          <el-select v-model="systemTenantItemFrom.systemTenantPackageId" clearable placeholder="请选择租户套餐">
            <el-option v-for="item in tenantPackageEnum" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系人" prop="contactName">
          <el-input v-model="systemTenantItemFrom.contactName" />
        </el-form-item>
        <el-form-item label="联系手机" prop="contactMobile">
          <el-input v-model="systemTenantItemFrom.contactMobile" />
        </el-form-item>
        <el-form-item label="账号额度" prop="accountCount">
          <el-input-number v-model="systemTenantItemFrom.accountCount" />
        </el-form-item>
        <el-form-item label="过期时间" prop="expireDate">
          <el-date-picker
            v-model="systemTenantItemFrom.expireDate"
            clearable
            placeholder="请选择过期时间"
            type="date"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD" />
        </el-form-item>
        <el-form-item label="绑定域名" prop="domain">
          <el-input v-model="systemTenantItemFrom.domain" placeholder="请输入绑定域名" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="systemTenantItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemTenantItemFrom)">取消</el-button>
          <el-button type="primary" @click="submitForm(refSystemTenantItemFrom)" :loading="loading">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx" name="systemTenant">
import { ref, reactive, onMounted } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh, Connection } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemTenant } from "@/api/interface/systemTenant";
import {
  getSystemTenantListApi,
  deleteSystemTenantApi,
  recoverSystemTenantApi,
  getSystemTenantItemApi,
  addSystemTenantApi,
  updateSystemTenantApi,
  getSystemTenantUserSearchApi,
  getSystemTenantLogin
} from "@/api/modules/systemTenant";
import { FormInstance, FormRules, ElMessage } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { SystemUser } from "@/api/interface/systemUser";
import { getSystemUserItemApi } from "@/api/modules/systemUser";
import { SystemTenantPackage } from "@/api/interface/systemTenantPackage";
import { getSystemTenantPackageSearchApi } from "@/api/modules/systemTenantPackage";
import { useRouter } from "vue-router";
import { HOME_URL } from "@/config";
import { getTimeState } from "@/utils";
import { ElNotification } from "element-plus";
import { useUserStore } from "@/stores/modules/user";
import { useTabsStore } from "@/stores/modules/tabs";
import { useKeepAliveStore } from "@/stores/modules/keepAlive";
import { initDynamicRouter } from "@/routers/modules/dynamicRouter";
import { HasPermission } from "@/utils/permission";

const router = useRouter();
const userStore = useUserStore();
const tabsStore = useTabsStore();
const keepAliveStore = useKeepAliveStore();
//加载
const loading = ref(false);
//弹出层标题
const title = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
//数据接口
const systemTenantItemFrom = ref<SystemTenant.ResSystemTenantItem>({
  id: 0, // bigint 租户编号,PRI
  name: "", // varchar 租户名
  systemUserId: undefined, // bigint 联系人的用户编号
  contactName: "", // varchar 联系人
  contactMobile: "", // varchar 联系手机
  status: 0, // tinyint 租户状态（0正常 1停用）
  domain: "", // varchar 绑定域名
  expireDate: "", // datetime 过期时间
  accountCount: 0, // int 账号数量
  systemTenantPackageId: 0, // bigint 套餐 ID
  deleted: 0, // tinyint 是否删除
  username: undefined, // varchar 用户名
  password: undefined // varchar 密码
});
//校验
const refSystemTenantItemFrom = ref<FormInstance>();
//校验
const rulesSystemTenantItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "租户名不能为空", trigger: "blur" }],
  systemTenantPackageId: [{ required: true, message: "租户套餐不能为空", trigger: "blur" }],
  status: [{ required: true, message: "租户状态不能为空", trigger: "blur" }],
  accountCount: [{ required: true, message: "账号额度不能为空", trigger: "blur" }],
  expireDate: [{ required: true, message: "过期时间不能为空", trigger: "blur" }],
  domain: [{ required: true, message: "绑定域名不能为空", trigger: "blur" }]
});

const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.expireDate && (newParams.beginExpireDate = newParams.expireDate[0]);
  newParams.expireDate && (newParams.finishExpireDate = newParams.expireDate[1]);
  delete newParams.expireDate;
  return getSystemTenantListApi(newParams);
};

const statusEnum = getIntDictOptions("status");
const deletedEnum = getIntDictOptions("delete");
const tenantPackageEnum = ref<SystemTenantPackage.ResSystemTenantPackageItem[]>([]);
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("tenant.SystemTenantDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);
const columns: ColumnProps<SystemTenant.ResSystemTenantItem>[] = [
  { prop: "id", label: "编号", width: 100 },
  { prop: "name", label: "租户名", search: { el: "input", span: 2 } },
  {
    prop: "systemTenantPackageId",
    label: "租户套餐",
    tag: true,
    enum: tenantPackageEnum,
    fieldNames: { label: "name", value: "id" },
    search: { el: "select", span: 2 }
  },
  { prop: "contactMobile", label: "联系手机" },
  { prop: "accountCount", label: "账号额度" },
  {
    prop: "expireDate",
    label: "过期时间",
    search: {
      el: "date-picker",
      span: 4,
      props: { type: "daterange", valueFormat: "YYYY-MM-DD" }
    }
  },
  { prop: "domain", label: "绑定域名" },
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
    width: 240,
    fixed: "right",
    isShow: HasPermission(
      "tenant.SystemTenantLogin",
      "tenant.SystemTenantRecover",
      "tenant.SystemTenantDelete",
      "tenant.SystemTenantUpdate"
    )
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
  { prop: "id", label: "编码" },
  { prop: "username", label: "用户名", search: { el: "input", span: 2, props: { placeholder: "请输入用户名/昵称/手机号" } } },
  { prop: "nickname", label: "用户昵称" },
  {
    prop: "operation",
    label: "操作",
    fixed: "right"
  }
];
// 重置数据
const reset = () => {
  loading.value = false;
  systemTenantItemFrom.value = {
    id: 0, // bigint 租户编号,PRI
    name: "", // varchar 租户名
    systemUserId: undefined, // bigint 联系人的用户编号
    contactName: "", // varchar 联系人
    contactMobile: "", // varchar 联系手机
    status: 0, // tinyint 租户状态（0正常 1停用）
    domain: "", // varchar 绑定域名
    expireDate: "", // datetime 过期时间
    accountCount: 0, // int 账号数量
    systemTenantPackageId: 0, // bigint 套餐 ID
    deleted: 0, // tinyint 是否删除
    username: undefined, // varchar 用户名
    password: undefined // varchar 密码
  };
  userItem.value = "点击选择";
  isUserOpenPopover.value = false;
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
    loading.value = false;
    const data = systemTenantItemFrom.value as unknown as SystemTenant.ResSystemTenantItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemTenantApi, data.id, data, "修改租户");
    } else {
      await useHandleData(addSystemTenantApi, data, "添加租户");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

// 删除按钮
const handleDelete = async (row: SystemTenant.ResSystemTenantItem) => {
  await useHandleData(deleteSystemTenantApi, Number(row.id), "删除租户");
  proTable.value?.getTableList();
};

const handleLogin = async (row: SystemTenant.ResSystemTenantItem) => {
  try {
    const { data } = await getSystemTenantLogin(Number(row.id));
    userStore.setToken(data.accessToken);
    // 2.添加动态路由
    await initDynamicRouter();
    // 3.清空 tabs、keepAlive 数据
    userStore.setUserInfo({ name: data.nickname });
    // 4.清空 tabs、keepAlive 数据
    tabsStore.setTabs([]);
    keepAliveStore.setKeepAliveName([]);

    // 5.跳转到首页
    router.push(HOME_URL);
    ElNotification({
      title: getTimeState(),
      message: "欢迎登录 Geeker-Admin",
      type: "success",
      duration: 3000
    });
  } catch {
    ElMessage.error({ message: "登录失败" });
  }
};

// 恢复按钮
const handleRecover = async (row: SystemTenant.ResSystemTenantItem) => {
  await useHandleData(recoverSystemTenantApi, Number(row.id), "恢复租户");
  proTable.value?.getTableList();
};

// 添加按钮
const handleAdd = () => {
  title.value = "新增租户";
  centerDialogVisible.value = true;
  reset();
};

// 编辑按钮
const handleUpdate = async (row: SystemTenant.ResSystemTenantItem) => {
  title.value = "编辑租户";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemTenantItemApi(Number(row.id));
  systemTenantItemFrom.value = data;
  if (Number(data.systemUserId) !== 0) {
    const user = await getSystemUserItemApi(Number(data.systemUserId));
    userItem.value = user.data.nickname;
  }
};

// 在 el-dialog 上添加点击事件监听器
const handleDialogClick = () => {
  isUserOpenPopover.value = false;
};

// 列表数据接口
const getSystemUserSearch = (params: any) => {
  const newParams = Object.assign(params, JSON.parse(JSON.stringify(initParam)));
  return getSystemTenantUserSearchApi(Number(systemTenantItemFrom.value.id), newParams);
};
// 当用户被选择
const handleUser = (row: SystemUser.ResSystemUserItem) => {
  userItem.value = row.nickname;
  systemTenantItemFrom.value.systemUserId = Number(row.id);
  isUserOpenPopover.value = false;
};

onMounted(async () => {
  // 获取接口数据
  let { data } = await getSystemTenantPackageSearchApi();
  if (data === null) {
    data = [] as SystemTenantPackage.ResSystemTenantPackageItem[];
  }
  data.push({
    id: 0, //bigint 套餐编号,PRI
    name: "内置套餐", //varchar 套餐名称
    status: 0, //tinyint 状态（0正常 1停用）
    systemMenuIds: [], //json 目录编号
    remark: "", //varchar 备注
    deleted: 0
  });
  tenantPackageEnum.value = data;
});
</script>

<style lang="scss">
@import "@/styles/custom.scss";
@import "@/styles/popover.scss";
</style>
