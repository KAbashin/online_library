<!-- пример:   -->
<template>
  <div class="mt-6">
    <h3 class="text-lg font-semibold">Комментарии</h3>
    <div v-if="!comments.length" class="text-gray-500">Нет комментариев</div>
    <ul>
      <li v-for="comment in comments" :key="comment.id" class="border-b py-2">
        <strong>{{ comment.user_name }}</strong>: {{ comment.text }}
      </li>
    </ul>
  </div>
</template>
<script setup>
const props = defineProps({ comments: Array })
</script>

