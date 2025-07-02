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
        <ParentCategories :categories="categories" />
      </section>

      <!-- Тестовый запрос -->
      <section class="mt-8">
        <h2 class="text-xl font-semibold mb-2">Тестовый запрос к бекенду</h2>
        <div v-if="testResponse" class="p-4 bg-green-100 rounded">
          Ответ: {{ testResponse }}
        </div>
        <div v-else class="p-4 bg-yellow-100 rounded">
          Нет данных
        </div>
      </section>
    </main>

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
const categories = ref([])
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

async function fetchTest() {
  try {
    const { data } = await axios.get('/test')
    testResponse.value = JSON.stringify(data)
  } catch (err) {
    testResponse.value = 'Ошибка: ' + err.message
  }
}

onMounted(() => {
  fetchNewBooks()
  fetchCategories()
  fetchTest()
})
</script>