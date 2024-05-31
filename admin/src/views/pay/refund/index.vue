<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="退款订单列表"
      row-key="id"
      :columns="columns"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12">
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="delete" :value="scope.row.deleted" />
      </template>
      <!-- 支付金额 -->
      <template #payPrice="scope">
        <span>￥{{ parseFloat(String(scope.row.payPrice / 100)).toFixed(2) }}</span>
      </template>
      <!-- 退款金额 -->
      <template #refundPrice="scope">
        <span>￥{{ parseFloat(String(scope.row.refundPrice / 100)).toFixed(2) }}</span>
      </template>
      <!-- 退款订单号 -->
      <template #customOrder="scope">
        <p class="order-font"><el-tag size="small">商户</el-tag> {{ scope.row.merchantRefundId }}</p>
        <p class="order-font"><el-tag size="small" type="warning">退款</el-tag> {{ scope.row.no }}</p>
        <p class="order-font" v-if="scope.row.channelRefundNo">
          <el-tag size="small" type="success">渠道</el-tag> {{ scope.row.channelRefundNo }}
        </p>
      </template>
      <!-- 支付订单号 -->
      <template #merchantOrder="scope">
        <p class="order-font"><el-tag size="small">商户</el-tag> {{ scope.row.merchantOrderId }}</p>
        <p class="order-font"><el-tag size="small" type="success">渠道</el-tag> {{ scope.row.channelOrderNo }}</p>
      </template>
      <template #status="scope">
        <DictTag type="pay.refundStatus" :value="scope.row.status" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button type="primary" link :icon="View" @click="handleItem(scope.row)"> 查看 </el-button>
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'refund.PayRefundDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'refund.PayRefundRecover'"
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
        <el-descriptions-item label="商户退款单号">
          <el-tag size="small">{{ payRefundItemFrom.merchantRefundId }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="渠道退款单号">
          <el-tag type="success" size="small" v-if="payRefundItemFrom.channelRefundNo">{{
            payRefundItemFrom.channelRefundNo
          }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="商户支付单号">
          <el-tag size="small">{{ payRefundItemFrom.merchantOrderId }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="渠道支付单号">
          <el-tag type="success" size="small">{{ payRefundItemFrom.channelOrderNo }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="应用编号">{{ payRefundItemFrom.appId }}</el-descriptions-item>
        <el-descriptions-item label="应用名称">{{
          payAppList.find(item => item.id === payRefundItemFrom.appId)?.name
        }}</el-descriptions-item>
        <el-descriptions-item label="支付金额">
          <el-tag type="success" size="small"> ￥{{ (payRefundItemFrom.payPrice / 100.0).toFixed(2) }} </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="退款金额">
          <el-tag size="mini" type="danger"> ￥{{ (payRefundItemFrom.refundPrice / 100.0).toFixed(2) }} </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="退款状态">
          <DictTag type="pay.refundStatus" :value="payRefundItemFrom.status" />
        </el-descriptions-item>
        <el-descriptions-item label="退款时间">
          {{ payRefundItemFrom.successTime }}
        </el-descriptions-item>
        <el-descriptions-item label="创建时间">
          {{ payRefundItemFrom.createTime }}
        </el-descriptions-item>
        <el-descriptions-item label="更新时间">
          {{ payRefundItemFrom.updateTime }}
        </el-descriptions-item>
      </el-descriptions>
      <!-- 分割线 -->
      <el-divider />
      <el-descriptions :column="2" border>
        <el-descriptions-item label="退款渠道">
          <DictTag type="payChannel.code" :value="payRefundItemFrom.channelCode" />
        </el-descriptions-item>
        <el-descriptions-item label="退款原因">{{ payRefundItemFrom.reason }}</el-descriptions-item>
        <el-descriptions-item label="退款 IP">{{ payRefundItemFrom.userIp }}</el-descriptions-item>
        <el-descriptions-item label="通知 URL">{{ payRefundItemFrom.notifyUrl }}</el-descriptions-item>
      </el-descriptions>
      <!-- 分割线 -->
      <el-divider />
      <el-descriptions :column="2" border>
        <el-descriptions-item label="渠道错误码">
          {{ payRefundItemFrom.channelErrorCode }}
        </el-descriptions-item>
        <el-descriptions-item label="渠道错误码描述">
          {{ payRefundItemFrom.channelErrorMsg }}
        </el-descriptions-item>
      </el-descriptions>
      <el-descriptions :column="1" direction="vertical" border>
        <el-descriptions-item label="支付通道异步回调内容">
          {{ payRefundItemFrom.channelNotifyData }}
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>
<script setup lang="ts" name="payRefund">
import { ref, reactive, onMounted } from "vue";
import { ProTableInstance, ColumnProps, SearchProps } from "@/components/ProTable/interface";
import { View, CirclePlus, Delete, Refresh } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { PayRefund } from "@/api/interface/payRefund";
import { getPayRefundListApi, deletePayRefundApi, recoverPayRefundApi, getPayRefundItemApi } from "@/api/modules/payRefund";
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
const payRefundItemFrom = ref<PayRefund.ResPayRefundItem>({
  id: 0, // 支付退款编号
  no: "", // 退款单号
  appId: 0, // 应用编号
  channelId: 0, // 渠道编号
  channelCode: "", // 渠道编码
  orderId: 0, // 支付订单编号 pay_order 表id
  orderNo: "", // 支付订单 no
  merchantOrderId: "", // 商户订单编号（商户系统生成）
  merchantRefundId: "", // 商户退款订单号（商户系统生成）
  notifyUrl: "", // 异步通知商户地址
  status: 0, // 退款状态
  payPrice: 0, // 支付金额,单位分
  refundPrice: 0, // 退款金额,单位分
  reason: "", // 退款原因
  userIp: undefined, // ip
  channelOrderNo: "", // 渠道订单号，pay_order 中的 channel_order_no 对应
  channelRefundNo: undefined, // 渠道退款单号，渠道返
  successTime: undefined, // 退款成功时间
  channelErrorCode: undefined, // 渠道调用报错时，错误码
  channelErrorMsg: undefined, // 渠道调用报错时，错误信息
  channelNotifyData: undefined, // 支付渠道异步通知的内容
  deleted: 0, // 删除
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined // 更新时间
});

const payAppList = ref<PayApp.ResPayAppItem[]>([]);
//删除状态
const deletedEnum = getIntDictOptions("delete");
//退款状态
const payRefundStatusEnum = getIntDictOptions("pay.refundStatus");
// 支付渠道
const payChannel = getStrDictOptions("payChannel.code");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("refund.PayRefundDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);

const columns: ColumnProps<PayRefund.ResPayRefundItem>[] = [
  { prop: "id", label: "编号" },
  {
    prop: "createTime",
    label: "创建时间",
    search: {
      el: "date-picker",
      span: 4,
      props: { type: "datetimerange", valueFormat: "YYYY-MM-DD HH:mm:ss" }
    }
  },
  { prop: "payPrice", label: "支付金额" },
  { prop: "refundPrice", label: "退款金额" },
  { prop: "customOrder", label: "退款订单号" },
  { prop: "merchantOrder", label: "支付订单号" },
  { prop: "status", label: "退款状态", tag: true, enum: payRefundStatusEnum, search: { el: "select", span: 2 } },
  { prop: "channelCode", label: "渠道编码", tag: true, enum: payChannel, search: { el: "select", span: 2 } },
  {
    prop: "appId",
    label: "应用名称",
    enum: payAppList,
    fieldNames: { label: "name", value: "id" },
    search: { el: "select", span: 2 }
  },
  {
    prop: "orderId",
    label: "支付订单号",
    isShow: false,
    search: { el: "input", span: 2, props: { placeholder: "请输入支付订单号" } }
  },
  {
    prop: "merchantOrderId",
    label: "商户订单号",
    isShow: false,
    search: { el: "input", span: 2, props: { placeholder: "请输入商户订单号" } }
  },
  {
    prop: "channelOrderNo",
    label: "渠道订单号",
    isShow: false,
    search: { el: "input", span: 2, props: { placeholder: "请输入渠道订单号" } }
  },
  { prop: "deleted", label: "删除", tag: true, enum: deletedEnum, search: deleteSearch },
  {
    prop: "operation",
    label: "操作",
    width: 150,
    fixed: "right",
    isShow: HasPermission("refund.PayRefundUpdate", "refund.PayRefundDelete", "refund.PayRefundRecover")
  }
];

// 重置数据
const reset = () => {
  loading.value = false;
  payRefundItemFrom.value = {
    id: 0, // 支付退款编号
    no: "", // 退款单号
    appId: 0, // 应用编号
    channelId: 0, // 渠道编号
    channelCode: "", // 渠道编码
    orderId: 0, // 支付订单编号 pay_order 表id
    orderNo: "", // 支付订单 no
    merchantOrderId: "", // 商户订单编号（商户系统生成）
    merchantRefundId: "", // 商户退款订单号（商户系统生成）
    notifyUrl: "", // 异步通知商户地址
    status: 0, // 退款状态
    payPrice: 0, // 支付金额,单位分
    refundPrice: 0, // 退款金额,单位分
    reason: "", // 退款原因
    userIp: undefined, // ip
    channelOrderNo: "", // 渠道订单号，pay_order 中的 channel_order_no 对应
    channelRefundNo: undefined, // 渠道退款单号，渠道返
    successTime: undefined, // 退款成功时间
    channelErrorCode: undefined, // 渠道调用报错时，错误码
    channelErrorMsg: undefined, // 渠道调用报错时，错误信息
    channelNotifyData: undefined, // 支付渠道异步通知的内容
    deleted: 0, // 删除
    tenantId: 0, // 租户
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined // 更新时间
  };
};

// 删除按钮
const handleDelete = async (row: PayRefund.ResPayRefundItem) => {
  await useHandleData(deletePayRefundApi, Number(row.id), "删除退款订单");
  proTable.value?.getTableList();
};
// 恢复按钮
const handleRecover = async (row: PayRefund.ResPayRefundItem) => {
  await useHandleData(recoverPayRefundApi, Number(row.id), "恢复退款订单");
  proTable.value?.getTableList();
};

const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  newParams.createTime && (newParams.beginCreateTimeTime = newParams.createTime[0]);
  newParams.createTime && (newParams.finishCreateTimeTime = newParams.createTime[1]);
  delete newParams.createTime;
  return getPayRefundListApi(newParams);
};

// 编辑按钮
const handleItem = async (row: PayRefund.ResPayRefundItem) => {
  title.value = "编辑退款订单";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getPayRefundItemApi(Number(row.id));
  payRefundItemFrom.value = data;
};

onMounted(async () => {
  const { data } = await getPayAppListApi();
  payAppList.value = data;
});
async () => {
  const { data } = await getPayAppListApi();
  payAppList.value = data;
};
</script>
<style lang="scss">
@import "@/styles/custom.scss";
</style>
