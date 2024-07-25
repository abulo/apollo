<template>
  <div class="main-box">
    <TreeFilter label="name" :data="deptList" :default-value="initParam.deptId" @change="changeDept" />
    <div class="table-box">
      <ProTable
        ref="proTable"
        title="文件列表"
        row-key="id"
        :columns="columns"
        :request-api="getSystemFileListApi"
        :request-auto="true"
        :init-param="initParam"
        :pagination="true"
        :search-col="12">
        <!-- 表格 header 按钮 -->
        <template #tableHeader>
          <UploadFile
            v-auth="'file.SystemFileCreate'"
            :api="addSystemFileApi"
            :callback="onSuccess"
            :file-type="[
              'image/gif',
              'image/jpeg',
              'image/png',
              'application/vnd.ms-excel',
              'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
              'application/msword',
              'application/vnd.openxmlformats-officedocument.wordprocessingml.document',
              'application/pdf',
              'application/vnd.ms-powerpoint',
              'application/vnd.openxmlformats-officedocument.presentationml.presentation'
            ]">
            <template #text>
              <el-button type="primary" :icon="CirclePlus">新增</el-button>
            </template>
          </UploadFile>
        </template>
        <!-- 删除状态 -->
        <template #deleted="scope">
          <DictTag type="delete" :value="scope.row.deleted" />
        </template>
        <template #fileSize="scope">
          {{ FileSize(scope.row.fileSize) }}
        </template>
        <template #view="scope">
          <el-button type="primary" link @click="handleView(scope.row)" v-if="scope.row.fileType != 4">
            <FileIcons :name="scope.row.filePath" :is-folder="false" width="30" height="30" />
          </el-button>
        </template>
        <!-- 菜单操作 -->
        <template #operation="scope">
          <el-button v-auth="'file.SystemFile'" type="primary" link :icon="View" @click="handleItem(scope.row)"> 查看 </el-button>
          <el-dropdown trigger="click">
            <el-button
              v-auth="[
                'file.SystemFileUpdate',
                'file.SystemFileDelete',
                'file.SystemFileRecover',
                'file.SystemFileDrop',
                'file.SystemFileReUpload'
              ]"
              type="primary"
              link
              :icon="DArrowRight">
              更多
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item v-if="scope.row.deleted === 0" v-auth="'file.SystemFileReUpload'" :icon="Upload">
                  <UploadFile
                    @click="handleBeForeUpload(scope.row)"
                    :api="uploadSystemFileApi"
                    :callback="onSuccess"
                    :file-type="[
                      'image/gif',
                      'image/jpeg',
                      'image/png',
                      'application/vnd.ms-excel',
                      'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
                      'application/msword',
                      'application/vnd.openxmlformats-officedocument.wordprocessingml.document',
                      'application/pdf',
                      'application/vnd.ms-powerpoint',
                      'application/vnd.openxmlformats-officedocument.presentationml.presentation'
                    ]">
                    <template #text> 重传 </template>
                  </UploadFile>
                </el-dropdown-item>
                <el-dropdown-item v-auth="'file.SystemFileUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                  编辑
                </el-dropdown-item>
                <el-dropdown-item
                  v-if="scope.row.deleted === 0"
                  v-auth="'file.SystemFileDelete'"
                  :icon="Delete"
                  @click="handleDelete(scope.row)">
                  删除
                </el-dropdown-item>
                <el-dropdown-item
                  v-if="scope.row.deleted === 1"
                  v-auth="'file.SystemFileRecover'"
                  :icon="Refresh"
                  @click="handleRecover(scope.row)">
                  恢复
                </el-dropdown-item>
                <el-dropdown-item
                  v-if="scope.row.deleted === 1"
                  v-auth="'file.SystemFileDrop'"
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
        <el-form ref="refSystemFileItemFrom" :model="systemFileItemFrom" :rules="rulesSystemFileItemFrom" label-width="100px">
          <el-form-item label="文件名称" prop="fileName">
            <el-input v-model="systemFileItemFrom.fileName" :disabled="disabled" />
          </el-form-item>
        </el-form>
        <template #footer v-if="!disabled">
          <span class="dialog-footer">
            <el-button @click="resetForm(refSystemFileItemFrom)">取消</el-button>
            <el-button type="primary" :loading="loading" @click="submitForm(refSystemFileItemFrom)">确定</el-button>
          </span>
        </template>
      </el-dialog>
      <el-dialog
        v-model="viewDialogVisible"
        :title="viewTitle"
        width="50%"
        destroy-on-close
        align-center
        center
        append-to-body
        draggable
        v-loading="viewLoading"
        :lock-scroll="false"
        class="dialog-settings dialog-settings-none">
        <vue-office-pdf
          :src="FileUrl(systemFileItemFrom.filePath)"
          style="height: 100vh"
          @rendered="renderedHandler"
          @error="errorHandler"
          v-if="systemFileItemFrom.fileType == 5" />
        <vue-office-excel
          :src="FileUrl(systemFileItemFrom.filePath)"
          style="height: 100vh"
          @rendered="renderedHandler"
          @error="errorHandler"
          v-if="systemFileItemFrom.fileType == 3" />
        <vue-office-docx
          :src="FileUrl(systemFileItemFrom.filePath)"
          style="height: 100vh"
          @rendered="renderedHandler"
          @error="errorHandler"
          v-if="systemFileItemFrom.fileType == 2" />
        <el-image :src="FileUrl(systemFileItemFrom.filePath)" style="height: 100vh" v-if="systemFileItemFrom.fileType == 1" />
      </el-dialog>
    </div>
  </div>
