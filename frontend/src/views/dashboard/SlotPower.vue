<script setup lang="ts">
import {request} from '@/utils/axios';
import {ArrowDown, ArrowUp} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {useI18n} from 'vue-i18n';

interface SubPower {
  id: string;
  active: boolean;
}

const subPowers = ref<SubPower[][][]>([])
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
    <div class="flex gap-4">
      <div v-for="pane in subPowers" class="flex flex-col gap-2 rounded-2xl p-4 bg-gray-100">
        <div v-for="drawer in pane" class="flex flex-col gap-2">
          <div class="flex gap-2 items-center">
            <el-button-group size="small">
              <el-button :icon="ArrowUp" @click="drawer.forEach(slot => request('POST', `/slot/${slot.id}/power-on`).then(loadSubPowers))" />
              <el-button :icon="ArrowDown" @click="drawer.forEach(slot => request('POST', `/slot/${slot.id}/power-off`).then(loadSubPowers))" />
            </el-button-group>
            <el-button
              v-for="slot in drawer"
              :key="slot.id"
              class="m-0!"
              :type="slot.active ? 'success' : 'warning'"
              @click="request('POST', `/slot/${slot.id}/power-${slot.active ? 'off' : 'on'}`).then(loadSubPowers)"
            >
              {{slot.id}} {{slot.active ? t('powered') : t('notPowered')}}
            </el-button>
          </div>
          <div class="h-px bg-gray-300" />
        </div>
      </div>
    </div>
  </div>
</template>
