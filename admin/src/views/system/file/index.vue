<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="文件列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemFileListApi"
      :request-auto="true"
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
          ]" />
      </template>
      <!-- 状态-->
      <template #fileSize="scope">
        {{ FileSize(scope.row.fileSize) }}
      </template>
      <template #fileName="scope">
        <span>{{ scope.row.fileName }}</span>
        <div v-if="scope.row.fileType == 1">
          <el-image style="width: 100px" :src="FileUrl(scope.row.filePath)" />
        </div>
      </template>
      <template #view="scope">
        <el-button type="primary" link :icon="View" @click="handleItem(scope.row)" v-if="scope.row.fileType != 4">
          预览
        </el-button>
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'file.SystemFileUpdate'" type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)">
          编辑
        </el-button>
        <el-button v-auth="'file.SystemFileDelete'" type="primary" link :icon="Delete" @click="handleDelete(scope.row)">
          删除
        </el-button>
      </template>
    </ProTable>
    <el-dialog
      v-model="centerDialogVisible"
      :title="title"
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
</template>
<script setup lang="ts" name="systemFile">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh, View } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemFile } from "@/api/interface/systemFile";
import { getSystemFileListApi, deleteSystemFileApi, addSystemFileApi, updateSystemFileApi } from "@/api/modules/systemFile";
import { FormInstance, FormRules, ElMessage, ElMessageBox } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import { FileSize, FileUrl } from "@/utils/file";
import UploadFile from "@/components/Upload/File.vue";
import VueOfficeDocx from "@vue-office/docx";
import VueOfficeExcel from "@vue-office/excel";
import VueOfficePdf from "@vue-office/pdf";
import "@vue-office/docx/lib/index.css";
import "@vue-office/excel/lib/index.css";
//弹出层标题
const title = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
// 是否显示加载
const viewLoading = ref(false);
//数据接口
const systemFileItemFrom = ref<SystemFile.ResSystemFileItem>({
  id: 0, //编号
  fileName: "", //文件名称
  fileType: 0, //文件类型
  fileMimeType: "", //Mime类型
  fileSize: 0, //文件大小
  filePath: "" //文件路径
});

//文件类型
const typeEnum = getIntDictOptions("file.type");
const columns: ColumnProps<SystemFile.ResSystemFileItem>[] = [
  { prop: "id", label: "编号", width: 100 },
  { prop: "fileName", label: "文件名称", search: { el: "input", span: 2, props: { placeholder: "请输入标题" } } },
  { prop: "view", label: "预览" },
  { prop: "fileType", label: "文件类型", tag: true, enum: typeEnum, search: { el: "select", span: 2 } },
  { prop: "fileSize", label: "文件大小" },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission("file.SystemFileUpdate", "file.SystemFileDelete", "file.SystemFileRecover")
  }
];

// 重置数据
const reset = () => {
  centerDialogVisible.value = false;
  systemFileItemFrom.value = {
    id: 0, //编号
    fileName: "", //文件名称
    fileType: 0, //文件类型
    fileMimeType: "", //Mime类型
    fileSize: 0, //文件大小
    filePath: "" //文件路径
  };
};

const renderedHandler = () => {
  console.log("渲染完成");
  viewLoading.value = false;
};

const errorHandler = () => {
  console.log("渲染失败");
  viewLoading.value = false;
  centerDialogVisible.value = false;
};

// 编辑按钮
const handleItem = async (row: SystemFile.ResSystemFileItem) => {
  title.value = "文件预览";
  reset();
  centerDialogVisible.value = true;
  systemFileItemFrom.value = row;
};

// 上传成功之后的回调
const onSuccess = () => {
  proTable.value?.getTableList();
};

// 删除按钮
const handleDelete = async (row: SystemFile.ResSystemFileItem) => {
  await useHandleData(deleteSystemFileApi, Number(row.id), "删除文件");
  proTable.value?.getTableList();
};

const handleUpdate = (row: SystemFile.ResSystemFileItem) => {
  systemFileItemFrom.value = row;
  ElMessageBox.prompt("请输入名称", "温馨提示", {
    confirmButtonText: "确认",
    cancelButtonText: "取消"
  })
    .then(({ value }) => {
      if (value === "") {
        ElMessage.error({ message: "名称不能为空" });
        return;
      }
      systemFileItemFrom.value.fileName = value;
      const data = systemFileItemFrom.value as unknown as SystemFile.ResSystemFileItem;
      updateSystemFileApi(Number(row.id), data).then(() => {
        ElMessage.success({ message: `修改成功！` });
        proTable.value?.getTableList();
      });
    })
    .catch(() => {
      ElMessage.info({ message: "已取消" });
    });
};
</script>

<style lang="scss">
@import "@/styles/custom.scss";
</style>
