<template>
  <div class="mt-6 space-y-4">
    <div v-if="extras.similar_books?.length">
      <h2 class="text-xl font-semibold mb-2">Похожие книги</h2>
      <BookGrid :books="extras.similar_books" />
    </div>

    <div v-if="extras.comments?.length">
      <h2 class="text-xl font-semibold mb-2">Комментарии</h2>
      <div class="space-y-2">
        <div v-for="comment in extras.comments" :key="comment.id" class="p-2 border rounded">
          <div class="font-semibold">{{ comment.user_name }}</div>
          <div class="text-sm text-gray-600">{{ comment.content }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import BookGrid from './BookGrid.vue'
const props = defineProps({ extras: Object })
</script>