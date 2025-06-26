<template>
  <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-2 rounded relative">
    <span>{{ message }}</span>
    <button class="ml-4 underline" @click="$emit('retry')">Повторить</button>
  </div>
</template>

<script setup>
defineProps({ message: String })
</script>