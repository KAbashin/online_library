<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-gray-100">
    <div class="bg-white p-6 rounded shadow-md w-full max-w-sm">
      <h1 class="text-xl font-bold mb-4 text-center">Регистрация</h1>
      <form @submit.prevent="register">
        <div class="mb-4">
          <label class="block text-gray-700 mb-1">Имя</label>
          <input v-model="name" type="text" required class="w-full px-3 py-2 border rounded" />
        </div>
        <div class="mb-4">
          <label class="block text-gray-700 mb-1">Email</label>
          <input v-model="email" type="email" required class="w-full px-3 py-2 border rounded" />
        </div>
        <div class="mb-4">
          <label class="block text-gray-700 mb-1">Пароль</label>
          <input v-model="password" type="password" required class="w-full px-3 py-2 border rounded" />
        </div>
        <div class="mb-4">
          <label class="block text-gray-700 mb-1">О себе</label>
          <textarea v-model="bio" class="w-full px-3 py-2 border rounded" rows="3" />
        </div>
        <button type="submit" class="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700">
          Зарегистрироваться
        </button>
      </form>
      <p v-if="error" class="text-red-500 mt-4 text-sm text-center">{{ error }}</p>

      <div class="mt-6 text-center">
        <p class="text-gray-600 text-sm">Уже есть аккаунт?</p>
        <button @click="goToLogin" class="mt-1 text-blue-600 hover:underline text-sm">
          Войти
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const name = ref('')
const email = ref('')
const password = ref('')
const bio = ref('')
const error = ref('')
const router = useRouter()
const API_URL = import.meta.env.VITE_API_URL

const register = async () => {
  error.value = ''

  try {
    const response = await fetch(`${API_URL}/auth/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        name: name.value,
        email: email.value,
        password: password.value,
        bio: bio.value,
      }),
    })

    if (!response.ok) {
      const errData = await response.json()
      throw new Error(errData.error || 'Ошибка регистрации')
    }

    // После регистрации перенаправляем на логин
    await router.push('/login')
  } catch (err) {
    error.value = err.message || 'Ошибка регистрации'
  }
}

const goToLogin = () => {
  router.push('/login')
}
</script>