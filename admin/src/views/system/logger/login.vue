<template>
  <div class="main-box">
    <TreeFilter label="name" :data="deptList" :default-value="initParam.deptId" @change="changeDept" />
    <div class="table-box">
      <ProTable
        ref="proTable"
        title="登录日志列表"
        row-key="id"
        :columns="columns"
        :request-api="getCustomSystemLoginLogListApi"
        :request-auto="true"
        :pagination="true"
        :init-params="initParam"
        :search-col="12">
        <!-- 表格 header 按钮 -->
        <template #tableHeader="scope">
          <el-dropdown trigger="click" class="table-el-dropdown">
            <el-button
              type="danger"
              v-auth="[
                'logger.SystemLoginLogMultipleDelete',
                'logger.SystemLoginLogMultipleRecover',
                'logger.SystemLoginLogMultipleDrop'
              ]"
              :icon="ArrowDownBold"
              plain
              :disabled="!scope.isSelected">
              操作
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  v-auth="'logger.SystemLoginLogMultipleDelete'"
                  :icon="Delete"
                  @click="handleMultipleDelete(scope.selectedListIds)">
                  删除
                </el-dropdown-item>
                <el-dropdown-item
                  v-auth="'logger.SystemLoginLogMultipleDrop'"
                  :icon="DeleteFilled"
                  @click="handleMultipleDrop(scope.selectedListIds)">
                  清理
                </el-dropdown-item>
                <el-dropdown-item
                  v-auth="'logger.SystemLoginLogMultipleRecover'"
                  :icon="Refresh"
                  @click="handleMultipleRecover(scope.selectedListIds)">
                  恢复
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
        <!-- 删除状态 -->
        <template #deleted="scope">
          <DictTag type="delete" :value="scope.row.deleted" />
        </template>
        <!-- 菜单操作 -->
        <template #operation="scope">
          <el-button v-auth="'logger.SystemLoginLog'" type="primary" link :icon="View" @click="handleItem(scope.row)">
            查看
          </el-button>
          <el-dropdown trigger="click">
            <el-button
              v-auth="['logger.SystemLoginLogDelete', 'logger.SystemLoginLogRecover', 'logger.SystemLoginLogDrop']"
              type="primary"
              link
              :icon="DArrowRight">
              更多
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  v-if="scope.row.deleted === 0"
                  v-auth="'logger.SystemLoginLogDelete'"
                  :icon="Delete"
                  @click="handleDelete(scope.row)">
                  删除
                </el-dropdown-item>
                <el-dropdown-item
                  v-if="scope.row.deleted === 1"
                  v-auth="'logger.SystemLoginLogRecover'"
                  :icon="Refresh"
                  @click="handleRecover(scope.row)">
                  恢复
                </el-dropdown-item>
                <el-dropdown-item
                  v-if="scope.row.deleted === 1"
                  v-auth="'logger.SystemLoginLogDrop'"
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
        <el-descriptions :column="1" border>
          <el-descriptions-item label="日志主键" min-width="120">
            {{ systemLoginLogItemFrom.id }}
          </el-descriptions-item>
          <el-descriptions-item label="用户账号">
            {{ systemLoginLogItemFrom.username }}
          </el-descriptions-item>
          <el-descriptions-item label="用户ip">
            {{ systemLoginLogItemFrom.userIp }}
          </el-descriptions-item>
          <el-descriptions-item label="UA">
            {{ systemLoginLogItemFrom.userAgent }}
          </el-descriptions-item>
          <el-descriptions-item label="渠道">
            {{ systemLoginLogItemFrom.channel }}
          </el-descriptions-item>
          <el-descriptions-item label="登录时间">
            {{ systemLoginLogItemFrom.loginTime }}
          </el-descriptions-item>
        </el-descriptions>
      </el-dialog>
    </div>
  </div>
</template>
<script setup lang="ts" name="systemLoginLog">
import { ref, reactive, onMounted } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { Delete, Refresh, DeleteFilled, View, DArrowRight, ArrowDownBold } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemLoginLog } from "@/api/interface/systemLoginLog";
import {
  getSystemLoginLogListApi,
  deleteSystemLoginLogApi,
  deleteSystemLoginLogMultipleApi,
  dropSystemLoginLogApi,
  dropSystemLoginLogMultipleApi,
  recoverSystemLoginLogApi,
  recoverSystemLoginLogMultipleApi,
  getSystemLoginLogApi
} from "@/api/modules/systemLoginLog";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import TreeFilter from "@/components/TreeFilter/index.vue";
import { SystemDept } from "@/api/interface/systemDept";
import { getSystemDeptListSimpleApi } from "@/api/modules/systemDept";
//弹出层标题
const title = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
//数据接口
const systemLoginLogItemFrom = ref<SystemLoginLog.ResSystemLoginLogItem>({
  id: 0, // 主键
  username: "", // 用户账号
  userIp: "", // 用户ip
  userAgent: undefined, // UA
  loginTime: "", // 登录时间
  channel: "", // 渠道
  userId: 0, //
  deleted: 0, // 删除
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined // 更新时间
});
//删除状态
const deletedEnum = getIntDictOptions("delete");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("logger.SystemLoginLogDelete")
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

