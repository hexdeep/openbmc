<script setup lang="ts">
import {request} from '@/utils/axios';
import {formatSize} from '@/utils/utils';
import {ref} from 'vue';
import {useI18n} from 'vue-i18n';

const { t } = useI18n({ messages: {
  zh: {
    slot: '插槽',
    statusSummaryTitle: '状态总览',
    deviceListTitle: '通电接口列表',
    temperature: '温度',
    disk: '硬盘',
    powerStatus: '电源状态',
    operation: '操作',
    powerOn: '上电',
    powerOff: '下电',
    powered: '在线',
    notPowered: '下电',
    detail: '详情',
    updateSystem: '升级',
    enterTerminal: '终端',
  },
} })

interface SubPower {
  id: string;
  active: boolean;
  mac: string;
  temp: string;
}

const subPowers = ref<SubPower[]>([])
const loadSubPowers = () => request('GET', '/sub-power').then(v => subPowers.value = v)
loadSubPowers()

interface PoweredSlot {
  slot: number;
  active: boolean;
  drawer: number;
  diskUsed: number;
  diskTotal: number;
  cpuUsage: number;
  ip: string;
  temperature: number;
}

const poweredSlots = ref<PoweredSlot[]>([])
const loadPoweredSlots = () => request<PoweredSlot[]>('GET', '/powered-slot').then(v => poweredSlots.value = v.sort((a, b) => a.slot - b.slot))
loadPoweredSlots()

const load = () => {
  loadSubPowers()
  loadPoweredSlots()
}
</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">

    <div class="card flex flex-col gap-4">
      <div class="flex items-center gap-4">
        <div class="text-lg">
          {{t('statusSummaryTitle')}}
        </div>
        <el-button @click="loadSubPowers">
          {{t('refresh')}}
        </el-button>
      </div>
      <div class="grid grid-cols-6 grid-rows-8 gap-4 grid-flow-col">
        <el-button
          class="m-0!"
          v-for="p in subPowers"
          :type="p.active ? 'success' : 'warning'"
          @click="request('POST', `/slot/${p.id}/power-${p.active ? 'off' : 'on'}`).then(load)"
        >
          {{p.id}} {{p.active ? t('powered') : t('notPowered')}}
        </el-button>
      </div>
    </div>

    <div class="card flex flex-col gap-4">
      <div class="text-lg">
        {{t('deviceListTitle')}}
      </div>
      <el-table :data="poweredSlots" class="rounded-1xl">
        <el-table-column label="t('slot')" prop="slot" />
        <el-table-column label="status">
          <template #default="{ row }">
            <el-tag :type="row.active ? 'success' : 'danger'">
              {{row.active ? '运行' : '异常'}}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="IP" prop="ip" />
        <el-table-column label="MAC" prop="mac" />
        <el-table-column label="CPU" prop="cpuUsage" :formatter="({ cpuUsage }) => `${cpuUsage} %`" />
        <el-table-column :label="t('disk')" width="200">
          <template #default="{ row }">
            {{formatSize(row.diskUsed)}} /
            {{formatSize(row.diskTotal)}} /
            {{Math.trunc(row.diskUsed / row.diskTotal * 100)}} %
          </template>
        </el-table-column>
        <el-table-column
          :label="t('temperature')"
          prop="temp"
          :formatter="({ temp }) => `${temp} °C`"
        />
        <el-table-column :label="t('operation')" width="200">
          <template #default="{ row }">
            <el-button type="danger" size="small" @click="request('POST', `/slot/${row.slot}/power-off`).then(load)">
              {{t('powerOff')}}
            </el-button>
            <el-button type="success" size="small">
              {{t('updateSystem')}}
            </el-button>
            <el-button type="primary" size="small">
              {{t('enterTerminal')}}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

  </div>
</template>
