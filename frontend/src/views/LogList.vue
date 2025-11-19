<script setup lang="ts">
import {request} from '@/utils/axios';
import type {List} from '@/utils/utils';
import {useUrlSearchParams} from '@vueuse/core';
import {ref, watchEffect} from 'vue';
import {useI18n} from 'vue-i18n';

interface Log {
  id: number;
  createdAt: string;
  clientIp: string;
  path: string;
  method: string;
  status: number;
}

const params = useUrlSearchParams<any>('history')

const { t } = useI18n({ messages: {
  zh: {
    logListTitle: '日志列表',
  },
} })

const logs = ref<List<Log> | null>(null)

watchEffect(async () => logs.value = await request<any>('GET', '/logs', params))

const statusTypeMap: Record<any, any> = {
  200: 'success',
  500: 'danger',
  401: 'warning',
}

</script>

<template>
  <div class="p-4 max-w-5xl mx-auto flex flex-col gap-4">
    <div v-if="logs" class="card flex flex-col gap-4">
      <div class="text-lg">
        {{t('logListTitle')}}
      </div>
      <el-table :data="logs.data">
        <el-table-column :label="t('createdAt')" prop="createdAt" />
        <el-table-column :label="t('clientIp')" prop="clientIp" />
        <el-table-column :label="t('method')" prop="method" />
        <el-table-column :label="t('path')" prop="path" />
        <el-table-column :label="t('status')">
          <template #default="{ row: { status } }">
            <el-tag :type="statusTypeMap[status]">
              {{status}}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column />
      </el-table>
      <el-pagination layout="prev, pager, next, total" :total="logs.total" />
    </div>
  </div>
</template>
