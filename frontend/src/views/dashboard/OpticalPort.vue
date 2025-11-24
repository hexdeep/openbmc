<script setup lang="ts">
import {request} from '@/utils/axios';
import {Refresh} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {useI18n} from 'vue-i18n';

const { t } = useI18n({ messages: {
  zh: {
    opticalPortTitle: '光口状态',
    connected: '已连接',
    notConnected: '未连接',
  },
} })

interface OpticalPort {
  port: string;
  status: string;
}

const opticalPorts = ref<OpticalPort[]>([])
const loadOpticalPort = () => request<any>('GET', '/optical-port').then(v => opticalPorts.value = v)
loadOpticalPort()

</script>

<template>
  <div class="card flex flex-col gap-4">
    <div class="flex items-center gap-4">
      <div class="text-lg">{{t('opticalPortTitle')}}</div>
      <el-button :icon="Refresh" @click="loadOpticalPort">
        {{t('refresh')}}
      </el-button>
    </div>
    <div class="grid grid-cols-2 gap-4">
      <div v-for="port in opticalPorts" :key="port.port" class="card flex items-center gap-4">
        <div>{{port.port}}</div>
        <el-tag :type="port.status === 'up' ? 'success' : 'danger'">
          {{port.status === 'up' ? t('connected') : t('notConnected')}}
        </el-tag>
      </div>
    </div>
  </div>
</template>
