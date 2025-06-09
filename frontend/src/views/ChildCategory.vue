<!--  книги из конкретной категории + новинки -->
<template>
  <div class="flex gap-6">
    <FilterSidebar :filters="filters" @apply="applyFilters" />
    <div class="flex-1 space-y-6">
      <BookGrid :books="books" />
      <NewReleases :books="newBooks" title="Новинки" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import FilterSidebar from '@/components/FilterSidebar.vue'
import BookGrid from '@/components/BookGrid.vue'
import NewReleases from '@/components/NewReleases.vue'

const route = useRoute()
const childId = Number(route.params.childId)

const filters = ref({})
const books = ref([])
const newBooks = ref([])

onMounted(async () => {
  const res = await fetch(`/api/books?category_id=${childId}&limit=20`)
  books.value = await res.json()

  const newRes = await fetch(`/api/books/new?category_id=${childId}`)
  newBooks.value = await newRes.json()
})

function applyFilters(newFilters) {
  filters.value = newFilters
  // Перезапросить книги с фильтрами
}
</script>