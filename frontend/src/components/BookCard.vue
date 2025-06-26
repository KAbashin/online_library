<!-- пример:  карточка товара -->
<template>
  <div class="border p-4 rounded shadow hover:shadow-lg transition">
    <h2 class="font-bold text-lg">{{ book.title }}</h2>
    <p class="text-sm text-gray-600">Авторов: {{ book.authors?.map(a => a.name).join(', ') }}</p>
    <p class="mt-2">{{ book.description || 'Без описания' }}</p>
    <router-link :to="`/book/${book.id}`" class="text-blue-500 mt-4 inline-block">Подробнее</router-link>
  </div>
</template>

<script setup>
defineProps({
  book: Object
})
</script>