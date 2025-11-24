<script setup lang="ts">
import {useI18n} from 'vue-i18n';


const { t } = useI18n({ messages: {
  zh: {
    deviceInfoTitle: '设备信息',
    systemVersion: '当前管理系统版本',
    deviceName: '设备名称',
  },
} })
</script>

<template>
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
</template>
