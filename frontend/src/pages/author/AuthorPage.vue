<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from '@/api/axios'
import BookGrid from '@/components/BookGrid.vue'
import SkeletonLoader from '@/components/SkeletonLoader.vue'

const route = useRoute()
const author = ref({})
const books = ref([])
const loading = ref(true)

onMounted(async () => {
  const res1 = await axios.get(`/authors/${route.params.id}`)
  const res2 = await axios.get(`/authors/${route.params.id}/books`)
  author.value = res1.data
  books.value = res2.data
  loading.value = false
})
</script>

<template>
  <div v-if="loading"><SkeletonLoader /></div>
  <div v-else>
    <h1>{{ author.name }}</h1>
    <BookGrid :books="books" />
  </div>
</template>