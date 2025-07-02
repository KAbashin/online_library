<template>
  <div class="p-4 space-y-6">
    <div v-if="loading" class="text-gray-500">Загрузка...</div>

    <div v-else-if="error">
      <ErrorBanner :message="error" @retry="loadTag" />
    </div>

    <div v-else>
      <h1 class="text-2xl font-bold mb-2" :style="{ color: tag?.color || '#000' }">
        Тег: {{ tag?.name }}
      </h1>
      <p class="text-gray-600 mb-4" v-if="tag?.description">{{ tag.description }}</p>

      <div v-if="books.length === 0" class="text-gray-500">Нет книг с этим тегом.</div>

      <BookGrid :books="books" />
    </div>
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
      fetchTagBooks(tagId),
    ])
    tag.value = tagRes.data
    books.value = booksRes.data
  } catch (err) {
    error.value = err.response?.data?.message || err.message || 'Ошибка загрузки'
  } finally {
    loading.value = false
  }
}

onMounted(loadTag)
</script>