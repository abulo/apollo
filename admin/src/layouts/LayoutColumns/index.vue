<!-- 分栏布局 -->
<template>
  <el-container class="layout">
    <div class="aside-split">
      <div class="logo flx-center">
        <img class="logo-img" src="@/assets/images/logo.svg" alt="logo" />
      </div>
      <el-scrollbar>
        <div class="split-list">
          <div
            v-for="item in menuList"
            :key="item.meta.id"
            class="split-item"
            :class="{ 'split-active': splitActive === String(item.meta.id) }"
            @click="changeSubMenu(item)">
            <el-icon>
              <component :is="item.meta.icon" />
            </el-icon>
            <span class="title">{{ item.meta.title }}</span>
          </div>
        </div>
      </el-scrollbar>
    </div>
    <el-aside :class="{ 'not-aside': !subMenuList.length }" :style="{ width: isCollapse ? '65px' : '210px' }">
      <div class="logo flx-center">
        <span v-show="subMenuList.length" class="logo-text">{{ isCollapse ? "G" : title }}</span>
      </div>
      <el-scrollbar>
        <el-menu
          :router="false"
          :default-active="activeMenu"
          :collapse="isCollapse"
          :unique-opened="accordion"
          :collapse-transition="false">
          <SubMenu :menu-list="subMenuList" />
        </el-menu>
      </el-scrollbar>
    </el-aside>
    <el-container>
      <el-header>
        <ToolBarLeft />
        <ToolBarRight />
      </el-header>
      <Main />
    </el-container>
  </el-container>
</template>

<script setup lang="ts" name="layoutColumns">
import { ref, computed, watch, onMounted } from "vue";
import { useRoute, useRouter, RouteLocationMatched } from "vue-router";
import { useAuthStore } from "@/stores/modules/auth";
import { useGlobalStore } from "@/stores/modules/global";
import Main from "@/layouts/components/Main/index.vue";
import ToolBarLeft from "@/layouts/components/Header/ToolBarLeft.vue";
import ToolBarRight from "@/layouts/components/Header/ToolBarRight.vue";
import SubMenu from "@/layouts/components/Menu/SubMenu.vue";
import { findRootMenuByPath, getShowMenuItem } from "@/utils/index";
const title = import.meta.env.VITE_GLOB_APP_TITLE;

const route = useRoute();
const router = useRouter();
const authStore = useAuthStore();
const globalStore = useGlobalStore();
const accordion = computed(() => globalStore.accordion);
const isCollapse = computed(() => globalStore.isCollapse);
const menuList = computed(() => authStore.showMenuListGet);
// const activeMenu = computed(() => (route.meta.activeMenu ? route.meta.activeMenu : String(route.meta.id)) as string);

const subMenuList = ref<Menu.MenuOptions[]>([]);
const splitActive = ref();
const activeMenu = ref();

const flatMenuList = computed(() => authStore.flatMenuListGet);
const getRouteItem = () => {
  const name = route.name;
  let routeItemId: number = 0;
  for (const item of route.matched) {
    if (item.name == name) {
      routeItemId = Number(item.meta.id);
      break;
    }
  }
  return flatMenuList.value.find(item => item.meta.id === routeItemId);
};

const getActiveHeaderMenu = () => {
  const routeItem = getRouteItem() as Menu.MenuOptions;
  const menuItem = findRootMenuByPath(authStore.authMenuList, routeItem);
  return String(menuItem?.meta.id) || "";
};

watch(
  () => [menuList, route],
  () => {
    // 当前菜单没有数据直接 return
    if (!menuList.value.length) return;
    console.log("dd");
    activeMenu.value = route.meta?.activeMenu ? route.meta?.activeMenu : String(route.meta.id);
    console.log(activeMenu.value);
    splitActive.value = getActiveHeaderMenu();
    const routeItem = getRouteItem() as Menu.MenuOptions;
    const menuItem = getShowMenuItem(findRootMenuByPath(authStore.authMenuList, routeItem) as Menu.MenuOptions);
    if (menuItem?.children?.length) return (subMenuList.value = menuItem.children);
    subMenuList.value = [];
  },
  {
    deep: true,
    immediate: true
  }
);

// change SubMenu
const changeSubMenu = (item: Menu.MenuOptions) => {
  activeMenu.value = route.meta?.activeMenu ? route.meta?.activeMenu : String(route.meta.id);
  let actionRouteItem = findRootMenuByPath(menuList.value, item);
  splitActive.value = String(actionRouteItem?.meta.id) || "";
  if (item?.children?.length) return (subMenuList.value = item.children);
  subMenuList.value = [];
  router.push(item.path);
};

onMounted(() => {
  activeMenu.value = route.meta?.activeMenu ? route.meta?.activeMenu : String(route.meta.id);
  console.log("m", activeMenu.value);
});
</script>

<style scoped lang="scss">
@import "./index.scss";
</style>
