<template>
  <div class="table-box">
    <VirtualTable
      ref="proTable"
      title="菜单列表"
      row-key="id"
      :columns="columns"
      :request-api="getSystemMenuListApi"
      :request-auto="true"
      :pagination="false"
      height="auto"
      :column-config="{ resizable: true }"
      :tree-config="{ transform: true, iconOpen: 'vxe-icon-arrow-down', iconClose: 'vxe-icon-arrow-right' }"
      :scroll-y="{ enabled: true, gt: 20 }"
      :init-param="initParam"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button type="primary" :icon="CirclePlus" @click="handleAdd()" v-auth="'menu.SystemMenuCreate'">新增</el-button>
        <el-button type="primary" :icon="Sort" @click="toggleExpandAll">展开/折叠</el-button>
      </template>
      <!-- 菜单图标 -->
      <template #icon="scope">
        <el-icon :size="18">
          <component :is="scope.row.icon"></component>
        </el-icon>
      </template>
      <!-- 类别 -->
      <template #type="scope">
        <DictTag type="menu.type" :value="scope.row.type" />
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
        <el-button type="primary" link :icon="CirclePlus" @click="handleAdd(scope.row)" v-auth="'menu.SystemMenuCreate'">
          新增
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            type="primary"
            link
            :icon="DArrowRight"
            v-auth="['menu.SystemMenuUpdate', 'menu.SystemMenuDelete', 'menu.SystemMenuRecover']">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item :icon="EditPen" @click="handleUpdate(scope.row)" v-auth="'menu.SystemMenuUpdate'">
                编辑
              </el-dropdown-item>
              <el-dropdown-item
                :icon="Delete"
                v-if="scope.row.deleted === 0"
                @click="handleDelete(scope.row)"
                v-auth="'menu.SystemMenuDelete'">
                删除
              </el-dropdown-item>
              <el-dropdown-item
                :icon="Refresh"
                v-if="scope.row.deleted === 1"
                @click="handleRecover(scope.row)"
                v-auth="'menu.SystemMenuRecover'">
                恢复
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </template>
    </VirtualTable>
    <!-- 新增/编辑弹窗 -->
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
      <el-form ref="refSystemMenuItemFrom" :model="systemMenuItemFrom" :rules="rulesSystemMenuItemFrom" label-width="100px">
        <el-form-item label="上级菜单" prop="parentId">
          <el-tree-select
            v-model="systemMenuItemFrom.parentId"
            :data="menuOptions"
            :props="{ value: 'id', label: 'name' }"
            value-key="id"
            node-key="id"
            placeholder="请选择"
            check-strictly
            :render-after-expand="false" />
        </el-form-item>
        <el-form-item label="菜单名称" prop="name">
          <el-input v-model="systemMenuItemFrom.name" placeholder="请输入菜单名称" />
        </el-form-item>
        <el-form-item label="菜单类型" prop="type">
          <el-radio-group v-model="systemMenuItemFrom.type">
            <el-radio-button v-for="dict in typeEnum" :key="dict.value" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="systemMenuItemFrom.type !== 3" label="菜单图标" prop="icon">
          <SelectIcon v-model="systemMenuItemFrom.icon" title="请选择图标" placeholder="搜索图标" :show-icon-name="true" />
        </el-form-item>
        <el-form-item label="显示排序" prop="sort">
          <el-input-number v-model="systemMenuItemFrom.sort" controls-position="right" :min="0" />
        </el-form-item>
        <el-form-item v-if="systemMenuItemFrom.type !== 3" label="路由地址" prop="path">
          <el-input v-model="systemMenuItemFrom.path" placeholder="请输入路由地址" />
        </el-form-item>
        <el-form-item v-if="systemMenuItemFrom.type !== 1" label="权限标识" prop="permission">
          <el-input v-model="systemMenuItemFrom.permission" placeholder="请输入权限标识" />
        </el-form-item>
        <el-form-item v-if="systemMenuItemFrom.type === 2" label="组件路径" prop="component">
          <el-input v-model="systemMenuItemFrom.component" placeholder="请输入组件路径" />
        </el-form-item>
        <el-form-item v-if="systemMenuItemFrom.type === 2" label="组件名称" prop="componentName">
          <el-input v-model="systemMenuItemFrom.componentName" placeholder="请输入组件名称" />
        </el-form-item>
        <el-form-item v-if="systemMenuItemFrom.type === 2" label="激活链接" prop="activePath">
          <el-input v-model="systemMenuItemFrom.activePath" placeholder="请输入激活链接" />
        </el-form-item>
        <el-form-item v-if="systemMenuItemFrom.type !== 3" label="重定向" prop="redirect">
          <el-input v-model="systemMenuItemFrom.redirect" placeholder="请输入重定向链接" />
        </el-form-item>
        <el-form-item label="菜单状态" prop="status">
          <el-radio-group v-model="systemMenuItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="systemMenuItemFrom.type !== 3" label="隐藏状态" prop="hide">
          <el-radio-group v-model="systemMenuItemFrom.hide">
            <el-radio-button v-for="dict in hideEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="systemMenuItemFrom.type !== 3" label="Tab显示" prop="affix">
          <el-radio-group v-model="systemMenuItemFrom.affix">
            <el-radio-button v-for="dict in affixEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="systemMenuItemFrom.type === 2" label="缓存状态" prop="keepAlive">
          <el-radio-group v-model="systemMenuItemFrom.keepAlive">
            <el-radio-button v-for="dict in keepAliveEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="systemMenuItemFrom.type === 2" label="全屏" prop="fullScreen">
          <el-radio-group v-model="systemMenuItemFrom.fullScreen">
            <el-radio-button v-for="dict in fullScreenEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSystemMenuItemFrom)">取消</el-button>
          <el-button type="primary" @click="submitForm(refSystemMenuItemFrom)" :loading="loading">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx" name="systemMenu">
