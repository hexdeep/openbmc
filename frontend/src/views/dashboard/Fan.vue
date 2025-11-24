<script setup lang="ts">
import {request} from '@/utils/axios';
import {ref} from 'vue';
import {useI18n} from 'vue-i18n';

const { t } = useI18n({ messages: {
  zh: {
    fanSpeedTitle: '风扇速度',
  },
} })

interface FanSpeed {
  id: number;
  speed: number;
}

const fanSpeeds = ref<FanSpeed[]>([])
request('GET', '/fan-speeds').then(v => fanSpeeds.value = v)

</script>

<template>

      <div class="card flex flex-col gap-4">
        <div class="text-lg">
          {{t('fanSpeedTitle')}}
        </div>
        <div class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
          <div v-for="s in fanSpeeds" :key="s.id">
            <div>{{s.id}} - {{s.speed}} %</div>
            <el-slider v-model="s.speed">
            </el-slider>
          </div>
        </div>
      </div>

</template>
