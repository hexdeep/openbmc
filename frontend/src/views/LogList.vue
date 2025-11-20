<script setup lang="ts">
import {request} from '@/utils/axios';
import type {List} from '@/utils/utils';
import {Search} from '@element-plus/icons-vue';
import {useUrlSearchParams} from '@vueuse/core';
import dayjs from 'dayjs';
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
    createdAt: '创建时间',
    clientIp: '客户端IP',
    path: '路径',
    method: '请求方法',
    status: '状态码',
  },
} })

const logs = ref<List<Log> | null>(null)

watchEffect(async () => logs.value = await request<any>('GET', '/logs', params))

const statusTypeMap: Record<any, any> = {
  200: 'success',
  500: 'danger',
  401: 'warning',
}

const methodTypeMap: Record<any, any> = {
  GET: "success",
  POST: "primary",
  PUT: "warning",
  DELETE: "danger",
  PATCH: "info",
  OPTIONS: "info",
  HEAD: "info",
}

</script>

<template>
  <div class="p-4 max-w-5xl mx-auto flex flex-col gap-4">
    <div v-if="logs" class="card flex flex-col gap-4">
      <div class="text-lg">
        {{t('logListTitle')}}
      </div>
      <div class="flex gap-4">
        <el-date-picker
          :model-value="params.time"
          @update:model-value="v => params.time = v?.map((e: Date) => e.toISOString())"
          type="datetimerange"
        />
        <el-input v-model="params.path" :placeholder="t('path')" class="!w-48" :suffix-icon="Search" />
        <el-select v-model="params.method" :placeholder="t('method')" class="!w-48" :empty-values="['', undefined]">
          <el-option :label="t('none')" value="" />
          <el-option v-for="v in ['GET', 'POST', 'PATCH', 'PUT', 'DELETE', 'OPTIONS']" :key="v" :label="v" :value="v" />
        </el-select>
        <el-select v-model="params.status" :placeholder="t('status')" class="!w-48" :empty-values="[0, undefined]">
          <el-option :label="t('none')" :value="0" />
          <el-option v-for="v in [200, 401, 500]" :key="v" :label="v" :value="v" />
        </el-select>
      </div>
      <el-table :data="logs.data">
        <el-table-column
          :label="t('createdAt')"
          prop="createdAt"
          :formatter="({createdAt}) => dayjs(createdAt).format('YYYY-MM-DD hh:mm:ss')"
        />
        <el-table-column :label="t('clientIp')" prop="clientIp" />
        <el-table-column :label="t('method')" prop="method">
          <template #default="{ row: { method } }">
            <el-tag :type="methodTypeMap[method]">
              {{method}}
            </el-tag>
          </template>
        </el-table-column>
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
