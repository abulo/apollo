<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="应用列表"
      row-key="id"
      :columns="columns"
      :request-api="getPayAppListApi"
      :request-auto="true"
      :pagination="false"
      :search-col="12">
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button v-auth="'app.PayAppCreate'" type="primary" :icon="CirclePlus" @click="handleAdd"> 新增 </el-button>
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
        <el-button v-auth="'app.PayAppUpdate'" type="primary" link :icon="EditPen" @click="handleUpdate(scope.row)">
          编辑
        </el-button>
        <el-button
          v-if="scope.row.deleted === 0"
          v-auth="'app.PayAppDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)">
          删除
        </el-button>
        <el-button
          v-if="scope.row.deleted === 1"
          v-auth="'app.PayAppRecover'"
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
      <el-form ref="refPayAppItemFrom" :model="payAppItemFrom" :rules="rulesPayAppItemFrom" label-width="100px">
        <el-form-item label="应用名称" prop="name">
          <el-input v-model="payAppItemFrom.name" />
        </el-form-item>
        <el-form-item label="开启状态" prop="status">
          <el-radio-group v-model="payAppItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :value="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="支付结果的回调地址" prop="orderNotifyUrl">
          <el-input v-model="payAppItemFrom.orderNotifyUrl" placeholder="请输入支付结果的回调地址" />
        </el-form-item>
        <el-form-item label="退款结果的回调地址" prop="refundNotifyUrl">
          <el-input v-model="payAppItemFrom.refundNotifyUrl" placeholder="请输入退款结果的回调地址" />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="payAppItemFrom.remark" placeholder="请输入备注" />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="payAppItemFrom.remark" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refPayAppItemFrom)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refPayAppItemFrom)">确定</el-button>
        </span>
      </template>
    </el-dialog>
    <el-dialog
      :title="titleChannel"
      v-model="centerMockDialogVisible"
      width="40%"
      destroy-on-close
      align-center
      center
      append-to-body
      draggable
      :lock-scroll="false"
      class="dialog-settings">
      <el-form ref="refPayChannelItemFrom" :model="payChannelItemFrom" :rules="rulesPayChannelMockItemFrom" label-width="180px">
        <el-form-item label="渠道费率" prop="feeRate">
          <el-input v-model="payChannelItemFrom.feeRate" placeholder="请输入渠道费率" clearable>
            <template #append>%</template>
          </el-input>
        </el-form-item>
        <el-form-item label="渠道状态" prop="status">
          <el-radio-group v-model="payChannelItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="payChannelItemFrom.remark" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetChannelForm(refPayChannelItemFrom)">取消</el-button>
          <el-button type="primary" @click="submitChannelForm(refPayChannelItemFrom)" :loading="loadingChannel"> 确定 </el-button>
        </span>
      </template>
    </el-dialog>
    <el-dialog
      :title="titleChannel"
      v-model="centerWeChatDialogVisible"
      width="40%"
      destroy-on-close
      align-center
      center
      append-to-body
      draggable
      :lock-scroll="false"
      class="dialog-settings">
      <el-form ref="refPayChannelItemFrom" :model="payChannelItemFrom" :rules="rulesPayChannelWeChatItemFrom" label-width="180px">
        <el-form-item label="渠道费率" prop="feeRate">
          <el-input v-model="payChannelItemFrom.feeRate" placeholder="请输入渠道费率" clearable>
            <template #append>%</template>
          </el-input>
        </el-form-item>
        <el-form-item label="渠道状态" prop="status">
          <el-radio-group v-model="payChannelItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="公众号 APPID" prop="config.appId">
          <el-input
            v-model="payChannelItemConfigFrom.appId"
            placeholder="请输入公众号 APPID"
            clearable
            :style="{ width: '100%' }" />
        </el-form-item>
        <el-form-item label="商户号" prop="config.mchId">
          <el-input v-model="payChannelItemConfigFrom.mchId" :style="{ width: '100%' }" />
        </el-form-item>
        <el-form-item label="API 版本" prop="config.apiVersion">
          <el-radio-group v-model="payChannelItemConfigFrom.apiVersion">
            <el-radio label="v2">v2</el-radio>
            <el-radio label="v3">v3</el-radio>
          </el-radio-group>
        </el-form-item>
        <div v-if="payChannelItemConfigFrom.apiVersion === 'v2'">
          <el-form-item label-width="180px" label="商户密钥" prop="config.mchKey">
            <el-input
              v-model="payChannelItemConfigFrom.mchKey"
              placeholder="请输入商户密钥"
              clearable
              :style="{ width: '100%' }"
              type="textarea"
              :autosize="{ minRows: 8, maxRows: 8 }" />
          </el-form-item>
          <el-form-item label="apiclient_cert.p12 证书" prop="config.keyContent">
            <el-input
              v-model="payChannelItemConfigFrom.keyContent"
              type="textarea"
              placeholder="请上传 apiclient_cert.p12 证书"
              readonly
              :autosize="{ minRows: 8, maxRows: 8 }"
              :style="{ width: '100%' }" />
          </el-form-item>
          <el-form-item label="">
            <el-upload :limit="1" accept=".p12" action="" :before-upload="p12FileBeforeUpload" :http-request="keyContentUpload">
              <el-button type="primary">点击上传</el-button>
            </el-upload>
          </el-form-item>
        </div>
        <div v-if="payChannelItemConfigFrom.apiVersion === 'v3'">
          <el-form-item label="API V3 密钥" prop="config.apiV3Key">
            <el-input
              v-model="payChannelItemConfigFrom.apiV3Key"
              placeholder="请输入 API V3 密钥"
              clearable
              :style="{ width: '100%' }"
              type="textarea"
              :autosize="{ minRows: 8, maxRows: 8 }" />
          </el-form-item>
          <el-form-item label="apiclient_key.pem 证书" prop="config.privateKeyContent">
            <el-input
              v-model="payChannelItemConfigFrom.privateKeyContent"
              type="textarea"
              placeholder="请上传 apiclient_key.pem 证书"
              readonly
              :autosize="{ minRows: 8, maxRows: 8 }"
              :style="{ width: '100%' }" />
          </el-form-item>
          <el-form-item label="" prop="privateKeyContentFile">
            <el-upload
              ref="privateKeyContentFile"
              :limit="1"
              accept=".pem"
              action=""
              :before-upload="pemFileBeforeUpload"
              :http-request="privateKeyContentUpload">
              <el-button type="primary"> 点击上传 </el-button>
            </el-upload>
          </el-form-item>
          <el-form-item label="apiclient_cert.pem证书" prop="config.privateCertContent">
            <el-input
              v-model="payChannelItemConfigFrom.privateCertContent"
              type="textarea"
              placeholder="请上传apiclient_cert.pem证书"
              readonly
              :autosize="{ minRows: 8, maxRows: 8 }"
              :style="{ width: '100%' }" />
          </el-form-item>
          <el-form-item label="" prop="privateCertContentFile">
            <el-upload
              ref="privateCertContentFile"
              :limit="1"
              accept=".pem"
              action=""
              :before-upload="pemFileBeforeUpload"
              :http-request="privateCertContentUpload">
              <el-button type="primary"> 点击上传 </el-button>
            </el-upload>
          </el-form-item>
        </div>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="payChannelItemFrom.remark" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetChannelForm(refPayChannelItemFrom)">取消</el-button>
          <el-button type="primary" @click="submitChannelForm(refPayChannelItemFrom)" :loading="loadingChannel"> 确定 </el-button>
        </span>
      </template>
    </el-dialog>
    <el-dialog
      :title="titleChannel"
      v-model="centerAliDialogVisible"
      width="40%"
      destroy-on-close
      align-center
      center
      append-to-body
      draggable
      :lock-scroll="false"
      class="dialog-settings">
      <el-form ref="refPayChannelItemFrom" :model="payChannelItemFrom" :rules="rulesPayChannelAliItemFrom" label-width="180px">
        <el-form-item label="渠道费率" prop="feeRate">
          <el-input v-model="payChannelItemFrom.feeRate" placeholder="请输入渠道费率" clearable>
            <template #append>%</template>
          </el-input>
        </el-form-item>
        <el-form-item label="渠道状态" prop="status">
          <el-radio-group v-model="payChannelItemFrom.status">
            <el-radio-button v-for="dict in statusEnum" :key="Number(dict.value)" :label="dict.value">
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="开放平台 APPID" prop="config.appId">
          <el-input v-model="payChannelItemConfigFrom.appId" placeholder="请输入开放平台 APPID" clearable />
        </el-form-item>
        <el-form-item label="网关地址" prop="config.serverUrl">
          <el-radio-group v-model="payChannelItemConfigFrom.serverUrl">
            <el-radio label="https://openapi.alipay.com/gateway.do">线上环境</el-radio>
            <el-radio label="https://openapi-sandbox.dl.alipaydev.com/gateway.do"> 沙箱环境 </el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="算法类型" prop="config.signType">
          <el-radio-group v-model="payChannelItemConfigFrom.signType">
            <el-radio key="RSA2" label="RSA2">RSA2</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="公钥类型" prop="config.mode">
          <el-radio-group v-model="payChannelItemConfigFrom.mode">
            <el-radio key="公钥模式" :label="1">公钥模式</el-radio>
            <el-radio key="证书模式" :label="2">证书模式</el-radio>
          </el-radio-group>
        </el-form-item>
        <div v-if="payChannelItemConfigFrom.mode === 1">
          <el-form-item label="应用私钥" prop="config.privateKey">
            <el-input
              type="textarea"
              :autosize="{ minRows: 8, maxRows: 8 }"
              v-model="payChannelItemConfigFrom.privateKey"
              placeholder="请输入应用私钥"
              clearable
              :style="{ width: '100%' }" />
          </el-form-item>
          <el-form-item label="支付宝公钥" prop="config.alipayPublicKey">
            <el-input
              type="textarea"
              :autosize="{ minRows: 8, maxRows: 8 }"
              v-model="payChannelItemConfigFrom.alipayPublicKey"
              placeholder="请输入支付宝公钥"
              clearable
              :style="{ width: '100%' }" />
          </el-form-item>
        </div>
        <div v-if="payChannelItemConfigFrom.mode === 2">
          <el-form-item label="商户公钥应用证书" prop="config.appCertContent">
            <el-input
              v-model="payChannelItemConfigFrom.appCertContent"
              type="textarea"
              placeholder="请上传商户公钥应用证书"
              readonly
              :autosize="{ minRows: 8, maxRows: 8 }"
              :style="{ width: '100%' }" />
          </el-form-item>
          <el-form-item label="">
            <el-upload
              action=""
              ref="privateKeyContentFile"
              :limit="1"
              :accept="fileAccept"
              :http-request="appCertUpload"
              :before-upload="fileBeforeUploadAli">
              <el-button type="primary"> 点击上传 </el-button>
            </el-upload>
          </el-form-item>
          <el-form-item label="支付宝公钥证书" prop="config.alipayPublicCertContent">
            <el-input
              v-model="payChannelItemConfigFrom.alipayPublicCertContent"
              type="textarea"
              placeholder="请上传支付宝公钥证书"
              readonly
              :autosize="{ minRows: 8, maxRows: 8 }"
              :style="{ width: '100%' }" />
          </el-form-item>
          <el-form-item label="">
            <el-upload
              ref="privateCertContentFile"
              action=""
              :limit="1"
              :accept="fileAccept"
              :before-upload="fileBeforeUploadAli"
              :http-request="alipayPublicCertUpload">
              <el-button type="primary"> 点击上传 </el-button>
            </el-upload>
          </el-form-item>
          <el-form-item label="根证书" prop="config.rootCertContent">
            <el-input
              v-model="payChannelItemConfigFrom.rootCertContent"
              type="textarea"
              placeholder="请上传根证书"
              readonly
              :autosize="{ minRows: 8, maxRows: 8 }"
              :style="{ width: '100%' }" />
          </el-form-item>
          <el-form-item label="">
            <el-upload
              ref="privateCertContentFile"
              :limit="1"
              :accept="fileAccept"
              action=""
              :before-upload="fileBeforeUploadAli"
              :http-request="rootCertUpload">
              <el-button type="primary"> 点击上传 </el-button>
            </el-upload>
          </el-form-item>
        </div>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="payChannelItemFrom.remark" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetChannelForm(refPayChannelItemFrom)">取消</el-button>
          <el-button type="primary" @click="submitChannelForm(refPayChannelItemFrom)" :loading="loadingChannel"> 确定 </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="tsx" name="payApp">
