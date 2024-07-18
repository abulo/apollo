<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="支付订单列表"
      row-key="id"
      :columns="columns"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 支付金额 -->
      <template #price="scope"> ￥{{ parseFloat(String(scope.row.price / 100)).toFixed(2) }} </template>
      <!-- 退款金额 -->
      <template #refundPrice="scope"> ￥{{ parseFloat(String(scope.row.refundPrice / 100)).toFixed(2) }} </template>
      <!-- 手续金额 -->
      <template #channelFeePrice="scope"> ￥{{ parseFloat(String(scope.row.channelFeePrice / 100)).toFixed(2) }} </template>
      <!-- 订单信息 -->
      <template #customOrder="scope">
        <p class="order-font"><el-tag size="small"> 商户</el-tag> {{ scope.row.merchantOrderId }}</p>
        <p class="order-font" v-if="scope.row.no"><el-tag size="small" type="warning">支付</el-tag> {{ scope.row.no }}</p>
        <p class="order-font" v-if="scope.row.channelOrderNo">
          <el-tag size="small" type="success">渠道</el-tag> {{ scope.row.channelOrderNo }}
        </p>
      </template>
      <template #status="scope">
        <DictTag type="pay.orderStatus" :value="scope.row.status" />
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button type="primary" link :icon="View" @click="handleItem(scope.row)"> 查看 </el-button>
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'order.PayOrderDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'order.PayOrderRecover'"
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
      <el-descriptions :column="2" border>
        <el-descriptions-item label="商户单号">
          <el-tag size="small">{{ payOrderItemFrom.merchantOrderId }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="支付单号">
          <el-tag type="warning" size="small" v-if="payOrderItemFrom.no">{{ payOrderItemFrom.no }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="应用编号">{{ payOrderItemFrom.appId }}</el-descriptions-item>
        <el-descriptions-item label="应用名称">
          {{ payAppList.find(item => item.id === payOrderItemFrom.appId)?.name }}
        </el-descriptions-item>
        <el-descriptions-item label="支付状态">
          <DictTag type="pay.orderStatus" :value="payOrderItemFrom.status" />
        </el-descriptions-item>
        <el-descriptions-item label="支付金额">
          <el-tag type="success" size="small">￥{{ (payOrderItemFrom.price / 100.0).toFixed(2) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="手续费">
          <el-tag type="warning" size="small"> ￥{{ (payOrderItemFrom.channelFeePrice / 100.0).toFixed(2) }} </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="手续费比例">
          {{ (payOrderItemFrom.channelFeeRate / 100.0).toFixed(2) }}%
        </el-descriptions-item>
        <el-descriptions-item label="支付时间">
          {{ payOrderItemFrom.successTime }}
        </el-descriptions-item>
        <el-descriptions-item label="失效时间">
          {{ payOrderItemFrom.expireTime }}
        </el-descriptions-item>
        <el-descriptions-item label="创建时间">
          {{ payOrderItemFrom.createTime }}
        </el-descriptions-item>
        <el-descriptions-item label="更新时间">
          {{ payOrderItemFrom.updateTime }}
        </el-descriptions-item>
      </el-descriptions>
      <!-- 分割线 -->
      <el-divider />
      <el-descriptions :column="2" border>
        <el-descriptions-item label="商品标题">{{ payOrderItemFrom.subject }}</el-descriptions-item>
        <el-descriptions-item label="商品描述">{{ payOrderItemFrom.body }}</el-descriptions-item>
        <el-descriptions-item label="支付渠道">
          <DictTag type="payChannel.code" :value="payOrderItemFrom.channelCode" />
        </el-descriptions-item>
        <el-descriptions-item label="支付 IP">{{ payOrderItemFrom.userIp }}</el-descriptions-item>
        <el-descriptions-item label="渠道单号">
          <el-tag type="success" v-if="payOrderItemFrom.channelOrderNo">
            {{ payOrderItemFrom.channelOrderNo }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="渠道用户">{{ payOrderItemFrom.channelUserId }}</el-descriptions-item>
        <el-descriptions-item label="退款金额">
          <el-tag type="danger"> ￥{{ (payOrderItemFrom.refundPrice / 100.0).toFixed(2) }} </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="通知 URL">{{ payOrderItemFrom.notifyUrl }}</el-descriptions-item>
      </el-descriptions>
      <el-divider />
      <el-descriptions :column="1" direction="vertical" border>
        <el-descriptions-item label="支付通道异步回调内容">
          <el-text>{{ payOrderItemFrom?.extension?.channelNotifyData }}</el-text>
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="payOrder">
import { ref, reactive, onMounted } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { View, CirclePlus, Delete, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { PayOrder } from "@/api/interface/payOrder";
import { getPayOrderListApi, deletePayOrderApi, recoverPayOrderApi, getPayOrderItemApi } from "@/api/modules/payOrder";
import { PayApp } from "@/api/interface/payApp";
import { getPayAppListApi } from "@/api/modules/payApp";
import { FormInstance, FormRules } from "element-plus";
import { getIntDictOptions, getStrDictOptions } from "@/utils/dict";
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
const payOrderItemFrom = ref<PayOrder.ResPayOrderItem>({
  id: 0, // 支付订单编号
  appId: 0, // 应用编号
  channelId: 0, // 渠道编号
  channelCode: "", // 渠道编码
  merchantOrderId: "", // 商户订单编号
  subject: "", // 商品标题
  body: "", // 商品描述
  notifyUrl: "", // 异步通知地址
  price: 0, // 支付金额，单位：分
  channelFeeRate: 0, // 渠道手续费，单位：百分比
  channelFeePrice: 0, // 渠道手续金额，单位：分
  status: 0, // 支付状态
  userIp: undefined, // 用户 IP
  expireTime: "", // 订单失效时间
  successTime: undefined, // 订单支付成功时间
  extensionId: undefined, // 支付成功的订单拓展单编号
  no: undefined, // 支付订单号
  refundPrice: 0, // 退款总金额，单位：分
  channelUserId: undefined, // 渠道用户编号
  channelOrderNo: undefined, // 渠道订单号
  deleted: 0, // 删除
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
  extension: undefined
});
const payAppList = ref<PayApp.ResPayAppItem[]>([]);

//删除状态
const deletedEnum = getIntDictOptions("delete");
//支付状态
const payOrderStatusEnum = getIntDictOptions("pay.orderStatus");
// 支付渠道
const payChannel = getStrDictOptions("payChannel.code");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("order.PayOrderDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);

const columns: ColumnProps<PayOrder.ResPayOrderItem>[] = [
  { prop: "id", label: "编号", fixed: "left" },
  {
    prop: "createTime",
    label: "创建时间",
    search: {
      el: "date-picker",
      span: 4,
      props: { type: "datetimerange", valueFormat: "YYYY-MM-DD HH:mm:ss" }
    }
  },
  { prop: "price", label: "支付金额" },
  { prop: "refundPrice", label: "退款金额" },
  { prop: "channelFeePrice", label: "手续金额" },
  { prop: "customOrder", label: "订单号" },
  { prop: "status", label: "支付状态", tag: true, enum: payOrderStatusEnum, search: { el: "select", span: 2 } },
  { prop: "channelCode", label: "渠道编码", tag: true, enum: payChannel, search: { el: "select", span: 2 } },
  {
    prop: "merchantOrderId",
    label: "商户单号",
    isShow: false,
    search: { el: "input", span: 2, props: { placeholder: "请输入商户单号" } }
  },
  { prop: "no", label: "支付单号", isShow: false, search: { el: "input", span: 2, props: { placeholder: "请输入支付单号" } } },
  {
    prop: "channelOrderNo",
    label: "渠道单号",
    isShow: false,
    search: { el: "input", span: 2, props: { placeholder: "请输入渠道单号" } }
  },
  { prop: "successTime", label: "支付时间" },
  {
    prop: "appId",
    label: "应用名称",
    enum: payAppList,
    fieldNames: { label: "name", value: "id" },
    search: { el: "select", span: 2 }
  },
  { prop: "subject", label: "商品标题" },
  { prop: "deleted", label: "删除", tag: true, enum: deletedEnum, search: deleteSearch },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission("order.PayOrderDelete", "order.PayOrderRecover")
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  payOrderItemFrom.value = {
    id: 0, // 支付订单编号
    appId: 0, // 应用编号
    channelId: 0, // 渠道编号
    channelCode: "", // 渠道编码
    merchantOrderId: "", // 商户订单编号
    subject: "", // 商品标题
    body: "", // 商品描述
    notifyUrl: "", // 异步通知地址
    price: 0, // 支付金额，单位：分
    channelFeeRate: 0, // 渠道手续费，单位：百分比
    channelFeePrice: 0, // 渠道手续金额，单位：分
    status: 0, // 支付状态
    userIp: undefined, // 用户 IP
    expireTime: "", // 订单失效时间
    successTime: undefined, // 订单支付成功时间
    extensionId: undefined, // 支付成功的订单拓展单编号
    no: undefined, // 支付订单号
    refundPrice: 0, // 退款总金额，单位：分
    channelUserId: undefined, // 渠道用户编号
    channelOrderNo: undefined, // 渠道订单号
    deleted: 0, // 删除
    tenantId: 0, // 租户
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined, // 更新时间
    extension: undefined
  };
};

// 删除按钮
const handleDelete = async (row: PayOrder.ResPayOrderItem) => {
  await useHandleData(deletePayOrderApi, Number(row.id), "删除支付订单");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: PayOrder.ResPayOrderItem) => {
  await useHandleData(recoverPayOrderApi, Number(row.id), "恢复支付订单");
  proTable.value?.getTableList();
};
// 编辑按钮
const handleItem = async (row: PayOrder.ResPayOrderItem) => {
  title.value = "支付订单";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getPayOrderItemApi(Number(row.id));
  payOrderItemFrom.value = data;
};

const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.createTime && (newParams.beginCreateTimeTime = newParams.createTime[0]);
  newParams.createTime && (newParams.finishCreateTimeTime = newParams.createTime[1]);
  delete newParams.createTime;
  return getPayOrderListApi(newParams);
};

onMounted(async () => {
  const { data } = await getPayAppListApi();
  payAppList.value = data;
});
</script>
<style lang="scss">
@import "@/styles/custom";
</style>
