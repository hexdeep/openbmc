<script setup lang="ts">
import {request} from '@/utils/axios';
import {formatSize} from '@/utils/utils';
import {useEventSource} from '@vueuse/core';
import {computed} from 'vue';
import {useI18n} from 'vue-i18n';

const { t } = useI18n({ messages: {
  zh: {
    slot: '插槽',
    deviceListTitle: '通电接口列表',
    temperature: '温度',
    mem: '内存',
    powerStatus: '电源状态',
    operation: '操作',
    powerOn: '上电',
    powerOff: '下电',
    detail: '详情',
    updateSystem: '升级',
    enterTerminal: '终端',
  },
} })

const { data } = useEventSource('/api/powered-slot')
const poweredSlots = computed(() => {
  if (!data.value) return 0
  try {
    return JSON.parse(data.value)
  } catch {
    return 0
  }
})
/*
const poweredSlots = ref<PoweredSlot[]>([])
const loadPoweredSlots = () => request<PoweredSlot[]>('GET', '/powered-slot').then(v => poweredSlots.value = v.sort((a, b) => a.slot - b.slot))
loadPoweredSlots()
*/
</script>

<template>
  <div class="flex flex-col gap-4 p-4">

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
        <el-table-column :label="t('mem')" width="200">
          <template #default="{ row }">
            {{formatSize(row.memUsed)}} /
            {{formatSize(row.memTotal)}} /
            {{Math.trunc((row.memUsed / row.memTotal) * 100)}} %
          </template>
        </el-table-column>
        <el-table-column
          :label="t('temperature')"
          prop="temp"
          :formatter="({ temp }) => `${temp} °C`"
        />
        <el-table-column :label="t('operation')" width="200">
          <template #default="{ row }">
            <el-button type="danger" size="small" @click="request('POST', `/slot/${row.slot}/power-off`)">
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
