<template>
  <div>
    <div class="card mb10">
      <h4 class="title">基础信息</h4>
      <Grid :gap="[20, 0]" :cols="4">
        <GridItem v-for="(item, index) in monitorTrendItem.head" :index="index" :key="index">
          <el-card shadow="never">
            <template #header>
              <div class="flex items-center justify-between">
                <span>{{ item.title }}</span>
                <el-icon>
                  <component :is="item.iconClass" />
                </el-icon>
              </div>
            </template>
            <div style="height: 130px" class="flex flex-col justify-between">
              <div class="flex flex-col justify-center">
                <span class="text-xxl">{{ item.data }}</span>
              </div>
              <div class="flex flex-col justify-center flex-1">
                <template v-if="index === 0">
                  <div class="margin-top-lg">
                    <div>{{ item.extra.data }}</div>
                    <div class="margin-top-sm">{{ item.extra.data1 }}</div>
                  </div>
                </template>
                <template v-else-if="index === 1">
                  <div class="margin-top" style="position: relative">
                    <div>已用内存：{{ item.extra.data }}</div>
                    <div class="margin-top-sm">剩余内存：{{ item.extra.data1 }}</div>
                    <div class="stack-avatar-wrapper"></div>
                  </div>
                </template>
                <template v-else-if="index === 2">
                  <el-progress :percentage="item.extra.data" />
                </template>
                <template v-else-if="index === 3">
                  <LoadChart ref="loadChartRef" :data-model="monitorTrendItem.load" />
                </template>
              </div>
              <div class="divide"></div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-grey">{{ item.bottomTitle }}</span>
                <span class="text-sm text-grey">{{ item.totalSum }}</span>
              </div>
            </div>
          </el-card>
        </GridItem>
      </Grid>
    </div>
    <div class="card mb10">
      <h4 class="title">实时网卡流量（全部网卡）</h4>
      <Grid :gap="[20, 0]">
        <GridItem>
          <el-card shadow="never" class="card-border">
            <div class="text-number">{{ last?.up }} KB</div>
            <div class="title-text">上行</div>
          </el-card>
        </GridItem>
        <GridItem>
          <el-card shadow="never" class="card-border">
            <div class="text-number">{{ last?.down ?? 0 }} KB</div>
            <div class="title-text">下行</div>
          </el-card>
        </GridItem>
        <GridItem>
          <el-card shadow="never" class="card-border">
            <div class="text-number">{{ last?.bytesSent ?? 0 }}</div>
            <div class="title-text">总发送</div>
          </el-card>
        </GridItem>
        <GridItem>
          <el-card shadow="never" class="card-border">
            <div class="text-number">{{ last?.bytesRecv ?? 0 }}</div>
            <div class="title-text">总接收</div>
          </el-card>
        </GridItem>
      </Grid>
      <FullYearSalesChart ref="fullYearSalesChartRef" :data-model="monitorTrendItem.net" />
    </div>
    <div class="card mb10">
      <h4 class="title">服务器信息</h4>
      <el-descriptions :column="3" border>
        <el-descriptions-item label="服务器名称">
          <el-tag type="info">{{ monitorRunItem.hostname }} </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="操作系统">
          <el-tag type="info">{{ monitorRunItem.os }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="平台">
          <el-tag type="info">{{ monitorRunItem.platform }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="架构">
          <el-tag type="info">{{ monitorRunItem.arch }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="公网 IP">
          <el-tag type="info">{{ monitorRunItem.publicIP }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="内网 IP">
          <el-tag type="info">{{ monitorRunItem.intranetIP }}</el-tag>
        </el-descriptions-item>
      </el-descriptions>
    </div>
    <div class="card mb10">
      <h4 class="title">服务信息</h4>
      <el-descriptions :column="4" border>
        <el-descriptions-item label="Go 版本">
          <el-tag type="info">{{ monitorRunItem.goVersion }} </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="Go 架构">
          <el-tag type="info">{{ monitorRunItem.goArch }} </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="Go 运行内存">
          <el-tag type="info">{{ monitorRunItem.goMem }} </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="Go 协程数">
          <el-tag type="info">{{ monitorRunItem.goroutine }} </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="Ratel 版本">
          <el-tag type="info">{{ monitorRunItem.ratelVersion }} </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="编译版本">
          <el-tag type="info">{{ monitorRunItem.buildVersion }} </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="编译时间">
          <el-tag type="info">
            {{ formatToDateTime(monitorRunItem.buildTime) }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="启动时间">
          <el-tag type="info">
            {{ formatToDateTime(monitorRunItem.launchTime) }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="运行时长">
          <el-tag type="info">
            {{ formatBefore(new Date(monitorRunItem.launchTime)) }}
          </el-tag>
        </el-descriptions-item>
      </el-descriptions>
    </div>
    <div class="card mb10">
      <h4 class="title">生产环境依赖</h4>
      <el-descriptions :column="3" border>
        <el-descriptions-item v-for="(value, key) in dependencies" :key="key" width="400px" :label="key">
          <el-tag type="info">
            {{ value }}
          </el-tag>
        </el-descriptions-item>
      </el-descriptions>
    </div>
    <div class="card">
      <h4 class="title">开发环境依赖</h4>
      <el-descriptions :column="3" border>
        <el-descriptions-item v-for="(value, key) in devDependencies" :key="key" width="400px" :label="key">
          <el-tag type="info">
            {{ value }}
          </el-tag>
        </el-descriptions-item>
      </el-descriptions>
    </div>
  </div>
</template>

<script setup lang="ts" name="home">
import { ref } from "vue";
import { useIntervalFn } from "@vueuse/core";
import { fetchEventData } from "@/api/sse";
import { useRoute } from "vue-router";
import Grid from "@/components/Grid/index.vue";
import GridItem from "@/components/Grid/components/GridItem.vue";
import { Monitor } from "@/api/interface/monitor";
import { getMonitorRunApi, getMonitorTrendApi } from "@/api/modules/monitor";
import { formatBefore, formatToDateTime } from "@/utils/date";
import LoadChart from "./components/chart/LoadChart.vue";
import FullYearSalesChart from "./components/chart/FullYearSalesChart.vue";
const { pkg, lastBuildTime } = __APP_INFO__;
const { dependencies, devDependencies, version } = pkg;
const route = useRoute();
const loadChartRef = ref<InstanceType<typeof LoadChart>>();
const fullYearSalesChartRef = ref<InstanceType<typeof FullYearSalesChart>>();
const last = ref<any>({
  bytesSent: "0B",
  bytesRecv: "0B",
  down: "0",
  up: "0"
});
const monitorRunItem = ref<Monitor.MonitorRunItem>({
  arch: "",
  buildTime: "",
  buildVersion: "",
  goArch: "",
  goMem: "",
  goSize: "",
  goVersion: "",
  goroutine: 0,
  hostname: "",
  intranetIP: "",
  launchTime: "",
  os: "",
  platform: "",
  publicIP: "",
  pwd: "",
  ratelVersion: "",
  rootPath: "",
  runTime: 0
});
const monitorTrendItem = ref<Monitor.MonitorTrendItem>({
  head: [
    {
      title: "CPU",
      data: "0%",
      bottomTitle: "CPU数量",
      totalSum: "",
      iconClass: "Cpu",
      extra: {
        data: "",
        data1: ""
      }
    },
    {
      title: "内存",
      data: "0%",
      bottomTitle: "总内存",
      totalSum: "0GB",
      iconClass: "PieChart",
      extra: {
        data: "0GB",
        data1: "0GB"
      }
    },
    {
      title: "磁盘",
      data: "已用 0GB",
      bottomTitle: "总容量",
      totalSum: "0GB",
      iconClass: "PieChart",
      extra: {
        data: "",
        data1: ""
      }
    },
    {
      title: "负载",
      data: "0%",
      bottomTitle: "总进程数",
      totalSum: "0个",
      iconClass: "PieChart",
      extra: {
        data: "",
        data1: ""
      }
    }
  ],
  load: [],
  net: []
});
useIntervalFn(() => {
  if (route.name === "home") {
    runFun();
    trendFun();
  }
}, 3000);

const runFun = async () => {
  const { data } = await getMonitorRunApi();
  monitorRunItem.value = data;
};

const trendFun = async () => {
  const { data } = await getMonitorTrendApi();
  monitorTrendItem.value = data;
  last.value = data.net[data.net.length - 1];
  monitorTrendItem.value.load = newLoad(data.load);
  monitorTrendItem.value.net = newNet(data.net);
};
const newNet = (rows: Monitor.Net[]) => {
  // 将 rows 里的time 装换类型, 其他的不变
  return rows.map(row => {
    return {
      ...row,
      time: formatToDateTime(row.time)
    };
  });
};

const newLoad = (rows: Monitor.Load[]) => {
  // 将 rows 里的time 装换类型, 其他的不变
  return rows.map(row => {
    return {
      ...row,
      time: formatToDateTime(row.time)
    };
  });
};
</script>

<style lang="scss" scoped>
@import "@/styles/custom";
.card {
  .title {
    margin: 0 0 15px;
    font-size: 17px;
    font-weight: bold;
    color: var(--el-text-color-primary);
  }
  .text {
    font-size: 15px;
    line-height: 25px;
    color: var(--el-text-color-regular);
    .el-link {
      font-size: 15px;
    }
  }
}
.justify-between {
  justify-content: space-between;
}
.flex-col {
  flex-direction: column;
}
.flex {
  display: flex;
}
.justify-center {
  justify-content: center;
}
.flex-1 {
  flex: 1 1 0%;
}
.grid-content {
  height: 100px;
}
:deep(.el-card) {
  font-size: 14px;
  border-radius: unset;
}
:deep(.el-card__body) {
  padding: 8px 11px;
}
:deep(.el-card__footer) {
  padding: 8px 11px;

  // font-size: 14px;
}
:deep(.el-card__header) {
  padding: 8px 11px;

  // font-size: 14px;
}
.card-border {
  border: unset;
}
</style>
