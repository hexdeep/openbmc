<script setup lang="ts">
import Usage from './Usage.vue';
import Fan from './Fan.vue';
import Description from './Description.vue';
import Power from './Power.vue';
import OpticalPort from './OpticalPort.vue';

</script>

<template>
  <div class="flex flex-col lg:flex-row gap-4 p-4">

    <div class="basis-2/3 flex flex-col gap-4">

      <usage />

      <fan />

    </div>

    <div class="basis-1/3 flex flex-col gap-4">

      <description />

      <power />

      <optical-port />

    </div>

  </div>
</template>