</template>
<script setup lang="ts" name="systemFile">
import { ref, reactive, onMounted } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh, DeleteFilled, View, DArrowRight, Upload } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemFile } from "@/api/interface/systemFile";
import {
  getSystemFileListApi,
  deleteSystemFileApi,
  dropSystemFileApi,
  recoverSystemFileApi,
  getSystemFileApi,
  addSystemFileApi,
  updateSystemFileApi,
  reUploadSystemFileApi
} from "@/api/modules/systemFile";
import { FormInstance, FormRules, ElMessage } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import { FileSize, FileUrl } from "@/utils/file";
import UploadFile from "@/components/Upload/File.vue";
import VueOfficeDocx from "@vue-office/docx";
import VueOfficeExcel from "@vue-office/excel";
import VueOfficePdf from "@vue-office/pdf";
import "@vue-office/docx/lib/index.css";
import "@vue-office/excel/lib/index.css";
import FileIcons from "file-icons-vue";
import { SystemDept } from "@/api/interface/systemDept";
import { getSystemDeptListSimpleApi } from "@/api/modules/systemDept";
import TreeFilter from "@/components/TreeFilter/index.vue";

const disabled = ref(true);
//加载
const loading = ref(false);
// 预览加载
const viewLoading = ref(false);
//弹出层标题
const title = ref();
// 预览层标题
const viewTitle = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
//是否显示预览弹出层
const viewDialogVisible = ref(false);
//数据接口
const systemFileItemFrom = ref<SystemFile.ResSystemFileItem>({
  id: 0, // 编号
  fileName: "", // 文件名称
  fileType: 0, // 文件类型
  fileMimeType: "", // Mime类型
  fileSize: 0, // 文件大小
  filePath: "", // 文件路径
  userId: 0, // 用户 ID
  deleted: 0, // 是否删除
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined // 更新时间
});
//校验
const refSystemFileItemFrom = ref<FormInstance>();
//校验
const rulesSystemFileItemFrom = reactive<FormRules>({
  fileName: [{ required: true, message: "文件名称不能为空", trigger: "blur" }]
});
//删除状态
const deletedEnum = getIntDictOptions("delete");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("file.SystemFileDelete")
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
// 文件类型
const typeEnum = getIntDictOptions("file.type");
const columns: ColumnProps<SystemFile.ResSystemFileItem>[] = [
  { prop: "id", label: "编号", width: 100, fixed: "left" },
  { prop: "fileName", label: "文件名称", search: { el: "input", span: 2, props: { placeholder: "请输入标题" } } },
  { prop: "view", label: "预览" },
  { prop: "fileType", label: "文件类型", tag: true, enum: typeEnum, search: { el: "select", span: 2 }, isShow: false },
  { prop: "fileSize", label: "文件大小" },
  {
    prop: "deleted",
    label: "删除",
    tag: true,
    enum: deletedEnum,
    search: deleteSearch
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
      "file.SystemFileUpdate",
      "file.SystemFileDelete",
      "file.SystemFileRecover",
      "file.SystemFileDrop",
      "file.SystemFileReUpload"
    )
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  viewLoading.value = false;
  systemFileItemFrom.value = {
    id: 0, // 编号
    fileName: "", // 文件名称
    fileType: 0, // 文件类型
    fileMimeType: "", // Mime类型
    fileSize: 0, // 文件大小
    filePath: "", // 文件路径
    userId: 0, // 用户 ID
    deleted: 0, // 是否删除
    tenantId: 0, // 租户
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined // 更新时间
  };
  disabled.value = true;
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
    const data = systemFileItemFrom.value as unknown as SystemFile.ResSystemFileItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemFileApi, data.id, data, "修改文件");
    } else {
      ElMessage({
        type: "warning",
        message: "暂不支持"
      });
      return;
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};
// 清理按钮
const handleDrop = async (row: SystemFile.ResSystemFileItem) => {
  await useHandleData(dropSystemFileApi, Number(row.id), "清理文件");
  proTable.value?.getTableList();
};
// 删除按钮
const handleDelete = async (row: SystemFile.ResSystemFileItem) => {
  await useHandleData(deleteSystemFileApi, Number(row.id), "删除文件");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: SystemFile.ResSystemFileItem) => {
  await useHandleData(recoverSystemFileApi, Number(row.id), "恢复文件");
  proTable.value?.getTableList();
};
// 编辑按钮
const handleUpdate = async (row: SystemFile.ResSystemFileItem) => {
  title.value = "编辑文件";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemFileApi(Number(row.id));
  systemFileItemFrom.value = data;
  disabled.value = false;
};
// 查看按钮
const handleItem = async (row: SystemFile.ResSystemFileItem) => {
  title.value = "查看文件";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemFileApi(Number(row.id));
  systemFileItemFrom.value = data;
  disabled.value = true;
};