import { ref, reactive, watch } from "vue";
import { ProTableInstance, ColumnProps, SearchProps, RenderScope } from "@/components/ProTable/interface";
import { EditPen, CirclePlus, Delete, Refresh, Check, Close } from "@element-plus/icons-vue";
import ProTable from "@/components/ProTable/index.vue";
import { PayApp } from "@/api/interface/payApp";
import {
  getPayAppListApi,
  deletePayAppApi,
  recoverPayAppApi,
  getPayAppItemApi,
  addPayAppApi,
  updatePayAppApi
} from "@/api/modules/payApp";
import { PayChannel } from "@/api/interface/payChannel";
import { setPayChannelApi, getPayChannelItemApi } from "@/api/modules/payChannel";
import { FormInstance, FormRules, ElNotification } from "element-plus";
import { getIntDictOptions, getStrDictOptions, getDictLabel } from "@/utils/dict";
import { DictTag } from "@/components/DictTag";
import { useHandleData, useHandleSet, useMultiHandleSet } from "@/hooks/useHandleData";
import { HasPermission } from "@/utils/permission";
import type { UploadProps, UploadRequestOptions } from "element-plus";
//加载
const loading = ref(false);
//弹出层标题
const title = ref();
//table数据
const proTable = ref<ProTableInstance>();
//是否显示弹出层
const centerDialogVisible = ref(false);
//数据接口
const payAppItemFrom = ref<PayApp.ResPayAppItem>({
  id: 0, // 应用编号
  name: "", // 应用名称
  status: 0, // 开启状态
  remark: "", // 备注
  orderNotifyUrl: "", // 支付结果的回调地址
  refundNotifyUrl: "", // 退款结果的回调地址
  transferNotifyUrl: "", // 转账结果的回调地址
  deleted: 0, // 删除
  channelList: undefined // 支付渠道列表
});

