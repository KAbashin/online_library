<!-- пример:  Страница тэгов -->
<template>
  <div>
    <h1 class="text-2xl font-bold mb-2">{{ tag?.name }}</h1>
    <p class="text-gray-600 mb-4">{{ tag?.description }}</p>

    <ErrorBanner
        v-if="error"
        :message="error"
        @retry="loadTag"
    />

    <BookGrid :books="books" :loading="loading" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { fetchTag, fetchTagBooks } from '@/api/tag'
import BookGrid from '@/components/BookGrid.vue'
import ErrorBanner from '@/components/ErrorBanner.vue'

const route = useRoute()
const tagId = route.params.id

const tag = ref(null)
const books = ref([])
const loading = ref(true)
const error = ref(null)

async function loadTag() {
  error.value = null
  loading.value = true
  try {
    const [tagRes, booksRes] = await Promise.all([
      fetchTag(tagId),
      fetchTagBooks(tagId)
    ])
    tag.value = tagRes.data
    books.value = booksRes.data
  } catch (err) {
    error.value = 'Не удалось загрузить тег'
  } finally {
    loading.value = false
  }
}

onMounted(loadTag)
</script>
