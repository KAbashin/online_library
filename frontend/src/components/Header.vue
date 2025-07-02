<template>
  <header class="bg-gray-800 text-white p-4 grid grid-cols-[1fr_auto] items-start">
    <!-- Левая колонка: Онлайн библиотека -->
    <div>
      <router-link to="/" class="text-2xl font-bold leading-tight hover:underline">
        Онлайн<br />библиотека
      </router-link>
    </div>

    <!-- Правая колонка: навигация по ролям -->
    <div class="grid grid-rows-3 gap-1 text-right">
      <router-link
          v-if="isAdmin"
          to="/adminbackdoor"
          class="hover:underline"
      >
        Админ
      </router-link>

      <router-link
          v-if="role === 'user' || isAdmin"
          to="/profile"
          class="hover:underline"
      >
        Профиль
      </router-link>

      <button
          v-if="role"
          @click="logout"
          class="text-sm text-red-300 hover:underline"
      >
        Выйти
      </button>
    </div>
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
