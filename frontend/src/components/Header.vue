<!-- пример:   -->
<template>
  <header class="bg-gray-800 text-white p-4 flex justify-between">
    <div class="font-bold text-xl">
      Онлайн библиотека
    </div>
    <nav class="space-x-4">
      <router-link to="/" class="hover:underline">Главная</router-link>
      <router-link v-if="role === 'user' || isAdmin" to="/profile" class="hover:underline">Профиль</router-link>
      <router-link v-if="isAdmin" to="/adminbackdoor" class="hover:underline">Админ</router-link>
      <button @click="logout" class="hover:underline">Выйти</button>
    </nav>
  </header>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()
const role = localStorage.getItem('role')
const isAdmin = role === 'admin' || role === 'superadmin'

function logout() {
  localStorage.removeItem('token')
  localStorage.removeItem('role')
  router.push('/login')
}
</script>