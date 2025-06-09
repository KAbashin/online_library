<!-- пример:  (страница книги) -->
<template>
  <div class="container mx-auto p-4">
    <div class="flex flex-col md:flex-row">
      <img :src="book.cover_url" class="w-48 h-72 object-cover mb-4 md:mr-6">
      <div>
        <h1 class="text-2xl font-bold">{{ book.title }}</h1>
        <p v-if="book.description">{{ book.description }}</p>
        <p><strong>Автор:</strong> <router-link :to="`/author/${author.name}-${author.id}`">{{ author.name }}</router-link></p>
        <p><strong>Год:</strong> {{ book.publish_year }}</p>
        <TagList :tags="book.tags" />
        <BookFiles :files="book.files" />
      </div>
    </div>
    <CommentList :comments="book.comments" />
  </div>
</template>
<script setup>
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import BookFiles from '@/components/BookFiles.vue'
import CommentList from '@/components/CommentList.vue'
import TagList from '@/components/TagList.vue'

const route = useRoute()
const book = ref({})
const author = ref({})

onMounted(async () => {
  const id = route.params.id.split('-').at(-1)
  const res = await fetch(`/api/books/${id}`)
  const data = await res.json()
  book.value = data
  author.value = data.authors[0] || {}
})
</script>
