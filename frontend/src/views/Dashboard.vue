<script setup lang="ts">
import {request} from '@/utils/axios';
import {computed, ref} from 'vue';
import {useI18n} from 'vue-i18n';
import {formatSize} from '@/utils/utils';
import {Refresh} from '@element-plus/icons-vue';
import {useEventSource} from '@vueuse/core';

const { t } = useI18n({ messages: {
  zh: {
    cpuUsage: 'CPU 使用率',
    used: '使用',
    total: '总量',
    memory: '内存',
    disk: '硬盘',
    statusSummaryTitle: '状态总览',
    deviceInfoTitle: '设备信息',
    fanSpeedTitle: '风扇速度',
    systemVersion: '当前管理系统版本',
    deviceName: '设备名称',
    powerTitle: '电源状态',
    opticalPortTitle: '光口状态',
    connectStatus: '连接状态',
    connected: '已连接',
    notConnected: '未连接',
    poweredStatus: '上电状态',
    powered: '已上电',
    notPowered: '未上电',
    powerOn: '上电',
    powerOff: '下电',
    runningStatus: '运行状态',
    running: '运行中',
    notRunning: '未运行',
  },
} })

interface Power {
  id: number;
  name: string;
  active: boolean;
}

const { data } = useEventSource('/api/usage')

const usage = computed(() => {
  if (!data.value) return 0
  try {
    return JSON.parse(data.value)
  } catch {
    return 0
  }
})


const powers = ref<Power[]>([])
const loadMainPowers = () => request<any>('GET', '/main-power').then(v => powers.value = v)
loadMainPowers()

interface OpticalPort {
  port: string;
  status: string;
}

const opticalPorts = ref<OpticalPort[]>([])
const loadOpticalPort = () => request<any>('GET', '/optical-port').then(v => opticalPorts.value = v)
loadOpticalPort()

interface FanSpeed {
  id: number;
  speed: number;
}

const fanSpeeds = ref<FanSpeed[]>([])
request('GET', '/fan-speeds').then(v => fanSpeeds.value = v)

const diskStatus = ref({
  used: 10000000000,
  total: 30000000000,
})

</script>

<template>
  <div class="flex flex-col lg:flex-row gap-4 p-4">

    <div class="basis-2/3 flex flex-col gap-4">

      <div class="card flex flex-col gap-4">
        <div class="text-lg">
          {{t('statusSummaryTitle')}}
        </div>
        <div class="flex flex-wrap gap-4">
          <el-progress type="dashboard" :percentage="usage.cpu">
            <div>CPU</div>
            <div class="text-xs mt-2">{{usage.cpu}} %</div>
          </el-progress>
          <el-progress type="dashboard" :percentage="(usage.memUsed / usage.memTotal) * 100">
            <div>{{t('memory')}}</div>
            <div class="text-xs mt-2">{{t('used')}} {{formatSize(usage.memUsed)}}</div>
            <div class="text-xs mt-1">{{t('total')}} {{formatSize(usage.memTotal)}}</div>
          </el-progress>
          <el-progress type="dashboard" :percentage="(diskStatus.used / diskStatus.total) * 100">
            <div>{{t('disk')}}</div>
            <div class="text-xs mt-2">{{t('used')}} {{formatSize(diskStatus.used)}}</div>
            <div class="text-xs mt-1">{{t('total')}} {{formatSize(diskStatus.total)}}</div>
          </el-progress>
        </div>
      </div>

      <div class="card flex flex-col gap-4">
        <div class="text-lg">
          {{t('fanSpeedTitle')}}
        </div>
        <div class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
          <div v-for="s in fanSpeeds" :key="s.id">
            <div>{{s.id}} - {{s.speed}} %</div>
            <el-slider v-model="s.speed">
            </el-slider>
          </div>
        </div>
      </div>

    </div>

    <div class="basis-1/3 flex flex-col gap-4">

      <div class="card flex flex-col gap-4">
        <div class="text-lg">
          {{t('deviceInfoTitle')}}
        </div>

        <el-descriptions border :column="1">
          <el-descriptions-item :label="t('systemVersion')">
          </el-descriptions-item>
          <el-descriptions-item :label="t('deviceName')">
          </el-descriptions-item>
          <el-descriptions-item label="IP">
          </el-descriptions-item>
          <el-descriptions-item label="MAC1">
          </el-descriptions-item>
          <el-descriptions-item label="MAC2">
          </el-descriptions-item>
        </el-descriptions>
      </div>

      <div class="card flex flex-col gap-4">
        <div class="flex items-center gap-4">
          <div class="text-lg">{{t('powerTitle')}}</div>
          <el-button @click="loadMainPowers" :icon="Refresh">
            {{t('refresh')}}
          </el-button>
        </div>
        <div class="flex flex-col gap-4">
          <div v-for="p in powers" class="card flex items-center gap-4">
            <div>{{p.name}}</div>
            <el-tag :type="p.active ? 'success' : 'warning'">
              {{p.active ? t('powered') : t('notPowered')}}
            </el-tag>
            <el-button-group>
              <el-button @click="request('POST', `/main-power/${p.id}/on`)">
                {{t('powerOn')}}
              </el-button>
              <el-button @click="request('POST', `/main-power/${p.id}/off`)">
                {{t('powerOff')}}
              </el-button>
            </el-button-group>
          </div>
        </div>
      </div>

      <div class="card flex flex-col gap-4">
        <div class="flex items-center gap-4">
          <div class="text-lg">{{t('opticalPortTitle')}}</div>
          <el-button :icon="Refresh" @click="loadOpticalPort">
            {{t('refresh')}}
          </el-button>
        </div>
        <div class="grid grid-cols-2 gap-4">
          <div v-for="port in opticalPorts" :key="port.port" class="card flex items-center gap-4">
            <div>{{port.port}}</div>
            <el-tag :type="port.status === 'up' ? 'success' : 'danger'">
              {{port.status === 'up' ? t('connected') : t('notConnected')}}
            </el-tag>
          </div>
        </div>
      </div>

    </div>

  </div>
</template>