import { ref, reactive } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/VirtualTable/interface";
import { EditPen, CirclePlus, Sort, Delete, Refresh, DArrowRight } from "@element-plus/icons-vue";
import VirtualTable from "@/components/VirtualTable/index.vue";
import { SystemMenu } from "@/api/interface/systemMenu";
import {
  getSystemMenuListApi,
  getSystemMenuItemApi,
  addSystemMenuApi,
  updateSystemMenuApi,
  deleteSystemMenuApi,
  recoverSystemMenuApi
} from "@/api/modules/systemMenu";
import { FormInstance, FormRules } from "element-plus";
import SelectIcon from "@/components/SelectIcon/index.vue";
import { getIntDictOptions } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
const initParam = reactive({ tree: 0 });
//加载
const loading = ref(false);
//table数据
const proTable = ref<ProTableInstance>();
//是否展开，默认全部折叠
const isExpandAll = ref(false);
//弹出层标题
const title = ref();
//是否显示弹出层
const centerDialogVisible = ref(false);
//菜单树选项
const menuOptions = ref<SystemMenu.ResSystemMenuList[]>([]);
//数据接口
const systemMenuItemFrom = ref<SystemMenu.ResSystemMenuItem>({
  id: 0, //菜单ID
  name: "", //菜单名称
  permission: "", //权限标识
  type: 0, //菜单类型(1:目录/2: 菜单/3: 按钮)
  sort: 0, //显示顺序
  parentId: 0, //父菜单ID
  path: "", //路由地址
  icon: "", //菜单图标
  component: "", //组件路径
  componentName: "", //组件名
  status: 0, //菜单状态(0开启/1关闭)
  hide: 0, //是否隐藏(0:否/1是)
  link: "", //路由外链时填写的访问地址
  keepAlive: 0, //是否缓存(0不/ 是)
  affix: 0, //是否总是显示(0 不显示/1 显示)
  activePath: "", //激活链接
  fullScreen: 0, //是否全屏
  redirect: "", //路由重定向地址
  creator: "", //创建者
  createTime: "", //创建时间
  updater: "", //更新者
  updateTime: "", //更新时间
  deleted: 0 //是否删除
});
//校验
const refSystemMenuItemFrom = ref<FormInstance>();
//校验
const rulesSystemMenuItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "请输入菜单名称", trigger: "blur" }]
});

// 重置数据
const reset = () => {
  systemMenuItemFrom.value = {
    id: 0, //菜单ID
    name: "", //菜单名称
    permission: "", //权限标识
    type: 0, //菜单类型(1:目录/2: 菜单/3: 按钮)
    sort: 0, //显示顺序
    parentId: 0, //父菜单ID
    path: "", //路由地址
    icon: "", //菜单图标
    component: "", //组件路径
    componentName: "", //组件名
    status: 0, //菜单状态(0开启/1关闭)
    hide: 0, //是否隐藏(0:否/1是)
    link: "", //路由外链时填写的访问地址
    keepAlive: 0, //是否缓存(0不/ 是)
    affix: 0, //是否总是显示(0 不显示/1 显示)
    activePath: "", //激活链接
    fullScreen: 0, //是否全屏
    redirect: "", //路由重定向地址
    creator: "", //创建者
    createTime: "", //创建时间
    updater: "", //更新者
    updateTime: "", //更新时间
    deleted: 0 //是否删除
  };
};

