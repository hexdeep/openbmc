<script setup lang="ts">
import {request} from '@/utils/axios';

</script>

<template>
  <div class="card">
    <el-button @click="request('POST', '/bmc/opentty')">
      开启BMC TTY
    </el-button>
    <el-button @click="request('POST', '/bmc/closetty')">
      关闭BMC TTY
    </el-button>
    <el-button @click="request('POST', '/switch/opentty')">
      开启Switch TTY
    </el-button>
    <el-button @click="request('POST', '/switch/closetty')">
      关闭Switch TTY
    </el-button>
  </div>
</template>
