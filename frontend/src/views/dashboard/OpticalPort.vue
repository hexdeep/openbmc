<script setup lang="ts">
import {request} from '@/utils/axios';
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
  id: string;
  status: boolean;
}

const opticalPorts = ref<OpticalPort[][]>([])
const loadOpticalPort = () => request<any>('GET', '/optical-port').then(v => opticalPorts.value = v)
loadOpticalPort()

</script>

<template>
  <div class="flex flex-col gap-2">
    <div v-for="group in opticalPorts" class="flex gap-2">
      <div v-for="port in group" :key="port.id" class="card flex items-center gap-4">
        <div>{{port.id}}</div>
        <el-tag :type="port.status ? 'success' : 'danger'">
          {{port.status ? t('connected') : t('notConnected')}}
        </el-tag>
      </div>
    </div>
  </div>
</template>
