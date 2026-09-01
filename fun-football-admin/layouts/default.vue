<template>
  <div class="min-h-screen bg-gray-50 flex font-sans">
    
    <!-- Sidebar Admin (Penuh ke bawah & Sticky) -->
    <aside class="w-64 h-screen sticky top-0 bg-gradient-to-b from-gray-900 via-gray-900 to-red-950 text-white flex flex-col shadow-xl z-20">
      
      <!-- Logo & Judul -->
      <div class="p-5 flex items-center gap-3 border-b border-gray-800">
        <div class="w-10 h-10 bg-white rounded-full p-0.5 shadow flex items-center justify-center overflow-hidden border border-orange-500 shrink-0">
          <img src="/logo-jmt.png" class="w-full h-full object-cover" alt="Logo" />
        </div>
        <div>
          <h1 class="font-extrabold text-base tracking-wide leading-tight">JMT Admin</h1>
          <p class="text-[10px] text-orange-400">Panel Pengurus</p>
        </div>
      </div>

      <!-- Navigasi Menu -->
      <nav class="flex-1 p-4 space-y-2 overflow-y-auto">
        <NuxtLink to="/" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-300 hover:bg-white/10 transition" exactActiveClass="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-md">
          <Icon name="ph:squares-four-bold" class="text-xl" />
          Dashboard
        </NuxtLink>
        <NuxtLink to="/jadwal" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-300 hover:bg-white/10 transition" exactActiveClass="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-md">
          <Icon name="ph:calendar-plus-bold" class="text-xl" />
          Jadwal Main
        </NuxtLink>
        <NuxtLink to="/verifikasi" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-300 hover:bg-white/10 transition" exactActiveClass="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-md">
          <Icon name="ph:receipt-bold" class="text-xl" />
          Verifikasi Bayar
        </NuxtLink>

        <!-- Menu Verifikasi Deposit -->
        <NuxtLink to="/verifikasi-deposit" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-400 hover:text-white transition-colors" exact-active-class="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-lg">
          <Icon name="ph:wallet-bold" class="text-xl" />
          Verifikasi Deposit
        </NuxtLink>

        <!-- Garis Pemisah -->
        <hr class="border-gray-800 my-4" />

        <NuxtLink to="/admin-keuangan" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-300 hover:bg-white/10 transition" exactActiveClass="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-md">
          <Icon name="ph:wallet-bold" class="text-xl" />
          Dashboard Kas
        </NuxtLink>

        <!-- Menu Profil Admin -->
        <NuxtLink to="/profil-admin" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-400 hover:text-white transition-colors" exact-active-class="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-lg">
          <Icon name="ph:user-circle-bold" class="text-xl" />
          Profil Saya
        </NuxtLink>

        <!-- Garis Pemisah Sebelum Logout -->
        <hr class="border-gray-800 my-4" />

        <!-- Tombol Logout Admin -->
        <button @click="logoutAdmin" class="w-full flex items-center gap-3 px-4 py-3 rounded-xl text-red-400 hover:bg-red-500/10 transition font-medium text-left">
          <Icon name="ph:sign-out-bold" class="text-xl" />
          Keluar (Logout)
        </button>
      </nav>

      <div class="p-4 border-t border-gray-800 text-xs text-gray-500 text-center">
        JMT Sport v1.0
      </div>
    </aside>

    <!-- Konten Utama -->
    <div class="flex-1 flex flex-col min-w-0">
      <header class="bg-white border-b border-gray-200 h-16 flex items-center justify-between px-8 shadow-sm sticky top-0 z-10">
        <p class="text-sm font-semibold text-gray-600">Selamat datang, Pengurus! 👋</p>
        <div class="w-9 h-9 bg-gradient-to-tr from-red-600 to-orange-600 text-white rounded-full flex items-center justify-center font-bold shadow">
          A
        </div>
      </header>

      <main class="flex-1 p-8 overflow-y-auto">
        <slot />
      </main>
    </div>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const isLoggedIn = ref(false)

onMounted(() => {
  const status = localStorage.getItem('admin_logged_in')
  if (status === 'true') {
    isLoggedIn.value = true
  } else {
    // Jika belum login, arahkan ke halaman login
    navigateTo('/admin-login')
  }
})

const logoutAdmin = () => {
  if (confirm('Yakin ingin keluar dari Panel Admin?')) {
    localStorage.removeItem('admin_logged_in')
    navigateTo('/admin-login')
  }
}
</script>