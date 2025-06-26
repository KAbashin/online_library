<template>
  <div class="flex flex-col md:flex-row gap-4">
    <img
        :src="book.cover_url || placeholder"
        alt="Обложка"
        class="w-full md:w-60 object-cover rounded"
    />
    <div class="flex-1 space-y-2">
      <h1 class="text-2xl font-bold">{{ book.title }}</h1>
      <div class="text-gray-700 italic">{{ book.description }}</div>
      <div>
        <strong>Авторы:</strong>
        <router-link
            v-for="author in book.authors"
            :key="author.id"
            :to="`/author/${author.name}-${author.id}`"
            class="text-blue-600 hover:underline mr-2"
        >{{ author.name }}</router-link>
      </div>
      <div>
        <strong>Теги:</strong>
        <router-link
            v-for="tag in book.tags"
            :key="tag.id"
            :to="`/${tag.name}-${tag.id}`"
            class="bg-blue-100 text-blue-800 text-sm px-2 py-1 rounded mr-2"
        >{{ tag.name }}</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({ book: Object })
const placeholder = '/placeholder.jpg'
</script>