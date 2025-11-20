<script setup lang="ts">
import {request} from '@/utils/axios';
import {formatSize} from '@/utils/utils';
import {ref} from 'vue';
import {useI18n} from 'vue-i18n';

const { t } = useI18n({ messages: {
  zh: {
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

interface SOMStatus {
  id: number;
  status: boolean;
}

const somStatuses = ref<SOMStatus[][][]>([])
request('GET', '/som-statuses').then(v => somStatuses.value = v)

interface Interface {
  id: number;
  active: boolean;
  drawer: number;
  diskUsed: number;
  diskTotal: number;
  cpuUsage: number;
  ip: string;
  temperature: number;
}

const interfaces = ref<Interface[]>([])
request('GET', '/powered-interfaces').then(v => interfaces.value = v)
</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">

    <div class="card flex flex-col gap-4">
      <div class="text-lg">
        {{t('statusSummaryTitle')}}
      </div>
      <div class="flex gap-8">
        <div v-for="panel in somStatuses" class="flex flex-col gap-4">
          <div v-for="drawer in panel" class="flex gap-2">
            <el-button class="!m-0" v-for="slot in drawer" :type="slot.status ? 'success' : 'warning'">
              {{slot.id}} {{slot.status ? t('powered') : t('notPowered')}}
            </el-button>
          </div>
        </div>
      </div>
    </div>

    <div class="card flex flex-col gap-4">
      <div class="text-lg">
        {{t('deviceListTitle')}}
      </div>
      <el-table :data="interfaces" class="rounded-2xl">
            {{}}
        <el-table-column label="ID" prop="id" />
        <el-table-column label="status">
          <template #default="{ row }">
            <el-tag :type="row.active ? 'success' : 'danger'">
              {{row.active ? '运行' : '异常'}}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="IP" prop="ip" />
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
          prop="temperature"
          :formatter="({ temperature }) => `${temperature} °C`"
        />
        <el-table-column :label="t('operation')" width="200">
          <el-button type="danger" size="small">
            {{t('powerOff')}}
          </el-button>
          <el-button type="success" size="small">
            {{t('updateSystem')}}
          </el-button>
          <el-button type="primary" size="small">
            {{t('enterTerminal')}}
          </el-button>
        </el-table-column>
      </el-table>
    </div>

  </div>
</template>
