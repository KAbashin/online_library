<!--  дочерние категории и книги из них + новинки -->
<template>
  <div class="flex gap-6">
    <FilterSidebar :filters="filters" @apply="applyFilters" />
    <div class="flex-1 space-y-6">
      <CategoryBlock
          v-for="child in childCategories"
          :key="child.id"
          :category="child"
          :books="child.books"
      />
      <NewReleases :books="newBooks" title="Новинки" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import FilterSidebar from '@/components/FilterSidebar.vue'
import CategoryBlock from '@/components/CategoryBlock.vue'
import NewReleases from '@/components/NewReleases.vue'

const route = useRoute()
const parentId = Number(route.params.parentId)

const filters = ref({})
const childCategories = ref([])
const newBooks = ref([])

onMounted(async () => {
  const childrenRes = await fetch(`/api/categories/${parentId}/children?withBooks=true`)
  childCategories.value = await childrenRes.json()

  const newBooksRes = await fetch(`/api/books/new?parent_category_id=${parentId}`)
  newBooks.value = await newBooksRes.json()
})

function applyFilters(newFilters) {
  filters.value = newFilters
  // Перезапросить книги с фильтрами
}
</script>