// resetForm
const resetForm = (formEl: FormInstance | undefined) => {
  centerDialogVisible.value = false;
  if (!formEl) return;
  formEl.resetFields();
};

//菜单类别
const typeEnum = getIntDictOptions("menu.type");
//菜单状态
const statusEnum = getIntDictOptions("status");
//菜单删除状态
const deletedEnum = getIntDictOptions("delete");
//是否隐藏
const hideEnum = getIntDictOptions("menu.hide");
//总是显示
const affixEnum = getIntDictOptions("menu.affix");
//缓存
const keepAliveEnum = getIntDictOptions("menu.keepAlive");
//全屏
const fullScreenEnum = getIntDictOptions("menu.fullScreen");

// 设置展开合并
const toggleExpandAll = () => {
  isExpandAll.value = !isExpandAll.value;
  proTable.value?.element?.setAllTreeExpand(isExpandAll.value);
};

// 添加按钮
const handleAdd = (row?: SystemMenu.ResSystemMenuItem) => {
  title.value = "新增菜单";
  centerDialogVisible.value = true;
  reset();
  getTreeSelect();
  if (row != null && row.id) {
    systemMenuItemFrom.value.parentId = row.id;
  } else {
    systemMenuItemFrom.value.parentId = 0;
  }
};

// 编辑按钮
const handleUpdate = async (row: SystemMenu.ResSystemMenuItem) => {
  title.value = "编辑菜单";
  centerDialogVisible.value = true;
  reset();
  getTreeSelect();
  const { data } = await getSystemMenuItemApi(Number(row.id));
  systemMenuItemFrom.value = data;
  if (data.type === 3) {
    systemMenuItemFrom.value.icon = "";
    systemMenuItemFrom.value.path = "";
    systemMenuItemFrom.value.component = "";
    systemMenuItemFrom.value.componentName = "";
  }
};

// 删除按钮
const handleDelete = async (row: SystemMenu.ResSystemMenuItem) => {
  await useHandleData(deleteSystemMenuApi, Number(row.id), "删除菜单");
  proTable.value?.getTableList();
};

// 恢复按钮
const handleRecover = async (row: SystemMenu.ResSystemMenuItem) => {
  await useHandleData(recoverSystemMenuApi, Number(row.id), "恢复菜单");
  proTable.value?.getTableList();
};

// 获取菜单树选项
const getTreeSelect = async () => {
  menuOptions.value = [];
  const res = await getSystemMenuListApi();
  let obj: SystemMenu.ResSystemMenuList = {
    id: 0,
    name: "主类目",
    children: res.data,
    permission: "",
    type: 0,
    sort: 0,
    parentId: 0,
    path: "",
    icon: "",
    component: "",
    componentName: "",
    status: 0,
    hide: 0,
    link: "",
    keepAlive: 0,
    affix: 0,
    activePath: "",
    fullScreen: 0,
    redirect: "",
    creator: "",
    createTime: "",
    updater: "",
    updateTime: "",
    deleted: 0
  };
  menuOptions.value.push(obj);
};

// 提交数据
const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loading.value = true;
    const data = systemMenuItemFrom.value as unknown as SystemMenu.ResSystemMenuItem;
    if (data.id !== 0) {
      await useHandleSet(updateSystemMenuApi, data.id, data, "修改菜单");
    } else {
      await useHandleData(addSystemMenuApi, data, "添加菜单");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("menu.SystemMenuDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);

const columns: ColumnProps<SystemMenu.ResSystemMenuList>[] = [
  { field: "id", title: "编号", width: 100 },
  { field: "name", title: "菜单名称", align: "left", treeNode: true },
  { field: "type", title: "菜单类别", tag: true, enum: typeEnum, width: 100 },
  { field: "icon", title: "菜单图标", width: 100 },
  { field: "sort", title: "排序", width: 100 },
  { field: "path", title: "路由地址" },
  { field: "permission", title: "权限标识" },
  { field: "component", title: "组件路径" },
  { field: "componentName", title: "组件别名" },
  { field: "status", title: "状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 } },
  {
    field: "deleted",
    title: "删除",
    tag: true,
    enum: deletedEnum,
    search: deleteSearch,
    width: 100
  },
  {
    field: "operation",
    title: "操作",
    width: 160,
    fixed: "right",
    isShow: HasPermission("menu.SystemMenuUpdate", "menu.SystemMenuDelete", "menu.SystemMenuRecover", "menu.SystemMenuCreate")
  }
];
</script>
<style lang="scss">
@import "@/styles/custom.scss";
</style>
