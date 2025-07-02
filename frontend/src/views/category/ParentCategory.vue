<template>
  <div>
    <h1 class="text-2xl font-bold mb-2">{{ category?.name }}</h1>
    <p class="mb-4 text-gray-600">{{ category?.description }}</p>

    <ErrorBanner
        v-if="error"
        :message="error"
        @retry="loadCategory"
    />

    <BookGrid :books="books" :loading="loading" />
  </div>
</template>

<script setup>
import {ref, onMounted} from 'vue'
import {useRoute} from 'vue-router'
import {fetchCategoryBooks, fetchCategoryInfo} from '@/api/category.js'
import BookGrid from '@/components/BookGrid.vue'
import ErrorBanner from '@/components/ErrorBanner.vue'

const route = useRoute()
const categoryId = route.params.parentId

const category = ref(null)
const books = ref([])
const loading = ref(true)
const error = ref(null)

async function loadCategory() {
  error.value = null
  loading.value = true
  try {
    const [catRes, booksRes] = await Promise.all([
      fetchCategoryInfo(categoryId),
      fetchCategoryBooks(categoryId)
    ])
    category.value = catRes.data
    books.value = booksRes.data
  } catch (err) {
    error.value = 'Не удалось загрузить категорию'
  } finally {
    loading.value = false
  }
}

onMounted(loadCategory)
</script>