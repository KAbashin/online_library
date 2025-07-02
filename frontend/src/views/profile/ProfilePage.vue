<template>
  <div class="p-4 max-w-5xl mx-auto space-y-8">
    <h1 class="text-2xl font-bold">Профиль пользователя</h1>

    <!-- Ошибка -->
    <ErrorBanner v-if="error" :message="error" @retry="fetchProfile" />

    <!-- Скелетон -->
    <div v-if="loading" class="space-y-4">
      <SkeletonLoader />
    </div>

    <!-- Профиль -->
    <div v-else>
      <div class="bg-white rounded shadow p-4">
        <p><strong>Имя:</strong> {{ user.name }}</p>
        <p><strong>Email:</strong> {{ user.email }}</p>
        <p><strong>Роль:</strong> {{ user.role }}</p>
      </div>

      <!-- Мои книги -->
      <section>
        <h2 class="text-xl font-semibold mb-2">Мои книги</h2>
        <BookGrid :books="books" />
      </section>

      <!-- Избранное -->
      <section>
        <h2 class="text-xl font-semibold mb-2">Избранное</h2>
        <BookGrid :books="favorites" />
      </section>

      <!-- Комментарии -->
      <section>
        <h2 class="text-xl font-semibold mb-2">Мои комментарии</h2>
        <ul class="list-disc list-inside text-gray-700">
          <li v-for="comment in comments" :key="comment.id">
            {{ comment.text }}
          </li>
        </ul>
      </section>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import {
  getUserProfile,
  getUserBooks,
  getUserFavoriteBooks,
  getUserComments
} from '/src/api/user'

import BookGrid from '@/components/BookGrid.vue'
import ErrorBanner from '@/components/ErrorBanner.vue'
import SkeletonLoader from '@/components/SkeletonLoader.vue'

const route = useRoute()
const userId = route.params.id

const user = ref({})
const books = ref([])
const favorites = ref([])
const comments = ref([])
const loading = ref(true)
const error = ref(null)

async function fetchProfile() {
  error.value = null
  loading.value = true
  try {
    const [userData, bookData, favoriteData, commentData] = await Promise.all([
      getUserProfile(userId),
      getUserBooks(userId),
      getUserFavoriteBooks(),
      getUserComments(userId)
    ])
    user.value = userData
    books.value = bookData
    favorites.value = favoriteData
    comments.value = commentData
  } catch (err) {
    error.value = err.response?.data?.message || err.message
  } finally {
    loading.value = false
  }
}

onMounted(fetchProfile)
</script>