//校验
const refPayAppItemFrom = ref<FormInstance>();
//校验
const rulesPayAppItemFrom = reactive<FormRules>({
  name: [{ required: true, message: "应用名称不能为空", trigger: "blur" }],
  status: [{ required: true, message: "应用状态不能为空", trigger: "change" }]
});

//状态
const statusEnum = getIntDictOptions("status");
//删除状态
const deletedEnum = getIntDictOptions("delete");
// 支付渠道
const payChannel = getStrDictOptions("payChannel.code");
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasPermission("app.PayAppDelete")
    ? {
        el: "switch",
        span: 2
      }
    : {}
);

const isChannelExists = (channels: Array<string>, channelCode: string) => {
  if (!channels) {
    return false;
  }
  return channels.indexOf(channelCode) !== -1;
};

// 自定义渲染表头（使用tsx语法）
const channelRender = (scope: RenderScope<PayApp.ResPayAppItem>, channelCode: string) => {
  return isChannelExists(scope.row.channelList, channelCode) ? (
    <el-button type="success" circle onClick={() => openChannel(scope.row.id, channelCode)}>
      <el-icon>
        <Check />
      </el-icon>
    </el-button>
  ) : (
    <el-button type="danger" circle onClick={() => openChannel(scope.row.id, channelCode)}>
      <el-icon>
        <Close />
      </el-icon>
    </el-button>
  );
};

