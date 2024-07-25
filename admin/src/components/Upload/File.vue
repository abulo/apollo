<template>
  <div class="upload-box">
    <el-upload
      :id="uuid"
      action="#"
      :multiple="false"
      :show-file-list="false"
      :http-request="handleHttpUpload"
      :before-upload="beforeUpload"
      :on-success="uploadSuccess"
      :on-error="uploadError"
      :accept="fileType.join(',')">
      <slot name="text"></slot>
    </el-upload>
  </div>
</template>

<script setup lang="ts" name="UploadFile">
import { ref, computed, inject } from "vue";
import { generateUUID } from "@/utils";
import { ElNotification, formContextKey, formItemContextKey } from "element-plus";
import type { UploadProps, UploadRequestOptions } from "element-plus";
import { CirclePlus } from "@element-plus/icons-vue";

interface UploadFileProps {
  file?: any; // 文件对象
  api: (params: any) => Promise<any>; // 上传文件的 api 方法，一般项目上传都是同一个 api 方法，在组件里直接引入即可 ==> 非必传
  fileSize?: number; // 文件大小限制 ==> 非必传（默认为 5M）
  fileType?: File.MimeType[]; // 文件类型限制 ==> 非必传（默认为 ["image/jpeg", "image/png", "image/gif"]）
  // btnText?: string; // 按钮文字 ==> 非必传（默认为 "上传"）
  callback?: () => void; // 上传成功回调 ==> 非必传
}

// 接受父组件参数
const props = withDefaults(defineProps<UploadFileProps>(), {
  fileSize: 5,
  fileType: () => ["image/jpeg", "image/png", "image/gif"]
  // btnText: "上传"
});

// 生成组件唯一id
const uuid = ref("id-" + generateUUID());

// 获取 el-form 组件上下文
const formContext = inject(formContextKey, void 0);
// 获取 el-form-item 组件上下文
const formItemContext = inject(formItemContextKey, void 0);
// 获取上传按钮的文字
// const btnText = computed(() => {
//   return props.btnText;
// });

/**
 * @description 文件上传
 * @param options upload 所有配置项
 * */
const emit = defineEmits<{
  "update:file": [value: any];
}>();
const handleHttpUpload = async (options: UploadRequestOptions) => {
  let formData = new FormData();
  formData.append("file", options.file);
  try {
    const api = props.api;
    const { data } = await api(formData);
    emit("update:file", data);
    // 调用 el-form 内部的校验方法（可自动校验）
    formItemContext?.prop && formContext?.validateField([formItemContext.prop as string]);
  } catch (error) {
    options.onError(error as any);
  }
};

/**
 * @description 文件上传之前判断
 * @param rawFile 选择的文件
 * */
const beforeUpload: UploadProps["beforeUpload"] = rawFile => {
  const fileSize = rawFile.size / 1024 / 1024 < props.fileSize;
  const fileType = props.fileType.includes(rawFile.type as File.MimeType);

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
        message: `上传文件大小不能超过 ${props.fileSize}M！`,
        type: "warning"
      });
    }, 0);
  return fileType && fileSize;
};

/**
 * @description 文件上传成功
 * */
const uploadSuccess = () => {
  ElNotification({
    title: "温馨提示",
    message: "文件上传成功！",
    type: "success"
  });
  props.callback && props.callback();
};

/**
 * @description 文件上传错误
 * */
const uploadError = () => {
  ElNotification({
    title: "温馨提示",
    message: "文件上传失败，请您重新上传！",
    type: "error"
  });
};
</script>
