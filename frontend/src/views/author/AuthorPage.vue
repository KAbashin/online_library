<template>
  <div class="p-4 max-w-5xl mx-auto space-y-6">
    <div v-if="loading">
      <SkeletonLoader />
    </div>

    <div v-else>
      <ErrorBanner v-if="error" :message="error" @retry="loadAuthor" />

      <div v-if="author" class="space-y-2">
        <h1 class="text-2xl font-bold">{{ author.name }}</h1>
        <p class="text-gray-600 whitespace-pre-line">{{ author.bio }}</p>
      </div>

      <section>
        <h2 class="text-xl font-semibold mb-2">Книги автора:</h2>
        <BookGrid :books="books" />
      </section>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { fetchAuthor, fetchAuthorBooks } from '/src/api/author'
import BookGrid from '@/components/BookGrid.vue'
import ErrorBanner from '@/components/ErrorBanner.vue'
import SkeletonLoader from '@/components/SkeletonLoader.vue'

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
    error.value = err.response?.data?.message || err.message
  } finally {
    loading.value = false
  }
}

onMounted(loadAuthor)
</script>