const columns: ColumnProps<SystemLoginLog.ResSystemLoginLogItem>[] = [
  { type: "selection", fixed: "left", width: 70 },
  { prop: "id", label: "编号" },
  { prop: "username", label: "用户账号", search: { el: "input", span: 2 } },
  { prop: "userIp", label: "用户ip" },
  { prop: "userAgent", label: "UA" },
  { prop: "channel", label: "渠道", search: { el: "input", span: 2 } },
  {
    prop: "loginTime",
    label: "登录时间",
    search: {
      el: "date-picker",
      span: 4,
      props: { type: "datetimerange", valueFormat: "YYYY-MM-DD HH:mm:ss" }
    }
  },
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
    width: 150,
    fixed: "right",
    isShow: HasPermission(
      "logger.SystemLoginLogDelete",
      "logger.SystemLoginLogDrop",
      "logger.SystemLoginLogRecover",
      "logger.SystemLoginLog"
    )
  }
];

// 重置数据
const reset = () => {
  systemLoginLogItemFrom.value = {
    id: 0, // 主键
    username: "", // 用户账号
    userIp: "", // 用户ip
    userAgent: undefined, // UA
    loginTime: "", // 登录时间
    channel: "", // 渠道
    userId: 0, //
    deleted: 0, // 删除
    tenantId: 0, // 租户
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined // 更新时间
  };
};
// 清理按钮
const handleDrop = async (row: SystemLoginLog.ResSystemLoginLogItem) => {
  await useHandleData(dropSystemLoginLogApi, Number(row.id), "清理登录日志");
  proTable.value?.getTableList();
};
// 删除按钮
const handleDelete = async (row: SystemLoginLog.ResSystemLoginLogItem) => {
  await useHandleData(deleteSystemLoginLogApi, Number(row.id), "删除登录日志");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: SystemLoginLog.ResSystemLoginLogItem) => {
  await useHandleData(recoverSystemLoginLogApi, Number(row.id), "恢复登录日志");
  proTable.value?.getTableList();
};
// 查看按钮
const handleItem = async (row: SystemLoginLog.ResSystemLoginLogItem) => {
  title.value = "查看登录日志";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemLoginLogApi(Number(row.id));
  systemLoginLogItemFrom.value = data;
};

const initParam = reactive({ deptId: "" });
// 部门树选择
const deptList = ref<SystemDept.ResSystemDeptItem[]>([]);

const getTreeFilter = async () => {
  const { data } = await getSystemDeptListSimpleApi();
  deptList.value = data;
};
// 树形筛选切换
const changeDept = (val: string) => {
  initParam.deptId = val;
  proTable.value!.pageable.pageNum = 1;
  proTable.value!.searchInitParam.deptId = val;
  proTable.value?.getTableList();
};

// 自定义搜索
const getCustomSystemLoginLogListApi = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.loginTime && (newParams.beginLoginTime = newParams.loginTime[0]);
  newParams.loginTime && (newParams.finishLoginTime = newParams.loginTime[1]);
  delete newParams.loginTime;
  return getSystemLoginLogListApi(newParams);
};

// 批量删除
const handleMultipleDelete = async (ids: string[]) => {
  const data = ref<SystemLoginLog.ReqSystemLoginLogMultiple>({});
  if (ids.length > 0) {
    data.value.ids = ids.map(Number);
  }
  await useHandleData(deleteSystemLoginLogMultipleApi, data.value, "删除登录日志");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

// 批量恢复
const handleMultipleRecover = async (ids: string[]) => {
  const data = ref<SystemLoginLog.ReqSystemLoginLogMultiple>({});
  if (ids.length > 0) {
    data.value.ids = ids.map(Number);
  }
  await useHandleData(recoverSystemLoginLogMultipleApi, data.value, "恢复登录日志");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

// 批量清理
const handleMultipleDrop = async (ids: string[]) => {
  const data = ref<SystemLoginLog.ReqSystemLoginLogMultiple>({});
  if (ids.length > 0) {
    data.value.ids = ids.map(Number);
  }
  await useHandleData(dropSystemLoginLogMultipleApi, data.value, "清理登录日志");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

onMounted(() => {
  getTreeFilter();
});
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
