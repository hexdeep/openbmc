<script setup lang="ts">
import {useI18n} from 'vue-i18n';


const { t } = useI18n({ messages: {
  zh: {
    statusSummaryTitle: '状态总览',
    deviceListTitle: '设备列表',
    temperature: '温度',
    diskCapacity: '硬盘容量',
    powerStatus: '电源状态',
    operation: '操作',
    powerOn: '上电',
    powerOff: '下电',
    detail: '详情',
  },
} })

interface Device {
  id: number;
  ip: string;
  mac: string;
  capacity: number;
  temperature: number;
}

const testStatus = [
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,1,1,0,0,0,0,0,
].map((v, i) => ({ id: i, status: v === 0 ? false : true }))

const testData: Device[] = [
  {
    id: 1,
    ip: '192.168.2.2',
    mac: 'aaaaaa',
    capacity: 12341234234,
    temperature: 123412342134,
  },
]
</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">

    <div class="card flex flex-col gap-4">
      <div class="text-lg">
        {{t('statusSummaryTitle')}}
      </div>
      <div class="flex gap-4">
        <el-tag v-for="s in testStatus" :key="s.id" :type="s.status ? 'success' : 'danger'">
          {{s.id}}
        </el-tag>
      </div>
    </div>

    <div class="card flex flex-col gap-4">
      <div class="text-lg">
        {{t('deviceListTitle')}}
      </div>
      <el-table :data="testData" class="rounded-2xl">
        <el-table-column label="ID" prop="id" />
        <el-table-column label="IP" prop="ip" />
        <el-table-column label="MAC" prop="mac" />
        <el-table-column :label="t('diskCapacity')" prop="capacity" />
        <el-table-column :label="t('temperature')" prop="temperature" />
        <el-table-column :label="t('operation')">
          <template #default="{row}">
            <el-button type="primary" size="small">
              {{t('powerOn')}}
            </el-button>
            <el-button type="danger" size="small">
              {{t('powerOff')}}
            </el-button>
            <el-button type="warning" size="small" @click="$router.push(`/devices/${row.id}`)">
              {{t('detail')}}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination layout="prev, pager, next, total" :total="10" />
    </div>

  </div>
</template>
