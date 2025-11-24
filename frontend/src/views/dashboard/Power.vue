<script setup lang="ts">
import {request} from '@/utils/axios';
import {Refresh} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {useI18n} from 'vue-i18n';

const { t } = useI18n({ messages: {
  zh: {
    powerTitle: '电源状态',
    powered: '已上电',
    notPowered: '未上电',
    powerOn: '上电',
    powerOff: '下电',
  },
} })

interface Power {
  id: number;
  name: string;
  active: boolean;
}

const powers = ref<Power[]>([])
const loadMainPowers = () => request<any>('GET', '/main-power').then(v => powers.value = v)
loadMainPowers()
</script>

<template>
  <div class="card flex flex-col gap-4">
    <div class="flex items-center gap-4">
      <div class="text-lg">{{t('powerTitle')}}</div>
      <el-button @click="loadMainPowers" :icon="Refresh">
        {{t('refresh')}}
      </el-button>
    </div>
    <div class="flex flex-col gap-4">
      <div v-for="p in powers" class="card flex items-center gap-4">
        <div>{{p.name}}</div>
        <el-tag :type="p.active ? 'success' : 'warning'">
          {{p.active ? t('powered') : t('notPowered')}}
        </el-tag>
        <el-button-group>
          <el-button @click="request('POST', `/main-power/${p.id}/on`)">
            {{t('powerOn')}}
          </el-button>
          <el-button @click="request('POST', `/main-power/${p.id}/off`)">
            {{t('powerOff')}}
          </el-button>
        </el-button-group>
      </div>
    </div>
  </div>
</template>
