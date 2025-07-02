<!--  книги из конкретной категории + новинки -->
<template>
  <div class="p-4 space-y-6 max-w-4xl mx-auto">
    <div v-if="loading" class="text-gray-500">Загрузка...</div>
    <div v-else-if="error" class="text-red-500">
      Ошибка: {{ error }}
      <button @click="loadTag" class="ml-4 underline">Повторить</button>
    </div>
    <div v-else>
      <h1 class="text-2xl font-bold mb-4" :style="{ color: tag.color || '#000' }">
        {{ tag.name }}
      </h1>
      <p class="mb-6 text-gray-600" v-if="tag.description">{{ tag.description }}</p>

      <h2 class="text-xl font-semibold mb-2">Книги с этим тегом:</h2>
      <div v-if="books.length === 0" class="text-gray-500">Нет книг с этим тегом.</div>
      <ul class="space-y-4">
        <li
            v-for="book in books"
            :key="book.id"
            class="flex items-center space-x-4 border p-3 rounded hover:bg-gray-50"
        >
          <router-link :to="`/book/${book.id}`" class="flex items-center space-x-4">
            <img
                v-if="book.cover_image_url || book.cover_url"
                :src="book.cover_image_url || book.cover_url"
                alt="Обложка книги"
                class="w-16 h-24 object-cover rounded"
            />
            <div>
              <div class="text-lg font-semibold text-blue-600 hover:underline">
                {{ book.title }}
                <span v-if="book.publish_year"> ({{ book.publish_year }})</span>
              </div>
              <div class="text-sm text-gray-600">
                {{ authorsList(book.authors) }}
              </div>
            </div>
          </router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { fetchTag, fetchTagBooks } from '/src/api/tag'

const route = useRoute()
const tagId = route.params.id

const tag = ref(null)
const books = ref([])
const loading = ref(true)
const error = ref(null)

function authorsList(authors) {
  if (!authors || authors.length === 0) return 'Автор неизвестен'
  return authors.map(a => a.name).join(', ')
}

async function loadTag() {
  loading.value = true
  error.value = null
  try {
    const [tagRes, booksRes] = await Promise.all([
      fetchTag(tagId),
      fetchTagBooks(tagId),
    ])
    tag.value = tagRes.data
    books.value = booksRes.data
  } catch (err) {
    error.value = err.response?.data?.message || err.message || 'Ошибка загрузки данных'
  } finally {
    loading.value = false
  }
}

onMounted(loadTag)
</script>