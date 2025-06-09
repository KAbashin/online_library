<template>
  <div class="flex flex-col min-h-screen bg-gray-100">
    <Header />

    <main class="flex-grow container mx-auto px-4 py-6 space-y-10">
      <!-- Новинки -->
      <section>
        <NewReleases :books="newBooks" />
      </section>

      <!-- Родительские категории -->
      <section>
        <h2 class="text-2xl font-semibold mb-4">Категории</h2>
        <CategoryBlock :categories="categories" />
      </section>

      <!-- Все книги -->
      <section>
        <h2 class="text-2xl font-semibold mb-4">Популярные книги</h2>
        <BookGrid :books="books" :loading="loading" />
        <div class="flex justify-center mt-4">
          <button
              @click="loadMore"
              class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
          >
            Показать ещё
          </button>
        </div>
      </section>
    </main>

    <Footer />
  </div>
</template>

<script setup>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import NewReleases from '@/components/NewReleases.vue'
import BookGrid from '@/components/BookGrid.vue'
import CategoryBlock from '@/components/CategoryBlock.vue'
import axios from 'axios'
import { ref, onMounted } from 'vue'

const newBooks = ref([])
const books = ref([])
const categories = ref([])
const loading = ref(true)
const page = ref(1)
const pageSize = 12

async function fetchBooks() {
  loading.value = true
  const { data } = await axios.get(`/api/books?limit=${pageSize}&page=${page.value}`)
  books.value.push(...data)
  loading.value = false
}

async function fetchNewBooks() {
  const { data } = await axios.get(`/api/books/new`)
  newBooks.value = data
}

async function fetchCategories() {
  const { data } = await axios.get(`/api/categories?level=1`)
  categories.value = data
}

function loadMore() {
  page.value++
  fetchBooks()
}

onMounted(() => {
  fetchNewBooks()
  fetchCategories()
  fetchBooks()
})
</script>
