<template>
  <div v-if="loading"><SkeletonLoader /></div>
  <CategoryBlock v-else :categories="categories" />
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from '@/api/axios'
import CategoryBlock from '@/components/CategoryBlock.vue'
import SkeletonLoader from '@/components/SkeletonLoader.vue'

const categories = ref([])
const loading = ref(true)

onMounted(async () => {
  const res = await axios.get('/categories?parentId=null')
  categories.value = res.data.map(c => ({
    ...c,
    to: `/category/${c.name}-${c.id}`
  }))
  loading.value = false
})
</script>