<script setup lang="ts">
import {useI18n} from 'vue-i18n';


const { t } = useI18n({ messages: {
  zh: {
    deviceInfoTitle: '设备信息',
    systemVersion: '当前管理系统版本',
    deviceName: '设备名称',
    opticalPort: '光口 {number}',
  },
} })
</script>

<template>
  <div class="flex flex-col gap-4 p-4">

    <div class="text-lg font-bold">
      {{t('deviceInfoTitle')}}
    </div>

    <el-descriptions class="card w-xl" border :column="1">
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
      <el-descriptions-item :label="t('opticalPort', { number: 1 })">
      </el-descriptions-item>
      <el-descriptions-item :label="t('opticalPort', { number: 2 })">
      </el-descriptions-item>
      <el-descriptions-item :label="t('opticalPort', { number: 3 })">
      </el-descriptions-item>
      <el-descriptions-item :label="t('opticalPort', { number: 4 })">
      </el-descriptions-item>
    </el-descriptions>
  </div>
</template>
