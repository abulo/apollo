import { autoComplete, Plugin as importToCDN } from "vite-plugin-cdn-import";

/**
 * @description 打包时采用`cdn`模式，仅限外网使用（默认不采用，如果需要采用cdn模式，请在 .env.production 文件，将 VITE_CDN 设置成true）
 * 平台采用国内cdn：https://www.bootcdn.cn，当然你也可以选择 https://unpkg.com 或者 https://www.jsdelivr.com
 * 提醒：mockjs不能用cdn模式引入，会报错。正确的方式是，生产环境删除mockjs，使用真实的后端请求
 * 注意：上面提到的仅限外网使用也不是完全肯定的，如果你们公司内网部署的有相关js、css文件，也可以将下面配置对应改一下，整一套内网版cdn
 */

// cdn 地址
const cdnUrl = (viteCdnWall: boolean) => {
  return viteCdnWall ? "https://cdn.bootcdn.net/ajax/libs/{name}/{version}/{path}" : "https://unpkg.com/{name}@{version}/{path}";
};

const cdnVue = (viteCdnWall: boolean) => {
  return viteCdnWall
    ? {
        name: "vue",
        var: "Vue",
        path: "vue.global.prod.min.js"
      }
    : autoComplete("vue");
};

const cdnVueRouter = (viteCdnWall: boolean) => {
  return viteCdnWall
    ? {
        name: "vue-router",
        var: "VueRouter",
        path: "vue-router.global.min.js"
      }
    : {
        name: "vue-router",
        var: "VueRouter",
        path: "dist/vue-router.global.js"
      };
};

const cdnVueI18n = (viteCdnWall: boolean) => {
  return viteCdnWall
    ? {
        name: "vue-i18n",
        var: "VueI18n",
        path: "vue-i18n.global.prod.min.js"
      }
    : {
        name: "vue-i18n",
        var: "VueI18n",
        path: "dist/vue-i18n.global.js"
      };
};
const cdnVueDemi = (viteCdnWall: boolean) => {
  return viteCdnWall
    ? {
        name: "vue-demi",
        var: "VueDemi",
        path: "index.iife.min.js"
      }
    : {
        name: "vue-demi",
        var: "VueDemi",
        path: "lib/index.iife.js"
      };
};

const cdnPinia = (viteCdnWall: boolean) => {
  return viteCdnWall
    ? {
        name: "pinia",
        var: "Pinia",
        path: "pinia.iife.min.js"
      }
    : {
        name: "pinia",
        var: "Pinia",
        path: "dist/pinia.iife.js"
      };
};
const cdnElementPlus = (viteCdnWall: boolean) => {
  return viteCdnWall
    ? {
        name: "element-plus",
        var: "ElementPlus",
        path: "index.full.min.js",
        css: "index.min.css"
      }
    : {
        name: "element-plus",
        var: "ElementPlus",
        path: "dist/index.full.js",
        css: "dist/index.css"
      };
};

const cdnElementPlusIconsVue = () => {
  return {
    name: "@element-plus/icons-vue",
    var: "ElementPlusIconsVue",
    path: "dist/index.iife.min.js"
  };
};

const cdnWangEditor = () => {
  return {
    name: "@wangeditor/editor",
    var: "WangEditor",
    path: "dist/index.js",
    css: "dist/css/style.css"
  };
};

const cdnAxios = (viteCdnWall: boolean) => {
  return viteCdnWall
    ? {
        name: "axios",
        var: "axios",
        path: "axios.min.js"
      }
    : autoComplete("axios");
};
const cdnDayjs = (viteCdnWall: boolean) => {
  return viteCdnWall
    ? {
        name: "dayjs",
        var: "dayjs",
        path: "dayjs.min.js"
      }
    : {
        name: "dayjs",
        var: "dayjs",
        path: "dayjs.min.js"
      };
};
const cdnEcharts = (viteCdnWall: boolean) => {
  return viteCdnWall
    ? {
        name: "echarts",
        var: "echarts",
        path: "echarts.min.js"
      }
    : {
        name: "echarts",
        var: "echarts",
        path: "dist/echarts.js"
      };
};

const cdnModulesList = (viteCdnWall: boolean) => {
  const list = [
    cdnVue(viteCdnWall),
    cdnVueRouter(viteCdnWall),
    cdnVueI18n(viteCdnWall),
    cdnVueDemi(viteCdnWall),
    cdnPinia(viteCdnWall),
    cdnElementPlus(viteCdnWall),
    // cdnElementPlusIconsVue(viteCdnWall),
    // cdnWangEditor(viteCdnWall),
    cdnAxios(viteCdnWall),
    cdnDayjs(viteCdnWall),
    cdnEcharts(viteCdnWall)
  ];
  if (!viteCdnWall) {
    list.push(cdnElementPlusIconsVue());
  }
  if (!viteCdnWall) {
    list.push(cdnWangEditor());
  }

  return list;
};

export const cdn = (viteCdnWall: boolean) => {
  return importToCDN({
    //（prodUrl解释： name: 对应下面modules的name，version: 自动读取本地package.json中dependencies依赖中对应包的版本号，path: 对应下面modules的path，当然也可写完整路径，会替换prodUrl）
    prodUrl: cdnUrl(viteCdnWall),
    // prodUrl: "https://unpkg.com/{name}@{version}/{path}",
    // modules: [autoComplete("vue"), autoComplete("axios"), autoComplete("@vueuse/core")]
    modules: cdnModulesList(viteCdnWall)
  });
};
