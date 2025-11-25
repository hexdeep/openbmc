<script setup lang="ts">
import {request} from '@/utils/axios';
import {ref} from 'vue';
import {useI18n} from 'vue-i18n';

interface SubPower {
  id: string;
  active: boolean;
  mac: string;
  temp: string;
  memUsed: number;
  memTotal: number;
}

const subPowers = ref<SubPower[]>([])
const loadSubPowers = () => request('GET', '/sub-power').then(v => subPowers.value = v)
loadSubPowers()


const { t } = useI18n({ messages: {
  zh: {
    statusSummaryTitle: '状态总览',
    powered: '在线',
    notPowered: '下电',
  },
} })
</script>

<template>
  <div class="card flex flex-col gap-4">
    <div class="flex items-center gap-4">
      <div class="text-lg">
        {{t('statusSummaryTitle')}}
      </div>
      <el-button @click="loadSubPowers">
        {{t('refresh')}}
      </el-button>
    </div>
    <div class="grid grid-cols-6 grid-rows-8 gap-4 grid-flow-col">
      <el-button
        class="m-0!"
        v-for="p in subPowers"
        :type="p.active ? 'success' : 'warning'"
        @click="request('POST', `/slot/${p.id}/power-${p.active ? 'off' : 'on'}`).then(loadSubPowers)"
      >
        {{p.id}} {{p.active ? t('powered') : t('notPowered')}}
      </el-button>
    </div>
  </div>
</template>
