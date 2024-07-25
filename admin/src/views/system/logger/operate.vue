<template>
  <div class="main-box">
    <TreeFilter label="name" :data="deptList" :default-value="initParam.deptId" @change="changeDept" />
    <div class="table-box">
      <ProTable
        ref="proTable"
        title="操作日志列表"
        row-key="id"
        :columns="columns"
        :request-api="getCustomSystemOperateLogListApi"
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
                'logger.SystemOperateLogMultipleDelete',
                'logger.SystemOperateLogMultipleRecover',
                'logger.SystemOperateLogMultipleDrop'
              ]"
              :icon="ArrowDownBold"
              plain
              :disabled="!scope.isSelected">
              操作
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  v-auth="'logger.SystemOperateLogMultipleDelete'"
                  :icon="Delete"
                  @click="handleMultipleDelete(scope.selectedListIds)">
                  删除
                </el-dropdown-item>
                <el-dropdown-item
                  v-auth="'logger.SystemOperateLogMultipleDrop'"
                  :icon="DeleteFilled"
                  @click="handleMultipleDrop(scope.selectedListIds)">
                  清理
                </el-dropdown-item>
                <el-dropdown-item
                  v-auth="'logger.SystemOperateLogMultipleRecover'"
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
          <el-button v-auth="'logger.SystemOperateLog'" type="primary" link :icon="View" @click="handleItem(scope.row)">
            查看
          </el-button>
          <el-dropdown trigger="click">
            <el-button
              v-auth="['logger.SystemOperateLogDelete', 'logger.SystemOperateLogRecover', 'logger.SystemOperateLogDrop']"
              type="primary"
              link
              :icon="DArrowRight">
              更多
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  v-if="scope.row.deleted === 0"
                  v-auth="'logger.SystemOperateLogDelete'"
                  :icon="Delete"
                  @click="handleDelete(scope.row)">
                  删除
                </el-dropdown-item>
                <el-dropdown-item
                  v-if="scope.row.deleted === 1"
                  v-auth="'logger.SystemOperateLogRecover'"
                  :icon="Refresh"
                  @click="handleRecover(scope.row)">
                  恢复
                </el-dropdown-item>
                <el-dropdown-item
                  v-if="scope.row.deleted === 1"
                  v-auth="'logger.SystemOperateLogDrop'"
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
          <el-descriptions-item label="主键" min-width="120">
            {{ systemOperateLogItemFrom.id }}
          </el-descriptions-item>
          <el-descriptions-item label="用户账号" min-width="120">
            {{ systemOperateLogItemFrom.username }}
          </el-descriptions-item>
          <el-descriptions-item label="模块名称" min-width="120">
            {{ systemOperateLogItemFrom.module }}
          </el-descriptions-item>
          <el-descriptions-item label="请求方法" min-width="120">
            {{ systemOperateLogItemFrom.requestMethod }}
          </el-descriptions-item>
          <el-descriptions-item label="请求地址" min-width="120">
            {{ systemOperateLogItemFrom.requestUrl }}
          </el-descriptions-item>
          <el-descriptions-item label="用户ip" min-width="120">
            {{ systemOperateLogItemFrom.userIp }}
          </el-descriptions-item>
          <el-descriptions-item label="UA" min-width="120">
            {{ systemOperateLogItemFrom.userAgent }}
          </el-descriptions-item>
          <el-descriptions-item label="方法" min-width="120">
            {{ systemOperateLogItemFrom.goMethod }}
          </el-descriptions-item>
          <el-descriptions-item label="方法参数" min-width="120">
            {{ systemOperateLogItemFrom.goMethodArgs }}
          </el-descriptions-item>
          <el-descriptions-item label="日志时间" min-width="120">
            {{ systemOperateLogItemFrom.startTime }}
          </el-descriptions-item>
          <el-descriptions-item label="耗时" min-width="120">
            {{ systemOperateLogItemFrom.duration }}
          </el-descriptions-item>
          <el-descriptions-item label="渠道" min-width="120">
            {{ systemOperateLogItemFrom.channel }}
          </el-descriptions-item>
          <el-descriptions-item label="结果" min-width="120">
            {{ systemOperateLogItemFrom.result }}
          </el-descriptions-item>
        </el-descriptions>
      </el-dialog>
    </div>
  </div>
