<script setup lang="ts">
import {request} from '@/utils/axios';
import type {List} from '@/utils/utils';
import {Search} from '@element-plus/icons-vue';
import {useUrlSearchParams} from '@vueuse/core';
import dayjs from 'dayjs';
import {reactive} from 'vue';
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
    clearLog: '清理日志',
    byTime: '按时间',
    byMethod: '按请求方法',
    byStatus: '按状态',
    OPTIONS: '预检',
    GET: '查询',
    POST: '提交',
    DELETE: '删除',
    '/login': '登录',
    200: '成功',
    500: '未知错误',
    400: '非法请求',
    401: '未认证',
    403: '权限不足',
  },
} })

const logs = ref<List<Log> | null>(null)

const load = async () => logs.value = await request<any>('GET', '/logs', params)
watchEffect(load)
load()

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

const isDialogOpen = ref(false)

const clearForm = reactive<{
  method: string[],
  status: number[],
}>({
  method: [],
  status: [],
})

</script>

<template>
  <div class="p-4 max-w-5xl mx-auto flex flex-col gap-4">
    <div v-if="logs" class="card flex flex-col gap-4">
      <div class="text-lg">
        {{t('logListTitle')}}
      </div>
      <div class="flex gap-4">
        <el-date-picker
          :model-value="[params.from && new Date(params.from), params.to && new Date(params.to)]"
          type="datetimerange"
          @update:model-value="v => {
            if (!v) {
              params.from = undefined
              params.to = undefined
            } else {
              params.from = v[0]?.toISOString()
              params.to = v[1]?.toISOString()
            }
          }"
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
        <el-button @click="isDialogOpen = true">
          {{t('clearLog')}}
        </el-button>
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
              {{t(method)}}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column :label="t('path')" :formatter="row => t(row.path)" />
        <el-table-column :label="t('status')">
          <template #default="{ row: { status } }">
            <el-tag :type="statusTypeMap[status]">
              {{t(status)}}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column />
      </el-table>
      <el-pagination
        layout="prev, pager, next, total"
        :total="logs.total"
        :current-page="Number(params.page || 1)"
        :page-size="Number(params.size || 9)"
        @update:current-page="v => params.page = String(v)"
        @update:page-size="v => params.size = String(v)"
      />
    </div>
  </div>
  <el-dialog v-model="isDialogOpen" :title="t('clearLog')">
    <el-form>
      <el-form-item :label="t('byTime')">
        <el-date-picker
          type="datetimerange"
        />
      </el-form-item>
      <el-form-item :label="t('byMethod')">
        <el-checkbox-group v-model="clearForm.method">
          <el-checkbox v-for="m in ['GET', 'POST', 'DELETE', 'OPTIONS']" :key="m" :value="m">
            {{t(m)}}
          </el-checkbox>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item :label="t('byStatus')">
        <el-checkbox-group v-model="clearForm.status">
          <el-checkbox v-for="s in [200, 400, 401, 403, 500]" :key="s" :value="s">
            {{t(s)}}
          </el-checkbox>
        </el-checkbox-group>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="request('POST', '/clear-logs', clearForm).then(ok => ok && load())">
        {{t('confirm')}}
      </el-button>
    </template>
  </el-dialog>
</template>