const columns: ColumnProps<PayApp.ResPayAppItem>[] = [
  { prop: "id", label: "编号", width: 100 },
  { prop: "name", label: "应用名称", search: { el: "input", span: 2, props: { placeholder: "请输入名称" } } },
  {
    prop: "-",
    label: "支付宝配置",
    _children: [
      {
        prop: "alipay_app",
        label: getDictLabel("payChannel.code", "alipay_app"),
        render: scope => {
          return channelRender(scope, "alipay_app");
        }
      },
      {
        prop: "alipay_pc",
        label: getDictLabel("payChannel.code", "alipay_pc"),
        render: scope => {
          return channelRender(scope, "alipay_pc");
        }
      },
      {
        prop: "alipay_wap",
        label: getDictLabel("payChannel.code", "alipay_wap"),
        render: scope => {
          return channelRender(scope, "alipay_wap");
        }
      },
      {
        prop: "alipay_qr",
        label: getDictLabel("payChannel.code", "alipay_qr"),
        render: scope => {
          return channelRender(scope, "alipay_qr");
        }
      },
      {
        prop: "alipay_bar",
        label: getDictLabel("payChannel.code", "alipay_bar"),
        render: scope => {
          return channelRender(scope, "alipay_bar");
        }
      }
    ]
  },
  {
    prop: "-",
    label: "微信配置",
    _children: [
      {
        prop: "wx_app",
        label: getDictLabel("payChannel.code", "wx_app"),
        render: scope => {
          return channelRender(scope, "wx_app");
        }
      },
      {
        prop: "wx_lite",
        label: getDictLabel("payChannel.code", "wx_lite"),
        render: scope => {
          return channelRender(scope, "wx_lite");
        }
      },
      {
        prop: "wx_pub",
        label: getDictLabel("payChannel.code", "wx_pub"),
        render: scope => {
          return channelRender(scope, "wx_pub");
        }
      }
    ]
  },
  {
    prop: "-",
    label: "模拟配置",
    _children: [
      {
        prop: "mock",
        label: getDictLabel("payChannel.code", "mock"),
        render: scope => {
          return channelRender(scope, "mock");
        }
      }
    ]
  },
  { prop: "status", label: "开启状态", tag: true, enum: statusEnum, search: { el: "select", span: 2 } },
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
    isShow: HasPermission("app.PayAppUpdate", "app.PayAppDelete", "app.PayAppRecover")
  }
];
// 重置数据
const reset = () => {
  loading.value = false;
  payAppItemFrom.value = {
    id: 0, // 应用编号
    name: "", // 应用名称
    status: 0, // 开启状态
    remark: "", // 备注
    orderNotifyUrl: "", // 支付结果的回调地址
    refundNotifyUrl: "", // 退款结果的回调地址
    transferNotifyUrl: "", // 转账结果的回调地址
    deleted: 0, // 删除
    channelList: undefined // 支付渠道列表
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
    const data = payAppItemFrom.value as unknown as PayApp.ResPayAppItem;
    if (data.id !== 0) {
      await useHandleSet(updatePayAppApi, data.id, data, "修改应用");
    } else {
      await useHandleData(addPayAppApi, data, "添加应用");
    }
    resetForm(formEl);
    loading.value = false;
    proTable.value?.getTableList();
  });
};