// uploadSystemFileApi 自定义上传
const uploadSystemFileApi = (params: FormData) => {
  return reUploadSystemFileApi(Number(systemFileItemFrom.value.id), params);
};
// handleBeForeUpload  上传之前
const handleBeForeUpload = async (row: SystemFile.ResSystemFileItem) => {
  const { data } = await getSystemFileApi(Number(row.id));
  systemFileItemFrom.value = data;
};

// 上传成功之后的回调
const onSuccess = () => {
  proTable.value?.getTableList();
};

// 预览
const handleView = async (row: SystemFile.ResSystemFileItem) => {
  viewTitle.value = "文件预览";
  reset();
  viewDialogVisible.value = true;
  const { data } = await getSystemFileApi(Number(row.id));
  systemFileItemFrom.value = data;
};
const renderedHandler = () => {
  viewLoading.value = false;
};

const errorHandler = () => {
  viewLoading.value = false;
  viewDialogVisible.value = false;
};

const initParam = reactive({ deptId: "" });
// 部门树选择
const deptList = ref<SystemDept.ResSystemDeptItem[]>([]);
// 树形筛选切换
const changeDept = (val: string) => {
  initParam.deptId = val;
  proTable.value!.pageable.pageNum = 1;
  proTable.value!.searchInitParam.deptId = val;
  proTable.value?.getTableList();
};

const getTreeFilter = async () => {
  const { data } = await getSystemDeptListSimpleApi();
  deptList.value = data;
};

onMounted(() => {
  getTreeFilter();
});
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
