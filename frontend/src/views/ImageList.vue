<script setup lang="ts">
import {useI18n} from 'vue-i18n';


const { t } = useI18n({ messages: {
  zh: {
    imageListTitle: '镜像列表',
    imageName: '镜像名称',
    imageSize: '镜像大小',
    createdAt: '创建时间',
    operation: '操作',
    delete: '删除',
  },
} })

const testData = [
  {
    name: 'test',
    size: '1GB',
    createdAt: '2006-01-02 15:04:05',
  },
]
</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">

    <div class="card flex flex-col gap-4">
      <div class="text-lg">
        {{t('imageListTitle')}}
      </div>
      <el-table :data="testData">
        <el-table-column :label="t('imageName')" prop="name" />
        <el-table-column :label="t('imageSize')" prop="size" />
        <el-table-column :label="t('createdAt')" prop="createdAt" />
        <el-table-column :label="t('operation')">
          <el-button type="danger" size="small">
            {{t('delete')}}
          </el-button>
        </el-table-column>
      </el-table>
      <el-pagination layout="prev, pager, next, total" :total="10" />
    </div>

  </div>
</template>
