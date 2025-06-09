<!-- пример: список родительских категорий  -->
<template>
  <div class="space-y-4">
    <h1 class="text-2xl font-bold">Категории</h1>
    <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
      <router-link
          v-for="category in categories"
          :key="category.id"
          :to="`/category/${category.name}-${category.id}`"
          class="block p-4 bg-white shadow hover:shadow-md rounded"
      >
        {{ category.name }}
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const categories = ref([])

onMounted(async () => {
  const res = await fetch('/api/categories?level=parent')
  categories.value = await res.json()
})
</script>
