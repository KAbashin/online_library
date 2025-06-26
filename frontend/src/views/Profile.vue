<!-- пример:  страница профиля пользователя -->
<template>
  <div class="p-4 max-w-4xl mx-auto">
    <h1 class="text-2xl font-bold mb-4">Профиль пользователя</h1>

    <div v-if="loading" class="space-y-4">
      <div class="h-6 bg-gray-200 rounded w-1/2 animate-pulse"></div>
      <div class="h-4 bg-gray-200 rounded w-1/3 animate-pulse"></div>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div v-for="n in 8" :key="n" class="h-64 bg-gray-200 rounded animate-pulse"></div>
      </div>
    </div>

    <div v-else>
      <div class="mb-6">
        <p><strong>Имя:</strong> {{ user.name }}</p>
        <p><strong>Email:</strong> {{ user.email }}</p>
        <p><strong>Роль:</strong> {{ user.role }}</p>
      </div>

      <h2 class="text-xl font-semibold mb-2">Мои книги</h2>
      <BookGrid :books="userBooks" :loading="loadingBooks" />
    </div>

    <ErrorBanner v-if="error" :message="error" @retry="fetchProfile" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getUserProfile } from '@/api/user'
import BookGrid from '@/components/BookGrid.vue'
import ErrorBanner from '@/components/ErrorBanner.vue'

const route = useRoute()
const userId = route.params.id

const user = ref({})
const userBooks = ref([])
const loading = ref(true)
const loadingBooks = ref(true)
const error = ref(null)

async function fetchProfile() {
  loading.value = true
  error.value = null
  try {
    const data = await getUserProfile(userId)
    user.value = data.user
    userBooks.value = data.books
  } catch (err) {
    error.value = err.message || 'Ошибка загрузки профиля'
  } finally {
    loading.value = false
    loadingBooks.value = false
  }
}

onMounted(fetchProfile)
</script>

<style scoped>
</style>
