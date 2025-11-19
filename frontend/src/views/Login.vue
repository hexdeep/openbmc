<script setup lang="ts">
import {request} from '@/utils/axios';
import {Setting} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {useI18n} from 'vue-i18n';


const { t } = useI18n({ messages: {
  zh: {
    title: '欢迎使用 OpenBMC !',
    password: '密码',
  },
} })

const password = ref('')
</script>

<template>
  <div class="h-screen flex justify-center items-center bg-[url('/login-bg.png')] bg-center bg-cover">
    <div class="card flex flex-col gap-4">
      <div class="flex justify-between items-center">
        <div class="text-lg">{{t('title')}}</div>
        <el-button :icon="Setting" circle @click="$router.push('/settings')" />
      </div>
      <el-form>
        <el-form-item :label="t('password')">
          <el-input v-model="password" type="password" show-password />
        </el-form-item>
      </el-form>
      <div class="flex justify-end">
        <el-button @click="request('POST', '/login', {}, { password })">
          {{t('confirm')}}
        </el-button>
      </div>
    </div>
  </div>
</template>
