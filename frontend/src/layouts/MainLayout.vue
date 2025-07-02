<template>
  <div class="flex flex-col min-h-screen bg-gray-100">
    <Header />
    <main class="flex-grow container mx-auto px-4 py-6 space-y-10">
      <slot />                      <!-- сюда попадут все children, в том числе Home.vue -->
    </main>
    <Footer />
  </div>
</template>

<script setup>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
</script>