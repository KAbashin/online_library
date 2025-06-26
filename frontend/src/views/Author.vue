<template>
  <div>
    <h1 class="text-2xl font-bold mb-2">{{ author?.name }}</h1>
    <p class="text-gray-600 mb-4">{{ author?.bio }}</p>

    <ErrorBanner
        v-if="error"
        :message="error"
        @retry="loadAuthor"
    />

    <BookGrid :books="books" :loading="loading" />
  </div>
</template>

<script setup>
import {ref, onMounted} from 'vue'
import {useRoute} from 'vue-router'
import {fetchAuthor, fetchAuthorBooks} from '@/api/author'
import BookGrid from '@/components/BookGrid.vue'
import ErrorBanner from '@/components/ErrorBanner.vue'

const route = useRoute()
const authorId = route.params.id

const author = ref(null)
const books = ref([])
const loading = ref(true)
const error = ref(null)

async function loadAuthor() {
  error.value = null
  loading.value = true
  try {
    const [authorRes, booksRes] = await Promise.all([
      fetchAuthor(authorId),
      fetchAuthorBooks(authorId)
    ])
    author.value = authorRes.data
    books.value = booksRes.data
  } catch (err) {
    error.value = 'Не удалось загрузить автора'
  } finally {
    loading.value = false
  }
}

onMounted(loadAuthor)
</script>
