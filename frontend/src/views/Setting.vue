<script setup lang="ts">
import {ArrowLeft} from '@element-plus/icons-vue';
import {useColorMode} from '@vueuse/core';
import {ref} from 'vue';
import {useI18n} from 'vue-i18n';

const i18n = useI18n({ messages: {
  zh: {
    settingsTitle: '设置',
    theme: '主题',
    dark: '暗色',
    light: '亮色',
    auto: '自动',
    resetPassword: '重设密码',
  },
} })

const { t } = i18n

const mode = useColorMode()

const isResetPasswordDialogOpen = ref(false)
</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    <div class="card flex items-center gap-4">
      <el-button :icon="ArrowLeft" @click="$router.back">
        {{t('back')}}
      </el-button>
      <div class="text-lg">
        {{t('settingsTitle')}}
      </div>
    </div>
    <el-form class="card" label-width="auto">
      <el-form-item :label="t('theme')">
        <el-radio-group v-model="mode">
          <el-radio-button :label="t('auto')" />
          <el-radio-button :label="t('light')" />
          <el-radio-button :label="t('dark')" />
        </el-radio-group>
      </el-form-item>
      <el-form-item label="语言 / Language">
        <el-radio-group v-model="i18n.locale.value">
          <el-radio-button label="简体中文" value="zh" />
          <el-radio-button label="English" value="en" />
        </el-radio-group>
      </el-form-item>
      <el-form-item>
        <el-button @click="isResetPasswordDialogOpen = !isResetPasswordDialogOpen">
          {{t('resetPassword')}}
        </el-button>
      </el-form-item>
    </el-form>
  </div>

  <el-dialog v-model="isResetPasswordDialogOpen" :title="t('resetPassword')">
  </el-dialog>
</template>
