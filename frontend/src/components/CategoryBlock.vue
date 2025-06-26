<!--  отображение дочерней категории  -->
<template>
  <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
    <router-link
        v-for="category in categories"
        :key="category.id"
        :to="category.to"
        class="p-4 bg-white rounded shadow hover:bg-gray-100"
    >
      {{ category.name }}
    </router-link>
  </div>
</template>

<script setup>
defineProps({ categories: Array })
</script>