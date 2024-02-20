<template>
  <!-- 分页组件 -->
  <el-pagination
    :background="true"
    :current-page="pageable.pageNum"
    :page-size="pageable.pageSize"
    :page-sizes="[5, 10, 25, 50, 100]"
    :total="pageable.total"
    :layout="layout"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"></el-pagination>
</template>

<script setup lang="ts" name="Pagination">
interface Pageable {
  pageNum: number;
  pageSize: number;
  total: number;
}

interface PaginationProps {
  pageable: Pageable;
  layout: string;
  handleSizeChange: (size: number) => void;
  handleCurrentChange: (currentPage: number) => void;
}

defineProps<PaginationProps>();
</script>
