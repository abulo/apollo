<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="公告列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemNoticeListApi"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'notice.SystemNoticeCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
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
        <el-button v-auth="'notice.SystemNoticeUpdate'" type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)">
          编辑
        </el-button>
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'notice.SystemNoticeDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'notice.SystemNoticeRecover'"
          type="primary"
          link
          :icon="Refresh"
          @click="handleRecover(scope.row)">
          恢复
        </el-button>
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
      <el-form ref="refSystemNoticeItemFrom" :model="systemNoticeItemFrom" :rules="rulesSystemNoticeItemFrom" label-width="100px">
        <el-form-item label="公告标题" prop="title">
          <el-input v-model="systemNoticeItemFrom.title" />
        </el-form-item>
        <el-form-item label="公告内容" prop="content">
          <el-input v-model="systemNoticeItemFrom.content" type="textarea" />
        </el-form-item>
        <el-form-item label="公告类型" prop="type">
          <el-radio-group v-model="systemNoticeItemFrom.type">
            <el-radio-button v-for="dict in typeEnum" :key="Number(dict.value)" :value="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="公告状态" prop="status">
          <el-radio-group v-model="systemNoticeItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemNoticeItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSystemNoticeItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="systemNotice">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { SystemNotice } from "@/api/interface/systemNotice";
import {
  getSystemNoticeListApi,
  deleteSystemNoticeApi,
  recoverSystemNoticeApi,
  getSystemNoticeItemApi,
  addSystemNoticeApi,
  updateSystemNoticeApi
} from "@/api/modules/systemNotice";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
//加载
const loading = ref(false);
//弹出层标题
const title = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
//数据接口
const systemNoticeItemFrom = ref<SystemNotice.ResSystemNoticeItem>({
  id: 0, // 公告ID
  title: "", // 公告标题
  content: "", // 公告内容
  type: 0, // 公告类型（1通知 2公告）
  status: 0, // 公告状态（0正常 1关闭）
  deleted: 0 // 删除状态（0正常 1删除）
});
//校验
const refSystemNoticeItemFrom = ref<FormInstance>();
//校验
const rulesSystemNoticeItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "公告标题不能为空", trigger: "blur" }],
  type: [{ required: true, message: "公告类别不能为空", trigger: "blur" }],
  content: [{ required: true, message: "公告内容不能为空", trigger: "blur" }],
  status: [{ required: true, message: "公告状态不能为空", trigger: "change" }]
});

//状态
const statusEnum = getIntDictOptions("status");
//删除状态
const deletedEnum = getIntDictOptions("delete");
//公告类型
const typeEnum = getIntDictOptions("notice.type");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("notice.SystemNoticeDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);
const columns: ColumnProps<SystemNotice.ResSystemNoticeItem>[] = [
  { prop: "id", label: "编号", width: 100 },
  { prop: "title", label: "公告标题", search: { el: "input", span: 2, props: { placeholder: "请输入标题" } } },
  { prop: "type", label: "公告类型", tag: true, enum: typeEnum, search: { el: "select", span: 2 } },
  { prop: "status", label: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 } },
  {
    prop: "deleted",
    label: "删除",
    tag: true,
    enum: deletedEnum,
    search: deleteSearch
  },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission("notice.SystemNoticeUpdate", "notice.SystemNoticeDelete", "notice.SystemNoticeRecover")
  }
];
// 重置数据
const reset = () => {
  systemNoticeItemFrom.value = {
    id: 0, // 公告ID
    title: "", // 公告标题
    content: "", // 公告内容
    type: 0, // 公告类型（1通知 2公告）
    status: 0, // 公告状态（0正常 1关闭）
    deleted: 0 // 删除状态（0正常 1删除）
  };
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
    const data = systemNoticeItemFrom.value as unknown as SystemNotice.ResSystemNoticeItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemNoticeApi, data.id, data, "修改公告");
    } else {
      await useHandleData(addSystemNoticeApi, data, "添加公告");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

// 删除按钮
const handleDelete = async (row: SystemNotice.ResSystemNoticeItem) => {
  await useHandleData(deleteSystemNoticeApi, Number(row.id), "删除公告");
  proTable.value?.getTableList();
};

// 恢复按钮
const handleRecover = async (row: SystemNotice.ResSystemNoticeItem) => {
  await useHandleData(recoverSystemNoticeApi, Number(row.id), "恢复公告");
  proTable.value?.getTableList();
};

// 添加按钮
const handleAdd = () => {
  title.value = "新增公告";
  centerDialogVisible.value = true;
  reset();
};

// 编辑按钮
const handleUpdate = async (row: SystemNotice.ResSystemNoticeItem) => {
  title.value = "编辑公告";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getSystemNoticeItemApi(Number(row.id));
  systemNoticeItemFrom.value = data;
};
</script>

<style lang="scss">
@import "@/styles/custom.scss";
</style>
