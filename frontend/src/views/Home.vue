<template>
  <div class="flex flex-col min-h-screen bg-gray-100">
    <Header />
    <br>
    <main class="flex-grow container mx-auto px-4 py-6 space-y-10">
      <!-- Новинки -->
      <section>
        <NewReleases :books="newBooks" />
      </section>

      <!-- Родительские категории -->
      <section>
        <h2 class="text-2xl font-semibold mb-4">Категории</h2>
        <ParentCategories :categories="categories" />
      </section>

    </main>

    <br>
    <Footer />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import NewReleases from '@/components/NewReleases.vue'
import ParentCategories from '@/components/ParentCategories.vue'

import { getNewReleases } from '/src/api/book'
import { fetchRootCategories } from '/src/api/category'
import axios from '/src/api/axios'

const newBooks = ref([])
const books = ref([])
const categories = ref([])
const loading = ref(true)
const page = ref(1)
const pageSize = 12
const testResponse = ref(null)


async function fetchNewBooks() {
  try {
    newBooks.value = await getNewReleases()
  } catch (err) {
    console.error('Ошибка загрузки новинок:', err)
  }
}

async function fetchCategories() {
  try {
    categories.value = await fetchRootCategories()
  } catch (err) {
    console.error('Ошибка загрузки категорий:', err)
  }
}


onMounted(() => {
  fetchNewBooks()
  fetchCategories()

})
</script>