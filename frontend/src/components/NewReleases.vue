<!-- пример: новинки  -->
<template>
  <div>
    <h2 class="text-xl font-bold mb-3">Новинки</h2>
    <BookGrid :books="books" />
  </div>
</template>
<script setup>
import BookGrid from './BookGrid.vue'
const props = defineProps({ books: Array })
</script>
