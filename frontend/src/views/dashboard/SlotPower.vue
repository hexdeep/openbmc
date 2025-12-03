<script setup lang="ts">
import {request} from '@/utils/axios';
import {ArrowDown, ArrowUp} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {useI18n} from 'vue-i18n';

interface Pane {
  name: string;
  drawers: {
    name: string;
    slots: {
      id: string;
      active: boolean;
    }[];
  }[];
}

const subPowers = ref<Pane[]>([])
const loadSubPowers = () => request('GET', '/sub-power').then(v => subPowers.value = v)
loadSubPowers()

interface OpticalPort {
  id: string;
  status: boolean;
}

const opticalPorts = ref<OpticalPort[][]>([])
const loadOpticalPort = () => request<any>('GET', '/optical-port').then(v => opticalPorts.value = v)
loadOpticalPort()

const { t } = useI18n({ messages: {
  zh: {
    statusSummaryTitle: '状态总览',
    powered: '在线',
    notPowered: '下电',
    opticalPortTitle: '光口状态',
    connected: '已连接',
    notConnected: '未连接',
    fanSpeed: '风扇速度',
  },
} })

const fanSpeed = ref(50)
</script>

<template>
  <div class="card flex flex-col gap-4">

    <div class="flex items-center gap-4">
      <div class="text-lg">
        {{t('statusSummaryTitle')}}
      </div>
      <el-button @click="loadSubPowers">
        {{t('refresh')}}
      </el-button>
    </div>

    <div class="flex items-center gap-4">
      <div>{{t('fanSpeed')}}</div>
      <el-slider class="w-64!" :model-value="fanSpeed" @update:model-value="v => request('POST', `/fan-speed/${v}`).then(res => fanSpeed = res)" />
    </div>

    <div class="flex items-end gap-4">
      <div v-for="pane in subPowers" class="flex flex-col gap-2 rounded-2xl p-4 bg-gray-100">
        <div>{{pane.name}}</div>
        <div v-for="drawer in pane.drawers" class="flex flex-col gap-2">
          <div class="flex gap-2 items-center">
            <div>{{drawer.name}}</div>
            <el-button-group size="small">
              <el-button :icon="ArrowUp" @click="drawer.slots.forEach(slot => request('POST', `/slot/${slot.id}/power-on`).then(loadSubPowers))" />
              <el-button :icon="ArrowDown" @click="drawer.slots.forEach(slot => request('POST', `/slot/${slot.id}/power-off`).then(loadSubPowers))" />
            </el-button-group>
            <el-button
              v-for="slot in drawer.slots"
              :key="slot.id"
              class="m-0!"
              :type="slot.active ? 'success' : 'warning'"
              @click="request('POST', `/slot/${slot.id}/power-${slot.active ? 'off' : 'on'}`).then(loadSubPowers)"
            >
              {{slot.id}} {{slot.active ? t('powered') : t('notPowered')}}
            </el-button>
          </div>
          <div class="h-px bg-gray-300" />
        </div>
      </div>
    </div>

    <div class="flex flex-col gap-2">
      <div v-for="group in opticalPorts" class="flex gap-2">
        <div v-for="port in group" :key="port.id" class="card flex items-center gap-4">
          <div>{{port.id}}</div>
          <el-tag :type="port.status ? 'success' : 'danger'">
            {{port.status ? t('connected') : t('notConnected')}}
          </el-tag>
        </div>
      </div>
    </div>

  </div>
</template>
