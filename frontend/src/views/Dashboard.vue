<script setup lang="ts">
import {request} from '@/utils/axios';
import {ref} from 'vue';
import {useI18n} from 'vue-i18n';
import {formatSize} from '@/utils/utils';

const { t } = useI18n({ messages: {
  zh: {
    cpuUsage: 'CPU 使用率',
    used: '使用',
    total: '总量',
    memory: '内存',
    disk: '硬盘',
    statusSummaryTitle: '状态总览',
    deviceInfoTitle: '设备信息',
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
    runningStatus: '运行状态',
    running: '运行中',
    notRunning: '未运行',
  },
} })

interface Power {
  id: number;
  name: string;
  powered: boolean;
  running: boolean;
}

const powers = ref<Power[]>([])
request<any>('GET', '/powers').then(v => powers.value = v)

interface OpticalPort {
  name: string;
  connected: boolean;
}

const opticalPorts = ref<OpticalPort[]>([])
request<any>('GET', '/optical-ports').then(v => opticalPorts.value = v)

const cpuPercentage = ref(20)
const memoryStatus = ref({
  used: 1000000000,
  total: 3000000000,
})
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
          <el-progress type="dashboard" :percentage="cpuPercentage">
            <div>CPU</div>
            <div class="text-xs mt-2">{{Math.trunc(cpuPercentage)}} %</div>
          </el-progress>
          <el-progress type="dashboard" :percentage="(memoryStatus.used / memoryStatus.total) * 100">
            <div>{{t('memory')}}</div>
            <div class="text-xs mt-2">{{t('used')}} {{formatSize(memoryStatus.used)}}</div>
            <div class="text-xs mt-1">{{t('total')}} {{formatSize(memoryStatus.total)}}</div>
          </el-progress>
          <el-progress type="dashboard" :percentage="(diskStatus.used / diskStatus.total) * 100">
            <div>{{t('disk')}}</div>
            <div class="text-xs mt-2">{{t('used')}} {{formatSize(diskStatus.used)}}</div>
            <div class="text-xs mt-1">{{t('total')}} {{formatSize(diskStatus.total)}}</div>
          </el-progress>
        </div>
      </div>

      <div class="card flex flex-col gap-4">
        <div class="text-lg">{{t('powerTitle')}}</div>
        <el-table :data="powers">
          <el-table-column label="ID" prop="id" />
          <el-table-column :label="t('name')" prop="name" />
          <el-table-column :label="t('poweredStatus')">
            <template #default="{ row: { powered } }">
              <el-tag :type="powered ? 'success' : 'warning'">
                {{powered ? t('powered') : t('notPowered')}}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column :label="t('runningStatus')">
            <template #default="{ row: { running } }">
              <el-tag :type="running ? 'success' : 'warning'">
                {{running ? t('running') : t('notRunning')}}
              </el-tag>
            </template>
          </el-table-column>
        </el-table>
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
        <div class="text-lg">{{t('opticalPortTitle')}}</div>
        <el-table :data="opticalPorts">
          <el-table-column :label="t('name')" prop="name" />
          <el-table-column :label="t('connectStatus')">
            <template #default="{ row: { connected } }">
              <el-tag :type="connected ? 'success' : 'warning'">
                {{connected ? t('connected') : t('notConnected')}}
              </el-tag>
            </template>
          </el-table-column>
        </el-table>
      </div>

    </div>

  </div>
</template>
