<!-- Скелетон (заглушка) на время загрузки.  -->
<template>
  <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
    <div v-for="n in count" :key="n" class="animate-pulse bg-gray-200 h-32 rounded"></div>
  </div>
</template>

<script setup>
const props = defineProps({
  count: {
    type: Number,
    default: 8
  }
})
</script>