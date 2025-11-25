<script setup lang="ts">
import {request} from '@/utils/axios';
import {formatSize} from '@/utils/utils';
import {useEventSource} from '@vueuse/core';
import {computed} from 'vue';
import {useI18n} from 'vue-i18n';

const { t } = useI18n({ messages: {
  zh: {
    slot: '插槽',
    status: '状态',
    deviceListTitle: '通电接口列表',
    temperature: '温度',
    mem: '内存',
    powerStatus: '电源状态',
    operation: '操作',
    load: '负载',
    powerOn: '上电',
    powerOff: '下电',
    detail: '详情',
    updateSystem: '升级',
    enterTerminal: '终端',
    uptime: '开机时间',
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

function formatDuration(seconds: number): string {
    const hours = Math.floor(seconds / 3600);
    const minutes = Math.floor((seconds % 3600) / 60);

    const parts: string[] = [];

    if (hours > 0) {
        parts.push(`${hours}h`);
    }

    if (minutes > 0 || hours === 0) {
        // Show minutes always if there are no hours, e.g. "0min"
        parts.push(`${minutes}min`);
    }

    return parts.join(" ");
}

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
        <el-table-column :label="t('slot')" prop="slot" />
        <el-table-column :label="t('status')">
          <template #default="{ row }">
            <el-tag :type="row.active ? 'success' : 'danger'">
              {{row.active ? '运行' : '异常'}}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="IP" prop="ip" />
        <el-table-column label="MAC" prop="mac" />
        <el-table-column :label="t('uptime')" prop="uptime" :formatter="({ uptime }) => formatDuration(uptime)" />
        <el-table-column :label="t('load')" prop="load" :formatter="({ load }) => `${load} %`" />
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
