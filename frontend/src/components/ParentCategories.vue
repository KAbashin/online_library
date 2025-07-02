<template>
  <div>
    <h2 class="text-xl font-bold mb-3">Родительские категории</h2>
    <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
      <router-link
          v-for="cat in categories"
          :key="cat.id"
          :to="`/category/${cat.name}-${cat.id}`"
          class="block p-4 bg-white shadow rounded hover:bg-gray-50"
      >
        <div class="text-lg font-semibold">{{ cat.name }}</div>
      </router-link>
    </div>
  </div>
</template>

<script setup>
defineProps({
  categories: {
    type: Array,
    default: () => []
  }
})
</script>