// 删除按钮
const handleDelete = async (row: PayApp.ResPayAppItem) => {
  await useHandleData(deletePayAppApi, Number(row.id), "删除应用");
  proTable.value?.getTableList();
};

// 恢复按钮
const handleRecover = async (row: PayApp.ResPayAppItem) => {
  await useHandleData(recoverPayAppApi, Number(row.id), "恢复应用");
  proTable.value?.getTableList();
};

// 添加按钮
const handleAdd = () => {
  title.value = "新增应用";
  centerDialogVisible.value = true;
  reset();
};

// 编辑按钮
const handleUpdate = async (row: PayApp.ResPayAppItem) => {
  title.value = "编辑应用";
  centerDialogVisible.value = true;
  reset();
  const { data } = await getPayAppItemApi(Number(row.id));
  payAppItemFrom.value = data;
};

//加载
const loadingChannel = ref(false);
//弹出层标题
const titleChannel = ref();
//校验
const refPayChannelItemFrom = ref<FormInstance>();
//是否显示弹出层
const centerAliDialogVisible = ref(false);
const centerWeChatDialogVisible = ref(false);
const centerMockDialogVisible = ref(false);
const fileAccept = ref(".crt");

//支付宝校验
const rulesPayChannelAliItemFrom = reactive<FormRules>({
  feeRate: [{ required: true, message: "请输入渠道费率", trigger: "blur" }],
  status: [{ required: true, message: "渠道状态不能为空", trigger: "blur" }],
  "config.appId": [{ required: true, message: "请输入开放平台上创建的应用的 ID", trigger: "blur" }],
  "config.serverUrl": [{ required: true, message: "请传入网关地址", trigger: "blur" }],
  "config.signType": [{ required: true, message: "请传入签名算法类型", trigger: "blur" }],
  "config.mode": [{ required: true, message: "公钥类型不能为空", trigger: "blur" }],
  "config.privateKey": [{ required: true, message: "请输入商户私钥", trigger: "blur" }],
  "config.alipayPublicKey": [{ required: true, message: "请输入支付宝公钥字符串", trigger: "blur" }],
  "config.appCertContent": [{ required: true, message: "请上传商户公钥应用证书", trigger: "blur" }],
  "config.alipayPublicCertContent": [{ required: true, message: "请上传支付宝公钥证书", trigger: "blur" }],
  "config.rootCertContent": [{ required: true, message: "请上传指定根证书", trigger: "blur" }]
});
// 微信
const rulesPayChannelWeChatItemFrom = reactive<FormRules>({
  feeRate: [{ required: true, message: "请输入渠道费率", trigger: "blur" }],
  status: [{ required: true, message: "渠道状态不能为空", trigger: "blur" }],
  "config.mchId": [{ required: true, message: "请传入商户号", trigger: "blur" }],
  "config.appId": [{ required: true, message: "请输入公众号APPID", trigger: "blur" }],
  "config.apiVersion": [{ required: true, message: "API版本不能为空", trigger: "blur" }],
  "config.mchKey": [{ required: true, message: "请输入商户密钥", trigger: "blur" }],
  "config.keyContent": [{ required: true, message: "请上传 apiclient_cert.p12 证书", trigger: "blur" }],
  "config.privateKeyContent": [{ required: true, message: "请上传 apiclient_key.pem 证书", trigger: "blur" }],
  "config.privateCertContent": [{ required: true, message: "请上传 apiclient_cert.pem证 书", trigger: "blur" }],
  "config.apiV3Key": [{ required: true, message: "请上传 api V3 密钥值", trigger: "blur" }]
});
// 模拟
const rulesPayChannelMockItemFrom = reactive<FormRules>({
  feeRate: [{ required: true, message: "请输入渠道费率", trigger: "blur" }],
  status: [{ required: true, message: "渠道状态不能为空", trigger: "blur" }]
});

