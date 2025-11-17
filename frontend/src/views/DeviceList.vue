<script setup lang="ts">
import {useI18n} from 'vue-i18n';


const { t } = useI18n({ messages: {
  zh: {
    temperature: '温度',
    diskCapacity: '硬盘容量',
    powerStatus: '电源状态',
    operation: '操作',
    deviceListTitle: '设备列表',
  },
} })
</script>

<template>
  <div class="flex flex-col gap-4 p-4">
    <div class="card flex flex-col gap-4">
      <div class="text-lg">
        {{t('deviceListTitle')}}
      </div>
      <el-table class="rounded-2xl">
        <el-table-column label="IP" />
        <el-table-column label="MAC" />
        <el-table-column :label="t('diskCapacity')" />
        <el-table-column :label="t('temperature')" />
        <el-table-column :label="t('operation')" />
      </el-table>
    </div>
  </div>
</template>
