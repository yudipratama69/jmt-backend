<template>
  <div class="min-h-screen theme-bg-page theme-text-main flex flex-col md:flex-row font-sans transition-colors duration-300">
    
    <!-- Mobile Overlay -->
    <div v-if="isSidebarOpen" @click="isSidebarOpen = false" class="fixed inset-0 bg-black/50 z-20 md:hidden"></div>

    <!-- Sidebar Admin -->
    <aside 
      :class="isSidebarOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0'"
      class="w-64 h-screen fixed md:sticky top-0 bg-gradient-to-b from-gray-900 via-gray-900 to-red-950 text-white flex flex-col shadow-xl z-30 transition-transform duration-300 ease-in-out">
      
      <!-- Logo & Judul -->
      <div class="p-5 flex items-center justify-between border-b border-gray-800">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 bg-white rounded-full p-0.5 shadow flex items-center justify-center overflow-hidden border border-orange-500 shrink-0">
            <img :src="'/logo-jmt.png'" class="w-full h-full object-contain" alt="Logo" />
          </div>
          <div>
            <h1 class="font-extrabold text-base tracking-wide leading-tight">JMT Admin</h1>
            <p class="text-[10px] text-orange-400">Panel Pengurus</p>
          </div>
        </div>

        <!-- Tombol Tutup di Mobile -->
        <button @click="isSidebarOpen = false" class="md:hidden text-gray-400 hover:text-white text-lg">
          ✕
        </button>
      </div>

      <!-- Navigasi Menu -->
      <nav class="flex-1 p-4 space-y-2 overflow-y-auto">
        <NuxtLink to="/admin" @click="isSidebarOpen = false" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-300 hover:bg-white/10 transition" exactActiveClass="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-md">
          <Icon name="ph:squares-four-bold" class="text-xl" />
          Dashboard
        </NuxtLink>
        <NuxtLink to="/jadwal" @click="isSidebarOpen = false" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-300 hover:bg-white/10 transition" exactActiveClass="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-md">
          <Icon name="ph:calendar-plus-bold" class="text-xl" />
          Jadwal Main
        </NuxtLink>
        <NuxtLink to="/verifikasi" @click="isSidebarOpen = false" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-300 hover:bg-white/10 transition" exactActiveClass="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-md">
          <Icon name="ph:receipt-bold" class="text-xl" />
          Verifikasi Bayar
        </NuxtLink>

        <!-- Menu Verifikasi Deposit -->
        <NuxtLink to="/verifikasi-deposit" @click="isSidebarOpen = false" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-400 hover:text-white transition-colors" exact-active-class="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-lg">
          <Icon name="ph:wallet-bold" class="text-xl" />
          Verifikasi Deposit
        </NuxtLink>

        <!-- Garis Pemisah -->
        <hr class="border-gray-800 my-4" />

        <NuxtLink to="/admin-keuangan" @click="isSidebarOpen = false" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-300 hover:bg-white/10 transition" exactActiveClass="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-md">
          <Icon name="ph:wallet-bold" class="text-xl" />
          Dashboard Kas
        </NuxtLink>

        <!-- Menu Profil Admin -->
        <NuxtLink to="/profil-admin" @click="isSidebarOpen = false" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-400 hover:text-white transition-colors" exact-active-class="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-lg">
          <Icon name="ph:user-circle-bold" class="text-xl" />
          Profil Saya
        </NuxtLink>

        <!-- Menu Master Notifikasi Broadcast -->
        <NuxtLink to="/master-notifikasi" @click="isSidebarOpen = false" class="flex items-center gap-3 px-4 py-3 rounded-xl text-gray-400 hover:text-white transition-colors" exact-active-class="bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold shadow-lg">
          <Icon name="ph:broadcast-bold" class="text-xl text-orange-400" />
          Master Notifikasi
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
      <header class="theme-bg-surface border-b theme-border h-16 flex items-center justify-between px-4 md:px-8 shadow-sm sticky top-0 z-10 transition-colors duration-300">
        <div class="flex items-center gap-3">
          <!-- Hamburger Button di Mobile -->
          <button @click="isSidebarOpen = !isSidebarOpen" class="md:hidden p-2 theme-text-muted hover:bg-orange-500/10 rounded-lg">
            <Icon name="ph:list-bold" class="text-2xl" />
          </button>
          <p class="text-sm font-bold theme-text-main">Selamat datang, Pengurus! 👋</p>
        </div>
        <div class="flex items-center gap-3">
          <!-- Theme Toggle Button -->
          <button 
            @click="toggleTheme" 
            :title="isDark ? 'Ganti ke Light Mode' : 'Ganti ke Dark Mode'"
            class="p-2.5 rounded-xl border theme-border theme-text-main hover:bg-orange-500/10 transition active:scale-95">
            <Icon :name="isDark ? 'ph:sun-bold' : 'ph:moon-bold'" class="text-xl text-amber-400" />
          </button>

          <div class="w-9 h-9 bg-gradient-to-tr from-red-600 to-orange-600 text-white rounded-full flex items-center justify-center font-bold shadow">
            A
          </div>
        </div>
      </header>

      <main class="flex-1 p-4 md:p-8 overflow-y-auto theme-bg-page transition-colors duration-300">
        <slot />
      </main>
    </div>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const isLoggedIn = ref(false)
const isSidebarOpen = ref(false)

const { isDark, toggleTheme } = useTheme()
const toast = useToast()

onMounted(() => {
  const status = localStorage.getItem('admin_logged_in')
  if (status === 'true') {
    isLoggedIn.value = true
  } else {
    navigateTo('/admin-login')
  }
})

const logoutAdmin = () => {
  toast.confirm({
    title: 'Keluar Panel Admin',
    message: 'Apakah Anda yakin ingin keluar dari Panel Pengurus?',
    confirmText: 'Ya, Keluar',
    cancelText: 'Batal',
    onConfirm: () => {
      localStorage.removeItem('admin_logged_in')
      toast.info('Keluar dari Panel Admin')
      navigateTo('/admin-login')
    }
  })
}
</script>