<!-- пример:   список тегов-->
<template>
  <div class="mt-4">
    <h3 class="font-bold">Теги:</h3>
    <div class="flex flex-wrap gap-2 mt-1">
      <router-link
          v-for="tag in tags"
          :key="tag.id"
          :to="`/${tag.name}-${tag.id}/`"
          class="bg-gray-200 text-sm px-2 py-1 rounded hover:bg-gray-300"
      >
        #{{ tag.name }}
      </router-link>
    </div>
  </div>
</template>
<script setup>
const props = defineProps({ tags: Array })
</script>

