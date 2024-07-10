<template>
  <el-form ref="loginFormRef" :model="loginForm" :rules="loginRules" size="large">
    <el-form-item prop="username">
      <el-input v-model="loginForm.username" placeholder="用户名">
        <template #prefix>
          <el-icon class="el-input__icon"><User /></el-icon>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item prop="password">
      <el-input v-model="loginForm.password" type="password" placeholder="密码" show-password autocomplete="new-password">
        <template #prefix>
          <el-icon class="el-input__icon"><Lock /></el-icon>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item prop="verifyCode">
      <el-input v-model="loginForm.captchaCode" placeholder="验证码">
        <template #prefix>
          <el-icon class="el-input__icon"><Key /></el-icon>
        </template>
        <template #append>
          <el-image class="captchaImg" :src="resCaptcha.captchaImage" fit="fill" @click="createCaptcha" />
        </template>
      </el-input>
    </el-form-item>
  </el-form>
  <div class="login-btn">
    <el-button :icon="CircleClose" round size="large" @click="resetForm(loginFormRef)">重置</el-button>
    <el-button :icon="UserFilled" round size="large" type="primary" :loading="loading" @click="login(loginFormRef)">
      登录
    </el-button>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { useRouter } from "vue-router";
import { HOME_URL } from "@/config";
import { getTimeState } from "@/utils";
import { SystemUser } from "@/api/interface/systemUser";
import { Captcha } from "@/api/interface/captcha";
import { ElNotification } from "element-plus";
import { postSystemUserLoginApi } from "@/api/modules/systemUser";
import { captchaGenerateApi } from "@/api/modules/captcha";
import { useUserStore } from "@/stores/modules/user";
import { useTabsStore } from "@/stores/modules/tabs";
import { useKeepAliveStore } from "@/stores/modules/keepAlive";
import { initDynamicRouter } from "@/routers/modules/dynamicRouter";
import { CircleClose, UserFilled } from "@element-plus/icons-vue";
import { ElForm } from "element-plus";
import md5 from "md5";

const router = useRouter();
const userStore = useUserStore();
const tabsStore = useTabsStore();
const keepAliveStore = useKeepAliveStore();

type FormInstance = InstanceType<typeof ElForm>;
const loginFormRef = ref<FormInstance>();
const loginRules = reactive({
  username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
  password: [{ required: true, message: "请输入密码", trigger: "blur" }],
  captchaCode: [{ required: true, message: "请输入验证码", trigger: "blur" }]
});

const loading = ref(false);
const loginForm = reactive<SystemUser.ReqSystemUserLogin>({
  username: "", // 用户名
  password: "", // 密码
  captchaCode: "", // 验证码
  captchaId: "" // 验证码id
});

//验证码
const resCaptcha = reactive<Captcha.ResCaptcha>({
  captchaId: "", // 验证码id
  captchaImage: "" // 验证码图片
});

// 获取验证码
const createCaptcha = async () => {
  if (loading.value) return;
  const { data } = await captchaGenerateApi();
  resCaptcha.captchaId = data.captchaId;
  resCaptcha.captchaImage = data.captchaImage;
};

// login
const login = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (!valid) return;
    loading.value = true;
    try {
      // 1.执行登录接口
      const { data } = await postSystemUserLoginApi({
        ...loginForm,
        password: md5(loginForm.password),
        captchaId: resCaptcha.captchaId
      });
      userStore.setToken(data.accessToken);
      // 2.添加动态路由
      await initDynamicRouter();
      // 3.清空 tabs、keepAlive 数据
      userStore.setUserInfo({ name: data.nickname });
      // 4.清空 tabs、keepAlive 数据
      tabsStore.setTabs([]);
      keepAliveStore.setKeepAliveName([]);

      // 5.跳转到首页
      router.push(HOME_URL);
      ElNotification({
        title: getTimeState(),
        message: "欢迎登录 Geeker-Admin",
        type: "success",
        duration: 3000
      });
    } finally {
      loading.value = false;
      //登录失败，重置验证码
      createCaptcha();
    }
  });
};

// resetForm
const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};

onMounted(() => {
  // 1.获取验证码
  createCaptcha();
  // 监听 enter 事件（调用登录）
  document.onkeydown = (e: KeyboardEvent) => {
    // e = (window.event as KeyboardEvent) || e;
    if (e.code === "Enter" || e.code === "enter" || e.code === "NumpadEnter") {
      if (loading.value) return;
      login(loginFormRef.value);
    }
  };
});
</script>

<style scoped lang="scss">
@import "../index";
:deep(.el-input-group__append) {
  padding: 0;
}
</style>
