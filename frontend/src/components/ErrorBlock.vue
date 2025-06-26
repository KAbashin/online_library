<template>
  <div class="bg-red-100 text-red-700 p-4 rounded mb-4">
    <p class="mb-2">{{ message }}</p>
    <button @click="$emit('retry')" class="text-sm underline">Повторить</button>
  </div>
</template>

<script setup>
const props = defineProps({
  message: String
})
</script>