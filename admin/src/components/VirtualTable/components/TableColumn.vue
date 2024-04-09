<template>
  <RenderTableColumn v-bind="column" />
</template>

<script setup lang="tsx" name="TableColumn">
import { inject, ref, useSlots } from "vue";
import { ColumnProps, RenderScope, HeaderRenderScope } from "@/components/VirtualTable/interface";
import { filterEnum, formatValue, handleProp, handleRowAccordingToProp } from "@/utils";

defineProps<{ column: ColumnProps }>();

const slots = useSlots();

const enumMap = inject("enumMap", ref(new Map()));

// 渲染表格数据
const renderCellData = (item: ColumnProps, scope: RenderScope<any>) => {
  return enumMap.value.get(item.field) && item.isFilterEnum
    ? filterEnum(handleRowAccordingToProp(scope.row, item.field!), enumMap.value.get(item.field)!, item.fieldNames)
    : formatValue(handleRowAccordingToProp(scope.row, item.field!));
};

// 获取 tag 类型
const getTagType = (item: ColumnProps, scope: RenderScope<any>) => {
  return filterEnum(handleRowAccordingToProp(scope.row, item.field!), enumMap.value.get(item.field), item.fieldNames, "tag");
};

const RenderTableColumn = (item: ColumnProps) => {
  return (
    <>
      {item.isShow && (
        <vxe-column {...item} align={item.align ?? "center"}>
          {{
            default: (scope: RenderScope<any>) => {
              if (item._children) return item._children.map(child => RenderTableColumn(child));
              if (item.render) return item.render(scope);
              if (slots[handleProp(item.field!)]) return slots[handleProp(item.field!)]!(scope);
              if (item.tag) return <el-tag type={getTagType(item, scope)}>{renderCellData(item, scope)}</el-tag>;
              return renderCellData(item, scope);
            },
            header: (scope: HeaderRenderScope<any>) => {
              if (item.headerRender) return item.headerRender(scope);
              if (slots[`${handleProp(item.field!)}Header`]) return slots[`${handleProp(item.field!)}Header`]!(scope);
              return item.title;
            }
          }}
        </vxe-column>
      )}
    </>
  );
};
</script>
