<script setup lang="ts">
import {DocumentAdd, DocumentCopy, Monitor, Phone, Picture, Setting} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {useI18n} from 'vue-i18n';

const isCollapse = ref(false)

const { t } = useI18n({ messages: {
  zh: {
    dashboard: '仪表板',
    deviceList: '设备列表',
    imageList: '镜像列表',
    hideNavbar: '隐藏侧边栏',
    fileManage: '文件管理',
    operationLogs: '操作日志',
    settings: '设置',
  },
} })

interface NavbarItem {
  label: string;
  icon: object;
  path: string;
}

const navbarItems: NavbarItem[] = [
  {
    label: t('dashboard'),
    icon: Monitor,
    path: '/dashboard',
  },
  {
    label: t('deviceList'),
    icon: Phone,
    path: '/devices',
  },
  {
    label: t('imageList'),
    icon: Picture,
    path: '/images',
  },
  {
    label: t('fileManage'),
    icon: DocumentCopy,
    path: '/file',
  },
  {
    label: t('operationLogs'),
    icon: DocumentAdd,
    path: '/logs',
  },
  {
    label: t('settings'),
    icon: Setting,
    path: '/settings',
  },
]
</script>

<template>
  <div class="h-screen flex">
    <el-menu router class="shrink-0 flex flex-col" :collapse="isCollapse">
      <div class="text-lg font-bold p-4 text-center">
        <el-icon v-if="isCollapse">
          <Menu />
        </el-icon>
        <span v-else>HexDeepBMC</span>
      </div>
      <el-menu-item v-for="item in navbarItems" :key="item.path" :index="item.path">
        <el-icon>
          <component :is="item.icon" />
        </el-icon>
        <span>{{t(item.label)}}</span>
      </el-menu-item>
      <el-menu-item class="mt-auto" @click="isCollapse = !isCollapse">
        <el-icon>
          <Hide v-if="isCollapse" />
          <View v-else />
        </el-icon>
        <span>{{t('hideNavbar')}}</span>
      </el-menu-item>
    </el-menu>
    <div class="grow min-w-0 overflow-auto">
      <router-view>
      </router-view>
    </div>
  </div>
</template>
