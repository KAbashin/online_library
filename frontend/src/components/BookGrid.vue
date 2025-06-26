<!--  сетка книг  -->
<template>
  <div class="grid grid-cols-2 sm:grid-cols-4 gap-4">
    <div v-for="book in books" :key="book.id" class="bg-white rounded shadow p-2">
      <router-link :to="`/book/${book.title}-${book.id}`">
        <img :src="book.cover_url" alt="" class="w-full h-auto" />
        <p class="mt-2">{{ book.title }}</p>
      </router-link>
    </div>
  </div>
</template>

<script setup>
defineProps({ books: Array })
</script>
