<!-- пример: фильтры  -->
<template>
  <aside class="w-full md:w-64 p-4 bg-gray-100">
    <h2 class="text-lg font-bold mb-2">Фильтры</h2>
    <div class="mb-4">
      <label class="block mb-1">Год издания</label>
      <input type="number" v-model="filters.year" class="w-full border p-1">
    </div>
    <div class="mb-4">
      <label class="block mb-1">Язык</label>
      <input type="text" v-model="filters.language" class="w-full border p-1">
    </div>
    <button class="bg-blue-500 text-white px-4 py-1" @click="$emit('apply', filters)">Применить</button>
  </aside>
</template>
<script setup>
import { reactive } from 'vue'
const filters = reactive({ year: '', language: '' })
</script>

