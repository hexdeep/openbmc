<script setup lang="ts">
import {usePersistedStore} from '@/stores/persisted';
import {request} from '@/utils/axios';
import {ArrowLeft, ArrowRight, Delete, Download} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {watchEffect} from 'vue';
import {useI18n} from 'vue-i18n';
import {useRoute} from 'vue-router';

const route = useRoute()
const persisted = usePersistedStore()

interface Item {
  name: string;
  path: string;
  size: number;
  isDir: boolean;
}

const { t } = useI18n({ messages: {
  zh: {
    fileListTitle: '文件列表',
    name: '名称',
    path: '路径',
    size: '大小',
    operation: '操作',
    folderName: '文件夹名称',
    createFolder: '创建文件夹',
    uploadFile: '上传文件',
  },
} })

const items = ref<Item[]>([])
const folderName = ref('')

const load = async () => items.value = await request('GET', '/folder', { path: route.query.path })

load()
watchEffect(load)

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
      <div class="flex gap-4 items-center">
        <el-button
          v-if="$route.query.path"
          circle
          :icon="ArrowLeft"
          @click="$router.push(`/file?path=${removeLastSegment($route.path)}`)"
        />
        <div class="font-bold text-sm text-subtle">
          {{route.query.path}}
        </div>
        <el-input v-model="folderName" class="!w-64" :placeholder="t('folderName')">
          <template #append>
            <el-button @click="request('POST', '/folder', { name: folderName }).then(ok => ok && load())">
              {{t('createFolder')}}
            </el-button>
          </template>
        </el-input>
        <el-upload
          class="ms-auto"
          :action="`${persisted.serverAddr}/file`"
          :on-success="load"
        >
          <el-button>
            {{t('uploadFile')}}
          </el-button>
        </el-upload>
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
        <el-table-column :label="t('size')" prop="size" :formatter="formatSize" />
        <el-table-column :label="t('operation')">
          <template #default="{ row }">
            <el-button
              v-if="row.isDir"
              @click="$router.push(`/file?path=${row.path}`)"
              :icon="ArrowRight"
              circle
            />
            <el-button v-else :icon="Download" circle />
            <el-button
              @click="request('POST', '/delete-folder', { path: row.path }).then(ok => ok && load())"
              type="danger"
              :icon="Delete"
              circle
            />
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>
