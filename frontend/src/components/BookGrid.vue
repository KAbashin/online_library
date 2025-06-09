<!-- пример: сетка книг  -->
<template>
  <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
    <template v-if="loading">
      <div v-for="n in 8" :key="n" class="animate-pulse bg-gray-200 h-64 rounded"></div>
    </template>

    <template v-else>
      <div v-for="book in books" :key="book.id" class="border rounded p-2 bg-white shadow">
        <router-link :to="`/book/${book.title}-${book.id}`" class="block hover:underline">
          <img
              v-lazy="book.cover_url || placeholder"
              alt="cover"
              class="w-full h-48 object-cover mb-2"
          >
          <div class="font-semibold">{{ book.title }}</div>
        </router-link>
      </div>
    </template>
  </div>
</template>

<script setup>
const props = defineProps({
  books: Array,
  loading: Boolean
})
const placeholder = '/placeholder.jpg'
</script>
