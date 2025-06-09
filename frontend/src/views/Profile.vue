<!-- пример:  страница профиля пользователя -->
<template>
  <div class="container mx-auto p-4">
    <h1 class="text-2xl font-bold mb-4">Профиль</h1>
    <h2 class="text-xl font-semibold mt-6">Избранное</h2>
    <BookGrid :books="favorites" />
    <h2 class="text-xl font-semibold mt-6">Созданные книги</h2>
    <BookGrid :books="createdBooks" />
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import BookGrid from '@/components/BookGrid.vue'

const route = useRoute()
const userId = route.params.id.split('-').at(-1)
const favorites = ref([])
const createdBooks = ref([])

onMounted(async () => {
  const [favRes, createdRes] = await Promise.all([
    fetch(`/api/users/${userId}/favorites`),
    fetch(`/api/users/${userId}/created-books`)
  ])
  favorites.value = await favRes.json()
  createdBooks.value = await createdRes.json()
})
</script>