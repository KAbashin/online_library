<!-- пример:   отображение дочерней категории  -->
<template>
  <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
    <router-link
        v-for="cat in categories"
        :key="cat.id"
        :to="`/category/${cat.name}-${cat.id}`"
        class="block p-4 bg-white shadow rounded hover:bg-gray-50"
    >
      <div class="text-lg font-semibold">{{ cat.name }}</div>
      <p class="text-sm text-gray-600 mt-1">{{ cat.description }}</p>
    </router-link>
  </div>
</template>

<script setup>
const props = defineProps({ categories: Array })
</script>