</template>
<script setup lang="ts" name="systemOperateLog">
import { ref, reactive, onMounted } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { Delete, Refresh, DeleteFilled, View, DArrowRight, ArrowDownBold } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemOperateLog } from "@/api/interface/systemOperateLog";
import {
  getSystemOperateLogListApi,
  deleteSystemOperateLogApi,
  deleteSystemOperateLogMultipleApi,
  dropSystemOperateLogApi,
  dropSystemOperateLogMultipleApi,
  recoverSystemOperateLogApi,
  recoverSystemOperateLogMultipleApi,
  getSystemOperateLogApi
} from "@/api/modules/systemOperateLog";
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
const systemOperateLogItemFrom = ref<SystemOperateLog.ResSystemOperateLogItem>({
  id: 0, // 主键
  username: "", // 用户账号
  module: "", // 模块名称
  requestMethod: "", // 请求方法名
  requestUrl: "", // 请求地址
  userIp: "", // 用户 ip
  userAgent: undefined, // UA
  goMethod: "", // 方法名
  goMethodArgs: undefined, // 方法的参数
  startTime: "", // 操作开始时间
  duration: 0, // 执行时长
  channel: "", // 渠道
  result: 0, // 结果(0 成功/1 失败)
  userId: 0, // 用户 ID
  deleted: 0, // 删除
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined // 更新时间
});
// 操作状态
const operateStatusEnum = getIntDictOptions("operate.status");
//删除状态
const deletedEnum = getIntDictOptions("delete");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("logger.SystemOperateLogDelete")
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

const columns: ColumnProps<SystemOperateLog.ResSystemOperateLogItem>[] = [
  { type: "selection", fixed: "left", width: 70 },
  { prop: "id", label: "编号" },
  { prop: "username", label: "用户账号", search: { el: "input", span: 2 } },
  { prop: "module", label: "模块名称" },
  { prop: "requestMethod", label: "请求方法" },
  { prop: "requestUrl", label: "请求地址" },
  { prop: "userIp", label: "用户ip" },
  { prop: "userAgent", label: "UA" },
  { prop: "goMethod", label: "方法" },
  {
    prop: "startTime",
    label: "日志时间",
    search: {
      el: "date-picker",
      span: 4,
      props: { type: "datetimerange", valueFormat: "YYYY-MM-DD HH:mm:ss" }
    }
  },
  { prop: "duration", label: "耗时" },
  { prop: "channel", label: "渠道" },
  { prop: "result", label: "结果", enum: operateStatusEnum },
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
      "logger.SystemOperateLogDelete",
      "logger.SystemOperateLogDrop",
      "logger.SystemOperateLogRecover",
      "logger.SystemOperateLog"
    )
  }
];

// 重置数据
const reset = () => {
  systemOperateLogItemFrom.value = {
    id: 0, // 主键
    username: "", // 用户账号
    module: "", // 模块名称
    requestMethod: "", // 请求方法名
    requestUrl: "", // 请求地址
    userIp: "", // 用户 ip
    userAgent: undefined, // UA
    goMethod: "", // 方法名
    goMethodArgs: undefined, // 方法的参数
    startTime: "", // 操作开始时间
    duration: 0, // 执行时长
    channel: "", // 渠道
    result: 0, // 结果(0 成功/1 失败)
    userId: 0, // 用户 ID
    deleted: 0, // 删除
    tenantId: 0, // 租户
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined // 更新时间
  };
};

// 清理按钮
const handleDrop = async (row: SystemOperateLog.ResSystemOperateLogItem) => {
  await useHandleData(dropSystemOperateLogApi, Number(row.id), "清理操作日志");
  proTable.value?.getTableList();
};
// 删除按钮
const handleDelete = async (row: SystemOperateLog.ResSystemOperateLogItem) => {
  await useHandleData(deleteSystemOperateLogApi, Number(row.id), "删除操作日志");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: SystemOperateLog.ResSystemOperateLogItem) => {
  await useHandleData(recoverSystemOperateLogApi, Number(row.id), "恢复操作日志");
  proTable.value?.getTableList();
};
// 查看按钮
const handleItem = async (row: SystemOperateLog.ResSystemOperateLogItem) => {
  title.value = "查看操作日志";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemOperateLogApi(Number(row.id));
  systemOperateLogItemFrom.value = data;
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
const getCustomSystemOperateLogListApi = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.startTime && (newParams.beginStartTime = newParams.startTime[0]);
  newParams.startTime && (newParams.finishStartTime = newParams.startTime[1]);
  delete newParams.startTime;
  return getSystemOperateLogListApi(newParams);
};

// 批量删除
const handleMultipleDelete = async (ids: string[]) => {
  const data = ref<SystemOperateLog.ReqSystemOperateLogMultiple>({});
  if (ids.length > 0) {
    data.value.ids = ids.map(Number);
  }
  await useHandleData(deleteSystemOperateLogMultipleApi, data.value, "删除登录日志");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

// 批量恢复
const handleMultipleRecover = async (ids: string[]) => {
  const data = ref<SystemOperateLog.ReqSystemOperateLogMultiple>({});
  if (ids.length > 0) {
    data.value.ids = ids.map(Number);
  }
  await useHandleData(recoverSystemOperateLogMultipleApi, data.value, "恢复登录日志");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

// 批量清理
const handleMultipleDrop = async (ids: string[]) => {
  const data = ref<SystemOperateLog.ReqSystemOperateLogMultiple>({});
  if (ids.length > 0) {
    data.value.ids = ids.map(Number);
  }
  await useHandleData(dropSystemOperateLogMultipleApi, data.value, "清理登录日志");
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
