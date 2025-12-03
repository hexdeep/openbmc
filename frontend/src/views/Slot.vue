<script setup lang="ts">
import {request} from '@/utils/axios';
import {formatSize} from '@/utils/utils';
import {useEventSource} from '@vueuse/core';
import {computed, ref} from 'vue';
import {useI18n} from 'vue-i18n';

const { t } = useI18n({ messages: {
  zh: {
    slot: '插槽',
    status: '状态',
    deviceListTitle: '通电接口列表',
    temperature: '温度',
    mem: '内存',
    powerStatus: '电源状态',
    operation: '操作',
    load: '负载',
    powerOn: '上电',
    powerOff: '下电',
    detail: '详情',
    updateSystem: '升级',
    openTerminal: '打开终端',
    closeTerminal: '关闭终端',
    enterTerminal: '进入',
    uptime: '开机时间',
    powerOffConfirmTitle: '你确定要下电吗？',
    flash: '刷机',
    flashSelected: '刷机选中项',
    selectedSlot: '已选择插槽',
    image: '镜像',
    selectImage: '选择镜像',
  },
} })

const selected = ref<string[]>([])

const { data } = useEventSource('/api/powered-slot')
const poweredSlots = computed(() => {
  if (!data.value) return 0
  try {
    return JSON.parse(data.value)
  } catch {
    return 0
  }
})

function formatDuration(seconds: number): string {
    const hours = Math.floor(seconds / 3600);
    const minutes = Math.floor((seconds % 3600) / 60);

    const parts: string[] = [];

    if (hours > 0) {
        parts.push(`${hours}h`);
    }

    if (minutes > 0 || hours === 0) {
        // Show minutes always if there are no hours, e.g. "0min"
        parts.push(`${minutes}min`);
    }

    return parts.join(" ");
}

const toTerminal = (port: string) => window.open(`http://${window.location.hostname}:${port}`, '_blank')

const isDialogOpen = ref(false)

</script>

<template>
  <div class="flex flex-col gap-4 p-4">

    <div class="card flex flex-col gap-4">
      <div class="flex items-center gap-4">
        <div class="text-lg">{{t('deviceListTitle')}}</div>
        <el-button @click="isDialogOpen = true">
          {{t('flashSelected')}}
        </el-button>
      </div>
      <el-table :data="poweredSlots" class="rounded-1xl">
        <el-table-column :width="50">
          <template #default="{ row }">
            <el-checkbox
              :model-value="selected.some(v => v === row.slot)"
              @update:model-value="v => v ? selected.push(row.slot) : selected.splice(selected.indexOf(row.slot), 1)"
            />
          </template>
        </el-table-column>
        <el-table-column :label="t('slot')" prop="slot" :width="50" />
        <el-table-column :label="t('status')">
          <template #default="{ row }">
            <el-tag :type="row.active ? 'success' : 'danger'">
              {{row.active ? '运行' : row.ttyActive ? '终端开启' : '异常'}}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="IP" prop="ip" />
        <el-table-column label="MAC" prop="mac" />
        <el-table-column :label="t('uptime')" prop="uptime" :formatter="({ uptime }) => formatDuration(uptime)" />
        <el-table-column :label="t('load')" prop="load" :formatter="({ load }) => `${load} %`" />
        <el-table-column :label="t('mem')" width="200">
          <template #default="{ row }">
            {{formatSize(row.memUsed)}} /
            {{formatSize(row.memTotal)}} /
            {{Math.trunc((row.memUsed / row.memTotal) * 100)}} %
          </template>
        </el-table-column>
        <el-table-column
          :label="t('temperature')"
          prop="temp"
          :formatter="({ temp }) => `${temp} °C`"
        />
        <el-table-column :label="t('operation')" width="400">
          <template #default="{ row }">
            <el-popconfirm :title="t('powerOffConfirmTitle')" @confirm="request('POST', `/slot/${row.slot}/power-off`)">
              <template #reference>
                <el-button type="danger" size="small">
                  {{t('powerOff')}}
                </el-button>
              </template>
            </el-popconfirm>
            <el-button type="success" size="small" @click="request('POST', `/slot/${row.slot}/flash`)">
              {{t('updateSystem')}}
            </el-button>
            <el-button-group class="ml-2">
              <el-button v-if="row.ttyActive" type="warning" size="small" @click="request('POST', `/slot/${row.slot}/closetty`)">
                {{t('closeTerminal')}}
              </el-button>
              <el-button v-else type="primary" size="small" @click="request('POST', `/slot/${row.slot}/opentty`)">
                {{t('openTerminal')}}
              </el-button>
              <el-button type="success" size="small" @click="toTerminal(row.port)">
                {{t('enterTerminal')}}
              </el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </div>

  </div>
  <el-dialog v-model="isDialogOpen" :title="t('flash')">
    <el-form>
      <el-form-item :label="t('selectedSlot')">
        {{selected.join(', ')}}
      </el-form-item>
      <el-form-item :label="t('image')">
        <el-select :placeholder="t('selectImage')">
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="isDialogOpen = false">
        {{t('cancel')}}
      </el-button>
      <el-button>
        {{t('confirm')}}
      </el-button>
    </template>
  </el-dialog>
</template>
