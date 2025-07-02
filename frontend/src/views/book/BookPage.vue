<template>
  <div class="p-4 space-y-6">
    <template v-if="loading">
      <div class="animate-pulse h-64 bg-gray-200 rounded"></div>
    </template>

    <template v-else-if="error">
      <ErrorBanner :message="error" @retry="fetchBook" />
    </template>

    <template v-else>
      <BookDetails :book="book" />
      <BookExtras v-if="extras" :extras="extras" />
    </template>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { getBook, getBookExtras } from '/src/api/book'
import BookDetails from '@/components/BookDetails.vue'
import BookExtras from '@/components/BookExtras.vue'
import ErrorBanner from '@/components/ErrorBanner.vue'

const route = useRoute()
const id = route.params.id

const book = ref(null)
const extras = ref(null)
const loading = ref(true)
const error = ref(null)

async function fetchBook() {
  loading.value = true
  error.value = null
  try {
    book.value = await getBook(id)

    // Загружаем дополнительные данные (extras) параллельно
    getBookExtras(id)
        .then(data => {
          extras.value = data
        })
        .catch(err => {
          console.warn('Ошибка загрузки extras:', err)
        })
  } catch (err) {
    error.value = 'Не удалось загрузить книгу'
    console.error(err)
  } finally {
    loading.value = false
  }
}

onMounted(fetchBook)
</script>