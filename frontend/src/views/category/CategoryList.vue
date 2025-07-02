<template>
  <div v-if="loading">
    <SkeletonLoader />
  </div>
  <div v-else>
    <CategoryBlock :categories="categories" />
  </div>
</template>

<script setup async>
import { ref } from 'vue'
import CategoryBlock from '@/components/CategoryBlock.vue'
import SkeletonLoader from '@/components/SkeletonLoader.vue'
import { fetchRootCategories } from '/src/api/category'

const loading = ref(true)
let categories = []

try {
  const rawCategories = await fetchRootCategories()
  categories = rawCategories.map(c => ({
    ...c,
    to: `/category/${c.name}-${c.id}`,
  }))
} catch (e) {
  console.error('Ошибка при загрузке категорий:', e)
} finally {
  loading.value = false
}
</script>