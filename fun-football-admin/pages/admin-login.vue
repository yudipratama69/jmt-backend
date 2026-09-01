<template>
  <div class="min-h-screen bg-gray-900 flex items-center justify-center p-5 font-sans">
    
    <!-- Kotak Form Login -->
    <div class="bg-white w-full max-w-md rounded-3xl p-8 shadow-2xl relative overflow-hidden">
      
      <!-- Aksen Gradasi di Atas Card -->
      <div class="absolute top-0 left-0 right-0 h-2 bg-gradient-to-r from-red-600 via-orange-600 to-amber-500"></div>

      <!-- Header Logo / Judul -->
      <div class="text-center mb-6 mt-2">
        <div class="w-28 h-28 mx-auto flex items-center justify-center mb-2">
          <img src="/logo-jmt.png" class="w-full h-full object-contain" alt="Logo JMT" />
        </div>
        <h2 class="text-2xl font-extrabold text-gray-800">Admin JMT Sport</h2>
        <p class="text-xs text-gray-400 mt-1">Masukkan password pengurus untuk masuk</p>
      </div>

      <!-- Form -->
      <form @submit.prevent="handleLogin" class="space-y-5">
        <div>
          <label class="block text-xs font-bold uppercase tracking-wider text-gray-500 mb-2">Password Admin</label>
          <input v-model="password" type="password" placeholder="••••••••" class="w-full border border-gray-200 rounded-xl p-3.5 focus:ring-2 focus:ring-orange-500 outline-none text-sm bg-gray-50/50" required />
        </div>

        <button type="submit" class="w-full bg-gradient-to-r from-red-600 to-orange-600 hover:from-red-700 hover:to-orange-700 text-white font-bold py-4 rounded-xl shadow-md transition duration-200 text-sm active:scale-95">
          Masuk ke Panel Admin
        </button>
      </form>

      <div class="mt-8 text-center">
        <NuxtLink to="/player/login" class="text-xs text-gray-400 hover:text-orange-600 font-medium transition">
          ← Kembali ke Halaman Player
        </NuxtLink>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

// MEMATIKAN LAYOUT ADMIN DI HALAMAN INI
definePageMeta({ layout: false })

const password = ref('')

const handleLogin = () => {
  if (password.value === 'admin123') {
    localStorage.setItem('admin_logged_in', 'true')
    alert('Login Admin Berhasil!')
    navigateTo('/')
  } else {
    alert('Password Admin Salah!')
  }
}
</script>