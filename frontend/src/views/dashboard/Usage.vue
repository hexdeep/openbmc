<script setup lang="ts">
import {formatSize} from '@/utils/utils'
import {useEventSource} from '@vueuse/core'
import {computed, ref} from 'vue'
import {useI18n} from 'vue-i18n'

const { t } = useI18n({ messages: {
  zh: {
    statusSummaryTitle: 'BMC状态',
    memory: '内存',
    disk: '硬盘',
    used: '使用',
    total: '总量',
  },
} })

const { data } = useEventSource('/api/usage')

const usage = computed(() => {
  if (!data.value) return 0
  try {
    return JSON.parse(data.value)
  } catch {
    return 0
  }
})

const diskStatus = ref({
  used: 10000000000,
  total: 30000000000,
})

</script>

<template>
  <div class="card flex flex-col gap-4">
    <div class="text-lg">
      {{t('statusSummaryTitle')}}
    </div>
    <div class="flex flex-wrap gap-4">
      <el-progress type="dashboard" :percentage="usage.cpu">
        <div>CPU</div>
        <div class="text-xs mt-2">{{usage.cpu}} %</div>
      </el-progress>
      <el-progress type="dashboard" :percentage="(usage.memUsed / usage.memTotal) * 100">
        <div>{{t('memory')}}</div>
        <div class="text-xs mt-2">{{t('used')}} {{formatSize(usage.memUsed)}}</div>
        <div class="text-xs mt-1">{{t('total')}} {{formatSize(usage.memTotal)}}</div>
      </el-progress>
      <el-progress type="dashboard" :percentage="(diskStatus.used / diskStatus.total) * 100">
        <div>{{t('disk')}}</div>
        <div class="text-xs mt-2">{{t('used')}} {{formatSize(diskStatus.used)}}</div>
        <div class="text-xs mt-1">{{t('total')}} {{formatSize(diskStatus.total)}}</div>
      </el-progress>
    </div>
  </div>
</template>
