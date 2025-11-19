<script setup lang="ts">
import {request} from '@/utils/axios';
import {ArrowLeft} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {watchEffect} from 'vue';
import {useI18n} from 'vue-i18n';
import {useRoute} from 'vue-router';

const route = useRoute()

interface Item {
  name: string;
  path: string;
  size: number;
  isDir: boolean;
}

const { t } = useI18n({ messages: {
  zh: {
    fileListTitle: '文件列表',
    enter: '进入',
    download: '下载',
    name: '名称',
    path: '路径',
    size: '大小',
    operation: '操作',
  },
} })

const items = ref<Item[]>([])

watchEffect(async () => items.value = await request('GET', '/folder', { path: route.query.path }))

function removeLastSegment(path: string): string {
  const index = path.lastIndexOf("/");
  return index === -1 ? path : path.slice(0, index);
}

function formatSize(row: Item) {
  if (row.isDir) return '';

  const size = row.size;
  if (size < 1024) return size + ' B';
  if (size < 1024 * 1024) return (size / 1024).toFixed(1) + ' KB';
  if (size < 1024 * 1024 * 1024) return (size / 1024 / 1024).toFixed(1) + ' MB';
  return (size / 1024 / 1024 / 1024).toFixed(1) + ' GB';
}

</script>

<template>
  <div class="p-4 max-w-5xl mx-auto flex flex-col gap-4">
    <div class="card flex flex-col gap-4">
      <div class="text-lg">
        {{t('fileListTitle')}}
      </div>
      <div class="flex gap-4">
        <el-button
          v-if="$route.query.path !== ''"
          circle
          :icon="ArrowLeft"
          @click="$router.push(`/file?path=${removeLastSegment($route.path)}`)"
        />
      </div>
      <el-table :data="items">
        <el-table-column :label="t('name')" prop="name">
          <template #default="{ row }">
            <el-icon class="mr-1">
              <Folder v-if="row.isDir" />
              <Document v-else />
            </el-icon>
            {{ row.name }}
          </template>
        </el-table-column>
        <el-table-column :label="t('path')" prop="path" />
        <el-table-column :label="t('size')" prop="size" :formatter="formatSize" />
        <el-table-column :label="t('operation')">
          <template #default="{ row }">
            <el-button v-if="row.isDir" @click="$router.push(`/file?path=${row.path}`)">
              {{t('enter')}}
            </el-button>
            <el-button v-else>
              {{t('download')}}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>
