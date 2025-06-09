<!-- пример:  Страница тэгов -->
<template>
  <div class="container mx-auto p-4">
    <h1 class="text-2xl font-bold mb-4">Тег: {{ tag.name }}</h1>
    <BookGrid :books="books" />
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import BookGrid from '@/components/BookGrid.vue'

const route = useRoute()
const books = ref([])
const tag = ref({})

onMounted(async () => {
  const id = route.params.id.split('-').at(-1)
  const res = await fetch(`/api/tags/${id}`)
  const data = await res.json()
  books.value = data.books
  tag.value = data
})
</script>