<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from '@/api/axios'
import BookGrid from '@/components/BookGrid.vue'
import SkeletonLoader from '@/components/SkeletonLoader.vue'

const route = useRoute()
const books = ref([])
const loading = ref(true)

onMounted(async () => {
  const res = await axios.get(`/books?categoryId=${route.params.childId}`)
  books.value = res.data
  loading.value = false
})
</script>

<template>
  <div v-if="loading"><SkeletonLoader /></div>
  <BookGrid v-else :books="books" />
</template>
