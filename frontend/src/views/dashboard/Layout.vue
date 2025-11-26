<script setup lang="ts">
import Usage from './Usage.vue';
import Fan from './Fan.vue';
import Power from './Power.vue';
import OpticalPort from './OpticalPort.vue';
import SlotPower from './SlotPower.vue';
import Temporary from './Temporary.vue';

</script>

<template>
  <div class="flex flex-col gap-4 p-4">

    <usage />

    <power />

    <slot-power />

    <fan />

    <temporary />

    <optical-port />

  </div>
</template>
