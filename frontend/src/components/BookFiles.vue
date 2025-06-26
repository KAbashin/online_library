<!-- пример:  файлы -->
<template>
  <div v-if="files.length" class="mt-4">
    <h3 class="text-lg font-semibold">Файлы</h3>
    <ul>
      <li v-for="file in files" :key="file.id">
        <a :href="file.url" target="_blank" class="text-blue-600 hover:underline">{{ file.name }}</a>
      </li>
    </ul>
  </div>
</template>
<script setup>
const props = defineProps({ files: Array })
</script>
