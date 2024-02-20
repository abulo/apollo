<template>
  <div class="home card">
    <img class="home-bg" src="@/assets/images/welcome.png" alt="welcome" />
    <UploadBtn v-model:file="file" :callback="onSuccessCallback" @click="ddd" />

    <UploadImg
      v-model:file="file"
      v-model:attachmentId="file.id"
      width="135px"
      height="135px"
      :file-size="3"
      :callback="onSuccessCallback"
      :preview-teleported="true">
      <template #empty>
        <el-icon><Avatar /></el-icon>
        <span>请上传头像</span>
      </template>
      <template #tip> 头像大小不能超过 3M </template>
    </UploadImg>

    <UploadImgs
      v-model:file-list="fileList"
      :limit="3"
      :callback="onSuccessCallback"
      height="140px"
      width="140px"
      border-radius="50%">
      <template #empty>
        <el-icon><Picture /></el-icon>
        <span>请上传照片</span>
      </template>
      <template #tip> 最多上传 3 张照片 </template>
    </UploadImgs>
  </div>
</template>

<script setup lang="ts" name="home">
import UploadImg from "@/components/Upload/UploadImg.vue";
import UploadImgs from "@/components/Upload/UploadImgs.vue";
import UploadBtn from "@/components/Upload/UploadBtn.vue";
import { SystemAttachment } from "@/api/interface/systemAttachment";
import { ref } from "vue";

// 定义一个 SystemAttachment.ResSystemAttachmentItem 数组
const fileList = ref<SystemAttachment.ResSystemAttachmentItem[]>([]);

const file = ref<SystemAttachment.ResSystemAttachmentItem>({
  id: 0,
  fileName: "",
  fileType: "",
  fileMimeType: "",
  fileSize: 0,
  filePath: "",
  creator: "",
  createTime: "",
  updater: "",
  updateTime: "",
  tenantId: 0,
  uid: 0,
  name: "",
  url: ""
});

const ddd = () => {
  console.log("ddd");
};
const onSuccessCallback = (res: SystemAttachment.ResSystemAttachmentItem) => {
  console.log("callback", res);
};
</script>

<style scoped lang="scss">
@import "./index.scss";
</style>