const payChannelItemConfigFrom = ref<PayChannel.ResPayChannelConfig>({});
// 构造数据绑定
const payChannelItemFrom = ref<PayChannel.ResPayChannelItem>({
  id: 0, // 渠道编号
  code: "", // 渠道编码
  status: 0, // 开启状态
  remark: "", // 备注
  feeRate: 0, // 渠道费率，单位：百分比
  appId: 0, // 应用编号
  config: {}, // 支付渠道配置
  deleted: 0, // 删除
  tenantId: 0 // 租户
});

const openChannel = async (id: number, channelCode: string) => {
  loadingChannel.value = false;
  centerAliDialogVisible.value = false;
  centerWeChatDialogVisible.value = false;
  centerMockDialogVisible.value = false;
  const { data } = await getPayChannelItemApi(id, channelCode);
  payChannelItemFrom.value = data;
  payChannelItemFrom.value.code = channelCode;
  payChannelItemFrom.value.appId = id;
  if (channelCode.indexOf("alipay_") === 0) {
    centerAliDialogVisible.value = true;
  }
  if (channelCode.indexOf("wx_") === 0) {
    centerWeChatDialogVisible.value = true;
    return;
  }
  if (channelCode.indexOf("mock") === 0) {
    centerMockDialogVisible.value = true;
  }
  if (channelCode.indexOf("wallet") === 0) {
    centerMockDialogVisible.value = true;
  }
  titleChannel.value = "编辑支付渠道";
  const config = data.config as unknown as PayChannel.ResPayChannelConfig;
  // 遍历 config的值赋值给payChannelItemConfigFrom
  for (const key in config) {
    if (Object.prototype.hasOwnProperty.call(config, key)) {
      payChannelItemConfigFrom.value[key] = config[key];
    }
  }
};

