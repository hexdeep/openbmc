<script setup lang="ts">
import {request} from '@/utils/axios';
import {useI18n} from 'vue-i18n';

const toTerminal = (port: string) => window.open(`http://${window.location.hostname}:${port}`, '_blank')

const { t } = useI18n({
  messages: {
    zh: {
      switchTty: '交换机终端',
      open: '开启',
      close: '关闭',
      enter: '进入',
    },
  },
})
</script>

<template>
  <div class="p-4 flex flex-col gap-4">
        <div class="flex flex-col gap-2">
          <div>{{t('switchTty')}}</div>
          <el-button-group>
            <el-button @click="request('POST', '/switch/opentty')" type="primary">
              {{t('open')}}
            </el-button>
            <el-button @click="request('POST', '/switch/closetty')" type="warning">
              {{t('close')}}
            </el-button>
            <el-button @click=" toTerminal('7600')" type="success">
              {{t('enter')}}
            </el-button>
          </el-button-group>
        </div>

  </div>
</template>
