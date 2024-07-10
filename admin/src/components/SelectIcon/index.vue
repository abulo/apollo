<template>
  <el-popover
    :placement="placement"
    trigger="focus"
    :hide-after="0"
    :width="state.selectorWidth"
    :visible="state.popoverVisible"
    :show-arrow="false">
    <div
      class="icon-selector"
      @mouseover.stop="state.iconSelectorMouseover = true"
      @mouseout.stop="state.iconSelectorMouseover = false">
      <transition name="el-zoom-in-center">
        <div v-if="show" class="icon-selector-box">
          <div class="selector-header">
            <div class="selector-title">{{ title ? title : "" }}</div>
          </div>
          <div class="selector-body">
            <el-scrollbar ref="selectorScrollbarRef" height="20vh">
              <div v-if="renderFontIconNames.length > 0">
                <div
                  v-for="(item, key) in renderFontIconNames"
                  :key="key"
                  class="icon-selector-item"
                  :title="item"
                  @click="onIcon(item)">
                  <el-button :icon="customIcons[item]" />
                </div>
              </div>
            </el-scrollbar>
          </div>
        </div>
      </transition>
    </div>
    <template #reference>
      <el-input
        ref="selectorInput"
        v-model="state.inputValue"
        :size="size"
        :disabled="disabled"
        :placeholder="placeholder"
        :class="'size-' + size"
        @focus="onInputFocus"
        @blur="onInputBlur">
        <template #prepend>
          <div class="icon-prepend">
            <el-button
              :key="'icon' + state.iconKey"
              :icon="customIcons[state.prependIcon ? state.prependIcon : state.defaultModelValue]" />
            <div v-if="showIconName" class="name">
              {{ state.prependIcon ? state.prependIcon : state.defaultModelValue }}
            </div>
          </div>
        </template>
        <template #append>
          <el-button :icon="Icons.Refresh" @click="onInputRefresh" />
        </template>
      </el-input>
    </template>
  </el-popover>
</template>

<script setup lang="ts" name="SelectIcon">
import { reactive, ref, onMounted, onUnmounted, nextTick, watch, computed } from "vue";
import { useEventListener } from "@vueuse/core";
import { Placement } from "element-plus";
import * as Icons from "@element-plus/icons-vue";

interface SelectIconProps {
  size?: "default" | "small" | "large";
  disabled?: boolean;
  title?: string;
  placement?: Placement;
  modelValue?: string;
  showIconName?: boolean;
  placeholder?: string;
  show?: boolean;
}

const props = withDefaults(defineProps<SelectIconProps>(), {
  size: "default",
  disabled: false,
  title: "",
  placement: "bottom",
  modelValue: "",
  showIconName: false,
  placeholder: "",
  show: true
});

const emits = defineEmits<{
  (e: "update:modelValue", value: string): void;
  (e: "change", value: string): void;
}>();

const selectorInput = ref();
const selectorScrollbarRef = ref();
const state: {
  selectorWidth: number;
  popoverVisible: boolean;
  inputFocus: boolean;
  iconSelectorMouseover: boolean;
  fontIconNames: string[];
  inputValue: string;
  prependIcon: string;
  defaultModelValue: string;
  iconKey: number;
} = reactive({
  selectorWidth: 0,
  popoverVisible: false,
  inputFocus: false,
  iconSelectorMouseover: false,
  fontIconNames: [],
  inputValue: "",
  prependIcon: props.modelValue,
  defaultModelValue: props.modelValue || "",
  iconKey: 0 // 给icon标签准备个key，以随时使用 h 函数重新生成元素
});

const onInputFocus = () => {
  state.inputFocus = state.popoverVisible = true;
};

const onInputBlur = () => {
  state.inputFocus = false;
  state.popoverVisible = state.iconSelectorMouseover;
};

const onInputRefresh = () => {
  state.iconKey++;
  state.prependIcon = state.defaultModelValue;
  state.inputValue = "";
  emits("update:modelValue", state.defaultModelValue);
  emits("change", state.defaultModelValue);
};

const renderFontIconNames = computed(() => {
  if (!state.inputValue) return state.fontIconNames;
  let inputValue = state.inputValue.trim().toLowerCase();
  return state.fontIconNames.filter((icon: string) => {
    if (icon.toLowerCase().indexOf(inputValue) !== -1) {
      return icon;
    }
  });
});

const customIcons: { [key: string]: any } = Icons;

// 获取 input 的宽度
const getInputWidth = () => {
  nextTick(() => {
    state.selectorWidth = selectorInput.value.$el.offsetWidth < 260 ? 260 : selectorInput.value.$el.offsetWidth;
  });
};

const popoverVisible = () => {
  state.popoverVisible = state.inputFocus || state.iconSelectorMouseover ? true : false;
};

const elementPlusIconfontNames = () => {
  return new Promise<string[]>((resolve, reject) => {
    nextTick(() => {
      // 定义 list 为一个空数组
      const list: string[] = [];
      const icons = Icons as any;
      for (const i in icons) {
        list.push(`${icons[i].name}`);
      }
      if (list.length > 0) {
        resolve(list);
      } else {
        reject("No ElementPlus Icons");
      }
    });
  });
};

const onIcon = (icon: string) => {
  state.iconSelectorMouseover = state.popoverVisible = false;
  state.iconKey++;
  state.prependIcon = icon;
  state.inputValue = "";
  emits("update:modelValue", icon);
  emits("change", icon);
  nextTick(() => {
    selectorInput.value.blur();
  });
};
watch(
  () => props.modelValue,
  () => {
    state.iconKey++;
    if (props.modelValue != state.prependIcon) state.defaultModelValue = props.modelValue;
    if (props.modelValue == "") state.defaultModelValue = "";
    state.prependIcon = props.modelValue;
  }
);

onMounted(() => {
  getInputWidth();
  useEventListener(document, "click", popoverVisible);
  elementPlusIconfontNames().then(res => {
    state.fontIconNames = res;
  });
  window.onresize = () => {
    getInputWidth();
  };
});

onUnmounted(() => {
  window.onresize = null;
});
</script>

<style scoped lang="scss">
@import "./index";
</style>