const fileBeforeUploadAli: UploadProps["beforeUpload"] = rawFile => {
  fileAccept.value = ".crt";
  fileBeforeUpload(rawFile);
};

const p12FileBeforeUpload: UploadProps["beforeUpload"] = rawFile => {
  fileAccept.value = ".p12";
  fileBeforeUpload(rawFile);
};

const pemFileBeforeUpload: UploadProps["beforeUpload"] = rawFile => {
  fileAccept.value = ".pem";
  fileBeforeUpload(rawFile);
};
/**
 * @description 文件上传之前判断
 * @param rawFile 选择的文件
 * */
const fileBeforeUpload: UploadProps["beforeUpload"] = rawFile => {
  const fileSize = rawFile.size / 1024 / 1024 < 2;
  const fileType = rawFile.name.endsWith(fileAccept.value);
  if (!fileType)
    ElNotification({
      title: "温馨提示",
      message: "上传文件不符合所需的格式！",
      type: "warning"
    });
  if (!fileSize)
    setTimeout(() => {
      ElNotification({
        title: "温馨提示",
        message: "上传文件大小不能超过2M！",
        type: "warning"
      });
    }, 0);
  return fileType && fileSize;
};

const appCertUpload = async (event: UploadRequestOptions) => {
  const readFile = new FileReader();
  readFile.onload = (e: any) => {
    payChannelItemConfigFrom.value.appCertContent = e.target.result as string;
  };
  readFile.readAsText(event.file);
};

const alipayPublicCertUpload = async (event: UploadRequestOptions) => {
  const readFile = new FileReader();
  readFile.onload = (e: any) => {
    payChannelItemConfigFrom.value.alipayPublicCertContent = e.target.result as string;
  };
  readFile.readAsText(event.file);
};

const rootCertUpload = async (event: UploadRequestOptions) => {
  const readFile = new FileReader();
  readFile.onload = (e: any) => {
    payChannelItemConfigFrom.value.rootCertContent = e.target.result as string;
  };
  readFile.readAsText(event.file);
};

/**
 * 读取 apiclient_key.pem 到 privateKeyContent 字段
 */
const privateKeyContentUpload = async (event: UploadRequestOptions) => {
  const readFile = new FileReader();
  readFile.onload = (e: any) => {
    payChannelItemConfigFrom.value.privateKeyContent = e.target.result as string;
  };
  readFile.readAsText(event.file);
};

/**
 * 读取 apiclient_cert.pem 到 privateCertContent 字段
 */
const privateCertContentUpload = async (event: UploadRequestOptions) => {
  const readFile = new FileReader();
  readFile.onload = (e: any) => {
    payChannelItemConfigFrom.value.privateCertContent = e.target.result as string;
  };
  readFile.readAsText(event.file);
};

/**
 * 读取 apiclient_cert.p12 到 keyContent 字段
 */
const keyContentUpload = async (event: UploadRequestOptions) => {
  const readFile = new FileReader();
  readFile.onload = (e: any) => {
    payChannelItemConfigFrom.value.keyContent = e.target.result.split(",")[1] as string;
  };
  readFile.readAsDataURL(event.file); // 读成 base64
};

const resetChannelForm = (formEl: FormInstance | undefined) => {
  centerAliDialogVisible.value = false;
  centerWeChatDialogVisible.value = false;
  centerMockDialogVisible.value = false;
  if (!formEl) return;
  formEl.resetFields();
};
const submitChannelForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loadingChannel.value = true;
    payChannelItemFrom.value.feeRate = Number(payChannelItemFrom.value.feeRate);
    const data = payChannelItemFrom.value as unknown as PayChannel.ResPayChannelItem;
    await useMultiHandleSet(setPayChannelApi, data.appId, data.code, data, "支付设置");
    resetChannelForm(formEl);
    loadingChannel.value = false;
    proTable.value?.getTableList();
  });
};
watch(
  payChannelItemConfigFrom,
  value => {
    payChannelItemFrom.value.config = value;
  },
  {
    deep: true // 深度监听的参数
  }
);
</script>

<style lang="scss">
@import "@/styles/custom.scss";
</style>
