<script setup lang="ts">
import Usage from './Usage.vue';
import Power from './Power.vue';
import SlotPower from './SlotPower.vue';

</script>

<template>
  <div class="flex flex-col gap-4 p-4">

    <usage />

    <power />

    <slot-power />

  </div>
</template>
