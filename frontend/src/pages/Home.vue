<template>
  <div class="container mx-auto p-4">
    <!-- Новинки книг -->
    <NewReleases :books="newReleases" />

    <!-- Родительские категории -->
    <ParentCategories />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from '@/api/axios'

import NewReleases from '@/components/NewReleases.vue'
import ParentCategories from '@/components/ParentCategories.vue'

const newReleases = ref([])

onMounted(async () => {
  try {
    const response = await axios.get('/api/books/new-releases')  // пример эндпоинта для новинок
    newReleases.value = response.data
  } catch (error) {
    console.error('Ошибка при загрузке новинок:', error)
  }
})
</script>