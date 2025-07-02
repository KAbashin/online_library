<!-- пример:   -->
<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-gray-100">
    <div class="bg-white p-6 rounded shadow-md w-full max-w-sm">
      <h1 class="text-xl font-bold mb-4 text-center">Вход</h1>
      <form @submit.prevent="login">
        <div class="mb-4">
          <label class="block text-gray-700 mb-1">Email</label>
          <input v-model="email" type="email" required class="w-full px-3 py-2 border rounded" />
        </div>
        <div class="mb-4">
          <label class="block text-gray-700 mb-1">Пароль</label>
          <input v-model="password" type="password" required class="w-full px-3 py-2 border rounded" />
        </div>
        <button type="submit" class="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700">
          Войти
        </button>
      </form>
      <p v-if="error" class="text-red-500 mt-4 text-sm text-center">{{ error }}</p>

      <div class="mt-6 text-center">
        <p class="text-gray-600 text-sm">Нет аккаунта?</p>
        <button @click="goToRegister" class="mt-1 text-blue-600 hover:underline text-sm">
          Зарегистрироваться
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { jwtDecode } from 'jwt-decode'

const email = ref('')
const password = ref('')
const error = ref('')
const router = useRouter()
const API_URL = import.meta.env.VITE_API_URL


const login = async () => {
  error.value = ''

  try {
    const response = await fetch(`${API_URL}/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email: email.value, password: password.value }),
    })

    if (!response.ok) throw new Error('Неверные данные')

    const data = await response.json()
    const token = data.token

    const decoded = jwtDecode(token)
    const role = decoded.role

    localStorage.setItem('token', token)
    localStorage.setItem('role', role)

    await router.push('/')
  } catch (err) {
    error.value = err.message || 'Ошибка входа'
  }
}

const goToRegister = () => {
  router.push('/register')
}
</script>
